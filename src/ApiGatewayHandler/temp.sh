#!/usr/bin/env bash
curl -v -X POST \
    "https://ho60mx1at5.execute-api.us-east-1.amazonaws.com/test" \
    -H 'content-type: application/json' \
    -H 'day:Thursday' \
    -H 'x-amz-docs-region: us-east-1' \
    -d '{"Name": "haozi"}'