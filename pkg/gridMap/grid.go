package gridMap

import (
    "fmt"
    "math"
    "github.com/ZazaRy/unicornXE/pkg/characters"
    _"github.com/ZazaRy/unicornXE/pkg/utils"

)




type Grid struct {
    Width int
    Height int
    Occupants []characters.BaseCombatant
}


func (g *Grid) GetL2(xAtk, yAtk, xDef, yDef int) int{
    gridDistance := math.Sqrt(math.Pow(float64(xDef-xAtk), 2) + math.Pow(float64(yDef - yAtk), 2))
    convertedDistance := int(gridDistance)
    return convertedDistance
}



func (g *Grid) GetNearestEnemy(attacker map[int]characters.Coords, opposingTeam map[int][]int) characters.BaseCombatant{
    maxDistance := math.MaxInt32
    var nearestCombatant characters.BaseCombatant
    var attackerX, attackerY int
    for _, attackerPos := range attacker {
        attackerX = attackerPos.X
        attackerY = attackerPos.Y
    }
    for combatantIndex, pos := range opposingTeam {
        distance := g.GetL2(attackerX, attackerY, pos[0], pos[1])

        if distance <  maxDistance{
            maxDistance = distance
            nearestCombatant = g.Occupants[combatantIndex]
        }
        fmt.Printf("Attacker positions x: %d and y: %d\n", attackerX, attackerY)
        fmt.Printf("Defender x: %d and y: %d\n", pos[0], pos[1])
        fmt.Printf("Distance between attacker and defender is: %d\n", distance)
    }
    return nearestCombatant
}
