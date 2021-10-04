:: TODO add conditions to make Makefile usable in linux
all: bin/main.exe

bin/main.exe: main.go crypto/hotp.go
	go build -o "./bin/main.exe" main.go   

:: Add the parameters here to test with make run
:: target: Target api
:: interval: Interval between each new hotp
run: bin/main.exe
	bin\main.exe -interval=10

clean:
	del bin\main.exe
