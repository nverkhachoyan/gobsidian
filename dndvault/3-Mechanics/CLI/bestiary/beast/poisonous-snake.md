---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mm
- monster/cr/1-8
- monster/environment/coastal
- monster/environment/desert
- monster/environment/forest
- monster/environment/grassland
- monster/environment/hill
- monster/environment/swamp
- monster/size/tiny
- monster/type/beast
aliases: ["Poisonous Snake"]
NoteIcon: monster
BestiaryType: beast
SourceType: Bestiary
BookSource: Monster Manual p. 334. Available in the SRD and the Basic Rules.
---
# [Poisonous Snake](3-Mechanics\CLI\bestiary\beast/poisonous-snake.md)
*Source: Monster Manual p. 334. Available in the SRD and the Basic Rules.*  

```statblock
"name": "Poisonous Snake"
"size": "Tiny"
"type": "beast"
"alignment": "Unaligned"
"ac": !!int "13"
"hp": !!int "2"
"hit_dice": "1d4"
"stats":
- !!int "2"
- !!int "16"
- !!int "11"
- !!int "1"
- !!int "10"
- !!int "3"
"speed": "30 ft., swim 30 ft."
"senses": "blindsight 10 ft., passive Perception 10"
"languages": ""
"cr": "1/8"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+5 (+5) to hit, reach 5 ft., one target.\
    \ Hit: 1 piercing damage, and the target must make a DC 10 Constitution saving\
    \ throw, taking dice:2d4|text(5) (2d4) poison damage on a failed save, or\
    \ half as much damage on a successful one."
  "name": "Bite"
"source":
- "MM"
- "TftYP"
- "ToA"
- "GoS"
- "IMR"
- "QftIS"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MM/Poisonous%20Snake.webp"
```
^statblock

```encounter-table
name: Poisonous Snake
creatures:
 - 1: Poisonous Snake
```

## Environment

grassland, forest, swamp, hill, desert, coastal