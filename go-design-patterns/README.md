# Go Design Patterns

This project is a study of various design patterns implemented in Go, organized in a gamified structure. The goal is to understand and apply different design patterns through a simple game engine.

## Project Structure

```
go-design-patterns
├── cmd
│   └── main.go               # Entry point of the application
├── pkg
│   ├── patterns
│   │   ├── creational
│   │   │   └── singleton.go  # Singleton design pattern implementation
│   │   ├── structural
│   │   │   └── adapter.go     # Adapter design pattern implementation
│   │   └── behavioral
│   │       └── observer.go    # Observer design pattern implementation
├── game
│   ├── engine.go             # Game engine logic
│   ├── player.go             # Player structure and actions
│   └── game.go               # Game flow orchestration
├── go.mod                     # Module definition and dependencies
└── README.md                  # Project documentation
```

## Getting Started

To run the project, follow these steps:

1. Clone the repository:

   ```
   git clone <repository-url>
   cd go-design-patterns
   ```

2. Install dependencies:

   ```
   go mod tidy
   ```

3. Run the application:

   ```
   go run cmd/main.go
   ```

## Design Patterns Implemented

- **Creational Patterns**
  - **Singleton**: Ensures a class has only one instance and provides a global point of access to it.

- **Structural Patterns**
  - **Adapter**: Allows incompatible interfaces to work together by converting the interface of a class into another interface that clients expect.

- **Behavioral Patterns**
  - **Observer**: Defines a one-to-many dependency between objects so that when one object changes state, all its dependents are notified and updated automatically.

## Contributing

Feel free to contribute by adding more design patterns or improving the existing implementations. Create a new branch for your feature and submit a pull request.

## License

This project is licensed under the MIT License. See the LICENSE file for details.

You can create a simple text-based adventure game. Here's a basic idea:

### Game Idea: "Adventure Quest"

#### Overview

- The player navigates through different rooms in a dungeon.
- Each room can have different items, enemies, or puzzles.
- The player can collect items, fight enemies, and solve puzzles to progress.

#### Design Patterns to Use

1. **Singleton**: Manage the game state (e.g., player's health, inventory).
2. **Factory Method**: Create different types of rooms (e.g., treasure room, enemy room).
3. **Observer**: Notify the player when certain events occur (e.g., enemy appears, item found).
4. **Adapter**: Integrate different types of input methods (e.g., keyboard, voice commands).

#### Example Structure

```filetree
go-design-patterns
└── go-design-patterns
|   ├── cmd
|       └── main.go
|   ├── pkg
|       └── patterns
|           ├── creational
|               └── singleton.go
|           ├── structural
|               └── adapter.go
|           └── behavioral
|               └── observer.go
|   ├── game
|       ├── engine.go
|       ├── player.go
|       ├── room.go
|       └── game.go
|   ├── go.mod
|   └── README.md
```

#### Example Code Snippets

**Singleton for Game State:**

```go
package creational

type GameState struct {
 Health   int
 Inventory []string
}

var instance *GameState

func GetGameState() *GameState {
 if instance == nil {
  instance = &GameState{Health: 100, Inventory: []string{}}
 }
 return instance
}
```

**Factory Method for Rooms:**

```go
package game

type Room interface {
 Enter() string
}

type TreasureRoom struct{}

func (r *TreasureRoom) Enter() string {
 return "You found a treasure!"
}

type EnemyRoom struct{}

func (r *EnemyRoom) Enter() string {
 return "An enemy appears!"
}

func RoomFactory(roomType string) Room {
 switch roomType {
 case "treasure":
  return &TreasureRoom{}
 case "enemy":
  return &EnemyRoom{}
 default:
  return nil
 }
}
```

**Observer for Events:**

```go
package behavioral

type Observer interface {
 Update(event string)
}

type Subject struct {
 observers []Observer
}

func (s *Subject) Attach(observer Observer) {
 s.observers = append(s.observers, observer)
}

func (s *Subject) Notify(event string) {
 for _, observer := range s.observers {
  observer.Update(event)
 }
}
```

**Adapter for Input Methods:**

```go
package structural

type InputMethod interface {
 GetInput() string
}

type Keyboard struct{}

func (k *Keyboard) GetInput() string {
 return "Keyboard input"
}

type VoiceCommand struct{}

func (v *VoiceCommand) GetInput() string {
 return "Voice command input"
}

type InputAdapter struct {
 method InputMethod
}

func NewInputAdapter(method InputMethod) *InputAdapter {
 return &InputAdapter{method: method}
}

func (a *InputAdapter) GetInput() string {
 return a.method.GetInput()
}
```

This structure and these patterns will help you build a robust and extensible text-based adventure game.
