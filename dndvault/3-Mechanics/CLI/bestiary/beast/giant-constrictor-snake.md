---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mm
- monster/cr/2
- monster/environment/desert
- monster/environment/forest
- monster/environment/swamp
- monster/environment/underdark
- monster/environment/underwater
- monster/size/huge
- monster/type/beast
aliases: ["Giant Constrictor Snake"]
NoteIcon: monster
BestiaryType: beast
SourceType: Bestiary
BookSource: Monster Manual p. 324, Dragon of Icespire Peak, Storm Lord's Wrath. Available in the SRD and the Basic Rules.
---
# [Giant Constrictor Snake](3-Mechanics\CLI\bestiary\beast/giant-constrictor-snake.md)
*Source: Monster Manual p. 324, Dragon of Icespire Peak, Storm Lord's Wrath. Available in the SRD and the Basic Rules.*  

```statblock
"name": "Giant Constrictor Snake"
"size": "Huge"
"type": "beast"
"alignment": "Unaligned"
"ac": !!int "12"
"hp": !!int "60"
"hit_dice": "8d12 + 8"
"stats":
- !!int "19"
- !!int "14"
- !!int "12"
- !!int "1"
- !!int "10"
- !!int "3"
"speed": "30 ft., swim 30 ft."
"skillsaves":
  "Perception": !!int "2"
"senses": "blindsight 10 ft., passive Perception 12"
"languages": ""
"cr": "2"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+6 (+6) to hit, reach 10 ft., one creature.\
    \ Hit: dice:2d6 + 4|text(11) (2d6 + 4) piercing damage."
  "name": "Bite"
- "desc": "Melee Weapon Attack: dice: d20+6 (+6) to hit, reach 5 ft., one creature.\
    \ Hit: dice:2d8 + 4|text(13) (2d8 + 4) bludgeoning damage, and the target\
    \ is [grappled](/3-Mechanics/CLI/rules/conditions.md#grappled) (escape DC 16).\
    \ Until this grapple ends, the creature is [restrained](/3-Mechanics/CLI/rules/conditions.md#restrained),\
    \ and the snake can't constrict another target."
  "name": "Constrict"
"source":
- "MM"
- "ToA"
- "GoS"
- "DIP"
- "SLW"
- "EGW"
- "WBtW"
- "PaBTSO"
- "QftIS"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MM/Giant%20Constrictor%20Snake.webp"
```
^statblock

```encounter-table
name: Giant Constrictor Snake
creatures:
 - 1: Giant Constrictor Snake
```

## Environment

underwater, underdark, forest, swamp, desert