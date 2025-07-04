---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mm
- monster/cr/1-2
- monster/environment/forest
- monster/size/medium
- monster/type/beast
aliases: ["Black Bear"]
NoteIcon: monster
BestiaryType: beast
SourceType: Bestiary
BookSource: Monster Manual p. 318. Available in the SRD and the Basic Rules.
---
# [Black Bear](3-Mechanics\CLI\bestiary\beast/black-bear.md)
*Source: Monster Manual p. 318. Available in the SRD and the Basic Rules.*  

```statblock
"name": "Black Bear"
"size": "Medium"
"type": "beast"
"alignment": "Unaligned"
"ac": !!int "11"
"ac_class": "natural armor"
"hp": !!int "19"
"hit_dice": "3d8 + 6"
"stats":
- !!int "15"
- !!int "10"
- !!int "14"
- !!int "2"
- !!int "12"
- !!int "7"
"speed": "40 ft., climb 30 ft."
"skillsaves":
  "Perception": !!int "3"
"senses": "passive Perception 13"
"languages": ""
"cr": "1/2"
"traits":
- "desc": "The bear has advantage on Wisdom ([Perception](/3-Mechanics/CLI/rules/skills.md#Perception))\
    \ checks that rely on smell."
  "name": "Keen Smell"
"actions":
- "desc": "The bear makes two attacks: one with its bite and one with its claws."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+4 (+4) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d6 + 2|text(5) (1d6 + 2) piercing damage."
  "name": "Bite"
- "desc": "Melee Weapon Attack: dice: d20+4 (+4) to hit, reach 5 ft., one target.\
    \ Hit: dice:2d4 + 2|text(7) (2d4 + 2) slashing damage."
  "name": "Claws"
"source":
- "MM"
- "PotA"
- "SKT"
- "WDMM"
- "GoS"
- "IMR"
- "SatO"
- "ToFW"
- "QftIS"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MM/Black%20Bear.webp"
```
^statblock

```encounter-table
name: Black Bear
creatures:
 - 1: Black Bear
```

## Environment

forest