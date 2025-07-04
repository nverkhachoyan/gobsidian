---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/7
- monster/environment/urban
- monster/size/medium
- monster/type/fiend/devil
aliases: ["Black Abishai"]
NoteIcon: monster
BestiaryType: fiend (devil)
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 38, Mordenkainen's Tome of Foes p. 160
---
# [Black Abishai](3-Mechanics\CLI\bestiary\fiend/black-abishai-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 38, Mordenkainen's Tome of Foes p. 160*  

```statblock
"name": "Black Abishai (MPMM)"
"size": "Medium"
"type": "fiend"
"subtype": "devil"
"alignment": "Typically  Lawful Evil"
"ac": !!int "15"
"ac_class": "natural armor"
"hp": !!int "58"
"hit_dice": "9d8 + 18"
"stats":
- !!int "14"
- !!int "17"
- !!int "14"
- !!int "13"
- !!int "16"
- !!int "11"
"speed": "30 ft., fly 40 ft."
"saves":
  "Dexterity": !!int "6"
  "Wisdom": !!int "6"
"skillsaves":
  "Stealth": !!int "6"
  "Perception": !!int "6"
"damage_resistances": "cold; bludgeoning, piercing, slashing from nonmagical attacks\
  \ that aren't silvered"
"damage_immunities": "acid, fire, poison"
"condition_immunities": "[poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)"
"senses": "darkvision 120 ft., passive Perception 16"
"languages": "Draconic, Infernal, telepathy 120 ft."
"cr": "7"
"traits":
- "desc": "The abishai casts [darkness](/3-Mechanics/CLI/spells/darkness.md) at a\
    \ point within 120 feet of it, requiring no spell components or [concentration](/3-Mechanics/CLI/rules/conditions.md#concentration).\
    \ Wisdom is its spellcasting ability for this spell. While the spell persists,\
    \ the abishai can move the area of darkness up to 60 feet as a bonus action.\n"
  "name": "Creeping Darkness (Recharge 6)"
- "desc": "Magical darkness doesn't impede the abishai's [darkvision](/3-Mechanics/CLI/rules/senses.md#darkvision)."
  "name": "Devil's Sight"
- "desc": "The abishai has advantage on saving throws against spells and other magical\
    \ effects."
  "name": "Magic Resistance"
"actions":
- "desc": "The abishai makes one Bite attack and two Scimitar attacks."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+6 (+6) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d10 + 3|text(8) (1d10 + 3) piercing damage plus dice:2d8|text(9)\
    \ (2d8) acid damage."
  "name": "Bite"
- "desc": "Melee Weapon Attack: dice: d20+6 (+6) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d6 + 3|text(6) (1d6 + 3) force damage."
  "name": "Scimitar"
"bonus_actions":
- "desc": "While in dim light or darkness, the abishai takes the [Hide](/3-Mechanics/CLI/rules/actions.md#Hide)\
    \ action."
  "name": "Shadow Stealth"
"source":
- "MPMM"
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Black%20Abishai.webp"
```
^statblock

```encounter-table
name: Black Abishai
creatures:
 - 1: Black Abishai
```

## Environment

urban