---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mm
- monster/cr/0
- monster/environment/desert
- monster/environment/grassland
- monster/size/small
- monster/type/beast
aliases: ["Jackal"]
NoteIcon: monster
BestiaryType: beast
SourceType: Bestiary
BookSource: Monster Manual p. 331. Available in the SRD and the Basic Rules.
---
# [Jackal](3-Mechanics\CLI\bestiary\beast/jackal.md)
*Source: Monster Manual p. 331. Available in the SRD and the Basic Rules.*  

```statblock
"name": "Jackal"
"size": "Small"
"type": "beast"
"alignment": "Unaligned"
"ac": !!int "12"
"hp": !!int "3"
"hit_dice": "1d6"
"stats":
- !!int "8"
- !!int "15"
- !!int "11"
- !!int "3"
- !!int "12"
- !!int "6"
"speed": "40 ft."
"skillsaves":
  "Perception": !!int "3"
"senses": "passive Perception 13"
"languages": ""
"cr": "0"
"traits":
- "desc": "The jackal has advantage on Wisdom ([Perception](/3-Mechanics/CLI/rules/skills.md#Perception))\
    \ checks that rely on hearing or smell."
  "name": "Keen Hearing and Smell"
- "desc": "The jackal has advantage on an attack roll against a creature if at least\
    \ one of the jackal's allies is within 5 feet of the creature and the ally isn't\
    \ [incapacitated](/3-Mechanics/CLI/rules/conditions.md#incapacitated)."
  "name": "Pack Tactics"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+1 (+1) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d4 - 1|text(1) (1d4 - 1) piercing damage."
  "name": "Bite"
"source":
- "MM"
- "GoS"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MM/Jackal.webp"
```
^statblock

```encounter-table
name: Jackal
creatures:
 - 1: Jackal
```

## Environment

grassland, desert