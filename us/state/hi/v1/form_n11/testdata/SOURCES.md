# Where this file came from

The line numbers in `mapping.textproto` are claims about a document the
Hawaii Department of Taxation publishes. This file is text extracted from that
document, so the claims can be checked.

Retrieved 2026-11-14 with `pdftotext -layout` from poppler 24.02.0, unfiltered —
the file is the whole form, with Schedule CR, Schedule X and its other pages.

| File | Source | sha256 of the PDF |
|---|---|---|
| `n11-2025.txt` | [2025 Form N-11](https://files.hawaii.gov/tax/forms/2025/n11_f.pdf) | `e0e707e1d2b3c58ec1e7bc837fce54ec87e61b8933fd2795dd8014147692b999` |

```bash
curl -o n11_f.pdf https://files.hawaii.gov/tax/forms/2025/n11_f.pdf
pdftotext -layout n11_f.pdf n11-2025.txt
```
