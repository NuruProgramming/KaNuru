package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/NuruProgramming/KaNuru/repl"
)

var (
	Version = "v0.5.18"
	Author  = "by Nuru Org"
	NewLogo = Author + " | " + Version
	Help    = fmt.Sprintf(`💡 Namna ya kutumia Nuru:
	%s: Kuanza programu ya Nuru
	%s: Kuendesha faili la Nuru
	%s: Kusoma nyaraka za Nuru
	%s: Kufahamu toleo la Nuru
`,
		"nuru",
		"nuru jinaLaFile.nr",
		"nuru --nyaraka",
		"nuru --toleo")
)

func main() {

	args := os.Args
	if len(args) < 2 {

		fmt.Println(NewLogo)
		fmt.Println("💡 Tumia exit() au toka() kuondoka")
		repl.Start(os.Stdin, os.Stdout)
		return
	}

	if len(args) == 2 {
		switch args[1] {
		case "msaada", "-msaada", "--msaada", "help", "-help", "--help", "-h":
			fmt.Println(Help)
		case "version", "-version", "--version", "-v", "v", "--toleo", "-toleo":
			fmt.Println(NewLogo)
		default:
			file := args[1]

			if strings.HasSuffix(file, "nr") || strings.HasSuffix(file, ".sw") {
				contents, err := os.ReadFile(file)
				if err != nil {
					fmt.Println("Error: Nuru imeshindwa kusoma faili: ", args[1])
					os.Exit(1)
				}

				repl.Read(string(contents))
			} else {
				fmt.Println("'"+file+"'", "sii faili sahihi. Tumia faili la '.nr' au '.sw'")
				os.Exit(1)
			}
		}
	} else {
		fmt.Println("Error: Operesheni imeshindikana")
		fmt.Println(Help)
		os.Exit(1)
	}
}
