# compliance-as-code-attestors
compliance-as-code-attestors


```rego

opa eval --data policy.rego --input output3.json 'data.documentationMerged.failure_msg' --format raw --fail-defined

```