package main

import (
	"context"
	"encoding/json"
	"fmt"
	"plugin"

	"github.com/kelindar/lua"
)

func luaCall(data map[string]interface{}) interface{} {
	s, _ := lua.FromString("xx.lua", `
	local json = require("json")
	function main(input)
		return json.encode(input)
	end
`)
	res, err := s.Run(context.Background(), data)
	if err != nil {
		panic(err)
	}
	return res.String()
}

var fGlobal func(map[string]interface{}) interface{}

func pluginInit() {
	p, err := plugin.Open("./talaria_go_function.so")
	if err != nil {
		panic(err)
	}
	f, err := p.Lookup("ComputeRow")
	if err != nil {
		panic(err)
	}
	fGlobal = f.(func(map[string]interface{}) interface{})
}

func pluginCallAfterInit(input map[string]interface{}) interface{} {
	return fGlobal(input)
}

func pluginCall(input map[string]interface{}) interface{} {
	p, err := plugin.Open("./talaria_go_function.so")
	if err != nil {
		panic(err)
	}
	f, err := p.Lookup("ComputeRow")
	if err != nil {
		panic(err)
	}
	return f.(func(map[string]interface{}) interface{})(input)
}

var facetsDefined = []string{
	"time", "data", "id", "uuid", "ios", "android", "apple",
	"city", "country", "device", "date", "mobile",
	"newcity", "newcountry", "newdevice", "newdate", "newmobile",
}

func isKeyInFacet(key string) bool {
	for _, facet := range facetsDefined {
		if key == facet {
			return true
		}
	}
	return false
}

//save it in data
func nativeFunctionCall(rowArg map[string]interface{}) interface{} {
	data := make(map[string]interface{})
	for key, val := range rowArg {
		if !isKeyInFacet(key) {
			data[key] = val
			delete(rowArg, key)
		}
	}
	res, _ := json.Marshal(data)
	return string(res)
}

func main() {
	fmt.Println("\n====================")
	fmt.Println("\n====================")
}
