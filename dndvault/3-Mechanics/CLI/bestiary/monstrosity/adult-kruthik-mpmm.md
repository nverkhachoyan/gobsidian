---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/2
- monster/environment/desert
- monster/environment/mountain
- monster/environment/underdark
- monster/size/medium
- monster/type/monstrosity
aliases: ["Adult Kruthik"]
NoteIcon: monster
BestiaryType: monstrosity
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 169, Mordenkainen's Tome of Foes p. 212
---
# [Adult Kruthik](3-Mechanics\CLI\bestiary\monstrosity/adult-kruthik-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 169, Mordenkainen's Tome of Foes p. 212*  

```statblock
"name": "Adult Kruthik (MPMM)"
"size": "Medium"
"type": "monstrosity"
"alignment": "Unaligned"
"ac": !!int "18"
"ac_class": "natural armor"
"hp": !!int "39"
"hit_dice": "6d8 + 12"
"stats":
- !!int "15"
- !!int "16"
- !!int "15"
- !!int "7"
- !!int "12"
- !!int "8"
"speed": "40 ft., burrow 20 ft., climb 40 ft."
"skillsaves":
  "Perception": !!int "5"
"senses": "darkvision 60 ft., tremorsense 60 ft., passive Perception 15"
"languages": "Kruthik"
"cr": "2"
"traits":
- "desc": "The kruthik has advantage on an attack roll against a creature if at least\
    \ one of the kruthik's allies is within 5 feet of the creature and the ally isn't\
    \ [incapacitated](/3-Mechanics/CLI/rules/conditions.md#incapacitated)."
  "name": "Pack Tactics"
- "desc": "The kruthik can burrow through solid rock at half its burrowing speed and\
    \ leaves a 5-foot-diameter tunnel in its wake."
  "name": "Tunneler"
"actions":
- "desc": "The kruthik makes two Stab or Spike attacks."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+5 (+5) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d6 + 3|text(6) (1d6 + 3) piercing damage."
  "name": "Stab"
- "desc": "Ranged Weapon Attack: dice: d20+5 (+5) to hit, range 20/60 ft., one\
    \ target. Hit: dice:1d4 + 3|text(5) (1d4 + 3) piercing damage."
  "name": "Spike"
"source":
- "MPMM"
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Adult%20Kruthik.webp"
```
^statblock

```encounter-table
name: Adult Kruthik
creatures:
 - 1: Adult Kruthik
```

## Environment

desert, mountain, underdark