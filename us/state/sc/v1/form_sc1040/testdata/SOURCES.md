# Where this file came from

The line numbers in `mapping.textproto` are claims about a document the
South Carolina Department of Revenue publishes. This file is text extracted
from that document, so the claims can be checked.

Retrieved 2026-11-14 with `pdftotext -layout` from poppler 24.02.0, unfiltered —
the file is the whole form.

| File | Source | sha256 of the PDF |
|---|---|---|
| `sc1040-2025.txt` | [2025 SC1040](https://dor.sc.gov/sites/dor/files/forms/SC1040_2025.pdf) | `f1cc1b2edf2143fabe5f4e5a1b19226fa2b6029c61e9a833309a92c3d963770a` |

```bash
curl -o SC1040_2025.pdf https://dor.sc.gov/sites/dor/files/forms/SC1040_2025.pdf
pdftotext -layout SC1040_2025.pdf sc1040-2025.txt
```
