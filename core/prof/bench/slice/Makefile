PROF_OUT=prof
CPU_OUT=${PROF_OUT}/cpu.out
MEM_OUT=${PROF_OUT}/mem.out
BIN_OUT=${PROF_OUT}/slice.test

test:
	go test -bench . -benchmem ./...

out:
	go test -bench . -benchmem -cpuprofile=${CPU_OUT} -memprofile=${MEM_OUT} -memprofilerate=1 -o ${BIN_OUT} ./...

refresh:
	rm -rf ${PROF_OUT} && mkdir ${PROF_OUT}

echo:
	@(echo "go tool pprof ${BIN_OUT} ${CPU_OUT} -- top, list <function>, web/png/gif")
	@(echo "go tool pprof ${BIN_OUT} ${MEM_OUT} -- top, list <function>, alloc_space, web")
	@(echo "top, list <function>, web/png/gif")

