---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mm
- monster/cr/1-4
- monster/environment/forest
- monster/environment/underdark
- monster/size/large
- monster/type/beast
aliases: ["Giant Bat"]
NoteIcon: monster
BestiaryType: beast
SourceType: Bestiary
BookSource: Monster Manual p. 323, Tasha's Cauldron of Everything. Available in the SRD and the Basic Rules.
---
# [Giant Bat](3-Mechanics\CLI\bestiary\beast/giant-bat.md)
*Source: Monster Manual p. 323, Tasha's Cauldron of Everything. Available in the SRD and the Basic Rules.*  

```statblock
"name": "Giant Bat"
"size": "Large"
"type": "beast"
"alignment": "Unaligned"
"ac": !!int "13"
"hp": !!int "22"
"hit_dice": "4d10"
"stats":
- !!int "15"
- !!int "16"
- !!int "11"
- !!int "2"
- !!int "12"
- !!int "6"
"speed": "10 ft., fly 60 ft."
"senses": "blindsight 60 ft., passive Perception 11"
"languages": ""
"cr": "1/4"
"traits":
- "desc": "The bat can't use its blindsight while [deafened](/3-Mechanics/CLI/rules/conditions.md#deafened)."
  "name": "Echolocation"
- "desc": "The bat has advantage on Wisdom ([Perception](/3-Mechanics/CLI/rules/skills.md#Perception))\
    \ checks that rely on hearing."
  "name": "Keen Hearing"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+4 (+4) to hit, reach 5 ft., one creature.\
    \ Hit: dice:1d6 + 2|text(5) (1d6 + 2) piercing damage."
  "name": "Bite"
"source":
- "MM"
- "PotA"
- "WDMM"
- "MOT"
- "TCE"
- "PSX"
- "QftIS"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MM/Giant%20Bat.webp"
```
^statblock

```encounter-table
name: Giant Bat
creatures:
 - 1: Giant Bat
```

## Environment

underdark, forest