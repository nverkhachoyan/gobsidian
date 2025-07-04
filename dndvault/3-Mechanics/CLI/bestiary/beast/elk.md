---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mm
- monster/cr/1-4
- monster/environment/forest
- monster/environment/grassland
- monster/environment/hill
- monster/size/large
- monster/type/beast
aliases: ["Elk"]
NoteIcon: monster
BestiaryType: beast
SourceType: Bestiary
BookSource: Monster Manual p. 322. Available in the SRD and the Basic Rules.
---
# [Elk](3-Mechanics\CLI\bestiary\beast/elk.md)
*Source: Monster Manual p. 322. Available in the SRD and the Basic Rules.*  

```statblock
"name": "Elk"
"size": "Large"
"type": "beast"
"alignment": "Unaligned"
"ac": !!int "10"
"hp": !!int "13"
"hit_dice": "2d10 + 2"
"stats":
- !!int "16"
- !!int "10"
- !!int "12"
- !!int "2"
- !!int "10"
- !!int "6"
"speed": "50 ft."
"senses": "passive Perception 10"
"languages": ""
"cr": "1/4"
"traits":
- "desc": "If the elk moves at least 20 feet straight toward a target and then hits\
    \ it with a ram attack on the same turn, the target takes an extra dice:2d6|text(7)\
    \ (2d6) damage. If the target is a creature, it must succeed on a DC 13 Strength\
    \ saving throw or be knocked [prone](/3-Mechanics/CLI/rules/conditions.md#prone)."
  "name": "Charge"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+5 (+5) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d6 + 3|text(6) (1d6 + 3) bludgeoning damage."
  "name": "Ram"
- "desc": "Melee Weapon Attack: dice: d20+5 (+5) to hit, reach 5 ft., one [prone](/3-Mechanics/CLI/rules/conditions.md#prone)\
    \ creature. Hit: dice:2d4 + 3|text(8) (2d4 + 3) bludgeoning damage."
  "name": "Hooves"
"source":
- "MM"
- "HotDQ"
- "PotA"
- "SKT"
- "EGW"
- "IDRotF"
- "QftIS"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MM/Elk.webp"
```
^statblock

```encounter-table
name: Elk
creatures:
 - 1: Elk
```

## Environment

grassland, forest, hill