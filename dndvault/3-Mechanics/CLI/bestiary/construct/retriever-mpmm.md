---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/14
- monster/environment/desert
- monster/environment/forest
- monster/environment/underdark
- monster/size/large
- monster/type/construct
aliases: ["Retriever"]
NoteIcon: monster
BestiaryType: construct
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 209, Mordenkainen's Tome of Foes p. 222
---
# [Retriever](3-Mechanics\CLI\bestiary\construct/retriever-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 209, Mordenkainen's Tome of Foes p. 222*  

```statblock
"name": "Retriever (MPMM)"
"size": "Large"
"type": "construct"
"alignment": "Typically  Lawful Evil"
"ac": !!int "19"
"ac_class": "natural armor"
"hp": !!int "210"
"hit_dice": "20d10 + 100"
"stats":
- !!int "22"
- !!int "16"
- !!int "20"
- !!int "3"
- !!int "11"
- !!int "4"
"speed": "40 ft., climb 40 ft."
"saves":
  "Dexterity": !!int "8"
  "Wisdom": !!int "5"
  "Constitution": !!int "10"
"skillsaves":
  "Stealth": !!int "8"
  "Perception": !!int "5"
"damage_immunities": "necrotic; poison; psychic; bludgeoning, piercing, slashing from\
  \ nonmagical attacks that aren't adamantine"
"condition_immunities": "[charmed](/3-Mechanics/CLI/rules/conditions.md#charmed),\
  \ [exhaustion](/3-Mechanics/CLI/rules/conditions.md#exhaustion), [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened),\
  \ [paralyzed](/3-Mechanics/CLI/rules/conditions.md#paralyzed), [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)"
"senses": "blindsight 30 ft., darkvision 60 ft., passive Perception 15"
"languages": "understands Abyssal, Elvish, and Undercommon but can't speak"
"cr": "14"
"traits":
- "desc": "The retriever casts one of the following spells, requiring no material\
    \ components and using Wisdom as the spellcasting ability (spell save DC 13):\n\
    \n3/day each: [plane shift](/3-Mechanics/CLI/spells/plane-shift.md) (only\
    \ self and up to one incapacitated creature, which is considered willing for the\
    \ spell), [web](/3-Mechanics/CLI/spells/web.md)"
  "name": "Spellcasting"
- "desc": "The retriever is given a quarry by its master. The quarry can be a specific\
    \ creature or object the master is personally acquainted with, or it can be a\
    \ general type of creature or object the master has seen before. The retriever\
    \ knows the direction and distance to its quarry as long as the two of them are\
    \ on the same plane of existence. The retriever can have only one such quarry\
    \ at a time. The retriever also always knows the location of its master."
  "name": "Faultless Tracker"
"actions":
- "desc": "The retriever makes two Foreleg attacks, and it uses Force Beam or Paralyzing\
    \ Beam, if available."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+11 (+11) to hit, reach 10 ft., one\
    \ target. Hit: dice:2d8 + 6|text(15) (2d8 + 6) slashing damage."
  "name": "Foreleg"
- "desc": "The retriever targets one creature it can see within 60 feet of it. The\
    \ target must make a DC 16 Dexterity saving throw, taking dice:5d10|text(27)\
    \ (5d10) force damage on a failed save, or half as much damage on a successful\
    \ one."
  "name": "Force Beam"
- "desc": "The retriever targets one creature it can see within 60 feet of it. The\
    \ target must succeed on a DC 18 Constitution saving throw or be [paralyzed](/3-Mechanics/CLI/rules/conditions.md#paralyzed)\
    \ for 1 minute. The [paralyzed](/3-Mechanics/CLI/rules/conditions.md#paralyzed)\
    \ target can repeat the saving throw at the end of each of its turns, ending the\
    \ effect on itself on a success.\n\nIf the [paralyzed](/3-Mechanics/CLI/rules/conditions.md#paralyzed)\
    \ creature is Medium or smaller, the retriever can pick it up as part of the retriever's\
    \ move and walk or climb with it at full speed."
  "name": "Paralyzing Beam (Recharge 5-6)"
"source":
- "MPMM"
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Retriever.webp"
```
^statblock

```encounter-table
name: Retriever
creatures:
 - 1: Retriever
```

## Environment

desert, forest, underdark