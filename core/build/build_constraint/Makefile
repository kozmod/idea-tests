out := bin

clean:
	rm -rf ${out}

build:
	go build --tags a -o ${out}/main_a .
	go build --tags b -o ${out}/main_b .

run:
	${out}/main_a
	@(echo "")
	${out}/main_b

