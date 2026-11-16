# Where this file came from

The line numbers in `mapping.textproto` are claims about a document the
Virginia Department of Taxation publishes. This file is text extracted from
that document, so the claims can be checked.

Retrieved 2026-10-20 with `pdftotext -layout` from poppler 24.02.0: the form,
in page layout.

| File | Source | sha256 of the PDF |
|---|---|---|
| `760c-2025.txt` | [Form 760C, Underpayment of Virginia Estimated Tax by Individuals, Estates and Trusts, 2025 (Rev. 04/26)](https://www.tax.virginia.gov/sites/default/files/taxforms/individual-income-tax/2025/760c-2025.pdf) | `a917d491de2c55a3029bac3c625c3de4f226b0a3046bb75fd4adc392fb676f0b` |

```bash
curl -o 760c-2025.pdf https://www.tax.virginia.gov/sites/default/files/taxforms/individual-income-tax/2025/760c-2025.pdf
pdftotext -layout 760c-2025.pdf 760c-2025.txt
```
