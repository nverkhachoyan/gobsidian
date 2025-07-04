---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mm
- monster/cr/0
- monster/environment/forest
- monster/environment/hill
- monster/size/small
- monster/type/beast
aliases: ["Baboon"]
NoteIcon: monster
BestiaryType: beast
SourceType: Bestiary
BookSource: Monster Manual p. 318. Available in the SRD and the Basic Rules.
---
# [Baboon](3-Mechanics\CLI\bestiary\beast/baboon.md)
*Source: Monster Manual p. 318. Available in the SRD and the Basic Rules.*  

```statblock
"name": "Baboon"
"size": "Small"
"type": "beast"
"alignment": "Unaligned"
"ac": !!int "12"
"hp": !!int "3"
"hit_dice": "1d6"
"stats":
- !!int "8"
- !!int "14"
- !!int "11"
- !!int "4"
- !!int "12"
- !!int "6"
"speed": "30 ft., climb 30 ft."
"senses": "passive Perception 11"
"languages": ""
"cr": "0"
"traits":
- "desc": "The baboon has advantage on an attack roll against a creature if at least\
    \ one of the baboon's allies is within 5 feet of the creature and the ally isn't\
    \ [incapacitated](/3-Mechanics/CLI/rules/conditions.md#incapacitated)."
  "name": "Pack Tactics"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+1 (+1) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d4 - 1|text(1) (1d4 - 1) piercing damage."
  "name": "Bite"
"source":
- "MM"
- "TftYP"
- "ToA"
- "CM"
- "CoS"
- "WBtW"
- "KftGV"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MM/Baboon.webp"
```
^statblock

```encounter-table
name: Baboon
creatures:
 - 1: Baboon
```

## Environment

forest, hill