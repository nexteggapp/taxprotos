# Where this file came from

The line numbers in `mapping.textproto` and `mapping_periods.textproto` are
claims about a document Maine Revenue Services publishes. This file is
text extracted from that document, so the claims can be checked.

Retrieved 2026-11-16 with `pdftotext -layout` from poppler 24.02.0. Maine
Revenue Services' site refused the download, so the PDF is the Internet
Archive's copy of that URL, captured 2026-02-12. The PDF holds the form and
its instructions; the extract is the form's front page, to line 17.

| File | Source | sha256 of the PDF |
|---|---|---|
| `2210me-2025.txt` | [Form 2210ME, Underpayment of Estimated Tax by Individuals, Estates, or Trusts, 2025](https://www.maine.gov/revenue/sites/maine.gov.revenue/files/inline-files/25_form_2210_fillable.pdf) | `94575b40b384f7ae552eef83b081e4cd10f68ecf06c0ebbd4d3e85883920eb35` |

```bash
curl -sS -o 2210me.pdf 'https://web.archive.org/web/20260212004453id_/https://www.maine.gov/revenue/sites/maine.gov.revenue/files/inline-files/25_form_2210_fillable.pdf'
pdftotext -layout 2210me.pdf 2210me.L.txt
awk '/FORM 2210ME INSTRUCTIONS/{exit} {print}' 2210me.L.txt > 2210me-2025.txt
```

The 2026 form is not published. The 2025 form is the latest.
