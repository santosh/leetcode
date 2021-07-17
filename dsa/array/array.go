package main

import "fmt"

type DVD struct {
	Name        string
	ReleaseYear int
}

type dvdCollection []DVD

func (d DVD) toString() {
	if d.Name == "" && d.ReleaseYear == 0 {
		fmt.Println(nil)
	} else {
		fmt.Printf("%s, released in %d.\n", d.Name, d.ReleaseYear)
	}
}

func (d dvdCollection) traverse() {
	for _, v := range d {
		v.toString()
	}
}

func (d *dvdCollection) append(item DVD)  {
	d.push(item)
}

func (d *dvdCollection) push(item DVD)  {
	*d = append(*d, item)
}

func (d *dvdCollection) pop() (lastItem DVD) {
	v := *d
	lastItem = v[len(v)-1]
	(*d)[len(v)-1] = DVD{}
	return
}

func (d *dvdCollection) insert(index int, elem DVD) {
	v := *d
	for _, v := range v {
		fmt.Println(v)
	}
}

// DeleteFromIndex: (index)
// https://www.youtube.com/watch?v=Bnjbun-hiBk&list=PLdo5W4Nhv31bbKJzrsKfMpo_grxuLl8LU&index=2
func main() {
	coll := dvdCollection{}

	incrediblesDVD := DVD{Name: "The Incredibles", ReleaseYear: 2004}
	coll.append(incrediblesDVD)
	avengersDVD := DVD{Name: "The Avengers", ReleaseYear: 2012}
	coll.push(avengersDVD)
	findingDoryDVD := DVD{Name: "Finding Dory", ReleaseYear: 2016}
	coll.push(findingDoryDVD)
	lionKingDVD := DVD{Name: "The Lion King", ReleaseYear: 2019}
	coll.append(lionKingDVD)


	// coll.traverse()
	// fmt.Println("============")

	// coll.pop()
	starWarsDVD := DVD{Name: "Star Wars", ReleaseYear: 1977}
	coll.insert(0, starWarsDVD)

	coll.traverse()
}
