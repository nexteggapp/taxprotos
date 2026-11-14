# Where this file came from

The line numbers in `mapping.textproto` are claims about a document the
Georgia Department of Revenue publishes. This file is text extracted from that
document, so the claims can be checked.

The Department prints the 2025 Form 500 inside the IT-511 Individual Income
Tax Booklet, on PDF pages 41 to 45. The extract is those five pages.

Retrieved 2026-11-12 with `pdftotext -layout` from poppler 24.02.0.

| File | Source | sha256 of the PDF |
|---|---|---|
| `form500-2025.txt` | [2025 IT-511 Individual Income Tax Booklet](https://dor.georgia.gov/document/document/2025-it-511-individual-income-tax-booklet/download) | `ae43633e68ca423609607b0f93b1ebcd6056bce06ffb91829c448292e48a693f` |

```bash
curl -L -o it511.pdf https://dor.georgia.gov/document/document/2025-it-511-individual-income-tax-booklet/download
pdftotext -layout -f 41 -l 45 it511.pdf form500-2025.txt
```
