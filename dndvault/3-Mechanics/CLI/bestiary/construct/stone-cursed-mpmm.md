---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/1
- monster/environment/desert
- monster/environment/mountain
- monster/environment/urban
- monster/size/medium
- monster/type/construct
aliases: ["Stone Cursed"]
NoteIcon: monster
BestiaryType: construct
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 233, Mordenkainen's Tome of Foes p. 240
---
# [Stone Cursed](3-Mechanics\CLI\bestiary\construct/stone-cursed-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 233, Mordenkainen's Tome of Foes p. 240*  

```statblock
"name": "Stone Cursed (MPMM)"
"size": "Medium"
"type": "construct"
"alignment": "Typically  Lawful Evil"
"ac": !!int "17"
"ac_class": "natural armor"
"hp": !!int "19"
"hit_dice": "3d8 + 6"
"stats":
- !!int "16"
- !!int "5"
- !!int "14"
- !!int "5"
- !!int "8"
- !!int "7"
"speed": "10 ft."
"damage_vulnerabilities": "bludgeoning"
"damage_immunities": "poison"
"condition_immunities": "[charmed](/3-Mechanics/CLI/rules/conditions.md#charmed),\
  \ [exhaustion](/3-Mechanics/CLI/rules/conditions.md#exhaustion), [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened),\
  \ [petrified](/3-Mechanics/CLI/rules/conditions.md#petrified), [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)"
"senses": "passive Perception 9"
"languages": "the languages it knew in life"
"cr": "1"
"traits":
- "desc": "The stone cursed has advantage on the attack rolls of opportunity attacks."
  "name": "Cunning Opportunist"
- "desc": "If the stone cursed is motionless at the start of combat, it has advantage\
    \ on its initiative roll. Moreover, if a creature hasn't observed the stone cursed\
    \ move or act, that creature must succeed on a DC 18 Intelligence ([Investigation](/3-Mechanics/CLI/rules/skills.md#Investigation))\
    \ check to discern that the stone cursed isn't a statue."
  "name": "False Appearance"
- "desc": "The stone cursed doesn't require air, food, drink, or sleep."
  "name": "Unusual Nature"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+5 (+5) to hit, reach 5 ft., one target.\
    \ Hit: dice:2d8 + 3|text(12) (2d8 + 3) slashing damage. If the target is\
    \ a creature, it must succeed on a DC 12 Constitution saving throw, or it begins\
    \ to turn to stone and is [restrained](/3-Mechanics/CLI/rules/conditions.md#restrained)\
    \ until the end of its next turn, when it must repeat the saving throw. The effect\
    \ ends if the second save is successful; otherwise the target is [petrified](/3-Mechanics/CLI/rules/conditions.md#petrified)\
    \ for 24 hours."
  "name": "Petrifying Claws"
"source":
- "MPMM"
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Stone%20Cursed.webp"
```
^statblock

```encounter-table
name: Stone Cursed
creatures:
 - 1: Stone Cursed
```

## Environment

desert, mountain, urban