# compliance-as-code-attestors
compliance-as-code-attestors


```rego

opa eval --data policy.rego --input output3.json 'data.documentationMerged.failure_msg' --format raw --fail-defined

```

## How to run locally

```bash

go run main.go prAttest --repos l3montree-dev/devguard,l3montree-dev/devguard-web,l3montree-dev/devguard-documentation \
            --pull_request_title "1277 organization wide dependency search" \
            --pull_request_number "8"

```

```bash

docker run compliance-as-code-attestors-1 prAttest --repos l3montree-dev/devguard,l3montree-dev/devguard-web,l3montree-dev/devguard-documentation \
            --pull_request_title "1277 organization wide dependency search" \
            --pull_request_number "8"

```

