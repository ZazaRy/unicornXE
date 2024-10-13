package characters
import (
    "fmt"
    "math/rand"
)

type Coords struct{
    X int
    Y int
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



func (m BaseCombatant) Attack(target BaseCombatant) BaseCombatant {
    attackRoll := m.GetAttackRoll()
    if attackRoll >= target.AC {
        fmt.Printf("%s of Team %d attacks %s of Team %d and coords %v for %f damage\n", m.Name, m.Team, target.Name, target.Team, target.Coords, m.Damage)
        fmt.Printf("Target HP before attack: %d\n", target.HP)

        target = target.TakeDamage(m.Damage)

        fmt.Printf("Target HP after attack: %d\n", target.HP)
    } else {
        fmt.Printf("%s misses %s\n", m.Name, target.Name)
    }
    return target
}




func (p *BaseCombatant) GetPos() []int{
    return []int{p.Coords.X, p.Coords.Y}
}

func (p *BaseCombatant) SetPos(x, y int){
    p.Coords.X = x
    p.Coords.Y = y
}

func (p* BaseCombatant) MoveLeft(x int){
    p.Coords.X -= x
}

func (p* BaseCombatant) MoveRight(x int){
    p.Coords.X += x
}

func (p* BaseCombatant) MoveUp(y int){
    p.Coords.Y += y
}

func (p* BaseCombatant) MoveDown(y int){
    p.Coords.Y -= y
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

func (m BaseCombatant) TakeDamage(damage float32) BaseCombatant {
    fmt.Printf("HP of %s before damage is %d\n", m.Name, m.HP)
    m.HP -= int(damage)
    if m.HP < 0 {
        m.HP = 0
    }
    fmt.Printf("%s takes %f damage. HP is now %d\n", m.Name, damage, m.HP)
    return m
}


type Movement interface{
    GetPos() Coords
    SetPos(x, y int)
    MoveLeft(x int)
    MoveRight(x int)
    MoveUp(y int)
    MoveDown(y int)
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
