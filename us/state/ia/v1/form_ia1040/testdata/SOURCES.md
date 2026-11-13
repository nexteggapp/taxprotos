# Where this file came from

The line numbers in `mapping.textproto` are claims about a document the
Iowa Department of Revenue publishes. This file is text extracted from that
document, so the claims can be checked.

Retrieved 2026-11-12 with `pdftotext -layout` from poppler 24.02.0, unfiltered —
the file is the whole form, six pages including IA 1040 Schedule 1.

| File | Source | sha256 of the PDF |
|---|---|---|
| `ia1040-2025.txt` | [2025 IA 1040 (41-001)](https://revenue.iowa.gov/media/4401/download?inline) | `623a3d5afbcec87d74d3596c426ae3bab109304af23efe02dc165544a3cd2f6b` |

```bash
curl -L -o ia1040.pdf 'https://revenue.iowa.gov/media/4401/download?inline'
pdftotext -layout ia1040.pdf ia1040-2025.txt
```

## Short numbering in Steps 2 and 3

The filing status boxes are numbered 1 to 5 and the exemption lines a to d, so
the numbers 1, 3 and 8 are each claimed by two concepts in the mapping, each
with a note saying which.
