# Where this file came from

The line numbers in `mapping.textproto` are claims about a document the
Connecticut Department of Revenue Services publishes. This file is text
extracted from that document, so the claims can be checked.

Retrieved 2026-11-14 with `pdftotext -layout` from poppler 24.02.0, unfiltered —
the file is the whole form, including Schedules 1 to 5.

| File | Source | sha256 of the PDF |
|---|---|---|
| `ct1040-2025.txt` | [2025 Form CT-1040 (Rev. 12/25)](https://portal.ct.gov/-/media/drs/forms/2025/income/ct-1040_1225.pdf) | `fadbe2acfd629a61c59266fd5d5f80cb0f717da556203b8cd73f6ae49d538b9b` |

```bash
curl -o ct-1040_1225.pdf https://portal.ct.gov/-/media/drs/forms/2025/income/ct-1040_1225.pdf
pdftotext -layout ct-1040_1225.pdf ct1040-2025.txt
```

## Line 17 repeats line 16

Line 17 carries line 16 to the payments section. The mapping records the
figure once, on line 16.
