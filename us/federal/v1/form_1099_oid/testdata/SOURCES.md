# Where these files came from

The line numbers in `mapping.textproto` are claims about documents the IRS
publishes. These files are text extracted from those documents, so the claims
can be checked.

Retrieved 2026-10-24 with `pdftotext -layout` from poppler 24.02.0, unfiltered —
each file is the whole form.

| File | Source | sha256 of the PDF |
|---|---|---|
| `f1099oid-2025.txt` | [f1099oid.pdf](https://www.irs.gov/pub/irs-pdf/f1099oid.pdf) | `cc88d18964125c3f69a6fd15afcb5cd8b90fb1c3ea68c7b3b7b91a106e17d756` |

```bash
curl -o f1099oid.pdf https://www.irs.gov/pub/irs-pdf/f1099oid.pdf
pdftotext -layout f1099oid.pdf f1099oid-2025.txt
```
