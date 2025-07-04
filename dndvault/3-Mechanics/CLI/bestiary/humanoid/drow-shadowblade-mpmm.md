---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/11
- monster/environment/underdark
- monster/size/medium
- monster/type/humanoid/elf
aliases: ["Drow Shadowblade"]
NoteIcon: monster
BestiaryType: humanoid (elf)
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 105, Mordenkainen's Tome of Foes p. 187
---
# [Drow Shadowblade](3-Mechanics\CLI\bestiary\humanoid/drow-shadowblade-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 105, Mordenkainen's Tome of Foes p. 187*  

```statblock
"name": "Drow Shadowblade (MPMM)"
"size": "Medium"
"type": "humanoid"
"subtype": "elf"
"alignment": "Any alignment"
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
- "desc": "The drow casts one of the following spells, requiring no material components\
    \ and using Charisma as the spellcasting ability (spell save DC 13):\n\nAt will:\
    \ [dancing lights](/3-Mechanics/CLI/spells/dancing-lights.md), [darkness](/3-Mechanics/CLI/spells/darkness.md)\n\
    \n1/day each: [faerie fire](/3-Mechanics/CLI/spells/faerie-fire.md), [levitate](/3-Mechanics/CLI/spells/levitate.md)\
    \ (self only)"
  "name": "Spellcasting"
- "desc": "Magical darkness doesn't impede the drow's [darkvision](/3-Mechanics/CLI/rules/senses.md#darkvision)."
  "name": "Devil's Sight"
- "desc": "The drow has advantage on saving throws against being [charmed](/3-Mechanics/CLI/rules/conditions.md#charmed),\
    \ and magic can't put the drow to sleep."
  "name": "Fey Ancestry"
- "desc": "While in sunlight, the drow has disadvantage on attack rolls, as well as\
    \ on Wisdom ([Perception](/3-Mechanics/CLI/rules/skills.md#Perception)) checks\
    \ that rely on sight."
  "name": "Sunlight Sensitivity"
"actions":
- "desc": "The drow makes three Shadow Sword attacks. One of the attacks can be replaced\
    \ by a Hand Crossbow attack. The drow can also use Spellcasting to cast darkness."
  "name": "Multiattack"
- "desc": "Melee or Ranged Weapon Attack: dice: d20+9 (+9) to hit, reach 5 ft.\
    \ or range 30/60 ft., one target. Hit: dice:7d6 + 5|text(27) (7d6 + 5) necrotic\
    \ damage."
  "name": "Shadow Sword"
- "desc": "Ranged Weapon Attack: dice: d20+9 (+9) to hit, range 30/120 ft.,\
    \ one target. Hit: dice:1d6 + 5|text(8) (1d6 + 5) piercing damage, and the\
    \ target must succeed on a DC 13 Constitution saving throw or be [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)\
    \ for 1 hour. If the saving throw fails by 5 or more, the target is also [unconscious](/3-Mechanics/CLI/rules/conditions.md#unconscious)\
    \ while [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned) in this way.\
    \ The target regains consciousness if it takes damage or if another creature takes\
    \ an action to shake it."
  "name": "Hand Crossbow"
"bonus_actions":
- "desc": "While in dim light or darkness, the drow teleports, along with any equipment\
    \ it is wearing or carrying, up to 60 feet to an unoccupied space it can see that\
    \ is also in dim light or darkness. It then has advantage on the first melee attack\
    \ it makes before the end of the turn."
  "name": "Shadow Step"
"source":
- "MPMM"
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Drow%20Shadowblade.webp"
```
^statblock

```encounter-table
name: Drow Shadowblade
creatures:
 - 1: Drow Shadowblade
```

## Environment

underdark