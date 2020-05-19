
ITERATIONS :=10

all:
	go test -bench=BeginsWithCopy -count=${ITERATIONS} -cpuprofile=cpuprof_copy.out
	go test -bench=BeginsWithArgCopy -count=${ITERATIONS} -cpuprofile=cpuprof_copy_arg.out
	go test -bench=BeginsWithPointer -count=${ITERATIONS} -cpuprofile=cpuprof_pointer.out
	go test -bench=BeginsWithArgPointer -count=${ITERATIONS} -cpuprofile=cpuprof_pointer.out

#  Open https://www.speedscope.app/ and open the .out files which will show you the flame graph.