# Where this file came from

The line numbers in `mapping.textproto` are claims about a document the New
York State Department of Taxation and Finance publishes. This file is text
extracted from that document, so the claims can be checked.

Retrieved 2026-11-14 with `pdftotext -layout` from poppler 24.02.0, unfiltered —
the file is the whole form.

| File | Source | sha256 of the PDF |
|---|---|---|
| `it201-2025.txt` | [2025 Form IT-201](https://www.tax.ny.gov/pdf/current_forms/it/it201_fill_in.pdf) | `cfe63039ffa4374f5fff67b7d2b37c5bdb802211fac55c69386e8f15c20d3e08` |

```bash
curl -o it201_fill_in.pdf https://www.tax.ny.gov/pdf/current_forms/it/it201_fill_in.pdf
pdftotext -layout it201_fill_in.pdf it201-2025.txt
```

The URL is the department's "current forms" copy, which will become the 2026
form. The sha256 identifies the 2025 one.

## Items rather than lines

Filing status, the dependent box, New York City residence and the dependents
are items A, C, F and H on page 1, not numbered lines. The mapping records
them by their item letter.
