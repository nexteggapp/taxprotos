# Where these files came from

The line numbers in `mapping_*.textproto` are claims about documents the Oregon
Department of Revenue publishes. These files are text extracted from those
documents, so the claims can be checked.

Retrieved 2026-08-26 with `pdftotext -layout` from poppler 24.02.0, unfiltered —
each file is the whole form.

| File | Source | sha256 of the PDF |
|---|---|---|
| `or40-2025.txt` | [2025 Form OR-40](https://www.oregon.gov/dor/forms/FormsPubs/form-or-40_101-040_2025.pdf) | `eaa5ca450083c7717e46846ce69d1b71ea64bcc58ae48d7f9c25ed9c8e9a56e6` |
| `ora-2025.txt` | [2025 Schedule OR-A](https://www.oregon.gov/dor/forms/FormsPubs/schedule-or-a_101-007_2025.pdf) | `f4a342d9dc71178ad4094cca71ab7753c431d1e512a1f10602ca378254d4f204` |
| `orasc-2025.txt` | [2025 Schedule OR-ASC](https://www.oregon.gov/dor/forms/FormsPubs/schedule-or-asc_101-063_2025.pdf) | `9fd4842a70033f6a78788cd3433000308a66c65d3ec2df12d7ded6f346907f7f` |

```bash
curl -o or-40.pdf    https://www.oregon.gov/dor/forms/FormsPubs/form-or-40_101-040_2025.pdf
curl -o or-a.pdf     https://www.oregon.gov/dor/forms/FormsPubs/schedule-or-a_101-007_2025.pdf
curl -o or-asc.pdf   https://www.oregon.gov/dor/forms/FormsPubs/schedule-or-asc_101-063_2025.pdf
pdftotext -layout or-40.pdf  or40-2025.txt
pdftotext -layout or-a.pdf   ora-2025.txt
pdftotext -layout or-asc.pdf orasc-2025.txt
```

## Oregon prints a full stop after its line numbers

The form says "7." where the federal forms say "7". `lineAppears` in
`coverage_test.go` allows one trailing stop for that reason, and the mappings
record `7` rather than `7.` — the number is the identifier and the stop is how
this year's PDF was typeset.

Schedule OR-ASC numbers by section instead, so its lines are `A5`, `B7`, `C5`,
`D16`, `E9` and `F7`. Those are checked the same way.

## Why extract at all

Because the alternative is writing line numbers from memory. Oregon's return has
a line for federal tax liability, which most states' do not, and a line for the
kicker, which no other state has at all. Neither would have been guessed.

Nothing here is parsed into structured data. `coverage_test.go` only asks
whether a line number appears in the text at all, so there is no column layout
to misread and no rows to drop silently.

## What this does not prove

That a line was *understood*. The extract confirms line 10 exists and mentions
federal tax liability; it cannot confirm the right concept was mapped to it.

Only the section totals of Schedule OR-ASC are modelled. Each section is a list
of numbered codes — several hundred across the six — and Form OR-40 consumes
only the totals. The codes are the Department's own list and this schema does
not duplicate it.
