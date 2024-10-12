package main

import (
	"fmt"
	_"math/rand"

	"github.com/ZazaRy/unicornXE/pkg/characters"
	"github.com/ZazaRy/unicornXE/pkg/combat"
	_ "github.com/ZazaRy/unicornXE/pkg/combat"
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
    }
    player2 := characters.BaseCombatant{
        Name: "Player 2",
        HP: 10,
        AC: 10,
        Attack_roll: 10,
        Damage: 4.5,
        Initiative: 2,
        Team: 1,
    }
    player3 := characters.BaseCombatant{
        Name: "Player 3",
        HP: 10,
        AC: 10,
        Attack_roll: 10,
        Damage: 4.5,
        Initiative: 2,
        Team: 2,
    }

    player4 := characters.BaseCombatant{
        Name: "Player 4",
        HP: 10,
        AC: 10,
        Attack_roll: 10,
        Damage: 4.5,
        Initiative: 2,
        Team: 2,
    }

    player5 := characters.BaseCombatant{
        Name: "Player 5",
        HP: 10,
        AC: 10,
        Attack_roll: 10,
        Damage: 4.5,
        Initiative: 2,
        Team: 1,
    }

    player6 := characters.BaseCombatant{
        Name: "Player 6",
        HP: 10,
        AC: 10,
        Attack_roll: 10,
        Damage: 4.5,
        Initiative: 2,
        Team: 1,
    }

    player7 := characters.BaseCombatant{
        Name: "Player 7",
        HP: 10,
        AC: 10,
        Attack_roll: 10,
        Damage: 4.5,
        Initiative: 2,
        Team: 2,
    }

    player8 := characters.BaseCombatant{
        Name: "Player 8",
        HP: 10,
        AC: 10,
        Attack_roll: 10,
        Damage: 4.5,
        Initiative: 2,
        Team: 2,
    }

    player9 := characters.BaseCombatant{
        Name: "Player 9",
        HP: 10,
        AC: 10,
        Attack_roll: 10,
        Damage: 4.5,
        Initiative: 2,
        Team: 1,
    }

    player10 := characters.BaseCombatant{
        Name: "Player 10",
        HP: 10,
        AC: 10,
        Attack_roll: 10,
        Damage: 4.5,
        Initiative: 2,
        Team: 1,
    }
    combatants := make([]characters.BaseCombatant, 0, 10)
    combatants = append(combatants, player1, player2, player3, player4, player5, player6, player7, player8, player9, player10)
    team := combat.Teams{Combatants: combatants}
    order := combat.InitOrder(combatants, len(combatants))
    turnFunc := combat.NextTurn(order)

    aliveFunc := team.AAA

    for aliveFunc() == true{

        nextAttacker := turnFunc()
        defenderIndex, err := team.RandomPicker(nextAttacker.GetTeam())
        defender := &team.Combatants[defenderIndex]
        if err!=nil{
            fmt.Println(err)
            break
        }else{
        nextAttacker.Attack(defender)
        }
    }
}
