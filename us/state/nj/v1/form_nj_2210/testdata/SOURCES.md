# Where this file came from

The line numbers in `mapping.textproto` are claims about a document the New
Jersey Division of Taxation publishes. This file is text extracted from that
document, so the claims can be checked.

Retrieved 2026-10-20 with `pdftotext -layout` from poppler 24.02.0: the form,
its page 2 worksheets and its instructions, with Calculating the Interest, in
page layout.

| File | Source | sha256 of the PDF |
|---|---|---|
| `nj-2210-2025.txt` | [Form NJ-2210, 2025 Underpayment of Estimated Tax by Individuals, Estates, or Trusts, with instructions](https://www.nj.gov/treasury/taxation/pdf/current/2210.pdf) | `376d6913ab27304556a95d242f34c972c386ad16758ce35a0eee68fa172bbb14` |

```bash
curl -o 2210.pdf https://www.nj.gov/treasury/taxation/pdf/current/2210.pdf
pdftotext -layout 2210.pdf nj-2210-2025.txt
```

The "current" URL serves the 2025 form until the Division publishes the
next; the hash identifies the edition.
