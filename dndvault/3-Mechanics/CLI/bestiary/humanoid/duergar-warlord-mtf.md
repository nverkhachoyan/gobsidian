---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mtf
- monster/cr/6
- monster/environment/mountain
- monster/environment/underdark
- monster/size/medium
- monster/type/humanoid/dwarf
aliases: ["Duergar Warlord"]
NoteIcon: monster
BestiaryType: humanoid (dwarf)
SourceType: Bestiary
BookSource: Mordenkainen's Tome of Foes p. 192
---
# [Duergar Warlord](3-Mechanics\CLI\bestiary\humanoid/duergar-warlord-mtf.md)
*Source: Mordenkainen's Tome of Foes p. 192*  

## Duergar Warlord

A warlord is cunning, inspiring, and cruel in equal parts. A skilled warrior who leads duergar into battle, the warlord can use spikes of psionic energy to compel duergar warriors to fight harder.

> [!quote]- A quote from Mordenkainen  
> 
> You think dwarves are dour and unpleasant? Wait until you meet a duergar.

The cruel duergar plot not only to defeat other dwarves, but to cast down the entire dwarven pantheon in revenge for their being abandoned and left to be enslaved by mind flayers. To this end, duergar train warriors to fulfill a variety of roles.

```statblock
"name": "Duergar Warlord (MTF)"
"size": "Medium"
"type": "humanoid"
"subtype": "dwarf"
"alignment": "Lawful Evil"
"ac": !!int "20"
"ac_class": "[plate armor](/3-Mechanics/CLI/items/plate-armor.md), [shield](/3-Mechanics/CLI/items/shield.md)"
"hp": !!int "75"
"hit_dice": "10d8 + 30"
"stats":
- !!int "18"
- !!int "11"
- !!int "17"
- !!int "12"
- !!int "12"
- !!int "14"
"speed": "25 ft."
"damage_resistances": "poison"
"senses": "darkvision 120 ft., passive Perception 11"
"languages": "Dwarvish, Undercommon"
"cr": "6"
"traits":
- "desc": "The duergar has advantage on saving throws against poison, spells, and\
    \ illusions, as well as to resist being [charmed](/3-Mechanics/CLI/rules/conditions.md#charmed)\
    \ or [paralyzed](/3-Mechanics/CLI/rules/conditions.md#paralyzed)."
  "name": "Duergar Resilience"
- "desc": "While in sunlight, the duergar has disadvantage on attack rolls, as well\
    \ as on Wisdom ([Perception](/3-Mechanics/CLI/rules/skills.md#Perception)) checks\
    \ that rely on sight."
  "name": "Sunlight Sensitivity"
"actions":
- "desc": "The duergar makes three hammer or javelin attacks and uses Call to Attack,\
    \ or Enlarge if it is available."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+7 (+7) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d10 + 4|text(9) (1d10 + 4) bludgeoning damage plus dice:1d10|text(5)\
    \ (1d10) psychic damage, or dice:2d10 + 4|text(15) (2d10 + 4) bludgeoning\
    \ damage plus dice:1d10|text(5) (1d10) psychic damage while enlarged,"
  "name": "Psychic-Attuned Hammer"
- "desc": "Melee or Ranged Weapon Attack: dice: d20+7 (+7) to hit, reach 5 ft.\
    \ or range 30/120 ft., one target. Hit: dice:1d6 + 4|text(7) (1d6 + 4) piercing\
    \ damage, or dice:2d6 + 4|text(11) (2d6 + 4) piercing damage while enlarged."
  "name": "Javelin"
- "desc": "Up to three allied duergar within 120 feet of this duergar that can hear\
    \ it can each use their reaction to make one weapon attack."
  "name": "Call to Attack"
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
  "name": "Invisibility (Recharge 4-6)"
"reactions":
- "desc": "When an ally that the duergar can see makes a dice: d20|avg|noform (d20)\
    \ roll, the duergar can roll a dice: 1d6|avg|noform (1d6) and the ally can\
    \ add the number rolled to the dice: d20|avg|noform (d20) roll by taking dice:1d6|text(3)\
    \ (1d6) psychic damage. A creature immune to psychic damage can't be affected\
    \ by Scouring Instruction."
  "name": "Scouring Instruction"
"source":
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MTF/Duergar%20Warlord.webp"
```
^statblock

```encounter-table
name: Duergar Warlord
creatures:
 - 1: Duergar Warlord
```

## Environment

mountain, underdark