---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/10
- monster/environment/arctic
- monster/environment/forest
- monster/size/medium
- monster/type/fey/elf
aliases: ["Winter Eladrin"]
NoteIcon: monster
BestiaryType: fey (elf)
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 117, Mordenkainen's Tome of Foes p. 197
---
# [Winter Eladrin](3-Mechanics\CLI\bestiary\fey/winter-eladrin-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 117, Mordenkainen's Tome of Foes p. 197*  

```statblock
"name": "Winter Eladrin (MPMM)"
"size": "Medium"
"type": "fey"
"subtype": "elf"
"alignment": "Typically  Chaotic Neutral"
"ac": !!int "19"
"ac_class": "natural armor"
"hp": !!int "165"
"hit_dice": "22d8 + 66"
"stats":
- !!int "11"
- !!int "16"
- !!int "16"
- !!int "18"
- !!int "17"
- !!int "13"
"speed": "30 ft."
"damage_resistances": "cold"
"senses": "darkvision 60 ft., passive Perception 13"
"languages": "Common, Elvish, Sylvan"
"cr": "10"
"traits":
- "desc": "The eladrin casts one of the following spells, requiring no material components\
    \ and using Intelligence as the spellcasting ability (spell save DC 16):\n\nAt\
    \ will: [fog cloud](/3-Mechanics/CLI/spells/fog-cloud.md), [gust of wind](/3-Mechanics/CLI/spells/gust-of-wind.md),\
    \ [sleet storm](/3-Mechanics/CLI/spells/sleet-storm.md)"
  "name": "Spellcasting"
- "desc": "The eladrin has advantage on saving throws against spells and other magical\
    \ effects."
  "name": "Magic Resistance"
- "desc": "Any non-eladrin creature that starts its turn within 60 feet of the eladrin\
    \ must make a DC 13 Wisdom saving throw. On a failed save, the creature becomes\
    \ [charmed](/3-Mechanics/CLI/rules/conditions.md#charmed) by the eladrin for 1\
    \ minute. While [charmed](/3-Mechanics/CLI/rules/conditions.md#charmed) in this\
    \ way, the creature has disadvantage on ability checks and saving throws. The\
    \ [charmed](/3-Mechanics/CLI/rules/conditions.md#charmed) creature can repeat\
    \ the saving throw at the end of each of its turns, ending the effect on itself\
    \ on a success. If a creature's saving throw is successful or the effect ends\
    \ for it, the creature is immune to any eladrin's Sorrowful Presence for the next\
    \ 24 hours.\n\nWhenever the eladrin deals damage to the [charmed](/3-Mechanics/CLI/rules/conditions.md#charmed)\
    \ creature, the [charmed](/3-Mechanics/CLI/rules/conditions.md#charmed) creature\
    \ can repeat the saving throw, ending the effect on itself on a success."
  "name": "Sorrowful Presence"
"actions":
- "desc": "The eladrin makes two Longsword or Longbow attacks. It can replace one\
    \ attack with a use of Spellcasting."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+4 (+4) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d8|text(4) (1d8) slashing damage, or dice:1d10|text(5) (1d10)\
    \ slashing damage if used with two hands, plus dice:3d8|text(13) (3d8) cold\
    \ damage."
  "name": "Longsword"
- "desc": "Ranged Weapon Attack: dice: d20+7 (+7) to hit, range 150/600 ft.,\
    \ one target. Hit: dice:1d8 + 3|text(7) (1d8 + 3) piercing damage plus dice:3d8|text(13)\
    \ (3d8) cold damage."
  "name": "Longbow"
"bonus_actions":
- "desc": "The eladrin teleports, along with any equipment it is wearing or carrying,\
    \ up to 30 feet to an unoccupied space it can see."
  "name": "Fey Step (Recharge 4-6)"
"reactions":
- "desc": "When the eladrin takes damage from a creature the eladrin can see within\
    \ 60 feet of it, the eladrin can force that creature to make a DC 16 Constitution\
    \ saving throw. On a failed save, the creature takes dice:2d10|text(11) (2d10)\
    \ cold damage."
  "name": "Frigid Rebuke"
"source":
- "MPMM"
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Winter%20Eladrin.webp"
```
^statblock

```encounter-table
name: Winter Eladrin
creatures:
 - 1: Winter Eladrin
```

## Environment

arctic, forest