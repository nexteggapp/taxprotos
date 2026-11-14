# Where this file came from

The line numbers in `mapping.textproto` are claims about a document the
Nebraska Department of Revenue publishes. This file is text extracted from that
document, so the claims can be checked.

Retrieved 2026-11-12 with `pdftotext -layout` from poppler 24.02.0, unfiltered —
the file is the whole form.

| File | Source | sha256 of the PDF |
|---|---|---|
| `f1040n-2025.txt` | [2025 Form 1040N](https://revenue.nebraska.gov/sites/default/files/doc/tax-forms/2025/f_1040N.pdf) | `db5850008b04901cce2a467707cc8db19e0f9194548d610247403ae8a9e078a3` |

```bash
curl -o f_1040N.pdf https://revenue.nebraska.gov/sites/default/files/doc/tax-forms/2025/f_1040N.pdf
pdftotext -layout f_1040N.pdf f1040n-2025.txt
```

## Lines printed in parts

Lines 4 and 16 are printed in parts, "4 a" to "4 c" and "16 a" to "16 c".
The extract has "4a", "4c", "16a" and "16c" in the instructions printed on the
form; 4a, 4b and 16b are followed by punctuation, so the mapping records them
as lines 4 and 16 with a note saying which part they are.
