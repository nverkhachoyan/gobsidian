---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/5
- monster/environment/forest
- monster/environment/grassland
- monster/size/gargantuan
- monster/type/beast/dinosaur
aliases: ["Brontosaurus"]
NoteIcon: monster
BestiaryType: beast (dinosaur)
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 95, Volo's Guide to Monsters p. 139
---
# [Brontosaurus](3-Mechanics\CLI\bestiary\beast/brontosaurus-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 95, Volo's Guide to Monsters p. 139*  

```statblock
"name": "Brontosaurus (MPMM)"
"size": "Gargantuan"
"type": "beast"
"subtype": "dinosaur"
"alignment": "Unaligned"
"ac": !!int "15"
"ac_class": "natural armor"
"hp": !!int "121"
"hit_dice": "9d20 + 27"
"stats":
- !!int "21"
- !!int "9"
- !!int "17"
- !!int "2"
- !!int "10"
- !!int "7"
"speed": "30 ft."
"saves":
  "Constitution": !!int "6"
"senses": "passive Perception 10"
"languages": ""
"cr": "5"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+8 (+8) to hit, reach 20 ft., one target.\
    \ Hit: dice:5d8 + 5|text(27) (5d8 + 5) bludgeoning damage, and the target\
    \ must succeed on a DC 14 Strength saving throw or be knocked [prone](/3-Mechanics/CLI/rules/conditions.md#prone)."
  "name": "Stomp"
- "desc": "Melee Weapon Attack: dice: d20+8 (+8) to hit, reach 20 ft., one target.\
    \ Hit: dice:6d8 + 5|text(32) (6d8 + 5) bludgeoning damage"
  "name": "Tail"
"source":
- "MPMM"
- "VGM"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Brontosaurus.webp"
```
^statblock

```encounter-table
name: Brontosaurus
creatures:
 - 1: Brontosaurus
```

## Environment

forest, grassland