---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/vgm
- monster/cr/9
- monster/environment/desert
- monster/environment/urban
- monster/size/medium
- monster/type/humanoid/any-race
aliases: ["Champion"]
NoteIcon: monster
BestiaryType: humanoid (any race)
SourceType: Bestiary
BookSource: Volo's Guide to Monsters p. 212
---
# [Champion](3-Mechanics\CLI\bestiary\humanoid/champion-vgm.md)
*Source: Volo's Guide to Monsters p. 212*  

Champions are mighty warriors who honed their fighting skills in wars or gladiatorial pits. To soldiers and other people who fight for a living, champions are as influential as nobles, and their presence is courted as a sign of status among rulers.

```statblock
"name": "Champion (VGM)"
"size": "Medium"
"type": "humanoid"
"subtype": "any race"
"alignment": "Any alignment"
"ac": !!int "18"
"ac_class": "[plate armor](/3-Mechanics/CLI/items/plate-armor.md)"
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
- "desc": "As a bonus action, the champion can regain 20 hit points."
  "name": "Second Wind (Recharges after a Short or Long Rest)"
"actions":
- "desc": "The champion makes three attacks with its greatsword or its shortbow."
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
"source":
- "VGM"
- "WDMM"
- "TftYP"
- "ToA"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/VGM/Champion.webp"
```
^statblock

```encounter-table
name: Champion
creatures:
 - 1: Champion
```

## Environment

desert, urban