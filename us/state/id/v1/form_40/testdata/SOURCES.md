# Where this file came from

The line numbers in `mapping.textproto` are claims about a document the
Idaho State Tax Commission publishes. This file is text extracted from that
document, so the claims can be checked.

Retrieved 2026-11-12 with `pdftotext -layout` from poppler 24.02.0, unfiltered —
the file is the whole PDF: the three pages of the form, then the Form 40
instructions that the same PDF carries on its later pages.

| File | Source | sha256 of the PDF |
|---|---|---|
| `form40-2025.txt` | [2025 Form 40 (EFO00089, 03-02-2026)](https://tax.idaho.gov/wp-content/uploads/forms/EFO00089/EFO00089_03-02-2026.pdf) | `1390489138a7df67095155aaf34e6e8f39a4fd36041594f95b766f5e33bb8a13` |

```bash
curl -o form40.pdf https://tax.idaho.gov/wp-content/uploads/forms/EFO00089/EFO00089_03-02-2026.pdf
pdftotext -layout form40.pdf form40-2025.txt
```
