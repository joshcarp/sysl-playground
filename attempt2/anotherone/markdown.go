package main

import (
	"fmt"

	"github.com/Joshcarp/sysl-playground/attempt2/files"
	"github.com/Joshcarp/sysl_testing/pkg/command"
	"github.com/sirupsen/logrus"
)

func main() {
	input := `
MobileApp:
	Login:
		Server <- Login
	!type LoginData:
		username <: string
		password <: string
	!type LoginResponse:
		message <: string
Server:
	Login(data <: MobileApp.LoginData):
		return MobileApp.LoginResponse`
	Render(input)

}

// Render implements the vecty.Component interface.
func Render(input string) {
	// Render the markdown input into HTML using Blackfriday.
	// unsafeHTML := blackfriday.Run([]byte(m.Input))

	fs := files.NewFs("root")
	f, err := fs.Create("tmp.sysl")
	check(err)

	_, e := f.Write([]byte(input))
	check(e)

	var logger = logrus.New()
	fmt.Println(logger)
	fmt.Println("this")
	// rc := 0
	rc := command.Main2([]string{"sysl", "pb", "-o", "project.textpb", "tmp.sysl"}, fs, logger, command.Main3)
	if rc != 0 {
		// panic(rc)
	}
	// g, err := fs.Create("project.svg")
	// check(err)
	// defer g.Close()
	svg, err := fs.Open("project.textpb")
	check(err)
	fmt.Println(svg)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
