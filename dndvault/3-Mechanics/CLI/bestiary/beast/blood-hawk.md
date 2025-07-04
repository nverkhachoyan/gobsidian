---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mm
- monster/cr/1-8
- monster/environment/arctic
- monster/environment/coastal
- monster/environment/forest
- monster/environment/grassland
- monster/environment/hill
- monster/environment/mountain
- monster/size/small
- monster/type/beast
aliases: ["Blood Hawk"]
NoteIcon: monster
BestiaryType: beast
SourceType: Bestiary
BookSource: Monster Manual p. 319. Available in the SRD and the Basic Rules.
---
# [Blood Hawk](3-Mechanics\CLI\bestiary\beast/blood-hawk.md)
*Source: Monster Manual p. 319. Available in the SRD and the Basic Rules.*  

Taking its name from its crimson feathers and aggressive nature, the blood hawk fearlessly attacks almost any animal, stabbing it with its daggerlike beak. Blood hawks flock together in large numbers, attacking as a pack to take down prey.

```statblock
"name": "Blood Hawk"
"size": "Small"
"type": "beast"
"alignment": "Unaligned"
"ac": !!int "12"
"hp": !!int "7"
"hit_dice": "2d6"
"stats":
- !!int "6"
- !!int "14"
- !!int "10"
- !!int "3"
- !!int "14"
- !!int "5"
"speed": "10 ft., fly 60 ft."
"skillsaves":
  "Perception": !!int "4"
"senses": "passive Perception 14"
"languages": ""
"cr": "1/8"
"traits":
- "desc": "The hawk has advantage on Wisdom ([Perception](/3-Mechanics/CLI/rules/skills.md#Perception))\
    \ checks that rely on sight."
  "name": "Keen Sight"
- "desc": "The hawk has advantage on an attack roll against a creature if at least\
    \ one of the hawk's allies is within 5 feet of the creature and the ally isn't\
    \ [incapacitated](/3-Mechanics/CLI/rules/conditions.md#incapacitated)."
  "name": "Pack Tactics"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+4 (+4) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d4 + 2|text(4) (1d4 + 2) piercing damage."
  "name": "Beak"
"source":
- "MM"
- "PotA"
- "SKT"
- "GoS"
- "EGW"
- "CM"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MM/Blood%20Hawk.webp"
```
^statblock

```encounter-table
name: Blood Hawk
creatures:
 - 1: Blood Hawk
```

## Environment

mountain, grassland, forest, hill, coastal, arctic