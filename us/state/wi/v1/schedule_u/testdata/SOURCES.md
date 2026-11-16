# Where this file came from

The line numbers in the mapping files are claims about a document the
Wisconsin Department of Revenue publishes. This file is text extracted from
that document, so the claims can be checked.

Retrieved 2026-11-16 with curl 8.5.0 and no special headers, and extracted
with `pdftotext -layout` from poppler 24.02.0, unfiltered: both pages of the
schedule.

| File | Source | sha256 of the PDF |
|---|---|---|
| `schedule-u-2025.txt` | [2025 Schedule U, Underpayment of Estimated Tax by Individuals, Partnerships, and Fiduciaries (D-104, R. 05-25)](https://www.revenue.wi.gov/TaxForms2025/2025-ScheduleU.pdf) | `90a843565015d1a81ff0a4a857ee7b7617cf57d5a0e51578691cc21db9500506` |

```bash
curl -sS -o schedule-u.pdf 'https://www.revenue.wi.gov/TaxForms2025/2025-ScheduleU.pdf'
pdftotext -layout schedule-u.pdf schedule-u-2025.txt
```

The 2026 schedule is not published. The 2025 schedule is the latest.
