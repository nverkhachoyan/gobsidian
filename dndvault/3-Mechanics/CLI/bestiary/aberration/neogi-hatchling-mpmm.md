---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/1-8
- monster/environment/hill
- monster/environment/underdark
- monster/size/tiny
- monster/type/aberration
aliases: ["Neogi Hatchling"]
NoteIcon: monster
BestiaryType: aberration
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 191, Volo's Guide to Monsters p. 179
---
# [Neogi Hatchling](3-Mechanics\CLI\bestiary\aberration/neogi-hatchling-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 191, Volo's Guide to Monsters p. 179*  

```statblock
"name": "Neogi Hatchling (MPMM)"
"size": "Tiny"
"type": "aberration"
"alignment": "Typically  Lawful Evil"
"ac": !!int "11"
"hp": !!int "7"
"hit_dice": "3d4"
"stats":
- !!int "3"
- !!int "13"
- !!int "10"
- !!int "6"
- !!int "10"
- !!int "9"
"speed": "20 ft., climb 20 ft."
"senses": "darkvision 60 ft., passive Perception 10"
"languages": ""
"cr": "1/8"
"traits":
- "desc": "The neogi has advantage on saving throws against being [charmed](/3-Mechanics/CLI/rules/conditions.md#charmed)\
    \ or [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened), and magic\
    \ can't put the neogi to sleep."
  "name": "Mental Fortitude"
- "desc": "The neogi can climb difficult surfaces, including upside down on ceilings,\
    \ without needing to make an ability check."
  "name": "Spider Climb"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+3 (+3) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d4 + 1|text(3) (1d4 + 1) piercing damage plus dice:1d6|text(3)\
    \ (1d6) poison damage, and the target must succeed on a DC 10 Constitution saving\
    \ throw or become [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned) for\
    \ 1 minute. A target can repeat the saving throw at the end of each of its turns,\
    \ ending the effect on itself on a success."
  "name": "Bite"
"source":
- "MPMM"
- "VGM"
- "SjA"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Neogi%20Hatchling.webp"
```
^statblock

```encounter-table
name: Neogi Hatchling
creatures:
 - 1: Neogi Hatchling
```

## Environment

hill, underdark