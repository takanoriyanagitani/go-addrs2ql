#!/bin/sh

query='query NetIfaces {
  netIfaces {
    index
    mtu
    name
    addrs {
      network
      string
    }
  }
}'

time jq \
    -c \
    -n \
    --arg q "${query}" \
    '{
      query: $q
    }' |
    curl \
        --header 'Content-Type: application/json' \
        --silent \
        --show-error \
        --fail \
        --location \
        --data @- \
        http://localhost:8158/query |
    jq -c '.data.netIfaces.[]' |
    jq -c 'select("lo0" != .name)' |
    jq -s -c 'sort_by(.mtu).[]'
