#!/usr/bin/env bash

function runStack() {
  out=$1
  outbstat=$2
  iterations=$3
  go test ./... -bench=BenchmarkMemoryStack -benchmem -run=^$ -count=${iterations} 2>&1 | tee ${out} && benchstat ${out} 2>&1 | tee ${outbstat}
}

function runStackArg() {
  out=$1
  outbstat=$2
  iterations=$3
  go test ./... -bench=BenchmarkMemoryStackArg -benchmem -run=^$ -count=${iterations} 2>&1 | tee ${out} && benchstat ${out} 2>&1 | tee ${outbstat}
}

function runHeap() {
  out=$1
  outbstat=$2
  iterations=$3
  go test ./... -bench=BenchmarkMemoryHeap -benchmem -run=^$ -count=${iterations} 2>&1 | tee ${out}  && benchstat ${out}  2>&1 | tee ${outbstat}

}

function runHeapArg() {
  out=$1
  outbstat=$2
  iterations=$3
  go test ./... -bench=BenchmarkMemoryHeapArg -benchmem -run=^$ -count=${iterations} 2>&1 | tee ${out}  && benchstat ${out}  2>&1 | tee ${outbstat}

}

# Allows to call a memory_prof based on arguments passed to the script
$*