---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/vgm
- monster/cr/1-4
- monster/environment/grassland
- monster/environment/swamp
- monster/size/large
- monster/type/beast
aliases: ["Hadrosaurus"]
NoteIcon: monster
BestiaryType: beast
SourceType: Bestiary
BookSource: Volo's Guide to Monsters p. 140
---
# [Hadrosaurus](3-Mechanics\CLI\bestiary\beast/hadrosaurus-vgm.md)
*Source: Volo's Guide to Monsters p. 140*  

A hadrosaurus is a semi-quadrupedal herbivore recognizable by its bony head crests. If raised as a hatchling, it can be trained to carry a Small or Medium rider.

```statblock
"name": "Hadrosaurus (VGM)"
"size": "Large"
"type": "beast"
"alignment": "Unaligned"
"ac": !!int "11"
"ac_class": "natural armor"
"hp": !!int "19"
"hit_dice": "3d10 + 3"
"stats":
- !!int "15"
- !!int "10"
- !!int "13"
- !!int "2"
- !!int "10"
- !!int "5"
"speed": "40 ft."
"skillsaves":
  "Perception": !!int "2"
"senses": "passive Perception 12"
"languages": ""
"cr": "1/4"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+4 (+4) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d10 + 2|text(7) (1d10 + 2) bludgeoning damage."
  "name": "Tail"
"source":
- "VGM"
- "ToA"
- "PSX"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/VGM/Hadrosaurus.webp"
```
^statblock

```encounter-table
name: Hadrosaurus
creatures:
 - 1: Hadrosaurus
```

## Environment

grassland, swamp