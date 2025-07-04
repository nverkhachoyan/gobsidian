---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/7
- monster/environment/underdark
- monster/size/large
- monster/type/fiend/demon
aliases: ["Draegloth"]
NoteIcon: monster
BestiaryType: fiend (demon)
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 98, Volo's Guide to Monsters p. 141
---
# [Draegloth](3-Mechanics\CLI\bestiary\fiend/draegloth-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 98, Volo's Guide to Monsters p. 141*  

```statblock
"name": "Draegloth (MPMM)"
"size": "Large"
"type": "fiend"
"subtype": "demon"
"alignment": "Typically  Chaotic Evil"
"ac": !!int "15"
"ac_class": "natural armor"
"hp": !!int "123"
"hit_dice": "13d10 + 52"
"stats":
- !!int "20"
- !!int "15"
- !!int "18"
- !!int "13"
- !!int "11"
- !!int "11"
"speed": "30 ft."
"skillsaves":
  "Stealth": !!int "5"
  "Perception": !!int "3"
"damage_resistances": "cold, fire, lightning"
"damage_immunities": "poison"
"condition_immunities": "[poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)"
"senses": "darkvision 120 ft., passive Perception 13"
"languages": "Abyssal, Elvish, Undercommon"
"cr": "7"
"traits":
- "desc": "The draegloth casts one of the following spells, requiring no material\
    \ components and using Charisma as the spellcasting ability (spell save DC 11):\n\
    \nAt will: [dancing lights](/3-Mechanics/CLI/spells/dancing-lights.md), [darkness](/3-Mechanics/CLI/spells/darkness.md)\n\
    \n1/day each: [confusion](/3-Mechanics/CLI/spells/confusion.md), [faerie fire](/3-Mechanics/CLI/spells/faerie-fire.md)"
  "name": "Spellcasting"
- "desc": "The draegloth has advantage on saving throws against being [charmed](/3-Mechanics/CLI/rules/conditions.md#charmed),\
    \ and magic can't put it to sleep."
  "name": "Fey Ancestry"
"actions":
- "desc": "The draegloth makes one Bite attack and two Claw attacks."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+8 (+8) to hit, reach 5 ft., one creature.\
    \ Hit: dice:2d10 + 5|text(16) (2d10 + 5) piercing damage."
  "name": "Bite"
- "desc": "Melee Weapon Attack: dice: d20+8 (+8) to hit, reach 10 ft., one target.\
    \ Hit: dice:2d10 + 5|text(16) (2d10 + 5) slashing damage."
  "name": "Claw"
"source":
- "MPMM"
- "VGM"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Draegloth.webp"
```
^statblock

```encounter-table
name: Draegloth
creatures:
 - 1: Draegloth
```

## Environment

underdark