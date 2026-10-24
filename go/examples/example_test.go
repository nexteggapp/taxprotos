// Package examples is the worked example the README quotes.
//
// It is a test so that the README cannot drift from what these schemas
// actually do: the output below is checked on every run.
package examples_test

import (
	"bytes"
	"encoding/json"
	"fmt"

	commonv1 "github.com/nexteggapp/taxprotos/go/common/v1"
	formw2 "github.com/nexteggapp/taxprotos/go/us/federal/v1/form_w2"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// usd returns a Money for a whole number of dollars.
//
// Money is an integer count of minor units, so a dollar is a hundred of them.
// There is no float anywhere in this, on purpose.
func usd(dollars int64) *commonv1.Money {
	return &commonv1.Money{Minor: dollars * 100, Currency: "USD"}
}

// Example reads a W-2 the way something that had just parsed one would.
func Example() {
	w2 := &formw2.FormW2{
		TaxYear: 2025,
		Employer: &commonv1.Party{
			Name: "Acme Corporation",
			Tin:  "12-3456789",
		},
		Employee: &commonv1.Party{
			Name: "Dana Okonkwo",
			// A W-2 prints the employee's Social Security number in full,
			// unlike a 1099, which masks all but the last four digits.
			Tin: "123-45-6789",
		},
		WagesTipsOtherCompensation: usd(120_000),
		FederalTaxWithheld:         usd(18_400),
		SocialSecurityWages:        usd(120_000),
		SocialSecurityTaxWithheld:  usd(7_440),
		MedicareWagesAndTips:       usd(120_000),
		MedicareTaxWithheld:        usd(1_740),

		// Box 12 is a repeated code-and-amount pair rather than a field per
		// code, because the codes run past thirty and the form has room for
		// four.
		CodedAmounts: []*formw2.FormW2_CodedAmount{
			{Code: "D", Amount: usd(9_000)},   // 401(k) deferral
			{Code: "DD", Amount: usd(14_200)}, // employer-paid health cover
		},
		StateTaxWithholding: []*commonv1.StateTaxWithholding{
			{State: "CA", StateId: "123-4567-8", TaxWithheld: usd(6_100), StateIncome: usd(120_000)},
		},
	}

	fmt.Printf("%s paid %s\n", w2.GetEmployer().GetName(), w2.GetEmployee().GetName())
	fmt.Printf("box 1 wages    %s\n", dollars(w2.GetWagesTipsOtherCompensation()))
	fmt.Printf("box 2 withheld %s\n", dollars(w2.GetFederalTaxWithheld()))

	for _, c := range w2.GetCodedAmounts() {
		fmt.Printf("box 12 code %-2s %s\n", c.GetCode(), dollars(c.GetAmount()))
	}
	for _, s := range w2.GetStateTaxWithholding() {
		fmt.Printf("%s withheld    %s\n", s.GetState(), dollars(s.GetTaxWithheld()))
	}

	// It is an ordinary protobuf message, so it serialises like one. The
	// round trip is shown rather than the byte count: the count is a fact
	// about today's field set and would change for reasons that teach nobody
	// anything.
	wire, err := proto.Marshal(w2)
	if err != nil {
		panic(err)
	}
	var back formw2.FormW2
	if err := proto.Unmarshal(wire, &back); err != nil {
		panic(err)
	}
	fmt.Printf("\nround trip preserved box 1: %v\n",
		back.GetWagesTipsOtherCompensation().GetMinor() == w2.GetWagesTipsOtherCompensation().GetMinor())

	// Output:
	// Acme Corporation paid Dana Okonkwo
	// box 1 wages    $120000.00
	// box 2 withheld $18400.00
	// box 12 code D  $9000.00
	// box 12 code DD $14200.00
	// CA withheld    $6100.00
	//
	// round trip preserved box 1: true
}

// dollars renders a Money. Formatting for a screen is the caller's business, so
// the schemas ship no such helper and this is the whole of what one takes.
func dollars(m *commonv1.Money) string {
	minor := m.GetMinor()
	sign := ""
	if minor < 0 {
		sign, minor = "-", -minor
	}
	return fmt.Sprintf("%s$%d.%02d", sign, minor/100, minor%100)
}

// ExampleFormW2_json is the JSON the README shows.
//
// A README that prints a payload and is never run drifts from the schema
// quietly, and a reader has no way to tell. This renders the real message with
// the canonical proto3 JSON mapping, so what the README claims is what the
// schema produces.
//
// Note that Money.minor is a string. Proto3 JSON encodes 64-bit integers as
// strings, because a JSON number is a double and would lose cents above about
// $90 trillion. That is the same reasoning the type itself rests on.
func ExampleFormW2_json() {
	w2 := &formw2.FormW2{
		TaxYear:                    2025,
		Employer:                   &commonv1.Party{Name: "Acme Corporation", Tin: "12-3456789"},
		WagesTipsOtherCompensation: usd(120_000),
		FederalTaxWithheld:         usd(18_400),
		CodedAmounts: []*formw2.FormW2_CodedAmount{
			{Code: "D", Amount: usd(9_000)},
		},
	}

	out, err := protojson.Marshal(w2)
	if err != nil {
		panic(err)
	}

	// Re-indented through encoding/json rather than printed as protojson
	// returns it. protojson varies its whitespace between runs on purpose, to
	// discourage exactly the byte comparison this example needs. Passing the
	// bytes through json.Indent normalises the layout without touching field
	// order or any value.
	var pretty bytes.Buffer
	if err := json.Indent(&pretty, out, "", "  "); err != nil {
		panic(err)
	}
	fmt.Println(pretty.String())

	// Output:
	// {
	//   "taxYear": 2025,
	//   "employer": {
	//     "name": "Acme Corporation",
	//     "tin": "12-3456789"
	//   },
	//   "wagesTipsOtherCompensation": {
	//     "minor": "12000000",
	//     "currency": "USD"
	//   },
	//   "federalTaxWithheld": {
	//     "minor": "1840000",
	//     "currency": "USD"
	//   },
	//   "codedAmounts": [
	//     {
	//       "code": "D",
	//       "amount": {
	//         "minor": "900000",
	//         "currency": "USD"
	//       }
	//     }
	//   ]
	// }
}
