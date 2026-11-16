# Where this file came from

The line numbers in `mapping.textproto` are claims about a document the
Delaware Division of Revenue publishes. This file is text extracted from that
document, so the claims can be checked.

Retrieved 2026-10-20 with curl 8.5.0 and no special headers, and extracted
with `pdftotext -layout` from poppler 24.02.0, unfiltered: the file is the
whole form, all three pages. Every page has extractable text. The form is a
grid whose columns reading order scatters, so it is extracted with its
layout.

| File | Source | sha256 of the PDF |
|---|---|---|
| `pit-und-2025.txt` | [2025 Form PIT-UND, Delaware Underpayment of Estimated Taxes (revision 20250724)](https://revenuefiles.delaware.gov/2025/PITForms_Instructions/PIT-UND_2025-01_PaperInteractiveIPM.pdf) | `a485b8212e24a1b0faee8f5ffba2a3d7245d66cc92ab592d31084b6a9baf6654` |

```bash
curl -O https://revenuefiles.delaware.gov/2025/PITForms_Instructions/PIT-UND_2025-01_PaperInteractiveIPM.pdf
pdftotext -layout PIT-UND_2025-01_PaperInteractiveIPM.pdf pit-und-2025.txt
```

The 2026 form was not published when this was retrieved. The instructions,
which the field comments also follow (lines 29, 30 and 41 in particular), are
[PIT-UND Instructions, 2025 (revised 11/10/25)](https://revenuefiles.delaware.gov/2025/PITForms_Instructions/Instructions/PIT-UND_Instructions_2025-01.pdf).
