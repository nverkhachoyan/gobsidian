---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/3
- monster/environment/hill
- monster/environment/underdark
- monster/size/small
- monster/type/aberration
aliases: ["Neogi"]
NoteIcon: monster
BestiaryType: aberration
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 192, Volo's Guide to Monsters p. 180
---
# [Neogi](3-Mechanics\CLI\bestiary\aberration/neogi-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 192, Volo's Guide to Monsters p. 180*  

```statblock
"name": "Neogi (MPMM)"
"size": "Small"
"type": "aberration"
"alignment": "Typically  Lawful Evil"
"ac": !!int "15"
"ac_class": "natural armor"
"hp": !!int "33"
"hit_dice": "6d6 + 12"
"stats":
- !!int "6"
- !!int "16"
- !!int "14"
- !!int "13"
- !!int "12"
- !!int "15"
"speed": "30 ft., climb 30 ft."
"skillsaves":
  "Intimidation": !!int "4"
  "Perception": !!int "3"
"senses": "darkvision 60 ft., passive Perception 13"
"languages": "Common, Deep Speech, Undercommon"
"cr": "3"
"traits":
- "desc": "The neogi has advantage on saving throws against being [charmed](/3-Mechanics/CLI/rules/conditions.md#charmed)\
    \ or [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened), and magic\
    \ can't put the neogi to sleep."
  "name": "Mental Fortitude"
- "desc": "The neogi can climb difficult surfaces, including upside down on ceilings,\
    \ without needing to make an ability check."
  "name": "Spider Climb"
"actions":
- "desc": "The neogi makes one Bite attack and two Claw attacks."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+5 (+5) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d6 + 3|text(6) (1d6 + 3) piercing damage plus dice:4d6|text(14)\
    \ (4d6) poison damage, and the target must succeed on a DC 12 Constitution saving\
    \ throw or become [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned) for\
    \ 1 minute. A target can repeat the saving throw at the end of each of its turns,\
    \ ending the effect on itself on a success."
  "name": "Bite"
- "desc": "Melee Weapon Attack: dice: d20+5 (+5) to hit, reach 5 ft., one target.\
    \ Hit: dice:2d4 + 3|text(8) (2d4 + 3) slashing damage."
  "name": "Claw"
"bonus_actions":
- "desc": "The neogi targets one creature it can see within 30 feet of it. The target\
    \ must succeed on a DC 14 Wisdom saving throw or be magically [charmed](/3-Mechanics/CLI/rules/conditions.md#charmed)\
    \ by the neogi for 1 day, or until the neogi dies or is more than 1 mile from\
    \ the target. The [charmed](/3-Mechanics/CLI/rules/conditions.md#charmed) target\
    \ obeys the neogi's commands and can't take reactions, and the neogi and the target\
    \ can communicate telepathically with each other at a distance of up to 1 mile.\
    \ Whenever the [charmed](/3-Mechanics/CLI/rules/conditions.md#charmed) target\
    \ takes damage, it can repeat the saving throw, ending the effect on itself on\
    \ a success."
  "name": "Enslave (Recharges after a Short or Long Rest)"
"source":
- "MPMM"
- "VGM"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Neogi.webp"
```
^statblock

```encounter-table
name: Neogi
creatures:
 - 1: Neogi
```

## Environment

hill, underdark