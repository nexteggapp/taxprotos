# Where this file came from

The line numbers in `mapping.textproto` are claims about a document the
Pennsylvania Department of Revenue publishes. This file is text extracted
from that document, so the claims can be checked.

Retrieved 2026-10-20 with `pdftotext` from poppler 24.02.0, in reading
order: the form and its instructions.

| File | Source | sha256 of the PDF |
|---|---|---|
| `rev-1630-2025.txt` | [REV-1630, Underpayment of Estimated Tax by Individuals, 2025](https://www.pa.gov/content/dam/copapwp-pagov/en/revenue/documents/formsandpublications/formsforindividuals/pit/documents/2025/2025_rev-1630.pdf) | `49438a97fc8026ff763bd9d234d2493f0a988b9f9be4ca34391c9187ee710112` |

```bash
curl -o rev-1630.pdf https://www.pa.gov/content/dam/copapwp-pagov/en/revenue/documents/formsandpublications/formsforindividuals/pit/documents/2025/2025_rev-1630.pdf
pdftotext rev-1630.pdf rev-1630-2025.txt
```
