---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mm
- monster/cr/1-8
- monster/environment/forest
- monster/environment/grassland
- monster/environment/hill
- monster/size/medium
- monster/type/beast
aliases: ["Giant Weasel"]
NoteIcon: monster
BestiaryType: beast
SourceType: Bestiary
BookSource: Monster Manual p. 329. Available in the SRD and the Basic Rules.
---
# [Giant Weasel](3-Mechanics\CLI\bestiary\beast/giant-weasel.md)
*Source: Monster Manual p. 329. Available in the SRD and the Basic Rules.*  

```statblock
"name": "Giant Weasel"
"size": "Medium"
"type": "beast"
"alignment": "Unaligned"
"ac": !!int "13"
"hp": !!int "9"
"hit_dice": "2d8"
"stats":
- !!int "11"
- !!int "16"
- !!int "10"
- !!int "4"
- !!int "12"
- !!int "5"
"speed": "40 ft."
"skillsaves":
  "Stealth": !!int "5"
  "Perception": !!int "3"
"senses": "darkvision 60 ft., passive Perception 13"
"languages": ""
"cr": "1/8"
"traits":
- "desc": "The weasel has advantage on Wisdom ([Perception](/3-Mechanics/CLI/rules/skills.md#Perception))\
    \ checks that rely on hearing or smell."
  "name": "Keen Hearing and Smell"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+5 (+5) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d4 + 3|text(5) (1d4 + 3) piercing damage."
  "name": "Bite"
"source":
- "MM"
- "TftYP"
- "GoS"
- "QftIS"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MM/Giant%20Weasel.webp"
```
^statblock

```encounter-table
name: Giant Weasel
creatures:
 - 1: Giant Weasel
```

## Environment

grassland, forest, hill