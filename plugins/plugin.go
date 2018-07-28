package plugins

import "plugin"

func Plugin(input string) {
	var lang string
	switch input {
	case "thai":
		lang = "thai.so"
	case "eng":
		lang = "eng.so"
	default:
		lang = "eng.so"
	}

	//load module

	plug, err := plugin.Open(lang)
	if err != nil {
		panic(err)
	}

	plugHelloWorld, err := plug.Lookup("HelloWorld")
	if err != nil {
		panic(err)
	}

	helloWorld := plugHelloWorld.(HelloWorld)

	helloWorld.Print("Chatchai Vichai")
}
