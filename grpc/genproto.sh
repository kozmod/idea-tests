#!/usr/bin/env bash
protoc -I ./proto --go_out=plugins=grpc:./vendor com/data/dataA.proto
protoc -I ./proto --go_out=plugins=grpc:./vendor com/data/dataB.proto
protoc -I ./proto --go_out=plugins=grpc:./vendor com/msg/msg.proto