---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/3
- monster/environment/urban
- monster/size/medium
- monster/type/humanoid
aliases: ["Giff"]
NoteIcon: monster
BestiaryType: humanoid
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 138, Mordenkainen's Tome of Foes p. 204
---
# [Giff](3-Mechanics\CLI\bestiary\humanoid/giff-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 138, Mordenkainen's Tome of Foes p. 204*  

```statblock
"name": "Giff (MPMM)"
"size": "Medium"
"type": "humanoid"
"alignment": "Any alignment"
"ac": !!int "16"
"ac_class": "[breastplate](/3-Mechanics/CLI/items/breastplate.md)"
"hp": !!int "60"
"hit_dice": "8d8 + 24"
"stats":
- !!int "18"
- !!int "14"
- !!int "17"
- !!int "11"
- !!int "12"
- !!int "12"
"speed": "30 ft."
"senses": "passive Perception 11"
"languages": "Common"
"cr": "3"
"traits":
- "desc": "The giff's mastery of its weapons enables it to ignore the loading property\
    \ of muskets and pistols."
  "name": "Firearms Knowledge"
- "desc": "The giff can try to knock a creature over; if the giff moves at least 20\
    \ feet in a straight line and ends within 5 feet of a Large or smaller creature,\
    \ that creature must succeed on a DC 14 Strength saving throw or take dice:2d6|text(7)\
    \ (2d6) bludgeoning damage and be knocked [prone](/3-Mechanics/CLI/rules/conditions.md#prone)."
  "name": "Headfirst Charge"
"actions":
- "desc": "The giff makes two Longsword, Musket, or Pistol attacks."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+6 (+6) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d8 + 4|text(8) (1d8 + 4) slashing damage, or dice:1d10 + 4|text(9)\
    \ (1d10 + 4) slashing damage if used with two hands."
  "name": "Longsword"
- "desc": "Ranged Weapon Attack: dice: d20+4 (+4) to hit, range 40/120 ft.,\
    \ one target. Hit: dice:1d12 + 2|text(8) (1d12 + 2) piercing damage."
  "name": "Musket"
- "desc": "Ranged Weapon Attack: dice: d20+4 (+4) to hit, range 30/90 ft., one\
    \ target. Hit: dice:1d10 + 2|text(7) (1d10 + 2) piercing damage."
  "name": "Pistol"
- "desc": "The giff throws a grenade up to 60 feet, and the grenade explodes in a\
    \ 20-foot-radius sphere. Each creature in that area must make a DC 15 Dexterity\
    \ saving throw, taking dice:5d6|text(17) (5d6) piercing damage on a failed\
    \ save, or half as much damage on a successful one."
  "name": "Fragmentation Grenade (1/Day)"
"source":
- "MPMM"
- "MTF"
- "SjA"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Giff.webp"
```
^statblock

```encounter-table
name: Giff
creatures:
 - 1: Giff
```

## Environment

urban