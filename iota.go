package main

import "fmt"

func main(){
	const (
		a = iota
		b
		c
		d = "li"
		e
		f
		g = "hai"
		h
		i = iota
	)
	fmt.Println(a,b,c,d,e,f,g,h,i)
	const (
		p = iota
		k
		m
		n
	
	)
	fmt.Println(p,k,m,n)
}
