package characters
import (
    "fmt"
    "math/rand"
)


type Coords struct{
    X int
    Y int
}

type StatArray struct{
    Str int
    Dex int
    Con int
    Int int
    Wis int
    Cha int
}


type Weapon struct{
    Name string
    //Property string
    //Mastery
    Damage int
    Range int
    //DamageType string
}



type BaseCombatant struct {
    ID int
    Name string
    Statarray StatArray
    HP int
    AC int
    Proficiency int
    PrimaryStat int
    NoAttacks int
    Attack_roll int
    Weapon Weapon
    Damage int
    Initiative int
    Team int
    Coords Coords
}

type Caster struct {
    BaseCombatant
    Spells SpellBook
}

func (m *BaseCombatant) Attack(target BaseCombatant) (int, BaseCombatant){
    attackRoll := m.GetAttackRoll()
    if attackRoll >= target.AC {
        fmt.Printf("%s of Team %d attacks %s of Team %d and coords %v for %d damage\n", m.Name, m.Team, target.Name, target.Team, target.Coords, m.Damage)
        fmt.Printf("Target HP before attack: %d\n", target.HP)

        target = target.TakeDamage(m.Damage)

        fmt.Printf("Target HP after attack: %d\n", target.HP)
    } else {
        fmt.Printf("%s misses %s\n", m.Name, target.Name)
    }
    return target.ID, target
}



func (c *Caster) GetSaveDC(int) (int, error){
    var abilityscore int
    switch c.PrimaryStat {
    case STR:
        abilityscore = c.Statarray.Str
    case DEX:
        abilityscore = c.Statarray.Dex
    case CON:
        abilityscore = c.Statarray.Con
    case INT:
        abilityscore = c.Statarray.Int
    case WIS:
        abilityscore = c.Statarray.Wis
    case CHA:
        abilityscore = c.Statarray.Cha
    default:
        return 0, fmt.Errorf("Invalid PrimaryStat")
    }
    return 8 + c.Proficiency + (abilityscore-10)/2, nil
}

func (i *BaseCombatant) SetID(id int){
    i.ID = id
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

func (dmg *BaseCombatant) GetDMG() int{
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


func (m BaseCombatant) TakeDamage(damage int) BaseCombatant {
    damage = rand.Intn(damage)
    fmt.Printf("HP of %s before damage is %d\n", m.Name, m.HP)
    m.HP -= int(damage)
    if m.HP < 0 {
        m.HP = 0
    }
    fmt.Printf("%s takes %d damage. HP is now %d\n", m.Name, damage, m.HP)
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
    GetSaveDC() (int, error)
    GetName() string
    GetTeam() int
    RollInitiative() int
}

type Mage interface{
    GetSaveDC() (int, error)
}
