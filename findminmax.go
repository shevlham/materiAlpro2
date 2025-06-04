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
	fmt.Println("min =",min(a, ndata))
	fmt.Println("max =",max(a, ndata))
}

func min(a tabInt, n int)int {
	var i, min int
	min = 0
	for i = 1; i < n; i++{
		if a[min] > a[i]{
			min = i
		}
	}
	return a[min]
}

func max(a tabInt, n int)int {
	var i, max int
	max = 0
	for i = 1; i < n; i++{
		if a[max] < a[i]{
			max = i
		}
	}
	return a[max]
}