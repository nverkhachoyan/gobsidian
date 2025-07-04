---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mtf
- monster/cr/7
- monster/size/medium
- monster/type/elemental
aliases: ["Earth Elemental Myrmidon"]
NoteIcon: monster
BestiaryType: elemental
SourceType: Bestiary
BookSource: Mordenkainen's Tome of Foes p. 202
---
# [Earth Elemental Myrmidon](3-Mechanics\CLI\bestiary\elemental/earth-elemental-myrmidon-mtf.md)
*Source: Mordenkainen's Tome of Foes p. 202*  

## Earth Elemental Myrmidon

Elemental myrmidons are elementals conjured and bound by magic into ritually created suits of plate armor. In this form, they possess no recollection of their former existence as free elementals. They exist only to follow the commands of their creators.

```statblock
"name": "Earth Elemental Myrmidon (MTF)"
"size": "Medium"
"type": "elemental"
"alignment": "Neutral"
"ac": !!int "18"
"ac_class": "[plate armor](/3-Mechanics/CLI/items/plate-armor.md)"
"hp": !!int "127"
"hit_dice": "17d8 + 51"
"stats":
- !!int "18"
- !!int "10"
- !!int "17"
- !!int "8"
- !!int "10"
- !!int "10"
"speed": "30 ft."
"damage_resistances": "bludgeoning, piercing, slashing from nonmagical attacks"
"damage_immunities": "poison"
"condition_immunities": "[paralyzed](/3-Mechanics/CLI/rules/conditions.md#paralyzed),\
  \ [petrified](/3-Mechanics/CLI/rules/conditions.md#petrified), [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned),\
  \ [prone](/3-Mechanics/CLI/rules/conditions.md#prone)"
"senses": "darkvision 60 ft., passive Perception 10"
"languages": "Terran, one language of its creator's choice"
"cr": "7"
"traits":
- "desc": "The myrmidon's weapon attacks are magical."
  "name": "Magic Weapons"
"actions":
- "desc": "The myrmidon makes two maul attacks."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+7 (+7) to hit, reach 5 ft., one target.\
    \ Hit: dice:2d6 + 4|text(11) (2d6 + 4) bludgeoning damage."
  "name": "Maul"
- "desc": "The myrmidon makes one maul attack. On a hit, the target takes an extra\
    \ dice:3d10|text(16) (3d10) thunder damage, and the target must succeed on\
    \ a DC 14 Strength saving throw or be knocked [prone](/3-Mechanics/CLI/rules/conditions.md#prone)."
  "name": "Thunderous Strike (Recharge 6)"
"source":
- "MTF"
- "PotA"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MTF/Earth%20Elemental%20Myrmidon.webp"
```
^statblock

```encounter-table
name: Earth Elemental Myrmidon
creatures:
 - 1: Earth Elemental Myrmidon
```