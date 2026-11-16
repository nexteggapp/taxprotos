# Where these files came from

Washington prints no form for paying its capital gains tax, so the mappings
are checked against the statute: chapter 82.87 RCW for the payment, RCW
82.32.090 for the late payment penalty, RCW 82.32.050 for interest, and RCW
1.12.070 for a deadline that falls on a weekend or holiday.

The chapter extract is built as `../../capital_gains/testdata/SOURCES.md`
describes, and also carries RCW 82.87.115 as 2026 c 191 added it. The other
three are the Legislature's PDF prints, extracted with `pdftotext` from poppler
24.02.0 in reading order. All four were retrieved 2026-10-20 with curl 8.5.0.

| File | Source | sha256 of the source |
|---|---|---|
| `rcw-82-87-extract.txt` | [Chapter 82.87 RCW](https://app.leg.wa.gov/rcw/default.aspx?cite=82.87&full=true), sections 82.87.020 to 82.87.160 | `5d1ba6bc30e2a9b1b8afd98f24918f700edc611355cb28f202b34bf12b4cb3af` |
| `rcw-82-32-090.txt` | [RCW 82.32.090](https://app.leg.wa.gov/RCW/default.aspx?cite=82.32.090&pdf=true), certified 7/15/2026 | `4eb7d3bc868d6eff99d322fd8e383fd87b8c2b74399e50b4568137819f202c00` |
| `rcw-82-32-050.txt` | [RCW 82.32.050](https://app.leg.wa.gov/RCW/default.aspx?cite=82.32.050&pdf=true), certified 7/15/2026 | `bf4536eb3f1aad869c1e65531e3caee489826fdef5bb57e284afb4f479e831ec` |
| `rcw-1-12-070.txt` | [RCW 1.12.070](https://app.leg.wa.gov/RCW/default.aspx?cite=1.12.070&pdf=true) | `d3029b32acf21043f35b394eb4ba8049f9992b3a93a00983cb24f0fb620927b7` |

```bash
for c in 82.32.090 82.32.050 1.12.070; do
  curl -L -o "rcw-${c//./-}.pdf" "https://app.leg.wa.gov/RCW/default.aspx?cite=$c&pdf=true"
  pdftotext "rcw-${c//./-}.pdf" "rcw-${c//./-}.txt"
done
```

The department's guidance (the capital gains tax page, its April 7, 2023
interim statement on late payment penalties, and WAC 458-20-228 and 458-20-300)
predates parts of the 2025 act. Where it differs from the statute the rules
that fill this message say so; the message follows the statute.
