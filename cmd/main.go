package main


import (
    "fmt"
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
        Initiative: 1,
    }
    player2 := characters.BaseCombatant{
        Name: "Player 2",
        HP: 10,
        AC: 10,
        Attack_roll: 10,
        Damage: 4.5,
        Initiative: 2,
    }


    team := combat.Teams{}
    team.AddToTracker(player1)
    team.AddToTracker(player2)

    nextTurn := combat.RollInitiative(team)


    fmt.Println("Turn 1: ", nextTurn())

    player1.Attack(&player2)
    player2.Attack(&player1)
}
