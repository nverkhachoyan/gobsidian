---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/vgm
- monster/cr/4
- monster/environment/forest
- monster/environment/grassland
- monster/size/huge
- monster/type/beast
aliases: ["Stegosaurus"]
NoteIcon: monster
BestiaryType: beast
SourceType: Bestiary
BookSource: Volo's Guide to Monsters p. 140
---
# [Stegosaurus](3-Mechanics\CLI\bestiary\beast/stegosaurus-vgm.md)
*Source: Volo's Guide to Monsters p. 140*  

This heavily built dinosaur has rows of plates on its back and a flexible, spiked tail held high to strike predators. It tends to travel in herds of mixed ages.

```statblock
"name": "Stegosaurus (VGM)"
"size": "Huge"
"type": "beast"
"alignment": "Unaligned"
"ac": !!int "13"
"ac_class": "natural armor"
"hp": !!int "76"
"hit_dice": "8d12 + 24"
"stats":
- !!int "20"
- !!int "9"
- !!int "17"
- !!int "2"
- !!int "11"
- !!int "5"
"speed": "40 ft."
"senses": "passive Perception 10"
"languages": ""
"cr": "4"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+7 (+7) to hit, reach 10 ft., one target.\
    \ Hit: dice:6d6 + 5|text(26) (6d6 + 5) piercing damage."
  "name": "Tail"
"source":
- "VGM"
- "ToA"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/VGM/Stegosaurus.webp"
```
^statblock

```encounter-table
name: Stegosaurus
creatures:
 - 1: Stegosaurus
```

## Environment

grassland, forest