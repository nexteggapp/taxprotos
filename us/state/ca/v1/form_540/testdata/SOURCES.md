# Where these files came from

The line numbers in `mapping_*.textproto` are claims about documents the
California Franchise Tax Board publishes. These files are text extracted from
those documents, so the claims can be checked.

Retrieved 2026-08-26 with `pdftotext -layout` from poppler 24.02.0, unfiltered —
each file is the whole form.

| File | Source | sha256 of the PDF |
|---|---|---|
| `f540-2025.txt` | [2025 Form 540](https://www.ftb.ca.gov/forms/2025/2025-540.pdf) | `3be81dcc5d2be3d303596f6fa528e8fdb7cda6644afac6ccf869fc5ebeb7bf1c` |
| `f540ca-2025.txt` | [2025 Schedule CA (540)](https://www.ftb.ca.gov/forms/2025/2025-540-ca.pdf) | `0927485bff61e49f21623c2df4f1e92b3b4d20ed52377814684f9bf814980660` |

```bash
curl -o 2025-540.pdf    https://www.ftb.ca.gov/forms/2025/2025-540.pdf
curl -o 2025-540-ca.pdf https://www.ftb.ca.gov/forms/2025/2025-540-ca.pdf
pdftotext -layout 2025-540.pdf    f540-2025.txt
pdftotext -layout 2025-540-ca.pdf f540ca-2025.txt
```

## Why extract at all

Because the alternative is writing line numbers from memory, and memory is
wrong. It would have called line 62 the Mental Health Services Tax, which is
what it was until Proposition 1 renamed it for taxable years beginning on or
after 1 January 2025. The printed form says Behavioral Health Services Tax, and
the mapping records both the current name and the old one.

Nothing here is parsed into structured data. `coverage_test.go` only asks
whether a line number appears in the text at all, so there is no column layout
to misread and no rows to drop silently. What it buys is narrow and real: a
mapping cannot claim a line the form does not have.

## Schedule CA needs the note as much as the number

Part I of Schedule CA prints three columns against each line number — the
federal amount, California's subtractions, California's additions — so line 27
identifies three different figures. `TestSharedLinesCarryANote` requires a
concept sharing a line to say which column it is, because without that the
mapping is unusable rather than merely terse.

## What this does not prove

That a line was *understood*. The extract confirms line 62 exists and mentions a
behavioural health tax; it cannot confirm that the right concept was mapped to
it. Publication 1220 does that for the twelve federal information returns by
enumerating the fields, and there is no equivalent for a state return.

Schedule D (540) is not modelled. California taxes capital gain at the ordinary
rates, so the gain reaches Form 540 inside federal adjusted gross income and
through the Schedule CA adjustments rather than on a line of its own.
