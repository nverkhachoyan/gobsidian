---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/1
- monster/environment/forest
- monster/environment/mountain
- monster/environment/underdark
- monster/size/small
- monster/type/aberration
aliases: ["Choker"]
NoteIcon: monster
BestiaryType: aberration
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 76, Mordenkainen's Tome of Foes p. 123
---
# [Choker](3-Mechanics\CLI\bestiary\aberration/choker-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 76, Mordenkainen's Tome of Foes p. 123*  

```statblock
"name": "Choker (MPMM)"
"size": "Small"
"type": "aberration"
"alignment": "Typically  Chaotic Evil"
"ac": !!int "16"
"ac_class": "natural armor"
"hp": !!int "13"
"hit_dice": "3d6 + 3"
"stats":
- !!int "16"
- !!int "14"
- !!int "13"
- !!int "4"
- !!int "12"
- !!int "7"
"speed": "30 ft."
"skillsaves":
  "Stealth": !!int "6"
"senses": "darkvision 60 ft., passive Perception 11"
"languages": "Deep Speech"
"cr": "1"
"traits":
- "desc": "The choker can take an extra action on its turn."
  "name": "Aberrant Quickness (Recharges after a Short or Long Rest)"
- "desc": "The choker can move through and occupy a space as narrow as 4 inches wide\
    \ without squeezing."
  "name": "Boneless"
- "desc": "The choker can climb difficult surfaces, including upside down on ceilings,\
    \ without needing to make an ability check."
  "name": "Spider Climb"
"actions":
- "desc": "The choker makes two Tentacle attacks."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+5 (+5) to hit, reach 10 ft., one target.\
    \ Hit: dice:1d4 + 3|text(5) (1d4 + 3) piercing damage. If the target is\
    \ a Large or smaller creature, it is [grappled](/3-Mechanics/CLI/rules/conditions.md#grappled)\
    \ (escape DC 15). Until this grapple ends, the target is [restrained](/3-Mechanics/CLI/rules/conditions.md#restrained),\
    \ and the choker can't use this tentacle on another target. The choker has two\
    \ tentacles. If this attack is a critical hit, the target also can't breathe or\
    \ speak until the grapple ends."
  "name": "Tentacle"
"source":
- "MPMM"
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Choker.webp"
```
^statblock

```encounter-table
name: Choker
creatures:
 - 1: Choker
```

## Environment

forest, mountain, underdark