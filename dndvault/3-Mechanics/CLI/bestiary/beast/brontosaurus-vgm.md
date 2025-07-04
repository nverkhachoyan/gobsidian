---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/vgm
- monster/cr/5
- monster/environment/forest
- monster/environment/grassland
- monster/size/gargantuan
- monster/type/beast
aliases: ["Brontosaurus"]
NoteIcon: monster
BestiaryType: beast
SourceType: Bestiary
BookSource: Volo's Guide to Monsters p. 139
---
# [Brontosaurus](3-Mechanics\CLI\bestiary\beast/brontosaurus-vgm.md)
*Source: Volo's Guide to Monsters p. 139*  

This massive four-legged dinosaur is large enough that most predators leave it alone. Its deadly tail can drive away or kill smaller threats.

```statblock
"name": "Brontosaurus (VGM)"
"size": "Gargantuan"
"type": "beast"
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
    \ Hit: dice:6d8 + 5|text(32) (6d8 + 5) bludgeoning damage."
  "name": "Tail"
"source":
- "VGM"
- "ToA"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/VGM/Brontosaurus.webp"
```
^statblock

```encounter-table
name: Brontosaurus
creatures:
 - 1: Brontosaurus
```

## Environment

grassland, forest