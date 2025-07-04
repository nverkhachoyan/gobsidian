---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/6
- monster/environment/mountain
- monster/environment/underdark
- monster/size/medium
- monster/type/humanoid/dwarf
aliases: ["Duergar Warlord"]
NoteIcon: monster
BestiaryType: humanoid (dwarf)
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 111, Mordenkainen's Tome of Foes p. 192
---
# [Duergar Warlord](3-Mechanics\CLI\bestiary\humanoid/duergar-warlord-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 111, Mordenkainen's Tome of Foes p. 192*  

```statblock
"name": "Duergar Warlord (MPMM)"
"size": "Medium"
"type": "humanoid"
"subtype": "dwarf"
"alignment": "Any alignment"
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
- "desc": "The duergar has advantage on saving throws against spells and the [charmed](/3-Mechanics/CLI/rules/conditions.md#charmed),\
    \ [paralyzed](/3-Mechanics/CLI/rules/conditions.md#paralyzed), and [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)\
    \ conditions."
  "name": "Duergar Resilience"
- "desc": "While in sunlight, the duergar has disadvantage on attack rolls, as well\
    \ as on Wisdom ([Perception](/3-Mechanics/CLI/rules/skills.md#Perception)) checks\
    \ that rely on sight."
  "name": "Sunlight Sensitivity"
"actions":
- "desc": "The duergar makes three Psychic-Attuned Hammer or Javelin attacks and uses\
    \ Call to Attack."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+7 (+7) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d10 + 4|text(9) (1d10 + 4) bludgeoning damage, or dice:2d10\
    \ + 4|text(15) (2d10 + 4) bludgeoning damage while under the effect of Enlarge,\
    \ plus dice:1d10|text(5) (1d10) psychic damage."
  "name": "Psychic-Attuned Hammer"
- "desc": "Melee or Ranged Weapon Attack: dice: d20+7 (+7) to hit, reach 5 ft.\
    \ or range 30/120 ft., one target. Hit: dice:1d6 + 4|text(7) (1d6 + 4) piercing\
    \ damage, or dice:2d6 + 4|text(11) (2d6 + 4) piercing damage while under the\
    \ effect of Enlarge."
  "name": "Javelin"
- "desc": "Up to three allies within 120 feet of this duergar that can hear it can\
    \ each use their reaction to make one weapon attack."
  "name": "Call to Attack"
- "desc": "The duergar magically turns [invisible](/3-Mechanics/CLI/rules/conditions.md#invisible)\
    \ for up to 1 hour or until it attacks, it forces a creature to make a saving\
    \ throw, or its [concentration](/3-Mechanics/CLI/rules/conditions.md#concentration)\
    \ is broken (as if concentrating on a spell). Any equipment the duergar wears\
    \ or carries is [invisible](/3-Mechanics/CLI/rules/conditions.md#invisible) with\
    \ it."
  "name": "Invisibility (Recharge 4-6)"
"bonus_actions":
- "desc": "For 1 minute, the duergar magically increases in size, along with anything\
    \ it is wearing or carrying. While enlarged, the duergar is Large, doubles its\
    \ damage dice on Strength-based weapon attacks (included in the attacks), and\
    \ makes Strength checks and Strength saving throws with advantage. If the duergar\
    \ lacks the room to become Large, it attains the maximum size possible in the\
    \ space available."
  "name": "Enlarge (Recharges after a Short or Long Rest)"
"reactions":
- "desc": "When an ally that the duergar can see makes a dice: d20|avg|noform (d20)\
    \ roll, the duergar can roll a dice: d6|avg|noform (d6), and the ally can\
    \ add the number rolled to the dice: d20|avg|noform (d20) by taking dice:1d6|text(3)\
    \ (1d6) psychic damage."
  "name": "Scouring Instruction"
"source":
- "MPMM"
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Duergar%20Warlord.webp"
```
^statblock

```encounter-table
name: Duergar Warlord
creatures:
 - 1: Duergar Warlord
```

## Environment

mountain, underdark