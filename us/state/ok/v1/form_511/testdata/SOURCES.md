# Where this file came from

The line numbers in `mapping.textproto` are claims about a document the
Oklahoma Tax Commission publishes. This file is text extracted from that
document, so the claims can be checked.

The Commission publishes Form 511 inside its resident income tax packet,
together with the instructions and the tax table. Pages 42 to 49 of the packet
are the form itself, pages 1 to 8 of Form 511 including Schedules 511-A to
511-I. Retrieved 2026-11-12 with `pdftotext -layout` from poppler 24.02.0,
limited to those pages.

| File | Source | sha256 of the PDF |
|---|---|---|
| `form511-2025.txt` | [2025 Form 511 packet](https://oklahoma.gov/content/dam/ok/en/tax/documents/forms/individuals/current/511-Pkt.pdf) | `c90cf30381cca6262ec94050acde8db1ee58885360f0af4ec21e045d1cc5a572` |

```bash
curl -o 511-Pkt.pdf https://oklahoma.gov/content/dam/ok/en/tax/documents/forms/individuals/current/511-Pkt.pdf
pdftotext -layout -f 42 -l 49 511-Pkt.pdf form511-2025.txt
```
