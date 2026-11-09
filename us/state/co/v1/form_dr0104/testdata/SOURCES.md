# Where this file came from

The line numbers in `mapping.textproto` are claims about a document the
Colorado Department of Revenue publishes. This file is text extracted from that
document, so the claims can be checked.

| File | Source | sha256 of the PDF |
|---|---|---|
| `dr0104-2025.txt` | [2025 Form DR 0104](https://tax.colorado.gov/sites/tax/files/documents/DR0104_2025.pdf), revision 10/03/25 | `6deb6c181848797505c80ffc351ee973778cbb834ba18ea7007f78fa26ecb49c` |

## Not pdftotext

The department's site answers scripted downloads with a browser challenge, so
the PDF was opened in a browser on 2026-11-09 and its text extracted there with
pdf.js 4.4.168 rather than with `pdftotext -layout`. Each item of each page's
text content was grouped into rows by its rounded vertical position, ordered
left to right, and the rows joined top to bottom; pages are separated by a blank
line. The layout differs from the other states' extracts, but the line numbers
are the same text, and the coverage test asks only whether they appear.

The form prints its line numbers with a trailing full stop, which the coverage
test allows.
