---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/3
- monster/environment/urban
- monster/size/medium
- monster/type/humanoid
aliases: ["Illusionist Wizard"]
NoteIcon: monster
BestiaryType: humanoid
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 263, Volo's Guide to Monsters p. 214
---
# [Illusionist Wizard](3-Mechanics\CLI\bestiary\humanoid/illusionist-wizard-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 263, Volo's Guide to Monsters p. 214*  

```statblock
"name": "Illusionist Wizard (MPMM)"
"size": "Medium"
"type": "humanoid"
"alignment": "Any alignment"
"ac": !!int "12"
"ac_class": "15 with [mage armor](/3-Mechanics/CLI/spells/mage-armor.md)"
"hp": !!int "44"
"hit_dice": "8d8 + 8"
"stats":
- !!int "9"
- !!int "14"
- !!int "13"
- !!int "16"
- !!int "11"
- !!int "12"
"speed": "30 ft."
"saves":
  "Wisdom": !!int "2"
  "Intelligence": !!int "5"
"skillsaves":
  "History": !!int "5"
  "Arcana": !!int "5"
"senses": "passive Perception 10"
"languages": "any four languages"
"cr": "3"
"traits":
- "desc": "The illusionist casts one of the following spells, using Intelligence as\
    \ the spellcasting ability (spell save DC 13):\n\nAt will: [dancing lights](/3-Mechanics/CLI/spells/dancing-lights.md),\
    \ [mage hand](/3-Mechanics/CLI/spells/mage-hand.md), [minor illusion](/3-Mechanics/CLI/spells/minor-illusion.md)\n\
    \n2/day each: [disguise self](/3-Mechanics/CLI/spells/disguise-self.md), [invisibility](/3-Mechanics/CLI/spells/invisibility.md),\
    \ [mage armor](/3-Mechanics/CLI/spells/mage-armor.md), [major image](/3-Mechanics/CLI/spells/major-image.md),\
    \ [phantasmal force](/3-Mechanics/CLI/spells/phantasmal-force.md), [phantom steed](/3-Mechanics/CLI/spells/phantom-steed.md)"
  "name": "Spellcasting"
"actions":
- "desc": "The illusionist makes two Arcane Burst attacks."
  "name": "Multiattack"
- "desc": "Melee or Ranged Spell Attack: dice: d20+5 (+5) to hit, reach 5 ft.\
    \ or range 120 ft., one target. Hit: dice:2d10 + 3|text(14) (2d10 + 3) psychic\
    \ damage."
  "name": "Arcane Burst"
"bonus_actions":
- "desc": "The illusionist projects an illusion that makes the illusionist appear\
    \ to be standing in a place a few inches from its actual location, causing any\
    \ creature to have disadvantage on attack rolls against the illusionist. The effect\
    \ lasts for 1 minute, and it ends early if the illusionist takes damage, if it\
    \ is [incapacitated](/3-Mechanics/CLI/rules/conditions.md#incapacitated), or if\
    \ its speed becomes 0."
  "name": "Displacement (Recharge 5-6)"
"source":
- "MPMM"
- "VGM"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Illusionist%20Wizard.webp"
```
^statblock

```encounter-table
name: Illusionist Wizard
creatures:
 - 1: Illusionist Wizard
```

## Environment

urban