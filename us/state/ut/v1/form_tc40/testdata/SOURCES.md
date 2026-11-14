# Where this file came from

The line numbers in `mapping.textproto` are claims about a document the Utah
State Tax Commission publishes. These files are text taken from that document,
so the claims can be checked.

| File | Source | sha256 of the PDF |
|---|---|---|
| `tc40-2025.txt` | [2025 Form TC-40](https://files.tax.utah.gov/tax/forms/current/tc-40.pdf), page 1 | `0eac22fb3d1487c465d29a4351d5fe3a29c5a511febd5975e0fc2b6688637cdf` |
| `tc40-2025-page2.txt` | The same PDF, page 2 | the same |

Retrieved 2026-11-12.

```bash
curl -o tc-40.pdf https://files.tax.utah.gov/tax/forms/current/tc-40.pdf
pdftotext -layout tc-40.pdf tc40-2025.txt          # poppler 24.02.0
pdftoppm -f 2 -l 2 -r 110 -png tc-40.pdf page      # page-2.png, read by eye
```

The URL is the Commission's "current" copy, which will become the 2026 form.
The sha256 identifies the 2025 one.

## Page 2 is transcribed, not extracted

Only page 1 of the PDF carries text. Pages 2 and 3 are drawn as outlines, so
`pdftotext` and `mutool` return nothing for them. `tc40-2025-page2.txt` is a
transcription of page 2 (lines 23 to 44) from the rendered page, line for line,
with the wording as printed. Page 3 holds only names, codes and account details,
and no line of it is mapped.

The line numbers and wording of page 2 agree with the 2025 TC-40 instructions
(`https://files.tax.utah.gov/tax/forms/current/tc-40inst.pdf`, sha256
`9ea2823d81494434fb4a3d77be1ad103042ba4bdd1462dcfdca33f8a2fdcddda`), whose
text does extract, under "TC-40, Page 2".

## Box 2 is printed in parts

The dependent counts are printed as "2", with "a" to "d" set beside it, so
2a to 2d never appear as such. The mapping records them as line 2 with a note.
