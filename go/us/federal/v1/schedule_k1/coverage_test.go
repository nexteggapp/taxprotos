// Coverage of Schedule K-1 (Form 1065, Form 1120-S, Form 1041) against the printed forms.
package schedulek1_test

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
	_ "github.com/nexteggapp/taxprotos/go/us/federal/v1/schedule_k1"
	"google.golang.org/protobuf/encoding/prototext"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
)

// runfilesPrefix is the repository name every runfiles path starts with.
const runfilesPrefix = "nexteggapp_taxprotos"

const dir = "us/federal/v1/schedule_k1"

// forms pairs each mapping with the extract its line numbers must appear in.
var forms = []struct {
	mapping string
	extract string
}{
	{"mapping_1065.textproto", "f1065sk1-2025.txt"},
	{"mapping_1120s.textproto", "f1120ssk-2025.txt"},
	{"mapping_1041.textproto", "f1041sk1-2025.txt"},
}

// TestConceptsResolve checks that every concept a mapping names is a field on
// the message it says it maps.
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

func lineAppears(text, line string) bool {
	q := regexp.QuoteMeta(line)
	if regexp.MustCompile(`(?m)(^|\s)` + q + `(\s|$)`).MatchString(text) {
		return true
	}
	parts := strings.FieldsFunc(line, func(r rune) bool {
		return r == '-' || r == '–' || r == ','
	})
	if len(parts) > 1 {
		for _, p := range parts {
			p = strings.TrimSpace(p)
			if p != "" && !lineAppears(text, p) {
				return false
			}
		}
		return true
	}
	return false
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
