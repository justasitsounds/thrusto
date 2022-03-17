.PHONY : run

run:
	rm -rf ./thrusto
	go build
	./thrusto