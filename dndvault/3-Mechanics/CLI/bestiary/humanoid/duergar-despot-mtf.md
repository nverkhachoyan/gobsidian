---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mtf
- monster/cr/12
- monster/environment/mountain
- monster/environment/underdark
- monster/size/medium
- monster/type/humanoid/dwarf
aliases: ["Duergar Despot"]
NoteIcon: monster
BestiaryType: humanoid (dwarf)
SourceType: Bestiary
BookSource: Mordenkainen's Tome of Foes p. 188
---
# [Duergar Despot](3-Mechanics\CLI\bestiary\humanoid/duergar-despot-mtf.md)
*Source: Mordenkainen's Tome of Foes p. 188*  

## Duergar Despot

Duergar despots replace parts of their bodies with mechanical devices that they control through their psionic abilities.

> [!quote]- A quote from Mordenkainen  
> 
> You think dwarves are dour and unpleasant? Wait until you meet a duergar.

The cruel duergar plot not only to defeat other dwarves, but to cast down the entire dwarven pantheon in revenge for their being abandoned and left to be enslaved by mind flayers. To this end, duergar train warriors to fulfill a variety of roles.

```statblock
"name": "Duergar Despot (MTF)"
"size": "Medium"
"type": "humanoid"
"subtype": "dwarf"
"alignment": "Lawful Evil"
"ac": !!int "21"
"ac_class": "natural armor"
"hp": !!int "119"
"hit_dice": "14d8 + 56"
"stats":
- !!int "20"
- !!int "5"
- !!int "19"
- !!int "15"
- !!int "14"
- !!int "13"
"speed": "25 ft."
"saves":
  "Wisdom": !!int "6"
  "Constitution": !!int "8"
"damage_immunities": "poison"
"condition_immunities": "[charmed](/3-Mechanics/CLI/rules/conditions.md#charmed),\
  \ [exhaustion](/3-Mechanics/CLI/rules/conditions.md#exhaustion), [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened),\
  \ [paralyzed](/3-Mechanics/CLI/rules/conditions.md#paralyzed), [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)"
"senses": "darkvision 120 ft., passive Perception 12"
"languages": "Dwarvish, Undercommon"
"cr": "12"
"traits":
- "desc": "The duergar despot's innate spellcasting ability is Intelligence (spell\
    \ save DC 12). It can cast the following spells, requiring no components:\n\n\
    At will: [mage hand](/3-Mechanics/CLI/spells/mage-hand.md), [minor illusion](/3-Mechanics/CLI/spells/minor-illusion.md)\n\
    \n1/day each: [counterspell](/3-Mechanics/CLI/spells/counterspell.md), [misty\
    \ step](/3-Mechanics/CLI/spells/misty-step.md), [stinking cloud](/3-Mechanics/CLI/spells/stinking-cloud.md)"
  "name": "Innate Spellcasting (Psionics)"
- "desc": "The duergar has advantage on saving throws against spells and other magical\
    \ effects."
  "name": "Magic Resistance"
- "desc": "When the duergar despot suffers a critical hit or is reduced to 0 hit points,\
    \ psychic energy erupts from its frame to deal dice:4d6|text(14) (4d6) psychic\
    \ damage to each creature within 5 feet of it."
  "name": "Psychic Engine"
- "desc": "While in sunlight, the duergar despot has disadvantage on attack rolls,\
    \ as well as on Wisdom ([Perception](/3-Mechanics/CLI/rules/skills.md#Perception))\
    \ checks that rely on sight."
  "name": "Sunlight Sensitivity"
"actions":
- "desc": "The despot makes two iron fist attacks and two stomping foot attacks. It\
    \ can replace up to four of these attacks with uses of its Flame Jet."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+9 (+9) to hit, reach 5 ft., one target.\
    \ Hit: dice:2d8 + 5|text(14) (2d8 + 5) bludgeoning damage. If the target\
    \ is a Large or smaller creature, it must make a successful DC 17 Strength saving\
    \ throw or be thrown up to 30 feet away in a straight line. The target is knocked\
    \ [prone](/3-Mechanics/CLI/rules/conditions.md#prone) and then takes dice:3d6|text(10)\
    \ (3d6) bludgeoning damage."
  "name": "Iron Fist"
- "desc": "Melee Weapon Attack: dice: d20+9 (+9) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d8 + 5|text(9) (1d8 + 5) bludgeoning damage, or dice:3d8 +\
    \ 5|text(18) (3d8 + 5) bludgeoning damage to a [prone](/3-Mechanics/CLI/rules/conditions.md#prone)\
    \ target."
  "name": "Stomping Foot"
- "desc": "The duergar spews flames in a line 100 feet long and 5 feet wide. Each\
    \ creature in the line must make a DC 16 Dexterity saving throw, taking dice:4d8|text(18)\
    \ (4d8) fire damage on a failed save, or half as much damage on a successful\
    \ one."
  "name": "Flame Jet"
"source":
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MTF/Duergar%20Despot.webp"
```
^statblock

```encounter-table
name: Duergar Despot
creatures:
 - 1: Duergar Despot
```

## Environment

mountain, underdark