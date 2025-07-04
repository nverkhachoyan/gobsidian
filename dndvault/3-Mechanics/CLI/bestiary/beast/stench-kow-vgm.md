---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/vgm
- monster/cr/1-4
- monster/size/large
- monster/type/beast
aliases: ["Stench Kow"]
NoteIcon: monster
BestiaryType: beast
SourceType: Bestiary
BookSource: Volo's Guide to Monsters p. 208
---
# [Stench Kow](3-Mechanics\CLI\bestiary\beast/stench-kow-vgm.md)
*Source: Volo's Guide to Monsters p. 208*  

These orange and green misshapen bison are native to the Lower Planes.

```statblock
"name": "Stench Kow (VGM)"
"size": "Large"
"type": "beast"
"alignment": "Unaligned"
"ac": !!int "10"
"hp": !!int "15"
"hit_dice": "2d10 + 4"
"stats":
- !!int "18"
- !!int "10"
- !!int "14"
- !!int "2"
- !!int "10"
- !!int "4"
"speed": "30 ft."
"damage_resistances": "cold, fire, poison"
"senses": "darkvision 60 ft., passive Perception 10"
"languages": ""
"cr": "1/4"
"traits":
- "desc": "If the kow moves at least 20 feet straight toward a target and then hits\
    \ it with a gore attack on the same turn, the target takes an extra dice:2d6|text(7)\
    \ (2d6) piercing damage."
  "name": "Charge"
- "desc": "Any creature other than a stench kow that starts its turn within 5 feet\
    \ of the stench kow must succeed on a DC 12 Constitution saving throw or be [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)\
    \ until the start of the creature's next turn. On a successful saving throw, the\
    \ creature is immune to the stench of all stench kows for 1 hour."
  "name": "Stench"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+6 (+6) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d6 + 4|text(7) (1d6 + 4) piercing damage."
  "name": "Gore"
"source":
- "VGM"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/VGM/Stench%20Kow.webp"
```
^statblock

```encounter-table
name: Stench Kow
creatures:
 - 1: Stench Kow
```