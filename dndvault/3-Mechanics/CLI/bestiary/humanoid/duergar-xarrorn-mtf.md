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
aliases: ["Duergar Xarrorn"]
NoteIcon: monster
BestiaryType: humanoid (dwarf)
SourceType: Bestiary
BookSource: Mordenkainen's Tome of Foes p. 193
---
# [Duergar Xarrorn](3-Mechanics\CLI\bestiary\humanoid/duergar-xarrorn-mtf.md)
*Source: Mordenkainen's Tome of Foes p. 193*  

## Duergar Xarrorn

The xarrorn are specialists who construct weapons using a mixture of alchemy and psionics.

> [!quote]- A quote from Mordenkainen  
> 
> You think dwarves are dour and unpleasant? Wait until you meet a duergar.

The cruel duergar plot not only to defeat other dwarves, but to cast down the entire dwarven pantheon in revenge for their being abandoned and left to be enslaved by mind flayers. To this end, duergar train warriors to fulfill a variety of roles.

```statblock
"name": "Duergar Xarrorn (MTF)"
"size": "Medium"
"type": "humanoid"
"subtype": "dwarf"
"alignment": "Lawful Evil"
"ac": !!int "18"
"ac_class": "[plate armor](/3-Mechanics/CLI/items/plate-armor.md)"
"hp": !!int "26"
"hit_dice": "4d8 + 8"
"stats":
- !!int "16"
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
- "desc": "While in sunlight, the duergar has disadvantage on attack rolls, as well\
    \ as on Wisdom ([Perception](/3-Mechanics/CLI/rules/skills.md#Perception)) checks\
    \ that rely on sight."
  "name": "Sunlight Sensitivity"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+5 (+5) to hit (with disadvantage if\
    \ the target is within 5 feet of the duergar), reach 10 ft., one target. Hit:\
    \ dice:1d12 + 3|text(9) (1d12 + 3) piercing damage plus dice:1d6|text(3)\
    \ (1d6) fire damage, or dice:2d12 + 3|text(16) (2d12 + 3) piercing damage\
    \ plus dice:1d6|text(3) (1d6) fire damage while enlarged."
  "name": "Fire Lance"
- "desc": "From its fire lance, the duergar shoots a 15-foot cone of fire or a line\
    \ of fire 30 feet long and 5 feet wide. Each creature in that area must make a\
    \ DC 12 Dexterity saving throw, taking dice:3d6|text(10) (3d6) fire damage\
    \ on a failed save, or half as much damage on a successful one."
  "name": "Fire Spray (Recharge 5-6)"
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
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MTF/Duergar%20Xarrorn.webp"
```
^statblock

```encounter-table
name: Duergar Xarrorn
creatures:
 - 1: Duergar Xarrorn
```

## Environment

mountain, underdark