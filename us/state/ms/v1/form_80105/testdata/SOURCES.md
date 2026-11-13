# Where this file came from

The line numbers in `mapping.textproto` are claims about a document the
Mississippi Department of Revenue publishes. This file is text extracted from
that document, so the claims can be checked.

Retrieved 2026-11-12 with `pdftotext -layout` from poppler 24.02.0, unfiltered —
the file is the whole form.

| File | Source | sha256 of the PDF |
|---|---|---|
| `80105-2025.txt` | [2025 Form 80-105](https://www.dor.ms.gov/sites/default/files/tax-forms/individual/80105258%201.pdf) | `57fa91016bf72d1ad734b83528985d8ea0a85501e4a0dbf1f44db83c6827d4fa` |

```bash
curl -o 80105258.pdf 'https://www.dor.ms.gov/sites/default/files/tax-forms/individual/80105258%201.pdf'
pdftotext -layout 80105258.pdf 80105-2025.txt
```

The department lists the form at
<https://www.dor.ms.gov/individual/individual-income-tax-forms>. The form is
Form 80-105-25-8 (Rev. 09/25).

## Filing status

Lines 1 to 5 are the filing status boxes. The mapping records the status at
line 5 (single), with a note saying it stands for all five boxes.
