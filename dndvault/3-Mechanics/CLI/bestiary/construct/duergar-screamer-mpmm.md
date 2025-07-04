---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/3
- monster/environment/mountain
- monster/environment/underdark
- monster/size/medium
- monster/type/construct/dwarf
aliases: ["Duergar Screamer"]
NoteIcon: monster
BestiaryType: construct (dwarf)
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 111, Mordenkainen's Tome of Foes p. 190
---
# [Duergar Screamer](3-Mechanics\CLI\bestiary\construct/duergar-screamer-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 111, Mordenkainen's Tome of Foes p. 190*  

```statblock
"name": "Duergar Screamer (MPMM)"
"size": "Medium"
"type": "construct"
"subtype": "dwarf"
"alignment": "Any alignment"
"ac": !!int "15"
"ac_class": "natural armor"
"hp": !!int "38"
"hit_dice": "7d8 + 7"
"stats":
- !!int "18"
- !!int "7"
- !!int "12"
- !!int "5"
- !!int "5"
- !!int "5"
"speed": "20 ft."
"damage_immunities": "poison"
"condition_immunities": "[charmed](/3-Mechanics/CLI/rules/conditions.md#charmed),\
  \ [exhaustion](/3-Mechanics/CLI/rules/conditions.md#exhaustion), [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened),\
  \ [paralyzed](/3-Mechanics/CLI/rules/conditions.md#paralyzed), [petrified](/3-Mechanics/CLI/rules/conditions.md#petrified),\
  \ [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)"
"senses": "darkvision 60 ft., passive Perception 7"
"languages": "understands Dwarvish but can't speak"
"cr": "3"
"actions":
- "desc": "The screamer makes one Drill attack, and it uses Sonic Scream."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+6 (+6) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d12 + 4|text(10) (1d12 + 4) piercing damage."
  "name": "Drill"
- "desc": "The screamer emits destructive energy in a 15-foot cube. Each creature\
    \ in that area must succeed on a DC 11 Strength saving throw or take dice:2d6|text(7)\
    \ (2d6) thunder damage and be knocked [prone](/3-Mechanics/CLI/rules/conditions.md#prone)."
  "name": "Sonic Scream"
"reactions":
- "desc": "Immediately after a creature within 5 feet of the screamer hits it with\
    \ an attack roll, the screamer makes a Drill attack against that creature."
  "name": "Engine of Pain"
"source":
- "MPMM"
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Duergar%20Screamer.webp"
```
^statblock

```encounter-table
name: Duergar Screamer
creatures:
 - 1: Duergar Screamer
```

## Environment

mountain, underdark