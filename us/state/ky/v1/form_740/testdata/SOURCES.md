# Where this file came from

The line numbers in `mapping.textproto` are claims about a document the
Kentucky Department of Revenue publishes. These files are text extracted from
those documents, so the claims can be checked.

Retrieved 2026-11-12 with `pdftotext -layout` from poppler 24.02.0, unfiltered —
each file is the whole document.

| File | Source | sha256 of the PDF |
|---|---|---|
| `740-2025.txt` | [2025 Form 740](https://revenue.ky.gov/Forms/740%20(2025).pdf) | `cb68242eb8351dcd909425cd8edc032ef3169f7eaa86f17fa3ac6c1cb4136ca2` |
| `schedule-itc-2025.txt` | [2025 Schedule ITC and its instructions](https://revenue.ky.gov/Forms/Schedule%20ITC%20(2025).pdf) | `48b763b31d57bd601b475633a7f6780983311070ec107129b1fd13abfc98919b` |

```bash
curl -o 740-2025.pdf 'https://revenue.ky.gov/Forms/740%20(2025).pdf'
pdftotext -layout 740-2025.pdf 740-2025.txt
curl -o itc-2025.pdf 'https://revenue.ky.gov/Forms/Schedule%20ITC%20(2025).pdf'
pdftotext -layout itc-2025.pdf schedule-itc-2025.txt
```

The coverage test checks the mapping against `740-2025.txt`. Schedule ITC is
kept because it prints the personal tax credits (Section B) and the family size
tax credit table (Section C) that lines 17 and 21 take their figures from.

## Line numbers set off by an en space

The extract sets lines 31b, 31c and 31h off from their leaders with an en space
(U+2002) rather than an ASCII space, so the coverage test counts any Unicode
space as a separator.
