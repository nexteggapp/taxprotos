# Where these files came from

The line numbers in `mapping_1065.textproto`, `mapping_1120s.textproto`, and `mapping_1041.textproto` are claims about documents the IRS
publishes. These files are text extracted from those documents, so the claims
can be checked.

Retrieved 2026-10-24 with `pdftotext -layout` from poppler 24.02.0, unfiltered —
each file is the whole form.

| File | Source | sha256 of the PDF |
|---|---|---|
| `f1065sk1-2025.txt` | [f1065sk1--2025.pdf](https://www.irs.gov/pub/irs-prior/f1065sk1--2025.pdf) | `66098d4d48537ce2dac1f093d6351567957896e5843d8c524823b380068f6547` |
| `f1120ssk-2025.txt` | [f1120ssk--2025.pdf](https://www.irs.gov/pub/irs-prior/f1120ssk--2025.pdf) | `104ac7b1137cb78c9051d4e3424630dca233887f17928b311fef6fd578baacd0` |
| `f1041sk1-2025.txt` | [f1041sk1--2025.pdf](https://www.irs.gov/pub/irs-prior/f1041sk1--2025.pdf) | `d8d7b6eacabdf145474aee8385fcfeafe68bd2f69baaef52d9a8227ebcd45fc3` |

```bash
curl -o f1065sk1--2025.pdf https://www.irs.gov/pub/irs-prior/f1065sk1--2025.pdf
pdftotext -layout f1065sk1--2025.pdf f1065sk1-2025.txt

curl -o f1120ssk--2025.pdf https://www.irs.gov/pub/irs-prior/f1120ssk--2025.pdf
pdftotext -layout f1120ssk--2025.pdf f1120ssk-2025.txt

curl -o f1041sk1--2025.pdf https://www.irs.gov/pub/irs-prior/f1041sk1--2025.pdf
pdftotext -layout f1041sk1--2025.pdf f1041sk1-2025.txt
```
