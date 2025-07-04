---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/xge
- monster/cr/
- monster/size/tiny
- monster/type/construct
aliases: ["Tiny Servant"]
NoteIcon: monster
BestiaryType: construct
SourceType: Bestiary
BookSource: Xanathar's Guide to Everything p. 169
---
# [Tiny Servant](3-Mechanics\CLI\bestiary\construct/tiny-servant-xge.md)
*Source: Xanathar's Guide to Everything p. 169*  

```statblock
"name": "Tiny Servant (XGE)"
"size": "Tiny"
"type": "construct"
"alignment": "Unaligned"
"ac": !!int "15"
"ac_class": "natural armor"
"hp": !!int "10"
"hit_dice": "4d4"
"stats":
- !!int "4"
- !!int "16"
- !!int "10"
- !!int "2"
- !!int "10"
- !!int "1"
"speed": "30 ft., climb 30 ft."
"damage_immunities": "poison, psychic"
"condition_immunities": "[blinded](/3-Mechanics/CLI/rules/conditions.md#blinded),\
  \ [charmed](/3-Mechanics/CLI/rules/conditions.md#charmed), [deafened](/3-Mechanics/CLI/rules/conditions.md#deafened),\
  \ [exhaustion](/3-Mechanics/CLI/rules/conditions.md#exhaustion), [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened),\
  \ [paralyzed](/3-Mechanics/CLI/rules/conditions.md#paralyzed), [petrified](/3-Mechanics/CLI/rules/conditions.md#petrified),\
  \ [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)"
"senses": "blindsight 60 ft. (blind beyond this radius), passive Perception 10"
"languages": ""
"actions":
- "desc": "Melee Weapon Attack: dice: d20+5 (+5) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d4 + 3|text(5) (1d4 + 3) bludgeoning damage."
  "name": "Slam"
"source":
- "XGE"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/XGE/Tiny%20Servant.webp"
```
^statblock

```encounter-table
name: Tiny Servant
creatures:
 - 1: Tiny Servant
```