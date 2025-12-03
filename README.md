# compliance-as-code-attestors
compliance-as-code-attestors


```rego

opa eval --data policy.rego --input output3.json "data.documentationMerged.merge"


opa eval --data policy.rego --input output3.json --fail-defined \
  'not data.documentationMerged.merge'

```