---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mm
- monster/cr/5
- monster/environment/grassland
- monster/size/huge
- monster/type/beast
aliases: ["Triceratops"]
NoteIcon: monster
BestiaryType: beast
SourceType: Bestiary
BookSource: Monster Manual p. 80. Available in the SRD and the Basic Rules.
---
# [Triceratops](3-Mechanics\CLI\bestiary\beast/triceratops.md)
*Source: Monster Manual p. 80. Available in the SRD and the Basic Rules.*  

One of the most aggressive of the herbivorous dinosaurs, a triceratops has a skull that flares out to form a protective plate of bone. With its great horns and formidable speed, a triceratops gores and tramples would-be predators to death.

## Dinosaurs

Dinosaurs, or behemoths, are among the oldest reptiles in the world. Predatory dinosaurs are savage, territorial hunters. Herbivorous dinosaurs are less aggressive, but they might attack to defend their young, or if startled or harassed.

Dinosaurs come in many sizes and shapes. Larger varieties often have drab coloration, while smaller dinosaurs have colorful markings akin to birds. Dinosaurs roam rugged and isolated areas that humanoids seldom visit, including remote mountain valleys, inaccessible plateaus, tropical islands, and deep fens.

```statblock
"name": "Triceratops"
"size": "Huge"
"type": "beast"
"alignment": "Unaligned"
"ac": !!int "13"
"ac_class": "natural armor"
"hp": !!int "95"
"hit_dice": "10d12 + 30"
"stats":
- !!int "22"
- !!int "9"
- !!int "17"
- !!int "2"
- !!int "11"
- !!int "5"
"speed": "50 ft."
"senses": "passive Perception 10"
"languages": ""
"cr": "5"
"traits":
- "desc": "If the triceratops moves at least 20 feet straight toward a creature and\
    \ then hits it with a gore attack on the same turn, that target must succeed on\
    \ a DC 13 Strength saving throw or be knocked [prone](/3-Mechanics/CLI/rules/conditions.md#prone).\
    \ If the target is [prone](/3-Mechanics/CLI/rules/conditions.md#prone), the triceratops\
    \ can make one stomp attack against it as a bonus action."
  "name": "Trampling Charge"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+9 (+9) to hit, reach 5 ft., one target.\
    \ Hit: dice:4d8 + 6|text(24) (4d8 + 6) piercing damage."
  "name": "Gore"
- "desc": "Melee Weapon Attack: dice: d20+9 (+9) to hit, reach 5 ft., one [prone](/3-Mechanics/CLI/rules/conditions.md#prone)\
    \ creature. Hit: dice:3d10 + 6|text(22) (3d10 + 6) bludgeoning damage"
  "name": "Stomp"
"source":
- "MM"
- "ToA"
- "ToFW"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MM/Triceratops.webp"
```
^statblock

```encounter-table
name: Triceratops
creatures:
 - 1: Triceratops
```

## Environment

grassland