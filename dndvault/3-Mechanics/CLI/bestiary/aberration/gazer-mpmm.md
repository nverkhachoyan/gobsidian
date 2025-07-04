---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/1-2
- monster/environment/underdark
- monster/size/tiny
- monster/type/aberration/beholder
aliases: ["Gazer"]
NoteIcon: monster
BestiaryType: aberration (beholder)
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 134, Volo's Guide to Monsters p. 126
---
# [Gazer](3-Mechanics\CLI\bestiary\aberration/gazer-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 134, Volo's Guide to Monsters p. 126*  

```statblock
"name": "Gazer (MPMM)"
"size": "Tiny"
"type": "aberration"
"subtype": "beholder"
"alignment": "Typically  Neutral Evil"
"ac": !!int "13"
"hp": !!int "13"
"hit_dice": "3d4 + 6"
"stats":
- !!int "3"
- !!int "17"
- !!int "14"
- !!int "3"
- !!int "10"
- !!int "7"
"speed": "0 ft., fly 30 ft. (hover)"
"saves":
  "Wisdom": !!int "2"
"skillsaves":
  "Stealth": !!int "5"
  "Perception": !!int "4"
"condition_immunities": "[prone](/3-Mechanics/CLI/rules/conditions.md#prone)"
"senses": "darkvision 60 ft., passive Perception 14"
"languages": ""
"cr": "1/2"
"traits":
- "desc": "The gazer can mimic simple sounds of speech it has heard, in any language.\
    \ A creature that hears the sounds can tell they are imitations with a successful\
    \ DC 10 Wisdom ([Insight](/3-Mechanics/CLI/rules/skills.md#Insight)) check."
  "name": "Mimicry"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+5 (+5) to hit, reach 5 ft., one target.\
    \ Hit: 1 piercing damage."
  "name": "Bite"
- "desc": "The gazer shoots two of the following magical eye rays at random (roll\
    \ two dice: d4|avg|noform (d4)s, and reroll duplicates), choosing one or two\
    \ targets it can see within 60 feet of it:\n\n- 1 Dazing Ray. The targeted\
    \ creature must succeed on a DC 12 Wisdom saving throw or be [charmed](/3-Mechanics/CLI/rules/conditions.md#charmed)\
    \ until the start of the gazer's next turn. While the target is [charmed](/3-Mechanics/CLI/rules/conditions.md#charmed)\
    \ in this way, its speed is halved, and it has disadvantage on attack rolls. \
    \ \n- 2 Fear Ray. The targeted creature must succeed on a DC 12 Wisdom saving\
    \ throw or be [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened) until\
    \ the start of the gazer's next turn.  \n- 3 Frost Ray. The target must succeed\
    \ on a DC 12 Dexterity saving throw or take dice:3d6|text(10) (3d6) cold damage.\
    \  \n- 4 Telekinetic Ray. If the target is a creature that is Medium or smaller,\
    \ it must succeed on a DC 12 Strength saving throw or be moved up to 30 feet directly\
    \ away from the gazer. If the target is a Tiny object that isn't being worn or\
    \ carried, the gazer moves it up to 30 feet in any direction. The gazer can also\
    \ exert fine control on objects with this ray, such as manipulating a simple tool\
    \ or opening a container.  "
  "name": "Eye Rays"
"bonus_actions":
- "desc": "The gazer moves up to its speed toward a hostile creature that it can see."
  "name": "Aggressive"
"source":
- "MPMM"
- "VGM"
- "SjA"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Gazer.webp"
```
^statblock

```encounter-table
name: Gazer
creatures:
 - 1: Gazer
```

## Environment

underdark