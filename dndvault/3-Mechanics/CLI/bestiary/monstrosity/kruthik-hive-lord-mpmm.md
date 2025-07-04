---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/5
- monster/environment/desert
- monster/environment/mountain
- monster/environment/underdark
- monster/size/large
- monster/type/monstrosity
aliases: ["Kruthik Hive Lord"]
NoteIcon: monster
BestiaryType: monstrosity
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 169, Mordenkainen's Tome of Foes p. 212
---
# [Kruthik Hive Lord](3-Mechanics\CLI\bestiary\monstrosity/kruthik-hive-lord-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 169, Mordenkainen's Tome of Foes p. 212*  

```statblock
"name": "Kruthik Hive Lord (MPMM)"
"size": "Large"
"type": "monstrosity"
"alignment": "Unaligned"
"ac": !!int "20"
"ac_class": "natural armor"
"hp": !!int "102"
"hit_dice": "12d10 + 36"
"stats":
- !!int "19"
- !!int "16"
- !!int "17"
- !!int "10"
- !!int "14"
- !!int "10"
"speed": "40 ft., burrow 20 ft., climb 40 ft."
"skillsaves":
  "Perception": !!int "8"
"senses": "darkvision 60 ft., tremorsense 60 ft., passive Perception 18"
"languages": "Kruthik"
"cr": "5"
"traits":
- "desc": "The kruthik has advantage on an attack roll against a creature if at least\
    \ one of the kruthik's allies is within 5 feet of the creature and the ally isn't\
    \ [incapacitated](/3-Mechanics/CLI/rules/conditions.md#incapacitated)."
  "name": "Pack Tactics"
- "desc": "The kruthik can burrow through solid rock at half its burrowing speed and\
    \ leaves a 10-foot-diameter tunnel in its wake."
  "name": "Tunneler"
"actions":
- "desc": "The kruthik makes two Stab or Spike attacks."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+7 (+7) to hit, reach 10 ft., one target.\
    \ Hit: dice:1d10 + 4|text(9) (1d10 + 4) piercing damage."
  "name": "Stab"
- "desc": "Ranged Weapon Attack: dice: d20+6 (+6) to hit, range 30/120 ft.,\
    \ one target. Hit: dice:1d6 + 4|text(7) (1d6 + 4) piercing damage."
  "name": "Spike"
- "desc": "The kruthik sprays acid in a 15-foot cone. Each creature in that area must\
    \ make a DC 14 Dexterity saving throw, taking dice:4d10|text(22) (4d10) acid\
    \ damage on a failed save, or half as much damage on a successful one."
  "name": "Acid Spray (Recharge 5-6)"
"source":
- "MPMM"
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Kruthik%20Hive%20Lord.webp"
```
^statblock

```encounter-table
name: Kruthik Hive Lord
creatures:
 - 1: Kruthik Hive Lord
```

## Environment

desert, mountain, underdark