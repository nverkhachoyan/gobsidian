---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/9
- monster/environment/desert
- monster/environment/forest
- monster/environment/swamp
- monster/environment/underdark
- monster/size/large
- monster/type/giant
aliases: ["Rot Troll"]
NoteIcon: monster
BestiaryType: giant
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 247, Mordenkainen's Tome of Foes p. 244
---
# [Rot Troll](3-Mechanics\CLI\bestiary\giant/rot-troll-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 247, Mordenkainen's Tome of Foes p. 244*  

```statblock
"name": "Rot Troll (MPMM)"
"size": "Large"
"type": "giant"
"alignment": "Typically  Chaotic Evil"
"ac": !!int "16"
"ac_class": "natural armor"
"hp": !!int "138"
"hit_dice": "12d10 + 72"
"stats":
- !!int "18"
- !!int "13"
- !!int "22"
- !!int "5"
- !!int "8"
- !!int "4"
"speed": "30 ft."
"skillsaves":
  "Perception": !!int "3"
"damage_immunities": "necrotic"
"senses": "darkvision 60 ft., passive Perception 13"
"languages": "Giant"
"cr": "9"
"traits":
- "desc": "At the end of each of the troll's turns, each creature within 5 feet of\
    \ it takes dice:2d10|text(11) (2d10) necrotic damage, unless the troll has\
    \ taken acid or fire damage since the end of its last turn."
  "name": "Rancid Degeneration"
"actions":
- "desc": "The troll makes one Bite attack and two Claw attacks."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+8 (+8) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d10 + 4|text(9) (1d10 + 4) piercing damage plus dice:3d10|text(16)\
    \ (3d10) necrotic damage."
  "name": "Bite"
- "desc": "Melee Weapon Attack: dice: d20+8 (+8) to hit, reach 5 ft., one target.\
    \ Hit: dice:2d6 + 4|text(11) (2d6 + 4) slashing damage plus dice:2d6|text(7)\
    \ (2d6) necrotic damage."
  "name": "Claws"
"source":
- "MPMM"
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Rot%20Troll.webp"
```
^statblock

```encounter-table
name: Rot Troll
creatures:
 - 1: Rot Troll
```

## Environment

desert, forest, swamp, underdark