# Where this file came from

The line numbers in `mapping.textproto` are claims about a document the
Massachusetts Department of Revenue publishes. This file is text extracted from
that document, so the claims can be checked.

Retrieved 2026-11-14 and extracted with `pdftotext -layout` from poppler
24.02.0, unfiltered — the file is the whole form.

| File | Source | sha256 of the PDF |
|---|---|---|
| `form1-2025.txt` | [2025 Form 1](https://www.mass.gov/doc/2025-form-1-massachusetts-resident-income-tax-return/download) | `0bb62498fda14b515084b83fb4467f379f17d08a712f6390b1cd2fceecb04a07` |

mass.gov refuses scripted requests (HTTP 403 for `curl`, even with a browser
User-Agent and Referer), so the PDF was downloaded through a web browser and
then extracted:

```bash
pdftotext -layout 2025-form-1.pdf form1-2025.txt
```

The PDF is four pages. Schedule 4% Surtax, which line 28b names, is a separate
schedule and is not part of it.

## Lines printed with their figure inside

Lines 2b, 2c, 2d, 14, 23a and 23b each print a count or an amount and the
result of multiplying or dividing it on one line; the mapping gives both
concepts that line, each with a note. Items a and b at the top of page 1 carry
federal Form 1040 lines 9 and 11a.

## Lines printed against their dot leaders

Lines 2e, 2g, 6a, 8a, 8b and 38b are printed with the dot leader touching the
number (".2e"), so the coverage test cannot find them standing alone. The
mapping gives them their bare number (2, 6, 8, 38) with a note naming the
part.
