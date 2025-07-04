---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mtf
- monster/cr/2
- monster/environment/mountain
- monster/environment/underdark
- monster/size/medium
- monster/type/humanoid/dwarf
aliases: ["Duergar Kavalrachni"]
NoteIcon: monster
BestiaryType: humanoid (dwarf)
SourceType: Bestiary
BookSource: Mordenkainen's Tome of Foes p. 189
---
# [Duergar Kavalrachni](3-Mechanics\CLI\bestiary\humanoid/duergar-kavalrachni-mtf.md)
*Source: Mordenkainen's Tome of Foes p. 189*  

## Duergar Kavalrachni

The kavalrachni are duergar cavalry, trained to fight while riding steeders.

> [!quote]- A quote from Mordenkainen  
> 
> You think dwarves are dour and unpleasant? Wait until you meet a duergar.

The cruel duergar plot not only to defeat other dwarves, but to cast down the entire dwarven pantheon in revenge for their being abandoned and left to be enslaved by mind flayers. To this end, duergar train warriors to fulfill a variety of roles.

```statblock
"name": "Duergar Kavalrachni (MTF)"
"size": "Medium"
"type": "humanoid"
"subtype": "dwarf"
"alignment": "Lawful Evil"
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
- "desc": "When the duergar hits a target with a melee attack while mounted on a female\
    \ steeder, the steeder can make one melee attack against the same target as a\
    \ reaction."
  "name": "Cavalry Training"
- "desc": "The duergar has advantage on saving throws against poison, spells, and\
    \ illusions, as well as to resist being [charmed](/3-Mechanics/CLI/rules/conditions.md#charmed)\
    \ or [paralyzed](/3-Mechanics/CLI/rules/conditions.md#paralyzed)."
  "name": "Duergar Resilience"
- "desc": "While in sunlight, the duergar has disadvantage on attack rolls, as well\
    \ as on Wisdom ([Perception](/3-Mechanics/CLI/rules/skills.md#Perception)) checks\
    \ that rely on sight."
  "name": "Sunlight Sensitivity"
"actions":
- "desc": "The duergar makes two war pick attacks."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+4 (+4) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d8 + 2|text(6) (1d8 + 2) piercing damage plus dice:2d4|text(5)\
    \ (2d4) poison damage."
  "name": "War Pick"
- "desc": "Ranged Weapon Attack: dice: d20+2 (+2) to hit, range 100/400 ft.,\
    \ one target. Hit: dice:1d10|text(5) (1d10) piercing damage."
  "name": "Heavy Crossbow"
- "desc": "The duergar magically turns [invisible](/3-Mechanics/CLI/rules/conditions.md#invisible)\
    \ for up to 1 hour or until it attacks, it casts a spell, or its [concentration](/3-Mechanics/CLI/rules/conditions.md#concentration)\
    \ is broken (as if concentrating on a spell). Any equipment the duergar wears\
    \ or carries is [invisible](/3-Mechanics/CLI/rules/conditions.md#invisible) with\
    \ it. While the [invisible](/3-Mechanics/CLI/rules/conditions.md#invisible) duergar\
    \ is mounted on a female steeder, the steeder is [invisible](/3-Mechanics/CLI/rules/conditions.md#invisible)\
    \ as well. The invisibility ends early on the steeder immediately after it attacks."
  "name": "Shared Invisibility (Recharges after a Short or Long Rest)"
"source":
- "MTF"
- "OotA"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MTF/Duergar%20Kavalrachni.webp"
```
^statblock

```encounter-table
name: Duergar Kavalrachni
creatures:
 - 1: Duergar Kavalrachni
```

## Environment

mountain, underdark