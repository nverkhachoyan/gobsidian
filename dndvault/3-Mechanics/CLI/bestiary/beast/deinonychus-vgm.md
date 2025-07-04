---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/vgm
- monster/cr/1
- monster/environment/forest
- monster/environment/grassland
- monster/environment/hill
- monster/size/medium
- monster/type/beast
aliases: ["Deinonychus"]
NoteIcon: monster
BestiaryType: beast
SourceType: Bestiary
BookSource: Volo's Guide to Monsters p. 139
---
# [Deinonychus](3-Mechanics\CLI\bestiary\beast/deinonychus-vgm.md)
*Source: Volo's Guide to Monsters p. 139*  

This larger cousin of the velociraptor kills by gripping its target with its claws and feeding while the creature is still alive.

```statblock
"name": "Deinonychus (VGM)"
"size": "Medium"
"type": "beast"
"alignment": "Unaligned"
"ac": !!int "13"
"ac_class": "natural armor"
"hp": !!int "26"
"hit_dice": "4d8 + 8"
"stats":
- !!int "15"
- !!int "15"
- !!int "14"
- !!int "4"
- !!int "12"
- !!int "6"
"speed": "40 ft."
"skillsaves":
  "Perception": !!int "3"
"senses": "passive Perception 13"
"languages": ""
"cr": "1"
"traits":
- "desc": "If the deinonychus moves at least 20 feet straight toward a creature and\
    \ then hits it with a claw attack on the same turn, that target must succeed on\
    \ a DC 12 Strength saving throw or be knocked [prone](/3-Mechanics/CLI/rules/conditions.md#prone).\
    \ If the target is [prone](/3-Mechanics/CLI/rules/conditions.md#prone), the deinonychus\
    \ can make one bite attack against it as a bonus action."
  "name": "Pounce"
"actions":
- "desc": "The deinonychus makes three attacks: one with its bite and two with its\
    \ claws."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+4 (+4) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d8 + 2|text(6) (1d8 + 2) piercing damage."
  "name": "Bite"
- "desc": "Melee Weapon Attack: dice: d20+4 (+4) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d8 + 2|text(6) (1d8 + 2) slashing damage."
  "name": "Claw"
"source":
- "VGM"
- "ToA"
- "PSX"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/VGM/Deinonychus.webp"
```
^statblock

```encounter-table
name: Deinonychus
creatures:
 - 1: Deinonychus
```

## Environment

grassland, forest, hill