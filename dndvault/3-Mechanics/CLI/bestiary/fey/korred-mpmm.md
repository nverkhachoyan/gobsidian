---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/7
- monster/environment/forest
- monster/size/small
- monster/type/fey
aliases: ["Korred"]
NoteIcon: monster
BestiaryType: fey
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 166, Volo's Guide to Monsters p. 168
---
# [Korred](3-Mechanics\CLI\bestiary\fey/korred-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 166, Volo's Guide to Monsters p. 168*  

```statblock
"name": "Korred (MPMM)"
"size": "Small"
"type": "fey"
"alignment": "Typically  Chaotic Neutral"
"ac": !!int "17"
"ac_class": "natural armor"
"hp": !!int "93"
"hit_dice": "11d6 + 55"
"stats":
- !!int "23"
- !!int "14"
- !!int "20"
- !!int "10"
- !!int "15"
- !!int "9"
"speed": "30 ft., burrow 30 ft."
"skillsaves":
  "Athletics": !!int "9"
  "Stealth": !!int "5"
  "Perception": !!int "5"
"damage_resistances": "bludgeoning, piercing, slashing from nonmagical attacks"
"senses": "darkvision 120 ft., tremorsense 120 ft., passive Perception 15"
"languages": "Dwarvish, Gnomish, Sylvan, Terran, Undercommon"
"cr": "7"
"traits":
- "desc": "The korred casts one of the following spells, requiring no spell components\
    \ and using Wisdom as the spellcasting ability (spell save DC 13):\n\nAt will:\
    \ [commune with nature](/3-Mechanics/CLI/spells/commune-with-nature.md) (as an\
    \ action), [meld into stone](/3-Mechanics/CLI/spells/meld-into-stone.md), [stone\
    \ shape](/3-Mechanics/CLI/spells/stone-shape.md)\n\n1/day: [Otto's irresistible\
    \ dance](/3-Mechanics/CLI/spells/ottos-irresistible-dance.md)"
  "name": "Spellcasting"
- "desc": "The korred has advantage on Dexterity ([Stealth](/3-Mechanics/CLI/rules/skills.md#Stealth))\
    \ checks made to hide in rocky terrain."
  "name": "Stone Camouflage"
"actions":
- "desc": "The korred makes two Greatclub or Rock attacks."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+9 (+9) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d8 + 6|text(10) (1d8 + 6) bludgeoning damage, or dice:3d8\
    \ + 6|text(19) (3d8 + 6) bludgeoning damage if the korred is on the ground."
  "name": "Greatclub"
- "desc": "Ranged Weapon Attack: dice: d20+9 (+9) to hit, range 60/120 ft.,\
    \ one target. Hit: dice:1d8 + 6|text(10) (1d8 + 6) bludgeoning damage, or\
    \ dice:3d8 + 6|text(19) (3d8 + 6) bludgeoning damage if the korred is on the\
    \ ground."
  "name": "Rock"
"bonus_actions":
- "desc": "The korred has at least one 50-foot-long rope woven out of its hair. The\
    \ korred commands one such rope within 30 feet of it to move up to 20 feet and\
    \ entangle a Large or smaller creature that the korred can see. The target must\
    \ succeed on a DC 13 Dexterity saving throw or become [grappled](/3-Mechanics/CLI/rules/conditions.md#grappled)\
    \ by the rope (escape DC 13). Until this grapple ends, the target is [restrained](/3-Mechanics/CLI/rules/conditions.md#restrained).\
    \ The korred can use a bonus action to release the target, which is also freed\
    \ if the korred dies or becomes [incapacitated](/3-Mechanics/CLI/rules/conditions.md#incapacitated).\n\
    \nA rope of korred hair has AC 20 and 20 hit points. It regains 1 hit point at\
    \ the start of each of the korred's turns while the rope has at least 1 hit point\
    \ and the korred is alive. If the rope drops to 0 hit points, it is destroyed."
  "name": "Command Hair"
"source":
- "MPMM"
- "VGM"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Korred.webp"
```
^statblock

```encounter-table
name: Korred
creatures:
 - 1: Korred
```

## Environment

forest