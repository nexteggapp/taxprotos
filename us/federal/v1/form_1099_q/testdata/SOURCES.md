# Where these files came from

The line numbers in `mapping.textproto` are claims about documents the IRS
publishes. These files are text extracted from those documents, so the claims
can be checked.

Retrieved 2026-10-24 with `pdftotext -layout` from poppler 24.02.0, unfiltered —
each file is the whole form.

| File | Source | sha256 of the PDF |
|---|---|---|
| `f1099q-2025.txt` | [f1099q--2025.pdf](https://www.irs.gov/pub/irs-prior/f1099q--2025.pdf) | `19ad3c80de52d509c30eac7d5de2527806c043530f567234ea310899cba4b7bb` |

```bash
curl -o f1099q--2025.pdf https://www.irs.gov/pub/irs-prior/f1099q--2025.pdf
pdftotext -layout f1099q--2025.pdf f1099q-2025.txt
```
