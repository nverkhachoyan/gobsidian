---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/toa
- monster/cr/3
- monster/size/large
- monster/type/plant
aliases: ["Assassin Vine"]
NoteIcon: monster
BestiaryType: plant
SourceType: Bestiary
BookSource: Tomb of Annihilation p. 213, Sleeping Dragon's Wake
---
# [Assassin Vine](3-Mechanics\CLI\bestiary\plant/assassin-vine-toa.md)
*Source: Tomb of Annihilation p. 213, Sleeping Dragon's Wake*  

```statblock
"name": "Assassin Vine (ToA)"
"size": "Large"
"type": "plant"
"alignment": "Unaligned"
"ac": !!int "13"
"ac_class": "natural armor"
"hp": !!int "85"
"hit_dice": "10d10 + 30"
"stats":
- !!int "18"
- !!int "10"
- !!int "16"
- !!int "1"
- !!int "10"
- !!int "1"
"speed": "5 ft., climb 5 ft."
"damage_resistances": "cold, fire"
"condition_immunities": "[blinded](/3-Mechanics/CLI/rules/conditions.md#blinded),\
  \ [deafened](/3-Mechanics/CLI/rules/conditions.md#deafened), [exhaustion](/3-Mechanics/CLI/rules/conditions.md#exhaustion),\
  \ [prone](/3-Mechanics/CLI/rules/conditions.md#prone)"
"senses": "blindsight 30 ft., passive Perception 10"
"languages": ""
"cr": "3"
"traits":
- "desc": "While the assassin vine remains motionless, it is indistinguishable from\
    \ a normal plant."
  "name": "False Appearance"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+6 (+6) to hit, reach 20 ft., one creature.\
    \ Hit: The target takes dice:2d6 + 4|text(11) (2d6 + 4) bludgeoning damage,\
    \ and it is [grappled](/3-Mechanics/CLI/rules/conditions.md#grappled) (escape\
    \ DC 14). Until this grapple ends, the target is [restrained](/3-Mechanics/CLI/rules/conditions.md#restrained),\
    \ and it takes dice:6d6|text(21) (6d6) poison damage at the start of each\
    \ of its turns. The vine can constrict only one target at a time."
  "name": "Constrict"
- "desc": "The assassin vine can animate normal vines and roots on the ground in a\
    \ 15-foot square within 30 feet of it. These plants turn the ground in that area\
    \ into difficult terrain. A creature in that area when the effect begins must\
    \ succeed on a DC 13 Strength saving throw or be [restrained](/3-Mechanics/CLI/rules/conditions.md#restrained)\
    \ by entangling vines and roots. A creature [restrained](/3-Mechanics/CLI/rules/conditions.md#restrained)\
    \ by the plants can use its action to make a DC 13 Strength ([Athletics](/3-Mechanics/CLI/rules/skills.md#Athletics))\
    \ check, freeing itself on a successful check. The effect ends after 1 minute\
    \ or when the assassin vine dies or uses Entangling Vines again."
  "name": "Entangling Vines"
"source":
- "ToA"
- "GoS"
- "SDW"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/ToA/Assassin%20Vine.webp"
```
^statblock

```encounter-table
name: Assassin Vine
creatures:
 - 1: Assassin Vine
```