# Where this file came from

The line numbers in `mapping.textproto` are claims about a document the
Mississippi Department of Revenue publishes. This file is text extracted from
that document, so the claims can be checked.

Retrieved 2026-10-20 with `pdftotext -layout` from poppler 24.02.0: the
worksheet, its instructions and its worked example, in page layout.

| File | Source | sha256 of the PDF |
|---|---|---|
| `form80320-2025.txt` | [Form 80-320, Individual Income Tax Interest and Penalty Worksheet (Rev. 09/25)](https://www.dor.ms.gov/sites/default/files/tax-forms/individual/80320258%201.pdf) | `9c2c621d0a9f8927f0fd71d954a12998f6420af1e1b0f49e3e308d2bdf9ddc78` |

```bash
curl -o 80-320.pdf "https://www.dor.ms.gov/sites/default/files/tax-forms/individual/80320258%201.pdf"
pdftotext -layout 80-320.pdf form80320-2025.txt
```
