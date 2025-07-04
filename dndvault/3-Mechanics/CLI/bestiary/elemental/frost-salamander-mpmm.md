---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/9
- monster/environment/arctic
- monster/size/huge
- monster/type/elemental
aliases: ["Frost Salamander"]
NoteIcon: monster
BestiaryType: elemental
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 132, Mordenkainen's Tome of Foes p. 223
---
# [Frost Salamander](3-Mechanics\CLI\bestiary\elemental/frost-salamander-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 132, Mordenkainen's Tome of Foes p. 223*  

```statblock
"name": "Frost Salamander (MPMM)"
"size": "Huge"
"type": "elemental"
"alignment": "Unaligned"
"ac": !!int "17"
"ac_class": "natural armor"
"hp": !!int "168"
"hit_dice": "16d12 + 64"
"stats":
- !!int "20"
- !!int "12"
- !!int "18"
- !!int "7"
- !!int "11"
- !!int "7"
"speed": "60 ft., burrow 40 ft., climb 40 ft."
"saves":
  "Wisdom": !!int "4"
  "Constitution": !!int "8"
"skillsaves":
  "Perception": !!int "4"
"damage_vulnerabilities": "fire"
"damage_immunities": "cold"
"senses": "darkvision 60 ft., tremorsense 60 ft., passive Perception 14"
"languages": "Primordial"
"cr": "9"
"traits":
- "desc": "When the salamander takes fire damage, its\n\nFreezing Breath automatically\
    \ recharges."
  "name": "Burning Fury"
"actions":
- "desc": "The salamander makes one Bite attack and four Claw attacks."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+9 (+9) to hit, reach 15 ft., one target.\
    \ Hit: dice:1d8 + 5|text(9) (1d8 + 5) piercing damage plus dice:1d10|text(5)\
    \ (1d10) cold damage."
  "name": "Bite"
- "desc": "Melee Weapon Attack: dice: d20+9 (+9) to hit, reach 10 ft., one target.\
    \ Hit: dice:1d6 + 5|text(8) (1d6 + 5) piercing damage."
  "name": "Claw"
- "desc": "The salamander exhales chill wind in a 60-foot cone. Each creature in that\
    \ area must make a DC 17 Constitution saving throw, taking dice:8d10|text(44)\
    \ (8d10) cold damage on a failed save, or half as much damage on a successful\
    \ one."
  "name": "Freezing Breath (Recharge 6)"
"source":
- "MPMM"
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Frost%20Salamander.webp"
```
^statblock

```encounter-table
name: Frost Salamander
creatures:
 - 1: Frost Salamander
```

## Environment

arctic