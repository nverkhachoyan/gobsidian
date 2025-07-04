---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/5
- monster/environment/coastal
- monster/environment/underwater
- monster/size/medium
- monster/type/monstrosity
aliases: ["Kraken Priest"]
NoteIcon: monster
BestiaryType: monstrosity
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 167, Volo's Guide to Monsters p. 215
---
# [Kraken Priest](3-Mechanics\CLI\bestiary\monstrosity/kraken-priest-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 167, Volo's Guide to Monsters p. 215*  

```statblock
"name": "Kraken Priest (MPMM)"
"size": "Medium"
"type": "monstrosity"
"alignment": "Typically  Chaotic Evil"
"ac": !!int "15"
"ac_class": "natural armor"
"hp": !!int "75"
"hit_dice": "10d8 + 30"
"stats":
- !!int "12"
- !!int "10"
- !!int "16"
- !!int "10"
- !!int "15"
- !!int "14"
"speed": "30 ft., swim 30 ft."
"skillsaves":
  "Perception": !!int "5"
"senses": "passive Perception 15"
"languages": "any two languages"
"cr": "5"
"traits":
- "desc": "The priest casts one of the following spells, requiring no material components\
    \ and using Wisdom as the spellcasting ability (spell save DC 13):\n\nAt will:\
    \ [command](/3-Mechanics/CLI/spells/command.md), [create or destroy water](/3-Mechanics/CLI/spells/create-or-destroy-water.md)\n\
    \n1/day: [Evard's black tentacles](/3-Mechanics/CLI/spells/evards-black-tentacles.md)\n\
    \n3/day each: [control water](/3-Mechanics/CLI/spells/control-water.md), [darkness](/3-Mechanics/CLI/spells/darkness.md),\
    \ [water breathing](/3-Mechanics/CLI/spells/water-breathing.md), [water walk](/3-Mechanics/CLI/spells/water-walk.md)"
  "name": "Spellcasting"
- "desc": "The priest can breathe air and water."
  "name": "Amphibious"
"actions":
- "desc": "The priest makes two Thunderous Touch or Thunderbolt attacks."
  "name": "Multiattack"
- "desc": "Melee Spell Attack: dice: d20+5 (+5) to hit, reach 5 ft., one target.\
    \ Hit: dice:5d10|text(27) (5d10) thunder damage."
  "name": "Thunderous Touch"
- "desc": "Ranged Spell Attack: dice: d20+5 (+5) to hit, range 60 ft., one target.\
    \ Hit: dice:2d10|text(11) (2d10) lightning damage plus dice:2d10|text(11)\
    \ (2d10) thunder damage, and the target is knocked [prone](/3-Mechanics/CLI/rules/conditions.md#prone)."
  "name": "Thunderbolt"
- "desc": "A kraken speaks through the priest with a thunderous voice audible within\
    \ 300 feet. Creatures of the priest's choice that can hear the kraken's words\
    \ (which are spoken in Abyssal, Infernal, or Primordial) must succeed on a DC\
    \ 14 Wisdom saving throw or be [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened)\
    \ of the priest for 1 minute. A [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened)\
    \ target can repeat the saving throw at the end of each of its turns, ending the\
    \ effect on itself on a success."
  "name": "Voice of the Kraken (Recharges after a Short or Long Rest)"
"source":
- "MPMM"
- "VGM"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Kraken%20Priest.webp"
```
^statblock

```encounter-table
name: Kraken Priest
creatures:
 - 1: Kraken Priest
```

## Environment

coastal, underwater