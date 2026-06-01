package main 

import (
	"fmt"	
	"mail-infra-inspector-cli/internal/commands"


)

func main()  {
	flags := commands.NewCmdFlags()
	fmt.Println("Hello from main")
	flags.Execute()
}
