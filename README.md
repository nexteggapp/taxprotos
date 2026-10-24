# taxprotos

Protobuf schemas for US tax documents: the W-2 and 1099 forms a household
receives, the Form 1040 it files, the state returns for California, Oregon and
Washington, and `TaxYearFacts`, a household's own facts for one tax year.

A schema holds what the form prints and nothing else. It has no confidence
scores, page coordinates or other traces of how a document was read. Software
that reads documents keeps those in a record that wraps these messages.

## Example

```proto
import "taxprotos/us/federal/v1/form_w2/form_w2.proto";

message PayrollRecord {
  nexteggapp.taxprotos.us.federal.v1.FormW2 w2 = 1;
}
```

A W-2 in the canonical proto3 JSON encoding:

```json
{
  "taxYear": 2025,
  "employer": { "name": "Acme Corporation", "tin": "12-3456789" },
  "employee": { "name": "Dana Okonkwo", "tin": "123-45-6789" },

  "wagesTipsOtherCompensation": { "minor": "12000000", "currency": "USD" },
  "federalTaxWithheld":         { "minor": "1840000",  "currency": "USD" },
  "socialSecurityWages":        { "minor": "12000000", "currency": "USD" },
  "socialSecurityTaxWithheld":  { "minor": "744000",   "currency": "USD" },

  "codedAmounts": [
    { "code": "D",  "amount": { "minor": "900000",  "currency": "USD" } },
    { "code": "DD", "amount": { "minor": "1420000", "currency": "USD" } }
  ],

  "stateTaxWithholding": [
    {
      "state": "CA",
      "stateId": "123-4567-8",
      "taxWithheld": { "minor": "610000",   "currency": "USD" },
      "stateIncome": { "minor": "12000000", "currency": "USD" }
    }
  ]
}
```

Money is an integer count of the currency's minor unit, so `$120,000.00` is
`"minor": "12000000"`. `minor` is a string because proto3 JSON encodes 64-bit
integers as strings. Box 12 is a repeated code-and-amount pair rather than a
field per code. The JSON above is the output of `ExampleFormW2_json` in
[`go/examples/example_test.go`](go/examples/example_test.go), which the test
suite checks.

## Using the schemas

This repository ships `.proto` files and no generated code. You generate
bindings for your language from the schemas.

### With Bazel

```starlark
# MODULE.bazel
bazel_dep(name = "nexteggapp_taxprotos", version = "0.1.0")
git_override(
    module_name = "nexteggapp_taxprotos",
    remote = "https://github.com/nexteggapp/taxprotos.git",
    commit = "<commit>",
)
```

```starlark
# BUILD.bazel
proto_library(
    name = "payroll_proto",
    srcs = ["payroll.proto"],
    deps = ["@nexteggapp_taxprotos//us/federal/v1/form_w2:form_w2_proto"],
)
```

The schema packages load only protobuf rules, so depending on them adds
`protobuf` and `rules_proto` to your module graph and nothing else. Generate Go,
Java or any other bindings from the `proto_library` targets as usual.

### With protoc or buf

Every schema imports the others as `taxprotos/<path>`, so put a checkout of
this repository at `taxprotos/` under an include root:

```bash
git clone https://github.com/nexteggapp/taxprotos.git third_party/taxprotos
protoc -I third_party --go_out=. --go_opt=module=github.com/nexteggapp/taxprotos \
  third_party/taxprotos/us/federal/v1/form_w2/form_w2.proto
```

Each file's `go_package` is `github.com/nexteggapp/taxprotos/<directory>`.

## Forms

### Information returns — what a household receives

| Form | Schema | Reports |
|---|---|---|
| [W-2](https://www.irs.gov/pub/irs-pdf/fw2.pdf) | [`form_w2.proto`](us/federal/v1/form_w2/form_w2.proto) | Wages and tax withheld |
| [1098](https://www.irs.gov/pub/irs-pdf/f1098.pdf) | [`form_1098.proto`](us/federal/v1/form_1098/form_1098.proto) | Mortgage interest |
| [1098-E](https://www.irs.gov/pub/irs-pdf/f1098e.pdf) | [`form_1098_e.proto`](us/federal/v1/form_1098_e/form_1098_e.proto) | Student loan interest |
| [1098-T](https://www.irs.gov/pub/irs-pdf/f1098t.pdf) | [`form_1098_t.proto`](us/federal/v1/form_1098_t/form_1098_t.proto) | Tuition |
| [1099-B](https://www.irs.gov/pub/irs-pdf/f1099b.pdf) | [`form_1099_b.proto`](us/federal/v1/form_1099_b/form_1099_b.proto) | Broker proceeds |
| [1099-DIV](https://www.irs.gov/pub/irs-pdf/f1099div.pdf) | [`form_1099_div.proto`](us/federal/v1/form_1099_div/form_1099_div.proto) | Dividends and distributions |
| [1099-G](https://www.irs.gov/pub/irs-pdf/f1099g.pdf) | [`form_1099_g.proto`](us/federal/v1/form_1099_g/form_1099_g.proto) | Unemployment compensation and state or local refunds |
| [1099-INT](https://www.irs.gov/pub/irs-pdf/f1099int.pdf) | [`form_1099_int.proto`](us/federal/v1/form_1099_int/form_1099_int.proto) | Interest income |
| [1099-K](https://www.irs.gov/pub/irs-pdf/f1099k.pdf) | [`form_1099_k.proto`](us/federal/v1/form_1099_k/form_1099_k.proto) | Payment card and network transactions |
| [1099-MISC](https://www.irs.gov/pub/irs-pdf/f1099msc.pdf) | [`form_1099_misc.proto`](us/federal/v1/form_1099_misc/form_1099_misc.proto) | Miscellaneous information |
| [1099-NEC](https://www.irs.gov/pub/irs-pdf/f1099nec.pdf) | [`form_1099_nec.proto`](us/federal/v1/form_1099_nec/form_1099_nec.proto) | Non-employee compensation |
| [1099-R](https://www.irs.gov/pub/irs-pdf/f1099r.pdf) | [`form_1099_r.proto`](us/federal/v1/form_1099_r/form_1099_r.proto) | Retirement distributions |
| [1099-SA](https://www.irs.gov/pub/irs-pdf/f1099sa.pdf) | [`form_1099_sa.proto`](us/federal/v1/form_1099_sa/form_1099_sa.proto) | HSA and MSA distributions |
| [5498](https://www.irs.gov/pub/irs-pdf/f5498.pdf) | [`form_5498.proto`](us/federal/v1/form_5498/form_5498.proto) | IRA contributions |
| [5498-SA](https://www.irs.gov/pub/irs-pdf/f5498sa.pdf) | [`form_5498_sa.proto`](us/federal/v1/form_5498_sa/form_5498_sa.proto) | HSA and MSA contributions |
| [SSA-1099](https://www.ssa.gov/myaccount/replacement-SSA-1099.html) | [`form_ssa_1099.proto`](us/federal/v1/form_ssa_1099/form_ssa_1099.proto) | Social Security benefits |
| [W-2G](https://www.irs.gov/pub/irs-pdf/fw2g.pdf) | [`form_w2g.proto`](us/federal/v1/form_w2g/form_w2g.proto) | Gambling winnings |

### The return — what a household files

| Form | Schema | Reports |
|---|---|---|
| [1040](https://www.irs.gov/pub/irs-pdf/f1040.pdf) | [`form_1040.proto`](us/federal/v1/form_1040/form_1040.proto) | Income, deductions, tax, payments, refund or amount owed |
| [Schedule 1](https://www.irs.gov/pub/irs-pdf/f1040s1.pdf) | [`schedule_1.proto`](us/federal/v1/form_1040/schedule_1.proto) | Additional income, and adjustments to reach AGI |
| [Schedule 1-A](https://www.irs.gov/pub/irs-pdf/f1040s1a.pdf) | [`schedule_1a.proto`](us/federal/v1/form_1040/schedule_1a.proto) | Tips, overtime, vehicle loan interest and senior deductions |
| [Schedule 2](https://www.irs.gov/pub/irs-pdf/f1040s2.pdf) | [`schedule_2.proto`](us/federal/v1/form_1040/schedule_2.proto) | Additional taxes, including self-employment tax |
| [Schedule 3](https://www.irs.gov/pub/irs-pdf/f1040s3.pdf) | [`schedule_3.proto`](us/federal/v1/form_1040/schedule_3.proto) | Additional credits and payments |
| [Schedule A](https://www.irs.gov/pub/irs-pdf/f1040sa.pdf) | [`schedule_a.proto`](us/federal/v1/form_1040/schedule_a.proto) | Itemized deductions |
| [Schedule D](https://www.irs.gov/pub/irs-pdf/f1040sd.pdf) | [`schedule_d.proto`](us/federal/v1/form_1040/schedule_d.proto) | Capital gains and losses |
| [Schedule SE](https://www.irs.gov/pub/irs-pdf/f1040sse.pdf) | [`schedule_se.proto`](us/federal/v1/form_1040/schedule_se.proto) | Self-employment tax |
| [8889](https://www.irs.gov/pub/irs-pdf/f8889.pdf) | [`form_8889.proto`](us/federal/v1/form_8889/form_8889.proto) | Health savings accounts |

### State returns — what three states want as well

| Form | Schema | Reports |
|---|---|---|
| [540](https://www.ftb.ca.gov/forms/2025/2025-540.pdf) | [`form_540.proto`](us/state/ca/v1/form_540/form_540.proto) | The California resident return |
| [Schedule CA](https://www.ftb.ca.gov/forms/2025/2025-540-ca.pdf) | [`schedule_ca.proto`](us/state/ca/v1/form_540/schedule_ca.proto) | How a federal return becomes a Californian one |
| [OR-40](https://www.oregon.gov/dor/forms/FormsPubs/form-or-40_101-040_2025.pdf) | [`form_or40.proto`](us/state/or/v1/form_or40/form_or40.proto) | The Oregon resident return |
| [Schedule OR-A](https://www.oregon.gov/dor/forms/FormsPubs/schedule-or-a_101-007_2025.pdf) | [`schedule_or_a.proto`](us/state/or/v1/form_or40/schedule_or_a.proto) | Oregon itemized deductions |
| [Schedule OR-ASC](https://www.oregon.gov/dor/forms/FormsPubs/schedule-or-asc_101-063_2025.pdf) | [`schedule_or_asc.proto`](us/state/or/v1/form_or40/schedule_or_asc.proto) | Oregon additions, subtractions and credits |
| [Capital gains](https://app.leg.wa.gov/rcw/default.aspx?cite=82.87) | [`capital_gains_return.proto`](us/state/wa/v1/capital_gains/capital_gains_return.proto) | Washington's excise tax on long-term gains |

Washington levies no personal income tax, so `us/state/wa` has no income tax
return. Its capital gains return is filed electronically and has no printed
line numbers, so its mapping cites subsections of chapter 82.87 RCW.

[`docs/form-schemas.md`](docs/form-schemas.md) describes how the form schemas
are built: field naming, money, dates, tax years, the FDX annotations and the
mapping files. [`docs/facts.md`](docs/facts.md) describes `TaxYearFacts`.

## Contributing

The repository builds and tests on its own, from a clone, with no other
checkout. Run every command below from the repository root. Pull requests run
the same two checks in [`.github/workflows/ci.yml`](.github/workflows/ci.yml).

With Bazel (the version is in [`.bazelversion`](.bazelversion)):

```bash
bazel test //...
```

This compiles every schema and runs the suites under [`go/`](go/), which check
the schemas against IRS Publication 1220 and against text extracted from the
published forms.

With Go:

```bash
go test ./...
```

This compiles every schema with [protocompile](https://github.com/bufbuild/protocompile)
under its `taxprotos/...` import path and checks each `go_package`. The suites
under `go/` need generated code and run only under Bazel.

## License

Apache 2.0. See [LICENSE](LICENSE) and [NOTICE](NOTICE).

These schemas describe what a form prints. They do not interpret it, and they
are not tax advice.
