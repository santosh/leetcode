package main

import "fmt"

type DVD struct {
	Name        string
	ReleaseYear int
	Director    string
}

type dvdColletion [15]DVD

func (d DVD) toString() {
	if d.Name == "" && d.Director == "" && d.ReleaseYear == 0 {
		fmt.Println(nil)
	} else {
		fmt.Printf("%s, directed by %s, released in %d.\n", d.Name, d.Director, d.ReleaseYear)
	}

}

func main() {
	coll := dvdColletion{}

	avengersDVD := DVD{"The Avengers", 2012, "Joss Whedon"}
	coll[7] = avengersDVD
	incrediblesDVD := DVD{"The Incredibles", 2004, "Brad Bird"}
	coll[3] = incrediblesDVD
	findingDoryDVD := DVD{"Finding Dory", 2016, "Andrew Stanton"}
	coll[9] = findingDoryDVD
	lionKingDVD := DVD{"The Lion King", 2019, "Jon Favreau"}
	coll[2] = lionKingDVD

	starWarsDVD := DVD{"Star Wars", 1977, "George Lucus"}
	coll[3] = starWarsDVD

	coll[7].toString()
	coll[10].toString()
	coll[3].toString()
}
