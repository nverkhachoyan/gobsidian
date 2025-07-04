---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mm
- monster/cr/1-2
- monster/environment/underwater
- monster/size/large
- monster/type/beast
aliases: ["Giant Sea Horse"]
NoteIcon: monster
BestiaryType: beast
SourceType: Bestiary
BookSource: Monster Manual p. 328. Available in the SRD and the Basic Rules.
---
# [Giant Sea Horse](3-Mechanics\CLI\bestiary\beast/giant-sea-horse.md)
*Source: Monster Manual p. 328. Available in the SRD and the Basic Rules.*  

Like their smaller kin, giant sea horses are shy, colorful fish with elongated bodies and curled tails. Aquatic elves train them as mounts.

```statblock
"name": "Giant Sea Horse"
"size": "Large"
"type": "beast"
"alignment": "Unaligned"
"ac": !!int "13"
"ac_class": "natural armor"
"hp": !!int "16"
"hit_dice": "3d10"
"stats":
- !!int "12"
- !!int "15"
- !!int "11"
- !!int "2"
- !!int "12"
- !!int "5"
"speed": "0 ft., swim 40 ft."
"senses": "passive Perception 11"
"languages": ""
"cr": "1/2"
"traits":
- "desc": "If the sea horse moves at least 20 feet straight toward a target and then\
    \ hits it with a ram attack on the same turn, the target takes an extra dice:2d6|text(7)\
    \ (2d6) bludgeoning damage. If the target is a creature, it must succeed on\
    \ a DC 11 Strength saving throw or be knocked [prone](/3-Mechanics/CLI/rules/conditions.md#prone)."
  "name": "Charge"
- "desc": "The sea horse can breathe only underwater."
  "name": "Water Breathing"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+3 (+3) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d6 + 1|text(4) (1d6 + 1) bludgeoning damage."
  "name": "Ram"
"source":
- "MM"
- "GoS"
- "JttRC"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MM/Giant%20Sea%20Horse.webp"
```
^statblock

```encounter-table
name: Giant Sea Horse
creatures:
 - 1: Giant Sea Horse
```

## Environment

underwater