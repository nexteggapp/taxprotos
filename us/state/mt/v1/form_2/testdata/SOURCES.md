# Where this file came from

The line numbers in `mapping.textproto` are claims about a document the
Montana Department of Revenue publishes. This file is text extracted from that
document, so the claims can be checked.

Retrieved 2026-11-12 with curl and extracted with `pdftotext -layout` from
poppler 24.02.0, unfiltered: the file is the whole form, all schedules
included.

| File | Source | sha256 of the PDF |
|---|---|---|
| `form2-2025.txt` | [2025 Montana Form 2 (2025v3, 12/2025)](https://revenuefiles.mt.gov/files/Forms/Montana-Individual-Income-Tax-Return-Form-2/2025_Montana_Individual_Income_Tax_Return_Form_2.pdf) | `4671aef667749732670c07fd51cbeed6bc824dc6466cfe19597787b3d6ba2b37` |

```bash
curl -o form2.pdf https://revenuefiles.mt.gov/files/Forms/Montana-Individual-Income-Tax-Return-Form-2/2025_Montana_Individual_Income_Tax_Return_Form_2.pdf
pdftotext -layout form2.pdf form2-2025.txt
```

## The worksheet on page 2

Page 2 carries the resident tax worksheet, whose lines are numbered 1 to 13
again. The mapping notes, for every concept, whether its line is on page 1 or
the worksheet.
