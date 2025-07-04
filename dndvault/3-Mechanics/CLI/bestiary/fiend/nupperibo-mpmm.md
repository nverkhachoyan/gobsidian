---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/1-2
- monster/size/medium
- monster/type/fiend/devil
aliases: ["Nupperibo"]
NoteIcon: monster
BestiaryType: fiend (devil)
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 196, Mordenkainen's Tome of Foes p. 168
---
# [Nupperibo](3-Mechanics\CLI\bestiary\fiend/nupperibo-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 196, Mordenkainen's Tome of Foes p. 168*  

```statblock
"name": "Nupperibo (MPMM)"
"size": "Medium"
"type": "fiend"
"subtype": "devil"
"alignment": "Typically  Lawful Evil"
"ac": !!int "13"
"ac_class": "natural armor"
"hp": !!int "11"
"hit_dice": "2d8 + 2"
"stats":
- !!int "16"
- !!int "11"
- !!int "13"
- !!int "3"
- !!int "8"
- !!int "1"
"speed": "20 ft."
"skillsaves":
  "Perception": !!int "1"
"damage_resistances": "acid, cold"
"damage_immunities": "fire, poison"
"condition_immunities": "[blinded](/3-Mechanics/CLI/rules/conditions.md#blinded),\
  \ [charmed](/3-Mechanics/CLI/rules/conditions.md#charmed), [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened),\
  \ [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)"
"senses": "blindsight 20 ft. (blind beyond this radius), passive Perception 11"
"languages": "understands Infernal but can't speak"
"cr": "1/2"
"traits":
- "desc": "Any creature, other than a devil, that starts its turn within 20 feet of\
    \ one or more nupperibos must succeed on a DC 11 Constitution saving throw or\
    \ take dice:2d4|text(5) (2d4) acid damage. A creature within the areas of\
    \ two or more nupperibos makes the saving throw with disadvantage."
  "name": "Cloud of Vermin"
- "desc": "In the Nine Hells, the nupperibo can flawlessly track any creature that\
    \ has taken damage from any nupperibo's Cloud of Vermin within the previous 24\
    \ hours."
  "name": "Driven Tracker"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+5 (+5) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d6 + 3|text(6) (1d6 + 3) piercing damage."
  "name": "Bite"
"source":
- "MPMM"
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Nupperibo.webp"
```
^statblock

```encounter-table
name: Nupperibo
creatures:
 - 1: Nupperibo
```