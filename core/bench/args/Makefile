OUT_DIR :=out

STACK_OUT :=$(OUT_DIR)/stack.txt
STACK_OUT_BSTAT :=$(OUT_DIR)/stack-bstat.txt

HEAP_OUT :=$(OUT_DIR)/heap.txt
HEAP_OUT_BSTAT :=$(OUT_DIR)/heap-bstat.txt

ITERATIONS :=10

all:
	make clear
	make stack
	make heap

stack:
	@(make mout)
	@(./bench.sh runStack ${STACK_OUT} ${STACK_OUT_BSTAT} ${ITERATIONS})

heap:
	@(make mout)
	@(./bench.sh runHeap ${HEAP_OUT} ${HEAP_OUT_BSTAT} ${ITERATIONS})

mout:
	@([ ! -d ${OUT_DIR} ] && mkdir ${OUT_DIR} || echo "${OUT_DIR} exists")


clear:
	@([ -d ${OUT_DIR} ] & rm -rf ${OUT_DIR})

#  Open https://www.speedscope.app/ and open the .out files which will show you the flame graph.