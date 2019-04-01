package main

import (
	"log"

	sdk "github.com/gaia-pipeline/gosdk"
)

// PrintParam prints out all given params.
func PrintParam(args sdk.Arguments) error {
	for _, arg := range args {
		log.Printf("Key: %s;Value: %s;\n", arg.Key, arg.Value)
	}
	return nil
}

func main() {
	jobs := sdk.Jobs{
		sdk.Job{
			Handler:     PrintParam,
			Title:       "Print Parameters",
			Description: "This job prints out all given params.",
			Args: sdk.Arguments{
				sdk.Argument{
					Description: "Username for the database schema:",
					// TextFieldInp displays a text field in the UI.
					// You can also use sdk.TextAreaInp for text area,
					// sdk.BoolInp for boolean input.
					Type: sdk.TextFieldInp,
					Key:  "username",
				},
				sdk.Argument{
					Description: "Description for username:",
					// TextFieldInp displays a text field in the UI.
					// You can also use sdk.TextAreaInp for text area and
					// sdk.BoolInp for boolean input.
					Type: sdk.TextAreaInp,
					Key:  "usernamedesc",
				},
				sdk.Argument{
					Type: sdk.VaultInp,
					Key:  "userpassword",
				},
			},
		},
	}

	// Serve
	if err := sdk.Serve(jobs); err != nil {
		panic(err)
	}
}
