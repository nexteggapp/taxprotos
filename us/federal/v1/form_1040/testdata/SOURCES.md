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
| `f1040sb-2025.txt` | [f1040sb.pdf](https://www.irs.gov/pub/irs-pdf/f1040sb.pdf) | `dd1ec3719954532b...` |
| `f1040sc-2025.txt` | [f1040sc.pdf](https://www.irs.gov/pub/irs-pdf/f1040sc.pdf) | `ddf401dbe060467d...` |
| `f1040se-2025.txt` | [f1040se.pdf](https://www.irs.gov/pub/irs-pdf/f1040se.pdf) | `c5e2c8fa0ec5ee91...` |
| `f1040sse-2025.txt` | [f1040sse.pdf](https://www.irs.gov/pub/irs-pdf/f1040sse.pdf) | `05bc2b3e1dfca65d...` |
| `f8949-2025.txt` | [f8949.pdf](https://www.irs.gov/pub/irs-pdf/f8949.pdf) | `274513891e4e281d...` |
| `f1040x-2025.txt` | [f1040x--2025.pdf](https://www.irs.gov/pub/irs-prior/f1040x--2025.pdf) | `e4d8d7df37209218...` |
| `f1116-2025.txt` | [f1116--2025.pdf](https://www.irs.gov/pub/irs-prior/f1116--2025.pdf) | `2575162f86c37297...` |
| `f2210-2025.txt` | [f2210--2025.pdf](https://www.irs.gov/pub/irs-prior/f2210--2025.pdf) | `6899ce672648b280...` |
| `f2441-2025.txt` | [f2441--2025.pdf](https://www.irs.gov/pub/irs-prior/f2441--2025.pdf) | `6c3c2d19163fa4c4...` |
| `f2555-2025.txt` | [f2555--2025.pdf](https://www.irs.gov/pub/irs-prior/f2555--2025.pdf) | `979a90f0f83723a4...` |
| `f4562-2025.txt` | [f4562--2025.pdf](https://www.irs.gov/pub/irs-prior/f4562--2025.pdf) | `c05f9d1f5e26b1b2...` |
| `f4684-2025.txt` | [f4684--2025.pdf](https://www.irs.gov/pub/irs-prior/f4684--2025.pdf) | `911e0fe84adab9c7...` |
| `f4797-2025.txt` | [f4797--2025.pdf](https://www.irs.gov/pub/irs-prior/f4797--2025.pdf) | `e8ab787f2caab6a5...` |
| `f4952-2025.txt` | [f4952--2025.pdf](https://www.irs.gov/pub/irs-prior/f4952--2025.pdf) | `5113c1669c7ca4ab...` |
| `f5329-2025.txt` | [f5329--2025.pdf](https://www.irs.gov/pub/irs-prior/f5329--2025.pdf) | `2ad77023aebb280c...` |
| `f5695-2025.txt` | [f5695--2025.pdf](https://www.irs.gov/pub/irs-prior/f5695--2025.pdf) | `e2ac6a5c4b39e903...` |
| `f6198-2025.txt` | [f6198--2025.pdf](https://www.irs.gov/pub/irs-prior/f6198--2025.pdf) | `8ae169186ad9f629...` |
| `f6251-2025.txt` | [f6251--2025.pdf](https://www.irs.gov/pub/irs-prior/f6251--2025.pdf) | `6995bfd29c6fe1b8...` |
| `f6252-2025.txt` | [f6252--2025.pdf](https://www.irs.gov/pub/irs-prior/f6252--2025.pdf) | `b2d42f3194dc715c...` |
| `f7203-2025.txt` | [f7203.pdf](https://www.irs.gov/pub/irs-pdf/f7203.pdf) | `e138560a8c5d591f...` |
| `f7206-2025.txt` | [f7206--2025.pdf](https://www.irs.gov/pub/irs-prior/f7206--2025.pdf) | `822203c9bbd958e6...` |
| `f8582-2025.txt` | [f8582--2025.pdf](https://www.irs.gov/pub/irs-prior/f8582--2025.pdf) | `d929884fe07b58ac...` |
| `f8606-2025.txt` | [f8606--2025.pdf](https://www.irs.gov/pub/irs-prior/f8606--2025.pdf) | `748a2bf2ef45918e...` |
| `f8615-2025.txt` | [f8615--2025.pdf](https://www.irs.gov/pub/irs-prior/f8615--2025.pdf) | `f1a9ce61a62faa91...` |
| `f8801-2025.txt` | [f8801--2025.pdf](https://www.irs.gov/pub/irs-prior/f8801--2025.pdf) | `b82dff67ecf37406...` |
| `f8814-2025.txt` | [f8814--2025.pdf](https://www.irs.gov/pub/irs-prior/f8814--2025.pdf) | `d88849a33a390cee...` |
| `f8824-2025.txt` | [f8824--2025.pdf](https://www.irs.gov/pub/irs-prior/f8824--2025.pdf) | `0d61359bff5b3137...` |
| `f8829-2025.txt` | [f8829--2025.pdf](https://www.irs.gov/pub/irs-prior/f8829--2025.pdf) | `994834c9eeacdb37...` |
| `f8839-2025.txt` | [f8839--2025.pdf](https://www.irs.gov/pub/irs-prior/f8839--2025.pdf) | `3d3102f9b3dc3e4e...` |
| `f8863-2025.txt` | [f8863--2025.pdf](https://www.irs.gov/pub/irs-prior/f8863--2025.pdf) | `a251a1cfbd613d5d...` |
| `f8880-2025.txt` | [f8880--2025.pdf](https://www.irs.gov/pub/irs-prior/f8880--2025.pdf) | `ba2b28770bc82d31...` |
| `f8936-2025.txt` | [f8936--2025.pdf](https://www.irs.gov/pub/irs-prior/f8936--2025.pdf) | `de9c6988fbe34811...` |
| `f8959-2025.txt` | [f8959--2025.pdf](https://www.irs.gov/pub/irs-prior/f8959--2025.pdf) | `13e64004948331d2...` |
| `f8960-2025.txt` | [f8960--2025.pdf](https://www.irs.gov/pub/irs-prior/f8960--2025.pdf) | `9b323b7166b5845a...` |
| `f8962-2025.txt` | [f8962--2025.pdf](https://www.irs.gov/pub/irs-prior/f8962--2025.pdf) | `dc7d2c6b566cac66...` |
| `f8995-2025.txt` | [f8995--2025.pdf](https://www.irs.gov/pub/irs-prior/f8995--2025.pdf) | `55380ad230303e15...` |
| `f8995a-2025.txt` | [f8995a--2025.pdf](https://www.irs.gov/pub/irs-prior/f8995a--2025.pdf) | `3362db81b8ef60cf...` |
| `f1040s8-2025.txt` | [f1040s8--2025.pdf](https://www.irs.gov/pub/irs-prior/f1040s8--2025.pdf) | `6936462d67e2309d...` |
| `f1040sei-2025.txt` | [f1040sei--2025.pdf](https://www.irs.gov/pub/irs-prior/f1040sei--2025.pdf) | `b6dab52511ece16e...` |
| `f1040sf-2025.txt` | [f1040sf--2025.pdf](https://www.irs.gov/pub/irs-prior/f1040sf--2025.pdf) | `c6f7f19e1a4f5729...` |
| `f1040sh-2025.txt` | [f1040sh--2025.pdf](https://www.irs.gov/pub/irs-prior/f1040sh--2025.pdf) | `d04de932bd7d960b...` |
| `f1040sj-2025.txt` | [f1040sj--2025.pdf](https://www.irs.gov/pub/irs-prior/f1040sj--2025.pdf) | `5c888d02d8f12d99...` |
| `f8936sa-2025.txt` | [f8936sa--2025.pdf](https://www.irs.gov/pub/irs-prior/f8936sa--2025.pdf) | `8db7ab70c8ae5e79...` |

The attached forms below the Form 8949 row were retrieved 2026-11-01, the same
way. By then the IRS had moved the 2025 revisions to `irs-prior/`, which keeps
each year's form at a fixed URL. Form 7203 has no yearly revision; the December
2022 revision is the one in use for 2025, taken from `irs-pdf/`. Form 8936
Schedule A is a separate PDF, extracted beside Form 8936.

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
