---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/12
- monster/environment/coastal
- monster/environment/desert
- monster/environment/grassland
- monster/environment/mountain
- monster/size/large
- monster/type/celestial
aliases: ["Ki-rin"]
NoteIcon: monster
BestiaryType: celestial
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 162, Volo's Guide to Monsters p. 163
---
# [Ki-rin](3-Mechanics\CLI\bestiary\celestial/ki-rin-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 162, Volo's Guide to Monsters p. 163*  

```statblock
"name": "Ki-rin (MPMM)"
"size": "Large"
"type": "celestial"
"alignment": "Typically  Lawful Good"
"ac": !!int "20"
"ac_class": "natural armor"
"hp": !!int "153"
"hit_dice": "18d10 + 54"
"stats":
- !!int "21"
- !!int "16"
- !!int "16"
- !!int "19"
- !!int "20"
- !!int "20"
"speed": "60 ft., fly 120 ft. (hover)"
"skillsaves":
  "Religion": !!int "8"
  "Perception": !!int "9"
  "Insight": !!int "9"
"damage_immunities": "poison"
"condition_immunities": "[poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)"
"senses": "darkvision 120 ft., truesight 30 ft., passive Perception 19"
"languages": "all, telepathy 120 ft."
"cr": "12"
"traits":
- "desc": "The ki-rin casts one of the following spells, requiring no material components\
    \ and using Wisdom as the spellcasting ability (spell save DC 17):\n\nAt will:\
    \ [light](/3-Mechanics/CLI/spells/light.md), [major image](/3-Mechanics/CLI/spells/major-image.md)\
    \ (6th-level version), [thaumaturgy](/3-Mechanics/CLI/spells/thaumaturgy.md)\n\
    \n1/day each: [banishment](/3-Mechanics/CLI/spells/banishment.md), [calm emotions](/3-Mechanics/CLI/spells/calm-emotions.md),\
    \ [create food and water](/3-Mechanics/CLI/spells/create-food-and-water.md), [greater\
    \ restoration](/3-Mechanics/CLI/spells/greater-restoration.md), [plane shift](/3-Mechanics/CLI/spells/plane-shift.md),\
    \ [protection from evil and good](/3-Mechanics/CLI/spells/protection-from-evil-and-good.md),\
    \ [revivify](/3-Mechanics/CLI/spells/revivify.md), [wind walk](/3-Mechanics/CLI/spells/wind-walk.md)\n\
    \n3/day each: [cure wounds](/3-Mechanics/CLI/spells/cure-wounds.md), [dispel\
    \ magic](/3-Mechanics/CLI/spells/dispel-magic.md), [lesser restoration](/3-Mechanics/CLI/spells/lesser-restoration.md),\
    \ [sending](/3-Mechanics/CLI/spells/sending.md)"
  "name": "Spellcasting"
- "desc": "If the ki-rin fails a saving throw, it can choose to succeed instead."
  "name": "Legendary Resistance (3/Day)"
- "desc": "The ki-rin has advantage on saving throws against spells and other magical\
    \ effects."
  "name": "Magic Resistance"
"actions":
- "desc": "The ki-rin makes two Hoof attacks and one Horn attack, or it makes two\
    \ Sacred Fire attacks."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+9 (+9) to hit, reach 15 ft., one target.\
    \ Hit: dice:2d4 + 5|text(10) (2d4 + 5) force damage."
  "name": "Hoof"
- "desc": "Melee Weapon Attack: dice: d20+9 (+9) to hit, reach 5 ft., one target.\
    \ Hit: dice:2d8 + 5|text(14) (2d8 + 5) radiant damage."
  "name": "Horn"
- "desc": "Ranged Spell Attack: dice: d20+9 (+9) to hit, range 120 ft., one\
    \ target. Hit: dice:3d8 + 5|text(18) (3d8 + 5) radiant damage."
  "name": "Sacred Fire"
"legendary_actions":
- "desc": "The ki-rin moves up to half its speed without provoking opportunity attacks."
  "name": "Move"
- "desc": "The ki-rin makes one Hoof, Horn, or Sacred Fire attack."
  "name": "Smite"
"source":
- "MPMM"
- "VGM"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Ki-rin.webp"
```
^statblock

```encounter-table
name: Ki-rin
creatures:
 - 1: Ki-rin
```

## Environment

coastal, desert, grassland, mountain