# Tax year facts

`TaxYearFacts`, in `facts/v1/tax_year_facts.proto`, is one household's own
facts for one tax year: who is in it, where it lived, the choices and judgement
calls the law leaves to it, and last year's returns. A calculator computes the
year from this record together with its own models of each income source.

Documents and bank entries are not stored in `TaxYearFacts`. Each sits in the
part of a calculator's model that uses it, such as a paystub in an employer's
paystubs. [taxcalcs](https://github.com/nexteggapp/taxcalcs) defines those
models.

## What it holds

| Field | Message | Holds |
|---|---|---|
| `tax_year` | | The year, as four digits |
| `people` | `Person` | Everyone in the household, the filer included. Exactly one is `RELATIONSHIP_YOU`. |
| `residence` | `ResidencePeriod` | Where the household lived, as dated periods. A move is two periods. |
| `elections` | `Election` | Choices the law gives the taxpayer: filing status, standard or itemised deduction, lot method per broker, withholding dates, a penalty waiver request |
| `positions` | `Position` | Answers to tests the law sets about the facts, such as whether you provided more than half of a person's support |
| `carryovers` | `Carryover` | Figures last year's return carries into this year, such as an unused capital loss |
| `prior_year` | `PriorYearReturn` | Last year's return in each jurisdiction: the filed return, entered figures, or an estimate |
| `unassigned` | `Document` | Documents read but not yet attached to a model |
| `adjacent_payments` | `AdjacentPayment` | Read-only copies of payments from the year before or after that can count toward this year |

Treatment is never stored. Whether a child qualifies for a credit, for example,
is computed each time from the people, positions and the law.

## Positions

A `Position` names a question by id, the ids of the people or things it is
about, an answer and the basis for it. The answer is a `Reply` (`YES`, `NO` or
`NOT_SURE`), a choice, or an amount. The calculator defines which questions
exist and reports any it needed and found unanswered.

| Question id | About | Law |
|---|---|---|
| `dependent.gross_income_under_limit` | a person | 26 U.S.C. § 152(d)(1)(B) |
| `dependent.more_than_half_support` | a person | § 152(d)(1)(C) |
| `head_of_household.more_than_half_home_cost` | the household | § 2(b)(1) |
| `surviving_spouse.spouse_died_within_two_years` | the household | § 2(a) |
| `deduction.spouse_itemises_separately` | the household | § 63(c)(6)(A) |

## Documents

`facts/v1/documents.proto` defines the documents a household holds that are not
IRS forms, and the references that say where a document came from.

A document is its form message exactly as printed. A W-2's box 1 is stored as
the employer reported it, after 401(k) deferrals. The federal and state form
messages are used directly. These have messages of their own:

| Message | Document |
|---|---|
| `Paystub`, with `PayFigures` for this paycheck and year to date | A pay statement |
| `Form1099BStatement` | A broker's 1099-B, one `Form1099B` per sale |
| `MortgageStatement` | A monthly mortgage statement |
| `PropertyTaxBill` | A property tax bill and its installments |
| `Receipt` | A receipt |
| `FormW4` | An employee's withholding certificate, as filled in |
| `PaymentConfirmation` | A confirmation of an estimated, balance-due or extension payment |
| `FederalReturn`, `CaliforniaReturn`, `OregonReturn` | A filed return with its schedules |

`Document` pairs one of these, or a form message, with a `DocumentRef`.

`DocumentRef` sits beside a form wherever the form is stored. It carries
`vault_ref`, the storing application's opaque reference to the file (empty when
the file is not stored), a display name, the day it was received, `removed`
(taken out of this year; it counts for nothing) and `use_anyway` (used although
its year or payer differs from its model's). `LedgerRef` does the same for a
bank or brokerage entry, with the day it posted.

## Conventions

- Every id is 1 to 64 letters, digits, `.`, `_` or `-`, unique across the
  record and the calculator's models. Ids are minted once and kept, so they are
  the same in every export and across years.
- A jurisdiction is `US` or a state's postal code.
- Every `oneof` is required: a message that has one sets exactly one case. The
  exception is `PriorYearReturn.filed`, which is unset when last year's return
  is not available.
- Identifiers are left empty: every `tin`, account number, control number,
  state id, routing number, `Party.address` and `Party.telephone`. Nothing is
  computed from them.
- No box or line numbers are stored. The form's `mapping.textproto` gives them;
  see [form-schemas.md](form-schemas.md).
- An amount with no document, such as a benefit entered before its form
  arrives, belongs in the calculator's model of that source, not in a made-up
  document.
