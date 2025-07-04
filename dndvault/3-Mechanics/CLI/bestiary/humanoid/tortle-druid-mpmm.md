---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/2
- monster/environment/coastal
- monster/size/medium
- monster/type/humanoid
aliases: ["Tortle Druid"]
NoteIcon: monster
BestiaryType: humanoid
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 244, Mordenkainen's Tome of Foes p. 242
---
# [Tortle Druid](3-Mechanics\CLI\bestiary\humanoid/tortle-druid-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 244, Mordenkainen's Tome of Foes p. 242*  

```statblock
"name": "Tortle Druid (MPMM)"
"size": "Medium"
"type": "humanoid"
"alignment": "Any alignment"
"ac": !!int "17"
"ac_class": "natural armor"
"hp": !!int "33"
"hit_dice": "6d8 + 6"
"stats":
- !!int "14"
- !!int "10"
- !!int "12"
- !!int "11"
- !!int "15"
- !!int "12"
"speed": "30 ft."
"skillsaves":
  "Nature": !!int "2"
  "Animal Handling": !!int "4"
  "Survival": !!int "4"
"senses": "passive Perception 12"
"languages": "Aquan, Common"
"cr": "2"
"traits":
- "desc": "The tortle casts one of the following spells, using Wisdom as the spellcasting\
    \ ability (spell save DC 12):\n\nAt will: [druidcraft](/3-Mechanics/CLI/spells/druidcraft.md),\
    \ [guidance](/3-Mechanics/CLI/spells/guidance.md)\n\n2/day each: [cure wounds](/3-Mechanics/CLI/spells/cure-wounds.md),\
    \ [hold person](/3-Mechanics/CLI/spells/hold-person.md), [speak with animals](/3-Mechanics/CLI/spells/speak-with-animals.md),\
    \ [thunderwave](/3-Mechanics/CLI/spells/thunderwave.md)"
  "name": "Spellcasting"
- "desc": "The tortle can hold its breath for 1 hour."
  "name": "Hold Breath"
"actions":
- "desc": "The tortle makes four Claw attacks or two [Nature](/3-Mechanics/CLI/rules/skills.md#Nature)'s\
    \ Wrath attacks."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+4 (+4) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d4 + 2|text(4) (1d4 + 2) slashing damage."
  "name": "Claw"
- "desc": "Ranged Spell Attack: dice: d20+4 (+4) to hit, range 90 ft., one target.\
    \ Hit: dice:2d6 + 2|text(9) (2d6 + 2) damage of a type chosen by the tortle:\
    \ cold, fire, lightning, or thunder."
  "name": "Nature's Wrath"
- "desc": "The tortle withdraws into its shell. Until it emerges, it gains a +4 bonus\
    \ to AC and has advantage on Strength and Constitution saving throws. While in\
    \ its shell, the tortle is [prone](/3-Mechanics/CLI/rules/conditions.md#prone),\
    \ its speed is 0 and can't increase, it has disadvantage on Dexterity saving throws,\
    \ it can't take reactions, and the only action it can take is a bonus action to\
    \ emerge."
  "name": "Shell Defense"
"source":
- "MPMM"
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Tortle%20Druid.webp"
```
^statblock

```encounter-table
name: Tortle Druid
creatures:
 - 1: Tortle Druid
```

## Environment

coastal