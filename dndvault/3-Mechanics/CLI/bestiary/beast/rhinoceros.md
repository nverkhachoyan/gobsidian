---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mm
- monster/cr/2
- monster/environment/grassland
- monster/size/large
- monster/type/beast
aliases: ["Rhinoceros"]
NoteIcon: monster
BestiaryType: beast
SourceType: Bestiary
BookSource: Monster Manual p. 336. Available in the SRD and the Basic Rules.
---
# [Rhinoceros](3-Mechanics\CLI\bestiary\beast/rhinoceros.md)
*Source: Monster Manual p. 336. Available in the SRD and the Basic Rules.*  

```statblock
"name": "Rhinoceros"
"size": "Large"
"type": "beast"
"alignment": "Unaligned"
"ac": !!int "11"
"ac_class": "natural armor"
"hp": !!int "45"
"hit_dice": "6d10 + 12"
"stats":
- !!int "21"
- !!int "8"
- !!int "15"
- !!int "2"
- !!int "12"
- !!int "6"
"speed": "40 ft."
"senses": "passive Perception 11"
"languages": ""
"cr": "2"
"traits":
- "desc": "If the rhinoceros moves at least 20 feet straight toward a target and then\
    \ hits it with a gore attack on the same turn, the target takes an extra dice:2d8|text(9)\
    \ (2d8) bludgeoning damage. If the target is a creature, it must succeed on\
    \ a DC 15 Strength saving throw or be knocked [prone](/3-Mechanics/CLI/rules/conditions.md#prone)."
  "name": "Charge"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+7 (+7) to hit, reach 5 ft., one target.\
    \ Hit: dice:2d8 + 5|text(14) (2d8 + 5) bludgeoning damage."
  "name": "Gore"
"source":
- "MM"
- "ToA"
- "WDH"
- "IDRotF"
- "ToFW"
- "QftIS"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MM/Rhinoceros.webp"
```
^statblock

```encounter-table
name: Rhinoceros
creatures:
 - 1: Rhinoceros
```

## Environment

grassland