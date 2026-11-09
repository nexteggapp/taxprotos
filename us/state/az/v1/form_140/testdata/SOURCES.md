# Where this file came from

The line numbers in `mapping.textproto` are claims about a document the
Arizona Department of Revenue publishes. This file is text extracted from that
document, so the claims can be checked.

Retrieved 2026-11-09 with `pdftotext -layout` from poppler 24.02.0, unfiltered —
the file is the whole form, including the schedules on pages 3 to 6.

| File | Source | sha256 of the PDF |
|---|---|---|
| `form140-2025.txt` | [2025 Form 140](https://azdor.gov/sites/default/files/document/FORMS_INDIVIDUAL_2025_140_f.pdf) | `2acdda2a07ab46d7f6bd701618e63d1af8d45ac7921c38572a740940437edf84` |

```bash
curl -o 140.pdf https://azdor.gov/sites/default/files/document/FORMS_INDIVIDUAL_2025_140_f.pdf
pdftotext -layout 140.pdf form140-2025.txt
```
