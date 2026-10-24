# Where these files came from

The line numbers in `mapping.textproto` are claims about documents the IRS
publishes. These files are text extracted from those documents, so the claims
can be checked.

Retrieved 2026-10-24 with `pdftotext -layout` from poppler 24.02.0, unfiltered —
each file is the whole form.

| File | Source | sha256 of the PDF |
|---|---|---|
| `f1098e-2025.txt` | [f1098e.pdf](https://www.irs.gov/pub/irs-pdf/f1098e.pdf) | `bffb9b5ec4f06d47391a5fa7b773d1c4197c587428ace03a96914a9038389602` |

```bash
curl -o f1098e.pdf https://www.irs.gov/pub/irs-pdf/f1098e.pdf
pdftotext -layout f1098e.pdf f1098e-2025.txt
```
