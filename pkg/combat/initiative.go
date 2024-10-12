package combat

import (
    "fmt"
    "math/rand"

   "github.com/ZazaRy/unicornXE/pkg/characters"
)


type Teams struct {
    Combatants []characters.BaseCombatant
}

// Are All Alive?Come on, this one was easy..My only comment in the entire codebase, SMFH.....
func (t *Teams) AAA() bool{
    status := All(t.Combatants, func(x characters.BaseCombatant) bool{ return x.GetHP() > 0})
    return status
}

func (t *Teams) AAnA() bool{
    status := Any(t.Combatants, func(x characters.BaseCombatant) bool{ return x.GetHP() > 0})
    return status
}

func TeamSplitter(t *Teams) ([]characters.BaseCombatant, []characters.BaseCombatant){
    length := len(t.Combatants)
    teamOne := make([]characters.BaseCombatant, 0, length/2 )
    teamTwo := make([]characters.BaseCombatant, 0, length/2 )
    for i:=0;i<length;i++{
        if t.Combatants[i].Team == 1{
           teamOne = append(teamOne, t.Combatants[i])
        }else{
            teamTwo = append(teamTwo, t.Combatants[i])
        }
    }
    return teamOne, teamTwo
}


func (t *Teams) RandomPicker(attackerTeamID int ) (int, error){
    length := len(t.Combatants)
    toPickFrom := make([]int, 0, length/2)
    if t.AAnA(){
        for i:=0; i<length;i++{
            if t.Combatants[i].Team != attackerTeamID && t.Combatants[i].HP != 0{
                toPickFrom = append(toPickFrom, i)
            }
        }
        if len(toPickFrom) > 0{
            return toPickFrom[rand.Intn(len(toPickFrom))], nil
        }
    }
    return 0, fmt.Errorf("No valid targets, Team %v wins", attackerTeamID)
}

func InitOrder(teams []characters.BaseCombatant, size int) []characters.BaseCombatant{
    for i:=0; i<size; i++{
        teams[i].RollInitiative()
    }
    QuickSort(teams, 0, size-1)
    return teams
}


func NextTurn(teams []characters.BaseCombatant) func() characters.BaseCombatant{
    turnCount := 0
    return func() characters.BaseCombatant{
        combIndex := turnCount % len(teams)
        turnCount += 1
        fmt.Printf("Turn count is: %v'\n'", turnCount)
        return teams[combIndex]
    }
}

type InitUtil interface{
    AAA() bool
    AAnA() bool
}
