---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/tce
- monster/cr/
- monster/size/medium
- monster/type/construct
aliases: ["Steel Defender"]
NoteIcon: monster
BestiaryType: construct
SourceType: Bestiary
BookSource: Tasha's Cauldron of Everything p. 19
---
# [Steel Defender](3-Mechanics\CLI\bestiary\construct/steel-defender-tce.md)
*Source: Tasha's Cauldron of Everything p. 19*  

```statblock
"name": "Steel Defender (TCE)"
"size": "Medium"
"type": "construct"
"alignment": "Unaligned"
"ac": !!int "15"
"ac_class": "natural armor"
"stats":
- !!int "14"
- !!int "12"
- !!int "14"
- !!int "4"
- !!int "10"
- !!int "6"
"speed": "40 ft."
"saves":
  "Dexterity": !!int "0"
  "Constitution": !!int "0"
"skillsaves":
  "Athletics": !!int "0"
  "Perception": !!int "0"
"damage_immunities": "poison"
"condition_immunities": "[charmed](/3-Mechanics/CLI/rules/conditions.md#charmed),\
  \ [exhaustion](/3-Mechanics/CLI/rules/conditions.md#exhaustion), [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)"
"senses": "darkvision 60 ft., passive Perception 0"
"languages": "understands the languages you speak"
"traits":
- "desc": "The defender can't be [surprised](/3-Mechanics/CLI/rules/conditions.md#surprised)."
  "name": "Vigilant"
"actions":
- "desc": "Melee Weapon Attack: the summoner's spell attack modifier to hit, reach\
    \ 5 ft., one target you can see. Hit: 1d8 + PB force damage."
  "name": "Force-Empowered Rend"
- "desc": "The magical mechanisms inside the defender restore 2d8 + PB hit points\
    \ to itself or to one construct or object within 5 feet of it."
  "name": "Repair (3/Day)"
"reactions":
- "desc": "The defender imposes disadvantage on the attack roll of one creature it\
    \ can see that is within 5 feet of it, provided the attack roll is against a creature\
    \ other than the defender."
  "name": "Deflect Attack"
"source":
- "TCE"
- "ERLW"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/TCE/Steel%20Defender.webp"
```
^statblock

```encounter-table
name: Steel Defender
creatures:
 - 1: Steel Defender
```