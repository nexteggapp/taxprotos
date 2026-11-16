# Where this file came from

The line numbers in `mapping.textproto` and `mapping_periods.textproto` are
claims about a document the Comptroller of Maryland publishes. This file is
text extracted from that document, so the claims can be checked.

Retrieved 2026-11-16 with `pdftotext -layout` from poppler 24.02.0. The PDF
holds the form and its instructions; the extract is the form's front page,
from its heading to its revision line.

| File | Source | sha256 of the PDF |
|---|---|---|
| `502up-2025.txt` | [Form 502UP, Underpayment of Estimated Income Tax by Individuals, 2025](https://www.marylandcomptroller.gov/content/dam/mdcomp/tax/forms/2025/502up.pdf) (COM/RAD 017 09/25) | `4f2b10661a071fe8aa1d4329da5d08a20c00554c40f7ae6fbc9fa4a0694e8510` |

```bash
curl -sS -L -o 502up.pdf https://www.marylandcomptroller.gov/content/dam/mdcomp/tax/forms/2025/502up.pdf
pdftotext -layout 502up.pdf 502up.L.txt
awk '{print} /COM\/RAD 017/{exit}' 502up.L.txt > 502up-2025.txt
```

The 2026 form is not published. The 2025 form is the latest.
