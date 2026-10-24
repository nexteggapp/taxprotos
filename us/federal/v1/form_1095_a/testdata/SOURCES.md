# Where these files came from

The line numbers in `mapping.textproto` are claims about documents the IRS
publishes. These files are text extracted from those documents, so the claims
can be checked.

Retrieved 2026-10-24 with `pdftotext -layout` from poppler 24.02.0, unfiltered —
each file is the whole form.

| File | Source | sha256 of the PDF |
|---|---|---|
| `f1095a-2025.txt` | [f1095a--2025.pdf](https://www.irs.gov/pub/irs-prior/f1095a--2025.pdf) | `df83be599d3bdcda826862aa5125868c2f481191a53c8417ed5fba91eccc7473` |

```bash
curl -o f1095a--2025.pdf https://www.irs.gov/pub/irs-prior/f1095a--2025.pdf
pdftotext -layout f1095a--2025.pdf f1095a-2025.txt
```
