---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/4
- monster/environment/underdark
- monster/environment/urban
- monster/size/medium
- monster/type/fiend/demon
aliases: ["Babau"]
NoteIcon: monster
BestiaryType: fiend (demon)
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 52, Volo's Guide to Monsters p. 136
---
# [Babau](3-Mechanics\CLI\bestiary\fiend/babau-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 52, Volo's Guide to Monsters p. 136*  

```statblock
"name": "Babau (MPMM)"
"size": "Medium"
"type": "fiend"
"subtype": "demon"
"alignment": "Typically  Chaotic Evil"
"ac": !!int "16"
"ac_class": "natural armor"
"hp": !!int "82"
"hit_dice": "11d8 + 33"
"stats":
- !!int "19"
- !!int "16"
- !!int "16"
- !!int "11"
- !!int "12"
- !!int "13"
"speed": "40 ft."
"skillsaves":
  "Stealth": !!int "5"
  "Perception": !!int "5"
"damage_resistances": "cold; fire; lightning; bludgeoning, piercing, slashing from\
  \ nonmagical attacks"
"damage_immunities": "poison"
"condition_immunities": "[poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)"
"senses": "darkvision 120 ft., passive Perception 15"
"languages": "Abyssal"
"cr": "4"
"traits":
- "desc": "The babau casts one of the following spells, requiring no material components\
    \ and using Wisdom as the spellcasting ability (spell save DC 11):\n\nAt will:\
    \ [darkness](/3-Mechanics/CLI/spells/darkness.md), [dispel magic](/3-Mechanics/CLI/spells/dispel-magic.md),\
    \ [fear](/3-Mechanics/CLI/spells/fear.md), [heat metal](/3-Mechanics/CLI/spells/heat-metal.md),\
    \ [levitate](/3-Mechanics/CLI/spells/levitate.md)"
  "name": "Spellcasting"
"actions":
- "desc": "The babau makes two Claw attacks. It can replace one attack with a use\
    \ of Spellcasting or Weakening Gaze."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+6 (+6) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d4 + 4|text(6) (1d4 + 4) slashing damage plus dice:1d4|text(2)\
    \ (1d4) acid damage."
  "name": "Claw"
- "desc": "The babau targets one creature that it can see within 20 feet of it. The\
    \ target must make a DC 13 Constitution saving throw. On a failed save, the target\
    \ deals only half damage with weapon attacks that use Strength for 1 minute. The\
    \ target can repeat the saving throw at the end of each of its turns, ending the\
    \ effect on itself on a success."
  "name": "Weakening Gaze"
"source":
- "MPMM"
- "VGM"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Babau.webp"
```
^statblock

```encounter-table
name: Babau
creatures:
 - 1: Babau
```

## Environment

underdark, urban