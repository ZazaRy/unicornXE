package combat

import (
    "fmt"
    "sort"
    "math/rand"
   "github.com/ZazaRy/unicornXE/pkg/characters"
)


type Teams struct {
    Combatants []characters.BaseCombatant
}


func (t *Teams) AddToTracker(comb characters.BaseCombatant) {
   t.Combatants = append(t.Combatants, comb)

}

func RollInitiative(t Teams) func() string {
    sort.Slice(t.Combatants, func(i,j int) bool {
        return t.Combatants[i].GetInitiative() > t.Combatants[j].GetInitiative()
    })

    turnIndex := 0

    return func() string{
        currentCombatant := t.Combatants[turnIndex]
        turnIndex = (turnIndex + 1) % len(t.Combatants)
        return currentCombatant.GetName()
    }}


