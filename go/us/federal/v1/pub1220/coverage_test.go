// Package pub1220_test checks each US federal tax form schema against IRS
// Publication 1220, which is the IRS's own list of what may be filed
// electronically for each information return.
//
// Publication 1220 gives every form an Amount Code table: "Total ordinary
// dividends" is Amount Code 1 on a 1099-DIV, "Qualified dividends" is 2. The
// test reads that table out of a checked-in fixture and asserts that every
// Amount Code the IRS lists is claimed by a concept in the form's
// mapping.textproto, or listed there as deliberately unmapped with a reason.
// docs/form-schemas.md requires this.
//
// # The fixture
//
// testdata/p1220-2025-amount-codes.txt is pages 57 to 70 of Publication 1220
// for Tax Year 2025, extracted with:
//
//	pdftotext -layout -f 57 -l 70 p1220.pdf p1220-2025-amount-codes.txt
//
// It is checked in rather than downloaded, so the suite is hermetic and a
// failure means a schema changed rather than that the IRS republished.
// Refreshing it is a deliberate commit, which is also how a new tax year gets
// added.
//
// # Coverage is necessary and not sufficient
//
// Publication 1220 lists what carries an amount. A form also prints checkboxes,
// codes, CUSIPs and state identifiers, and none of those has an Amount Code. A
// form at full coverage can still be missing fields.
package pub1220_test

import (
	"os"
	"path"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"testing"

	"github.com/bazelbuild/rules_go/go/runfiles"
	commonv1 "github.com/nexteggapp/taxprotos/go/common/v1"
	// Registering each form's descriptors is what lets the mapping's
	// proto_message be looked up by name.
	_ "github.com/nexteggapp/taxprotos/go/us/federal/v1/form_1098"
	_ "github.com/nexteggapp/taxprotos/go/us/federal/v1/form_1098_e"
	_ "github.com/nexteggapp/taxprotos/go/us/federal/v1/form_1098_t"
	_ "github.com/nexteggapp/taxprotos/go/us/federal/v1/form_1099_b"
	_ "github.com/nexteggapp/taxprotos/go/us/federal/v1/form_1099_c"
	_ "github.com/nexteggapp/taxprotos/go/us/federal/v1/form_1099_div"
	_ "github.com/nexteggapp/taxprotos/go/us/federal/v1/form_1099_g"
	_ "github.com/nexteggapp/taxprotos/go/us/federal/v1/form_1099_int"
	_ "github.com/nexteggapp/taxprotos/go/us/federal/v1/form_1099_k"
	_ "github.com/nexteggapp/taxprotos/go/us/federal/v1/form_1099_misc"
	_ "github.com/nexteggapp/taxprotos/go/us/federal/v1/form_1099_nec"
	_ "github.com/nexteggapp/taxprotos/go/us/federal/v1/form_1099_oid"
	_ "github.com/nexteggapp/taxprotos/go/us/federal/v1/form_1099_patr"
	_ "github.com/nexteggapp/taxprotos/go/us/federal/v1/form_1099_q"
	_ "github.com/nexteggapp/taxprotos/go/us/federal/v1/form_1099_r"
	_ "github.com/nexteggapp/taxprotos/go/us/federal/v1/form_1099_s"
	_ "github.com/nexteggapp/taxprotos/go/us/federal/v1/form_1099_sa"
	_ "github.com/nexteggapp/taxprotos/go/us/federal/v1/form_3921"
	_ "github.com/nexteggapp/taxprotos/go/us/federal/v1/form_3922"
	_ "github.com/nexteggapp/taxprotos/go/us/federal/v1/form_5498"
	_ "github.com/nexteggapp/taxprotos/go/us/federal/v1/form_5498_sa"
	_ "github.com/nexteggapp/taxprotos/go/us/federal/v1/form_ssa_1099"
	_ "github.com/nexteggapp/taxprotos/go/us/federal/v1/form_w2"
	_ "github.com/nexteggapp/taxprotos/go/us/federal/v1/form_w2g"
	"google.golang.org/protobuf/encoding/prototext"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
)

// runfilesPrefix is the workspace name every runfiles path starts with.
const runfilesPrefix = "nexteggapp_taxprotos"

// taxYear is the year the fixture covers. A second year means a second fixture
// and a second entry here.
const taxYear = 2025

const fixture = "go/us/federal/v1/pub1220/testdata/p1220-2025-amount-codes.txt"

// schema is one form's proto and the mapping beside it.
type schema struct {
	// form is the form as Publication 1220 names it, which is the key the
	// fixture is parsed under.
	form string

	// dir is the package directory holding form.proto and mapping.textproto.
	dir string

	// filedWithSsa says the form goes to the Social Security Administration
	// rather than the IRS, so Publication 1220 has no table for it.
	filedWithSsa bool

	// issuedBySsa says the Social Security Administration issues the form,
	// so it is not an information return Publication 1220 specifies.
	issuedBySsa bool
}

var schemas = []schema{
	{form: "W-2", dir: "form_w2", filedWithSsa: true},
	{form: "1099-INT", dir: "form_1099_int"},
	{form: "1099-DIV", dir: "form_1099_div"},
	{form: "1099-B", dir: "form_1099_b"},
	{form: "1099-R", dir: "form_1099_r"},
	{form: "1099-MISC", dir: "form_1099_misc"},
	{form: "1099-NEC", dir: "form_1099_nec"},
	{form: "1099-K", dir: "form_1099_k"},
	{form: "1098", dir: "form_1098"},
	{form: "1098-E", dir: "form_1098_e"},
	{form: "1098-T", dir: "form_1098_t"},
	{form: "1099-SA", dir: "form_1099_sa"},
	{form: "5498", dir: "form_5498"},
	{form: "5498-SA", dir: "form_5498_sa"},
	{form: "1099-G", dir: "form_1099_g"},
	{form: "W-2G", dir: "form_w2g"},
	{form: "SSA-1099", dir: "form_ssa_1099", issuedBySsa: true},
	{form: "1099-OID", dir: "form_1099_oid"},
	{form: "1099-PATR", dir: "form_1099_patr"},
	{form: "1099-S", dir: "form_1099_s"},
	{form: "1099-C", dir: "form_1099_c"},
	{form: "1099-Q", dir: "form_1099_q"},
	{form: "3921", dir: "form_3921"},
	{form: "3922", dir: "form_3922"},
}

// TestPub1220Coverage checks every form's mapping against the IRS's own list.
func TestPub1220Coverage(t *testing.T) {
	tables := parseAmountCodeTables(t)

	for _, s := range schemas {
		t.Run(s.form, func(t *testing.T) {
			mapping := loadMapping(t, s)
			message := messageDescriptor(t, mapping.GetMessageName())

			checkConceptsResolve(t, mapping, message)
			checkYearsComplete(t, mapping)
			checkNoDuplicateAmountCodes(t, mapping)

			if s.filedWithSsa || s.issuedBySsa {
				// Publication 1220 specifies electronic filing for the forms
				// the IRS receives. A W-2 is filed with the Social Security
				// Administration and specified separately, and an SSA-1099 is
				// issued by it, so there is no table to check either against
				// and each needs another oracle. Nothing checks their box
				// lists today.
				if len(mapping.GetUnmappedAmountCodes()) != 0 {
					t.Errorf("%s carries unmapped_amount_codes, but Publication 1220 has no table for it", s.form)
				}
				t.Logf("%s: no Publication 1220 table; %d concepts unchecked", s.form, len(mapping.GetConcepts()))
				return
			}

			table := tables[s.form]
			if len(table) == 0 {
				t.Fatalf("the fixture has no Amount Code table for Form %s; the forms it does have are %v", s.form, sortedKeys(tables))
			}
			checkAmountCodes(t, s.form, mapping, table)
		})
	}
}

// checkConceptsResolve asserts every concept names a field the message has.
func checkConceptsResolve(t *testing.T, mapping *commonv1.FormMapping, message protoreflect.MessageDescriptor) {
	t.Helper()
	for _, concept := range mapping.GetConcepts() {
		if !fieldExists(message, concept.GetConcept()) {
			t.Errorf("concept %q is not a field on %s", concept.GetConcept(), message.FullName())
		}
	}
}

// checkYearsComplete asserts every concept says something about every tax year
// the mapping covers, so that a gap is reported rather than read as absence.
func checkYearsComplete(t *testing.T, mapping *commonv1.FormMapping) {
	t.Helper()
	for _, concept := range mapping.GetConcepts() {
		for _, year := range mapping.GetTaxYears() {
			if !hasYear(concept, year) {
				t.Errorf("concept %q says nothing about tax year %d; set absent to say the form did not report it", concept.GetConcept(), year)
			}
		}
	}
}

// checkNoDuplicateAmountCodes asserts two concepts never claim one code.
func checkNoDuplicateAmountCodes(t *testing.T, mapping *commonv1.FormMapping) {
	t.Helper()
	claimed := map[string]string{}
	for _, concept := range mapping.GetConcepts() {
		for _, year := range concept.GetYears() {
			code := year.GetAmountCode()
			if code == "" {
				continue
			}
			key := code + "/" + strconv.Itoa(int(year.GetTaxYear()))
			if first, ok := claimed[key]; ok {
				t.Errorf("amount code %s for tax year %d is claimed by both %q and %q", code, year.GetTaxYear(), first, concept.GetConcept())
				continue
			}
			claimed[key] = concept.GetConcept()
		}
	}
}

// checkAmountCodes compares the mapping against the IRS's table both ways: no
// concept may claim a code the IRS does not list, and no code the IRS lists may
// go unaccounted for.
func checkAmountCodes(t *testing.T, form string, mapping *commonv1.FormMapping, table map[string]string) {
	t.Helper()
	claimed := claimedCodes(mapping)

	for code := range claimed {
		if _, ok := table[code]; !ok {
			t.Errorf("%s maps amount code %s, which Publication 1220 does not list for this form", form, code)
		}
	}

	excused := map[string]bool{}
	for _, unmapped := range mapping.GetUnmappedAmountCodes() {
		if unmapped.GetTaxYear() != taxYear {
			continue
		}
		if unmapped.GetReason() == "" {
			t.Errorf("%s excuses amount code %s with no reason", form, unmapped.GetAmountCode())
		}
		excused[unmapped.GetAmountCode()] = true
	}

	var missing []string
	for _, code := range sortedKeys(table) {
		if claimed[code] != "" || excused[code] {
			continue
		}
		missing = append(missing, code+" ("+table[code]+")")
	}
	if len(missing) > 0 {
		t.Errorf("%s maps %d of %d amount types. Nothing maps to: %s",
			form, len(claimed), len(table), strings.Join(missing, "; "))
	}
	t.Logf("%s: %d of %d amount types mapped, %d excused", form, len(claimed), len(table), len(excused))
}

// claimedCodes is every amount code the mapping claims for the fixture's year.
func claimedCodes(mapping *commonv1.FormMapping) map[string]string {
	claimed := map[string]string{}
	for _, concept := range mapping.GetConcepts() {
		for _, year := range concept.GetYears() {
			if year.GetTaxYear() == taxYear && year.GetAmountCode() != "" {
				claimed[year.GetAmountCode()] = concept.GetConcept()
			}
		}
	}
	return claimed
}

var (
	// sectionRe opens a form's table: "For Reporting Payments on Form 1099-DIV:".
	sectionRe = regexp.MustCompile(`For Reporting (?:Payments|Distributions|Information) on Form ([A-Z0-9-]+):`)

	// rowRe is one row of a table: an amount type, a run of spaces, and a
	// single-character amount code at the end of the line.
	rowRe = regexp.MustCompile(`^\s*(\S.*?)\s{2,}([0-9A-Z])\s*$`)

	// columnGapRe splits the left-hand notes column from the amount type,
	// which pdftotext runs together on the pages that have both.
	columnGapRe = regexp.MustCompile(`\s{3,}`)
)

// parseAmountCodeTables reads the fixture into form name -> amount code ->
// amount type.
func parseAmountCodeTables(t *testing.T) map[string]map[string]string {
	t.Helper()
	resolved, err := runfiles.Rlocation(path.Join(runfilesPrefix, fixture))
	if err != nil {
		t.Fatalf("resolve %s from runfiles: %v", fixture, err)
	}
	raw, err := os.ReadFile(resolved)
	if err != nil {
		t.Fatalf("read %s: %v", resolved, err)
	}

	tables := map[string]map[string]string{}
	form := ""
	for _, line := range strings.Split(string(raw), "\n") {
		if match := sectionRe.FindStringSubmatch(line); match != nil {
			form = match[1]
			if _, ok := tables[form]; !ok {
				tables[form] = map[string]string{}
			}
			continue
		}
		if form == "" {
			continue
		}
		code, amountType, ok := parseRow(line)
		if !ok {
			continue
		}
		if _, seen := tables[form][code]; !seen {
			tables[form][code] = amountType
		}
	}
	if len(tables) == 0 {
		t.Fatalf("%s produced no Amount Code tables", resolved)
	}
	return tables
}

// parseRow reads one table row, or reports that the line is not one.
func parseRow(line string) (code, amountType string, ok bool) {
	match := rowRe.FindStringSubmatch(line)
	if match == nil {
		return "", "", false
	}
	columns := columnGapRe.Split(strings.TrimSpace(match[1]), -1)
	amountType = strings.TrimSpace(columns[len(columns)-1])
	if len(amountType) < 4 || strings.HasPrefix(amountType, "Note") {
		return "", "", false
	}
	return match[2], amountType, true
}

// loadMapping parses one form's mapping.textproto.
func loadMapping(t *testing.T, s schema) *commonv1.FormMapping {
	t.Helper()
	rel := path.Join("us/federal/v1", s.dir, "mapping.textproto")
	resolved, err := runfiles.Rlocation(path.Join(runfilesPrefix, rel))
	if err != nil {
		t.Fatalf("resolve %s from runfiles: %v", rel, err)
	}
	raw, err := os.ReadFile(resolved)
	if err != nil {
		t.Fatalf("read %s: %v", resolved, err)
	}
	mapping := &commonv1.FormMapping{}
	if err := prototext.Unmarshal(raw, mapping); err != nil {
		t.Fatalf("parse %s: %v", rel, err)
	}
	if mapping.GetForm() != s.form {
		t.Fatalf("%s says it is for Form %q, want %q", rel, mapping.GetForm(), s.form)
	}
	return mapping
}

// messageDescriptor looks up a message by the name the mapping gives.
func messageDescriptor(t *testing.T, name string) protoreflect.MessageDescriptor {
	t.Helper()
	found, err := protoregistry.GlobalFiles.FindDescriptorByName(protoreflect.FullName(name))
	if err != nil {
		t.Fatalf("look up %s: %v", name, err)
	}
	message, ok := found.(protoreflect.MessageDescriptor)
	if !ok {
		t.Fatalf("%s is a %T, not a message", name, found)
	}
	return message
}

// fieldExists walks a dotted concept path through nested messages.
func fieldExists(message protoreflect.MessageDescriptor, dotted string) bool {
	for _, part := range strings.Split(dotted, ".") {
		field := message.Fields().ByName(protoreflect.Name(part))
		if field == nil {
			return false
		}
		if field.Message() == nil {
			return true
		}
		message = field.Message()
	}
	return true
}

// hasYear reports whether a concept says anything about one tax year.
func hasYear(concept *commonv1.Concept, taxYear int32) bool {
	for _, year := range concept.GetYears() {
		if year.GetTaxYear() == taxYear {
			return true
		}
	}
	return false
}

// sortedKeys keeps failure messages stable between runs.
func sortedKeys[V any](m map[string]V) []string {
	keys := make([]string, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}
