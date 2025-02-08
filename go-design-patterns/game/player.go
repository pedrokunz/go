package game

type Player struct {
    Name  string
    score int
}

func (p *Player) Move(direction string) {
    // Logic for moving the player in the specified direction
}

func (p *Player) Score() int {
    return p.score
}

func (p *Player) UpdateScore(points int) {
    p.score += points
}