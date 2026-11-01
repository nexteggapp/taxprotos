# Where these files came from

The line numbers in `mapping.textproto` are claims about documents the IRS
publishes. These files are text extracted from those documents, so the claims
can be checked.

Retrieved 2026-11-01 with `pdftotext -layout` from poppler 24.02.0, unfiltered —
each file is the whole form.

| File | Source | sha256 of the PDF |
|---|---|---|
| `f3922-2025.txt` | [f3922--2025.pdf](https://www.irs.gov/pub/irs-prior/f3922--2025.pdf) | `98077229f4efb9d964caedd171847e6d6b8ae4b0a3cce60089410e9fb47f7164` |

```bash
curl -o f3922--2025.pdf https://www.irs.gov/pub/irs-prior/f3922--2025.pdf
pdftotext -layout f3922--2025.pdf f3922-2025.txt
```
