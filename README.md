# nojson
A Simple Golang library created in order to escape some Struct Fields easily while marshaling a custom struct object or an array of custom struct objects 
## Why I have created this library
While I was developing a REST API server I needed to hide some struct field values from the client if he was not allowed to see that part of the data.
I was using an ORM library so I needed an easy way to create a filtered json response.

##Example

```go
import (
	"fmt"
	
	"github.com/Maaarcocr/nojson"
)

type Post struct{
	User string
	Likes int `nojson:"user"`
	Comments []string `nojson:"user;mod"`
}

func main(){
	examplePost := Post{"Marco", 5, []string{"good", "i like it"}}
	resultForUser, err1 := nojson.MarshalAndFilterBy("user", examplePost)
	resultForMod, err2 := nojson.MarshalAndFilterBy("mod", examplePost)
	if err1 == nil {
		fmt.Println(string(resultForUser))
	}
	if err2 == nil {
		fmt.Println(string(resultForMod))
	}
}
```
User Result:
```json
{"User": "Marco"}
```
Mod Result: 
```
{"Likes": 5, "User": "Marco"}
```
##NB
1. The function MarshalAndFilterBy accept struct, slice and array as types. (If you use different types it won't complain at compile time, but you will have a nil result as result)
2. The name used in the JSON result will be the Field Name (so there is no point in using the json tag to set a different name)
3. The function MarshalAndFilterBy has []byte as return type.

##PS
You can find other examples in the example folder
