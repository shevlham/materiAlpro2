package main
import "fmt"

const NMAX int = 100

type tabInt [NMAX]int

func main(){
	var a tabInt
	var i, ndata, cari int
	
	fmt.Scan(&ndata)
	for i = 0; i < ndata; i++ {
		fmt.Scan(&a[i])
	}
	fmt.Scan(&cari)
	fmt.Println(seqsearch(a, ndata, cari))
}

func seqsearch(a tabInt, n, cari int)bool{
	var i int
	
	for i < n {
		if a[i] == cari {
			return true
		}
		i = i + 1
	}
	return false
}