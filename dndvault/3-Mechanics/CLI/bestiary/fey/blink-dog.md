---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mm
- monster/cr/1-4
- monster/environment/forest
- monster/size/medium
- monster/type/fey
aliases: ["Blink Dog"]
NoteIcon: monster
BestiaryType: fey
SourceType: Bestiary
BookSource: Monster Manual p. 318, Tasha's Cauldron of Everything. Available in the SRD and the Basic Rules.
---
# [Blink Dog](3-Mechanics\CLI\bestiary\fey/blink-dog.md)
*Source: Monster Manual p. 318, Tasha's Cauldron of Everything. Available in the SRD and the Basic Rules.*  

A blink dog takes its name from its ability to blink in and out of existence, a talent it uses to aid its attacks and to avoid harm. Blink dogs harbor a long-standing hatred for displacer beasts and attack them on sight.

```statblock
"name": "Blink Dog"
"size": "Medium"
"type": "fey"
"alignment": "Lawful Good"
"ac": !!int "13"
"hp": !!int "22"
"hit_dice": "4d8 + 4"
"stats":
- !!int "12"
- !!int "17"
- !!int "12"
- !!int "10"
- !!int "13"
- !!int "11"
"speed": "40 ft."
"skillsaves":
  "Stealth": !!int "5"
  "Perception": !!int "3"
"senses": "passive Perception 13"
"languages": "Blink Dog, understands Sylvan but can't speak it"
"cr": "1/4"
"traits":
- "desc": "The dog has advantage on Wisdom ([Perception](/3-Mechanics/CLI/rules/skills.md#Perception))\
    \ checks that rely on hearing or smell."
  "name": "Keen Hearing and Smell"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+3 (+3) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d6 + 1|text(4) (1d6 + 1) piercing damage."
  "name": "Bite"
- "desc": "The dog magically teleports, along with any equipment it is wearing or\
    \ carrying, up to 40 feet to an unoccupied space it can see. Before or after teleporting,\
    \ the dog can make one bite attack."
  "name": "Teleport (Recharge 4-6)"
"source":
- "MM"
- "WDMM"
- "GoS"
- "ERLW"
- "MOT"
- "TCE"
- "WBtW"
- "SatO"
- "BMT"
- "VEoR"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MM/Blink%20Dog.webp"
```
^statblock

```encounter-table
name: Blink Dog
creatures:
 - 1: Blink Dog
```

## Environment

forest