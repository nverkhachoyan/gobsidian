---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/7
- monster/environment/underdark
- monster/size/medium
- monster/type/fiend/yugoloth
aliases: ["Dhergoloth"]
NoteIcon: monster
BestiaryType: fiend (yugoloth)
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 94, Mordenkainen's Tome of Foes p. 248
---
# [Dhergoloth](3-Mechanics\CLI\bestiary\fiend/dhergoloth-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 94, Mordenkainen's Tome of Foes p. 248*  

```statblock
"name": "Dhergoloth (MPMM)"
"size": "Medium"
"type": "fiend"
"subtype": "yugoloth"
"alignment": "Typically  Neutral Evil"
"ac": !!int "15"
"ac_class": "natural armor"
"hp": !!int "119"
"hit_dice": "14d8 + 56"
"stats":
- !!int "17"
- !!int "10"
- !!int "19"
- !!int "7"
- !!int "10"
- !!int "9"
"speed": "30 ft."
"saves":
  "Strength": !!int "6"
"damage_resistances": "cold; fire; lightning; bludgeoning, piercing, slashing from\
  \ nonmagical attacks"
"damage_immunities": "acid, poison"
"condition_immunities": "[poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)"
"senses": "blindsight 60 ft., darkvision 60 ft., passive Perception 10"
"languages": "Abyssal, Infernal, telepathy 60 ft."
"cr": "7"
"traits":
- "desc": "The dhergoloth casts one of the following spells, requiring no material\
    \ components and using Charisma as the spellcasting ability (spell save DC 10):\n\
    \nAt will: [darkness](/3-Mechanics/CLI/spells/darkness.md), [fear](/3-Mechanics/CLI/spells/fear.md)"
  "name": "Spellcasting"
- "desc": "The dhergoloth has advantage on saving throws against spells and other\
    \ magical effects."
  "name": "Magic Resistance"
"actions":
- "desc": "The dhergoloth makes two Claw attacks."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+6 (+6) to hit, reach 5 ft., one target.\
    \ Hit: dice:2d8 + 3|text(12) (2d8 + 3) force damage."
  "name": "Claw"
- "desc": "The dhergoloth moves up to its speed in a straight line and targets each\
    \ creature within 5 feet of it during its movement. Each target must succeed on\
    \ a DC 14 Dexterity saving throw or take dice:3d12 + 3|text(22) (3d12 + 3)\
    \ force damage."
  "name": "Flailing Claws (Recharge 5-6)"
- "desc": "The dhergoloth teleports, along with any equipment it is wearing or carrying,\
    \ up to 60 feet to an unoccupied space it can see."
  "name": "Teleport"
"source":
- "MPMM"
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Dhergoloth.webp"
```
^statblock

```encounter-table
name: Dhergoloth
creatures:
 - 1: Dhergoloth
```

## Environment

underdark