# Where this file came from

The line numbers in `mapping.textproto` are claims about a document the
Michigan Department of Treasury publishes. This file is text extracted from that
document, so the claims can be checked.

Retrieved 2026-11-12 with `pdftotext -layout` from poppler 24.02.0, unfiltered —
the file is the whole form.

| File | Source | sha256 of the PDF |
|---|---|---|
| `mi1040-2025.txt` | [2025 Form MI-1040](https://www.michigan.gov/taxes/-/media/Project/Websites/taxes/Forms/IIT/TY2025/MI-1040.pdf) | `98e6d52b88e412a5d51529bc3144bd137eb24554e30c918f63fdcf82cf9cd01c` |

```bash
curl -o MI-1040.pdf https://www.michigan.gov/taxes/-/media/Project/Websites/taxes/Forms/IIT/TY2025/MI-1040.pdf
pdftotext -layout MI-1040.pdf mi1040-2025.txt
```

## Line 9 prints a count and an amount

Lines 9a to 9d each take a number of people and the amount that number is
worth, both under the same line number. The mapping records both with a note.
