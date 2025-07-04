---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mm
- monster/cr/0
- monster/environment/desert
- monster/environment/forest
- monster/environment/grassland
- monster/environment/urban
- monster/size/tiny
- monster/type/beast
aliases: ["Cat"]
NoteIcon: monster
BestiaryType: beast
SourceType: Bestiary
BookSource: Monster Manual p. 320, Tasha's Cauldron of Everything. Available in the SRD and the Basic Rules.
---
# [Cat](3-Mechanics\CLI\bestiary\beast/cat.md)
*Source: Monster Manual p. 320, Tasha's Cauldron of Everything. Available in the SRD and the Basic Rules.*  

```statblock
"name": "Cat"
"size": "Tiny"
"type": "beast"
"alignment": "Unaligned"
"ac": !!int "12"
"hp": !!int "2"
"hit_dice": "1d4"
"stats":
- !!int "3"
- !!int "15"
- !!int "10"
- !!int "3"
- !!int "12"
- !!int "7"
"speed": "40 ft., climb 30 ft."
"skillsaves":
  "Stealth": !!int "4"
  "Perception": !!int "3"
"senses": "passive Perception 13"
"languages": ""
"cr": "0"
"traits":
- "desc": "The cat has advantage on Wisdom ([Perception](/3-Mechanics/CLI/rules/skills.md#Perception))\
    \ checks that rely on smell."
  "name": "Keen Smell"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+0 (+0) to hit, reach 5 ft., one target.\
    \ Hit: 1 slashing damage."
  "name": "Claws"
"source":
- "MM"
- "CoS"
- "HotDQ"
- "SKT"
- "ToA"
- "WDH"
- "IMR"
- "IDRotF"
- "TCE"
- "CM"
- "WBtW"
- "JttRC"
- "DoDk"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MM/Cat.webp"
```
^statblock

```encounter-table
name: Cat
creatures:
 - 1: Cat
```

## Environment

grassland, forest, urban, desert