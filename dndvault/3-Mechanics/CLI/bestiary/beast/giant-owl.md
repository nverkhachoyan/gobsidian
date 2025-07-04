---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mm
- monster/cr/1-4
- monster/environment/arctic
- monster/environment/forest
- monster/environment/hill
- monster/size/large
- monster/type/beast
aliases: ["Giant Owl"]
NoteIcon: monster
BestiaryType: beast
SourceType: Bestiary
BookSource: Monster Manual p. 327. Available in the SRD and the Basic Rules.
---
# [Giant Owl](3-Mechanics\CLI\bestiary\beast/giant-owl.md)
*Source: Monster Manual p. 327. Available in the SRD and the Basic Rules.*  

Giant owls often befriend fey and other sylvan creatures and are guardians of their woodland realms.

```statblock
"name": "Giant Owl"
"size": "Large"
"type": "beast"
"alignment": "Neutral"
"ac": !!int "12"
"hp": !!int "19"
"hit_dice": "3d10 + 3"
"stats":
- !!int "13"
- !!int "15"
- !!int "12"
- !!int "8"
- !!int "13"
- !!int "10"
"speed": "5 ft., fly 60 ft."
"skillsaves":
  "Stealth": !!int "4"
  "Perception": !!int "5"
"senses": "darkvision 120 ft., passive Perception 15"
"languages": "Giant Owl, understands Common, Elvish, and Sylvan but can't speak them"
"cr": "1/4"
"traits":
- "desc": "The owl doesn't provoke opportunity attacks when it flies out of an enemy's\
    \ reach."
  "name": "Flyby"
- "desc": "The owl has advantage on Wisdom ([Perception](/3-Mechanics/CLI/rules/skills.md#Perception))\
    \ checks that rely on hearing or sight."
  "name": "Keen Hearing and Sight"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+3 (+3) to hit, reach 5 ft., one target.\
    \ Hit: dice:2d6 + 1|text(8) (2d6 + 1) slashing damage."
  "name": "Talons"
"source":
- "MM"
- "SKT"
- "ERLW"
- "IMR"
- "EGW"
- "BMT"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MM/Giant%20Owl.webp"
```
^statblock

```encounter-table
name: Giant Owl
creatures:
 - 1: Giant Owl
```

## Environment

forest, hill, arctic