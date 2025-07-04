---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/1-2
- monster/environment/underdark
- monster/size/small
- monster/type/monstrosity
aliases: ["Chitine"]
NoteIcon: monster
BestiaryType: monstrosity
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 75, Volo's Guide to Monsters p. 131
---
# [Chitine](3-Mechanics\CLI\bestiary\monstrosity/chitine-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 75, Volo's Guide to Monsters p. 131*  

```statblock
"name": "Chitine (MPMM)"
"size": "Small"
"type": "monstrosity"
"alignment": "Typically  Chaotic Evil"
"ac": !!int "14"
"ac_class": "[hide armor](/3-Mechanics/CLI/items/hide-armor.md)"
"hp": !!int "18"
"hit_dice": "4d6 + 4"
"stats":
- !!int "10"
- !!int "14"
- !!int "12"
- !!int "10"
- !!int "10"
- !!int "7"
"speed": "30 ft., climb 30 ft."
"skillsaves":
  "Athletics": !!int "4"
  "Stealth": !!int "4"
"senses": "darkvision 60 ft., passive Perception 10"
"languages": "Undercommon"
"cr": "1/2"
"traits":
- "desc": "The chitine has advantage on saving throws against being [charmed](/3-Mechanics/CLI/rules/conditions.md#charmed),\
    \ and magic can't put the chitine to sleep."
  "name": "Fey Ancestry"
- "desc": "While in sunlight, the chitine has disadvantage on attack rolls, as well\
    \ as on Wisdom ([Perception](/3-Mechanics/CLI/rules/skills.md#Perception)) checks\
    \ that rely on sight."
  "name": "Sunlight Sensitivity"
- "desc": "While in contact with a web, the chitine knows the exact location of any\
    \ other creature in contact with the same web."
  "name": "Web Sense"
- "desc": "The chitine ignores movement restrictions caused by webbing."
  "name": "Web Walker"
"actions":
- "desc": "The chitine makes three Dagger attacks."
  "name": "Multiattack"
- "desc": "Melee or Ranged Weapon Attack: dice: d20+4 (+4) to hit, reach 5 ft.\
    \ or range 20/60 ft., one target. Hit: dice:1d4 + 2|text(4) (1d4 + 2) piercing\
    \ damage."
  "name": "Dagger"
"source":
- "MPMM"
- "VGM"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Chitine.webp"
```
^statblock

```encounter-table
name: Chitine
creatures:
 - 1: Chitine
```

## Environment

underdark