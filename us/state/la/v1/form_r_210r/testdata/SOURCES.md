# Where this file came from

The line numbers in `mapping.textproto` and `mapping_period.textproto` are
claims about a document the Louisiana Department of Revenue publishes. This
file is text extracted from that document, so the claims can be checked.

Retrieved 2026-11-16 with curl 8.5.0 and no special headers, and extracted
with `pdftotext -layout` from poppler 24.02.0, unfiltered: the file is the
whole form, one page.

| File | Source | sha256 of the PDF |
|---|---|---|
| `r210r-2025.txt` | [Form R-210R, Underpayment of Individual Income Tax Penalty Computation, 2025 Taxable Year, Resident Filers](https://dam.ldr.la.gov/taxforms/210R-1-26-F.pdf), R-210R (1/26) | `b8f6908005857c03c6d8ae66f169bb5bae7f449b6aaa5131ef1df59f37f63209` |

```bash
curl -sS -L -o r210r.pdf https://dam.ldr.la.gov/taxforms/210R-1-26-F.pdf
pdftotext -layout r210r.pdf r210r-2025.txt
```

The field comments also follow the form's instructions, R-210R-i (1/26),
<https://dam.ldr.la.gov/taxforms/210Ri-1-26.pdf>. The 2026 form is not
published. The 2025 form is the latest.
