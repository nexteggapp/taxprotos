# Where this file came from

The line numbers in `mapping.textproto` are claims about a document the Ohio
Department of Taxation publishes. This file is text extracted from that
document, so the claims can be checked.

Retrieved 2026-10-20 with `pdftotext` from poppler 24.02.0: the form's four
pages, with the instructions on page 4, in reading order.

| File | Source | sha256 of the PDF |
|---|---|---|
| `itsd2210-2025.txt` | [Ohio IT/SD 2210, Interest Penalty on Underpayment of Tax, 2025](https://dam.assets.ohio.gov/image/upload/v1762982566/tax.ohio.gov/forms/ohio_individual/individual/2025/itsd2210-fi.pdf) | `58e7940049cf248709d4829a44ae79117426cac19c3c495e07f0d6accc3f9fde` |

```bash
curl -o 2210.pdf https://dam.assets.ohio.gov/image/upload/v1762982566/tax.ohio.gov/forms/ohio_individual/individual/2025/itsd2210-fi.pdf
pdftotext 2210.pdf itsd2210-2025.txt
```
