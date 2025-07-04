---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/3
- monster/environment/underdark
- monster/size/small
- monster/type/aberration/sorcerer
aliases: ["Derro Savant"]
NoteIcon: monster
BestiaryType: aberration (sorcerer)
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 92, Mordenkainen's Tome of Foes p. 159
---
# [Derro Savant](3-Mechanics\CLI\bestiary\aberration/derro-savant-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 92, Mordenkainen's Tome of Foes p. 159*  

```statblock
"name": "Derro Savant (MPMM)"
"size": "Small"
"type": "aberration"
"subtype": "sorcerer"
"alignment": "Typically  Chaotic Evil"
"ac": !!int "13"
"ac_class": "[leather armor](/3-Mechanics/CLI/items/leather-armor.md)"
"hp": !!int "36"
"hit_dice": "8d6 + 8"
"stats":
- !!int "9"
- !!int "14"
- !!int "12"
- !!int "11"
- !!int "5"
- !!int "14"
"speed": "30 ft."
"skillsaves":
  "Stealth": !!int "4"
"senses": "darkvision 120 ft., passive Perception 7"
"languages": "Dwarvish, Undercommon"
"cr": "3"
"traits":
- "desc": "The derro casts one of the following spells, using Charisma as the spellcasting\
    \ ability (spell save DC 12):\n\nAt will: [mage hand](/3-Mechanics/CLI/spells/mage-hand.md),\
    \ [message](/3-Mechanics/CLI/spells/message.md), [prestidigitation](/3-Mechanics/CLI/spells/prestidigitation.md)\n\
    \n1/day each: [invisibility](/3-Mechanics/CLI/spells/invisibility.md), [sleep](/3-Mechanics/CLI/spells/sleep.md),\
    \ [spider climb](/3-Mechanics/CLI/spells/spider-climb.md)"
  "name": "Spellcasting"
- "desc": "The derro has advantage on saving throws against spells and other magical\
    \ effects."
  "name": "Magic Resistance"
- "desc": "While in sunlight, the derro has disadvantage on attack rolls, as well\
    \ as on Wisdom ([Perception](/3-Mechanics/CLI/rules/skills.md#Perception)) checks\
    \ that rely on sight."
  "name": "Sunlight Sensitivity"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+1 (+1) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d6 - 1|text(2) (1d6 - 1) bludgeoning damage."
  "name": "Quarterstaff"
- "desc": "The derro launches a brilliant beam of magical energy in a 5-foot-wide\
    \ line that is 60 feet long. Each creature in the line must make a DC 12 Dexterity\
    \ saving throw, taking dice:6d6|text(21) (6d6) radiant damage on a failed\
    \ save, or half as much damage on a successful one."
  "name": "Chromatic Beam"
"source":
- "MPMM"
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Derro%20Savant.webp"
```
^statblock

```encounter-table
name: Derro Savant
creatures:
 - 1: Derro Savant
```

## Environment

underdark