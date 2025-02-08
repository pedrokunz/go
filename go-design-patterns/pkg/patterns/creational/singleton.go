package creational

// GameStateSingleton is a type that implements the GameStateSingleton design pattern.
type GameStateSingleton struct {
	//TODO create a struct for player
	Health    int      //TODO define a value object for health
	Inventory []string //TODO define a value object for inventory
}

var instance *GameStateSingleton

// GetGameState returns the single instance of the Singleton type.
func GetGameState() *GameStateSingleton {
	if instance == nil {
		instance = &GameStateSingleton{
			Health:    100,
			Inventory: make([]string, 0),
		}
	}

	return instance
}
