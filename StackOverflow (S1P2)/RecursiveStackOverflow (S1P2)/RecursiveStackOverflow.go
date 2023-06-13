package main

func Fib(n int64) int64 {
	if n == 1 {
		return 1
	}
	return Fib(n-1) + Fib(n-2)

}

func main() {
	// Вызываем функцию overflow с начальным значением n равным 1
	Fib(-1)
}
