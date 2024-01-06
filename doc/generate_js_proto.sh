#!/usr/bin/env bash

set -x
set -e

rm -rf ./protocol/js/*

# protoc --es_out ./../protocol/js --es_opt target=ts -I./protocol ./protocol/tc_egress.proto ./protocol/tc_ingress.proto ./protocol/tc_room.proto ./protocol/tc_webhook.proto ./protocol/tc_models.proto

protoc --es_out ./../protocol/js --es_opt target=ts -I./protocol ./protocol/tc_egress.proto ./protocol/tc_ingress.proto ./protocol/tc_room.proto ./protocol/tc_webhook.proto ./protocol/tc_models.proto ./protocol/tc_rtc.proto ./protocol/tc_sip.proto

# protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative --proto_path=/ /*.proto
# protoc --es_out src/proto --es_opt target=ts -I./protocol ./protocol/livekit_egress.proto ./protocol/livekit_ingress.proto ./protocol/livekit_room.proto ./protocol/livekit_webhook.proto ./protocol/livekit_models.proto