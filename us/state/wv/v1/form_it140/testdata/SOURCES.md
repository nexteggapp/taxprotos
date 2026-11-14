# Where this file came from

The line numbers in `mapping.textproto` are claims about a document the
West Virginia Tax Division publishes. This file is text extracted from that
document, so the claims can be checked.

The form is printed inside the 2025 Personal Income Tax Forms and Instructions
booklet. Retrieved 2026-11-14 with `pdftotext -layout` from poppler 24.02.0,
pages 3 to 6 of the booklet only: Form IT-140 pages 1 and 2 and Schedule M.

| File | Source | sha256 of the PDF |
|---|---|---|
| `it140-2025.txt` | [2025 Personal Income Tax Forms and Instructions](https://tax.wv.gov/Documents/PIT/2025/it140.PersonalIncomeTaxFormsAndInstructions.2025.pdf) | `6383bf9748fde18c0469b0eb1b836f4ff6999a2c541bf4e909502a8b3132e0b0` |

```bash
curl -o it140.pdf https://tax.wv.gov/Documents/PIT/2025/it140.PersonalIncomeTaxFormsAndInstructions.2025.pdf
pdftotext -f 3 -l 6 -layout it140.pdf it140-2025.txt
```

Form IT-140 lines 2 and 3 refer to "line 61" and "line 52" of Schedule M, but
the 2025 Schedule M prints its totals on lines 59 (additions) and 50
(subtractions).
