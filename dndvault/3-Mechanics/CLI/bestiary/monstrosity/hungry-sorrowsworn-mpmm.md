---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/11
- monster/environment/forest
- monster/environment/underdark
- monster/environment/urban
- monster/size/medium
- monster/type/monstrosity
aliases: ["Hungry Sorrowsworn"]
NoteIcon: monster
BestiaryType: monstrosity
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 223, Mordenkainen's Tome of Foes p. 232
---
# [Hungry Sorrowsworn](3-Mechanics\CLI\bestiary\monstrosity/hungry-sorrowsworn-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 223, Mordenkainen's Tome of Foes p. 232*  

```statblock
"name": "Hungry Sorrowsworn (MPMM)"
"size": "Medium"
"type": "monstrosity"
"alignment": "Typically  Neutral Evil"
"ac": !!int "17"
"ac_class": "natural armor"
"hp": !!int "225"
"hit_dice": "30d8 + 90"
"stats":
- !!int "19"
- !!int "10"
- !!int "17"
- !!int "6"
- !!int "11"
- !!int "6"
"speed": "30 ft."
"damage_resistances": "bludgeoning, piercing, slashing while in dim light or darkness"
"senses": "darkvision 60 ft., passive Perception 10"
"languages": "Common"
"cr": "11"
"traits":
- "desc": "If a creature within 60 feet of the sorrowsworn regains hit points, the\
    \ sorrowsworn gains two benefits until the end of its next turn: it has advantage\
    \ on attack rolls, and its Bite deals an extra dice:4d10|text(22) (4d10) necrotic\
    \ damage on a hit."
  "name": "Life Hunger"
"actions":
- "desc": "The sorrowsworn makes one Bite attack and one Claw attack."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+8 (+8) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d8 + 4|text(8) (1d8 + 4) piercing damage plus dice:3d8|text(13)\
    \ (3d8) necrotic damage."
  "name": "Bite"
- "desc": "Melee Weapon Attack: dice: d20+8 (+8) to hit, reach 10 ft., one target.\
    \ Hit: dice:4d6 + 4|text(18) (4d6 + 4) slashing damage. If the target is\
    \ Medium or smaller, it is [grappled](/3-Mechanics/CLI/rules/conditions.md#grappled)\
    \ (escape DC 16), and it is [restrained](/3-Mechanics/CLI/rules/conditions.md#restrained)\
    \ until the grapple ends. While grappling a creature, the sorrowsworn can't make\
    \ a Claw attack."
  "name": "Claw"
"source":
- "MPMM"
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Hungry%20Sorrowsworn.webp"
```
^statblock

```encounter-table
name: Hungry Sorrowsworn
creatures:
 - 1: Hungry Sorrowsworn
```

## Environment

forest, underdark, urban