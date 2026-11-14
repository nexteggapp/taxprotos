# Where this file came from

The line numbers in `mapping.textproto` are claims about a document the Rhode
Island Division of Taxation publishes. This file is text extracted from that
document, so the claims can be checked.

The form itself could not be fetched: tax.ri.gov's HTML pages, including the
forms index that links the 2025 RI-1040, answer a script with HTTP 403, so the
form's own URL could not be found. The resident instructions, whose URL is
known, answered with HTTP 200. They describe every line of the RI-1040 by
number ("Line 14a – Rhode Island Income Tax Withheld"), and that is what the
mapping is checked against.

Retrieved 2026-11-14 with curl 8.5.0 and extracted with `pdftotext -layout`
from poppler 24.02.0, unfiltered — the file is the whole instructions.

| File | Source | sha256 of the PDF |
|---|---|---|
| `ri1040-2025-instructions.txt` | [2025 RI-1040 instructions](https://tax.ri.gov/sites/g/files/xkgbur541/files/2025-12/2025%201040R%20Instructions%20122025.pdf) (December 2025) | `17c4d5abe85fe61e69da5daa349865576b5a66425a147dd5686a922b0c208336` |

```bash
curl -o 1040r.pdf "https://tax.ri.gov/sites/g/files/xkgbur541/files/2025-12/2025%201040R%20Instructions%20122025.pdf"
pdftotext -layout 1040r.pdf ri1040-2025-instructions.txt
```

## The filing status

The instructions head that section "FILING STATUS" rather than giving it a
line number, and the mapping uses the heading.

## The exemption worksheet

The instructions send filers whose line 3 exceeds $254,250 to an Exemption
Worksheet "on the inside of the back cover" of the resident booklet, which is
not part of this file.
