# Where these files came from

The line numbers in `mapping.textproto` are claims about documents the IRS
publishes. These files are text extracted from those documents, so the claims
can be checked.

Retrieved 2026-11-01 with `pdftotext -layout` from poppler 24.02.0, unfiltered —
each file is the whole form.

| File | Source | sha256 of the PDF |
|---|---|---|
| `f1099ptr-2025.txt` | [f1099ptr--2025.pdf](https://www.irs.gov/pub/irs-prior/f1099ptr--2025.pdf) | `4bdef5fc7c3bfb831af9be9365aaf6b1c691767ed3c9b787c15f49e85db557b3` |

```bash
curl -o f1099ptr--2025.pdf https://www.irs.gov/pub/irs-prior/f1099ptr--2025.pdf
pdftotext -layout f1099ptr--2025.pdf f1099ptr-2025.txt
```
