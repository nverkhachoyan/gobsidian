---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mtf
- monster/cr/9
- monster/environment/underdark
- monster/size/medium
- monster/type/humanoid/elf
aliases: ["Drow House Captain"]
NoteIcon: monster
BestiaryType: humanoid (elf)
SourceType: Bestiary
BookSource: Mordenkainen's Tome of Foes p. 184
---
# [Drow House Captain](3-Mechanics\CLI\bestiary\humanoid/drow-house-captain-mtf.md)
*Source: Mordenkainen's Tome of Foes p. 184*  

> [!quote]- A quote from Mordenkainen  
> 
> Many creatures enjoy torture, but the dark elves have made it into an exquisite art.

## Drow House Captain

Each drow noble house entrusts the leadership of its military forces to a house captain, a position normally held by the matriarch's first or second son. The house captain commands the drow and slaves making up the family's army and has made extensive study of strategy and tactics to become an effective leader in battle.

```statblock
"name": "Drow House Captain (MTF)"
"size": "Medium"
"type": "humanoid"
"subtype": "elf"
"alignment": "Neutral Evil"
"ac": !!int "16"
"ac_class": "[chain mail](/3-Mechanics/CLI/items/chain-mail.md)"
"hp": !!int "162"
"hit_dice": "25d8 + 50"
"stats":
- !!int "14"
- !!int "19"
- !!int "15"
- !!int "12"
- !!int "14"
- !!int "13"
"speed": "30 ft."
"saves":
  "Dexterity": !!int "8"
  "Wisdom": !!int "6"
  "Constitution": !!int "6"
"skillsaves":
  "Stealth": !!int "8"
  "Perception": !!int "6"
"senses": "darkvision 120 ft., passive Perception 16"
"languages": "Elvish, Undercommon"
"cr": "9"
"traits":
- "desc": "The drow's innate spellcasting ability is Charisma (spell save DC 13).\
    \ He can innately cast the following spells, requiring no material components:\n\
    \nAt will: [dancing lights](/3-Mechanics/CLI/spells/dancing-lights.md)\n\n\
    1/day each: [darkness](/3-Mechanics/CLI/spells/darkness.md), [faerie fire](/3-Mechanics/CLI/spells/faerie-fire.md),\
    \ [levitate](/3-Mechanics/CLI/spells/levitate.md) (self only)"
  "name": "Innate Spellcasting"
- "desc": "As a bonus action, the drow targets one ally he can see within 30 feet\
    \ of him. If the target can see or hear the drow, the target can use its reaction\
    \ to make one melee attack or to take the Dodge or Hide action."
  "name": "Battle Command"
- "desc": "The drow has advantage on saving throws against being [charmed](/3-Mechanics/CLI/rules/conditions.md#charmed),\
    \ and magic can't put the drow to sleep."
  "name": "Fey Ancestry"
- "desc": "While in sunlight, the drow has disadvantage on attack rolls, as well as\
    \ on Wisdom ([Perception](/3-Mechanics/CLI/rules/skills.md#Perception)) checks\
    \ that rely on sight."
  "name": "Sunlight Sensitivity"
"actions":
- "desc": "The drow makes three attacks: two with his scimitar and one with his whip\
    \ or his hand crossbow."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+8 (+8) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d6 + 4|text(7) (1d6 + 4) slashing damage plus dice:4d6|text(14)\
    \ (4d6) poison damage."
  "name": "Scimitar"
- "desc": "Melee Weapon Attack: dice: d20+8 (+8) to hit, reach 10 ft., one target.\
    \ Hit: dice:1d4 + 4|text(6) (1d4 + 4) slashing damage. If the target is\
    \ an ally, it has advantage on attack rolls until the end of its next turn."
  "name": "Whip"
- "desc": "Ranged Weapon Attack: dice: d20+8 (+8) to hit, range 30/120 ft.,\
    \ one target. Hit: dice:1d6 + 4|text(7) (1d6 + 4) piercing damage, and the\
    \ target must succeed on a DC 13 Constitution saving throw or be [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)\
    \ for 1 hour. If the saving throw fails by 5 or more, the target is also [unconscious](/3-Mechanics/CLI/rules/conditions.md#unconscious)\
    \ while [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned) in this way.\
    \ The target regains consciousness if it takes damage or if another creature takes\
    \ an action to shake it."
  "name": "Hand Crossbow"
"reactions":
- "desc": "The drow adds 3 to his AC against one melee attack that would hit him.\
    \ To do so, the drow must see the attacker and be wielding a melee weapon."
  "name": "Parry"
"source":
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MTF/Drow%20House%20Captain.webp"
```
^statblock

```encounter-table
name: Drow House Captain
creatures:
 - 1: Drow House Captain
```

## Environment

underdark