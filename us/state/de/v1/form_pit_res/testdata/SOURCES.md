# Where this file came from

The line numbers in `mapping.textproto` are claims about a document the
Delaware Division of Revenue publishes. This file is text extracted from that
document, so the claims can be checked.

Retrieved 2026-11-13 with curl 8.5.0 and no special headers, and extracted
with `pdftotext -layout` from poppler 24.02.0, unfiltered: the file is the
whole form, all four pages. Every page has extractable text.

| File | Source | sha256 of the PDF |
|---|---|---|
| `pit-res-2025.txt` | [2025 Form PIT-RES (revision 20260407)](https://revenuefiles.delaware.gov/2025/PITForms_Instructions/PIT-RES_2025-01_PaperInteractiveIPM.pdf) | `7b7ae51e2ede914167ced0c7278c4896e02612f40913545aa722befe078145da` |

```bash
curl -O https://revenuefiles.delaware.gov/2025/PITForms_Instructions/PIT-RES_2025-01_PaperInteractiveIPM.pdf
pdftotext -layout PIT-RES_2025-01_PaperInteractiveIPM.pdf pit-res-2025.txt
```

The same server rejected requests sent with a browser's User-Agent header,
answering with an HTML "Request Rejected" page. Plain curl worked.
