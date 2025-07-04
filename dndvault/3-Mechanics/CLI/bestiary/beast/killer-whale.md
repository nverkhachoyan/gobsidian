---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mm
- monster/cr/3
- monster/environment/underwater
- monster/size/huge
- monster/type/beast
aliases: ["Killer Whale"]
NoteIcon: monster
BestiaryType: beast
SourceType: Bestiary
BookSource: Monster Manual p. 331. Available in the SRD and the Basic Rules.
---
# [Killer Whale](3-Mechanics\CLI\bestiary\beast/killer-whale.md)
*Source: Monster Manual p. 331. Available in the SRD and the Basic Rules.*  

```statblock
"name": "Killer Whale"
"size": "Huge"
"type": "beast"
"alignment": "Unaligned"
"ac": !!int "12"
"ac_class": "natural armor"
"hp": !!int "90"
"hit_dice": "12d12 + 12"
"stats":
- !!int "19"
- !!int "10"
- !!int "13"
- !!int "3"
- !!int "12"
- !!int "7"
"speed": "swim 60 ft."
"skillsaves":
  "Perception": !!int "3"
"senses": "blindsight 120 ft., passive Perception 13"
"languages": ""
"cr": "3"
"traits":
- "desc": "The whale can't use its blindsight while [deafened](/3-Mechanics/CLI/rules/conditions.md#deafened)."
  "name": "Echolocation"
- "desc": "The whale can hold its breath for 30 minutes."
  "name": "Hold Breath"
- "desc": "The whale has advantage on Wisdom ([Perception](/3-Mechanics/CLI/rules/skills.md#Perception))\
    \ checks that rely on hearing."
  "name": "Keen Hearing"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+6 (+6) to hit, reach 5 ft., one target.\
    \ Hit: dice:5d6 + 4|text(21) (5d6 + 4) piercing damage."
  "name": "Bite"
"source":
- "MM"
- "SKT"
- "GoS"
- "EGW"
- "MOT"
- "IDRotF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MM/Killer%20Whale.webp"
```
^statblock

```encounter-table
name: Killer Whale
creatures:
 - 1: Killer Whale
```

## Environment

underwater