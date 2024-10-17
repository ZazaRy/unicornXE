package characters


import (
    "fmt"
    _"math/rand"

)

const(
    Acid = 1 << iota
    Bludgeoning 
    Cold
    Fire
    Force
    Lightning
    Necrotic
    Piercing
    Poison
    Psychic
    Radiant
    Slashing
    Thunder
)

const(
    STR = 1 << iota
    DEX
    CON
    INT
    WIS
    CHA
)

type SpellAttack struct {
    AttackRoll int
    Range int
}

type SpellSave struct {
    SaveDC int
    SaveType int
}

type Damage struct {
    DamageRoll int
    DiceCount int
    DamageType int
}

type Spell struct {
    Damage Damage
    Attack *SpellAttack
    SpellSave func(combatant Caster) *SpellSave

}

type SpellBook struct {
    Spells map[string]Spell
}

var SpellList = make(map[string]map[string]Spell)

func (s *Spell) InitSpellList(pStat int) {
    SpellList["Fire Bolt"]["Attack"] = Spell{Damage: Damage{10, 1, Fire}, Attack: &SpellAttack{AttackRoll: 5, Range: 120}}
    SpellList["Toll the Dead"]["Save"] = Spell{Damage: Damage{8, 1, Necrotic}, SpellSave: func(c Caster) *SpellSave {
        dc, err := c.GetSaveDC(pStat)
        if err != nil {
            fmt.Println("Error calculating save DC", err)
            return nil
        }
        return &SpellSave{
            SaveDC: dc,
            SaveType: c.PrimaryStat,}
    }}

}


type Arcane interface {
    InitSpellList() int
}



