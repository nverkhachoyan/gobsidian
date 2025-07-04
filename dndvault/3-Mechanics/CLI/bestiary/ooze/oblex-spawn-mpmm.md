---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/1-4
- monster/environment/swamp
- monster/environment/underdark
- monster/environment/urban
- monster/size/tiny
- monster/type/ooze
aliases: ["Oblex Spawn"]
NoteIcon: monster
BestiaryType: ooze
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 197, Mordenkainen's Tome of Foes p. 217
---
# [Oblex Spawn](3-Mechanics\CLI\bestiary\ooze/oblex-spawn-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 197, Mordenkainen's Tome of Foes p. 217*  

```statblock
"name": "Oblex Spawn (MPMM)"
"size": "Tiny"
"type": "ooze"
"alignment": "Typically  Lawful Evil"
"ac": !!int "13"
"hp": !!int "18"
"hit_dice": "4d4 + 8"
"stats":
- !!int "8"
- !!int "16"
- !!int "15"
- !!int "14"
- !!int "11"
- !!int "10"
"speed": "20 ft."
"saves":
  "Charisma": !!int "2"
  "Intelligence": !!int "4"
"condition_immunities": "[blinded](/3-Mechanics/CLI/rules/conditions.md#blinded),\
  \ [charmed](/3-Mechanics/CLI/rules/conditions.md#charmed), [deafened](/3-Mechanics/CLI/rules/conditions.md#deafened),\
  \ [exhaustion](/3-Mechanics/CLI/rules/conditions.md#exhaustion), [prone](/3-Mechanics/CLI/rules/conditions.md#prone)"
"senses": "blindsight 60 ft. (blind beyond this radius), passive Perception 12"
"languages": ""
"cr": "1/4"
"traits":
- "desc": "The oblex can move through a space as narrow as 1 inch wide without squeezing."
  "name": "Amorphous"
- "desc": "If the oblex takes fire damage, it has disadvantage on attack rolls and\
    \ ability checks until the end of its next turn."
  "name": "Aversion to Fire"
- "desc": "The oblex doesn't require sleep."
  "name": "Unusual Nature"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+5 (+5) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d4 + 3|text(5) (1d4 + 3) bludgeoning damage plus dice:1d4|text(2)\
    \ (1d4) psychic damage."
  "name": "Pseudopod"
"source":
- "MPMM"
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Oblex%20Spawn.webp"
```
^statblock

```encounter-table
name: Oblex Spawn
creatures:
 - 1: Oblex Spawn
```

## Environment

swamp, underdark, urban