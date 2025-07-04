---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/10
- monster/environment/underdark
- monster/size/large
- monster/type/aberration/beholder
aliases: ["Death Kiss"]
NoteIcon: monster
BestiaryType: aberration (beholder)
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 85, Volo's Guide to Monsters p. 124
---
# [Death Kiss](3-Mechanics\CLI\bestiary\aberration/death-kiss-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 85, Volo's Guide to Monsters p. 124*  

```statblock
"name": "Death Kiss (MPMM)"
"size": "Large"
"type": "aberration"
"subtype": "beholder"
"alignment": "Typically  Neutral Evil"
"ac": !!int "15"
"ac_class": "natural armor"
"hp": !!int "142"
"hit_dice": "15d10 + 60"
"stats":
- !!int "18"
- !!int "14"
- !!int "18"
- !!int "10"
- !!int "12"
- !!int "10"
"speed": "0 ft., fly 30 ft. (hover)"
"saves":
  "Wisdom": !!int "5"
  "Constitution": !!int "8"
"skillsaves":
  "Perception": !!int "5"
"damage_immunities": "lightning"
"condition_immunities": "[prone](/3-Mechanics/CLI/rules/conditions.md#prone)"
"senses": "darkvision 120 ft., passive Perception 15"
"languages": "Deep Speech, Undercommon"
"cr": "10"
"traits":
- "desc": "A creature within 5 feet of the death kiss takes dice:1d10|text(5) (1d10)\
    \ lightning damage whenever it hits the death kiss with a melee attack that deals\
    \ piercing or slashing damage."
  "name": "Lightning Blood"
"actions":
- "desc": "The death kiss makes three Tentacle attacks. Up to three of these attacks\
    \ can be replaced by Blood Drain—one replacement per tentacle grappling a creature."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+8 (+8) to hit, reach 20 ft., one target.\
    \ Hit: dice:2d6 + 4|text(11) (2d6 + 4) piercing damage, and the target is\
    \ [grappled](/3-Mechanics/CLI/rules/conditions.md#grappled) (escape DC 14) if\
    \ it is a Huge or smaller creature. Until this grapple ends, the target is [restrained](/3-Mechanics/CLI/rules/conditions.md#restrained),\
    \ and the death kiss can't use the same tentacle on another target. The death\
    \ kiss has ten tentacles."
  "name": "Tentacle"
- "desc": "One creature [grappled](/3-Mechanics/CLI/rules/conditions.md#grappled)\
    \ by a tentacle of the death kiss must make a DC 16 Constitution saving throw.\
    \ On a failed save, the target takes dice:4d10|text(22) (4d10) lightning damage,\
    \ and the death kiss regains half as many hit points."
  "name": "Blood Drain"
"source":
- "MPMM"
- "VGM"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Death%20Kiss.webp"
```
^statblock

```encounter-table
name: Death Kiss
creatures:
 - 1: Death Kiss
```

## Environment

underdark