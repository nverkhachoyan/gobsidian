---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mm
- monster/cr/0
- monster/environment/desert
- monster/environment/grassland
- monster/environment/hill
- monster/size/medium
- monster/type/beast
aliases: ["Vulture"]
NoteIcon: monster
BestiaryType: beast
SourceType: Bestiary
BookSource: Monster Manual p. 339. Available in the SRD and the Basic Rules.
---
# [Vulture](3-Mechanics\CLI\bestiary\beast/vulture.md)
*Source: Monster Manual p. 339. Available in the SRD and the Basic Rules.*  

```statblock
"name": "Vulture"
"size": "Medium"
"type": "beast"
"alignment": "Unaligned"
"ac": !!int "10"
"hp": !!int "5"
"hit_dice": "1d8 + 1"
"stats":
- !!int "7"
- !!int "10"
- !!int "13"
- !!int "2"
- !!int "12"
- !!int "4"
"speed": "10 ft., fly 50 ft."
"skillsaves":
  "Perception": !!int "3"
"senses": "passive Perception 13"
"languages": ""
"cr": "0"
"traits":
- "desc": "The vulture has advantage on Wisdom ([Perception](/3-Mechanics/CLI/rules/skills.md#Perception))\
    \ checks that rely on sight or smell."
  "name": "Keen Sight and Smell"
- "desc": "The vulture has advantage on an attack roll against a creature if at least\
    \ one of the vulture's allies is within 5 feet of the creature and the ally isn't\
    \ [incapacitated](/3-Mechanics/CLI/rules/conditions.md#incapacitated)."
  "name": "Pack Tactics"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+2 (+2) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d4|text(2) (1d4) piercing damage."
  "name": "Beak"
"source":
- "MM"
- "RoT"
- "ToA"
- "BGDIA"
- "CM"
- "WBtW"
- "BMT"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MM/Vulture.webp"
```
^statblock

```encounter-table
name: Vulture
creatures:
 - 1: Vulture
```

## Environment

grassland, hill, desert