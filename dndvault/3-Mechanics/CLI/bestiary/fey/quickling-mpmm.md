---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/1
- monster/environment/forest
- monster/size/tiny
- monster/type/fey
aliases: ["Quickling"]
NoteIcon: monster
BestiaryType: fey
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 207, Volo's Guide to Monsters p. 187
---
# [Quickling](3-Mechanics\CLI\bestiary\fey/quickling-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 207, Volo's Guide to Monsters p. 187*  

```statblock
"name": "Quickling (MPMM)"
"size": "Tiny"
"type": "fey"
"alignment": "Typically  Chaotic Evil"
"ac": !!int "16"
"hp": !!int "10"
"hit_dice": "3d4 + 3"
"stats":
- !!int "4"
- !!int "23"
- !!int "13"
- !!int "10"
- !!int "12"
- !!int "7"
"speed": "120 ft."
"skillsaves":
  "Sleight of Hand": !!int "8"
  "Stealth": !!int "8"
  "Perception": !!int "5"
  "Acrobatics": !!int "8"
"senses": "darkvision 60 ft., passive Perception 15"
"languages": "Common, Sylvan"
"cr": "1"
"traits":
- "desc": "Attack rolls against the quickling have disadvantage unless it is [incapacitated](/3-Mechanics/CLI/rules/conditions.md#incapacitated)\
    \ or its speed is 0."
  "name": "Blurred Movement"
- "desc": "If the quickling is subjected to an effect that allows it to make a Dexterity\
    \ saving throw to take only half damage, it instead takes no damage if it succeeds\
    \ on the saving throw and only half damage if it fails, provided it isn't [incapacitated](/3-Mechanics/CLI/rules/conditions.md#incapacitated)."
  "name": "Evasion"
"actions":
- "desc": "The quickling makes three Dagger attacks."
  "name": "Multiattack"
- "desc": "Melee or Ranged Weapon Attack: dice: d20+8 (+8) to hit, reach 5 ft.\
    \ or range 20/60 ft., one target. Hit: dice:1d4 + 6|text(8) (1d4 + 6) piercing\
    \ damage."
  "name": "Dagger"
"source":
- "MPMM"
- "VGM"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Quickling.webp"
```
^statblock

```encounter-table
name: Quickling
creatures:
 - 1: Quickling
```

## Environment

forest