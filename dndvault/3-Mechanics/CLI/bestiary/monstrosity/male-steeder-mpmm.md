---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/1-4
- monster/environment/underdark
- monster/size/medium
- monster/type/monstrosity
aliases: ["Male Steeder"]
NoteIcon: monster
BestiaryType: monstrosity
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 231, Mordenkainen's Tome of Foes p. 238
---
# [Male Steeder](3-Mechanics\CLI\bestiary\monstrosity/male-steeder-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 231, Mordenkainen's Tome of Foes p. 238*  

```statblock
"name": "Male Steeder (MPMM)"
"size": "Medium"
"type": "monstrosity"
"alignment": "Unaligned"
"ac": !!int "12"
"ac_class": "natural armor"
"hp": !!int "13"
"hit_dice": "2d8 + 4"
"stats":
- !!int "15"
- !!int "12"
- !!int "14"
- !!int "2"
- !!int "10"
- !!int "3"
"speed": "30 ft., climb 30 ft."
"skillsaves":
  "Stealth": !!int "5"
  "Perception": !!int "4"
"senses": "darkvision 120 ft., passive Perception 14"
"languages": ""
"cr": "1/4"
"traits":
- "desc": "The distance of the steeder's long jumps is tripled; every foot of its\
    \ walking speed that it spends on the jump allows it to jump 3 feet."
  "name": "Extraordinary Leap"
- "desc": "The steeder can climb difficult surfaces, including upside down on ceilings,\
    \ without needing to make an ability check."
  "name": "Spider Climb"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+4 (+4) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d8 + 2|text(6) (1d8 + 2) piercing damage plus dice:1d8|text(4)\
    \ (1d8) poison damage."
  "name": "Bite"
- "desc": "Melee Weapon Attack: dice: d20+4 (+4) to hit, reach 5 ft., one Small\
    \ or Tiny creature. Hit: The target is stuck to the steeder's leg and [grappled](/3-Mechanics/CLI/rules/conditions.md#grappled)\
    \ (escape DC 12). The steeder can have only one creature [grappled](/3-Mechanics/CLI/rules/conditions.md#grappled)\
    \ at a time."
  "name": "Sticky Leg"
"source":
- "MPMM"
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Male%20Steeder.webp"
```
^statblock

```encounter-table
name: Male Steeder
creatures:
 - 1: Male Steeder
```

## Environment

underdark