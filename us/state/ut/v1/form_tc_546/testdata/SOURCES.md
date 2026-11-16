# Where this file came from

The line numbers in `mapping.textproto` are claims about a document the Utah
State Tax Commission publishes. This file is text extracted from that
document, so the claims can be checked.

Retrieved 2026-10-20 with `pdftotext -layout` from poppler 24.02.0: the
Payment Worksheet of TC-546, which is set in the right-hand column of a
two-column page, so it is cut from column 77 and bounded by its own first and
last lines. The PDF prints a control character after the dash on line 5,
which the extract keeps.

| File | Source | sha256 of the PDF |
|---|---|---|
| `tc-546-worksheet.txt` | [TC-546, Individual Income Tax Prepayment Coupon (Rev. 10/26)](https://files.tax.utah.gov/tax/forms/current/tc-546.pdf) | `8c3fa7cd582f6f870a622188ff4d9356e4b1b548bf3f797b54589a2413ff2f80` |

```bash
curl -L -o tc-546.pdf https://files.tax.utah.gov/tax/forms/current/tc-546.pdf
pdftotext -layout tc-546.pdf - | cut -c77- \
  | awk '/Payment Worksheet/{f=1} f; /subtract line 6 and line 7 from line 5/{if(f)exit}' > tc-546-worksheet.txt
```

The extension penalty's steps follow [Publication 58, Utah Interest and
Penalties (Rev. 11/25)](https://files.tax.utah.gov/tax/forms/pubs/pub-58.pdf),
Insufficient Prepayment (Extension Penalty), which numbers no lines.
