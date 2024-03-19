package main

type Value struct {
	data int
}

func printSome(v Value) {
	println(v.data)
}

func main() {
	x := Value{123}

	defer func() {
		printSome(x)
	}()

	x.data = 222
}
