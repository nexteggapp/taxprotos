# Where this file came from

The line numbers in `mapping.textproto` and `mapping_part_iii.textproto` are
claims about a document the California Franchise Tax Board publishes. This file
is text extracted from that document, so the claims can be checked.

Retrieved 2026-10-20 with `pdftotext -layout` from poppler 24.02.0, unfiltered —
the file is the whole form.

| File | Source | sha256 of the PDF |
|---|---|---|
| `ftb5805-2025.txt` | [2025 Form FTB 5805](https://www.ftb.ca.gov/forms/2025/2025-5805.pdf) | `cf9a00556393ef16056d2f84dd2d579f845a8f8978580c75c5df4451fd675467` |

```bash
curl -o 2025-5805.pdf https://www.ftb.ca.gov/forms/2025/2025-5805.pdf
pdftotext -layout 2025-5805.pdf ftb5805-2025.txt
```

## What is not here

The instructions, which print Worksheet I (Prorated Tax for Form 540NR Filers)
and Worksheet II (Regular Method to Figure Your Underpayment and Penalty), are
an HTML page, [2025 Instructions for Form FTB
5805](https://www.ftb.ca.gov/forms/2025/2025-5805-instructions.html). Their
lines are lettered (A to P) or numbered with a trailing period, which the
coverage test's line search does not match, so the worksheets are described in
the schema's comments and not mapped.

The 2026 form is not published. The 2025 form is the latest.
