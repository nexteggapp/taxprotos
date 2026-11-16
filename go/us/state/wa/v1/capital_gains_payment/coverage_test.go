// Coverage of the Washington capital gains tax payment against the statute.
//
// Washington prints no form for the payment: the return is filed
// electronically and the tax is one payment, so there are no line numbers to
// check. As for the capital gains return beside it, the oracle is the statute.
// Each concept's mapping cites the sections that set it, and this test
// requires every cited section to be in the extracts checked in under
// testdata: chapter 82.87 RCW, RCW 82.32.050 and 82.32.090, and RCW 1.12.070.
package capitalgainspayment_test

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
	_ "github.com/nexteggapp/taxprotos/go/us/state/wa/v1/capital_gains_payment"
	"google.golang.org/protobuf/encoding/prototext"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
)

// runfilesPrefix is the repository name every runfiles path starts with. It is
// this module's own name rather than _main, so the path resolves both when this
// module is built on its own and when another module builds it.
const runfilesPrefix = "nexteggapp_taxprotos"

const dir = "us/state/wa/v1/capital_gains_payment"

// mappings are the payment's own concepts, a step of the late payment
// penalty, and an interest period.
var mappings = []string{"mapping.textproto", "mapping_steps.textproto", "mapping_interest.textproto"}

// extracts are the statute texts the citations must be found in.
var extracts = []string{"rcw-82-87-extract.txt", "rcw-82-32-050.txt", "rcw-82-32-090.txt", "rcw-1-12-070.txt"}

// citation matches an RCW section, with or without a subsection.
var citation = regexp.MustCompile(`RCW (\d+\.\d+\.\d+)`)

// TestConceptsResolve checks that every concept a mapping names is a field on
// the message it says it maps, which is what catches a rename.
func TestConceptsResolve(t *testing.T) {
	for _, name := range mappings {
		t.Run(name, func(t *testing.T) {
			m := readMapping(t, name)
			md := messageDescriptor(t, m.GetMessageName())
			for _, c := range m.GetConcepts() {
				if md.Fields().ByName(protoreflect.Name(c.GetConcept())) == nil {
					t.Errorf("concept %q is not a field on %s", c.GetConcept(), m.GetMessageName())
				}
			}
		})
	}
}

// TestEveryConceptCitesTheStatute checks that each concept cites at least one
// RCW section, that every section it cites is in the checked-in extracts, and
// that no concept claims a line number, since there is no form to have one.
func TestEveryConceptCitesTheStatute(t *testing.T) {
	statute := readStatute(t)
	for _, name := range mappings {
		t.Run(name, func(t *testing.T) {
			m := readMapping(t, name)
			for _, c := range m.GetConcepts() {
				for _, y := range c.GetYears() {
					if y.GetAbsent() {
						continue
					}
					if y.GetIrsBox() != "" {
						t.Errorf("concept %q carries a line number %q; the payment has no printed form", c.GetConcept(), y.GetIrsBox())
					}
					found := citation.FindAllStringSubmatch(y.GetNote(), -1)
					if len(found) == 0 {
						t.Errorf("concept %q cites no RCW section", c.GetConcept())
						continue
					}
					for _, f := range found {
						if !strings.Contains(statute, "RCW "+f[1]) {
							t.Errorf("concept %q cites %s, which is not in the checked-in extracts", c.GetConcept(), f[1])
						}
					}
				}
			}
		})
	}
}

// TestEveryMoneyFieldIsMapped reports fields carrying an amount that no mapping
// accounts for: a field nobody can trace to the rule that sets it.
func TestEveryMoneyFieldIsMapped(t *testing.T) {
	for _, name := range mappings {
		t.Run(name, func(t *testing.T) {
			m := readMapping(t, name)
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

// TestTheStatuteStillSaysWhatTheSchemaAssumes pins the sentences the shape of
// this message depends on.
func TestTheStatuteStillSaysWhatTheSchemaAssumes(t *testing.T) {
	statute := readStatute(t)
	for _, tc := range []struct {
		says  string
		about string
	}{
		{"on or before the date fixed for the filing of the return, regardless of any filing extension", "one payment by the due date, whatever extension"},
		{"may submit payment of the tax beginning six months before the due date", "the first day a payment can be made"},
		{"there is assessed a penalty of\nnine percent of the amount of the tax", "the first step of the late payment penalty"},
		{"No penalty so added may be less than $5.", "the $5 minimum"},
	} {
		if !strings.Contains(statute, tc.says) {
			t.Errorf("the statute no longer states %s\n  wanted: %q", tc.about, tc.says)
		}
	}
}

func readMapping(t *testing.T, name string) *commonv1.FormMapping {
	t.Helper()
	m := &commonv1.FormMapping{}
	if err := prototext.Unmarshal(readRunfile(t, path.Join(dir, name)), m); err != nil {
		t.Fatalf("parse %s: %v", name, err)
	}
	return m
}

func readStatute(t *testing.T) string {
	t.Helper()
	var b strings.Builder
	for _, name := range extracts {
		b.Write(readRunfile(t, path.Join(dir, "testdata", name)))
		b.WriteString("\n")
	}
	return b.String()
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
