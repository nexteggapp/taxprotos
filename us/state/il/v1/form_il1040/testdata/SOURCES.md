# Where this file came from

The line numbers in `mapping.textproto` are claims about a document the
Illinois Department of Revenue publishes. This file is text extracted from that
document, so the claims can be checked.

Retrieved 2026-11-09 with `pdftotext -layout` from poppler 24.02.0, unfiltered —
the file is the whole form.

| File | Source | sha256 of the PDF |
|---|---|---|
| `il1040-2025.txt` | [2025 Form IL-1040](https://tax.illinois.gov/content/dam/soi/en/web/tax/forms/incometax/documents/currentyear/individual/il-1040.pdf) | `7af662f99b31f6aa0d00df5ae93c2a50c4c15e01a0c1dcf5f722fc0bde0799cd` |

```bash
curl -o il-1040.pdf https://tax.illinois.gov/content/dam/soi/en/web/tax/forms/incometax/documents/currentyear/individual/il-1040.pdf
pdftotext -layout il-1040.pdf il1040-2025.txt
```

The URL is the department's "current year" copy, which will become the 2026
form. The sha256 identifies the 2025 one.

## Line 10 is printed in parts

The exemption lines are printed as "10", with "a" to "d" set beside it, so only
"10a" and "10d" appear as such, in the instruction to add them. The mapping
records 10b and 10c as line 10 with a note.
