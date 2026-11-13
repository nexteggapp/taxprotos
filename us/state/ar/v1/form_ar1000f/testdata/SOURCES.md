# Where this file came from

The line numbers in `mapping.textproto` are claims about a document the
Arkansas Department of Finance and Administration publishes. This file is text
extracted from that document, so the claims can be checked.

Retrieved 2026-11-12 with `pdftotext -layout` from poppler 24.02.0, unfiltered —
the file is the whole form.

| File | Source | sha256 of the PDF |
|---|---|---|
| `ar1000f-2025.txt` | [2025 Form AR1000F](https://www.dfa.arkansas.gov/wp-content/uploads/2025_AR1000F_FullYearResidentIndividualIncomeTaxReturn.pdf) | `b8165006622e496020b03aa946e36208abc3290b48b8c57d7ff6206a6d8e06f7` |

```bash
curl -o ar1000f.pdf https://www.dfa.arkansas.gov/wp-content/uploads/2025_AR1000F_FullYearResidentIndividualIncomeTaxReturn.pdf
pdftotext -layout ar1000f.pdf ar1000f-2025.txt
```

## Most labels do not extract

The PDF sets most line labels in a font whose glyphs map to the wrong
characters, so pdftotext prints them shifted ("6LQJOH" for "Single"). The line
numbers, which the mapping relies on, extract cleanly; the labels were decoded by
shifting each character back 29 code points. Filing status boxes 1 and 4 have no extractable
number, so the filing status is not mapped to a line.
