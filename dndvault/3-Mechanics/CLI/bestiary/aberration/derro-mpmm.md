---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/1-4
- monster/environment/mountain
- monster/environment/underdark
- monster/size/small
- monster/type/aberration
aliases: ["Derro"]
NoteIcon: monster
BestiaryType: aberration
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 91, Mordenkainen's Tome of Foes p. 158
---
# [Derro](3-Mechanics\CLI\bestiary\aberration/derro-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 91, Mordenkainen's Tome of Foes p. 158*  

```statblock
"name": "Derro (MPMM)"
"size": "Small"
"type": "aberration"
"alignment": "Typically  Chaotic Evil"
"ac": !!int "13"
"ac_class": "[leather armor](/3-Mechanics/CLI/items/leather-armor.md)"
"hp": !!int "13"
"hit_dice": "3d6 + 3"
"stats":
- !!int "10"
- !!int "14"
- !!int "12"
- !!int "11"
- !!int "5"
- !!int "9"
"speed": "30 ft."
"skillsaves":
  "Stealth": !!int "4"
"senses": "darkvision 120 ft., passive Perception 7"
"languages": "Dwarvish, Undercommon"
"cr": "1/4"
"traits":
- "desc": "The derro has advantage on saving throws against spells and other magical\
    \ effects."
  "name": "Magic Resistance"
- "desc": "While in sunlight, the derro has disadvantage on attack rolls, as well\
    \ as on Wisdom ([Perception](/3-Mechanics/CLI/rules/skills.md#Perception)) checks\
    \ that rely on sight."
  "name": "Sunlight Sensitivity"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+2 (+2) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d6|text(3) (1d6) piercing damage. If the target is Medium or\
    \ smaller, the derro can choose to deal no damage and knock it [prone](/3-Mechanics/CLI/rules/conditions.md#prone)."
  "name": "Hooked Spear"
- "desc": "Ranged Weapon Attack: dice: d20+4 (+4) to hit, range 80/320 ft.,\
    \ one target. Hit: dice:1d8 + 2|text(6) (1d8 + 2) piercing damage."
  "name": "Light Crossbow"
"source":
- "MPMM"
- "MTF"
- "QftIS"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Derro.webp"
```
^statblock

```encounter-table
name: Derro
creatures:
 - 1: Derro
```

## Environment

mountain, underdark