# Where this file came from

The line numbers in `mapping.textproto` are claims about a document Maine
Revenue Services publishes. This file is text extracted from that document, so
the claims can be checked.

Retrieved 2026-11-14 with `pdftotext -layout` from poppler 24.02.0, unfiltered —
the file is the whole form.

| File | Source | sha256 of the PDF |
|---|---|---|
| `1040me-2025.txt` | [2025 Form 1040ME](https://www.maine.gov/revenue/sites/maine.gov.revenue/files/inline-files/25_1040ME_fillable.pdf) | `6692450e683dc565a1f6605ce76eebb9244cbee7f02151943f603391593be197` |

```bash
curl -o 25_1040ME_fillable.pdf https://www.maine.gov/revenue/sites/maine.gov.revenue/files/inline-files/25_1040ME_fillable.pdf
pdftotext -layout 25_1040ME_fillable.pdf 1040me-2025.txt
```

The site refuses requests that do not look like a browser's; a browser user
agent and a referrer from the forms page were needed.

## Lines printed after leader dots

Lines 25b to 25f and 34a and 34b are printed only at the end of a run of
leader dots, with a bare "b." to "f." beside line 25, so the mapping records
them as lines 25 and 34 with a note.
