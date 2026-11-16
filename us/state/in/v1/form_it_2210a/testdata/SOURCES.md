# Where this file came from

The line numbers in `mapping.textproto` and `mapping_period.textproto` are
claims about a document the Indiana Department of Revenue publishes. This
file is text extracted from that document, so the claims can be checked.

Retrieved 2026-11-16 with curl 8.5.0 and no special headers, and extracted
with `pdftotext -layout` from poppler 24.02.0, unfiltered: the file is the
whole schedule and its instructions, all five pages.

| File | Source | sha256 of the PDF |
|---|---|---|
| `it2210a-2025.txt` | [Schedule IT-2210A, 2025 Annualized Income Schedule for the Underpayment of Estimated Tax by Individuals](https://forms.in.gov/Download.aspx?id=16928), State Form 48437 (R25 / 9-25) | `a3d9f70671a50ae1974ff41f9a89ea65c42e325b7c518735b2193ea94347dbe7` |

```bash
curl -sS -L -o it2210a.pdf 'https://forms.in.gov/Download.aspx?id=16928'
pdftotext -layout it2210a.pdf it2210a-2025.txt
```

The 2026 schedule is not published. The 2025 schedule is the latest.
