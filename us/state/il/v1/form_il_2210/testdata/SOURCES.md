# Where this file came from

The line numbers in `mapping.textproto` are claims about a document the
Illinois Department of Revenue publishes. This file is text extracted from
that document, so the claims can be checked.

Retrieved 2026-10-20 with curl 8.5.0 and no special headers, and extracted
with `pdftotext -layout` from poppler 24.02.0, unfiltered: the file is the
whole form, all four pages. Every page has extractable text.

| File | Source | sha256 of the PDF |
|---|---|---|
| `il-2210-2025.txt` | [2025 IL-2210, Computation of Penalties for Individuals (R-12/25)](https://tax.illinois.gov/content/dam/soi/en/web/tax/forms/incometax/documents/currentyear/individual/il-2210.pdf) | `4bb1ba097544532e41657fb1c769654acadf77a7a1a8c57eb67d68510e678fe7` |

```bash
curl -O https://tax.illinois.gov/content/dam/soi/en/web/tax/forms/incometax/documents/currentyear/individual/il-2210.pdf
pdftotext -layout il-2210.pdf il-2210-2025.txt
```

The URL is the department's current-year copy, which a later year's form
will replace; the hash identifies the 2025 one. The instructions, which the
field comments also follow (lines 10a, 10b and Penalty Worksheet 1 in
particular), are the
[2025 Form IL-2210 Instructions](https://tax.illinois.gov/content/dam/soi/en/web/tax/forms/incometax/documents/currentyear/individual/il-2210-instr.pdf).
