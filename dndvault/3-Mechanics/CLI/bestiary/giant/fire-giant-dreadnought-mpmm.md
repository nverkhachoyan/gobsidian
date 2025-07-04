---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/14
- monster/environment/mountain
- monster/environment/underdark
- monster/size/huge
- monster/type/giant
aliases: ["Fire Giant Dreadnought"]
NoteIcon: monster
BestiaryType: giant
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 124, Volo's Guide to Monsters p. 147
---
# [Fire Giant Dreadnought](3-Mechanics\CLI\bestiary\giant/fire-giant-dreadnought-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 124, Volo's Guide to Monsters p. 147*  

```statblock
"name": "Fire Giant Dreadnought (MPMM)"
"size": "Huge"
"type": "giant"
"alignment": "Typically  Lawful Evil"
"ac": !!int "21"
"ac_class": "[plate](/3-Mechanics/CLI/items/plate-armor.md), [Dual Shields](/3-Mechanics/CLI/items/shield.md)"
"hp": !!int "187"
"hit_dice": "15d12 + 90"
"stats":
- !!int "27"
- !!int "9"
- !!int "23"
- !!int "8"
- !!int "10"
- !!int "11"
"speed": "30 ft."
"saves":
  "Charisma": !!int "5"
  "Dexterity": !!int "4"
  "Constitution": !!int "11"
"skillsaves":
  "Athletics": !!int "13"
  "Perception": !!int "5"
"damage_immunities": "fire"
"senses": "passive Perception 15"
"languages": "Giant"
"cr": "14"
"traits":
- "desc": "The giant carries two shields, which together give the giant +3 to its\
    \ AC (accounted for above)."
  "name": "Dual Shields"
"actions":
- "desc": "The giant makes two Fireshield or Rock attacks."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+13 (+13) to hit, reach 5 ft., one\
    \ target. Hit: dice:4d6 + 8|text(22) (4d6 + 8) bludgeoning damage plus dice:2d6|text(7)\
    \ (2d6) fire damage plus dice:2d6|text(7) (2d6) piercing damage."
  "name": "Fireshield"
- "desc": "Ranged Weapon Attack: dice: d20+13 (+13) to hit, range 60/240 ft.,\
    \ one target. Hit: dice:4d10 + 8|text(30) (4d10 + 8) bludgeoning damage."
  "name": "Rock"
- "desc": "The giant moves up to 30 feet in a straight line and can move through the\
    \ space of any creature smaller than Huge. The first time it enters a creature's\
    \ space during this move, that creature must succeed on a DC 21 Strength saving\
    \ throw or take dice:8d6 + 8|text(36) (8d6 + 8) bludgeoning damage plus dice:4d6|text(14)\
    \ (4d6) fire damage and be pushed up to 30 feet and knocked [prone](/3-Mechanics/CLI/rules/conditions.md#prone)."
  "name": "Shield Charge (Recharge 5-6)"
"source":
- "MPMM"
- "VGM"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Fire%20Giant%20Dreadnought.webp"
```
^statblock

```encounter-table
name: Fire Giant Dreadnought
creatures:
 - 1: Fire Giant Dreadnought
```

## Environment

mountain, underdark