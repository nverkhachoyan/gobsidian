---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/dip
- monster/cr/4
- monster/size/medium
- monster/type/humanoid/human
aliases: ["Falcon the Hunter"]
NoteIcon: monster
BestiaryType: humanoid (human)
SourceType: Bestiary
BookSource: Dragon of Icespire Peak p. 56
---
# [Falcon the Hunter](3-Mechanics\CLI\bestiary\npc/falcon-the-hunter-dip.md)
*Source: Dragon of Icespire Peak p. 56*  

Adventurers encounter Falcon the Hunter if they visit his hunting lodge in Neverwinter Wood.

```statblock
"name": "Falcon the Hunter (DIP)"
"size": "Medium"
"type": "humanoid"
"subtype": "human"
"alignment": "Neutral Good"
"ac": !!int "14"
"ac_class": "[studded leather](/3-Mechanics/CLI/items/studded-leather-armor.md)"
"hp": !!int "112"
"hit_dice": "15d8 + 45"
"stats":
- !!int "14"
- !!int "15"
- !!int "16"
- !!int "11"
- !!int "16"
- !!int "15"
"speed": "30 ft."
"saves":
  "Dexterity": !!int "4"
  "Wisdom": !!int "5"
"skillsaves":
  "Athletics": !!int "4"
  "Perception": !!int "7"
  "Survival": !!int "5"
"senses": "passive Perception 17"
"languages": "Common"
"cr": "4"
"traits":
- "desc": "A longbow or shortbow deals one extra die of its damage when Falcon hits\
    \ with it (included in his longbow attack)."
  "name": "Archer"
- "desc": "Falcon's ranged weapon attacks ignore half cover and three-quarters cover."
  "name": "Sharpshooter"
"actions":
- "desc": "Falcon makes three melee attacks or two ranged attacks."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+4 (+4) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d8 + 2|text(6) (1d8 + 2) slashing damage, or dice:1d10 + 2|text(7)\
    \ (1d10 + 2) slashing damage if used with two hands."
  "name": "Longsword"
- "desc": "Ranged Weapon Attack: dice: d20+4 (+4) to hit, range 150/600 ft.,\
    \ one target. Hit: dice:2d8 + 2|text(11) (2d8 + 2) piercing damage."
  "name": "Longbow"
"source":
- "DIP"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/DIP/Falcon%20the%20Hunter.webp"
```
^statblock

```encounter-table
name: Falcon the Hunter
creatures:
 - 1: Falcon the Hunter
```