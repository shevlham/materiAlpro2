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
	fmt.Println(binSearch(a, ndata, cari))
}

func binSearch(a tabInt, n, cari int)bool{
	var left, mid, right int
	
	left = 0
	right = n-1
	mid = (left + right)/2
	for left <= right && a[mid] != cari {
		if cari > a[mid] {
			left = mid + 1
		}else{
			right = mid - 1
		}
		mid = (left + right)/2
	}
	
	if a[mid] == cari {
		return true
	}else {
		return false
	}
}