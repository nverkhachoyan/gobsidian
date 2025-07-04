---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mm
- monster/cr/7
- monster/environment/forest
- monster/size/huge
- monster/type/beast
aliases: ["Giant Ape"]
NoteIcon: monster
BestiaryType: beast
SourceType: Bestiary
BookSource: Monster Manual p. 323. Available in the SRD and the Basic Rules.
---
# [Giant Ape](3-Mechanics\CLI\bestiary\beast/giant-ape.md)
*Source: Monster Manual p. 323. Available in the SRD and the Basic Rules.*  

```statblock
"name": "Giant Ape"
"size": "Huge"
"type": "beast"
"alignment": "Unaligned"
"ac": !!int "12"
"hp": !!int "157"
"hit_dice": "15d12 + 60"
"stats":
- !!int "23"
- !!int "14"
- !!int "18"
- !!int "7"
- !!int "12"
- !!int "7"
"speed": "40 ft., climb 40 ft."
"skillsaves":
  "Athletics": !!int "9"
  "Perception": !!int "4"
"senses": "passive Perception 14"
"languages": ""
"cr": "7"
"actions":
- "desc": "The ape makes two fist attacks."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+9 (+9) to hit, reach 10 ft., one target.\
    \ Hit: dice:3d10 + 6|text(22) (3d10 + 6) bludgeoning damage."
  "name": "Fist"
- "desc": "Ranged Weapon Attack: dice: d20+9 (+9) to hit, range 50/100 ft.,\
    \ one target. Hit: dice:7d6 + 6|text(30) (7d6 + 6) bludgeoning damage."
  "name": "Rock"
"source":
- "MM"
- "GoS"
- "GHLoE"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MM/Giant%20Ape.webp"
```
^statblock

```encounter-table
name: Giant Ape
creatures:
 - 1: Giant Ape
```

## Environment

forest