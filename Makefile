all: emerging

emerging: emerging.go cmap.go lmap.go
	go build emerging.go cmap.go lmap.go

test-chan: emerging
	./emerging -chan -readers=2 -askers=2 -askdelay=10 -infiles=data/pg1041.txt,data/pg1103.txt

test-lock: emerging
	./emerging -lock -readers=2 -askers=2 -askdelay=10 -infiles=data/pg1041.txt,data/pg1103.txt

.PHONY: clean

clean:
	rm emerging
