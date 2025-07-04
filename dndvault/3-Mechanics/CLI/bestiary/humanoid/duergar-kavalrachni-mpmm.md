---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/2
- monster/environment/mountain
- monster/environment/underdark
- monster/size/medium
- monster/type/humanoid/dwarf
aliases: ["Duergar Kavalrachni"]
NoteIcon: monster
BestiaryType: humanoid (dwarf)
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 107, Mordenkainen's Tome of Foes p. 189
---
# [Duergar Kavalrachni](3-Mechanics\CLI\bestiary\humanoid/duergar-kavalrachni-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 107, Mordenkainen's Tome of Foes p. 189*  

```statblock
"name": "Duergar Kavalrachni (MPMM)"
"size": "Medium"
"type": "humanoid"
"subtype": "dwarf"
"alignment": "Any alignment"
"ac": !!int "16"
"ac_class": "[scale mail](/3-Mechanics/CLI/items/scale-mail.md), [shield](/3-Mechanics/CLI/items/shield.md)"
"hp": !!int "26"
"hit_dice": "4d8 + 8"
"stats":
- !!int "14"
- !!int "11"
- !!int "14"
- !!int "11"
- !!int "10"
- !!int "9"
"speed": "25 ft."
"damage_resistances": "poison"
"senses": "darkvision 120 ft., passive Perception 10"
"languages": "Dwarvish, Undercommon"
"cr": "2"
"traits":
- "desc": "When the duergar hits a target with a melee attack while mounted, the mount\
    \ can use its reaction to make one melee attack against the same target."
  "name": "Cavalry Training"
- "desc": "The duergar has advantage on saving throws against spells and the [charmed](/3-Mechanics/CLI/rules/conditions.md#charmed),\
    \ [paralyzed](/3-Mechanics/CLI/rules/conditions.md#paralyzed), and [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)\
    \ conditions."
  "name": "Duergar Resilience"
- "desc": "While in sunlight, the duergar has disadvantage on attack rolls, as well\
    \ as on Wisdom ([Perception](/3-Mechanics/CLI/rules/skills.md#Perception)) checks\
    \ that rely on sight."
  "name": "Sunlight Sensitivity"
"actions":
- "desc": "The duergar makes two War Pick attacks."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+4 (+4) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d8 + 2|text(6) (1d8 + 2) piercing damage plus dice:2d4|text(5)\
    \ (2d4) poison damage."
  "name": "War Pick"
- "desc": "Ranged Weapon Attack: dice: d20+2 (+2) to hit, range 100/400 ft.,\
    \ one target. Hit: dice:1d10|text(5) (1d10) piercing damage."
  "name": "Heavy Crossbow"
- "desc": "The duergar magically turns [invisible](/3-Mechanics/CLI/rules/conditions.md#invisible)\
    \ for up to 1 hour or until it attacks, it forces a creature to make a saving\
    \ throw, or its [concentration](/3-Mechanics/CLI/rules/conditions.md#concentration)\
    \ is broken (as if concentrating on a spell). Any equipment the duergar wears\
    \ or carries is [invisible](/3-Mechanics/CLI/rules/conditions.md#invisible) with\
    \ it. While the [invisible](/3-Mechanics/CLI/rules/conditions.md#invisible) duergar\
    \ is mounted, the mount is [invisible](/3-Mechanics/CLI/rules/conditions.md#invisible)\
    \ as well. The invisibility ends early on the mount immediately after it attacks."
  "name": "Shared Invisibility (Recharges after a Short or Long Rest)"
"source":
- "MPMM"
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Duergar%20Kavalrachni.webp"
```
^statblock

```encounter-table
name: Duergar Kavalrachni
creatures:
 - 1: Duergar Kavalrachni
```

## Environment

mountain, underdark