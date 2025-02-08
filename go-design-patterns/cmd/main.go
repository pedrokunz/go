package main

import (
    "fmt"
    "go-design-patterns/game"
)

func main() {
    // Initialize the game engine
    engine := game.Engine{}
    engine.Start()

    // Create and initialize the game
    g := game.Game{}
    g.Initialize()

    // Start the game loop
    fmt.Println("Starting the game...")
    g.Play()

    // End the game
    g.End()
    engine.Stop()
}