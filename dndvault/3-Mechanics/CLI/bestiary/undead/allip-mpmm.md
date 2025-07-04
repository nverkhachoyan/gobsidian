---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/5
- monster/environment/swamp
- monster/environment/urban
- monster/size/medium
- monster/type/undead
aliases: ["Allip"]
NoteIcon: monster
BestiaryType: undead
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 45, Mordenkainen's Tome of Foes p. 116
---
# [Allip](3-Mechanics\CLI\bestiary\undead/allip-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 45, Mordenkainen's Tome of Foes p. 116*  

```statblock
"name": "Allip (MPMM)"
"size": "Medium"
"type": "undead"
"alignment": "Typically  Neutral Evil"
"ac": !!int "13"
"hp": !!int "40"
"hit_dice": "9d8"
"stats":
- !!int "6"
- !!int "17"
- !!int "10"
- !!int "17"
- !!int "15"
- !!int "16"
"speed": "0 ft., fly 40 ft. (hover)"
"saves":
  "Wisdom": !!int "5"
  "Intelligence": !!int "6"
"skillsaves":
  "Stealth": !!int "6"
  "Perception": !!int "5"
"damage_resistances": "acid; fire; lightning; thunder; bludgeoning, piercing, slashing\
  \ from nonmagical attacks"
"damage_immunities": "cold, necrotic, poison"
"condition_immunities": "[charmed](/3-Mechanics/CLI/rules/conditions.md#charmed),\
  \ [exhaustion](/3-Mechanics/CLI/rules/conditions.md#exhaustion), [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened),\
  \ [grappled](/3-Mechanics/CLI/rules/conditions.md#grappled), [paralyzed](/3-Mechanics/CLI/rules/conditions.md#paralyzed),\
  \ [petrified](/3-Mechanics/CLI/rules/conditions.md#petrified), [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned),\
  \ [prone](/3-Mechanics/CLI/rules/conditions.md#prone), [restrained](/3-Mechanics/CLI/rules/conditions.md#restrained)"
"senses": "darkvision 60 ft., passive Perception 15"
"languages": "the languages it knew in life"
"cr": "5"
"traits":
- "desc": "The allip can move through other creatures and objects as if they were\
    \ difficult terrain. It takes dice:1d10|text(5) (1d10) force damage if it\
    \ ends its turn inside an object."
  "name": "Incorporeal Movement"
- "desc": "The allip doesn't require air, food, drink, or sleep."
  "name": "Unusual Nature"
"actions":
- "desc": "Melee Spell Attack: dice: d20+6 (+6) to hit, reach 5 ft., one target.\
    \ Hit: dice:4d6 + 3|text(17) (4d6 + 3) psychic damage."
  "name": "Maddening Touch"
- "desc": "Each creature within 30 feet of the allip that can hear it must make a\
    \ DC 14 Wisdom saving throw. On a failed save, a target takes dice:2d8 + 3|text(12)\
    \ (2d8 + 3) psychic damage, and it is [stunned](/3-Mechanics/CLI/rules/conditions.md#stunned)\
    \ until the end of its next turn. On a successful save, it takes half as much\
    \ damage and isn't [stunned](/3-Mechanics/CLI/rules/conditions.md#stunned). Constructs\
    \ and Undead are immune to this effect."
  "name": "Howling Babble (Recharge 6)"
- "desc": "The allip chooses up to three creatures it can see within 60 feet of it.\
    \ Each target must succeed on a DC 14 Wisdom saving throw, or it takes dice:2d8\
    \ + 3|text(12) (2d8 + 3) psychic damage and must use its reaction to make a\
    \ melee weapon attack against one creature of the allip's choice that the allip\
    \ can see. Constructs and Undead are immune to this effect."
  "name": "Whispers of Compulsion"
"source":
- "MPMM"
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Allip.webp"
```
^statblock

```encounter-table
name: Allip
creatures:
 - 1: Allip
```

## Environment

swamp, urban