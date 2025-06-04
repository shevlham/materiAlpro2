package main
import "fmt"

const NMAX int = 100

type tabInt [NMAX]int

func main(){
	var a tabInt
	var i, ndata int
	
	fmt.Scan(&ndata)
	for i = 0; i < ndata; i++ {
		fmt.Scan(&a[i])
	}
	insertsort(&a, ndata)
	for i = 0; i < ndata; i++ {
		fmt.Print(a[i]," ")
	}
}

func insertsort(a *tabInt, n int){
	var i, pass, temp int
	
	pass = 1
	for pass < n {
		i = pass
		temp = a[pass]
		for i > 0 && temp < a[i-1]{ // ascending
			a[i] = a[i-1]
			i = i - 1
		}
		a[i] = temp
		pass = pass + 1
	}
}