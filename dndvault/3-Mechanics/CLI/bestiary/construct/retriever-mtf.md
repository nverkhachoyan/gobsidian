---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mtf
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
BookSource: Mordenkainen's Tome of Foes p. 222
---
# [Retriever](3-Mechanics\CLI\bestiary\construct/retriever-mtf.md)
*Source: Mordenkainen's Tome of Foes p. 222*  

## Retriever

The retriever is a potent, spider-like construct conceived and built by the drow for one original purpose—to prowl the Abyss and capture demons for the drow to enslave or use in their rituals. The automatons proved so effective and so fearsome that they now perform many different missions.

## Demon-Infused

Although each retriever is a metal-and-magic construct, it houses the imprisoned spirit of a bebilith. Most of the demon's intellect has been distilled away, leaving only its cruelty and cunning.

## Lethal Collectors

Though they were created to operate only in the Abyss, retrievers are sometimes dispatched when a powerful drow needs some creature or object captured and brought back alive and intact. Only under the rarest of circumstances is a retriever handed over or sold to a non-drow, since the dark elves don't want to take the chance that the construct will be turned against them.

```statblock
"name": "Retriever (MTF)"
"size": "Large"
"type": "construct"
"alignment": "Lawful Evil"
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
- "desc": "The retriever's innate spellcasting ability is Wisdom (spell save DC 13).\
    \ The retriever can innately cast the following spells, requiring no material\
    \ components.\n\n3/day each: [plane shift](/3-Mechanics/CLI/spells/plane-shift.md)\
    \ (only self and up to one incapacitated creature which is considered willing\
    \ for the spell), [web](/3-Mechanics/CLI/spells/web.md)"
  "name": "Innate Spellcasting"
- "desc": "The retriever is given a quarry by its master. The quarry can be a specific\
    \ creature or object the master is personally acquainted with, or it can be a\
    \ general type of creature or object the master has seen before. The retriever\
    \ knows the direction and distance to its quarry as long as the two of them are\
    \ on the same plane of existence. The retriever can have only one such quarry\
    \ at a time. The retriever also always knows the location of its master."
  "name": "Faultless Tracker"
"actions":
- "desc": "The retriever makes two foreleg attacks and uses its force or paralyzing\
    \ beam once, if available."
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
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MTF/Retriever.webp"
```
^statblock

```encounter-table
name: Retriever
creatures:
 - 1: Retriever
```

## Environment

desert, forest, underdark