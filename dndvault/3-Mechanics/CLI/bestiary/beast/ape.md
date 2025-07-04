---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mm
- monster/cr/1-2
- monster/environment/forest
- monster/size/medium
- monster/type/beast
aliases: ["Ape"]
NoteIcon: monster
BestiaryType: beast
SourceType: Bestiary
BookSource: Monster Manual p. 317. Available in the SRD and the Basic Rules.
---
# [Ape](3-Mechanics\CLI\bestiary\beast/ape.md)
*Source: Monster Manual p. 317. Available in the SRD and the Basic Rules.*  

```statblock
"name": "Ape"
"size": "Medium"
"type": "beast"
"alignment": "Unaligned"
"ac": !!int "12"
"hp": !!int "19"
"hit_dice": "3d8 + 6"
"stats":
- !!int "16"
- !!int "14"
- !!int "14"
- !!int "6"
- !!int "12"
- !!int "7"
"speed": "30 ft., climb 30 ft."
"skillsaves":
  "Athletics": !!int "5"
  "Perception": !!int "3"
"senses": "passive Perception 13"
"languages": ""
"cr": "1/2"
"actions":
- "desc": "The ape makes two fist attacks."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+5 (+5) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d6 + 3|text(6) (1d6 + 3) bludgeoning damage."
  "name": "Fist"
- "desc": "Ranged Weapon Attack: dice: d20+5 (+5) to hit, range 25/50 ft., one\
    \ target. Hit: dice:1d6 + 3|text(6) (1d6 + 3) bludgeoning damage."
  "name": "Rock"
"source":
- "MM"
- "TftYP"
- "ToA"
- "WDH"
- "GoS"
- "CM"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MM/Ape.webp"
```
^statblock

```encounter-table
name: Ape
creatures:
 - 1: Ape
```

## Environment

forest