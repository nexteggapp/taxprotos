# Where this file came from

The line numbers in `mapping.textproto` are claims about a document the Ohio
Department of Taxation publishes. This file is text extracted from that
document, so the claims can be checked.

Retrieved 2026-11-14 with `pdftotext -layout` from poppler 24.02.0. The PDF is
the department's IT 1040 bundle, which prints the return on pages 1 and 2 and
its schedules after; only those two pages are extracted, so that a line number
can only be found on the return itself.

| File | Source | sha256 of the PDF |
|---|---|---|
| `it1040-2025.txt` | [2025 Ohio IT 1040 bundle](https://dam.assets.ohio.gov/image/upload/tax.ohio.gov/forms/ohio_individual/individual/2025/1040-bundle.pdf) | `612775da1fdce580451a4c8b9d0e47c5a1a3ffbf0054a06f000edd9c0b5d343b` |

```bash
curl -o 1040-bundle.pdf https://dam.assets.ohio.gov/image/upload/tax.ohio.gov/forms/ohio_individual/individual/2025/1040-bundle.pdf
pdftotext -layout -f 1 -l 2 1040-bundle.pdf it1040-2025.txt
```
