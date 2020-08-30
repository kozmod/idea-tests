
ITERATIONS :=10

all:
	go test -bench=BenchmarkWithoutPipelineModule -count=${ITERATIONS}
	go test -bench=BenchmarkWithPipelineModule -count=${ITERATIONS}