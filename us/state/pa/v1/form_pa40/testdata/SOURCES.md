# Where this file came from

The line numbers in `mapping.textproto` are claims about a document the
Pennsylvania Department of Revenue publishes. This file is text extracted from
that document, so the claims can be checked.

Retrieved 2026-11-12 with `pdftotext -layout` from poppler 24.02.0, unfiltered —
the file is the whole form.

| File | Source | sha256 of the PDF |
|---|---|---|
| `pa40-2025.txt` | [2025 Form PA-40](https://www.pa.gov/content/dam/copapwp-pagov/en/revenue/documents/formsandpublications/formsforindividuals/pit/documents/2025/2025_pa-40.pdf) | `fcd7648601f3ce5a7a258be49c54aa1c604e1c61c1fe9175351451ac6087d8b3` |

```bash
curl -o 2025_pa-40.pdf https://www.pa.gov/content/dam/copapwp-pagov/en/revenue/documents/formsandpublications/formsforindividuals/pit/documents/2025/2025_pa-40.pdf
pdftotext -layout 2025_pa-40.pdf pa40-2025.txt
```

The older address under `www.revenue.pa.gov` now redirects to a search page
rather than the PDF.

The instructions (`2025_pa-40in.pdf`, sha256
`e8996504a4302a65055e697ba6301a3e41a411a4802645a5fc5f653f7538c172`, same
directory) were used for the meaning of each line and the Schedule SP
eligibility income tables; they are not checked in.

## Lines 32 to 36 are one figure here

The form prints five refund donation lines, each an organization code and an
amount. The schema carries their total as one field, mapped to line 32 with a
note.
