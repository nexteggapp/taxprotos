# Where this file came from

The line numbers in `mapping.textproto` are claims about a document the
Missouri Department of Revenue publishes. This file is text extracted from that
document, so the claims can be checked.

Retrieved 2026-11-12 with `pdftotext -layout` from poppler 24.02.0. The PDF is
the department's "fillable calculating" booklet, 32 pages, which also carries
Form MO-A and other schedules; the extract keeps pages 1 to 5, which hold
Form MO-1040 itself.

| File | Source | sha256 of the PDF |
|---|---|---|
| `mo1040-2025.txt` | [2025 Form MO-1040](https://dor.mo.gov/forms/MO-1040%20Fillable%20Calculating_2025.pdf) | `912b43015ea273d93b0c25fe6268d1415496fbadbb048a97783849c5dd2a8c91` |

```bash
curl -o mo-1040.pdf 'https://dor.mo.gov/forms/MO-1040%20Fillable%20Calculating_2025.pdf'
pdftotext -layout -f 1 -l 5 mo-1040.pdf mo1040-2025.txt
```

The figures the rules use (the tax chart, the standard deductions, the federal
tax percentage chart) come from the 2025 Form MO-1040 instructions,
`https://dor.mo.gov/forms/MO-1040%20Instructions_2025.pdf`, sha256
`99b1d607e1c9c7623ecf3cd7a8f6fba2a732861dcb54164cc7e7e3a82c5df540`.
