---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/vgm
- monster/cr/3
- monster/environment/forest
- monster/environment/urban
- monster/size/medium
- monster/type/humanoid/any-race
aliases: ["Archer"]
NoteIcon: monster
BestiaryType: humanoid (any race)
SourceType: Bestiary
BookSource: Volo's Guide to Monsters p. 210, Dragon of Icespire Peak, Storm Lord's Wrath
---
# [Archer](3-Mechanics\CLI\bestiary\humanoid/archer-vgm.md)
*Source: Volo's Guide to Monsters p. 210, Dragon of Icespire Peak, Storm Lord's Wrath*  

Archers defend castles, hunt wild game on the fringes of civilization, serve as artillery in military units, and occasionally make good coin as brigands or caravan guards.

```statblock
"name": "Archer (VGM)"
"size": "Medium"
"type": "humanoid"
"subtype": "any race"
"alignment": "Any alignment"
"ac": !!int "16"
"ac_class": "[studded leather](/3-Mechanics/CLI/items/studded-leather-armor.md)"
"hp": !!int "75"
"hit_dice": "10d8 + 30"
"stats":
- !!int "11"
- !!int "18"
- !!int "16"
- !!int "11"
- !!int "13"
- !!int "10"
"speed": "30 ft."
"skillsaves":
  "Perception": !!int "5"
  "Acrobatics": !!int "6"
"senses": "passive Perception 15"
"languages": "any one language (usually Common)"
"cr": "3"
"traits":
- "desc": "As a bonus action, the archer can add dice: 1d10|avg|noform (1d10)\
    \ to its next attack or damage roll with a longbow or shortbow."
  "name": "Archer's Eye (3/Day)"
"actions":
- "desc": "The archer makes two attacks with its longbow."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+6 (+6) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d6 + 4|text(7) (1d6 + 4) piercing damage."
  "name": "Shortsword"
- "desc": "Ranged Weapon Attack: dice: d20+6 (+6) to hit, range 150/600 ft.,\
    \ one target. Hit: dice:1d8 + 4|text(8) (1d8 + 4) piercing damage."
  "name": "Longbow"
"source":
- "VGM"
- "DIP"
- "SLW"
- "MOT"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/VGM/Archer.webp"
```
^statblock

```encounter-table
name: Archer
creatures:
 - 1: Archer
```

## Environment

forest, urban