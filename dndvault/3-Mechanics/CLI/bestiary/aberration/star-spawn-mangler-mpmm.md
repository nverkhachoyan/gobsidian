---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/5
- monster/size/medium
- monster/type/aberration
aliases: ["Star Spawn Mangler"]
NoteIcon: monster
BestiaryType: aberration
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 229, Mordenkainen's Tome of Foes p. 236
---
# [Star Spawn Mangler](3-Mechanics\CLI\bestiary\aberration/star-spawn-mangler-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 229, Mordenkainen's Tome of Foes p. 236*  

```statblock
"name": "Star Spawn Mangler (MPMM)"
"size": "Medium"
"type": "aberration"
"alignment": "Typically  Chaotic Evil"
"ac": !!int "14"
"hp": !!int "71"
"hit_dice": "13d8 + 13"
"stats":
- !!int "8"
- !!int "18"
- !!int "12"
- !!int "11"
- !!int "12"
- !!int "7"
"speed": "40 ft., climb 40 ft."
"saves":
  "Dexterity": !!int "7"
  "Constitution": !!int "4"
"skillsaves":
  "Stealth": !!int "7"
"damage_resistances": "cold"
"damage_immunities": "psychic"
"condition_immunities": "[charmed](/3-Mechanics/CLI/rules/conditions.md#charmed),\
  \ [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened), [prone](/3-Mechanics/CLI/rules/conditions.md#prone)"
"senses": "darkvision 60 ft., passive Perception 11"
"languages": "Deep Speech"
"cr": "5"
"traits":
- "desc": "The mangler has advantage on initiative rolls."
  "name": "Ambusher"
"actions":
- "desc": "The mangler makes two Claw attacks."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+7 (+7) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d8 + 4|text(8) (1d8 + 4) slashing damage. If the attack roll\
    \ has advantage, the target also takes dice:2d6|text(7) (2d6) psychic damage."
  "name": "Claw"
- "desc": "The mangler makes six Claw attacks. Either before or after these attacks,\
    \ it can move up to its speed without provoking opportunity attacks."
  "name": "Flurry of Claws (Recharge 5-6)"
"bonus_actions":
- "desc": "While in dim light or darkness, the mangler takes the [Hide](/3-Mechanics/CLI/rules/actions.md#Hide)\
    \ action."
  "name": "Shadow Stealth"
"source":
- "MPMM"
- "MTF"
- "BMT"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Star%20Spawn%20Mangler.webp"
```
^statblock

```encounter-table
name: Star Spawn Mangler
creatures:
 - 1: Star Spawn Mangler
```