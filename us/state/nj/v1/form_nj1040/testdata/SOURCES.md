# Where this file came from

The line numbers in `mapping.textproto` are claims about a document the
New Jersey Division of Taxation publishes. This file is text extracted from that
document, so the claims can be checked.

Retrieved 2026-11-14 with `pdftotext -layout` from poppler 24.02.0, unfiltered —
the file is the whole form.

| File | Source | sha256 of the PDF |
|---|---|---|
| `nj1040-2025.txt` | [2025 Form NJ-1040](https://www.nj.gov/treasury/taxation/pdf/current/1040.pdf) | `3d15ad99f3ced62a38eb53c25bd53830a0b9e3927abcb75030ef06ec4c7eb6a3` |

```bash
curl -o 1040.pdf https://www.nj.gov/treasury/taxation/pdf/current/1040.pdf
pdftotext -layout 1040.pdf nj1040-2025.txt
```

The URL is the division's "current" copy, which will become the 2026 form. The
sha256 identifies the 2025 one.

## Line 37 is printed in parts

The NJBEST, NJCLASS and higher education tuition deductions are printed as
"37a", with "b" and "c" set beside it, so 37b and 37c never appear as such. The
mapping records them as line 37a with a note.
