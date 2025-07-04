---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mm
- monster/cr/2
- monster/environment/forest
- monster/environment/grassland
- monster/environment/hill
- monster/environment/mountain
- monster/size/huge
- monster/type/beast
aliases: ["Giant Elk"]
NoteIcon: monster
BestiaryType: beast
SourceType: Bestiary
BookSource: Monster Manual p. 325. Available in the SRD and the Basic Rules.
---
# [Giant Elk](3-Mechanics\CLI\bestiary\beast/giant-elk.md)
*Source: Monster Manual p. 325. Available in the SRD and the Basic Rules.*  

The majestic giant elk is rare to the point that its appearance is often taken as a foreshadowing of an important event, such as the birth of a king. Legends tell of gods that take the form of giant elk when visiting the Material Plane. Many cultures therefore believe that to hunt these creatures is to invite divine wrath.

```statblock
"name": "Giant Elk"
"size": "Huge"
"type": "beast"
"alignment": "Unaligned"
"ac": !!int "14"
"ac_class": "natural armor"
"hp": !!int "42"
"hit_dice": "5d12 + 10"
"stats":
- !!int "19"
- !!int "16"
- !!int "14"
- !!int "7"
- !!int "14"
- !!int "10"
"speed": "60 ft."
"skillsaves":
  "Perception": !!int "4"
"senses": "passive Perception 14"
"languages": "Giant Elk, understands Common, Elvish, Sylvan but can't speak them"
"cr": "2"
"traits":
- "desc": "If the elk moves at least 20 feet straight toward a target and then hits\
    \ it with a ram attack on the same turn, the target takes an extra dice:2d6|text(7)\
    \ (2d6) damage. If the target is a creature, it must succeed on a DC 14 Strength\
    \ saving throw or be knocked [prone](/3-Mechanics/CLI/rules/conditions.md#prone)."
  "name": "Charge"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+6 (+6) to hit, reach 10 ft., one target.\
    \ Hit: dice:2d6 + 4|text(11) (2d6 + 4) bludgeoning damage."
  "name": "Ram"
- "desc": "Melee Weapon Attack: dice: d20+6 (+6) to hit, reach 5 ft., one [prone](/3-Mechanics/CLI/rules/conditions.md#prone)\
    \ creature. Hit: dice:4d8 + 4|text(22) (4d8 + 4) bludgeoning damage."
  "name": "Hooves"
"source":
- "MM"
- "SKT"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MM/Giant%20Elk.webp"
```
^statblock

```encounter-table
name: Giant Elk
creatures:
 - 1: Giant Elk
```

## Environment

mountain, grassland, forest, hill