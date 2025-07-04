---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/13
- monster/environment/underdark
- monster/size/gargantuan
- monster/type/aberration
aliases: ["Neothelid"]
NoteIcon: monster
BestiaryType: aberration
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 193, Volo's Guide to Monsters p. 181
---
# [Neothelid](3-Mechanics\CLI\bestiary\aberration/neothelid-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 193, Volo's Guide to Monsters p. 181*  

```statblock
"name": "Neothelid (MPMM)"
"size": "Gargantuan"
"type": "aberration"
"alignment": "Typically  Chaotic Evil"
"ac": !!int "16"
"ac_class": "natural armor"
"hp": !!int "232"
"hit_dice": "15d20 + 75"
"stats":
- !!int "27"
- !!int "7"
- !!int "21"
- !!int "3"
- !!int "16"
- !!int "12"
"speed": "30 ft."
"saves":
  "Charisma": !!int "6"
  "Wisdom": !!int "8"
  "Intelligence": !!int "1"
"senses": "blindsight 120 ft. (blind beyond this radius), passive Perception 13"
"languages": ""
"cr": "13"
"traits":
- "desc": "The neothelid casts one of the following spells, requiring no spell components\
    \ and using Wisdom as the spellcasting ability (spell save DC 16):\n\nAt will:\
    \ [levitate](/3-Mechanics/CLI/spells/levitate.md)\n\n1/day each: [confusion](/3-Mechanics/CLI/spells/confusion.md),\
    \ [feeblemind](/3-Mechanics/CLI/spells/feeblemind.md), [telekinesis](/3-Mechanics/CLI/spells/telekinesis.md)"
  "name": "Spellcasting (Psionics)"
- "desc": "The neothelid is aware of the presence of creatures within 1 mile of it\
    \ that have an Intelligence score of 4 or higher. It knows the distance and direction\
    \ to each creature, as well as each creature's Intelligence score, but can't sense\
    \ anything else about it. A creature protected by a [mind blank](/3-Mechanics/CLI/spells/mind-blank.md)\
    \ spell, a [nondetection](/3-Mechanics/CLI/spells/nondetection.md) spell, or similar\
    \ magic can't be perceived in this manner."
  "name": "Creature Sense"
- "desc": "The neothelid has advantage on saving throws against spells and other magical\
    \ effects."
  "name": "Magic Resistance"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+13 (+13) to hit, reach 15 ft., one\
    \ target. Hit: dice:3d8 + 8|text(21) (3d8 + 8) bludgeoning damage plus dice:2d10|text(11)\
    \ (2d10) psychic damage. If the target is a Large or smaller creature, it must\
    \ succeed on a DC 18 Strength saving throw or be swallowed by the neothelid. A\
    \ swallowed creature is [blinded](/3-Mechanics/CLI/rules/conditions.md#blinded)\
    \ and [restrained](/3-Mechanics/CLI/rules/conditions.md#restrained), it has total\
    \ cover against attacks and other effects outside the neothelid, and it takes\
    \ dice:6d6|text(21) (6d6) acid damage at the start of each of the neothelid's\
    \ turns.\n\nIf the neothelid takes 30 damage or more on a single turn from a creature\
    \ inside it, the neothelid must succeed on a DC 18 Constitution saving throw at\
    \ the end of that turn or regurgitate all swallowed creatures, which fall [prone](/3-Mechanics/CLI/rules/conditions.md#prone)\
    \ in a space within 10 feet of the neothelid. If the neothelid dies, a swallowed\
    \ creature is no longer [restrained](/3-Mechanics/CLI/rules/conditions.md#restrained)\
    \ by it and can escape from the corpse by using 20 feet of movement, exiting [prone](/3-Mechanics/CLI/rules/conditions.md#prone)."
  "name": "Tentacles"
- "desc": "The neothelid exhales acid in a 60-foot cone. Each creature in that area\
    \ must make a DC 18 Dexterity saving throw, taking dice:10d6|text(35) (10d6)\
    \ acid damage on a failed save, or half as much damage on a successful one."
  "name": "Acid Breath (Recharge 5-6)"
"source":
- "MPMM"
- "VGM"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Neothelid.webp"
```
^statblock

```encounter-table
name: Neothelid
creatures:
 - 1: Neothelid
```

## Environment

underdark