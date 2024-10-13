package gridMap

import (
    "fmt"
    "math"
    "github.com/ZazaRy/unicornXE/pkg/characters"

	"github.com/ZazaRy/unicornXE/pkg/combat"
)




type Grid struct {
    Width int
    Height int
    Occupants []combat.Teams
}


func (g *Grid) GetDistance(p1, p2 characters.BaseCombatant) float64{
    return math.Sqrt(math.Pow(float64(p2.Coords.X - p1.Coords.X), 2) + math.Pow(float64(p2.Coords.Y - p1.Coords.Y), 2))
}



func (g *Grid) GetNearestEnemy(attacker *characters.BaseCombatant) *characters.BaseCombatant{
    teamsLength := len(g.Occupants)
    checkPoses := make([]characters.BaseCombatant, 0, teamsLength )

}
