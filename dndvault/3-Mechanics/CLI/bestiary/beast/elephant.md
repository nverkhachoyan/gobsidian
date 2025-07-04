---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mm
- monster/cr/4
- monster/environment/grassland
- monster/size/huge
- monster/type/beast
aliases: ["Elephant"]
NoteIcon: monster
BestiaryType: beast
SourceType: Bestiary
BookSource: Monster Manual p. 322. Available in the SRD and the Basic Rules.
---
# [Elephant](3-Mechanics\CLI\bestiary\beast/elephant.md)
*Source: Monster Manual p. 322. Available in the SRD and the Basic Rules.*  

```statblock
"name": "Elephant"
"size": "Huge"
"type": "beast"
"alignment": "Unaligned"
"ac": !!int "12"
"ac_class": "natural armor"
"hp": !!int "76"
"hit_dice": "8d12 + 24"
"stats":
- !!int "22"
- !!int "9"
- !!int "17"
- !!int "3"
- !!int "11"
- !!int "6"
"speed": "40 ft."
"senses": "passive Perception 10"
"languages": ""
"cr": "4"
"traits":
- "desc": "If the elephant moves at least 20 feet straight toward a creature and then\
    \ hits it with a gore attack on the same turn, that target must succeed on a DC\
    \ 12 Strength saving throw or be knocked [prone](/3-Mechanics/CLI/rules/conditions.md#prone).\
    \ If the target is [prone](/3-Mechanics/CLI/rules/conditions.md#prone), the elephant\
    \ can make one stomp attack against it as a bonus action."
  "name": "Trampling Charge"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+8 (+8) to hit, reach 5 ft., one target.\
    \ Hit: dice:3d8 + 6|text(19) (3d8 + 6) piercing damage."
  "name": "Gore"
- "desc": "Melee Weapon Attack: dice: d20+8 (+8) to hit, reach 5 ft., one [prone](/3-Mechanics/CLI/rules/conditions.md#prone)\
    \ creature. Hit: dice:3d10 + 6|text(22) (3d10 + 6) bludgeoning damage."
  "name": "Stomp"
"source":
- "MM"
- "ToA"
- "SatO"
- "ToFW"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MM/Elephant.webp"
```
^statblock

```encounter-table
name: Elephant
creatures:
 - 1: Elephant
```

## Environment

grassland