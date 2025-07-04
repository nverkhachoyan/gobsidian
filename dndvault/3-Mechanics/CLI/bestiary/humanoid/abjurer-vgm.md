---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/vgm
- monster/cr/9
- monster/environment/urban
- monster/size/medium
- monster/type/humanoid/any-race
aliases: ["Abjurer"]
NoteIcon: monster
BestiaryType: humanoid (any race)
SourceType: Bestiary
BookSource: Volo's Guide to Monsters p. 209
---
# [Abjurer](3-Mechanics\CLI\bestiary\humanoid/abjurer-vgm.md)
*Source: Volo's Guide to Monsters p. 209*  

Abjurers are specialist wizards who feel secure when warded by layers of magical power. Kings, nobles, and other wealthy individuals commonly hire abjurers to cast protective spells on their homes and vaults.

```statblock
"name": "Abjurer (VGM)"
"size": "Medium"
"type": "humanoid"
"subtype": "any race"
"alignment": "Any alignment"
"ac": !!int "12"
"ac_class": "15 with [mage armor](/3-Mechanics/CLI/spells/mage-armor.md)"
"hp": !!int "84"
"hit_dice": "13d8 + 26"
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
- "desc": "The abjurer is a 13th-level spellcaster. Its spellcasting ability is Intelligence\
    \ (spell save DC 16, dice: d20+8 (+8) to hit with spell attacks). The abjurer\
    \ has the following wizard spells prepared:\n\nCantrips (at will): [blade\
    \ ward](/3-Mechanics/CLI/spells/blade-ward.md), [dancing lights](/3-Mechanics/CLI/spells/dancing-lights.md),\
    \ [mending](/3-Mechanics/CLI/spells/mending.md), [message](/3-Mechanics/CLI/spells/message.md),\
    \ [ray of frost](/3-Mechanics/CLI/spells/ray-of-frost.md)\n\n1st level (4 slots):\
    \ [alarm](/3-Mechanics/CLI/spells/alarm.md), [mage armor](/3-Mechanics/CLI/spells/mage-armor.md),\
    \ [magic missile](/3-Mechanics/CLI/spells/magic-missile.md), [shield](/3-Mechanics/CLI/spells/shield.md)\n\
    \n2nd level (3 slots): [arcane lock](/3-Mechanics/CLI/spells/arcane-lock.md),\
    \ [invisibility](/3-Mechanics/CLI/spells/invisibility.md)\n\n3rd level (3 slots):\
    \ [counterspell](/3-Mechanics/CLI/spells/counterspell.md), [dispel magic](/3-Mechanics/CLI/spells/dispel-magic.md),\
    \ [fireball](/3-Mechanics/CLI/spells/fireball.md)\n\n4th level (3 slots):\
    \ [banishment](/3-Mechanics/CLI/spells/banishment.md), [stoneskin](/3-Mechanics/CLI/spells/stoneskin.md)\n\
    \n5th level (2 slots): [cone of cold](/3-Mechanics/CLI/spells/cone-of-cold.md),\
    \ [wall of force](/3-Mechanics/CLI/spells/wall-of-force.md)\n\n6th level (1\
    \ slots): [flesh to stone](/3-Mechanics/CLI/spells/flesh-to-stone.md), [globe\
    \ of invulnerability](/3-Mechanics/CLI/spells/globe-of-invulnerability.md)\n\n\
    7th level (1 slots): [symbol](/3-Mechanics/CLI/spells/symbol.md), [teleport](/3-Mechanics/CLI/spells/teleport.md)\n\
    \nAbjuration spell of 1st level or higher"
  "name": "Spellcasting"
- "desc": "The abjurer has a magical ward that has 30 hit points. Whenever the abjurer\
    \ takes damage, the ward takes the damage instead. If the ward is reduced to 0\
    \ hit points, the abjurer takes any remaining damage. When the abjurer casts an\
    \ abjuration spell of 1st level or higher, the ward regains a number of hit points\
    \ equal to twice the level of the spell."
  "name": "Arcane Ward"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+3 (+3) to hit, reach 5 ft., one target.\
    \ Hit:* dice:1d6 - 1|text(2) (1d6 - 1) bludgeoning damage, or dice:1d8 -\
    \ 1|text(3) (1d8 - 1) bludgeoning damage if used with two hands."
  "name": "Quarterstaff"
"source":
- "VGM"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/VGM/Abjurer.webp"
```
^statblock

```encounter-table
name: Abjurer
creatures:
 - 1: Abjurer
```

## Environment

urban