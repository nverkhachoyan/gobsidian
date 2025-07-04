---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/18
- monster/environment/underdark
- monster/size/medium
- monster/type/humanoid/elf
- monster/type/humanoid/wizard
aliases: ["Drow Favored Consort"]
NoteIcon: monster
BestiaryType: humanoid (elf, wizard)
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 100, Mordenkainen's Tome of Foes p. 183
---
# [Drow Favored Consort](3-Mechanics\CLI\bestiary\humanoid/drow-favored-consort-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 100, Mordenkainen's Tome of Foes p. 183*  

```statblock
"name": "Drow Favored Consort (MPMM)"
"size": "Medium"
"type": "humanoid"
"subtype": "elf, wizard"
"alignment": "Any alignment"
"ac": !!int "15"
"hp": !!int "240"
"hit_dice": "32d8 + 96"
"stats":
- !!int "15"
- !!int "20"
- !!int "16"
- !!int "18"
- !!int "15"
- !!int "18"
"speed": "30 ft."
"saves":
  "Charisma": !!int "10"
  "Dexterity": !!int "11"
  "Constitution": !!int "9"
"skillsaves":
  "Athletics": !!int "8"
  "Stealth": !!int "11"
  "Perception": !!int "8"
  "Acrobatics": !!int "11"
"senses": "darkvision 120 ft., passive Perception 18"
"languages": "Elvish, Undercommon"
"cr": "18"
"traits":
- "desc": "The drow casts one of the following spells, requiring no material components\
    \ and using Intelligence as the spellcasting ability (spell save DC 18):\n\nAt\
    \ will: [dancing lights](/3-Mechanics/CLI/spells/dancing-lights.md), [mage armor](/3-Mechanics/CLI/spells/mage-armor.md),\
    \ [mage hand](/3-Mechanics/CLI/spells/mage-hand.md), [message](/3-Mechanics/CLI/spells/message.md)\n\
    \n1/day each: [darkness](/3-Mechanics/CLI/spells/darkness.md), [faerie fire](/3-Mechanics/CLI/spells/faerie-fire.md),\
    \ [levitate](/3-Mechanics/CLI/spells/levitate.md) (self only)\n\n3/day each:\
    \ [dimension door](/3-Mechanics/CLI/spells/dimension-door.md), [fireball](/3-Mechanics/CLI/spells/fireball.md),\
    \ [invisibility](/3-Mechanics/CLI/spells/invisibility.md)"
  "name": "Spellcasting"
- "desc": "The drow has advantage on saving throws against being [charmed](/3-Mechanics/CLI/rules/conditions.md#charmed),\
    \ and magic can't put the drow to sleep."
  "name": "Fey Ancestry"
- "desc": "While in sunlight, the drow has disadvantage on attack rolls, as well as\
    \ on Wisdom ([Perception](/3-Mechanics/CLI/rules/skills.md#Perception)) checks\
    \ that rely on sight."
  "name": "Sunlight Sensitivity"
"actions":
- "desc": "The drow makes three Scimitar or Arcane Eruption attacks. The drow can\
    \ replace one of the attacks with a use of Spellcasting."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+11 (+11) to hit, reach 5 ft., one\
    \ target. Hit: dice:1d6 + 5|text(8) (1d6 + 5) slashing damage plus dice:6d8|text(27)\
    \ (6d8) poison damage."
  "name": "Scimitar"
- "desc": "Ranged Spell Attack: dice: d20+10 (+10) to hit, range 120 ft., one\
    \ target. Hit: dice:8d8|text(36) (8d8) force damage, and the drow can push\
    \ the target up to 10 feet away if it is a Large or smaller creature."
  "name": "Arcane Eruption"
"reactions":
- "desc": "When the drow or a creature within 10 feet of it is hit by an attack roll,\
    \ the drow gives the target a +5 bonus to its AC until the start of the drow's\
    \ next turn, which can cause the triggering attack roll to miss."
  "name": "Protective Shield (3/Day)"
"source":
- "MPMM"
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Drow%20Favored%20Consort.webp"
```
^statblock

```encounter-table
name: Drow Favored Consort
creatures:
 - 1: Drow Favored Consort
```

## Environment

underdark