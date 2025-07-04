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
aliases: ["Cow"]
NoteIcon: monster
BestiaryType: beast
SourceType: Bestiary
BookSource: Volo's Guide to Monsters p. 207, Dragon of Icespire Peak
---
# [Cow](3-Mechanics\CLI\bestiary\beast/cow-vgm.md)
*Source: Volo's Guide to Monsters p. 207, Dragon of Icespire Peak*  

There are many kinds of cattle, from common oxen to more unusual, magical variants.

```statblock
"name": "Cow (VGM)"
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
"senses": "passive Perception 10"
"languages": ""
"cr": "1/4"
"traits":
- "desc": "If the cow moves at least 20 feet straight toward a target and then hits\
    \ it with a gore attack on the same turn, the target takes an extra dice:2d6|text(7)\
    \ (2d6) piercing damage."
  "name": "Charge"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+6 (+6) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d6 + 4|text(7) (1d6 + 4) piercing damage."
  "name": "Gore"
"source":
- "VGM"
- "DIP"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/VGM/Cow.webp"
```
^statblock

```encounter-table
name: Cow
creatures:
 - 1: Cow
```

## Environment

grassland