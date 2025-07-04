---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mtf
- monster/cr/1
- monster/environment/mountain
- monster/environment/underdark
- monster/size/medium
- monster/type/humanoid/dwarf
aliases: ["Duergar Soulblade"]
NoteIcon: monster
BestiaryType: humanoid (dwarf)
SourceType: Bestiary
BookSource: Mordenkainen's Tome of Foes p. 190
---
# [Duergar Soulblade](3-Mechanics\CLI\bestiary\humanoid/duergar-soulblade-mtf.md)
*Source: Mordenkainen's Tome of Foes p. 190*  

## Duergar Soulblade

Soulblades are duergar warriors whose mastery of psionics allows them to manifest blades of psychic energy to slice apart their foes.

> [!quote]- A quote from Mordenkainen  
> 
> You think dwarves are dour and unpleasant? Wait until you meet a duergar.

The cruel duergar plot not only to defeat other dwarves, but to cast down the entire dwarven pantheon in revenge for their being abandoned and left to be enslaved by mind flayers. To this end, duergar train warriors to fulfill a variety of roles.

```statblock
"name": "Duergar Soulblade (MTF)"
"size": "Medium"
"type": "humanoid"
"subtype": "dwarf"
"alignment": "Lawful Evil"
"ac": !!int "14"
"ac_class": "[leather armor](/3-Mechanics/CLI/items/leather-armor.md)"
"hp": !!int "18"
"hit_dice": "4d8"
"stats":
- !!int "11"
- !!int "16"
- !!int "10"
- !!int "11"
- !!int "10"
- !!int "12"
"speed": "25 ft."
"damage_resistances": "poison"
"senses": "darkvision 120 ft., passive Perception 10"
"languages": "Dwarvish, Undercommon"
"cr": "1"
"traits":
- "desc": "The duergar's innate spellcasting ability is Wisdom (spell save DC 12,\
    \ dice: d20+4 (+4) to hit with spell attacks). It can innately cast the following\
    \ spells, requiring no components:\n\nAt will: [blade ward](/3-Mechanics/CLI/spells/blade-ward.md),\
    \ [true strike](/3-Mechanics/CLI/spells/true-strike.md)\n\n3/day each: [jump](/3-Mechanics/CLI/spells/jump.md),\
    \ [hunter's mark](/3-Mechanics/CLI/spells/hunters-mark.md)"
  "name": "Innate Spellcasting (Psionics)"
- "desc": "The duergar has advantage on saving throws against poison, spells, and\
    \ illusions, as well as to resist being [charmed](/3-Mechanics/CLI/rules/conditions.md#charmed)\
    \ or [paralyzed](/3-Mechanics/CLI/rules/conditions.md#paralyzed)."
  "name": "Duergar Resilience"
- "desc": "As a bonus action, the duergar can create a shortsword-sized, visible blade\
    \ of psionic energy. The weapon appears in the duergar's hand and vanishes if\
    \ it leaves the duergar's grip, or if the duergar dies or is [incapacitated](/3-Mechanics/CLI/rules/conditions.md#incapacitated)."
  "name": "Create Soulblade"
- "desc": "While in sunlight, the duergar has disadvantage on attack rolls, as well\
    \ as on Wisdom ([Perception](/3-Mechanics/CLI/rules/skills.md#Perception)) checks\
    \ that rely on sight."
  "name": "Sunlight Sensitivity"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+5 (+5) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d6 + 3|text(6) (1d6 + 3) force damage, or dice:2d6 + 3|text(10)\
    \ (2d6 + 3) force damage while enlarged. If the soulblade has advantage on the\
    \ attack roll, the attack deals an extra dice:1d6|text(3) (1d6) force damage."
  "name": "Soulblade"
- "desc": "For 1 minute, the duergar magically increases in size, along with anything\
    \ it is wearing or carrying. While enlarged, the duergar is Large, doubles its\
    \ damage dice on Strength-based weapon attacks (included in the attacks), and\
    \ makes Strength checks and Strength saving throws with advantage. If the duergar\
    \ lacks the room to become Large, it attains the maximum size possible in the\
    \ space available."
  "name": "Enlarge (Recharges after a Short or Long Rest)"
- "desc": "The duergar magically turns [invisible](/3-Mechanics/CLI/rules/conditions.md#invisible)\
    \ for up to 1 hour or until it attacks, it casts a spell, it uses its Enlarge,\
    \ or its [concentration](/3-Mechanics/CLI/rules/conditions.md#concentration) is\
    \ broken (as if concentrating on a spell). Any equipment the duergar wears or\
    \ carries is [invisible](/3-Mechanics/CLI/rules/conditions.md#invisible) with\
    \ it."
  "name": "Invisibility (Recharges after a Short or Long Rest)"
"source":
- "MTF"
- "OotA"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MTF/Duergar%20Soulblade.webp"
```
^statblock

```encounter-table
name: Duergar Soulblade
creatures:
 - 1: Duergar Soulblade
```

## Environment

mountain, underdark