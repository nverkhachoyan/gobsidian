---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/tce
- monster/cr/
- monster/size/tiny
- monster/type/construct
aliases: ["Homunculus Servant"]
NoteIcon: monster
BestiaryType: construct
SourceType: Bestiary
BookSource: Tasha's Cauldron of Everything p. 22
---
# [Homunculus Servant](3-Mechanics\CLI\bestiary\construct/homunculus-servant-tce.md)
*Source: Tasha's Cauldron of Everything p. 22*  

```statblock
"name": "Homunculus Servant (TCE)"
"size": "Tiny"
"type": "construct"
"alignment": "Unaligned"
"ac": !!int "13"
"ac_class": "natural armor"
"stats":
- !!int "4"
- !!int "15"
- !!int "12"
- !!int "10"
- !!int "10"
- !!int "7"
"speed": "20 ft., fly 30 ft."
"saves":
  "Dexterity": !!int "0"
"skillsaves":
  "Stealth": !!int "0"
  "Perception": !!int "0"
"damage_immunities": "poison"
"condition_immunities": "[exhaustion](/3-Mechanics/CLI/rules/conditions.md#exhaustion),\
  \ [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)"
"senses": "darkvision 60 ft., passive Perception 0"
"languages": "understands the languages you speak"
"traits":
- "desc": "If the homunculus is subjected to an effect that allows it to make a Dexterity\
    \ saving throw to take only half damage, it instead takes no damage if it succeeds\
    \ on the saving throw, and only half damage if it fails. It can't use this trait\
    \ if it's [incapacitated](/3-Mechanics/CLI/rules/conditions.md#incapacitated)."
  "name": "Evasion"
"actions":
- "desc": "Ranged Weapon Attack: the summoner's spell attack modifier to hit, range\
    \ 30 ft., one target you can see. Hit: 1d4 + PB force damage."
  "name": "Force Strike"
"reactions":
- "desc": "The homunculus delivers a spell you cast that has a range of touch. The\
    \ homunculus must be within 120 feet of you."
  "name": "Channel Magic"
"source":
- "TCE"
- "ERLW"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/TCE/Homunculus%20Servant.webp"
```
^statblock

```encounter-table
name: Homunculus Servant
creatures:
 - 1: Homunculus Servant
```