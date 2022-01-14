package main

import (
	"context"
	"testing"

	"github.com/kelindar/lua"
)

var data = map[string]interface{}{
	"customKey1": 12, "time": 12999, "customKey2": "Key2Val", "city": "SG", "customKey3": 999, "android": "yes",
	"customKey15": 5555, "uuid": 12999, "customKey12": "Key2Val", "country": "SG", "customKey13": 999, "ios": "no",
	"customKey25": 5555, "newuuid": 12999, "customKey22": "Key2Val", "region": "SG", "customKey23": 999, "device": "ios",
	"date": "03031999",
}

func init() {
	pluginInit()
}
func BenchmarkNativeFunctionCall(b *testing.B) {
	for n := 0; n < b.N; n++ {
		nativeFunctionCall(data)
	}
}
func BenchmarkPluginCall(b *testing.B) {
	for n := 0; n < b.N; n++ {
		pluginCall(data)
	}
}

func BenchmarkPluginCallAfterInit(b *testing.B) {
	for n := 0; n < b.N; n++ {
		pluginCallAfterInit(data)
	}
}

func BenchmarkLua(b *testing.B) {
	for n := 0; n < b.N; n++ {
		luaCall(data)
	}
}
func BenchmarkLuaCallAfterInit(b *testing.B) {
	// Cold start, only happens once
	s, _ := lua.FromString("xx.lua", `
	local json = require("json")
	function main(input)
		return json.encode(input)
	end
`)

	b.ReportAllocs()
	b.ResetTimer()

	// Actual benchmark code that executes the function,
	for n := 0; n < b.N; n++ {
		_, err := s.Run(context.Background(), data)
		if err != nil {
			panic(err)
		}
	}
}
