# Where this file came from

The line numbers in `mapping.textproto` are claims about a document the
Georgia Department of Revenue publishes. This file is text extracted from that
document, so the claims can be checked.

Retrieved 2026-10-20 with `pdftotext` from poppler 24.02.0, in reading order:
the form and its instructions. The revision is marked "For tax years 2025 and
later".

| File | Source | sha256 of the PDF |
|---|---|---|
| `form500uet-2025.txt` | [Form 500 UET (Rev. 07/15/25)](https://dor.georgia.gov/document/document/beginning-or-after-january-1-2025-500-uet-underpayment-estimated-tax-individuals/download) | `364ea3b0644fbd26a9753732b3b91c16f0e6b8b7a5772451c70d08677432b9b0` |

```bash
curl -L -o uet.pdf https://dor.georgia.gov/document/document/beginning-or-after-january-1-2025-500-uet-underpayment-estimated-tax-individuals/download
pdftotext uet.pdf form500uet-2025.txt
```
