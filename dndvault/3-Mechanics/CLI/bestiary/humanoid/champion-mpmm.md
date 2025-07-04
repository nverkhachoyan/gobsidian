---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/9
- monster/environment/desert
- monster/environment/urban
- monster/size/medium
- monster/type/humanoid
aliases: ["Champion"]
NoteIcon: monster
BestiaryType: humanoid
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 74, Volo's Guide to Monsters p. 212
---
# [Champion](3-Mechanics\CLI\bestiary\humanoid/champion-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 74, Volo's Guide to Monsters p. 212*  

```statblock
"name": "Champion (MPMM)"
"size": "Medium"
"type": "humanoid"
"alignment": "Any alignment"
"ac": !!int "18"
"ac_class": "[plate](/3-Mechanics/CLI/items/plate-armor.md)"
"hp": !!int "143"
"hit_dice": "22d8 + 44"
"stats":
- !!int "20"
- !!int "15"
- !!int "14"
- !!int "10"
- !!int "14"
- !!int "12"
"speed": "30 ft."
"saves":
  "Strength": !!int "9"
  "Constitution": !!int "6"
"skillsaves":
  "Intimidation": !!int "5"
  "Athletics": !!int "9"
  "Perception": !!int "6"
"senses": "passive Perception 16"
"languages": "any one language (usually Common)"
"cr": "9"
"traits":
- "desc": "The champion rerolls a failed saving throw."
  "name": "Indomitable (2/Day)"
"actions":
- "desc": "The champion makes three Greatsword or Shortbow attacks."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+9 (+9) to hit, reach 5 ft., one target.\
    \ Hit: dice:2d6 + 5|text(12) (2d6 + 5) slashing damage, plus dice:2d6|text(7)\
    \ (2d6) slashing damage if the champion has more than half of its total hit\
    \ points remaining."
  "name": "Greatsword"
- "desc": "Ranged Weapon Attack: dice: d20+6 (+6) to hit, range 80/320 ft.,\
    \ one target. Hit: dice:1d6 + 2|text(5) (1d6 + 2) piercing damage, plus\
    \ dice:2d6|text(7) (2d6) piercing damage if the champion has more than half\
    \ of its total hit points remaining."
  "name": "Shortbow"
"bonus_actions":
- "desc": "The champion regains 20 hit points."
  "name": "Second Wind (Recharges after a Short or Long Rest)"
"source":
- "MPMM"
- "VGM"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Champion.webp"
```
^statblock

```encounter-table
name: Champion
creatures:
 - 1: Champion
```

## Environment

desert, urban