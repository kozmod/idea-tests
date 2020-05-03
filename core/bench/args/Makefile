OUT_DIR :=out

STACK_OUT :=$(OUT_DIR)/stack.txt
STACK_OUT_BSTAT :=$(OUT_DIR)/stack-bstat.txt

HEAP_OUT :=$(OUT_DIR)/heap.txt
HEAP_OUT_BSTAT :=$(OUT_DIR)/heap-bstat.txt

ITERATIONS :=10

all:
	make clear
	make mout
	@(./bench.sh runStack $(HEAP_OUT) $(HEAP_OUT_BSTAT) $(ITERATIONS))
	@(./bench.sh runStack $(STACK_OUT) $(STACK_OUT_BSTAT) $(ITERATIONS))

stack:
	make clear
	make mout
	@(./bench.sh runStack $(STACK_OUT) $(STACK_OUT_BSTAT) $(ITERATIONS))

heap:
	make clear
	make mout
	@(./bench.sh runHeap $(HEAP_OUT) $(HEAP_OUT_BSTAT) $(ITERATIONS))

mout:
	@([ ! -d "./$(OUT_DIR)" ] && mkdir $(OUT_DIR))


clear:
	@([ -d "./$(OUT_DIR)" ] && rm -r ./$(OUT_DIR))

#  Open https://www.speedscope.app/ and open the .out files which will show you the flame graph.