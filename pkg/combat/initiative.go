package combat

import (
	"fmt"
	"math/rand"

	"github.com/ZazaRy/unicornXE/pkg/characters"
	_"github.com/ZazaRy/unicornXE/pkg/gridMap"
    "github.com/ZazaRy/unicornXE/pkg/utils"
)


type Teams struct {
    Combatants []characters.BaseCombatant
}

// Are All Alive?Come on, this one was easy..My only comment in the entire codebase, SMFH.....
func (t *Teams) AAA() bool{
    status := utils.All(t.Combatants, func(x characters.BaseCombatant) bool{ return x.GetHP() > 0})
    return status
}

func (t *Teams) AAnA() bool{
    status := utils.Any(t.Combatants, func(x characters.BaseCombatant) bool{ return x.GetHP() > 0})
    return status
}

func (t *Teams) TeamSplitter() ([]characters.BaseCombatant, []characters.BaseCombatant){
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

func (t *Teams) GetOpposingTeam(attackerTeamID int) (map[int][]int, error){
    length := len(t.Combatants)
    opposingTeam := make(map[int][]int)
    if t.AAnA(){
        for i:=0; i < length; i++{
            if t.Combatants[i].Team != attackerTeamID && t.Combatants[i].HP > 0{
                opposingTeam[i] = []int{t.Combatants[i].Coords.X, t.Combatants[i].Coords.Y}
            }
        }
        if len(opposingTeam) > 0{
            return opposingTeam, nil
        }
    }
    return opposingTeam, fmt.Errorf("No valid targets, Team %v wins", attackerTeamID)
}


func (t *Teams) GetOppsingTeamByIndex(attackerTeamID int) ([]int, error){
    length := len(t.Combatants)
    opposingTeam := make([]int, 0, length/2)

    if t.AAnA(){
        for i:=0; i < length; i++{
            if t.Combatants[i].Team != attackerTeamID && t.Combatants[i].HP > 0{
                opposingTeam = append(opposingTeam, i)
            }
            if len(opposingTeam) > 0 {
                return opposingTeam, nil
            }
        }
    }
    return opposingTeam, fmt.Errorf("No valid targets, Team %v wins", attackerTeamID)
}


func (t *Teams) RandomPicker(attackerTeamID int) int {
    toPickFrom, err := t.GetOppsingTeamByIndex(attackerTeamID)
    if err!=nil{
        return 0
    }
    length := len(toPickFrom)
    return toPickFrom[rand.Intn(length)]

}

func InitOrder(teams []characters.BaseCombatant, size int) []characters.BaseCombatant{
    for i:=0; i<size; i++{
        teams[i].RollInitiative()
    }
    utils.QuickSort(teams, 0, size-1)
    return teams
}


func NextTurn(teams []characters.BaseCombatant) func() characters.BaseCombatant{
    turnCount := 0
    return func() characters.BaseCombatant{
        combIndex := turnCount % len(teams)
        turnCount += 1
        fmt.Printf("'\n'---Turn count is: %d---'\n'", turnCount)
        return teams[combIndex]
    }
}

type InitUtil interface{
    AAA() bool
    AAnA() bool
}
