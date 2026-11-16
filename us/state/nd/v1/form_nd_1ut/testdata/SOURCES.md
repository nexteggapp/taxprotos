# Where this file came from

The line numbers in `mapping.textproto` are claims about a document the
North Dakota Office of State Tax Commissioner publishes. This file is text
extracted from that document, so the claims can be checked.

Retrieved 2026-10-20 with `pdftotext` from poppler 24.02.0, in reading
order: the schedule and its instructions. The office lists the 2025
schedule as the latest.

| File | Source | sha256 of the PDF |
|---|---|---|
| `formnd1ut-2025.txt` | [Schedule ND-1UT, Underpayment or Late Payment of Estimated Income Tax for Individuals, SFN 28704 (12-2025)](https://www.tax.nd.gov/sites/www/files/documents/forms/individual/2025-iit/28704-schedule-nd-1ut-2025.pdf) | `3d6950869900148794caa94e612b701320e7efd83923ca24e0ecf12deaac5529` |

```bash
curl -o ut.pdf https://www.tax.nd.gov/sites/www/files/documents/forms/individual/2025-iit/28704-schedule-nd-1ut-2025.pdf
pdftotext ut.pdf formnd1ut-2025.txt
```
