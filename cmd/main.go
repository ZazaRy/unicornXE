package main

import (
	"fmt"
	_"math/rand"

	"github.com/ZazaRy/unicornXE/pkg/characters"
	"github.com/ZazaRy/unicornXE/pkg/combat"
    "github.com/ZazaRy/unicornXE/pkg/gridMap"
)

var player1 = characters.BaseCombatant{
    Name: "Player1",
    HP: 10,
    AC: 20,
    Statarray: characters.StatArray{
        Str: 10,
        Dex: 10,
        Con: 10,
        Int: 10,
        Wis: 10,
        Cha: 10,
    },
    Proficiency: 2,
    PrimaryStat: characters.STR,
    NoAttacks: 1,
    Attack_roll: 10,
    Weapon: characters.Weapon{
        Name: "Sword",
        Damage: 6,
        Range: 1,
    },
    Damage: 6,
    Initiative: 10,
    Team: 1,
    Coords: characters.Coords{X: 30, Y: 10},
}


var player1 = characters.BaseCombatant{
    Name: "Player1",
    HP: 10,
    AC: 20,
    Statarray: characters.StatArray{
        Str: 10,
        Dex: 10,
        Con: 10,
        Int: 10,
        Wis: 10,
        Cha: 10,
    },
    Proficiency: 2,
    PrimaryStat: characters.STR,
    NoAttacks: 1,
    Attack_roll: 10,
    Weapon: characters.Weapon{
        Name: "Sword",
        Damage: 6,
        Range: 1,
    },
    Damage: 6,
    Initiative: 10,
    Team: 1,
    Coords: characters.Coords{X: 30, Y: 10},
}


var player4 = characters.BaseCombatant{
    Name: "Player4",
    HP: 10,
    AC: 20,
    Statarray: characters.StatArray{
        Str: 10,
        Dex: 10,
        Con: 10,
        Int: 10,
        Wis: 10,
        Cha: 10,
    },
    Proficiency: 2,
    PrimaryStat: characters.STR,
    NoAttacks: 1,
    Attack_roll: 10,
    Weapon: characters.Weapon{
        Name: "Sword",
        Damage: 6,
        Range: 1,
    },
    Damage: 6,
    Initiative: 10,
    Team: 2,
    Coords: characters.Coords{X: 30, Y: 10},
}

func main() {


    combatants := make([]characters.BaseCombatant, 0, 4)
    combatants = append(combatants, player1, player2, player3, player4)

    combMap := make(map[int]characters.BaseCombatant, 4)
    for i, v := range combatants {
        combMap[i] = v
    }

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

        opposingTeam, errOP := team.GetTeamCoords(nextAttacker.Team)
        if errOP != nil {
            fmt.Println(errOP)
            break
        }
        nearestDefender := grid.GetNearestEnemy(nextAttacker.Coords, opposingTeam)

        upDefenderID, updatedDefender := nextAttacker.Attack(nearestDefender)

        combatants[upDefenderID] = updatedDefender
    }
}
