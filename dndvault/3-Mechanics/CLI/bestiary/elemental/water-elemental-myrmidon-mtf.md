---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mtf
- monster/cr/7
- monster/size/medium
- monster/type/elemental
aliases: ["Water Elemental Myrmidon"]
NoteIcon: monster
BestiaryType: elemental
SourceType: Bestiary
BookSource: Mordenkainen's Tome of Foes p. 203
---
# [Water Elemental Myrmidon](3-Mechanics\CLI\bestiary\elemental/water-elemental-myrmidon-mtf.md)
*Source: Mordenkainen's Tome of Foes p. 203*  

## Water Elemental Myrmidon

Elemental myrmidons are elementals conjured and bound by magic into ritually created suits of plate armor. In this form, they possess no recollection of their former existence as free elementals. They exist only to follow the commands of their creators.

```statblock
"name": "Water Elemental Myrmidon (MTF)"
"size": "Medium"
"type": "elemental"
"alignment": "Neutral"
"ac": !!int "18"
"ac_class": "[plate armor](/3-Mechanics/CLI/items/plate-armor.md)"
"hp": !!int "127"
"hit_dice": "17d8 + 51"
"stats":
- !!int "18"
- !!int "14"
- !!int "15"
- !!int "8"
- !!int "10"
- !!int "10"
"speed": "40 ft., swim 40 ft."
"damage_resistances": "acid; bludgeoning, piercing, slashing from nonmagical attacks"
"damage_immunities": "poison"
"condition_immunities": "[paralyzed](/3-Mechanics/CLI/rules/conditions.md#paralyzed),\
  \ [petrified](/3-Mechanics/CLI/rules/conditions.md#petrified), [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned),\
  \ [prone](/3-Mechanics/CLI/rules/conditions.md#prone)"
"senses": "darkvision 60 ft., passive Perception 10"
"languages": "Aquan, one language of its creator's choice"
"cr": "7"
"traits":
- "desc": "The myrmidon's weapon attacks are magical."
  "name": "Magic Weapons"
"actions":
- "desc": "The myrmidon makes three trident attacks."
  "name": "Multiattack"
- "desc": "Melee or Ranged Weapon Attack: dice: d20+7 (+7) to hit, reach 5 ft.\
    \ or range 20/60 ft., one target. Hit: dice:1d6 + 4|text(7) (1d6 + 4) piercing\
    \ damage, or dice:1d8 + 4|text(8) (1d8 + 4) piercing damage if used with two\
    \ hands to make a melee attack."
  "name": "Trident"
- "desc": "The myrmidon uses Multiattack. Each attack that hits deals an extra dice:1d10|text(5)\
    \ (1d10) cold damage. A target that is hit by one or more of these attacks has\
    \ its speed reduced by 10 feet until the end of the myrmidon's next turn."
  "name": "Freezing Strikes (Recharge 6)"
"source":
- "MTF"
- "LR"
- "PotA"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MTF/Water%20Elemental%20Myrmidon.webp"
```
^statblock

```encounter-table
name: Water Elemental Myrmidon
creatures:
 - 1: Water Elemental Myrmidon
```