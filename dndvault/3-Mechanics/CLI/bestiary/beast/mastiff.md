---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mm
- monster/cr/1-8
- monster/environment/forest
- monster/environment/hill
- monster/environment/urban
- monster/size/medium
- monster/type/beast
aliases: ["Mastiff"]
NoteIcon: monster
BestiaryType: beast
SourceType: Bestiary
BookSource: Monster Manual p. 332. Available in the SRD and the Basic Rules.
---
# [Mastiff](3-Mechanics\CLI\bestiary\beast/mastiff.md)
*Source: Monster Manual p. 332. Available in the SRD and the Basic Rules.*  

Mastiffs are impressive hounds prized by humanoids for their loyalty and keen senses. Mastiffs can be trained as guard dogs, hunting dogs, and war dogs. Halflings and other Small humanoids ride them as mounts.

```statblock
"name": "Mastiff"
"size": "Medium"
"type": "beast"
"alignment": "Unaligned"
"ac": !!int "12"
"hp": !!int "5"
"hit_dice": "1d8 + 1"
"stats":
- !!int "13"
- !!int "14"
- !!int "12"
- !!int "3"
- !!int "12"
- !!int "7"
"speed": "40 ft."
"skillsaves":
  "Perception": !!int "3"
"senses": "passive Perception 13"
"languages": ""
"cr": "1/8"
"traits":
- "desc": "The mastiff has advantage on Wisdom ([Perception](/3-Mechanics/CLI/rules/skills.md#Perception))\
    \ checks that rely on hearing or smell."
  "name": "Keen Hearing and Smell"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+3 (+3) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d6 + 1|text(4) (1d6 + 1) piercing damage. If the target is\
    \ a creature, it must succeed on a DC 11 Strength saving throw or be knocked [prone](/3-Mechanics/CLI/rules/conditions.md#prone)."
  "name": "Bite"
"source":
- "MM"
- "CoS"
- "SKT"
- "TftYP"
- "WDH"
- "BGDIA"
- "ERLW"
- "EGW"
- "DSotDQ"
- "KftGV"
- "ToFW"
- "BMT"
- "QftIS"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MM/Mastiff.webp"
```
^statblock

```encounter-table
name: Mastiff
creatures:
 - 1: Mastiff
```

## Environment

forest, hill, urban