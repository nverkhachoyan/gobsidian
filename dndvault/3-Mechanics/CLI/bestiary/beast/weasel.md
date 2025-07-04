---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mm
- monster/cr/0
- monster/size/tiny
- monster/type/beast
aliases: ["Weasel"]
NoteIcon: monster
BestiaryType: beast
SourceType: Bestiary
BookSource: Monster Manual p. 340. Available in the SRD and the Basic Rules.
---
# [Weasel](3-Mechanics\CLI\bestiary\beast/weasel.md)
*Source: Monster Manual p. 340. Available in the SRD and the Basic Rules.*  

```statblock
"name": "Weasel"
"size": "Tiny"
"type": "beast"
"alignment": "Unaligned"
"ac": !!int "13"
"hp": !!int "1"
"hit_dice": "1d4 - 1"
"stats":
- !!int "3"
- !!int "16"
- !!int "8"
- !!int "2"
- !!int "12"
- !!int "3"
"speed": "30 ft."
"skillsaves":
  "Stealth": !!int "5"
  "Perception": !!int "3"
"senses": "passive Perception 13"
"languages": ""
"cr": "0"
"traits":
- "desc": "The weasel has advantage on Wisdom ([Perception](/3-Mechanics/CLI/rules/skills.md#Perception))\
    \ checks that rely on hearing or smell."
  "name": "Keen Hearing and Smell"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+5 (+5) to hit, reach 5 ft., one creature.\
    \ Hit: 1 piercing damage."
  "name": "Bite"
"source":
- "MM"
- "IDRotF"
- "WBtW"
- "PaBTSO"
- "QftIS"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MM/Weasel.webp"
```
^statblock

```encounter-table
name: Weasel
creatures:
 - 1: Weasel
```