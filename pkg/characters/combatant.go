package characters
import (
    "fmt"
    "math/rand"
)

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
