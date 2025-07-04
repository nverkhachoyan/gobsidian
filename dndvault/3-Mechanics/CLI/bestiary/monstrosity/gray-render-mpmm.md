---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/12
- monster/environment/forest
- monster/environment/hill
- monster/size/large
- monster/type/monstrosity
aliases: ["Gray Render"]
NoteIcon: monster
BestiaryType: monstrosity
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 146, Mordenkainen's Tome of Foes p. 209
---
# [Gray Render](3-Mechanics\CLI\bestiary\monstrosity/gray-render-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 146, Mordenkainen's Tome of Foes p. 209*  

```statblock
"name": "Gray Render (MPMM)"
"size": "Large"
"type": "monstrosity"
"alignment": "Typically  Chaotic Neutral"
"ac": !!int "19"
"ac_class": "natural armor"
"hp": !!int "189"
"hit_dice": "18d10 + 90"
"stats":
- !!int "19"
- !!int "13"
- !!int "20"
- !!int "3"
- !!int "6"
- !!int "8"
"speed": "30 ft."
"saves":
  "Strength": !!int "8"
  "Constitution": !!int "9"
"skillsaves":
  "Perception": !!int "2"
"senses": "darkvision 60 ft., passive Perception 12"
"languages": ""
"cr": "12"
"actions":
- "desc": "The gray render makes one Bite attack and two Claw attacks."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+8 (+8) to hit, reach 5 ft., one target.\
    \ Hit: dice:2d12 + 4|text(17) (2d12 + 4) piercing damage. If the target\
    \ is Medium or smaller, the target must succeed on a DC 16 Strength saving throw\
    \ or be knocked [prone](/3-Mechanics/CLI/rules/conditions.md#prone)."
  "name": "Bite"
- "desc": "Melee Weapon Attack: dice: d20+8 (+8) to hit, reach 10 ft., one target.\
    \ Hit: dice:2d8 + 4|text(13) (2d8 + 4) slashing damage, plus dice:3d6|text(10)\
    \ (3d6) bludgeoning damage if the target is [prone](/3-Mechanics/CLI/rules/conditions.md#prone)."
  "name": "Claw"
"reactions":
- "desc": "When the gray render takes damage, it makes one Claw attack against a random\
    \ creature within its reach, other than its master."
  "name": "Bloody Rampage"
"source":
- "MPMM"
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Gray%20Render.webp"
```
^statblock

```encounter-table
name: Gray Render
creatures:
 - 1: Gray Render
```

## Environment

forest, hill