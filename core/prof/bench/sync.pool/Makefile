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


