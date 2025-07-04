---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/2
- monster/environment/forest
- monster/environment/hill
- monster/environment/swamp
- monster/size/medium
- monster/type/monstrosity
aliases: ["Shadow Mastiff"]
NoteIcon: monster
BestiaryType: monstrosity
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 215, Volo's Guide to Monsters p. 190
---
# [Shadow Mastiff](3-Mechanics\CLI\bestiary\monstrosity/shadow-mastiff-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 215, Volo's Guide to Monsters p. 190*  

```statblock
"name": "Shadow Mastiff (MPMM)"
"size": "Medium"
"type": "monstrosity"
"alignment": "Typically  Neutral Evil"
"ac": !!int "12"
"hp": !!int "33"
"hit_dice": "6d8 + 6"
"stats":
- !!int "16"
- !!int "14"
- !!int "13"
- !!int "5"
- !!int "12"
- !!int "5"
"speed": "40 ft."
"skillsaves":
  "Stealth": !!int "6"
  "Perception": !!int "5"
"damage_resistances": "bludgeoning, piercing, slashing from nonmagical attacks while\
  \ in dim light or darkness"
"senses": "darkvision 60 ft., passive Perception 15"
"languages": ""
"cr": "2"
"traits":
- "desc": "The shadow mastiff can see ethereal creatures and objects."
  "name": "Ethereal Awareness"
- "desc": "While in bright light created by sunlight, the shadow mastiff has disadvantage\
    \ on attack rolls, ability checks, and saving throws."
  "name": "Sunlight Weakness"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+5 (+5) to hit, reach 5 ft., one target.\
    \ Hit: dice:2d6 + 3|text(10) (2d6 + 3) piercing damage. If the target is\
    \ a creature, it must succeed on a DC 13 Strength saving throw or be knocked [prone](/3-Mechanics/CLI/rules/conditions.md#prone)."
  "name": "Bite"
"bonus_actions":
- "desc": "While in dim light or darkness, the shadow mastiff becomes [invisible](/3-Mechanics/CLI/rules/conditions.md#invisible),\
    \ along with anything it is wearing or carrying. The invisibility lasts until\
    \ the shadow mastiff uses a bonus action to end it or until the shadow mastiff\
    \ attacks, is in bright light, or is [incapacitated](/3-Mechanics/CLI/rules/conditions.md#incapacitated)."
  "name": "Shadow Blend"
"source":
- "MPMM"
- "VGM"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Shadow%20Mastiff.webp"
```
^statblock

```encounter-table
name: Shadow Mastiff
creatures:
 - 1: Shadow Mastiff
```

## Environment

forest, hill, swamp