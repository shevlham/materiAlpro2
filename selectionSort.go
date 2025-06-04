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
	selectsort(&a, ndata)
	for i = 0; i < ndata; i++ {
		fmt.Print(a[i]," ")
	}
}

func selectsort(a *tabInt, n int){
	var i, pass, temp, idx int
	
	pass = 1
	for pass < n {
		i = pass
		idx = pass - 1
		for i < n {
			if a[i] < a[idx]{ // ascending
				idx = i
			}
			i = i + 1
		}
		temp = a[pass-1]
		a[pass-1] = a[idx]
		a[idx] = temp
		pass = pass + 1
	}
}