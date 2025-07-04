---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/slw
- monster/cr/3
- monster/size/medium
- monster/type/fiend
aliases: ["Tooth-N-Claw"]
NoteIcon: monster
BestiaryType: fiend
SourceType: Bestiary
BookSource: Storm Lord's Wrath
---
# [Tooth-N-Claw](3-Mechanics\CLI\bestiary\npc/tooth-n-claw-slw.md)
*Source: Storm Lord's Wrath*  

```statblock
"name": "Tooth-N-Claw (SLW)"
"size": "Medium"
"type": "fiend"
"alignment": "Lawful Evil"
"ac": !!int "15"
"ac_class": "natural armor"
"hp": !!int "45"
"hit_dice": "7d8 + 14"
"stats":
- !!int "17"
- !!int "12"
- !!int "14"
- !!int "6"
- !!int "13"
- !!int "6"
"speed": "50 ft."
"skillsaves":
  "Perception": !!int "5"
"damage_immunities": "cold"
"senses": "darkvision 60 ft., passive Perception 15"
"languages": "understands Infernal but can't speak it"
"cr": "3"
"traits":
- "desc": "Tooth-N-Claw has advantage on Wisdom ([Perception](/3-Mechanics/CLI/rules/skills.md#Perception))\
    \ checks that rely on hearing or smell."
  "name": "Keen Hearing and Smell"
- "desc": "Tooth-N-Claw has advantage on an attack roll against a creature if at least\
    \ one of its allies is within 5 feet of the creature and the ally isn't [incapacitated](/3-Mechanics/CLI/rules/conditions.md#incapacitated)."
  "name": "Pack Tactics"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+5 (+5) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d8 + 3|text(7) (1d8 + 3) piercing damage plus dice:2d6|text(7)\
    \ (2d6) cold damage."
  "name": "Bite"
- "desc": "Tooth-N-Claw exhales an icy blast in a 15-foot cone. Each creature in that\
    \ area must make a DC 12 Dexterity saving throw, taking dice:6d6|text(21) (6d6)\
    \ cold damage on a failed save, or half as much damage on a successful one."
  "name": "Freezing Breath (Recharge 5-6)"
"source":
- "SLW"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/SLW/Tooth-N-Claw.webp"
```
^statblock

```encounter-table
name: Tooth-N-Claw
creatures:
 - 1: Tooth-N-Claw
```