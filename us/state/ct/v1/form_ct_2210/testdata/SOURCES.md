# Where this file came from

The line numbers in `mapping.textproto` are claims about a document the
Connecticut Department of Revenue Services publishes. This file is text
extracted from that document, so the claims can be checked.

Retrieved 2026-10-20 with `pdftotext` from poppler 24.02.0, in reading order:
the form, Schedules A and B, and the instructions, without Worksheet 1 (pages
5 and 6, the nonresidents' annualized tax). The 2026 Form CT-2210 is not yet
published.

| File | Source | sha256 of the PDF |
|---|---|---|
| `formct2210-2025.txt` | [Form CT-2210 (Rev. 12/25)](https://portal.ct.gov/-/media/drs/forms/2025/income/ct-2210_1225.pdf) | `9559ebc66de7385d50f7597909e505b13786944623c5b4d950d67b5ca9fe56b8` |

```bash
curl -o ct2210.pdf https://portal.ct.gov/-/media/drs/forms/2025/income/ct-2210_1225.pdf
pdftotext ct2210.pdf 2210.txt
awk '/^\f?Worksheet 1$/{skip=1} !skip; /^Page 6 of 10$/{skip=0}' 2210.txt > formct2210-2025.txt
```
