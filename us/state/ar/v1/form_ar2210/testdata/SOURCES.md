# Where this file came from

The line numbers in `mapping.textproto` and `mapping_part_iii.textproto` are
claims about a document the Arkansas Department of Finance and Administration
publishes. This file is text extracted from that document, so the claims can be
checked.

Retrieved 2026-10-20 with `pdftotext -layout` from poppler 24.02.0, unfiltered —
the file is the whole form.

| File | Source | sha256 of the PDF |
|---|---|---|
| `ar2210-2025.txt` | [AR2210, Penalty for Underpayment of Estimated Tax, 2025](https://www.dfa.arkansas.gov/wp-content/uploads/2025_AR2210_Penalty_for_Underpayment_of_EstimatedTax_1.pdf) (rev. 6/6/2025) | `0e1a6e610221e056461ad93845cac28799f9a9a99e2f862fd7aa438cee875938` |

```bash
curl -sS -L -o ar2210-2025.pdf https://www.dfa.arkansas.gov/wp-content/uploads/2025_AR2210_Penalty_for_Underpayment_of_EstimatedTax_1.pdf
pdftotext -layout ar2210-2025.pdf ar2210-2025.txt
```

The 2026 form is not published. The 2025 form is the latest.
