---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/slw
- monster/cr/1-2
- monster/size/medium
- monster/type/construct
aliases: ["Skull Flier"]
NoteIcon: monster
BestiaryType: construct
SourceType: Bestiary
BookSource: Storm Lord's Wrath
---
# [Skull Flier](3-Mechanics\CLI\bestiary\construct/skull-flier-slw.md)
*Source: Storm Lord's Wrath*  

```statblock
"name": "Skull Flier (SLW)"
"size": "Medium"
"type": "construct"
"alignment": "Unaligned"
"ac": !!int "16"
"ac_class": "natural armor"
"hp": !!int "24"
"hit_dice": "3d8"
"stats":
- !!int "10"
- !!int "14"
- !!int "10"
- !!int "1"
- !!int "10"
- !!int "3"
"speed": "10 ft., fly 50 ft."
"damage_immunities": "poison, psychic"
"condition_immunities": "[charmed](/3-Mechanics/CLI/rules/conditions.md#charmed),\
  \ [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened), [paralyzed](/3-Mechanics/CLI/rules/conditions.md#paralyzed),\
  \ [petrified](/3-Mechanics/CLI/rules/conditions.md#petrified), [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)"
"senses": "darkvision 60 ft., passive Perception 10"
"languages": ""
"cr": "1/2"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+4 (+4) to hit, reach 5 ft., one creature.\
    \ Hit: dice:1d6 + 2|text(5) (1d6 + 2) piercing damage, and the target must\
    \ make a DC 11 Constitution saving throw, taking dice:3d6|text(10) (3d6) poison\
    \ damage on a failed save, or half as much damage on a successful one. If the\
    \ poison damage reduces the target to 0 hit points, the target is stable but [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)\
    \ for 1 hour, even after regaining hit points, and is [paralyzed](/3-Mechanics/CLI/rules/conditions.md#paralyzed)\
    \ while [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned) in this way."
  "name": "Sting"
"source":
- "SLW"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/SLW/Skull%20Flier.webp"
```
^statblock

```encounter-table
name: Skull Flier
creatures:
 - 1: Skull Flier
```