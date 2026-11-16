# Where this file came from

The line numbers in `mapping.textproto` and `mapping_periods.textproto` are
claims about a document the District of Columbia Office of Tax and Revenue
publishes. This file is text extracted from that document, so the claims can be
checked.

Retrieved 2026-10-20 with `pdftotext -layout` from poppler 24.02.0. The PDF holds
the form and its instructions; the extract is the form's front, from the
exceptions to its revision line.

| File | Source | sha256 of the PDF |
|---|---|---|
| `d2210-2025.txt` | [2025 D-2210, Underpayment of Estimated Income Tax By Individuals](https://otr.cfo.dc.gov/sites/default/files/dc/sites/otr/publication/attachments/2025_D-2210_021226.pdf) (revised 05/2025) | `26dbe65554e7244da2c614cc5f34ba0b8af5bfe81b6eb82a620b6cdef25b6072` |

```bash
curl -sS -L -o d2210-2025.pdf https://otr.cfo.dc.gov/sites/default/files/dc/sites/otr/publication/attachments/2025_D-2210_021226.pdf
pdftotext -layout d2210-2025.pdf d2210.L.txt
awk '/No underpayment Interest Is due/{f=1} f; /Revised 05\/2025/{if(f)exit}' d2210.L.txt > d2210-2025.txt
```

The 2026 form is not published. The 2025 form is the latest.
