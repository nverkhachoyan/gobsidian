---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/18
- monster/environment/underdark
- monster/size/huge
- monster/type/fiend/demon
aliases: ["Sibriex"]
NoteIcon: monster
BestiaryType: fiend (demon)
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 217, Mordenkainen's Tome of Foes p. 137
---
# [Sibriex](3-Mechanics\CLI\bestiary\fiend/sibriex-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 217, Mordenkainen's Tome of Foes p. 137*  

```statblock
"name": "Sibriex (MPMM)"
"size": "Huge"
"type": "fiend"
"subtype": "demon"
"alignment": "Typically  Chaotic Evil"
"ac": !!int "19"
"ac_class": "natural armor"
"hp": !!int "150"
"hit_dice": "12d12 + 72"
"stats":
- !!int "10"
- !!int "3"
- !!int "23"
- !!int "25"
- !!int "24"
- !!int "25"
"speed": "0 ft., fly 20 ft. (hover)"
"saves":
  "Charisma": !!int "13"
  "Intelligence": !!int "13"
"skillsaves":
  "Perception": !!int "13"
  "History": !!int "13"
  "Arcana": !!int "13"
"damage_resistances": "cold; fire; lightning; bludgeoning, piercing, slashing from\
  \ nonmagical attacks"
"damage_immunities": "poison"
"condition_immunities": "[poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)"
"senses": "truesight 120 ft., passive Perception 23"
"languages": "all, telepathy 120 ft."
"cr": "18"
"traits":
- "desc": "The sibriex casts one of the following spells, requiring no material components\
    \ and using Charisma as the spellcasting ability (spell save DC 21):\n\nAt will:\
    \ [command](/3-Mechanics/CLI/spells/command.md), [dispel magic](/3-Mechanics/CLI/spells/dispel-magic.md),\
    \ [hold monster](/3-Mechanics/CLI/spells/hold-monster.md)\n\n1/day: [feeblemind](/3-Mechanics/CLI/spells/feeblemind.md)"
  "name": "Spellcasting"
- "desc": "The sibriex emits an aura of corruption 30 feet in every direction. Vegetation\
    \ withers in the aura, and the ground in the aura is difficult terrain for other\
    \ creatures. Any creature that starts its turn in the aura must succeed on a DC\
    \ 20 Constitution saving throw or take dice:4d6|text(14) (4d6) poison damage.\
    \ A creature that succeeds on the save is immune to this sibriex's Contamination\
    \ for 24 hours."
  "name": "Contamination"
- "desc": "If the sibriex fails a saving throw, it can choose to succeed instead."
  "name": "Legendary Resistance (3/Day)"
- "desc": "The sibriex has advantage on saving throws against spells and other magical\
    \ effects."
  "name": "Magic Resistance"
"actions":
- "desc": "The sibriex makes three Chain attacks, and it uses Squirt Bile."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+13 (+13) to hit, reach 15 ft., one\
    \ target. Hit: dice:2d12 + 7|text(20) (2d12 + 7) force damage."
  "name": "Chain"
- "desc": "The sibriex targets one creature it can see within 120 feet of it. The\
    \ target must succeed on a DC 20 Dexterity saving throw or take dice:9d6|text(31)\
    \ (9d6) acid damage."
  "name": "Squirt Bile"
- "desc": "The sibriex targets up to three creatures it can see within 120 feet of\
    \ it. Each target must make a DC 20 Constitution saving throw. On a successful\
    \ save, a creature becomes immune to this sibriex's Warp Creature. On a failed\
    \ save, the target is [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned),\
    \ which causes it to also gain 1 level of [exhaustion](/3-Mechanics/CLI/rules/conditions.md#exhaustion).\
    \ While [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned) in this way,\
    \ the target must repeat the saving throw at the start of each of its turns. Three\
    \ successful saves against the poison end it, and ending the poison removes any\
    \ levels of [exhaustion](/3-Mechanics/CLI/rules/conditions.md#exhaustion) caused\
    \ by it. Each failed save causes the target to gain another level of [exhaustion](/3-Mechanics/CLI/rules/conditions.md#exhaustion).\
    \ Once the target reaches 6 levels of [exhaustion](/3-Mechanics/CLI/rules/conditions.md#exhaustion),\
    \ it dies and instantly transforms into a living [manes](/3-Mechanics/CLI/bestiary/fiend/manes.md)\
    \ under the sibriex's control. The transformation of the body can be undone only\
    \ by a [wish](/3-Mechanics/CLI/spells/wish.md) spell."
  "name": "Warp Creature"
"legendary_actions":
- "desc": "The sibriex uses Spellcasting."
  "name": "Cast a Spell"
- "desc": "The sibriex uses Squirt Bile."
  "name": "Spray Bile"
- "desc": "The sibriex uses Warp Creature."
  "name": "Warp (Costs 2 Actions)"
"source":
- "MPMM"
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Sibriex.webp"
```
^statblock

```encounter-table
name: Sibriex
creatures:
 - 1: Sibriex
```

## Environment

underdark