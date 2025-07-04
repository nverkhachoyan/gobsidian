---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mtf
- monster/cr/11
- monster/environment/underdark
- monster/size/medium
- monster/type/humanoid/elf
aliases: ["Drow Shadowblade"]
NoteIcon: monster
BestiaryType: humanoid (elf)
SourceType: Bestiary
BookSource: Mordenkainen's Tome of Foes p. 187
---
# [Drow Shadowblade](3-Mechanics\CLI\bestiary\humanoid/drow-shadowblade-mtf.md)
*Source: Mordenkainen's Tome of Foes p. 187*  

> [!quote]- A quote from Mordenkainen  
> 
> Many creatures enjoy torture, but the dark elves have made it into an exquisite art.

## Drow Shadowblade

Drow shadowblades steal down the darkened passages of the Underdark, bound on errands of mayhem. Ruthless killers, shadowblades find employment with a noble house, usually involving the elimination of a rival in another house. Shadowblades also protect enclaves and Underdark cities from enemies and track down thieves who make off with prized treasures. In whatever role they serve, they move undetected until the moment they attack. And then they are the last thing their victims see.

A shadowblade harnesses a dark magic that is said to arise from a fiendish ritual in which the drow kills a lesser demon and mystically prevents it from reforming in the Abyss. This ritual creates a shadow demon and infuses the drow with shadow magic.

```statblock
"name": "Drow Shadowblade (MTF)"
"size": "Medium"
"type": "humanoid"
"subtype": "elf"
"alignment": "Neutral Evil"
"ac": !!int "17"
"ac_class": "[studded leather](/3-Mechanics/CLI/items/studded-leather-armor.md)"
"hp": !!int "150"
"hit_dice": "20d8 + 60"
"stats":
- !!int "14"
- !!int "21"
- !!int "16"
- !!int "12"
- !!int "14"
- !!int "13"
"speed": "30 ft."
"saves":
  "Dexterity": !!int "9"
  "Wisdom": !!int "6"
  "Constitution": !!int "7"
"skillsaves":
  "Stealth": !!int "9"
  "Perception": !!int "6"
"senses": "darkvision 120 ft., passive Perception 16"
"languages": "Elvish, Undercommon"
"cr": "11"
"traits":
- "desc": "The drow's innate spellcasting ability is Charisma (spell save DC 13).\
    \ It can innately cast the following spells, requiring no material components:\n\
    \nAt will: [dancing lights](/3-Mechanics/CLI/spells/dancing-lights.md)\n\n\
    1/day each: [darkness](/3-Mechanics/CLI/spells/darkness.md), [faerie fire](/3-Mechanics/CLI/spells/faerie-fire.md),\
    \ [levitate](/3-Mechanics/CLI/spells/levitate.md) (self only)"
  "name": "Innate Spellcasting"
- "desc": "The drow has advantage on saving throws against being [charmed](/3-Mechanics/CLI/rules/conditions.md#charmed),\
    \ and magic can't put the drow to sleep."
  "name": "Fey Ancestry"
- "desc": "While in dim light or darkness, the drow can teleport as a bonus action\
    \ up to 60 feet to an unoccupied space it can see that is also in dim light or\
    \ darkness. It then has advantage on the first melee attack it makes before the\
    \ end of the turn."
  "name": "Shadow Step"
- "desc": "While in sunlight, the drow has disadvantage on attack rolls, as well as\
    \ on Wisdom ([Perception](/3-Mechanics/CLI/rules/skills.md#Perception)) checks\
    \ that rely on sight."
  "name": "Sunlight Sensitivity"
"actions":
- "desc": "The drow makes two attacks with its shadow sword. If either attack hits\
    \ and the target is within 10 feet of a 5-foot cube of darkness created by the\
    \ shadow sword on a previous turn, the drow can dismiss that darkness and cause\
    \ the target to take dice:6d6|text(21) (6d6) necrotic damage. The drow can\
    \ dismiss darkness in this way no more than once per turn."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+9 (+9) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d6 + 5|text(8) (1d6 + 5) piercing damage plus dice:3d6|text(10)\
    \ (3d6) necrotic damage and dice:3d6|text(10) (3d6) poison damage. The drow\
    \ can then fill an unoccupied 5-foot cube within 5 feet of the target with magical\
    \ darkness, which remains for 1 minute."
  "name": "Shadow Sword"
- "desc": "Ranged Weapon Attack: dice: d20+9 (+9) to hit, range 30/120 ft.,\
    \ one target. Hit: dice:1d6 + 5|text(8) (1d6 + 5) piercing damage, and the\
    \ target must succeed on a DC 13 Constitution saving throw or be [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)\
    \ for 1 hour. If the saving throw fails by 5 or more, the target is also [unconscious](/3-Mechanics/CLI/rules/conditions.md#unconscious)\
    \ while [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned) in this way.\
    \ The target regains consciousness if it takes damage or if another creature takes\
    \ an action to shake it."
  "name": "Hand Crossbow"
"source":
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MTF/Drow%20Shadowblade.webp"
```
^statblock

```encounter-table
name: Drow Shadowblade
creatures:
 - 1: Drow Shadowblade
```

## Environment

underdark