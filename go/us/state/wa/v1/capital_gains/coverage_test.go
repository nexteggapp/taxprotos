// Coverage of the Washington capital gains return against the statute.
//
// Every other coverage test in this tree checks line numbers against a printed
// form. This one cannot: RCW 82.87.110 requires the return to be filed
// electronically and Washington prints no form, so there are no line numbers to
// check and no page to check them on.
//
// The oracle here is chapter 82.87 RCW itself, checked in under testdata. Each
// concept's mapping cites the subsection that creates it, and this test
// requires every cited section to exist in the statute.
//
// That is a different check from the federal and Californian ones, and a
// stronger one. A form extract shows that a line exists; it says nothing about
// what belongs on it. A statute enumerates what is taxed, what is exempt and
// what is deductible, so a citation that has drifted to the wrong subsection
// fails here rather than looking plausible forever.
package capitalgains_test

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
	_ "github.com/nexteggapp/taxprotos/go/us/state/wa/v1/capital_gains"
	"google.golang.org/protobuf/encoding/prototext"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
)

// runfilesPrefix is the repository name every runfiles path starts with. It is
// this module's own name rather than _main, so the path resolves both when this
// module is built on its own and when another module builds it.
const runfilesPrefix = "nexteggapp_taxprotos"

const dir = "us/state/wa/v1/capital_gains"

const mappingFile = "mapping.textproto"

// citation matches a chapter 82.87 RCW section, with or without a subsection.
var citation = regexp.MustCompile(`RCW (82\.87\.\d+)`)

// TestConceptsResolve checks that every concept the mapping names is a field on
// the message it says it maps.
func TestConceptsResolve(t *testing.T) {
	m := readMapping(t)
	md := messageDescriptor(t, m.GetMessageName())

	for _, c := range m.GetConcepts() {
		name := protoreflect.Name(c.GetConcept())
		if md.Fields().ByName(name) == nil {
			t.Errorf("concept %q is not a field on %s", c.GetConcept(), m.GetMessageName())
		}
	}
}

// TestEveryConceptCitesTheStatute checks that each concept names a subsection of
// chapter 82.87 RCW, and that the section it names is really in the chapter.
//
// This is what replaces the line-number check. A citation is a claim about a
// document, and it is exactly the kind of claim that nothing else would catch:
// the calculator beside this schema shipped three wrong ones — the standard
// deduction cited to the definitions section, the base rate to a subsection one
// level too shallow, and the surcharge to a subsection about state government
// funding — all of which read plausibly and all of which would have sent a
// reader to the wrong place.
func TestEveryConceptCitesTheStatute(t *testing.T) {
	m := readMapping(t)
	statute := readExtract(t)

	for _, c := range m.GetConcepts() {
		for _, y := range c.GetYears() {
			if y.GetAbsent() {
				continue
			}

			// Washington has no line numbers, so a mapping claiming one is a
			// mapping copied from another state without being read.
			if y.GetIrsBox() != "" {
				t.Errorf("concept %q carries a line number %q; the Washington return is filed electronically and has none",
					c.GetConcept(), y.GetIrsBox())
			}

			found := citation.FindAllStringSubmatch(y.GetNote(), -1)
			if len(found) == 0 {
				t.Errorf("concept %q cites no subsection of chapter 82.87 RCW", c.GetConcept())
				continue
			}
			for _, f := range found {
				if !strings.Contains(statute, "RCW "+f[1]) {
					t.Errorf("concept %q cites %s, which is not in the checked-in chapter",
						c.GetConcept(), f[1])
				}
			}
		}
	}
}

// TestEveryMoneyFieldIsMapped reports fields carrying an amount that the
// mapping does not account for.
//
// Coverage in the other direction: the tests above check that what is mapped is
// real, and this one checks that what is real is mapped. A field added to the
// .proto and forgotten in the mapping is a field nobody can trace back to the
// rule that creates it.
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
		t.Errorf("%s has %d money field(s) no mapping accounts for: %s",
			m.GetMessageName(), len(missing), strings.Join(missing, ", "))
	}
}

// TestTheStatuteStillSaysWhatTheSchemaAssumes pins the handful of sentences the
// shape of this message depends on.
//
// A schema is a set of claims about how a tax works, and these are the ones
// that would silently become wrong if the legislature moved them.
func TestTheStatuteStillSaysWhatTheSchemaAssumes(t *testing.T) {
	statute := readExtract(t)

	for _, tc := range []struct {
		says  string
		about string
	}{
		{
			`"Washington capital gains" means an individual's adjusted capital gain, as modified in RCW  82.87.060`,
			"the deductions being inside the defined term, which is why both rates apply after them",
		},
		{
			"2.90 percent multiplied by the portion of an individual's Washington capital gains exceeding $1,000,000",
			"the surcharge, and that it is measured on the post-deduction figure",
		},
		{
			"seven percent multiplied by an individual's Washington capital gains",
			"the base rate applying to all taxable gain rather than to a band",
		},
		{
			"their combined standard deduction is limited to $250,000, regardless of whether they file joint or separate returns",
			"one deduction per couple, which is why the standard deduction does not vary by filing status",
		},
	} {
		if !strings.Contains(statute, tc.says) {
			t.Errorf("the statute no longer states %s\n  wanted: %q", tc.about, tc.says)
		}
	}
}

func readMapping(t *testing.T) *commonv1.FormMapping {
	t.Helper()

	raw := readRunfile(t, path.Join(dir, mappingFile))
	m := &commonv1.FormMapping{}
	if err := prototext.Unmarshal(raw, m); err != nil {
		t.Fatalf("parse %s: %v", mappingFile, err)
	}
	return m
}

func readExtract(t *testing.T) string {
	t.Helper()
	return string(readRunfile(t, path.Join(dir, "testdata", "rcw-82-87-extract.txt")))
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
