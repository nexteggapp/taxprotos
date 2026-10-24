# Where these files came from

The line numbers in `mapping_*.textproto` are claims about documents the IRS
publishes. These files are text extracted from those documents, so the claims
can be checked.

Retrieved 2026-08-26 with `pdftotext -layout` from poppler 24.02.0, unfiltered —
each file is the whole form.

| File | Source | sha256 of the PDF |
|---|---|---|
| `f1040-2025.txt` | [f1040.pdf](https://www.irs.gov/pub/irs-pdf/f1040.pdf) | `3d31c226df0d189c...` |
| `f1040s1-2025.txt` | [f1040s1.pdf](https://www.irs.gov/pub/irs-pdf/f1040s1.pdf) | `8dafec719f6a4716...` |
| `f1040s1a-2025.txt` | [f1040s1a.pdf](https://www.irs.gov/pub/irs-pdf/f1040s1a.pdf) | `64f97b38ff4218e2...` |
| `f1040s2-2025.txt` | [f1040s2.pdf](https://www.irs.gov/pub/irs-pdf/f1040s2.pdf) | `64d867b683334cfc...` |
| `f1040s3-2025.txt` | [f1040s3.pdf](https://www.irs.gov/pub/irs-pdf/f1040s3.pdf) | `008cfd3fe3ebd086...` |
| `f1040sa-2025.txt` | [f1040sa.pdf](https://www.irs.gov/pub/irs-pdf/f1040sa.pdf) | `c14acf3478f4c33f...` |
| `f1040sd-2025.txt` | [f1040sd.pdf](https://www.irs.gov/pub/irs-pdf/f1040sd.pdf) | `90564c8b7e492803...` |
| `f1040sse-2025.txt` | [f1040sse.pdf](https://www.irs.gov/pub/irs-pdf/f1040sse.pdf) | `05bc2b3e1dfca65d...` |

```bash
curl -o f1040.pdf https://www.irs.gov/pub/irs-pdf/f1040.pdf
pdftotext -layout f1040.pdf f1040-2025.txt
```

## Why extract at all

Because the alternative is writing line numbers from memory, and memory is
wrong. For tax year 2025 it would have put the deduction on line 12, adjusted
gross income on line 11, and would not have known Schedule 1-A exists. The
printed forms say 12e, 11a, and it does — a schedule new that year carrying the
tips, overtime, vehicle loan interest and senior deductions.

Unlike the tax table in taxcalcs, nothing here is parsed into structured data.
`coverage_test.go` only asks whether a line number appears in the text at all,
so there is no column layout to misread and no rows to drop silently. What it
buys is narrow and real: a mapping cannot claim a line the form does not have.

## What this does not prove

That a line was *understood*. The extract confirms line 12e exists and mentions
a deduction; it cannot confirm that the right concept was mapped to it. That is
what Publication 1220 does for the twelve information returns, by enumerating
the fields, and there is no equivalent for a return.
