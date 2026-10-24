# Where these files came from

The line numbers in `mapping.textproto` are claims about documents the IRS
publishes. These files are text extracted from those documents, so the claims
can be checked.

Retrieved 2026-10-24 with `pdftotext -layout` from poppler 24.02.0, unfiltered —
each file is the whole form.

| File | Source | sha256 of the PDF |
|---|---|---|
| `f1099c-2025.txt` | [f1099c--2025.pdf](https://www.irs.gov/pub/irs-prior/f1099c--2025.pdf) | `826fe6811249e38b00df962b323c7ef838bc1e94a3c4522f9e10be4887f096a8` |

```bash
curl -o f1099c--2025.pdf https://www.irs.gov/pub/irs-prior/f1099c--2025.pdf
pdftotext -layout f1099c--2025.pdf f1099c-2025.txt
```
