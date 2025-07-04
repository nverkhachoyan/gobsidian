---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/16
- monster/environment/urban
- monster/size/large
- monster/type/construct
aliases: ["Steel Predator"]
NoteIcon: monster
BestiaryType: construct
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 232, Mordenkainen's Tome of Foes p. 239
---
# [Steel Predator](3-Mechanics\CLI\bestiary\construct/steel-predator-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 232, Mordenkainen's Tome of Foes p. 239*  

```statblock
"name": "Steel Predator (MPMM)"
"size": "Large"
"type": "construct"
"alignment": "Typically  Lawful Neutral"
"ac": !!int "20"
"ac_class": "natural armor"
"hp": !!int "207"
"hit_dice": "18d10 + 108"
"stats":
- !!int "24"
- !!int "17"
- !!int "22"
- !!int "4"
- !!int "14"
- !!int "6"
"speed": "40 ft."
"skillsaves":
  "Stealth": !!int "8"
  "Perception": !!int "7"
  "Survival": !!int "7"
"damage_resistances": "cold, lightning, necrotic, thunder"
"damage_immunities": "poison; psychic; bludgeoning, piercing, slashing from nonmagical\
  \ attacks"
"condition_immunities": "[charmed](/3-Mechanics/CLI/rules/conditions.md#charmed),\
  \ [exhaustion](/3-Mechanics/CLI/rules/conditions.md#exhaustion), [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened),\
  \ [paralyzed](/3-Mechanics/CLI/rules/conditions.md#paralyzed), [petrified](/3-Mechanics/CLI/rules/conditions.md#petrified),\
  \ [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned), [stunned](/3-Mechanics/CLI/rules/conditions.md#stunned)"
"senses": "blindsight 30 ft., darkvision 60 ft., passive Perception 17"
"languages": "understands Modron and the language of its owner but can't speak"
"cr": "16"
"traits":
- "desc": "The steel predator casts one of the following spells, requiring no spell\
    \ components and using Wisdom as the spellcasting ability:\n\n3/day each:\
    \ [dimension door](/3-Mechanics/CLI/spells/dimension-door.md) (self only), [plane\
    \ shift](/3-Mechanics/CLI/spells/plane-shift.md) (self only)"
  "name": "Spellcasting"
- "desc": "The steel predator has advantage on saving throws against spells and other\
    \ magical effects."
  "name": "Magic Resistance"
- "desc": "The steel predator doesn't require air, food, drink, or sleep."
  "name": "Unusual Nature"
"actions":
- "desc": "The steel predator makes one Bite attack and two Claw attacks."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+12 (+12) to hit, reach 5 ft., one\
    \ target. Hit: dice:2d10 + 7|text(18) (2d10 + 7) lightning damage."
  "name": "Bite"
- "desc": "Melee Weapon Attack: dice: d20+12 (+12) to hit, reach 5 ft., one\
    \ target. Hit: dice:2d8 + 7|text(16) (2d8 + 7) force damage."
  "name": "Claw"
- "desc": "The steel predator emits a roar in a 60-foot cone. Each creature in that\
    \ area must make a DC 19 Constitution saving throw. On a failed save, a creature\
    \ takes dice:6d10|text(33) (6d10) thunder damage, drops everything it's holding,\
    \ and is [stunned](/3-Mechanics/CLI/rules/conditions.md#stunned) for 1 minute.\
    \ The [stunned](/3-Mechanics/CLI/rules/conditions.md#stunned) creature can repeat\
    \ the saving throw at the end of each of its turns, ending the effect on itself\
    \ on a success. On a successful save, a creature takes half as much damage and\
    \ isn't [stunned](/3-Mechanics/CLI/rules/conditions.md#stunned)."
  "name": "Stunning Roar (Recharge 5-6)"
"source":
- "MPMM"
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Steel%20Predator.webp"
```
^statblock

```encounter-table
name: Steel Predator
creatures:
 - 1: Steel Predator
```

## Environment

urban