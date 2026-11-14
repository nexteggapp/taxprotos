# Where this file came from

The line numbers in `mapping.textproto` are claims about a document the
Kansas Department of Revenue publishes. This file is text extracted from that
document, so the claims can be checked.

| File | Source | sha256 of the PDF |
|---|---|---|
| `k40-2025.txt` | [2025 Form K-40](https://www.ksrevenue.gov/pdf/k-4025.pdf) | `dfff7caae33d2527ae9655f562a3469835556cc6ec6ab40297641caf10379fdb` |

Downloaded on 2026-11-12 and extracted with `pdftotext -layout` (poppler).

## The front page

Most of page 1 uses a font whose text does not extract: the filing status,
residency and exemption boxes come out as shifted characters. Those boxes have
no line numbers, so the mapping names them by their caption instead, taken from
the 2025 Form K-40 instructions (https://www.ksrevenue.gov/pdf/ip25.pdf, sha256
`8cbc2ea8d4f877a0d57e1e008c203ec06aa087fb3d4ab7319d91c157ee628871`), and the
coverage test has no line to check for them. Page 2, which holds every
numbered line, extracts cleanly.
