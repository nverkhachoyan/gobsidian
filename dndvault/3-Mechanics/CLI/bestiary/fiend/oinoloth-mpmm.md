---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/12
- monster/environment/desert
- monster/environment/underdark
- monster/size/medium
- monster/type/fiend/yugoloth
aliases: ["Oinoloth"]
NoteIcon: monster
BestiaryType: fiend (yugoloth)
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 202, Mordenkainen's Tome of Foes p. 251
---
# [Oinoloth](3-Mechanics\CLI\bestiary\fiend/oinoloth-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 202, Mordenkainen's Tome of Foes p. 251*  

```statblock
"name": "Oinoloth (MPMM)"
"size": "Medium"
"type": "fiend"
"subtype": "yugoloth"
"alignment": "Typically  Neutral Evil"
"ac": !!int "17"
"ac_class": "natural armor"
"hp": !!int "119"
"hit_dice": "14d8 + 56"
"stats":
- !!int "19"
- !!int "17"
- !!int "18"
- !!int "17"
- !!int "16"
- !!int "19"
"speed": "40 ft."
"saves":
  "Wisdom": !!int "7"
  "Constitution": !!int "8"
"skillsaves":
  "Intimidation": !!int "8"
  "Deception": !!int "8"
  "Perception": !!int "7"
"damage_resistances": "cold; fire; lightning; bludgeoning, piercing, slashing from\
  \ nonmagical attacks"
"damage_immunities": "acid, poison"
"condition_immunities": "[poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)"
"senses": "blindsight 60 ft., darkvision 60 ft., passive Perception 17"
"languages": "Abyssal, Infernal, telepathy 60 ft."
"cr": "12"
"traits":
- "desc": "The oinoloth casts one of the following spells, requiring no material components\
    \ and using Charisma as the spellcasting ability (spell save DC 16):\n\nAt will:\
    \ [darkness](/3-Mechanics/CLI/spells/darkness.md), [detect magic](/3-Mechanics/CLI/spells/detect-magic.md),\
    \ [dispel magic](/3-Mechanics/CLI/spells/dispel-magic.md), [hold monster](/3-Mechanics/CLI/spells/hold-monster.md),\
    \ [invisibility](/3-Mechanics/CLI/spells/invisibility.md) (self only)\n\n1/day\
    \ each: [feeblemind](/3-Mechanics/CLI/spells/feeblemind.md), [globe of invulnerability](/3-Mechanics/CLI/spells/globe-of-invulnerability.md)"
  "name": "Spellcasting"
- "desc": "The oinoloth has advantage on saving throws against spells and other magical\
    \ effects."
  "name": "Magic Resistance"
"actions":
- "desc": "The oinoloth makes two Claw attacks, and it uses Spellcasting or Teleport."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+8 (+8) to hit, reach 5 ft., one target.\
    \ Hit: dice:3d6 + 4|text(14) (3d6 + 4) slashing damage plus dice:4d10|text(22)\
    \ (4d10) necrotic damage."
  "name": "Claw"
- "desc": "The oinoloth touches one willing creature within 5 feet of it. The target\
    \ regains all its hit points. In addition, the oinoloth can end one disease on\
    \ the target or remove one of the following conditions from it: [blinded](/3-Mechanics/CLI/rules/conditions.md#blinded),\
    \ [deafened](/3-Mechanics/CLI/rules/conditions.md#deafened), [paralyzed](/3-Mechanics/CLI/rules/conditions.md#paralyzed),\
    \ or [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned). The target then\
    \ gains 1 level of [exhaustion](/3-Mechanics/CLI/rules/conditions.md#exhaustion),\
    \ and its hit point maximum is reduced by dice:2d6|text(7) (2d6). This reduction\
    \ can be removed only by a [wish](/3-Mechanics/CLI/spells/wish.md) spell or by\
    \ casting [greater restoration](/3-Mechanics/CLI/spells/greater-restoration.md)\
    \ on the target three times within the same hour. The target dies if its hit point\
    \ maximum is reduced to 0."
  "name": "Corrupted Healing (Recharge 6)"
- "desc": "The oinoloth teleports, along with any equipment it is wearing or carrying,\
    \ up to 60 feet to an unoccupied space it can see."
  "name": "Teleport"
"bonus_actions":
- "desc": "The oinoloth blights the area in a 30-foot-radius sphere centered on itself.\
    \ The blight lasts for 24 hours. While the area is blighted, all normal plants\
    \ there wither and die.\n\nFurthermore, when a creature moves into the blighted\
    \ area or starts its turn there, that creature must make a DC 16 Constitution\
    \ saving throw. On a failed save, the creature takes dice:4d6|text(14) (4d6)\
    \ poison damage and is [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned).\
    \ On a successful save, the creature is immune to the oinoloth's Bringer of Plagues\
    \ for the next 24 hours.\n\nThe [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)\
    \ creature can't regain hit points. After every 24 hours that elapse, the [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)\
    \ creature can repeat the saving throw. On a failed save, the creature's hit point\
    \ maximum is reduced by dice:1d10|text(5) (1d10). This reduction lasts until\
    \ the poison ends, and the target dies if its hit point maximum is reduced to\
    \ 0. The poison ends after the creature successfully saves against it three times."
  "name": "Bringer of Plagues (Recharge 5-6)"
"source":
- "MPMM"
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Oinoloth.webp"
```
^statblock

```encounter-table
name: Oinoloth
creatures:
 - 1: Oinoloth
```

## Environment

desert, underdark