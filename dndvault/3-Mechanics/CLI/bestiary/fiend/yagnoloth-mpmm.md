---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/11
- monster/environment/urban
- monster/size/large
- monster/type/fiend/yugoloth
aliases: ["Yagnoloth"]
NoteIcon: monster
BestiaryType: fiend (yugoloth)
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 268, Mordenkainen's Tome of Foes p. 252
---
# [Yagnoloth](3-Mechanics\CLI\bestiary\fiend/yagnoloth-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 268, Mordenkainen's Tome of Foes p. 252*  

```statblock
"name": "Yagnoloth (MPMM)"
"size": "Large"
"type": "fiend"
"subtype": "yugoloth"
"alignment": "Typically  Neutral Evil"
"ac": !!int "17"
"ac_class": "natural armor"
"hp": !!int "147"
"hit_dice": "14d10 + 70"
"stats":
- !!int "19"
- !!int "14"
- !!int "21"
- !!int "16"
- !!int "15"
- !!int "18"
"speed": "40 ft."
"saves":
  "Charisma": !!int "8"
  "Dexterity": !!int "6"
  "Wisdom": !!int "6"
  "Intelligence": !!int "7"
"skillsaves":
  "Deception": !!int "8"
  "Insight": !!int "6"
  "Perception": !!int "6"
  "Persuasion": !!int "8"
"damage_resistances": "cold; fire; lightning; bludgeoning, piercing, slashing from\
  \ nonmagical attacks"
"damage_immunities": "acid, poison"
"condition_immunities": "[poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)"
"senses": "blindsight 60 ft., darkvision 60 ft., passive Perception 16"
"languages": "Abyssal, Infernal, telepathy 60 ft."
"cr": "11"
"traits":
- "desc": "The yagnoloth casts one of the following spells, requiring no material\
    \ components and using Charisma as the spellcasting ability (spell save DC 16):\n\
    \nAt will: [darkness](/3-Mechanics/CLI/spells/darkness.md), [detect magic](/3-Mechanics/CLI/spells/detect-magic.md),\
    \ [dispel magic](/3-Mechanics/CLI/spells/dispel-magic.md), [invisibility](/3-Mechanics/CLI/spells/invisibility.md)\
    \ (self only), [suggestion](/3-Mechanics/CLI/spells/suggestion.md)\n\n3/day:\
    \ [lightning bolt](/3-Mechanics/CLI/spells/lightning-bolt.md)"
  "name": "Spellcasting"
- "desc": "The yagnoloth has advantage on saving throws against spells and other magical\
    \ effects."
  "name": "Magic Resistance"
"actions":
- "desc": "The yagnoloth makes one Electrified Touch attack and one Massive Arm attack,\
    \ or it makes one Massive Arm attack and uses Battlefield Cunning, if available,\
    \ or Teleport."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+8 (+8) to hit, reach 5 ft., one target.\
    \ Hit: dice:6d8|text(27) (6d8) lightning damage."
  "name": "Electrified Touch"
- "desc": "Melee Weapon Attack: dice: d20+8 (+8) to hit, reach 15 ft., one target.\
    \ Hit: dice:3d12 + 4|text(23) (3d12 + 4) force damage. If the target is\
    \ a creature, it must succeed on a DC 16 Constitution saving throw or become [stunned](/3-Mechanics/CLI/rules/conditions.md#stunned)\
    \ until the end of the yagnoloth's next turn."
  "name": "Massive Arm"
- "desc": "Up to two allied yugoloths within 60 feet of the yagnoloth that can hear\
    \ it can use their reactions to make one melee attack each."
  "name": "Battlefield Cunning (Recharge 4-6)"
- "desc": "The yagnoloth touches one [incapacitated](/3-Mechanics/CLI/rules/conditions.md#incapacitated)\
    \ creature within 15 feet of it. The target takes dice:7d8 + 4|text(36) (7d8\
    \ + 4) necrotic damage, and the yagnoloth gains temporary hit points equal to\
    \ half the damage dealt. The target must succeed on a DC 16 Constitution saving\
    \ throw, or its hit point maximum is reduced by an amount equal to half the necrotic\
    \ damage taken. This reduction lasts until the target finishes a long rest, and\
    \ the target dies if its hit point maximum is reduced to 0."
  "name": "Life Leech"
- "desc": "The yagnoloth teleports, along with any equipment it is wearing or carrying,\
    \ up to 60 feet to an unoccupied space it can see."
  "name": "Teleport"
"source":
- "MPMM"
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Yagnoloth.webp"
```
^statblock

```encounter-table
name: Yagnoloth
creatures:
 - 1: Yagnoloth
```

## Environment

urban