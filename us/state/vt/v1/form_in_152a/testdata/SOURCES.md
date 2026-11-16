# Where this file came from

The line numbers in `mapping.textproto` and `mapping_periods.textproto` are
claims about a document the Vermont Department of Taxes publishes. This file is
text extracted from that document, so the claims can be checked.

Retrieved 2026-11-16 with `pdftotext -layout` from poppler 24.02.0. The
Department's site answered a plain download with HTTP 403 ("Request
blocked."), so the PDF is the Internet Archive's copy of the published URL,
captured 2026-02-03. The extract is the whole worksheet, one page.

| File | Source | sha256 of the PDF |
|---|---|---|
| `in152a-2025.txt` | [Worksheet IN-152A, Annualized Income Installment Method for Underpayment of 2025 Estimated Tax by Individuals, Estates, and Trusts (Rev. 10/25)](https://tax.vermont.gov/sites/tax/files/documents/IN-152A-2025.pdf) | `fd76a50808402311cb39a46e8b77049f1b2db6f4075e2acb84f20e46c1d40d3a` |

```bash
curl -sS -L -o in152a.pdf 'https://web.archive.org/web/20260203213903id_/https://tax.vermont.gov/sites/tax/files/documents/IN-152A-2025.pdf'
pdftotext -layout in152a.pdf in152a-2025.txt
```

The 2026 worksheet is not published. The 2025 worksheet is the latest.
