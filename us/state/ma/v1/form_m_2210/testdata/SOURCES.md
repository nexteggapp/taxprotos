# Where this file came from

The line numbers in `mapping.textproto`, `mapping_installment.textproto` and
`mapping_period.textproto` are claims about a document the Massachusetts
Department of Revenue publishes. This file is text extracted from that
document, so the claims can be checked.

Retrieved 2026-11-16 with curl 8.5.0 and no special headers, and extracted
with `pdftotext -layout` from poppler 24.02.0, unfiltered: the file is the
whole form and its instructions, all four pages.

| File | Source | sha256 of the PDF |
|---|---|---|
| `m-2210-2025.txt` | [2025 Form M-2210, Underpayment of Massachusetts Estimated Income Tax](https://www.mass.gov/doc/2025-form-m-2210-underpayment-of-massachusetts-estimated-income-tax/download) | `96fd1bd71c8926847f2c95dc74b7bad5fd50155e4e12bbb88cad1f0855fcbaeb` |

```bash
curl -sS -L -o m2210.pdf 'https://www.mass.gov/doc/2025-form-m-2210-underpayment-of-massachusetts-estimated-income-tax/download'
pdftotext -layout m2210.pdf m-2210-2025.txt
```

A browser's user agent gets an HTML page from the same address; curl's own
gets the PDF. The final 2026 form is not published; the draft 2026 form
prints the same lines with the 2026 dates and rates.
