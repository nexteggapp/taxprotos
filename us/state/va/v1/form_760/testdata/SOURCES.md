# Where this file came from

The line numbers in `mapping.textproto` are claims about a document the
Virginia Department of Taxation publishes. This file is text extracted from
that document, so the claims can be checked.

Retrieved 2026-11-12 with `pdftotext -layout` from poppler 24.02.0, unfiltered —
the file is the whole form.

| File | Source | sha256 of the PDF |
|---|---|---|
| `760-2025.txt` | [2025 Form 760](https://www.tax.virginia.gov/sites/default/files/taxforms/individual-income-tax/2025/760-2025.pdf) | `bb271a15f9ad173ed900bb23dc7d154c0cb2056bdd05957fb25e5faedbc27575` |

```bash
curl -o 760-2025.pdf https://www.tax.virginia.gov/sites/default/files/taxforms/individual-income-tax/2025/760-2025.pdf
pdftotext -layout 760-2025.pdf 760-2025.txt
```
