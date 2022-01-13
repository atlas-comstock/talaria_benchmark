package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/kelindar/lua"
)

var keysToExclude = []string{
	"key1", "key2", "key3", "key4", "key5",
	"key6", "key7", "key8", "key9", "key10",
}

const jsonstr = `{"id": "userName", "key1": "11111", "city": "cityValue", "uuid": "99999", "loc": "12.676689147949219, 101.20865195568277", "addtion": "999999999999999999999999999999999", "key10": "10000000"}`

func trim() string {
	e := map[string]string{}
	if err := json.Unmarshal([]byte(jsonstr), &e); err != nil {
		panic(err)
	}
	for _, k := range keysToExclude {
		if _, exist := e[k]; exist {
			delete(e, k)
		}
	}
	s, err := json.Marshal(e)
	if err != nil {
		panic(err)
	}
	return string(s)
}

func luaTrim() string {
	s, err := lua.FromString("xx.lua", `
	local json = require("json")
	local keys = {
	'key1', 'key2', 'key3', 'key4', 'key5',
	'key6', 'key7', 'key8', 'key9', 'key10',
	}
	function main(input)
		local data = json.decode(input)
		for i, key in ipairs(keys) do
			data[key] = nil
		end
		return json.encode(data)
	end
`)
	if err != nil {
		panic(err)
	}

	result, err := s.Run(context.Background(), jsonstr)
	if err != nil {
		panic(err)
	}
	return result.String()
}

func main() {
	fmt.Println("=================")
	fmt.Println("trim: ", trim())
	fmt.Println("luaTrim: ", luaTrim())
}
