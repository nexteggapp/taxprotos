# Where these files came from

The line numbers in `mapping.textproto` are claims about documents the IRS
publishes. These files are text extracted from those documents, so the claims
can be checked.

Retrieved 2026-10-24 with `pdftotext -layout` from poppler 24.02.0, unfiltered —
each file is the whole form.

| File | Source | sha256 of the PDF |
|---|---|---|
| `f1099msc-2025.txt` | [f1099msc--2025.pdf](https://www.irs.gov/pub/irs-prior/f1099msc--2025.pdf) | `a56d3bc3b224eb38ab4b85c793beed7eeafa4efa278a4833ac53918f746340be` |

```bash
curl -o f1099msc--2025.pdf https://www.irs.gov/pub/irs-prior/f1099msc--2025.pdf
pdftotext -layout f1099msc--2025.pdf f1099msc-2025.txt
```
