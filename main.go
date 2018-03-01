package main

import (
	"os"
	"sort"
)

func main() {
	fi, _ := os.Open(os.Args[1])
	defer fi.Close()
	d := Data{}
	d.Fill(fi)
	//fmt.Println("% v", d)
	//println("----")
	vs := make([]Vahicle, d.F)

	/*
		for i := range d.RS {
			print(i, " ")
		}
		println()


	*/
	sort.Slice(d.RS, func(i, j int) bool {
		return d.RS[i].S[0]+d.RS[i].S[1] < d.RS[j].S[0]+d.RS[j].S[1]
	})

	rs := make([]int, len(d.RS))
	for j, r := range d.RS {
		rs[j] = r.I
		//print(r.I, " ")
	}
	//println()
	//fmt.Println(rs)
	j := 0
	var r Ride
	var v Vahicle
	for len(rs)-j > 0 { //j, r = range d.RS { //
		//fmt.Printf("%#v %#v\n", d.RS, rs)
		//fmt.Println(len(rs), j, len(rs)-j == 0)
		r = d.RS[rs[j]]
		vf := false
		for i := 0; i < d.F; i++ {
			v = vs[i]
			tol := Lenght(v.X, v.Y, r.S[0], r.S[1]) + r.Lenght() + r.ES
			//println("to ", tol, v.S, r.ES)
			//println("total ", tol, v.S, r.LF)
			if v.S+tol >= r.LF || v.S+tol >= d.T {
				continue
			}

			vs[i].S += tol
			vs[i].RS = append(vs[i].RS, rs[j])
			vs[i].X, vs[i].Y = r.F[0], r.F[1]
			rs = append(rs[:j], rs[j+1:]...)
			//fmt.Printf("%d %v %d\n", j, rs, len(rs))
			if j > 0 {
				j--
			}

			vf = true
			break
		}
		if !vf {
			j++
		}

	}

	for _, v := range vs {
		print(len(v.RS), " ")
		for _, r := range v.RS {
			print(r, " ")
		}
		println()
	}
}
