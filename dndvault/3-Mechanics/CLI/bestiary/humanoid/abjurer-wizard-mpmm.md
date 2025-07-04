---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/9
- monster/environment/urban
- monster/size/medium
- monster/type/humanoid
aliases: ["Abjurer Wizard"]
NoteIcon: monster
BestiaryType: humanoid
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 260, Volo's Guide to Monsters p. 209
---
# [Abjurer Wizard](3-Mechanics\CLI\bestiary\humanoid/abjurer-wizard-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 260, Volo's Guide to Monsters p. 209*  

```statblock
"name": "Abjurer Wizard (MPMM)"
"size": "Medium"
"type": "humanoid"
"alignment": "Any alignment"
"ac": !!int "12"
"ac_class": "15 with [mage armor](/3-Mechanics/CLI/spells/mage-armor.md)"
"hp": !!int "104"
"hit_dice": "16d8 + 32"
"stats":
- !!int "9"
- !!int "14"
- !!int "14"
- !!int "18"
- !!int "12"
- !!int "11"
"speed": "30 ft."
"saves":
  "Wisdom": !!int "5"
  "Intelligence": !!int "8"
"skillsaves":
  "History": !!int "8"
  "Arcana": !!int "8"
"senses": "passive Perception 11"
"languages": "any four languages"
"cr": "9"
"traits":
- "desc": "The abjurer casts one of the following spells, using Intelligence as the\
    \ spellcasting ability (spell save DC 16):\n\nAt will: [dancing lights](/3-Mechanics/CLI/spells/dancing-lights.md),\
    \ [mage hand](/3-Mechanics/CLI/spells/mage-hand.md), [message](/3-Mechanics/CLI/spells/message.md),\
    \ [prestidigitation](/3-Mechanics/CLI/spells/prestidigitation.md)\n\n1/day each:\
    \ [arcane lock](/3-Mechanics/CLI/spells/arcane-lock.md), [banishment](/3-Mechanics/CLI/spells/banishment.md),\
    \ [globe of invulnerability](/3-Mechanics/CLI/spells/globe-of-invulnerability.md),\
    \ [invisibility](/3-Mechanics/CLI/spells/invisibility.md), [wall of force](/3-Mechanics/CLI/spells/wall-of-force.md)\n\
    \n2/day each: [dispel magic](/3-Mechanics/CLI/spells/dispel-magic.md), [lightning\
    \ bolt](/3-Mechanics/CLI/spells/lightning-bolt.md), [mage armor](/3-Mechanics/CLI/spells/mage-armor.md)"
  "name": "Spellcasting"
"actions":
- "desc": "The abjurer makes three Arcane Burst attacks."
  "name": "Multiattack"
- "desc": "Melee or Ranged Spell Attack: dice: d20+6 (+6) to hit, reach 5 ft.\
    \ or range 120 ft., one target. Hit: dice:3d10 + 4|text(20) (3d10 + 4) force\
    \ damage."
  "name": "Arcane Burst"
- "desc": "Each creature in a 20-foot cube originating from the abjurer must make\
    \ a DC 16 Constitution saving throw. On a failed save, a creature takes dice:8d8|text(36)\
    \ (8d8) force damage and is pushed up to 10 feet away from the abjurer. On a\
    \ successful save, a creature takes half as much damage and isn't pushed."
  "name": "Force Blast"
"reactions":
- "desc": "When the abjurer or a creature it can see within 30 feet of it takes damage,\
    \ the abjurer magically creates a protective barrier around itself or the other\
    \ creature. The barrier reduces the damage to the protected creature by dice:4d10\
    \ + 4|text(26) (4d10 + 4), to a minimum of 0, and then vanishes."
  "name": "Arcane Ward (Recharge 4-6)"
"source":
- "MPMM"
- "VGM"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Abjurer%20Wizard.webp"
```
^statblock

```encounter-table
name: Abjurer Wizard
creatures:
 - 1: Abjurer Wizard
```

## Environment

urban