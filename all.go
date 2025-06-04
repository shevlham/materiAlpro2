package main
import "fmt"

const NMAX int = 100

type tabInt [NMAX]int

func main(){
	var a tabInt
	var i, ndata int
	//var cari int
	
	fmt.Scan(&ndata)
	for i = 0; i < ndata; i++ {
		fmt.Scan(&a[i])
	}
	
	//fmt.Scan(&cari)
	//fmt.Println(seqsearch(a, ndata, cari))
	//fmt.Println(findmax(a, ndata))
	//fmt.Println(findmin(a, ndata))
	//fmt.Println(binsearch(a, ndata, cari))
	//selsort(&a, ndata)
	insesort(&a, ndata)

	
	for i = 0; i < ndata; i++ {
		fmt.Print(a[i], " ")
	}
}

func seqsearch(a tabInt, n, cari int)int{
	var i,idx int
	idx = -1
	for i < n && idx == -1 {
		if a[i] == cari {
			idx = i
		}
		i++
	}
	return idx
}

func binsearch(a tabInt, n, cari int)int{
	var l,r,m int
	l = 0
	r = n-1
	m = (l+r)/2
	for l <= r && a[m] != cari {
		if cari > a[m] {
			l = m + 1 
		} else {
			r = m - 1
		}
		m = (l+r)/2
	}
	return m
}

func findmax(a tabInt, n int)int{
	var i, maks int
	maks = a[0]
	i = 1
	for i < n {
		if maks < a[i] {
			maks = a[i]
		}
		i++
	}
	return maks
}

func findmin(a tabInt, n int)int{
	var i, min int
	min = a[0]
	i = 1
	for i < n {
		if min > a[i] {
			min = a[i]
		}
		i++
	}
	return min
}

func selsort(a *tabInt, n int){
	var idx, temp, pass, i int
	
	pass = 1
	for pass < n {
		i = pass
		idx = pass - 1
		for i < n {
			if a[idx] > a[i]{
				idx = i
			}
			i++
		}
		temp = a[pass-1]
		a[pass-1] = a[idx]
		a[idx] = temp
		pass++
	}
}

func insesort(a *tabInt, n int){
	var i, pass, temp int
	
	pass = 1
	for pass < n {
		i = pass
		temp = a[pass]
		for i > 0 && temp < a[i-1] {
			a[i] = a[i-1]
			i--
		}
		a[i] = temp
		pass++
	}
}

// 7 5 2 9 8 
// _ 7 2 9 8
// 0 1 2 3 4

