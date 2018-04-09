package main

import "fmt"

func main() {
	var ds *Datastore
	ds.Save("client", "John Stoll")

	ds = &Datastore{}
	ds.Save("client", "Jamie Locke")
}

type Datastore struct{}

func (ds *Datastore) Save(k, v string) {
	if ds == nil {
		fmt.Println("saving to /dev/null")
	} else {
		fmt.Println("saving to the database")
	}
}
