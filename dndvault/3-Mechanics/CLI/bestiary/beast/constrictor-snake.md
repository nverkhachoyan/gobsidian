---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mm
- monster/cr/1-4
- monster/environment/desert
- monster/environment/forest
- monster/environment/swamp
- monster/environment/underwater
- monster/size/large
- monster/type/beast
aliases: ["Constrictor Snake"]
NoteIcon: monster
BestiaryType: beast
SourceType: Bestiary
BookSource: Monster Manual p. 320. Available in the SRD and the Basic Rules.
---
# [Constrictor Snake](3-Mechanics\CLI\bestiary\beast/constrictor-snake.md)
*Source: Monster Manual p. 320. Available in the SRD and the Basic Rules.*  

```statblock
"name": "Constrictor Snake"
"size": "Large"
"type": "beast"
"alignment": "Unaligned"
"ac": !!int "12"
"hp": !!int "13"
"hit_dice": "2d10 + 2"
"stats":
- !!int "15"
- !!int "14"
- !!int "12"
- !!int "1"
- !!int "10"
- !!int "3"
"speed": "30 ft., swim 30 ft."
"senses": "blindsight 10 ft., passive Perception 10"
"languages": ""
"cr": "1/4"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+4 (+4) to hit, reach 5 ft., one creature.\
    \ Hit: dice:1d6 + 2|text(5) (1d6 + 2) piercing damage."
  "name": "Bite"
- "desc": "Melee Weapon Attack: dice: d20+4 (+4) to hit, reach 5 ft., one creature.\
    \ Hit: dice:1d8 + 2|text(6) (1d8 + 2) bludgeoning damage, and the target\
    \ is [grappled](/3-Mechanics/CLI/rules/conditions.md#grappled) (escape DC 14).\
    \ Until this grapple ends, the creature is [restrained](/3-Mechanics/CLI/rules/conditions.md#restrained),\
    \ and the snake can't constrict another target."
  "name": "Constrict"
"source":
- "MM"
- "TftYP"
- "ToA"
- "WDMM"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MM/Constrictor%20Snake.webp"
```
^statblock

```encounter-table
name: Constrictor Snake
creatures:
 - 1: Constrictor Snake
```

## Environment

underwater, forest, swamp, desert