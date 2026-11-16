# Where this file came from

The line numbers in `mapping.textproto` are claims about a document the
Missouri Department of Revenue publishes. This file is text extracted from
that document, so the claims can be checked.

Retrieved 2026-10-20 with `pdftotext -layout` from poppler 24.02.0: the form
and its instructions, in page layout. The extract keeps the form-feed and BEL
control characters `pdftotext` writes; the coverage test matches line numbers
set off by whitespace and does not read them.

| File | Source | sha256 of the PDF |
|---|---|---|
| `formmo2210-2025.txt` | [Form MO-2210, 2025 Underpayment of Estimated Tax by Individuals](https://dor.mo.gov/forms/MO-2210_2025.pdf) | `b17d567571d9507937a88406243f08d8c3ff84c318485e2fcfda30ba8bbca6fd` |

```bash
curl -o 2210.pdf https://dor.mo.gov/forms/MO-2210_2025.pdf
pdftotext -layout 2210.pdf formmo2210-2025.txt
```
