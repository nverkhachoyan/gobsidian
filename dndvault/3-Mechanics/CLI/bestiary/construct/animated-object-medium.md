---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/phb
- monster/cr/
- monster/size/medium
- monster/type/construct
aliases: ["Animated Object (Medium)"]
NoteIcon: monster
BestiaryType: construct
SourceType: Bestiary
BookSource: Player's Handbook p. 213
---
# [Animated Object (Medium)](3-Mechanics\CLI\bestiary\construct/animated-object-medium.md)
*Source: Player's Handbook p. 213*  

```statblock
"name": "Animated Object (Medium) (PHB)"
"size": "Medium"
"type": "construct"
"alignment": "Unaligned"
"ac": !!int "13"
"ac_class": "natural armor"
"hp": !!int "40"
"stats":
- !!int "10"
- !!int "12"
- !!int "10"
- !!int "3"
- !!int "3"
- !!int "1"
"speed": "30 ft."
"senses": "blindsight 30 ft. (blind beyond this radius), passive Perception 6"
"languages": ""
"traits":
- "desc": "If the object lacks legs or other appendages it can use for locomotion,\
    \ it instead has a flying speed of 30 feet and can hover. If the object is securely\
    \ attached to a surface or larger object, such as a chain bolted to a wall, its\
    \ speed is 0.\n\nWhen the animated object drops to 0 hit points, it reverts to\
    \ its original object form, and any remaining damage carries over to its original\
    \ object form.\n\nThe DM might rule that a specific objects slam attack inflicts\
    \ slashing or piercing damage based on its form."
  "name": "Animated"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+5 (+5) to hit, reach 5 ft., one target.\
    \ Hit: dice:2d6 + 1|text(7) (2d6 + 1) bludgeoning damage."
  "name": "Slam"
"source":
- "PHB"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/PHB/Animated%20Object%20%28Medium%29.webp"
```
^statblock

```encounter-table
name: Animated Object (Medium)
creatures:
 - 1: Animated Object (Medium)
```