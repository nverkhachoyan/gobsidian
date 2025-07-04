---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/12
- monster/environment/mountain
- monster/environment/underdark
- monster/size/medium
- monster/type/humanoid/dwarf
aliases: ["Duergar Despot"]
NoteIcon: monster
BestiaryType: humanoid (dwarf)
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 107, Mordenkainen's Tome of Foes p. 188
---
# [Duergar Despot](3-Mechanics\CLI\bestiary\humanoid/duergar-despot-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 107, Mordenkainen's Tome of Foes p. 188*  

```statblock
"name": "Duergar Despot (MPMM)"
"size": "Medium"
"type": "humanoid"
"subtype": "dwarf"
"alignment": "Any alignment"
"ac": !!int "21"
"ac_class": "natural armor"
"hp": !!int "119"
"hit_dice": "14d8 + 56"
"stats":
- !!int "20"
- !!int "5"
- !!int "19"
- !!int "15"
- !!int "14"
- !!int "13"
"speed": "25 ft."
"saves":
  "Wisdom": !!int "6"
  "Constitution": !!int "8"
"damage_immunities": "poison"
"condition_immunities": "[charmed](/3-Mechanics/CLI/rules/conditions.md#charmed),\
  \ [exhaustion](/3-Mechanics/CLI/rules/conditions.md#exhaustion), [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened),\
  \ [paralyzed](/3-Mechanics/CLI/rules/conditions.md#paralyzed), [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)"
"senses": "darkvision 120 ft., passive Perception 12"
"languages": "Dwarvish, Undercommon"
"cr": "12"
"traits":
- "desc": "The duergar casts one of the following spells, requiring no spell components\
    \ and using Intelligence as the spellcasting ability (spell save DC 12):\n\nAt\
    \ will: [mage hand](/3-Mechanics/CLI/spells/mage-hand.md), [minor illusion](/3-Mechanics/CLI/spells/minor-illusion.md)\n\
    \n1/day: [stinking cloud](/3-Mechanics/CLI/spells/stinking-cloud.md)"
  "name": "Spellcasting (Psionics)"
- "desc": "The duergar has advantage on saving throws against spells and other magical\
    \ effects."
  "name": "Magic Resistance"
- "desc": "When the duergar suffers a critical hit or is reduced to 0 hit points,\
    \ psychic energy erupts from its frame to deal dice:4d6|text(14) (4d6) psychic\
    \ damage to each creature within 5 feet of it."
  "name": "Psychic Engine"
- "desc": "While in sunlight, the duergar has disadvantage on attack rolls, as well\
    \ as on Wisdom ([Perception](/3-Mechanics/CLI/rules/skills.md#Perception)) checks\
    \ that rely on sight."
  "name": "Sunlight Sensitivity"
"actions":
- "desc": "The duergar makes two Iron Fist attacks and two Stomping Foot attacks.\
    \ After one of the attacks, the duergar can move up to its speed without provoking\
    \ opportunity attacks. It can replace one of the attacks with a use of Flame Jet."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+9 (+9) to hit, reach 5 ft., one target.\
    \ Hit: dice:4d8 + 5|text(23) (4d8 + 5) bludgeoning damage. If the target\
    \ is a Large or smaller creature, it must succeed on a DC 17 Strength saving throw\
    \ or be pushed up to 30 feet away in a straight line and be knocked [prone](/3-Mechanics/CLI/rules/conditions.md#prone)."
  "name": "Iron Fist"
- "desc": "Melee Weapon Attack: dice: d20+9 (+9) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d10 + 5|text(10) (1d10 + 5) bludgeoning damage, or dice:3d10\
    \ + 5|text(21) (3d10 + 5) to a [prone](/3-Mechanics/CLI/rules/conditions.md#prone)\
    \ target."
  "name": "Stomping Foot"
- "desc": "The duergar spews flames in a line 100 feet long and 5 feet wide. Each\
    \ creature in the line must make a DC 16 Dexterity saving throw, taking dice:4d8|text(18)\
    \ (4d8) fire damage on a failed save, or half as much damage on a successful\
    \ one."
  "name": "Flame Jet"
"source":
- "MPMM"
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Duergar%20Despot.webp"
```
^statblock

```encounter-table
name: Duergar Despot
creatures:
 - 1: Duergar Despot
```

## Environment

mountain, underdark