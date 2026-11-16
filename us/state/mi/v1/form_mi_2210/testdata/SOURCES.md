# Where this file came from

The line numbers in `mapping.textproto` and `mapping_periods.textproto` are
claims about a document the Michigan Department of Treasury publishes. This file is
text extracted from that document, so the claims can be checked.

Retrieved 2026-11-16 with `pdftotext -layout` from poppler 24.02.0. Treasury's
site refused the download, so the PDF is the Internet Archive's copy of the
published URL (with Treasury's rev and hash query), captured 2026-05-17. The
PDF holds the form, its Annualized Income Worksheet and its instructions;
the extract is the form and the worksheet, pages 1 and 2.

| File | Source | sha256 of the PDF |
|---|---|---|
| `mi-2210-2025.txt` | [2025 MI-2210, Underpayment of Estimated Income Tax (Rev. 10-25)](https://www.michigan.gov/taxes/-/media/Project/Websites/taxes/Forms/IIT/TY2025/MI-2210.pdf) | `6df61c17b35372ed1b8f58d0bd9e1b3fdd4effab72ffb9e8359b65842ecd85dd` |

```bash
curl -sS -o mi-2210.pdf 'https://web.archive.org/web/20260517063622id_/https://www.michigan.gov/taxes/-/media/Project/Websites/taxes/Forms/IIT/TY2025/MI-2210.pdf?rev=4c4095ff482347d4a68311914f0c543e&hash=F5A4BC24313FBA31F6292E2ADA75CB31'
pdftotext -layout mi-2210.pdf mi-2210.L.txt
awk '/MI-2210, Page 3/{exit} {print}' mi-2210.L.txt > mi-2210-2025.txt
```

The 2026 form is not published. The 2025 form is the latest.
