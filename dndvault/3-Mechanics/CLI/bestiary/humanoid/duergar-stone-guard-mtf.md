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
aliases: ["Duergar Stone Guard"]
NoteIcon: monster
BestiaryType: humanoid (dwarf)
SourceType: Bestiary
BookSource: Mordenkainen's Tome of Foes p. 191
---
# [Duergar Stone Guard](3-Mechanics\CLI\bestiary\humanoid/duergar-stone-guard-mtf.md)
*Source: Mordenkainen's Tome of Foes p. 191*  

## Duergar Stone Guard

The stone guard are elite duergar troops, deployed in small numbers to bolster war bands of regulars or organized into elite strike forces for specific missions.

> [!quote]- A quote from Mordenkainen  
> 
> You think dwarves are dour and unpleasant? Wait until you meet a duergar.

The cruel duergar plot not only to defeat other dwarves, but to cast down the entire dwarven pantheon in revenge for their being abandoned and left to be enslaved by mind flayers. To this end, duergar train warriors to fulfill a variety of roles.

```statblock
"name": "Duergar Stone Guard (MTF)"
"size": "Medium"
"type": "humanoid"
"subtype": "dwarf"
"alignment": "Lawful Evil"
"ac": !!int "18"
"ac_class": "[chain mail](/3-Mechanics/CLI/items/chain-mail.md), [shield](/3-Mechanics/CLI/items/shield.md)"
"hp": !!int "39"
"hit_dice": "6d8 + 12"
"stats":
- !!int "18"
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
- "desc": "The duergar has advantage on saving throws against poison, spells, and\
    \ illusions, as well as to resist being [charmed](/3-Mechanics/CLI/rules/conditions.md#charmed)\
    \ or [paralyzed](/3-Mechanics/CLI/rules/conditions.md#paralyzed)."
  "name": "Duergar Resilience"
- "desc": "The duergar has advantage on attack rolls and Dexterity saving throws while\
    \ standing within 5 feet of a duergar ally wielding a shield."
  "name": "Phalanx Formation"
- "desc": "While in sunlight, the duergar has disadvantage on attack rolls, as well\
    \ as on Wisdom ([Perception](/3-Mechanics/CLI/rules/skills.md#Perception)) checks\
    \ that rely on sight."
  "name": "Sunlight Sensitivity"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+6 (+6) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d6 + 4|text(7) (1d6 + 4) piercing damage, or dice:2d6 + 4|text(11)\
    \ (2d6 + 4) piercing damage while enlarged."
  "name": "King's Knife (Shortsword)"
- "desc": "Melee or Ranged Weapon Attack: dice: d20+6 (+6) to hit, reach 5 ft.\
    \ or range 30/120 ft., one target. Hit: dice:1d6 + 4|text(7) (1d6 + 4) piercing\
    \ damage, or dice:2d6 + 4|text(11) (2d6 + 4) piercing damage while enlarged."
  "name": "Javelin"
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
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MTF/Duergar%20Stone%20Guard.webp"
```
^statblock

```encounter-table
name: Duergar Stone Guard
creatures:
 - 1: Duergar Stone Guard
```

## Environment

mountain, underdark