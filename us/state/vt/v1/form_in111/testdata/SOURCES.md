# Where this file came from

The line numbers in `mapping.textproto` are claims about a document the
Vermont Department of Taxes publishes. This file is text extracted from that
document, so the claims can be checked.

Retrieved 2026-11-14 with `pdftotext -layout` from poppler 24.02.0, unfiltered —
the file is the whole form.

| File | Source | sha256 of the PDF |
|---|---|---|
| `in111-2025.txt` | [2025 Form IN-111](https://tax.vermont.gov/sites/tax/files/documents/IN-111-2025.pdf) | `728c6d2a4f14d9a6c6fd5939e72a954c8c727db42daef4fd74568de6306827fc` |

```bash
curl -o IN-111-2025.pdf https://tax.vermont.gov/sites/tax/files/documents/IN-111-2025.pdf
pdftotext -layout IN-111-2025.pdf in111-2025.txt
```

## Numbers run into the entry rule

The form prints some line numbers straight into the rule for the amount, as
`18.____` and `24e.____`, so the coverage test accepts a full stop followed by
an underscore as the end of a line number.
