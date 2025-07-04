---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/3
- monster/environment/forest
- monster/environment/urban
- monster/size/medium
- monster/type/humanoid
aliases: ["Archer"]
NoteIcon: monster
BestiaryType: humanoid
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 49, Volo's Guide to Monsters p. 210
---
# [Archer](3-Mechanics\CLI\bestiary\humanoid/archer-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 49, Volo's Guide to Monsters p. 210*  

```statblock
"name": "Archer (MPMM)"
"size": "Medium"
"type": "humanoid"
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
"actions":
- "desc": "The archer makes two Shortsword or Longbow attacks."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+6 (+6) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d6 + 4|text(7) (1d6 + 4) piercing damage."
  "name": "Shortsword"
- "desc": "Ranged Weapon Attack: dice: d20+6 (+6) to hit, range 150/600 ft.,\
    \ one target. Hit: dice:1d8 + 4|text(8) (1d8 + 4) piercing damage."
  "name": "Longbow"
"bonus_actions":
- "desc": "Immediately after making an attack roll or a damage roll with a ranged\
    \ weapon, the archer can roll a dice: d10|avg|noform (d10) and add the number\
    \ rolled to the total."
  "name": "Archer's Eye (3/Day)"
"source":
- "MPMM"
- "VGM"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Archer.webp"
```
^statblock

```encounter-table
name: Archer
creatures:
 - 1: Archer
```

## Environment

forest, urban