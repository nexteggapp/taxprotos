# Form schemas

Each tax form is one protobuf message that holds what the printed form says.
This page describes the conventions every form schema follows, the files that
sit beside each schema, and how to add a form or a state.

## Layout

```
common/v1/                   Money, Quantity, Date, Party, annotations, FormMapping
us/federal/v1/<form>/        one federal form: <form>.proto, mapping.textproto, BUILD.bazel
us/state/<postal>/v1/<form>/ one state form, same files
facts/v1/                    TaxYearFacts and the document messages (see facts.md)
go/                          Go bindings and test suites, built only with Bazel
```

Proto packages follow the path: `nexteggapp.taxprotos.us.federal.v1`,
`nexteggapp.taxprotos.us.state.ca.v1`, and so on. Every file is imported as
`taxprotos/<path>`, which is the `import_prefix` on each `proto_library`. A
state directory uses the lowercase two-letter postal code.

A schema's `BUILD.bazel` holds a `proto_library` and loads nothing else.
The Go targets for the same package live at the mirrored path under `go/`, so a
consumer of the schemas never loads `rules_go`.

## Conventions

| Rule | Example |
|---|---|
| A field is named for the concept it holds, never for a box or line number. | `qualified_dividends`, not `box_1b` |
| Every amount is `common.v1.Money`: a signed `int64` count of minor units and an ISO 4217 code. No floats. | `{ "minor": "12000000", "currency": "USD" }` is $120,000.00 |
| Every rate is `common.v1.Quantity`: an integer and a scale. | 5.25% is `units: 525, scale: 2` |
| Dates are `common.v1.Date` and may be partial. | `{ "year": 2025 }` |
| A person or business is `common.v1.Party`. | employer, payer, lender |
| Subtotals the form prints get fields of their own. They are recorded as filed, not recomputed. | Form 1040 line 9, total income |
| One message serves every year of its form, and `tax_year` says which year a record is for. | |
| A field that applies only from or until some year carries `valid_from_tax_year` or `valid_through_tax_year`. | Schedule 1-A's deductions are valid from 2025 |
| When a box changes meaning, it becomes a new field and the old field number is reserved. | |
| A schema imports nothing from outside this repository. | |
| Each state has its own messages. No message abstracts over several states. | California's 540 and Oregon's OR-40 are unrelated messages |
| A state that does not levy a tax has no schema for it. | Washington has no income tax return |

Field numbers and names are stable: records written with these schemas are
stored, so a field is never renumbered or renamed.

`us/federal/v1/form_identity.proto` defines `FormIdentity`, which records what
the printed form said about itself, such as its revision marker and catalog
number. It is advisory; `tax_year` is the field that decides the layout and the
law.

## FDX annotations

The [Financial Data Exchange](https://taxdataexchange.org/docs/fdx/v5.0/index.html)
(FDX) tax profile names fields by concept, as these schemas do. Options from
`common/v1/annotations.proto` record the correspondence on each message and
field:

| Option | On | Meaning |
|---|---|---|
| `fdx_entity` | message | The FDX entity, such as `TaxW2` |
| `fdx_version` | message | The FDX version the correspondence was written against |
| `fdx_field` | field | The FDX field it corresponds to, such as `qualifiedDividends` |
| `fdx_deviation` | field | How the field differs from that FDX field |
| `fdx_absent` | field | FDX has no counterpart |
| `valid_from_tax_year`, `valid_through_tax_year` | field | The years the field applies to |

Money and rates differ from FDX everywhere: FDX carries both as JSON numbers.
Fields of those types do not carry `fdx_deviation`, because the difference
follows from the type.

## Mapping files

Box and line numbers live in `mapping.textproto` beside each schema, as a
`common.v1.FormMapping`, keyed by tax year:

```textproto
# proto-message: nexteggapp.taxprotos.common.v1.FormMapping
form: "1099-DIV"
message_name: "nexteggapp.taxprotos.us.federal.v1.Form1099DIV"
tax_years: 2025

concepts {
  concept: "qualified_dividends"
  years { tax_year: 2025  irs_box: "1b"  amount_code: "2" }
}
```

`concept` is a field name on the message. `irs_box` is the printed box or line.
`amount_code` is the code IRS Publication 1220 uses for electronic filing.
`absent` marks a concept the form did not have that year, and `note` holds
anything else, such as the statute subsection for Washington's return, which has
no printed line numbers. `unmapped_amount_codes` lists Publication 1220 codes
that no field holds, each with a reason.

## How the schemas are checked

Every claim about a form is checked against a document its issuing authority
publishes. Text extracted from that document is checked into the package's
`testdata/` directory, and `testdata/SOURCES.md` records the URL, the source
file's sha256, the extraction tool and the command.

| Check | Where | Catches |
|---|---|---|
| Every Publication 1220 amount code is mapped or listed as unmapped | `go/us/federal/v1/pub1220` | a box missing from an information return |
| Every concept a mapping names is a field on its message | each coverage test | a field renamed without its mapping |
| Every line a mapping names appears in the extracted form text | each coverage test | a mapping left stale by renumbering |
| Every money field is named by a mapping | each coverage test | a field nobody can trace to the form |
| Every element of the facts schemas has a doc comment | `go/facts/v1` | undocumented fields |
| Every schema compiles under its `taxprotos/...` path, with a matching `go_package` | `schemas_test.go` | a broken import or package option |

A printed form is a weaker check than Publication 1220. It shows a line exists,
not that the right concept was mapped to it. The Form 1040 and the state
returns rely on it because Publication 1220 covers only information returns.

## Adding a form or a state

1. Find out from the state's revenue authority what it levies and which returns
   an ordinary resident files. Model those returns and nothing else.
2. Download each form, extract its text (for a PDF, `pdftotext -layout`), and
   check the extract into `testdata/` with a `SOURCES.md` entry.
3. Write one message per form in `us/state/<postal>/v1/<form>/<form>.proto`,
   following the conventions above, and a `BUILD.bazel` with a `proto_library`
   using `import_prefix = "taxprotos"`.
4. Write `mapping.textproto` for every concept.
5. Add the Go targets at `go/us/state/<postal>/v1/<form>/BUILD.bazel`, copying
   `go/us/federal/v1/form_w2/BUILD.bazel`, and a coverage test copying
   `go/us/federal/v1/form_1040/coverage_test.go`.
6. Run `bazel test //...` and `go test ./...`.
