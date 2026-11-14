# Where this file came from

The line numbers in `mapping.textproto` are claims about a document the
District of Columbia Office of Tax and Revenue publishes. This file is text
extracted from that document, so the claims can be checked.

Retrieved 2026-11-14 with curl 8.5.0 and extracted with `pdftotext -layout`
from poppler 24.02.0. The form is printed inside the 2025 D-40 booklet, on its
pages 42 to 44; only those pages are extracted. The form's own text uses a
custom font encoding, so its instructions in small print come out as
unreadable characters. The line numbers and line captions extract cleanly,
and they are what the mapping is checked against.

| File | Source | sha256 of the PDF |
|---|---|---|
| `d40-2025.txt` | [2025 D-40 Booklet](https://otr.cfo.dc.gov/sites/default/files/dc/sites/otr/publication/attachments/2025_D40_Book_082026_v1.pdf), the form revised 11/2025 | `6c1df426c2a6b4cdbe5182f7b83b3f9c0cf31011ce177f1ce404ed566d1754d0` |

```bash
curl -o d40.pdf https://otr.cfo.dc.gov/sites/default/files/dc/sites/otr/publication/attachments/2025_D40_Book_082026_v1.pdf
pdftotext -f 42 -l 44 -layout d40.pdf d40-2025.txt
```
