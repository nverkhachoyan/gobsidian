---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/7
- monster/environment/forest
- monster/environment/underdark
- monster/environment/urban
- monster/size/medium
- monster/type/humanoid/elf
aliases: ["Shadar-kai Shadow Dancer"]
NoteIcon: monster
BestiaryType: humanoid (elf)
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 213, Mordenkainen's Tome of Foes p. 225
---
# [Shadar-kai Shadow Dancer](3-Mechanics\CLI\bestiary\humanoid/shadar-kai-shadow-dancer-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 213, Mordenkainen's Tome of Foes p. 225*  

```statblock
"name": "Shadar-kai Shadow Dancer (MPMM)"
"size": "Medium"
"type": "humanoid"
"subtype": "elf"
"alignment": "Any alignment"
"ac": !!int "15"
"ac_class": "[studded leather](/3-Mechanics/CLI/items/studded-leather-armor.md)"
"hp": !!int "71"
"hit_dice": "13d8 + 13"
"stats":
- !!int "12"
- !!int "16"
- !!int "13"
- !!int "11"
- !!int "12"
- !!int "12"
"speed": "30 ft."
"saves":
  "Charisma": !!int "4"
  "Dexterity": !!int "6"
"skillsaves":
  "Stealth": !!int "6"
"damage_resistances": "necrotic"
"condition_immunities": "[charmed](/3-Mechanics/CLI/rules/conditions.md#charmed),\
  \ [exhaustion](/3-Mechanics/CLI/rules/conditions.md#exhaustion)"
"senses": "darkvision 60 ft., passive Perception 11"
"languages": "Common, Elvish"
"cr": "7"
"traits":
- "desc": "The shadar-kai has advantage on saving throws against being [charmed](/3-Mechanics/CLI/rules/conditions.md#charmed),\
    \ and magic can't put it to sleep."
  "name": "Fey Ancestry"
"actions":
- "desc": "The shadar-kai makes three Spiked Chain attacks.\n\nIt can use Shadow Jump\
    \ after one of these attacks."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+6 (+6) to hit, reach 10 ft., one target.\
    \ Hit: dice:2d6 + 3|text(10) (2d6 + 3) piercing damage. The target must\
    \ succeed on a DC 14 Dexterity saving throw or suffer one of the following effects\
    \ (choose one or roll a dice: d6|avg|noform (d6)):\n\n- 1–2 Decay. The\
    \ target takes dice:4d10|text(22) (4d10) necrotic damage.  \n- 3–4 Grapple.\
    \ The target is [grappled](/3-Mechanics/CLI/rules/conditions.md#grappled) (escape\
    \ DC 14) if it is a Medium or smaller creature. Until the grapple ends, the target\
    \ is [restrained](/3-Mechanics/CLI/rules/conditions.md#restrained), and the shadar-kai\
    \ can't grapple another target.  \n- 5–6 Topple. The target is knocked [prone](/3-Mechanics/CLI/rules/conditions.md#prone).\
    \  "
  "name": "Spiked Chain"
"bonus_actions":
- "desc": "The shadar-kai teleports, along with any equipment is it wearing or carrying,\
    \ up to 30 feet to an unoccupied space it can see. Both the space it teleports\
    \ from and the space it teleports to must be in dim light or darkness."
  "name": "Shadow Jump"
"source":
- "MPMM"
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Shadar-kai%20Shadow%20Dancer.webp"
```
^statblock

```encounter-table
name: Shadar-kai Shadow Dancer
creatures:
 - 1: Shadar-kai Shadow Dancer
```

## Environment

forest, underdark, urban