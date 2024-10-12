package characters
import (
    "fmt"
    "math/rand"
)
type Coords struct{
    x int
    y int
}
type BaseCombatant struct {
    Name string
    HP int
    AC int
    Attack_roll int
    Damage float32
    Initiative int
    Team int
    Coords Coords
}



func (m *BaseCombatant) Attack(target *BaseCombatant) {
    attackRoll := m.GetAttackRoll()
    if attackRoll >= target.AC {
        fmt.Printf("%s of Team %d attacks %s of Team %d for %f damage\n", m.Name, m.Team, target.Name, target.Team, m.Damage)
        target.TakeDamage(m.Damage)
    } else {
        fmt.Printf("%s misses %s\n", m.Name, target.Name)
    }
}


func (p *BaseCombatant) GetPos() Coords{
    return p.Coords
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

func (t *BaseCombatant) GetTeam() int{
    return t.Team
}

func (t *BaseCombatant) RollInitiative() int{
    roll := t.Initiative + rand.Intn(20)
    t.Initiative = roll
    return roll
}

func (m *BaseCombatant) TakeDamage(damage float32) {
    m.HP -= int(damage)
    if m.HP < 0{
        m.HP = 0
    }
    fmt.Printf("%s takes %f damage. HP is now %d\n", m.Name, damage, m.HP)
}

type Combatant interface{
    Attack(target Combatant)
    TakeDamage(damage float32)
    GetHP() int
    GetAC() int
    GetAttackRoll() int
    GetInitiative() int
    GetName() string
    GetTeam() int
    RollInitiative() int
}
