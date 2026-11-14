# Where this file came from

The line numbers in `mapping.textproto` are claims about a document the
Minnesota Department of Revenue publishes. This file is text extracted from
that document, so the claims can be checked.

Retrieved 2026-11-14 with `pdftotext -layout` from poppler 24.02.0, unfiltered —
the file is the whole form.

| File | Source | sha256 of the PDF |
|---|---|---|
| `m1-2025.txt` | [2025 Form M1](https://www.revenue.state.mn.us/sites/default/files/2026-03/m1-25.pdf) | `6207b7b0ecbcf5d87effd4ead8f3c6e326b6807997b9b0ffe79bb755d115f268` |

```bash
curl -o m1-25.pdf https://www.revenue.state.mn.us/sites/default/files/2026-03/m1-25.pdf
pdftotext -layout m1-25.pdf m1-2025.txt
```

## Some line numbers are printed with a space

The right-hand column prints some line numbers as "1 0", "1 4 a" and so on.
Each of those lines is also printed whole at the left margin, which is where the
coverage test finds it.
