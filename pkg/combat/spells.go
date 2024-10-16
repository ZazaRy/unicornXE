package combat


import (
    "fmt"
    "math/rand"
    "github.com/ZazaRy/unicornXE/pkg/characters"
    _"github.com/ZazaRy/unicornXE/pkg/utils"

)

type SpellAttack struct {
    AttackRoll int
    Range int
}

type SpellSave struct {
    SaveDC int
    SaveType string
}

type Damage struct {
    DamageRoll int
    DamageType string
}

type Spell struct {
    Name string
    Damage Damage
    Attack SpellAttack
    SpellSave SpellSave

}


func (s *Spell) CastSpell(attacker, defender characters.Caster) int{
    attackRoll := rand.Intn(20) + attacker.Attack_roll
    if attackRoll >= s.Attack.AttackRoll{
        return s.Damage.DamageRoll
    }
    return 0
}
