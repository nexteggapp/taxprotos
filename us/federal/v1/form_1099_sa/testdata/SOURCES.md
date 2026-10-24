# Where these files came from

The line numbers in `mapping.textproto` are claims about documents the IRS
publishes. These files are text extracted from those documents, so the claims
can be checked.

Retrieved 2026-10-24 with `pdftotext -layout` from poppler 24.02.0, unfiltered —
each file is the whole form.

| File | Source | sha256 of the PDF |
|---|---|---|
| `f1099sa-2025.txt` | [f1099sa.pdf](https://www.irs.gov/pub/irs-pdf/f1099sa.pdf) | `cc2324629a9bd416180c1e10b66ce44d914937a6c4a11cc4bb8e17c22da44bd1` |

```bash
curl -o f1099sa.pdf https://www.irs.gov/pub/irs-pdf/f1099sa.pdf
pdftotext -layout f1099sa.pdf f1099sa-2025.txt
```
