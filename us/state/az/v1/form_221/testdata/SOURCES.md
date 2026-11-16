# Where this file came from

The line numbers in `mapping.textproto` are claims about a document the
Arizona Department of Revenue publishes. This file is text extracted from that
document, so the claims can be checked.

Retrieved 2026-10-20 with `pdftotext -layout` from poppler 24.02.0, unfiltered —
the file is the whole form, both pages, the Annualized Income Installment
Worksheet included. The 2026 Form 221 is not yet published.

| File | Source | sha256 of the PDF |
|---|---|---|
| `form221-2025.txt` | [2025 Form 221](https://azdor.gov/sites/default/files/document/FORMS_INDIVIDUAL_2025_221_f.pdf) | `d09ada6d2ebf1d5c827ab4fbb593efd16b6dfb82e684680d960ee4181844ac41` |

```bash
curl -o 221.pdf https://azdor.gov/sites/default/files/document/FORMS_INDIVIDUAL_2025_221_f.pdf
pdftotext -layout 221.pdf form221-2025.txt
```
