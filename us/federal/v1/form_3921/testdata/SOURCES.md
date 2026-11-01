# Where these files came from

The line numbers in `mapping.textproto` are claims about documents the IRS
publishes. These files are text extracted from those documents, so the claims
can be checked.

Retrieved 2026-11-01 with `pdftotext -layout` from poppler 24.02.0, unfiltered —
each file is the whole form.

| File | Source | sha256 of the PDF |
|---|---|---|
| `f3921-2025.txt` | [f3921--2025.pdf](https://www.irs.gov/pub/irs-prior/f3921--2025.pdf) | `8173947e5351009688ce609062ffff6eac2101a6183af45dcc63206e53f0b0fe` |

```bash
curl -o f3921--2025.pdf https://www.irs.gov/pub/irs-prior/f3921--2025.pdf
pdftotext -layout f3921--2025.pdf f3921-2025.txt
```
