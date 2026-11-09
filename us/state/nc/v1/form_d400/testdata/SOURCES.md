# Where this file came from

The line numbers in `mapping.textproto` are claims about a document the North
Carolina Department of Revenue publishes. This file is text extracted from that
document, so the claims can be checked.

Retrieved 2026-11-09 with `pdftotext -layout` from poppler 24.02.0, unfiltered —
the file is the whole form. The department publishes a web-fill and a
handwritten version with the same lines; this is the web-fill one.

| File | Source | sha256 of the PDF |
|---|---|---|
| `d400-2025.txt` | [2025 Form D-400](https://www.ncdor.gov/2025-d-400-web-filled-version/open) | `2bf6e0b44e628a05cc401d22f7cf6f8c7b39db529d345f4bcf7d93ccfe170d63` |

```bash
curl -L -o d-400.pdf https://www.ncdor.gov/2025-d-400-web-filled-version/open
pdftotext -layout d-400.pdf d400-2025.txt
```

## Lines printed in parts

Lines 20, 21 and 26 are printed as the number with "a." to "e." beside each
part. Where a part never appears joined to its number, the mapping records the
line with a note saying which part it is.
