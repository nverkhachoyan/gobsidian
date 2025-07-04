---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/17
- monster/environment/coastal
- monster/environment/urban
- monster/size/medium
- monster/type/fiend/devil
- monster/type/fiend/wizard
aliases: ["Blue Abishai"]
NoteIcon: monster
BestiaryType: fiend (devil, wizard)
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 39, Mordenkainen's Tome of Foes p. 161
---
# [Blue Abishai](3-Mechanics\CLI\bestiary\fiend/blue-abishai-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 39, Mordenkainen's Tome of Foes p. 161*  

```statblock
"name": "Blue Abishai (MPMM)"
"size": "Medium"
"type": "fiend"
"subtype": "devil, wizard"
"alignment": "Typically  Lawful Evil"
"ac": !!int "19"
"ac_class": "natural armor"
"hp": !!int "202"
"hit_dice": "27d8 + 81"
"stats":
- !!int "15"
- !!int "14"
- !!int "17"
- !!int "22"
- !!int "23"
- !!int "18"
"speed": "30 ft., fly 50 ft."
"saves":
  "Wisdom": !!int "12"
  "Intelligence": !!int "12"
"skillsaves":
  "Arcana": !!int "12"
"damage_resistances": "cold; bludgeoning, piercing, slashing from nonmagical attacks\
  \ that aren't silvered"
"damage_immunities": "fire, lightning, poison"
"condition_immunities": "[poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)"
"senses": "darkvision 120 ft., passive Perception 16"
"languages": "Draconic, Infernal, telepathy 120 ft."
"cr": "17"
"traits":
- "desc": "The abishai casts one of the following spells, using Intelligence as the\
    \ spellcasting ability (spell save DC 20):\n\nAt will: [disguise self](/3-Mechanics/CLI/spells/disguise-self.md),\
    \ [mage hand](/3-Mechanics/CLI/spells/mage-hand.md), [minor illusion](/3-Mechanics/CLI/spells/minor-illusion.md)\n\
    \n2/day each: [charm person](/3-Mechanics/CLI/spells/charm-person.md), [dispel\
    \ magic](/3-Mechanics/CLI/spells/dispel-magic.md), [greater invisibility](/3-Mechanics/CLI/spells/greater-invisibility.md),\
    \ [wall of force](/3-Mechanics/CLI/spells/wall-of-force.md)"
  "name": "Spellcasting"
- "desc": "Magical darkness doesn't impede the abishai's [darkvision](/3-Mechanics/CLI/rules/senses.md#darkvision)."
  "name": "Devil's Sight"
- "desc": "The abishai has advantage on saving throws against spells and other magical\
    \ effects."
  "name": "Magic Resistance"
"actions":
- "desc": "The abishai makes three Bite or Lightning Strike attacks."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+8 (+8) to hit, reach 5 ft., one target.\
    \ Hit: dice:2d10 + 2|text(13) (2d10 + 2) piercing damage plus dice:4d6|text(14)\
    \ (4d6) lightning damage."
  "name": "Bite"
- "desc": "Ranged Spell Attack: dice: d20+12 (+12) to hit, range 120 ft., one\
    \ target. Hit: dice:8d8|text(36) (8d8) lightning damage."
  "name": "Lightning Strike"
"bonus_actions":
- "desc": "The abishai teleports, along with any equipment it is wearing or carrying,\
    \ up to 30 feet to an unoccupied space that it can see."
  "name": "Teleport"
"source":
- "MPMM"
- "MTF"
- "VEoR"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Blue%20Abishai.webp"
```
^statblock

```encounter-table
name: Blue Abishai
creatures:
 - 1: Blue Abishai
```

## Environment

coastal, urban