# Where this file came from

The line numbers in `mapping.textproto` are claims about a document the New
York State Department of Taxation and Finance publishes. This file is text
extracted from that document, so the claims can be checked.

Retrieved 2026-11-16 with curl 8.5.0 and `pdftotext -layout` from poppler
24.02.0: both pages of the form, in page layout. The department publishes
the 2025 form only as a fill-in PDF; `it2105_9_2025.pdf` returns a 404 page.
Worksheet 1, which `AnnualizationPeriod` follows, is on page 6 of the
instructions, IT-2105.9-I (2025), at
https://www.tax.ny.gov/pdf/2025/inc/it2105_9i_2025.pdf.

| File | Source | sha256 of the PDF |
|---|---|---|
| `it-2105.9-2025.txt` | [Form IT-2105.9 (2025), Underpayment of Estimated Tax By Individuals and Fiduciaries](https://www.tax.ny.gov/pdf/2025/inc/it2105_9_2025_fill_in.pdf) | `5e86a7ea75f6f66a3019c03492fdd4836f95fb754feb1623ab006b63ffcf0324` |

```bash
curl -o it2105_9.pdf https://www.tax.ny.gov/pdf/2025/inc/it2105_9_2025_fill_in.pdf
pdftotext -layout it2105_9.pdf it-2105.9-2025.txt
```
