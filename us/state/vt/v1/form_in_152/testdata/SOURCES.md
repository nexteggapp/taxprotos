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
| `in152-2025.txt` | [Worksheet IN-152, Underpayment of 2025 Estimated Individual Income Tax (Rev. 10/25)](https://tax.vermont.gov/sites/tax/files/documents/IN-152-2025.pdf) | `7f0002794ea260211349d80e6dde47574d179eb99802b0217ad47730f3fe2979` |

```bash
curl -sS -L -o in152.pdf 'https://web.archive.org/web/20260203203854id_/https://tax.vermont.gov/sites/tax/files/documents/IN-152-2025.pdf'
pdftotext -layout in152.pdf in152-2025.txt
```

The 2026 worksheet is not published. The 2025 worksheet is the latest.
