# Where this file came from

The line numbers in `mapping.textproto` are claims about a document the New
Mexico Taxation and Revenue Department publishes. This file is text extracted
from that document, so the claims can be checked.

Retrieved 2026-11-12 with `pdftotext -layout` from poppler 24.02.0, unfiltered,
the file is the whole form.

| File | Source | sha256 of the PDF |
|---|---|---|
| `pit1-2025.txt` | [2025 Form PIT-1](https://realfile.tax.newmexico.gov/2025pit-1.pdf) | `70a37142468c02661f871cb53030a86fa3666492bfb94a09ea8fa801f2a08e03` |

```bash
curl -o 2025pit-1.pdf https://realfile.tax.newmexico.gov/2025pit-1.pdf
pdftotext -layout 2025pit-1.pdf pit1-2025.txt
```

The worksheets for lines 13 and 14 are in the
[2025 PIT-1 instructions](https://realfile.tax.newmexico.gov/2025pit-1-ins.pdf),
sha256 `7d578286a047513ef579659b3683cca5818e21e64448e6b442ae217dce27067a`.
