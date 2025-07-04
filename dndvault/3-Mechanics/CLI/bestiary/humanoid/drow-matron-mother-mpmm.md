---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/20
- monster/environment/underdark
- monster/size/medium
- monster/type/humanoid/cleric
- monster/type/humanoid/elf
aliases: ["Drow Matron Mother"]
NoteIcon: monster
BestiaryType: humanoid (cleric, elf)
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 104, Mordenkainen's Tome of Foes p. 186
---
# [Drow Matron Mother](3-Mechanics\CLI\bestiary\humanoid/drow-matron-mother-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 104, Mordenkainen's Tome of Foes p. 186*  

```statblock
"name": "Drow Matron Mother (MPMM)"
"size": "Medium"
"type": "humanoid"
"subtype": "cleric, elf"
"alignment": "Typically  Neutral Evil"
"ac": !!int "17"
"ac_class": "[half plate](/3-Mechanics/CLI/items/half-plate-armor.md)"
"hp": !!int "247"
"hit_dice": "33d8 + 99"
"stats":
- !!int "12"
- !!int "18"
- !!int "16"
- !!int "17"
- !!int "21"
- !!int "22"
"speed": "30 ft."
"saves":
  "Charisma": !!int "12"
  "Wisdom": !!int "11"
  "Constitution": !!int "9"
"skillsaves":
  "Stealth": !!int "10"
  "Religion": !!int "9"
  "Insight": !!int "11"
  "Perception": !!int "11"
"condition_immunities": "[charmed](/3-Mechanics/CLI/rules/conditions.md#charmed),\
  \ [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened), [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)"
"senses": "darkvision 120 ft., passive Perception 21"
"languages": "Elvish, Undercommon"
"cr": "20"
"traits":
- "desc": "The drow casts one of the following spells, requiring no material components\
    \ and using Charisma as the spellcasting ability (spell save DC 20):\n\nAt will:\
    \ [command](/3-Mechanics/CLI/spells/command.md), [dancing lights](/3-Mechanics/CLI/spells/dancing-lights.md),\
    \ [detect magic](/3-Mechanics/CLI/spells/detect-magic.md), [thaumaturgy](/3-Mechanics/CLI/spells/thaumaturgy.md)\n\
    \n1/day each: [clairvoyance](/3-Mechanics/CLI/spells/clairvoyance.md), [darkness](/3-Mechanics/CLI/spells/darkness.md),\
    \ [detect thoughts](/3-Mechanics/CLI/spells/detect-thoughts.md), [dispel magic](/3-Mechanics/CLI/spells/dispel-magic.md),\
    \ [faerie fire](/3-Mechanics/CLI/spells/faerie-fire.md), [gate](/3-Mechanics/CLI/spells/gate.md),\
    \ [levitate](/3-Mechanics/CLI/spells/levitate.md) (self only), [suggestion](/3-Mechanics/CLI/spells/suggestion.md)\n\
    \n2/day each: [banishment](/3-Mechanics/CLI/spells/banishment.md), [blade\
    \ barrier](/3-Mechanics/CLI/spells/blade-barrier.md), [cure wounds](/3-Mechanics/CLI/spells/cure-wounds.md),\
    \ [hold person](/3-Mechanics/CLI/spells/hold-person.md), [plane shift](/3-Mechanics/CLI/spells/plane-shift.md),\
    \ [silence](/3-Mechanics/CLI/spells/silence.md)"
  "name": "Spellcasting"
- "desc": "The drow has advantage on saving throws against being [charmed](/3-Mechanics/CLI/rules/conditions.md#charmed),\
    \ and magic can't put the drow to sleep."
  "name": "Fey Ancestry"
- "desc": "The drow wields a [tentacle rod](/3-Mechanics/CLI/items/tentacle-rod.md)."
  "name": "Special Equipment"
- "desc": "While in sunlight, the drow has disadvantage on attack rolls, as well as\
    \ on Wisdom ([Perception](/3-Mechanics/CLI/rules/skills.md#Perception)) checks\
    \ that rely on sight."
  "name": "Sunlight Sensitivity"
"actions":
- "desc": "The drow makes two Demon Staff attacks or one Demon Staff attack and three\
    \ Tentacle Rod attacks."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+10 (+10) to hit, reach 5 ft., one\
    \ target. Hit: dice:1d6 + 4|text(7) (1d6 + 4) bludgeoning damage, or dice:1d8\
    \ + 4|text(8) (1d8 + 4) bludgeoning damage if used with two hands, plus dice:4d6|text(14)\
    \ (4d6) psychic damage. The target must succeed on a DC 19 Wisdom saving throw\
    \ or become [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened) of the\
    \ drow for 1 minute. The [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened)\
    \ target can repeat the saving throw at the end of each of its turns, ending the\
    \ effect on itself on a success."
  "name": "Demon Staff"
- "desc": "Melee Weapon Attack: dice: d20+9 (+9) to hit, reach 15 ft., one creature.\
    \ Hit: dice:1d6|text(3) (1d6) bludgeoning damage. If the target is hit three\
    \ times by the [rod](/3-Mechanics/CLI/items/tentacle-rod.md) on one turn, the\
    \ target must succeed on a DC 15 Constitution saving throw or suffer the following\
    \ effects for 1 minute: the target's speed is halved, it has disadvantage on Dexterity\
    \ saving throws, and it can't use reactions. Moreover, on each of its turns, it\
    \ can take either an action or a bonus action, but not both. At the end of each\
    \ of its turns, it can repeat the saving throw, ending the effect on itself on\
    \ a success."
  "name": "Tentacle Rod"
- "desc": "A 10-foot-radius, 40-foot-high column of divine fire sprouts in an area\
    \ up to 120 feet away from the drow. Each creature in the column must make a DC\
    \ 20 Dexterity saving throw, taking dice:4d6|text(14) (4d6) fire damage and\
    \ dice:4d6|text(14) (4d6) radiant damage on a failed save, or half as much\
    \ damage on a successful one."
  "name": "Divine Flame (2/Day)"
"bonus_actions":
- "desc": "The drow bestows the Spider Queen's blessing on one ally she can see within\
    \ 30 feet of her. The ally takes dice:2d6|text(7) (2d6) psychic damage but\
    \ has advantage on the next attack roll it makes before the end of its next turn."
  "name": "Lolth's Fickle Favor"
- "desc": "The drow magically summons a [glabrezu](/3-Mechanics/CLI/bestiary/fiend/glabrezu.md)\
    \ or a [yochlol](/3-Mechanics/CLI/bestiary/fiend/yochlol.md). The summoned creature\
    \ appears in an unoccupied space within 60 feet of its summoner, acts as an ally\
    \ of its summoner, and can't summon other demons. It remains for 10 minutes, until\
    \ it or its summoner dies, or until its summoner dismisses it as an action."
  "name": "Summon Servant (1/Day)"
"legendary_actions":
- "desc": "An allied demon within 30 feet of the drow uses its reaction to make one\
    \ attack against a target of the drow's choice that she can see."
  "name": "Compel Demon"
- "desc": "The drow makes one Demon Staff attack."
  "name": "Demon Staff"
- "desc": "The drow uses Spellcasting."
  "name": "Cast a Spell (Costs 2 Actions)"
"source":
- "MPMM"
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Drow%20Matron%20Mother.webp"
```
^statblock

```encounter-table
name: Drow Matron Mother
creatures:
 - 1: Drow Matron Mother
```

## Environment

underdark