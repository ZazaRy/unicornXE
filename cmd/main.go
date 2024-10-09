package main


import (
    "fmt"
    "sort"
    "math/rand"
)



type Teams struct {
    Combatants []BaseCombatant
}

type BaseCombatant struct {
    Name string
    HP int
    AC int
    Attack_roll int
    Damage float32
    Initiative int
    Team string
}

func (m *BaseCombatant) Attack(target *BaseCombatant) {
    attackRoll := m.GetAttackRoll()
    if attackRoll >= target.AC {
        fmt.Printf("%s attacks %s for %f damage\n", m.Name, target.Name, m.Damage)
        target.TakeDamage(m.Damage)
    } else {
        fmt.Printf("%s misses %s\n", m.Name, target.Name)
    }
}


func (n *BaseCombatant) GetName() string{
    return n.Name
}
func (i *BaseCombatant) GetInitiative() int{
    return i.Initiative
}

func (hp *BaseCombatant) GetHP() int{
    return hp.HP
}

func (ac *BaseCombatant) GetAC() int{
    return ac.AC
}

func (atkRl *BaseCombatant) GetAttackRoll() int{
    return rand.Intn(20) + atkRl.Attack_roll
}

func (dmg *BaseCombatant) GetDMG() float32{
    return dmg.Damage
}

func (t *BaseCombatant) GetTeam() string{
    return t.Team
}

func (m *BaseCombatant) TakeDamage(damage float32) {
    m.HP -= int(damage)
    fmt.Printf("%s takes %f damage\n", m.Name, damage)
}

type Combatant interface{
    Attack(target Combatant)
    TakeDamage(damage float32)
    GetHP() int
    GetAC() int
    GetAttack_roll() int
    GetInitiative() int
    GetName() string
    GetTeam() string
}
func (t *Teams) AddToTracker(comb BaseCombatant) {
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

func main() {
    player1 := BaseCombatant{
        Name: "Player 1",
        HP: 10,
        AC: 10,
        Attack_roll: 10,
        Damage: 4.5,
        Initiative: 1,
    }
    player2 := BaseCombatant{
        Name: "Player 2",
        HP: 10,
        AC: 10,
        Attack_roll: 10,
        Damage: 4.5,
        Initiative: 2,
    }


    team := Teams{}
    team.AddToTracker(player1)
    team.AddToTracker(player2)

    nextTurn := RollInitiative(team)


    fmt.Println("Turn 1: ", nextTurn())

    player1.Attack(&player2)
    player2.Attack(&player1)
}
