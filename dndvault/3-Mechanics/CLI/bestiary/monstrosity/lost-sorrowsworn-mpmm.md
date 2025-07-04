---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/7
- monster/environment/arctic
- monster/environment/desert
- monster/environment/forest
- monster/environment/mountain
- monster/environment/swamp
- monster/environment/underdark
- monster/environment/urban
- monster/size/medium
- monster/type/monstrosity
aliases: ["Lost Sorrowsworn"]
NoteIcon: monster
BestiaryType: monstrosity
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 224, Mordenkainen's Tome of Foes p. 233
---
# [Lost Sorrowsworn](3-Mechanics\CLI\bestiary\monstrosity/lost-sorrowsworn-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 224, Mordenkainen's Tome of Foes p. 233*  

```statblock
"name": "Lost Sorrowsworn (MPMM)"
"size": "Medium"
"type": "monstrosity"
"alignment": "Typically  Neutral Evil"
"ac": !!int "15"
"ac_class": "natural armor"
"hp": !!int "78"
"hit_dice": "12d8 + 24"
"stats":
- !!int "17"
- !!int "12"
- !!int "15"
- !!int "6"
- !!int "7"
- !!int "5"
"speed": "30 ft."
"skillsaves":
  "Athletics": !!int "6"
"damage_resistances": "bludgeoning, piercing, slashing while in dim light or darkness"
"senses": "darkvision 60 ft., passive Perception 8"
"languages": "Common"
"cr": "7"
"actions":
- "desc": "The sorrowsworn makes two Arm Spike attacks."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+6 (+6) to hit, reach 10 ft., one target.\
    \ Hit: dice:2d10 + 3|text(14) (2d10 + 3) piercing damage."
  "name": "Arm Spike"
- "desc": "Melee Weapon Attack: dice: d20+6 (+6) to hit, reach 5 ft., one target.\
    \ Hit: dice:4d10 + 3|text(25) (4d10 + 3) piercing damage, and the target\
    \ is [grappled](/3-Mechanics/CLI/rules/conditions.md#grappled) (escape DC 14)\
    \ if it is a Medium or smaller creature. Until the grapple ends, the target is\
    \ [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened), and it takes\
    \ dice:6d8|text(27) (6d8) psychic damage at the end of each of its turns.\
    \ The sorrowsworn can grapple only one creature at a time."
  "name": "Embrace (Recharge 4-6)"
"reactions":
- "desc": "If the sorrowsworn takes damage, the creature [grappled](/3-Mechanics/CLI/rules/conditions.md#grappled)\
    \ by Embrace takes dice:4d8|text(18) (4d8) psychic damage."
  "name": "Tightening Embrace"
"source":
- "MPMM"
- "MTF"
- "AATM"
- "VEoR"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Lost%20Sorrowsworn.webp"
```
^statblock

```encounter-table
name: Lost Sorrowsworn
creatures:
 - 1: Lost Sorrowsworn
```

## Environment

arctic, desert, forest, mountain, swamp, underdark, urban