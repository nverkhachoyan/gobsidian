---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/vgm
- monster/cr/1-4
- monster/environment/underdark
- monster/size/medium
- monster/type/beast
aliases: ["Deep Rothé"]
NoteIcon: monster
BestiaryType: beast
SourceType: Bestiary
BookSource: Volo's Guide to Monsters p. 208
---
# [Deep Rothé](3-Mechanics\CLI\bestiary\beast/deep-rothe-vgm.md)
*Source: Volo's Guide to Monsters p. 208*  

Deep rothé are stunted Underdark variants of [rothé](/3-Mechanics/CLI/bestiary/beast/rothe-vgm.md).

```statblock
"name": "Deep Rothé (VGM)"
"size": "Medium"
"type": "beast"
"alignment": "Unaligned"
"ac": !!int "10"
"hp": !!int "13"
"hit_dice": "2d8 + 4"
"stats":
- !!int "18"
- !!int "10"
- !!int "14"
- !!int "2"
- !!int "10"
- !!int "4"
"speed": "30 ft."
"senses": "darkvision 60 ft., passive Perception 10"
"languages": ""
"cr": "1/4"
"traits":
- "desc": "The deep rothé's spellcasting ability is Charisma. It can innately cast\
    \ [dancing lights](/3-Mechanics/CLI/spells/dancing-lights.md) at will, requiring\
    \ no components.\n\nAt will: [dancing lights](/3-Mechanics/CLI/spells/dancing-lights.md)"
  "name": "Innate Spellcasting"
- "desc": "If the rothé moves at least 20 feet straight toward a target and then hits\
    \ it with a gore attack on the same turn, the target takes an extra dice:2d6|text(7)\
    \ (2d6) piercing damage."
  "name": "Charge"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+6 (+6) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d6 + 4|text(7) (1d6 + 4) piercing damage."
  "name": "Gore"
"source":
- "VGM"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/VGM/Deep%20Rothe.webp"
```
^statblock

```encounter-table
name: Deep Rothé
creatures:
 - 1: Deep Rothé
```

## Environment

underdark