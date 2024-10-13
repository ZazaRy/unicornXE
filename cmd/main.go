package main

import (
	"fmt"
	_"math/rand"

	"github.com/ZazaRy/unicornXE/pkg/characters"
	"github.com/ZazaRy/unicornXE/pkg/combat"
    "github.com/ZazaRy/unicornXE/pkg/gridMap"
)




func main() {
    player1 := characters.BaseCombatant{
        Name: "Player 1",
        HP: 10,
        AC: 10,
        Attack_roll: 10,
        Damage: 4.5,
        Initiative: 2,
        Team: 1,
        Coords: characters.Coords{X: 100, Y: 10},
    }
    player2 := characters.BaseCombatant{
        Name: "Player 2",
        HP: 10,
        AC: 10,
        Attack_roll: 10,
        Damage: 4.5,
        Initiative: 2,
        Team: 1,
        Coords: characters.Coords{X: 110, Y: 10},
    }
    player3 := characters.BaseCombatant{
        Name: "Player 3",
        HP: 10,
        AC: 10,
        Attack_roll: 10,
        Damage: 4.5,
        Initiative: 2,
        Team: 2,
        Coords: characters.Coords{X: 105, Y: 10},
    }

    player4 := characters.BaseCombatant{
        Name: "Player 4",
        HP: 10,
        AC: 10,
        Attack_roll: 10,
        Damage: 4.5,
        Initiative: 2,
        Team: 2,
        Coords: characters.Coords{X:115, Y: 10},
    }

    player5 := characters.BaseCombatant{
        Name: "Player 5",
        HP: 10,
        AC: 10,
        Attack_roll: 10,
        Damage: 4.5,
        Initiative: 2,
        Team: 1,
        Coords: characters.Coords{X:95, Y: 10},
    }

    player6 := characters.BaseCombatant{
        Name: "Player 6",
        HP: 10,
        AC: 10,
        Attack_roll: 10,
        Damage: 4.5,
        Initiative: 2,
        Team: 1,
        Coords: characters.Coords{X: 90, Y: 10},
    }

    player7 := characters.BaseCombatant{
        Name: "Player 7",
        HP: 10,
        AC: 10,
        Attack_roll: 10,
        Damage: 4.5,
        Initiative: 2,
        Team: 2,
        Coords: characters.Coords{X: 120, Y: 10},
    }

    player8 := characters.BaseCombatant{
        Name: "Player 8",
        HP: 10,
        AC: 10,
        Attack_roll: 10,
        Damage: 4.5,
        Initiative: 2,
        Team: 2,
        Coords: characters.Coords{X: 85, Y: 10},
    }

    player9 := characters.BaseCombatant{
        Name: "Player 9",
        HP: 10,
        AC: 10,
        Attack_roll: 10,
        Damage: 4.5,
        Initiative: 2,
        Team: 1,
        Coords: characters.Coords{X: 90, Y: 20},
    }

    player10 := characters.BaseCombatant{
        Name: "Player 10",
        HP: 10,
        AC: 10,
        Attack_roll: 10,
        Damage: 4.5,
        Initiative: 2,
        Team: 1,
        Coords: characters.Coords{X:100, Y: 20},
    }


    combatants := make([]characters.BaseCombatant, 0, 10)
    combatants = append(combatants, player1, player2, player3, player4, player5, player6, player7, player8, player9, player10)

    grid := gridMap.Grid{
        Width: 600,
        Height: 600,
        Occupants: combatants,
    }

    team := combat.Teams{Combatants: combatants}
    order := combat.InitOrder(combatants, len(combatants))

    turnFunc := combat.NextTurn(order)
    aliveFunc := team.AAA

    var count int


    for aliveFunc() == true {
        count += 1
        nextAttacker := turnFunc()

        opposingTeam, errOP := team.GetOpposingTeam(nextAttacker.Team)
        if errOP != nil {
            fmt.Println(errOP)
            break
        }

        nextAttackerCoords := map[int]characters.Coords{}
        nextAttackerCoords[nextAttacker.Team] = nextAttacker.Coords

        nearestDefender := grid.GetNearestEnemy(nextAttackerCoords, opposingTeam)

        updatedDefender := nextAttacker.Attack(nearestDefender)

        for i, comb := range grid.Occupants {
            if comb.Name == updatedDefender.Name {
                grid.Occupants[i] = updatedDefender
            }
        }
    }
}
