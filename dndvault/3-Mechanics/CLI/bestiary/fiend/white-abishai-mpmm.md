---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/6
- monster/environment/urban
- monster/size/medium
- monster/type/fiend/devil
aliases: ["White Abishai"]
NoteIcon: monster
BestiaryType: fiend (devil)
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 41, Mordenkainen's Tome of Foes p. 163
---
# [White Abishai](3-Mechanics\CLI\bestiary\fiend/white-abishai-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 41, Mordenkainen's Tome of Foes p. 163*  

```statblock
"name": "White Abishai (MPMM)"
"size": "Medium"
"type": "fiend"
"subtype": "devil"
"alignment": "Typically  Lawful Evil"
"ac": !!int "15"
"ac_class": "natural armor"
"hp": !!int "68"
"hit_dice": "8d8 + 32"
"stats":
- !!int "16"
- !!int "11"
- !!int "18"
- !!int "11"
- !!int "12"
- !!int "13"
"speed": "30 ft., fly 40 ft."
"saves":
  "Strength": !!int "6"
  "Constitution": !!int "7"
"damage_resistances": "bludgeoning, piercing, slashing from nonmagical attacks that\
  \ aren't silvered"
"damage_immunities": "cold, fire, poison"
"condition_immunities": "[poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)"
"senses": "darkvision 120 ft., passive Perception 11"
"languages": "Draconic, Infernal, telepathy 120 ft."
"cr": "6"
"traits":
- "desc": "Magical darkness doesn't impede the abishai's [darkvision](/3-Mechanics/CLI/rules/senses.md#darkvision)."
  "name": "Devil's Sight"
- "desc": "The abishai has advantage on saving throws against spells and other magical\
    \ effects."
  "name": "Magic Resistance"
- "desc": "At the start of its turn, the abishai can gain advantage on all melee weapon\
    \ attack rolls during that turn, but attack rolls against it have advantage until\
    \ the start of its next turn."
  "name": "Reckless"
"actions":
- "desc": "The abishai makes one Bite attack, one Claw attack, and one Longsword attack."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+6 (+6) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d4 + 3|text(5) (1d4 + 3) piercing damage plus dice:1d6|text(3)\
    \ (1d6) cold damage."
  "name": "Bite"
- "desc": "Melee Weapon Attack: dice: d20+6 (+6) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d10 + 3|text(8) (1d10 + 3) slashing damage."
  "name": "Claw"
- "desc": "Melee Weapon Attack: dice: d20+6 (+6) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d8 + 3|text(7) (1d8 + 3) force damage, or dice:1d10 + 3|text(8)\
    \ (1d10 + 3) force damage if used with two hands."
  "name": "Longsword"
"reactions":
- "desc": "In response to taking damage, the abishai makes one Bite attack against\
    \ a random creature within 5 feet of it. If no creature is within reach, the abishai\
    \ moves up to half its speed toward an enemy it can see, without provoking opportunity\
    \ attacks."
  "name": "Vicious Reprisal"
"source":
- "MPMM"
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/White%20Abishai.webp"
```
^statblock

```encounter-table
name: White Abishai
creatures:
 - 1: White Abishai
```

## Environment

urban