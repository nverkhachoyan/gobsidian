---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mm
- monster/cr/2
- monster/environment/underwater
- monster/size/large
- monster/type/beast
aliases: ["Hunter Shark"]
NoteIcon: monster
BestiaryType: beast
SourceType: Bestiary
BookSource: Monster Manual p. 330, Dragon of Icespire Peak. Available in the SRD and the Basic Rules.
---
# [Hunter Shark](3-Mechanics\CLI\bestiary\beast/hunter-shark.md)
*Source: Monster Manual p. 330, Dragon of Icespire Peak. Available in the SRD and the Basic Rules.*  

Smaller than a giant shark but larger and fiercer than a reef shark, a hunter shark haunts deep waters. It usually hunts alone, but multiple hunter sharks might feed in the same area. A fully grown hunter shark is 15 to 20 feet long.

```statblock
"name": "Hunter Shark"
"size": "Large"
"type": "beast"
"alignment": "Unaligned"
"ac": !!int "12"
"ac_class": "natural armor"
"hp": !!int "45"
"hit_dice": "6d10 + 12"
"stats":
- !!int "18"
- !!int "13"
- !!int "15"
- !!int "1"
- !!int "10"
- !!int "4"
"speed": "swim 40 ft."
"skillsaves":
  "Perception": !!int "2"
"senses": "blindsight 30 ft., passive Perception 12"
"languages": ""
"cr": "2"
"traits":
- "desc": "The shark has advantage on melee attack rolls against any creature that\
    \ doesn't have all its hit points."
  "name": "Blood Frenzy"
- "desc": "The shark can breathe only underwater."
  "name": "Water Breathing"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+6 (+6) to hit, reach 5 ft., one target.\
    \ Hit: dice:2d8 + 4|text(13) (2d8 + 4) piercing damage."
  "name": "Bite"
"source":
- "MM"
- "PotA"
- "SKT"
- "TftYP"
- "ToA"
- "GoS"
- "DIP"
- "EGW"
- "PSX"
- "SatO"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MM/Hunter%20Shark.webp"
```
^statblock

```encounter-table
name: Hunter Shark
creatures:
 - 1: Hunter Shark
```

## Environment

underwater