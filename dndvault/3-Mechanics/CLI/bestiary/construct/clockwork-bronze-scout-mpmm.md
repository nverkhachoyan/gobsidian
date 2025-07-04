---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/1
- monster/environment/forest
- monster/environment/grassland
- monster/environment/hill
- monster/environment/mountain
- monster/size/medium
- monster/type/construct
aliases: ["Clockwork Bronze Scout"]
NoteIcon: monster
BestiaryType: construct
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 79, Mordenkainen's Tome of Foes p. 125
---
# [Clockwork Bronze Scout](3-Mechanics\CLI\bestiary\construct/clockwork-bronze-scout-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 79, Mordenkainen's Tome of Foes p. 125*  

```statblock
"name": "Clockwork Bronze Scout (MPMM)"
"size": "Medium"
"type": "construct"
"alignment": "Unaligned"
"ac": !!int "13"
"hp": !!int "36"
"hit_dice": "8d8"
"stats":
- !!int "10"
- !!int "16"
- !!int "11"
- !!int "3"
- !!int "14"
- !!int "1"
"speed": "30 ft., burrow 30 ft."
"skillsaves":
  "Stealth": !!int "7"
  "Perception": !!int "6"
"damage_immunities": "poison"
"condition_immunities": "[charmed](/3-Mechanics/CLI/rules/conditions.md#charmed),\
  \ [exhaustion](/3-Mechanics/CLI/rules/conditions.md#exhaustion), [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened),\
  \ [paralyzed](/3-Mechanics/CLI/rules/conditions.md#paralyzed), [petrified](/3-Mechanics/CLI/rules/conditions.md#petrified),\
  \ [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)"
"senses": "darkvision 60 ft., passive Perception 16"
"languages": "understands one language of its creator but can't speak"
"cr": "1"
"traits":
- "desc": "The clockwork doesn't provoke opportunity attacks when it burrows."
  "name": "Earth Armor"
- "desc": "The clockwork has advantage on saving throws against spells and other magical\
    \ effects."
  "name": "Magic Resistance"
- "desc": "The clockwork doesn't require air, food, drink, or sleep."
  "name": "Unusual Nature"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+5 (+5) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d4 + 3|text(5) (1d4 + 3) piercing damage plus dice:1d6|text(3)\
    \ (1d6) lightning damage."
  "name": "Bite"
- "desc": "Each creature in contact with the ground within 15 feet of the clockwork\
    \ must make a DC 13 Dexterity saving throw, taking dice:4d6|text(14) (4d6)\
    \ lightning damage on a failed save, or half as much damage on a successful one."
  "name": "Lightning Flare (Recharges after a Short or Long Rest)"
"source":
- "MPMM"
- "MTF"
- "RtG"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Clockwork%20Bronze%20Scout.webp"
```
^statblock

```encounter-table
name: Clockwork Bronze Scout
creatures:
 - 1: Clockwork Bronze Scout
```

## Environment

forest, grassland, hill, mountain