type = -chan

all: emerging

emerging: emerging.go cmap.go lmap.go
	go build emerging.go cmap.go lmap.go

test-chan: emerging
	./emerging -chan -readers=2 -askers=2 -askdelay=10 -infiles=data/pg1041.txt,data/pg1103.txt

test-lock: emerging
	./emerging -lock -readers=2 -askers=2 -askdelay=10 -infiles=data/pg1041.txt,data/pg1103.txt

test-runs: emerging
	echo "1 reader and 1 asker\n" >> times.txt
	{ time ./emerging $(type) -readers=1 -askers=1 -askdelay=10 -infiles=data/pg1041.txt,data/pg1103.txt; } 2>> times.txt
	echo "\n16 readers and 2 askers\n" >> times.txt
	{ time ./emerging $(type) -readers=16 -askers=2 -askdelay=10 -infiles=data/pg1041.txt,data/pg1103.txt; } 2>> times.txt
	echo "\n4 readers 8 askers\n" >> times.txt
	{ time ./emerging $(type) -readers=4 -askers=8 -askdelay=10 -infiles=data/pg1041.txt,data/pg1103.txt; } 2>> times.txt
	echo "\n16 readers 32 askers\n" >> times.txt
	{ time ./emerging $(type) -readers=16 -askers=32 -askdelay=10 -infiles=data/pg1041.txt,data/pg1103.txt; } 2>> times.txt
	echo "\n64 readers 64 askers\n" >> times.txt
	{ time ./emerging $(type) -readers=64 -askers=64 -askdelay=10 -infiles=data/pg1041.txt,data/pg1103.txt; }  2>> times.txt


.PHONY: clean

clean:
	rm emerging
