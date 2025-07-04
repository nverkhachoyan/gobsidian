---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/23
- monster/size/large
- monster/type/fiend/demon
aliases: ["Fraz-Urb'luu"]
NoteIcon: monster
BestiaryType: fiend (demon)
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 129, Mordenkainen's Tome of Foes p. 146
---
# [Fraz-Urb'luu](3-Mechanics\CLI\bestiary\npc/fraz-urbluu-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 129, Mordenkainen's Tome of Foes p. 146*  

```statblock
"name": "Fraz-Urb'luu (MPMM)"
"size": "Large"
"type": "fiend"
"subtype": "demon"
"alignment": "Chaotic Evil"
"ac": !!int "18"
"ac_class": "natural armor"
"hp": !!int "337"
"hit_dice": "27d10 + 189"
"stats":
- !!int "29"
- !!int "12"
- !!int "25"
- !!int "26"
- !!int "24"
- !!int "26"
"speed": "40 ft., fly 40 ft."
"saves":
  "Dexterity": !!int "8"
  "Wisdom": !!int "14"
  "Intelligence": !!int "15"
  "Constitution": !!int "14"
"skillsaves":
  "Deception": !!int "15"
  "Stealth": !!int "8"
  "Perception": !!int "14"
"damage_resistances": "cold, fire, lightning"
"damage_immunities": "poison; bludgeoning, piercing, slashing that is nonmagical"
"condition_immunities": "[charmed](/3-Mechanics/CLI/rules/conditions.md#charmed),\
  \ [exhaustion](/3-Mechanics/CLI/rules/conditions.md#exhaustion), [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened),\
  \ [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)"
"senses": "truesight 120 ft., passive Perception 24"
"languages": "all, telepathy 120 ft."
"cr": "23"
"traits":
- "desc": "Fraz-Urb'luu casts one of the following spells, requiring no material components\
    \ and using Charisma as the spellcasting ability (spell save DC 23):\n\nAt will:\
    \ [alter self](/3-Mechanics/CLI/spells/alter-self.md) (can become Medium when\
    \ changing his appearance), [detect magic](/3-Mechanics/CLI/spells/detect-magic.md),\
    \ [dispel magic](/3-Mechanics/CLI/spells/dispel-magic.md), [phantasmal force](/3-Mechanics/CLI/spells/phantasmal-force.md)\n\
    \n1/day each: [modify memory](/3-Mechanics/CLI/spells/modify-memory.md), [project\
    \ image](/3-Mechanics/CLI/spells/project-image.md)\n\n3/day each: [mislead](/3-Mechanics/CLI/spells/mislead.md),\
    \ [programmed illusion](/3-Mechanics/CLI/spells/programmed-illusion.md), [seeming](/3-Mechanics/CLI/spells/seeming.md)"
  "name": "Spellcasting"
- "desc": "If Fraz-Urb'luu fails a saving throw, he can choose to succeed instead."
  "name": "Legendary Resistance (3/Day)"
- "desc": "Fraz-Urb'luu has advantage on saving throws against spells and other magical\
    \ effects."
  "name": "Magic Resistance"
- "desc": "Fraz-Urb'luu can't be targeted by divination magic, perceived through magical\
    \ scrying sensors, or detected by abilities that sense demons or Fiends."
  "name": "Undetectable"
"actions":
- "desc": "Fraz-Urb'luu makes one Bite attack and two Fist attacks, and he uses Phantasmal\
    \ Terror."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+16 (+16) to hit, reach 10 ft., one\
    \ target. Hit: dice:3d6 + 9|text(19) (3d6 + 9) force damage."
  "name": "Bite"
- "desc": "Melee Weapon Attack: dice: d20+16 (+16) to hit, reach 10 ft., one\
    \ target. Hit: dice:3d8 + 9|text(22) (3d8 + 9) force damage."
  "name": "Fist"
- "desc": "Fraz-Urb'luu targets one creature he can see within 120 feet of him. The\
    \ target must succeed on a DC 23 Wisdom saving throw, or it takes dice:3d10|text(16)\
    \ (3d10) psychic damage and is [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened)\
    \ of Fraz-Urb'luu until the end of its next turn."
  "name": "Phantasmal Terror"
"legendary_actions":
- "desc": "Melee Weapon Attack: dice: d20+16 (+16) to hit, reach 15 ft., one\
    \ target. Hit: dice:2d10 + 9|text(20) (2d10 + 9) force damage. If the target\
    \ is a Large or smaller creature, it is also [grappled](/3-Mechanics/CLI/rules/conditions.md#grappled)\
    \ (escape DC 24), and it is [restrained](/3-Mechanics/CLI/rules/conditions.md#restrained)\
    \ until the grapple ends. Fraz-Urb'luu can grapple only one creature with his\
    \ tail at a time."
  "name": "Tail"
- "desc": "Fraz-Urb'luu uses Phantasmal Terror."
  "name": "Terror (Costs 2 Actions)"
"source":
- "MPMM"
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Fraz-Urb%27luu.webp"
```
^statblock

```encounter-table
name: Fraz-Urb'luu
creatures:
 - 1: Fraz-Urb'luu
```