---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/10
- monster/environment/forest
- monster/environment/grassland
- monster/size/medium
- monster/type/fey/elf
aliases: ["Spring Eladrin"]
NoteIcon: monster
BestiaryType: fey (elf)
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 116, Mordenkainen's Tome of Foes p. 196
---
# [Spring Eladrin](3-Mechanics\CLI\bestiary\fey/spring-eladrin-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 116, Mordenkainen's Tome of Foes p. 196*  

```statblock
"name": "Spring Eladrin (MPMM)"
"size": "Medium"
"type": "fey"
"subtype": "elf"
"alignment": "Typically  Chaotic Neutral"
"ac": !!int "19"
"ac_class": "natural armor"
"hp": !!int "165"
"hit_dice": "22d8 + 66"
"stats":
- !!int "14"
- !!int "16"
- !!int "16"
- !!int "18"
- !!int "11"
- !!int "18"
"speed": "30 ft."
"skillsaves":
  "Deception": !!int "8"
  "Persuasion": !!int "8"
"damage_resistances": "psychic"
"senses": "darkvision 60 ft., passive Perception 10"
"languages": "Common, Elvish, Sylvan"
"cr": "10"
"traits":
- "desc": "The eladrin casts one of the following spells, requiring no material components\
    \ and using Charisma as the spellcasting ability (spell save DC 16):\n\nAt will:\
    \ [Tasha's hideous laughter](/3-Mechanics/CLI/spells/tashas-hideous-laughter.md)\n\
    \n1/day each: [major image](/3-Mechanics/CLI/spells/major-image.md), [suggestion](/3-Mechanics/CLI/spells/suggestion.md)"
  "name": "Spellcasting"
- "desc": "Any non-eladrin creature that starts its turn within 60 feet of the eladrin\
    \ must make a DC 16 Wisdom saving throw. On a failed save, the creature becomes\
    \ [charmed](/3-Mechanics/CLI/rules/conditions.md#charmed) by the eladrin for 1\
    \ minute. On a successful save, the creature becomes immune to any eladrin's Joyful\
    \ Presence for 24 hours.\n\nWhenever the eladrin deals damage to the [charmed](/3-Mechanics/CLI/rules/conditions.md#charmed)\
    \ creature, the [charmed](/3-Mechanics/CLI/rules/conditions.md#charmed) creature\
    \ can repeat the saving throw, ending the effect on itself on a success."
  "name": "Joyful Presence"
- "desc": "The eladrin has advantage on saving throws against spells and other magical\
    \ effects."
  "name": "Magic Resistance"
"actions":
- "desc": "The eladrin makes two Longsword or Longbow attacks. It can replace one\
    \ attack with a use of Spellcasting."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+6 (+6) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d8 + 2|text(6) (1d8 + 2) slashing damage, or dice:1d10 + 2|text(7)\
    \ (1d10 + 2) slashing damage if used with two hands, plus dice:5d8|text(22)\
    \ (5d8) psychic damage."
  "name": "Longsword"
- "desc": "Ranged Weapon Attack: dice: d20+7 (+7) to hit, range 150/600 ft.,\
    \ one target. Hit: dice:1d8 + 3|text(7) (1d8 + 3) piercing damage plus dice:5d8|text(22)\
    \ (5d8) psychic damage."
  "name": "Longbow"
"bonus_actions":
- "desc": "The eladrin teleports, along with any equipment it is wearing or carrying,\
    \ up to 30 feet to an unoccupied space it can see."
  "name": "Fey Step (Recharge 4-6)"
"source":
- "MPMM"
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Spring%20Eladrin.webp"
```
^statblock

```encounter-table
name: Spring Eladrin
creatures:
 - 1: Spring Eladrin
```

## Environment

forest, grassland