# Where this file came from

The line numbers in the `mapping*.textproto` files are claims about a document
the Arkansas Department of Finance and Administration publishes. This file is
text extracted from that document, so the claims can be checked.

Retrieved 2026-10-20 with `pdftotext -layout` from poppler 24.02.0, unfiltered —
the file is the whole form.

| File | Source | sha256 of the PDF |
|---|---|---|
| `ar2210a-2025.txt` | [AR2210A, Annualized Penalty for Underpayment of Estimated Income Tax, 2025](https://www.dfa.arkansas.gov/wp-content/uploads/2025_AR2210A_AnnualizedPenalty_for_Underpayment_of_EstimatedTax.pdf) (rev. 4/28/2025) | `964dddfe459f449f2a60006a614fcccd798b4f1cffb170bae80bb68184b56093` |

```bash
curl -sS -L -o ar2210a-2025.pdf https://www.dfa.arkansas.gov/wp-content/uploads/2025_AR2210A_AnnualizedPenalty_for_Underpayment_of_EstimatedTax.pdf
pdftotext -layout ar2210a-2025.pdf ar2210a-2025.txt
```

Some labels on page 1 are set in a font whose characters extract shifted
("LIÀOLQJVWDWXV" for "filing status"); the line numbers, which the mappings
rely on, extract cleanly. The text is left as `pdftotext` writes it.

The 2026 form is not published. The 2025 form is the latest.
