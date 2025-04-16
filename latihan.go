package main
import "fmt"

func main(){
	const NMAX int = 100
	var A[NMAX]int
	var n, i, x, jum int
	
	fmt.Scan(&n, &x)
	for i = 0; i < n; i++{
		fmt.Scan(&A[i])
	}
	for i = 0; i<n; i++{
		if A[i] == x{
			jum++
		}
	}
	fmt.Print(jum)
}