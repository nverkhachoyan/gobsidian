---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mtf
- monster/cr/8
- monster/environment/grassland
- monster/environment/swamp
- monster/size/medium
- monster/type/undead
aliases: ["Sword Wraith Commander"]
NoteIcon: monster
BestiaryType: undead
SourceType: Bestiary
BookSource: Mordenkainen's Tome of Foes p. 241, Divine Contention, Dragon of Icespire Peak, Sleeping Dragon's Wake
---
# [Sword Wraith Commander](3-Mechanics\CLI\bestiary\undead/sword-wraith-commander-mtf.md)
*Source: Mordenkainen's Tome of Foes p. 241, Divine Contention, Dragon of Icespire Peak, Sleeping Dragon's Wake*  

## Sword Wraith

When a glory-obsessed warrior dies in battle without earning the honor it desperately sought, its valor-hungry spirit might haunt the battlefield as a sword wraith.

## Brooding Spirits

The most likely spots for encountering sword wraiths are scenes of ancient ambushes, battlefields where soldiers were felled by magic with no chance to fight back, and sites where enemies were hemmed in and slaughtered without quarter.

## Honor Above All

Sword wraiths fly into a rage if anyone questions their valor. Conversely, they are easily appeased by praise. Little pleases them more than hearing a ballad performed in their honor. Towns located near ancient battlefields hold annual festivals of remembrance to keep sword wraiths there placated.

## Undead Nature

A sword wraith doesn't require air, food, drink, or sleep.

```statblock
"name": "Sword Wraith Commander (MTF)"
"size": "Medium"
"type": "undead"
"alignment": "Lawful Evil"
"ac": !!int "18"
"ac_class": "[breastplate](/3-Mechanics/CLI/items/breastplate.md), [shield](/3-Mechanics/CLI/items/shield.md)"
"hp": !!int "127"
"hit_dice": "15d8 + 60"
"stats":
- !!int "18"
- !!int "14"
- !!int "18"
- !!int "11"
- !!int "12"
- !!int "14"
"speed": "30 ft."
"skillsaves":
  "Perception": !!int "4"
"damage_resistances": "necrotic; bludgeoning, piercing, slashing from nonmagical attacks"
"damage_immunities": "poison"
"condition_immunities": "[exhaustion](/3-Mechanics/CLI/rules/conditions.md#exhaustion),\
  \ [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened), [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned),\
  \ [unconscious](/3-Mechanics/CLI/rules/conditions.md#unconscious)"
"senses": "darkvision 60 ft., passive Perception 14"
"languages": "the languages it knew in life"
"cr": "8"
"traits":
- "desc": "As a bonus action, the sword wraith can make one weapon attack, which deals\
    \ an extra dice:2d8|text(9) (2d8) necrotic damage on a hit. If it does so,\
    \ attack rolls against it have advantage until the start of its next turn."
  "name": "Martial Fury"
- "desc": "The sword wraith and any other sword wraiths within 30 feet of it have\
    \ advantage on saving throws against effects that turn undead."
  "name": "Turning Defiance"
"actions":
- "desc": "The sword wraith makes two weapon attacks."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+7 (+7) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d8 + 4|text(8) (1d8 + 4) slashing damage, or dice:1d10 + 4|text(9)\
    \ (1d10 + 4) slashing damage if used with two hands."
  "name": "Longsword"
- "desc": "Ranged Weapon Attack: dice: d20+5 (+5) to hit, range 150/600 ft.,\
    \ one target. Hit: dice:1d8 + 2|text(6) (1d8 + 2) piercing damage."
  "name": "Longbow"
- "desc": "To use this action, the sword wraith must have taken damage during the\
    \ current combat. If the sword wraith can use this action, it gives itself advantage\
    \ on attack rolls until the end of its next turn, and dice: 1d4 + 1|avg|noform\
    \ (1d4 + 1) [sword wraith warriors](/3-Mechanics/CLI/bestiary/undead/sword-wraith-warrior-mtf.md)\
    \ appear in unoccupied spaces within 30 feet of it. The warriors last until they\
    \ drop to 0 hit points, and they take their turns immediately after the commander's\
    \ turn on the same initiative count."
  "name": "Call to Honor (1/Day)"
"source":
- "MTF"
- "DC"
- "DIP"
- "SDW"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MTF/Sword%20Wraith%20Commander.webp"
```
^statblock

```encounter-table
name: Sword Wraith Commander
creatures:
 - 1: Sword Wraith Commander
```

## Environment

grassland, swamp