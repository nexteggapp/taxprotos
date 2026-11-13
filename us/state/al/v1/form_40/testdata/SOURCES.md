# Where this file came from

The line numbers in `mapping.textproto` are claims about a document the
Alabama Department of Revenue publishes. This file is text extracted from that
document, so the claims can be checked.

Retrieved 2026-11-12 with `pdftotext -layout` from poppler 24.02.0, unfiltered —
the file is the whole form.

| File | Source | sha256 of the PDF |
|---|---|---|
| `form40-2025.txt` | [2025 Form 40](https://www.revenue.alabama.gov/wp-content/uploads/2026/01/25f40blk.pdf) | `c001efc5771c1c2ce1da95e57c6a4c4e82cdf529e75024adc6be6cc81927cfe8` |

```bash
curl -o 25f40blk.pdf https://www.revenue.alabama.gov/wp-content/uploads/2026/01/25f40blk.pdf
pdftotext -layout 25f40blk.pdf form40-2025.txt
```

## Page 2 restarts its line numbers

Page 2 numbers Parts I to III from 1 again. Only Part III, line 1 (the number
of dependents) is mapped from page 2, with a note; Parts I and II are carried
as their totals on page 1, lines 7 and 9.
