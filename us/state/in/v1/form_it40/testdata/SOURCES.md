# Where this file came from

The line numbers in `mapping.textproto` are claims about a document the
Indiana Department of Revenue publishes. This file is text extracted from that
document, so the claims can be checked.

Retrieved 2026-11-12 with `pdftotext -layout` from poppler 24.02.0, unfiltered —
the file is the whole form.

| File | Source | sha256 of the PDF |
|---|---|---|
| `it40-2025.txt` | [2025 Form IT-40, State Form 154 (R24 / 9-25)](https://forms.in.gov/Download.aspx?id=16914) | `0b5d74cffea1dc758eff500a8d74bf020b53b977b57e04b12ce344316234b785` |

```bash
curl -L -o it40.pdf 'https://forms.in.gov/Download.aspx?id=16914'
pdftotext -layout it40.pdf it40-2025.txt
```

The link is the one the department's current individual income tax forms page
gives for "IT-40 Form"; the PDF it serves is headed 2025.
