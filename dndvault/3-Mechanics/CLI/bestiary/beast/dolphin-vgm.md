---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/vgm
- monster/cr/1-8
- monster/environment/coastal
- monster/environment/underwater
- monster/size/medium
- monster/type/beast
aliases: ["Dolphin"]
NoteIcon: monster
BestiaryType: beast
SourceType: Bestiary
BookSource: Volo's Guide to Monsters p. 208
---
# [Dolphin](3-Mechanics\CLI\bestiary\beast/dolphin-vgm.md)
*Source: Volo's Guide to Monsters p. 208*  

Dolphins are clever, social marine mammals that feed on small fish and squid. An adult specimen is between 5 and 6 feet long.

```statblock
"name": "Dolphin (VGM)"
"size": "Medium"
"type": "beast"
"alignment": "Unaligned"
"ac": !!int "12"
"ac_class": "natural armor"
"hp": !!int "11"
"hit_dice": "2d8 + 2"
"stats":
- !!int "14"
- !!int "13"
- !!int "13"
- !!int "6"
- !!int "12"
- !!int "7"
"speed": "0 ft., swim 60 ft."
"skillsaves":
  "Perception": !!int "3"
"senses": "blindsight 60 ft., passive Perception 13"
"languages": ""
"cr": "1/8"
"traits":
- "desc": "If the dolphin moves at least 30 feet straight toward a target and then\
    \ hits it with a slam attack on the same turn, the target takes an extra dice:1d6|text(3)\
    \ (1d6) bludgeoning damage."
  "name": "Charge"
- "desc": "The dolphin can hold its breath for 20 minutes."
  "name": "Hold Breath"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+4 (+4) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d6 + 2|text(5) (1d6 + 2) bludgeoning damage."
  "name": "Slam"
"source":
- "VGM"
- "GoS"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/VGM/Dolphin.webp"
```
^statblock

```encounter-table
name: Dolphin
creatures:
 - 1: Dolphin
```

## Environment

underwater, coastal