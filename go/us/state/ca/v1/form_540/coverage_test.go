// Coverage of Form 540 and Schedule CA (540) against the printed forms.
//
// The oracle is the same one the federal 1040 uses, for the same reason.
// Publication 1220 enumerates the fields the IRS expects electronically for the
// information returns it receives; a state return is not one of those, so that
// oracle is unavailable and the printed form is what is left.
//
// testdata holds pdftotext -layout output over the PDFs the Franchise Tax Board
// publishes. It is weaker than an enumerated field list — a page of text cannot
// tell us a line was misunderstood, only that it exists — and much stronger than
// trusting anybody's recollection of a tax form. Recollection would have put
// California's behavioural health surcharge under its old name and its
// standard deduction near the federal figure, and both would have been wrong.
package form540_test

import (
	"fmt"
	"os"
	"path"
	"regexp"
	"strings"
	"testing"

	"github.com/bazelbuild/rules_go/go/runfiles"
	commonv1 "github.com/nexteggapp/taxprotos/go/common/v1"
	// Registers the messages the mappings name.
	_ "github.com/nexteggapp/taxprotos/go/us/state/ca/v1/form_540"
	"google.golang.org/protobuf/encoding/prototext"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
)

// runfilesPrefix is the repository name every runfiles path starts with. It is
// this module's own name rather than _main, so the path resolves both when this
// module is built on its own and when another module builds it.
const runfilesPrefix = "nexteggapp_taxprotos"

const dir = "us/state/ca/v1/form_540"

// forms pairs each mapping with the extract its line numbers must appear in.
var forms = []struct {
	mapping string
	extract string
}{
	{"mapping_form540.textproto", "f540-2025.txt"},
	{"mapping_schedule_ca.textproto", "f540ca-2025.txt"},
}

// TestConceptsResolve checks that every concept a mapping names is a field on
// the message it says it maps.
//
// This is the check that catches a rename. A field renamed in the .proto with
// the mapping left alone leaves a mapping entry pointing at nothing, and
// nothing else would notice.
func TestConceptsResolve(t *testing.T) {
	for _, f := range forms {
		t.Run(f.mapping, func(t *testing.T) {
			m := readMapping(t, f.mapping)
			md := messageDescriptor(t, m.GetMessageName())

			for _, c := range m.GetConcepts() {
				name := protoreflect.Name(c.GetConcept())
				if md.Fields().ByName(name) == nil {
					t.Errorf("concept %q is not a field on %s", c.GetConcept(), m.GetMessageName())
				}
			}
		})
	}
}

// TestLinesExistOnTheForm checks that every line number a mapping names appears
// on the printed form.
//
// A line number is a claim about a document somebody else publishes, and it is
// the kind of claim that goes stale quietly. California renumbers, and a
// mapping that was right last year now points at a different line. This is what
// makes refreshing the extract a deliberate commit rather than something nobody
// notices until a return is wrong.
func TestLinesExistOnTheForm(t *testing.T) {
	for _, f := range forms {
		t.Run(f.mapping, func(t *testing.T) {
			m := readMapping(t, f.mapping)
			text := readExtract(t, f.extract)

			for _, c := range m.GetConcepts() {
				for _, y := range c.GetYears() {
					line := y.GetIrsBox()
					if line == "" || y.GetAbsent() {
						continue
					}
					if !lineAppears(text, line) {
						t.Errorf("concept %q maps to line %s, which does not appear in %s",
							c.GetConcept(), line, f.extract)
					}
				}
			}
		})
	}
}

// TestEveryMoneyFieldIsMapped reports fields carrying an amount that no mapping
// accounts for.
//
// Coverage in the other direction: the two tests above check that what is
// mapped is real, and this one checks that what is real is mapped. A field
// added to the .proto and forgotten in the mapping is a field nobody can trace
// back to a line on the form.
func TestEveryMoneyFieldIsMapped(t *testing.T) {
	for _, f := range forms {
		t.Run(f.mapping, func(t *testing.T) {
			m := readMapping(t, f.mapping)
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
				t.Errorf("%s has %d money field(s) no mapping accounts for: %s",
					m.GetMessageName(), len(missing), strings.Join(missing, ", "))
			}
		})
	}
}

// TestSharedLinesCarryANote checks that a concept sharing a line number with
// another says which figure it is.
//
// Schedule CA prints three columns against one line number, so line 27 alone
// identifies nothing. Where two concepts claim one line, the note is the only
// thing telling them apart, and a mapping that omits it is unusable rather than
// merely terse.
func TestSharedLinesCarryANote(t *testing.T) {
	for _, f := range forms {
		t.Run(f.mapping, func(t *testing.T) {
			m := readMapping(t, f.mapping)

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
		})
	}
}

// lineAppears reports whether the extract prints the given line number.
//
// The form sets a line number off from its text, either at the start of a line
// or surrounded by whitespace, so both are accepted. Anchoring on the number
// alone would match any figure that happened to equal it.
func lineAppears(text, line string) bool {
	q := regexp.QuoteMeta(line)
	re := regexp.MustCompile(`(?m)(^|\s)` + q + `(\s|$)`)
	return re.MatchString(text)
}

func readMapping(t *testing.T, name string) *commonv1.FormMapping {
	t.Helper()

	raw := readRunfile(t, path.Join(dir, name))
	m := &commonv1.FormMapping{}
	if err := prototext.Unmarshal(raw, m); err != nil {
		t.Fatalf("parse %s: %v", name, err)
	}
	return m
}

func readExtract(t *testing.T, name string) string {
	t.Helper()
	return string(readRunfile(t, path.Join(dir, "testdata", name)))
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
