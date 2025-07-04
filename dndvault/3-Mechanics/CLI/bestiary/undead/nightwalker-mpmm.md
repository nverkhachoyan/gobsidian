---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/20
- monster/environment/arctic
- monster/environment/desert
- monster/environment/swamp
- monster/environment/underdark
- monster/size/huge
- monster/type/undead
aliases: ["Nightwalker"]
NoteIcon: monster
BestiaryType: undead
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 194, Mordenkainen's Tome of Foes p. 216
---
# [Nightwalker](3-Mechanics\CLI\bestiary\undead/nightwalker-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 194, Mordenkainen's Tome of Foes p. 216*  

```statblock
"name": "Nightwalker (MPMM)"
"size": "Huge"
"type": "undead"
"alignment": "Typically  Chaotic Evil"
"ac": !!int "14"
"hp": !!int "337"
"hit_dice": "25d12 + 175"
"stats":
- !!int "22"
- !!int "19"
- !!int "24"
- !!int "6"
- !!int "9"
- !!int "8"
"speed": "40 ft., fly 40 ft."
"saves":
  "Constitution": !!int "13"
"damage_resistances": "acid; cold; fire; lightning; thunder; bludgeoning, piercing,\
  \ slashing from nonmagical attacks"
"damage_immunities": "necrotic, poison"
"condition_immunities": "[exhaustion](/3-Mechanics/CLI/rules/conditions.md#exhaustion),\
  \ [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened), [grappled](/3-Mechanics/CLI/rules/conditions.md#grappled),\
  \ [paralyzed](/3-Mechanics/CLI/rules/conditions.md#paralyzed), [petrified](/3-Mechanics/CLI/rules/conditions.md#petrified),\
  \ [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned), [prone](/3-Mechanics/CLI/rules/conditions.md#prone),\
  \ [restrained](/3-Mechanics/CLI/rules/conditions.md#restrained)"
"senses": "darkvision 120 ft., passive Perception 9"
"languages": ""
"cr": "20"
"traits":
- "desc": "Any creature that starts its turn within 30 feet of the nightwalker must\
    \ succeed on a DC 21 Constitution saving throw or take dice:6d6|text(21) (6d6)\
    \ necrotic damage. Undead are immune to this aura."
  "name": "Annihilating Aura"
- "desc": "A creature dies if reduced to 0 hit points by the nightwalker and can't\
    \ be revived except by a [wish](/3-Mechanics/CLI/spells/wish.md) spell."
  "name": "Life Eater"
- "desc": "The nightwalker doesn't require air, food, drink, or sleep."
  "name": "Unusual Nature"
"actions":
- "desc": "The nightwalker makes two Enervating Focus attacks, one of which can be\
    \ replaced by Finger of Doom, if available."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+12 (+12) to hit, reach 15 ft., one\
    \ target. Hit: dice:5d8 + 6|text(28) (5d8 + 6) necrotic damage. The target\
    \ must succeed on a DC 21 Constitution saving throw or its hit point maximum is\
    \ reduced by an amount equal to the necrotic damage taken. This reduction lasts\
    \ until the target finishes a long rest. The target dies if its hit point maximum\
    \ is reduced to 0."
  "name": "Enervating Focus"
- "desc": "The nightwalker points at one creature it can see within 300 feet of it.\
    \ The target must succeed on a DC 21 Wisdom saving throw or take dice:6d12|text(39)\
    \ (6d12) necrotic damage and become [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened)\
    \ until the end of the nightwalker's next turn. While [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened)\
    \ in this way, the creature is also [paralyzed](/3-Mechanics/CLI/rules/conditions.md#paralyzed).\
    \ If a target's saving throw is successful, the target is immune to the nightwalker's\
    \ Finger of Doom for the next 24 hours."
  "name": "Finger of Doom (Recharge 6)"
"source":
- "MPMM"
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Nightwalker.webp"
```
^statblock

```encounter-table
name: Nightwalker
creatures:
 - 1: Nightwalker
```

## Environment

arctic, desert, swamp, underdark