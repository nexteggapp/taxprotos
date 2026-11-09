// Coverage of Form D-400 against the printed form.
//
// The oracle is the same one the federal 1040 uses, for the same reason.
// Publication 1220 enumerates the fields the IRS expects electronically for the
// information returns it receives; a state return is not one of those, so that
// oracle is unavailable and the printed form is what is left.
//
// testdata holds the text of the PDF the North Carolina Department of Revenue publishes. It is weaker
// than an enumerated field list — a page of text cannot tell us a line was
// misunderstood, only that it exists — and much stronger than trusting
// anybody's recollection of a tax form.
package formd400_test

import (
	"fmt"
	"os"
	"path"
	"regexp"
	"strings"
	"testing"

	"github.com/bazelbuild/rules_go/go/runfiles"
	commonv1 "github.com/nexteggapp/taxprotos/go/common/v1"
	// Registers the message the mapping names.
	_ "github.com/nexteggapp/taxprotos/go/us/state/nc/v1/form_d400"
	"google.golang.org/protobuf/encoding/prototext"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
)

// runfilesPrefix is the repository name every runfiles path starts with. It is
// this module's own name rather than _main, so the path resolves both when this
// module is built on its own and when another module builds it.
const runfilesPrefix = "nexteggapp_taxprotos"

const (
	dir     = "us/state/nc/v1/form_d400"
	mapping = "mapping.textproto"
	extract = "d400-2025.txt"
)

// TestConceptsResolve checks that every concept the mapping names is a field
// on the message it says it maps, which is what catches a rename.
func TestConceptsResolve(t *testing.T) {
	m := readMapping(t)
	md := messageDescriptor(t, m.GetMessageName())

	for _, c := range m.GetConcepts() {
		if md.Fields().ByName(protoreflect.Name(c.GetConcept())) == nil {
			t.Errorf("concept %q is not a field on %s", c.GetConcept(), m.GetMessageName())
		}
	}
}

// TestLinesExistOnTheForm checks that every line number the mapping names
// appears on the printed form, so that a renumbering is noticed when the
// extract is refreshed.
func TestLinesExistOnTheForm(t *testing.T) {
	m := readMapping(t)
	text := string(readRunfile(t, path.Join(dir, "testdata", extract)))

	for _, c := range m.GetConcepts() {
		for _, y := range c.GetYears() {
			line := y.GetIrsBox()
			if line == "" || y.GetAbsent() {
				continue
			}
			if !lineAppears(text, line) {
				t.Errorf("concept %q maps to line %s, which does not appear in %s", c.GetConcept(), line, extract)
			}
		}
	}
}

// TestEveryMoneyFieldIsMapped reports fields carrying an amount that the
// mapping does not account for: a field nobody can trace to a line.
func TestEveryMoneyFieldIsMapped(t *testing.T) {
	m := readMapping(t)
	md := messageDescriptor(t, m.GetMessageName())

	mapped := map[string]bool{}
	for _, c := range m.GetConcepts() {
		mapped[c.GetConcept()] = true
	}

	var missing []string
	for i := 0; i < md.Fields().Len(); i++ {
		fd := md.Fields().Get(i)
		if fd.Kind() != protoreflect.MessageKind || fd.IsList() {
			continue
		}
		if fd.Message().FullName() != "nexteggapp.taxprotos.common.v1.Money" {
			continue
		}
		if !mapped[string(fd.Name())] {
			missing = append(missing, string(fd.Name()))
		}
	}
	if len(missing) > 0 {
		t.Errorf("%s has %d money field(s) the mapping does not account for: %s",
			m.GetMessageName(), len(missing), strings.Join(missing, ", "))
	}
}

// TestSharedLinesCarryANote checks that a concept sharing a line number with
// another says which figure it is.
func TestSharedLinesCarryANote(t *testing.T) {
	m := readMapping(t)

	claims := map[string][]*commonv1.Concept{}
	for _, c := range m.GetConcepts() {
		for _, y := range c.GetYears() {
			if line := y.GetIrsBox(); line != "" {
				claims[line] = append(claims[line], c)
			}
		}
	}

	for line, sharing := range claims {
		if len(sharing) < 2 {
			continue
		}
		for _, c := range sharing {
			for _, y := range c.GetYears() {
				if y.GetIrsBox() == line && y.GetNote() == "" {
					t.Errorf("line %s is claimed by %d concepts and %q says nothing about which figure it is",
						line, len(sharing), c.GetConcept())
				}
			}
		}
	}
}

// lineAppears reports whether the extract prints the given line number, set
// off by whitespace as a form sets it. A single trailing full stop is allowed,
// since some forms print "7." where others print "7".
func lineAppears(text, line string) bool {
	q := regexp.QuoteMeta(line)
	re := regexp.MustCompile(`(?m)(^|\s)` + q + `\.?(\s|$)`)
	return re.MatchString(text)
}

func readMapping(t *testing.T) *commonv1.FormMapping {
	t.Helper()

	m := &commonv1.FormMapping{}
	if err := prototext.Unmarshal(readRunfile(t, path.Join(dir, mapping)), m); err != nil {
		t.Fatalf("parse %s: %v", mapping, err)
	}
	return m
}

func readRunfile(t *testing.T, rel string) []byte {
	t.Helper()

	resolved, err := runfiles.Rlocation(path.Join(runfilesPrefix, rel))
	if err != nil {
		t.Fatalf("resolve %s from runfiles: %v", rel, err)
	}
	raw, err := os.ReadFile(resolved)
	if err != nil {
		t.Fatalf("read %s: %v", rel, err)
	}
	return raw
}

func messageDescriptor(t *testing.T, name string) protoreflect.MessageDescriptor {
	t.Helper()

	found, err := protoregistry.GlobalFiles.FindDescriptorByName(protoreflect.FullName(name))
	if err != nil {
		t.Fatalf("find %s: %v", name, err)
	}
	md, ok := found.(protoreflect.MessageDescriptor)
	if !ok {
		t.Fatal(fmt.Sprintf("%s is not a message", name))
	}
	return md
}
