---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/1
- monster/environment/forest
- monster/environment/hill
- monster/environment/mountain
- monster/environment/underdark
- monster/environment/urban
- monster/size/small
- monster/type/humanoid
aliases: ["Kobold Scale Sorcerer"]
NoteIcon: monster
BestiaryType: humanoid
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 165, Volo's Guide to Monsters p. 167
---
# [Kobold Scale Sorcerer](3-Mechanics\CLI\bestiary\humanoid/kobold-scale-sorcerer-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 165, Volo's Guide to Monsters p. 167*  

```statblock
"name": "Kobold Scale Sorcerer (MPMM)"
"size": "Small"
"type": "humanoid"
"alignment": "Any alignment"
"ac": !!int "15"
"ac_class": "natural armor"
"hp": !!int "27"
"hit_dice": "5d6 + 10"
"stats":
- !!int "7"
- !!int "15"
- !!int "14"
- !!int "10"
- !!int "9"
- !!int "14"
"speed": "30 ft."
"skillsaves":
  "Medicine": !!int "1"
  "Arcana": !!int "2"
"senses": "darkvision 60 ft., passive Perception 9"
"languages": "Common, Draconic"
"cr": "1"
"traits":
- "desc": "The kobold casts one of the following spells, requiring no material components\
    \ and using Charisma as the spellcasting ability (spell save DC 12):\n\nAt will:\
    \ [mage hand](/3-Mechanics/CLI/spells/mage-hand.md), [prestidigitation](/3-Mechanics/CLI/spells/prestidigitation.md)\n\
    \n2/day each: [charm person](/3-Mechanics/CLI/spells/charm-person.md), [fog\
    \ cloud](/3-Mechanics/CLI/spells/fog-cloud.md), [levitate](/3-Mechanics/CLI/spells/levitate.md)"
  "name": "Spellcasting"
- "desc": "The kobold has advantage on an attack roll against a creature if at least\
    \ one of the kobold's allies is within 5 feet of the creature and the ally isn't\
    \ [incapacitated](/3-Mechanics/CLI/rules/conditions.md#incapacitated)."
  "name": "Pack Tactics"
- "desc": "While in sunlight, the kobold has disadvantage on attack rolls, as well\
    \ as on Wisdom ([Perception](/3-Mechanics/CLI/rules/skills.md#Perception)) checks\
    \ that rely on sight."
  "name": "Sunlight Sensitivity"
"actions":
- "desc": "The kobold makes two Dagger or Chromatic Bolt attacks. It can replace one\
    \ attack with a use of Spellcasting."
  "name": "Multiattack"
- "desc": "Melee or Ranged Weapon Attack: dice: d20+4 (+4) to hit, reach 5 ft.\
    \ or range 20/60 ft., one target. Hit: dice:1d4 + 2|text(4) (1d4 + 2) piercing\
    \ damage."
  "name": "Dagger"
- "desc": "Ranged Spell Attack: dice: d20+4 (+4) to hit, range 60 feet, one\
    \ target. Hit: dice:2d6 + 2|text(9) (2d6 + 2) of a type of the kobold's\
    \ choice: acid, cold, fire, lightning, poison, or thunder."
  "name": "Chromatic Bolt"
"source":
- "MPMM"
- "VGM"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Kobold%20Scale%20Sorcerer.webp"
```
^statblock

```encounter-table
name: Kobold Scale Sorcerer
creatures:
 - 1: Kobold Scale Sorcerer
```

## Environment

forest, hill, mountain, underdark, urban