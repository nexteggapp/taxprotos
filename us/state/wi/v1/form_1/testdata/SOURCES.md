# Where this file came from

The line numbers in `mapping.textproto` are claims about a document the
Wisconsin Department of Revenue publishes. This file is text extracted from that
document, so the claims can be checked.

Retrieved 2026-11-14 with `pdftotext -layout` from poppler 24.02.0, unfiltered —
the file is the whole form.

| File | Source | sha256 of the PDF |
|---|---|---|
| `form1-2025.txt` | [2025 Form 1](https://www.revenue.wi.gov/TaxForms2025/2025-Form1f.pdf) | `a9b0d9b75bcb55ec0e20984c28872a893765fea66e0f1edd9a7c47b60ceae14e` |

```bash
curl -o 2025-Form1f.pdf https://www.revenue.wi.gov/TaxForms2025/2025-Form1f.pdf
pdftotext -layout 2025-Form1f.pdf form1-2025.txt
```
