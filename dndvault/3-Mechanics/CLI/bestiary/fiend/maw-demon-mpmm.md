---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/1
- monster/environment/underdark
- monster/size/medium
- monster/type/fiend/demon
aliases: ["Maw Demon"]
NoteIcon: monster
BestiaryType: fiend (demon)
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 176, Volo's Guide to Monsters p. 137
---
# [Maw Demon](3-Mechanics\CLI\bestiary\fiend/maw-demon-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 176, Volo's Guide to Monsters p. 137*  

```statblock
"name": "Maw Demon (MPMM)"
"size": "Medium"
"type": "fiend"
"subtype": "demon"
"alignment": "Typically  Chaotic Evil"
"ac": !!int "13"
"ac_class": "natural armor"
"hp": !!int "33"
"hit_dice": "6d8 + 6"
"stats":
- !!int "14"
- !!int "8"
- !!int "13"
- !!int "5"
- !!int "8"
- !!int "5"
"speed": "30 ft."
"damage_resistances": "cold, fire, lightning"
"damage_immunities": "poison"
"condition_immunities": "[charmed](/3-Mechanics/CLI/rules/conditions.md#charmed),\
  \ [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened), [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)"
"senses": "darkvision 60 ft., passive Perception 9"
"languages": "understands Abyssal but can't speak"
"cr": "1"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+4 (+4) to hit, reach 5 ft., one target.\
    \ Hit: dice:2d8 + 2|text(11) (2d8 + 2) piercing damage."
  "name": "Bite"
- "desc": "The demon vomits in a 15-foot cube. Each creature in that area must succeed\
    \ on a DC 11 Dexterity saving throw or take dice:2d10|text(11) (2d10) acid\
    \ damage and fall [prone](/3-Mechanics/CLI/rules/conditions.md#prone) in the spew."
  "name": "Disgorge (Recharge 6)"
"source":
- "MPMM"
- "VGM"
- "BMT"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Maw%20Demon.webp"
```
^statblock

```encounter-table
name: Maw Demon
creatures:
 - 1: Maw Demon
```

## Environment

underdark