---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/2
- monster/environment/desert
- monster/size/medium
- monster/type/aberration
aliases: ["Berbalang"]
NoteIcon: monster
BestiaryType: aberration
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 61, Mordenkainen's Tome of Foes p. 120
---
# [Berbalang](3-Mechanics\CLI\bestiary\aberration/berbalang-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 61, Mordenkainen's Tome of Foes p. 120*  

```statblock
"name": "Berbalang (MPMM)"
"size": "Medium"
"type": "aberration"
"alignment": "Typically  Neutral Evil"
"ac": !!int "14"
"ac_class": "natural armor"
"hp": !!int "49"
"hit_dice": "14d8 - 14"
"stats":
- !!int "9"
- !!int "16"
- !!int "9"
- !!int "17"
- !!int "11"
- !!int "10"
"speed": "30 ft., fly 40 ft."
"saves":
  "Dexterity": !!int "5"
  "Intelligence": !!int "5"
"skillsaves":
  "Religion": !!int "5"
  "Insight": !!int "2"
  "Perception": !!int "2"
  "History": !!int "5"
  "Arcana": !!int "5"
"senses": "truesight 120 ft., passive Perception 12"
"languages": "all"
"cr": "2"
"traits":
- "desc": "The berbalang casts one of the following spells, requiring no material\
    \ components and using Intelligence as the spellcasting ability:\n\nAt will:\
    \ [speak with dead](/3-Mechanics/CLI/spells/speak-with-dead.md)\n\n1/day:\
    \ [plane shift](/3-Mechanics/CLI/spells/plane-shift.md) (self only)"
  "name": "Spellcasting"
"actions":
- "desc": "The berbalang makes one Bite attack and one\n\nClaw attack."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+5 (+5) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d10 + 3|text(8) (1d10 + 3) piercing damage plus dice:1d8|text(4)\
    \ (1d8) psychic damage."
  "name": "Bite"
- "desc": "Melee Weapon Attack: dice: d20+5 (+5) to hit, reach 5 ft., one target.\
    \ Hit: dice:2d4 + 3|text(8) (2d4 + 3) slashing damage."
  "name": "Claw"
"bonus_actions":
- "desc": "The berbalang creates one spectral duplicate of itself in an unoccupied\
    \ space it can see within 60 feet of it. While the duplicate exists, the berbalang\
    \ is [unconscious](/3-Mechanics/CLI/rules/conditions.md#unconscious). A berbalang\
    \ can have only one duplicate at a time. The duplicate disappears when it or the\
    \ berbalang drops to 0 hit points or when the berbalang dismisses it (no action\
    \ required).\n\nThe duplicate has the same statistics and knowledge as the berbalang,\
    \ and everything experienced by the duplicate is known by the berbalang. All damage\
    \ dealt by the duplicate's attacks is psychic damage."
  "name": "Spectral Duplicate (Recharges after a Short or Long Rest)"
"source":
- "MPMM"
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Berbalang.webp"
```
^statblock

```encounter-table
name: Berbalang
creatures:
 - 1: Berbalang
```

## Environment

desert