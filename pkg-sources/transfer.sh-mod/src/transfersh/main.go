package main

import "transfersh/cmd"

func main() {
	app := cmd.New()
	app.RunAndExitOnError()
}
