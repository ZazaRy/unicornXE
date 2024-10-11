package main


import (
    _"fmt"
    "github.com/ZazaRy/unicornXE/pkg/characters"
    "github.com/ZazaRy/unicornXE/pkg/combat"
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

    team := combat.Teams{}
    team.AddToTracker(player1)
    team.AddToTracker(player2)
    team.AddToTracker(player3)
    team.AddToTracker(player4)


    alive := combat.All(team.Combatants, func(c interface{}) bool{
        combatant, ok := c.(characters.Combatant)
        if !ok {
            panic("All checks failed")
        }
        return combatant.GetHP() > 0
    })
}
