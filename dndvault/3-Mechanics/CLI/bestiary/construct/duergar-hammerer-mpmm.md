---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/2
- monster/environment/mountain
- monster/environment/underdark
- monster/size/medium
- monster/type/construct/dwarf
aliases: ["Duergar Hammerer"]
NoteIcon: monster
BestiaryType: construct (dwarf)
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 112, Mordenkainen's Tome of Foes p. 188
---
# [Duergar Hammerer](3-Mechanics\CLI\bestiary\construct/duergar-hammerer-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 112, Mordenkainen's Tome of Foes p. 188*  

```statblock
"name": "Duergar Hammerer (MPMM)"
"size": "Medium"
"type": "construct"
"subtype": "dwarf"
"alignment": "Any alignment"
"ac": !!int "17"
"ac_class": "natural armor"
"hp": !!int "33"
"hit_dice": "6d8 + 6"
"stats":
- !!int "17"
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
"cr": "2"
"traits":
- "desc": "The hammerer deals double damage to objects and structures."
  "name": "Siege Monster"
"actions":
- "desc": "The hammerer makes one Claw attack and one Hammer attack."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+5 (+5) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d6 + 3|text(6) (1d6 + 3) piercing damage."
  "name": "Claw"
- "desc": "Melee Weapon Attack: dice: d20+5 (+5) to hit, reach 5 ft., one target.\
    \ Hit: dice:2d6 + 3|text(10) (2d6 + 3) bludgeoning damage."
  "name": "Hammer"
"reactions":
- "desc": "Immediately after a creature within 5 feet of the hammerer hits it with\
    \ an attack roll, the hammerer makes a Hammer attack against that creature."
  "name": "Engine of Pain"
"source":
- "MPMM"
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Duergar%20Hammerer.webp"
```
^statblock

```encounter-table
name: Duergar Hammerer
creatures:
 - 1: Duergar Hammerer
```

## Environment

mountain, underdark