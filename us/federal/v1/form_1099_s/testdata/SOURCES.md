# Where these files came from

The line numbers in `mapping.textproto` are claims about documents the IRS
publishes. These files are text extracted from those documents, so the claims
can be checked.

Retrieved 2026-10-24 with `pdftotext -layout` from poppler 24.02.0, unfiltered —
each file is the whole form.

| File | Source | sha256 of the PDF |
|---|---|---|
| `f1099s-2025.txt` | [f1099s--2025.pdf](https://www.irs.gov/pub/irs-prior/f1099s--2025.pdf) | `d32d8864c2af930286022c6d57546f36c4061e5af1a642b13544de655c99e2cf` |

```bash
curl -o f1099s--2025.pdf https://www.irs.gov/pub/irs-prior/f1099s--2025.pdf
pdftotext -layout f1099s--2025.pdf f1099s-2025.txt
```
