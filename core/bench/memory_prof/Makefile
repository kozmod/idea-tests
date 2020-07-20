OUT_DIR :=out

STACK_OUT :=$(OUT_DIR)/stack.txt
STACK_ARG_OUT :=$(OUT_DIR)/stack_arg.txt
STACK_OUT_BSTAT :=$(OUT_DIR)/stack-bstat.txt
STACK_ARG_OUT_BSTAT :=$(OUT_DIR)/stack-arg-bstat.txt

HEAP_OUT :=$(OUT_DIR)/heap.txt
HEAP_ARG_OUT :=$(OUT_DIR)/heap_arg.txt
HEAP_OUT_BSTAT :=$(OUT_DIR)/heap-bstat.txt
HEAP_ARG_OUT_BSTAT :=$(OUT_DIR)/heap-arg-bstat.txt

ITERATIONS :=10

all:
	make clear
	make stack
	make heap

all.arg:
	make clear
	make stack.arg
	make heap.arg

stack:
	@(make mout)
	@(./bench.sh runStack ${STACK_OUT} ${STACK_OUT_BSTAT} ${ITERATIONS})

stack.arg:
	@(make mout)
	@(./bench.sh runStackArg ${STACK_ARG_OUT} ${STACK_ARG_OUT_BSTAT} ${ITERATIONS})

heap:
	@(make mout)
	@(./bench.sh runHeap ${HEAP_OUT} ${HEAP_OUT_BSTAT} ${ITERATIONS})

heap.arg:
	@(make mout)
	@(./bench.sh runHeapArg ${HEAP_ARG_OUT} ${HEAP_ARG_OUT_BSTAT} ${ITERATIONS})

mout:
	@([ ! -d ${OUT_DIR} ] && mkdir ${OUT_DIR} || echo "${OUT_DIR} exists")

trace.stack:
	go tool trace $(OUT_DIR)/stack.out

clear:
	@([ -d ${OUT_DIR} ] & rm -rf ${OUT_DIR})