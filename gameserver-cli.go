package main

import (
	"fmt"

	"github.com/ngoldack/gameserver-cli/internal"
)

func main() {
	fmt.Println("Hello world!")

	t, _ := internal.LoadTemplate("templates/minecraft.yaml")
	gs, _ := internal.GenerateServer(t, nil, nil)
	gs.Name = "minecraft-server"
	internal.ConvertToComposeFile(gs)
}
