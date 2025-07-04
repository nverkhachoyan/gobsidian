---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/3
- monster/environment/underdark
- monster/size/large
- monster/type/monstrosity
aliases: ["Trapper"]
NoteIcon: monster
BestiaryType: monstrosity
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 245, Volo's Guide to Monsters p. 194
---
# [Trapper](3-Mechanics\CLI\bestiary\monstrosity/trapper-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 245, Volo's Guide to Monsters p. 194*  

```statblock
"name": "Trapper (MPMM)"
"size": "Large"
"type": "monstrosity"
"alignment": "Unaligned"
"ac": !!int "13"
"ac_class": "natural armor"
"hp": !!int "68"
"hit_dice": "8d10 + 24"
"stats":
- !!int "17"
- !!int "10"
- !!int "17"
- !!int "2"
- !!int "13"
- !!int "4"
"speed": "20 ft., climb 20 ft."
"skillsaves":
  "Stealth": !!int "2"
"senses": "blindsight 30 ft., darkvision 60 ft., passive Perception 11"
"languages": ""
"cr": "3"
"traits":
- "desc": "If the trapper is motionless on a floor, wall, or ceiling at the start\
    \ of combat, it has advantage on its initiative roll. Moreover, if a creature\
    \ hasn't observed the trapper move or act, that creature must succeed on a DC\
    \ 18 Intelligence ([Investigation](/3-Mechanics/CLI/rules/skills.md#Investigation))\
    \ check to discern that the trapper isn't part of the floor, wall, or ceiling."
  "name": "False Appearance"
- "desc": "The trapper can climb difficult surfaces, including upside down on ceilings,\
    \ without needing to make an ability check."
  "name": "Spider Climb"
"actions":
- "desc": "One Large or smaller creature within 10 feet of the trapper must succeed\
    \ on a DC 13 Dexterity saving throw or be [grappled](/3-Mechanics/CLI/rules/conditions.md#grappled)\
    \ (escape DC 14). Until the grapple ends, the target takes dice:3d6 + 3|text(13)\
    \ (3d6 + 3) bludgeoning damage plus dice:1d6|text(3) (1d6) acid damage at\
    \ the start of each of its turns. While [grappled](/3-Mechanics/CLI/rules/conditions.md#grappled)\
    \ in this way, the target is [restrained](/3-Mechanics/CLI/rules/conditions.md#restrained),\
    \ [blinded](/3-Mechanics/CLI/rules/conditions.md#blinded), and deprived of air.\
    \ The trapper can smother only one creature at a time."
  "name": "Smother"
"source":
- "MPMM"
- "VGM"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Trapper.webp"
```
^statblock

```encounter-table
name: Trapper
creatures:
 - 1: Trapper
```

## Environment

underdark