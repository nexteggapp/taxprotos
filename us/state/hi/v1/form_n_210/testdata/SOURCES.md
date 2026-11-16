# Where this file came from

The line numbers in `mapping.textproto` are claims about a document the
Hawaii Department of Taxation publishes. This file is text extracted from
that document, so the claims can be checked.

Retrieved 2026-10-20 with curl 8.5.0 and no special headers, and extracted
with `pdftotext -layout` from poppler 24.02.0, unfiltered: the file is the
whole form, both pages. Every page has extractable text.

| File | Source | sha256 of the PDF |
|---|---|---|
| `n-210-2025.txt` | [Form N-210, Underpayment of Estimated Tax by Individuals, Estates, and Trusts (Rev. 2025)](https://files.hawaii.gov/tax/forms/2025/n210_i.pdf) | `61511c4353a2024caea1fcca29931b39ea6dcc5b3bf29c557399b64d6a6c03cc` |

```bash
curl -O https://files.hawaii.gov/tax/forms/2025/n210_i.pdf
pdftotext -layout n210_i.pdf n-210-2025.txt
```

The instructions, which the field comments also follow (lines 10, 18 and 19
in particular), are
[Instructions for Form N-210 (Rev. 2025)](https://files.hawaii.gov/tax/forms/2025/n210ins.pdf).
