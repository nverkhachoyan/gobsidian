---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/slw
- monster/cr/10
- monster/size/large
- monster/type/elemental
aliases: ["Statue of Talos"]
NoteIcon: monster
BestiaryType: elemental
SourceType: Bestiary
BookSource: Storm Lord's Wrath
---
# [Statue of Talos](3-Mechanics\CLI\bestiary\elemental/statue-of-talos-slw.md)
*Source: Storm Lord's Wrath*  

```statblock
"name": "Statue of Talos (SLW)"
"size": "Large"
"type": "elemental"
"alignment": "Chaotic Evil"
"ac": !!int "17"
"ac_class": "natural armor"
"hp": !!int "147"
"hit_dice": "14d10 + 70"
"stats":
- !!int "19"
- !!int "11"
- !!int "20"
- !!int "6"
- !!int "11"
- !!int "9"
"speed": "30 ft., fly 60 ft."
"saves":
  "Wisdom": !!int "4"
"skillsaves":
  "Perception": !!int "4"
"damage_resistances": "bludgeoning, piercing, slashing from nonmagical attacks that\
  \ aren't adamantine"
"damage_immunities": "poison"
"condition_immunities": "[exhaustion](/3-Mechanics/CLI/rules/conditions.md#exhaustion),\
  \ [petrified](/3-Mechanics/CLI/rules/conditions.md#petrified), [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)"
"senses": "darkvision 60 ft., passive Perception 14"
"languages": "Terran"
"cr": "10"
"traits":
- "desc": "While the statue remains motionless, it is indistinguishable from an inanimate\
    \ statue."
  "name": "False Appearance"
"actions":
- "desc": "The statue makes five attacks: one with its headbutt and four with its\
    \ lightning bolt blades."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+8 (+8) to hit, reach 5 ft., one target.\
    \ Hit: dice:2d6 + 4|text(11) (2d6 + 4) piercing damage."
  "name": "Headbutt"
- "desc": "Melee Weapon Attack: dice: d20+8 (+8) to hit, reach 5 ft., one target.\
    \ Hit: dice:2d4 + 4|text(9) (2d4 + 4) slashing damage."
  "name": "Lightning Bolt Blades"
"source":
- "SLW"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/SLW/Statue%20of%20Talos.webp"
```
^statblock

```encounter-table
name: Statue of Talos
creatures:
 - 1: Statue of Talos
```