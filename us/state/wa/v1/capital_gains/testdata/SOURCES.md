# Where this file came from

Every other mapping in this tree is checked against a printed form. This one is
checked against a statute, because Washington has no printed form to check.

Retrieved 2026-08-26.

| File | Source | sha256 of the source |
|---|---|---|
| `rcw-82-87-extract.txt` | [Chapter 82.87 RCW](https://app.leg.wa.gov/rcw/default.aspx?cite=82.87&full=true) | `4e8229ad4d649e81ce120371433d1bd17403c2166b8ca413fd004fb24ad530b8` |

```bash
curl -A "Mozilla/5.0" -o rcw8287.html "https://app.leg.wa.gov/rcw/default.aspx?cite=82.87&full=true"
sed -e 's/<[^>]*>//g' rcw8287.html \
  | sed 's/&nbsp;/ /g; s/&amp;/\&/g; s/&#8212;/-/g; s/^[[:space:]]*//' \
  | grep -v '^$' > rcw8287.txt
```

The twelve cited sections are then pulled out of that text. Each section body is
introduced by `PDFRCW`, which the chapter's table of contents entries are not,
and runs to the bracketed session-law history that closes it. Subsections are
put one per line so a diff points at a subsection rather than at a wall of text.

## Why the statute and not a form

RCW 82.87.110 requires the return to be filed electronically. There is no paper
form, no line numbers, and nothing to extract: what the Department of Revenue
publishes is a walkthrough of the filing wizard, and its field labels are inside
screenshots rather than in the PDF's text layer.

So `mapping.textproto` leaves `irs_box` empty and cites a subsection in the note
instead, and `coverage_test.go` checks those citations against this file.
`TestEveryConceptCitesTheStatute` also fails a mapping that *does* carry a line
number, because such a mapping was copied from another state without being read.

## This oracle is stronger than a form extract, not weaker

A form extract shows that a line exists. It says nothing about what belongs on
it, which is the limitation every other SOURCES.md in this tree records.

A statute enumerates: what is taxed, what is exempt, what is deductible, and at
what rate. So a citation that has drifted to the wrong subsection fails here,
where on a form it would look plausible indefinitely. The same kind of check
found three wrong citations in the `washingtontaxcalc` package of
[taxcalcs](https://github.com/nexteggapp/taxcalcs), including the standard
deduction cited to the chapter's definitions section.

## What this does not prove

That the schema's *shape* matches what the wizard collects. The statute says
what may be deducted; it does not say the wizard asks for donations one at a
time, which is why the repeated `CharitableDonation` message rests on the
Department's walkthrough rather than on chapter 82.87 RCW.

That a web page will still say this tomorrow. The hash lets somebody re-fetch
and see *that* the chapter changed, not what it used to say. A statute changes
far less often than a page of guidance, which is the one respect in which this
source is more stable than the Department's own capital gains tax page.
