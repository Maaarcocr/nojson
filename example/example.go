package main

import (
	"fmt"

	"github.com/Maaarcocr/nojson"
)

type Post struct {
	User     string
	Likes    int      `nojson:"user"`
	Comments []string `nojson:"user;mod"`
}

func main() {
	examplePost := Post{"Marco", 5, []string{"good", "i like it"}}
	resultForUser, _ := nojson.MarshalAndFilterBy("user", examplePost)
	fmt.Println(string(resultForUser))
	resultForMod, _ := nojson.MarshalAndFilterBy("mod", examplePost)
	fmt.Println(string(resultForMod))
	examplePost2 := Post{"Luke", 4, []string{"bad", "no"}}
	examplePostsSlice := []Post{examplePost, examplePost2}
	resultArray, _ := nojson.MarshalAndFilterBy("user", examplePostsSlice)
	fmt.Println(string(resultArray))
}
