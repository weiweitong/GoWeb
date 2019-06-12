package main

import (
	"encoding/json"
	"fmt"
	"os"
)

var Pl = fmt.Println

type Response1 struct {
	Page int
	Fruits []string
}

type Response2 struct {
	Page int	`json: "page"`
	Fruits []string `json: fruits`
}

func main()  {
	bolB, _ := json.Marshal(true)
	Pl(string(bolB))

	intB, _ := json.Marshal(1)
	Pl(string(intB))

	fltB, _ := json.Marshal(2.34)
	Pl(string(fltB))

	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))

	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	Pl(string(slcB))

	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	Pl(string(mapB))

	res1D := &Response1{
		Page:	1,
		Fruits:	[]string{"apple", "peach", "pear"},
	}
	res1B, _ := json.Marshal(res1D)
	Pl(string(res1B))

	res2D := Response2{
		Page:	1,
		Fruits:	[]string{"apple", "peach", "pear"},
	}
	res2B, _ := json.Marshal(res2D)
	Pl(string(res2B))

	byt := []byte(`{"num":6.13, "strs":["a", "b"]}`)

	var dat map[string]interface{}

	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	Pl(dat)

	num := dat["num"].(float64)
	Pl(num)

	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	Pl(str1)

	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := &Response2{}
	json.Unmarshal([]byte(str), &res)
	Pl(res)
	Pl(res.Fruits[0])

	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce":7}
	enc.Encode(d)

}

