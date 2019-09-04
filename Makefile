test:
	go test . -v

generate-map:
	go run ./mapgen -n 20 -m 30 -path ./testdata/map.txt

run:
	go run . -n 500 -path ./testdata/map.txt