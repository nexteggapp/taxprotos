# Where these files came from

The line numbers in `mapping.textproto` are claims about documents the IRS
publishes. These files are text extracted from those documents, so the claims
can be checked.

Retrieved 2026-10-24 with `pdftotext -layout` from poppler 24.02.0, unfiltered —
each file is the whole form.

| File | Source | sha256 of the PDF |
|---|---|---|
| `f5498-2025.txt` | [f5498.pdf](https://www.irs.gov/pub/irs-pdf/f5498.pdf) | `745fe725b8b59b5ab5d797045b2715416b7ef4bca86c7337fce210e064866f17` |

```bash
curl -o f5498.pdf https://www.irs.gov/pub/irs-pdf/f5498.pdf
pdftotext -layout f5498.pdf f5498-2025.txt
```
