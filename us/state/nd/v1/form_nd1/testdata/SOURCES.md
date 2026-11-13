# Where this file came from

The line numbers in `mapping.textproto` are claims about a document the
North Dakota Office of State Tax Commissioner publishes. This file is text
extracted from that document, so the claims can be checked.

Retrieved 2026-11-12 with `pdftotext -layout` from poppler 24.02.0, unfiltered —
the file is the whole form.

| File | Source | sha256 of the PDF |
|---|---|---|
| `nd1-2025.txt` | [2025 Form ND-1](https://www.tax.nd.gov/sites/www/files/documents/forms/individual/2025-iit/28702-form-nd-1-2025.pdf) | `e7c35d515b7b1e101bef3ceb04cd8363d5df3e44da19654982b8f26bfeed83da` |

```bash
curl -o 28702-form-nd-1-2025.pdf https://www.tax.nd.gov/sites/www/files/documents/forms/individual/2025-iit/28702-form-nd-1-2025.pdf
pdftotext -layout 28702-form-nd-1-2025.pdf nd1-2025.txt
```

## The text is printed twice

The PDF overlays each label on a copy of itself, so the extract repeats most
words. The line numbers are still set off by whitespace, which is all the
coverage test needs.
