---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/vgm
- monster/cr/1-4
- monster/environment/grassland
- monster/size/large
- monster/type/beast
aliases: ["Rothé"]
NoteIcon: monster
BestiaryType: beast
SourceType: Bestiary
BookSource: Volo's Guide to Monsters p. 208
---
# [Rothé](3-Mechanics\CLI\bestiary\beast/rothe-vgm.md)
*Source: Volo's Guide to Monsters p. 208*  

Ordinary rothé resemble musk oxen and have [darkvision](/3-Mechanics/CLI/rules/senses.md#darkvision) out to a range of 30 feet.

```statblock
"name": "Rothé (VGM)"
"size": "Large"
"type": "beast"
"alignment": "Unaligned"
"ac": !!int "10"
"hp": !!int "15"
"hit_dice": "2d10 + 4"
"stats":
- !!int "18"
- !!int "10"
- !!int "14"
- !!int "2"
- !!int "10"
- !!int "4"
"speed": "30 ft."
"senses": "darkvision 30 ft., passive Perception 10"
"languages": ""
"cr": "1/4"
"traits":
- "desc": "If the rothé moves at least 20 feet straight toward a target and then hits\
    \ it with a gore attack on the same turn, the target takes an extra dice:2d6|text(7)\
    \ (2d6) piercing damage."
  "name": "Charge"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+6 (+6) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d6 + 4|text(7) (1d6 + 4) piercing damage."
  "name": "Gore"
"source":
- "VGM"
- "CoA"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/VGM/Rothe.webp"
```
^statblock

```encounter-table
name: Rothé
creatures:
 - 1: Rothé
```

## Environment

grassland