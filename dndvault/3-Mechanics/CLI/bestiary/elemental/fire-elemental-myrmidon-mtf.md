---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mtf
- monster/cr/7
- monster/size/medium
- monster/type/elemental
aliases: ["Fire Elemental Myrmidon"]
NoteIcon: monster
BestiaryType: elemental
SourceType: Bestiary
BookSource: Mordenkainen's Tome of Foes p. 203
---
# [Fire Elemental Myrmidon](3-Mechanics\CLI\bestiary\elemental/fire-elemental-myrmidon-mtf.md)
*Source: Mordenkainen's Tome of Foes p. 203*  

## Fire Elemental Myrmidon

Elemental myrmidons are elementals conjured and bound by magic into ritually created suits of plate armor. In this form, they possess no recollection of their former existence as free elementals. They exist only to follow the commands of their creators.

```statblock
"name": "Fire Elemental Myrmidon (MTF)"
"size": "Medium"
"type": "elemental"
"alignment": "Neutral"
"ac": !!int "18"
"ac_class": "[plate armor](/3-Mechanics/CLI/items/plate-armor.md)"
"hp": !!int "123"
"hit_dice": "19d8 + 38"
"stats":
- !!int "13"
- !!int "18"
- !!int "15"
- !!int "9"
- !!int "10"
- !!int "10"
"speed": "40 ft."
"damage_resistances": "bludgeoning, piercing, slashing from nonmagical attacks"
"damage_immunities": "fire, poison"
"condition_immunities": "[paralyzed](/3-Mechanics/CLI/rules/conditions.md#paralyzed),\
  \ [petrified](/3-Mechanics/CLI/rules/conditions.md#petrified), [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned),\
  \ [prone](/3-Mechanics/CLI/rules/conditions.md#prone)"
"senses": "darkvision 60 ft., passive Perception 10"
"languages": "Ignan, one language of its creator's choice"
"cr": "7"
"traits":
- "desc": "The myrmidon sheds bright light in a 20-foot radius and dim light in a\
    \ 40-foot radius."
  "name": "Illumination"
- "desc": "The myrmidon's weapon attacks are magical."
  "name": "Magic Weapons"
- "desc": "For every 5 feet the myrmidon moves in 1 foot or more of water, it takes\
    \ dice:1d4|text(2) (1d4) cold damage."
  "name": "Water Susceptibility"
"actions":
- "desc": "The myrmidon makes three scimitar attacks."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+7 (+7) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d6 + 4|text(7) (1d6 + 4) slashing damage."
  "name": "Scimitar"
- "desc": "The myrmidon uses Multiattack. Each attack that hits deals an extra dice:1d10|text(5)\
    \ (1d10) fire damage."
  "name": "Fiery Strikes (Recharge 6)"
"source":
- "MTF"
- "PotA"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MTF/Fire%20Elemental%20Myrmidon.webp"
```
^statblock

```encounter-table
name: Fire Elemental Myrmidon
creatures:
 - 1: Fire Elemental Myrmidon
```