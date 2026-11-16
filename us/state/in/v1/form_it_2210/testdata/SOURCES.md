# Where this file came from

The line numbers in `mapping.textproto` and `mapping_section_e.textproto` are
claims about a document the Indiana Department of Revenue publishes. This file
is text extracted from that document, so the claims can be checked.

Retrieved 2026-11-16 with curl 8.5.0 and no special headers, and extracted
with `pdftotext -layout` from poppler 24.02.0, unfiltered: the file is the
whole schedule and its instructions, all four pages.

| File | Source | sha256 of the PDF |
|---|---|---|
| `it2210-2025.txt` | [Schedule IT-2210, 2025 Underpayment of Estimated Tax by Individuals](https://forms.in.gov/Download.aspx?id=16926), State Form 46002 (R27 / 9-25) | `9611265ce223f620162c8e1c79e2d68265b92c99b595494107cf9402a7da61b4` |

```bash
curl -sS -L -o it2210.pdf 'https://forms.in.gov/Download.aspx?id=16926'
pdftotext -layout it2210.pdf it2210-2025.txt
```

The 2026 schedule is not published. The 2025 schedule is the latest.
