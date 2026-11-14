# Where this file came from

The line numbers in `mapping.textproto` are claims about a document the
Comptroller of Maryland publishes. This file is text extracted from that
document, so the claims can be checked.

Retrieved 2026-11-14 with `pdftotext -layout` from poppler 24.02.0, unfiltered —
the file is the whole form.

| File | Source | sha256 of the PDF |
|---|---|---|
| `502-2025.txt` | [2025 Form 502](https://www.marylandcomptroller.gov/content/dam/mdcomp/tax/forms/2025/502.pdf) | `4f4d604f27cf6995cbab58b7650fd812c0e4ccc9f8d899f04261ded5a04a24cf` |

```bash
curl -o 502.pdf https://www.marylandcomptroller.gov/content/dam/mdcomp/tax/forms/2025/502.pdf
pdftotext -layout 502.pdf 502-2025.txt
```

## The exemptions are lettered

The exemption boxes on page 2 are lettered A to D rather than numbered; the
mapping uses the letters.
