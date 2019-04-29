
.DEFAULT: run

run: helloworld echo1 echo2 echo3 ex

.PHONY: helloworld
helloworld:
	go run helloworld.go

.PHONY: echo1
echo1:
	go run echo1.go 1 a

.PHONY: echo2
echo2:
	go run echo2.go 2 b

.PHONY: echo3
echo3:
	go run echo3.go 3 c

ex: ex1_1 ex1_2

.PHONY: ex1_1
ex1_1:
	cd ex; go run ex1.1.go 1

.PHONY: ex1_2
ex1_2:
	cd ex; go run ex1.2.go 2 b
