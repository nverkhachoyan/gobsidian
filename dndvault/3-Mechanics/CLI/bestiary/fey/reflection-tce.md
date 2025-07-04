---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/tce
- monster/cr/1-2
- monster/size/medium
- monster/type/fey
aliases: ["Reflection"]
NoteIcon: monster
BestiaryType: fey
SourceType: Bestiary
BookSource: Tasha's Cauldron of Everything p. 158
---
# [Reflection](3-Mechanics\CLI\bestiary\fey/reflection-tce.md)
*Source: Tasha's Cauldron of Everything p. 158*  

```statblock
"name": "Reflection (TCE)"
"size": "Medium"
"type": "fey"
"alignment": "Chaotic Evil"
"ac": !!int "12"
"hp": !!int "16"
"hit_dice": "3d8 + 3"
"stats":
- !!int "6"
- !!int "14"
- !!int "13"
- !!int "6"
- !!int "10"
- !!int "8"
"speed": "40 ft."
"skillsaves":
  "Stealth": !!int "4"
"damage_vulnerabilities": "bludgeoning"
"damage_resistances": "acid; cold; fire; lightning; thunder; piercing, slashing from\
  \ nonmagical attacks"
"damage_immunities": "necrotic, poison"
"condition_immunities": "[exhaustion](/3-Mechanics/CLI/rules/conditions.md#exhaustion),\
  \ [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened), [grappled](/3-Mechanics/CLI/rules/conditions.md#grappled),\
  \ [paralyzed](/3-Mechanics/CLI/rules/conditions.md#paralyzed), [petrified](/3-Mechanics/CLI/rules/conditions.md#petrified),\
  \ [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned), [prone](/3-Mechanics/CLI/rules/conditions.md#prone),\
  \ [restrained](/3-Mechanics/CLI/rules/conditions.md#restrained)"
"senses": "darkvision 60 ft., passive Perception 10"
"languages": ""
"cr": "1/2"
"traits":
- "desc": "The reflection can move through a space as narrow as 1 inch wide without\
    \ squeezing."
  "name": "Amorphous"
- "desc": "While in dim light or darkness, the reflection can take the Hide action\
    \ as a bonus action. Its stealth bonus is also improved to +6."
  "name": "Shadow Stealth"
- "desc": "While in sunlight, the reflection has disadvantage on attack rolls, ability\
    \ checks, and saving throws."
  "name": "Sunlight Weakness"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+4 (+4) to hit, reach 5 ft., one creature.\
    \ Hit: dice:2d6 + 2|text(9) (2d6 + 2) necrotic damage, and the target's\
    \ Strength score is reduced by dice: 1d4|avg|noform (1d4). The target dies\
    \ if this reduces its Strength to 0. Otherwise, the reduction lasts until the\
    \ target finishes a short or long rest.\n\nIf a non-evil humanoid dies from this\
    \ attack, a new reflection rises from the corpse dice: 1d4|avg|noform (1d4)\
    \ hours later."
  "name": "Strength Drain"
"source":
- "TCE"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/TCE/Reflection.webp"
```
^statblock

```encounter-table
name: Reflection
creatures:
 - 1: Reflection
```