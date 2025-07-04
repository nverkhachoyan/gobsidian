---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/1-8
- monster/environment/desert
- monster/environment/mountain
- monster/environment/underdark
- monster/size/small
- monster/type/monstrosity
aliases: ["Young Kruthik"]
NoteIcon: monster
BestiaryType: monstrosity
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 168, Mordenkainen's Tome of Foes p. 211
---
# [Young Kruthik](3-Mechanics\CLI\bestiary\monstrosity/young-kruthik-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 168, Mordenkainen's Tome of Foes p. 211*  

```statblock
"name": "Young Kruthik (MPMM)"
"size": "Small"
"type": "monstrosity"
"alignment": "Unaligned"
"ac": !!int "16"
"ac_class": "natural armor"
"hp": !!int "9"
"hit_dice": "2d6 + 2"
"stats":
- !!int "13"
- !!int "16"
- !!int "13"
- !!int "4"
- !!int "10"
- !!int "6"
"speed": "30 ft., burrow 10 ft., climb 30 ft."
"skillsaves":
  "Perception": !!int "4"
"senses": "darkvision 30 ft., tremorsense 60 ft., passive Perception 14"
"languages": "Kruthik"
"cr": "1/8"
"traits":
- "desc": "The kruthik has advantage on an attack roll against a creature if at least\
    \ one of the kruthik's allies is within 5 feet of the creature and the ally isn't\
    \ [incapacitated](/3-Mechanics/CLI/rules/conditions.md#incapacitated)."
  "name": "Pack Tactics"
- "desc": "The kruthik can burrow through solid rock at half its burrowing speed and\
    \ leaves a 2½-foot-diameter tunnel in its wake."
  "name": "Tunneler"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+5 (+5) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d4 + 3|text(5) (1d4 + 3) piercing damage."
  "name": "Stab"
"source":
- "MPMM"
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Young%20Kruthik.webp"
```
^statblock

```encounter-table
name: Young Kruthik
creatures:
 - 1: Young Kruthik
```

## Environment

desert, mountain, underdark