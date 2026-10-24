# Where these files came from

The line numbers in `mapping.textproto` are claims about documents the IRS
publishes. These files are text extracted from those documents, so the claims
can be checked.

Retrieved 2026-10-24 with `pdftotext -layout` from poppler 24.02.0, unfiltered —
each file is the whole form.

| File | Source | sha256 of the PDF |
|---|---|---|
| `f8889-2025.txt` | [f8889.pdf](https://www.irs.gov/pub/irs-pdf/f8889.pdf) | `15ed4587f75bec57b0b343dfd001b4016f09698af970afee633ca1a364d5f49f` |

```bash
curl -o f8889.pdf https://www.irs.gov/pub/irs-pdf/f8889.pdf
pdftotext -layout f8889.pdf f8889-2025.txt
```
