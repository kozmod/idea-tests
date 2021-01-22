LOG=./log

.EXPORT_ALL_VARIABLES:
LOG_FILE=${LOG}/$(shell date '+%d_%m_%Y__%H_%M_%S').log
SERVER_PORT=8090

run:
	mkdir -p ${LOG}
	go run ./main.go

clean-log:
	rm -rf ${LOG}