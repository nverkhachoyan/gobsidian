---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/7
- monster/size/medium
- monster/type/elemental
aliases: ["Air Elemental Myrmidon"]
NoteIcon: monster
BestiaryType: elemental
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 122, Mordenkainen's Tome of Foes p. 202
---
# [Air Elemental Myrmidon](3-Mechanics\CLI\bestiary\elemental/air-elemental-myrmidon-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 122, Mordenkainen's Tome of Foes p. 202*  

```statblock
"name": "Air Elemental Myrmidon (MPMM)"
"size": "Medium"
"type": "elemental"
"alignment": "Typically  Neutral"
"ac": !!int "18"
"ac_class": "[plate](/3-Mechanics/CLI/items/plate-armor.md)"
"hp": !!int "117"
"hit_dice": "18d8 + 36"
"stats":
- !!int "18"
- !!int "14"
- !!int "14"
- !!int "9"
- !!int "10"
- !!int "10"
"speed": "30 ft., fly 30 ft. (hover)"
"damage_resistances": "lightning; thunder; bludgeoning, piercing, slashing from nonmagical\
  \ attacks"
"damage_immunities": "poison"
"condition_immunities": "[paralyzed](/3-Mechanics/CLI/rules/conditions.md#paralyzed),\
  \ [petrified](/3-Mechanics/CLI/rules/conditions.md#petrified), [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned),\
  \ [prone](/3-Mechanics/CLI/rules/conditions.md#prone)"
"senses": "darkvision 60 ft., passive Perception 10"
"languages": "Auran, one language of its creator's choice"
"cr": "7"
"actions":
- "desc": "The myrmidon makes three Flail attacks."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+7 (+7) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d8 + 4|text(8) (1d8 + 4) force damage."
  "name": "Flail"
- "desc": "The myrmidon makes one Flail attack. On a hit, the target takes an extra\
    \ dice:4d8|text(18) (4d8) lightning damage, and the target must succeed on\
    \ a DC 13 Constitution saving throw or be [stunned](/3-Mechanics/CLI/rules/conditions.md#stunned)\
    \ until the end of the myrmidon's next turn."
  "name": "Lightning Strike (Recharge 6)"
"source":
- "MPMM"
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Air%20Elemental%20Myrmidon.webp"
```
^statblock

```encounter-table
name: Air Elemental Myrmidon
creatures:
 - 1: Air Elemental Myrmidon
```