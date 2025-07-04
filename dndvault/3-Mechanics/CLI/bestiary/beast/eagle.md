---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mm
- monster/cr/0
- monster/environment/coastal
- monster/environment/grassland
- monster/environment/hill
- monster/environment/mountain
- monster/size/small
- monster/type/beast
aliases: ["Eagle"]
NoteIcon: monster
BestiaryType: beast
SourceType: Bestiary
BookSource: Monster Manual p. 322. Available in the SRD and the Basic Rules.
---
# [Eagle](3-Mechanics\CLI\bestiary\beast/eagle.md)
*Source: Monster Manual p. 322. Available in the SRD and the Basic Rules.*  

```statblock
"name": "Eagle"
"size": "Small"
"type": "beast"
"alignment": "Unaligned"
"ac": !!int "12"
"hp": !!int "3"
"hit_dice": "1d6"
"stats":
- !!int "6"
- !!int "15"
- !!int "10"
- !!int "2"
- !!int "14"
- !!int "7"
"speed": "10 ft., fly 60 ft."
"skillsaves":
  "Perception": !!int "4"
"senses": "passive Perception 14"
"languages": ""
"cr": "0"
"traits":
- "desc": "The eagle has advantage on Wisdom ([Perception](/3-Mechanics/CLI/rules/skills.md#Perception))\
    \ checks that rely on sight."
  "name": "Keen Sight"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+4 (+4) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d4 + 2|text(4) (1d4 + 2) slashing damage."
  "name": "Talons"
"source":
- "MM"
- "CM"
- "ToFW"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MM/Eagle.webp"
```
^statblock

```encounter-table
name: Eagle
creatures:
 - 1: Eagle
```

## Environment

mountain, grassland, hill, coastal