# Where this file came from

The line numbers in `mapping.textproto` are claims about a document the
Louisiana Department of Revenue publishes. This file is text extracted from
that document, so the claims can be checked.

Retrieved 2026-11-12 with `pdftotext -layout` from poppler 24.02.0, unfiltered —
the file is the whole 2025 Form IT-540-WEB-BC, including its schedules
(Schedule E, Adjustments to Income, among them).

| File | Source | sha256 of the PDF |
|---|---|---|
| `it540-2025.txt` | [2025 Form IT-540-WEB-BC](https://dam.ldr.la.gov/taxforms/IT540WEB-BC(2025)%20F.pdf) | `4a500da8ec82f6cad99a059d7cf25eccafb82962a53499f260baf29238ad8627` |

```bash
curl -o it540.pdf 'https://dam.ldr.la.gov/taxforms/IT540WEB-BC(2025)%20F.pdf'
pdftotext -layout it540.pdf it540-2025.txt
```
