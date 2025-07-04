---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/3
- monster/environment/coastal
- monster/environment/urban
- monster/size/medium
- monster/type/humanoid
aliases: ["Swashbuckler"]
NoteIcon: monster
BestiaryType: humanoid
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 238, Volo's Guide to Monsters p. 217
---
# [Swashbuckler](3-Mechanics\CLI\bestiary\humanoid/swashbuckler-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 238, Volo's Guide to Monsters p. 217*  

```statblock
"name": "Swashbuckler (MPMM)"
"size": "Medium"
"type": "humanoid"
"alignment": "Any alignment"
"ac": !!int "17"
"ac_class": "[leather armor](/3-Mechanics/CLI/items/leather-armor.md), suave defense"
"hp": !!int "66"
"hit_dice": "12d8 + 12"
"stats":
- !!int "12"
- !!int "18"
- !!int "12"
- !!int "14"
- !!int "11"
- !!int "15"
"speed": "30 ft."
"skillsaves":
  "Athletics": !!int "5"
  "Acrobatics": !!int "8"
  "Persuasion": !!int "6"
"senses": "passive Perception 10"
"languages": "any one language (usually Common)"
"cr": "3"
"traits":
- "desc": "While the swashbuckler is wearing light or no armor and wielding no [shield](/3-Mechanics/CLI/items/shield.md),\
    \ its AC includes its Charisma modifier."
  "name": "Suave Defense"
"actions":
- "desc": "The swashbuckler makes one Dagger attack and two Rapier attacks."
  "name": "Multiattack"
- "desc": "Melee or Ranged Weapon Attack: dice: d20+6 (+6) to hit, reach 5 ft.\
    \ or range 20/60 ft., one target. Hit: dice:1d4 + 4|text(6) (1d4 + 4) piercing\
    \ damage."
  "name": "Dagger"
- "desc": "Melee Weapon Attack: dice: d20+6 (+6) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d8 + 4|text(8) (1d8 + 4) piercing damage."
  "name": "Rapier"
"bonus_actions":
- "desc": "The swashbuckler takes the [Dash](/3-Mechanics/CLI/rules/actions.md#Dash)\
    \ or [Disengage](/3-Mechanics/CLI/rules/actions.md#Disengage) action."
  "name": "Lightfooted"
"source":
- "MPMM"
- "VGM"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Swashbuckler.webp"
```
^statblock

```encounter-table
name: Swashbuckler
creatures:
 - 1: Swashbuckler
```

## Environment

coastal, urban