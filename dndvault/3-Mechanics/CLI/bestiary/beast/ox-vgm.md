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
aliases: ["Ox"]
NoteIcon: monster
BestiaryType: beast
SourceType: Bestiary
BookSource: Volo's Guide to Monsters p. 208
---
# [Ox](3-Mechanics\CLI\bestiary\beast/ox-vgm.md)
*Source: Volo's Guide to Monsters p. 208*  

An ox is mainly used for draft work rather than meat or milk.

```statblock
"name": "Ox (VGM)"
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
- "desc": "If the ox moves at least 20 feet straight toward a target and then hits\
    \ it with a gore attack on the same turn, the target takes an extra dice:2d6|text(7)\
    \ (2d6) piercing damage."
  "name": "Charge"
- "desc": "The oxen is considered to be a Huge animal for the purposes of determining\
    \ its carrying capacity."
  "name": "Beast of Burden"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+6 (+6) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d6 + 4|text(7) (1d6 + 4) piercing damage."
  "name": "Gore"
"source":
- "VGM"
- "PotA"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/VGM/Ox.webp"
```
^statblock

```encounter-table
name: Ox
creatures:
 - 1: Ox
```

## Environment

grassland