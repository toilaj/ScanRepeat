package main

import (
	"CheckRepeat/cmd"
	_ "CheckRepeat/scan"
)

func main() {
	cmd.InitCmd()
	cmd.Execute()
}
