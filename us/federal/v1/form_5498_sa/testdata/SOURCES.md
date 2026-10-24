# Where these files came from

The line numbers in `mapping.textproto` are claims about documents the IRS
publishes. These files are text extracted from those documents, so the claims
can be checked.

Retrieved 2026-10-24 with `pdftotext -layout` from poppler 24.02.0, unfiltered —
each file is the whole form.

| File | Source | sha256 of the PDF |
|---|---|---|
| `f5498sa-2025.txt` | [f5498sa--2025.pdf](https://www.irs.gov/pub/irs-prior/f5498sa--2025.pdf) | `3506c7849bd7faa89e445ccc800cd34e8fdbd4b970bdb16e7b2fd0c927c38aef` |

```bash
curl -o f5498sa--2025.pdf https://www.irs.gov/pub/irs-prior/f5498sa--2025.pdf
pdftotext -layout f5498sa--2025.pdf f5498sa-2025.txt
```
