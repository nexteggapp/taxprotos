# Where these files came from

The line numbers in `mapping.textproto` are claims about documents the IRS
publishes. These files are text extracted from those documents, so the claims
can be checked.

Retrieved 2026-10-24 with `pdftotext -layout` from poppler 24.02.0, unfiltered —
each file is the whole form.

| File | Source | sha256 of the PDF |
|---|---|---|
| `f1098t-2025.txt` | [f1098t.pdf](https://www.irs.gov/pub/irs-pdf/f1098t.pdf) | `f461d17ce14de4efb4d95936c8639434861dbf174da250ac777f50cd1883c692` |

```bash
curl -o f1098t.pdf https://www.irs.gov/pub/irs-pdf/f1098t.pdf
pdftotext -layout f1098t.pdf f1098t-2025.txt
```
