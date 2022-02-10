package main

import "github.com/chrislusf/gleam/gio"

func main() {
	gs := gio.NewSettings("ca.desrt.dconf-editor.Demo")
	str := gs.GetString("string")

	gs.SetString("string", str)

	int0 := gs.GetInt("integer-32-signed")

	gs.SetInt("integer-32-signed", int0)
	gs.Unref()

}
