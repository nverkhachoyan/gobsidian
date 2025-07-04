---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mtf
- monster/cr/3
- monster/environment/grassland
- monster/environment/swamp
- monster/size/medium
- monster/type/undead
aliases: ["Sword Wraith Warrior"]
NoteIcon: monster
BestiaryType: undead
SourceType: Bestiary
BookSource: Mordenkainen's Tome of Foes p. 241, Divine Contention, Dragon of Icespire Peak, Sleeping Dragon's Wake
---
# [Sword Wraith Warrior](3-Mechanics\CLI\bestiary\undead/sword-wraith-warrior-mtf.md)
*Source: Mordenkainen's Tome of Foes p. 241, Divine Contention, Dragon of Icespire Peak, Sleeping Dragon's Wake*  

## Sword Wraith

When a glory-obsessed warrior dies in battle without earning the honor it desperately sought, its valor—hungry spirit might haunt the battlefield as a sword wraith.

## Brooding Spirits

The most likely spots for encountering sword wraiths are scenes of ancient ambushes, battlefields where soldiers were felled by magic with no chance to fight back, and sites where enemies were hemmed in and slaughtered without quarter.

## Honor Above All

Sword wraiths fly into a rage if anyone questions their valor. Conversely, they are easily appeased by praise. Little pleases them more than hearing a ballad performed in their honor. Towns located near ancient battlefields hold annual festivals of remembrance to keep sword wraiths there placated.

## Undead Nature

A sword wraith doesn't require air, food, drink, or sleep.

```statblock
"name": "Sword Wraith Warrior (MTF)"
"size": "Medium"
"type": "undead"
"alignment": "Lawful Evil"
"ac": !!int "16"
"ac_class": "[chain shirt](/3-Mechanics/CLI/items/chain-shirt.md), [shield](/3-Mechanics/CLI/items/shield.md)"
"hp": !!int "45"
"hit_dice": "6d8 + 18"
"stats":
- !!int "18"
- !!int "12"
- !!int "17"
- !!int "6"
- !!int "9"
- !!int "10"
"speed": "30 ft."
"damage_resistances": "necrotic; bludgeoning, piercing, slashing from nonmagical attacks"
"damage_immunities": "poison"
"condition_immunities": "[exhaustion](/3-Mechanics/CLI/rules/conditions.md#exhaustion),\
  \ [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened), [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned),\
  \ [unconscious](/3-Mechanics/CLI/rules/conditions.md#unconscious)"
"senses": "darkvision 60 ft., passive Perception 9"
"languages": "the languages it knew in life"
"cr": "3"
"traits":
- "desc": "As a bonus action, the sword wraith can make one weapon attack. If it does\
    \ so, attack rolls against it have advantage until the start of its next turn."
  "name": "Martial Fury"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+6 (+6) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d8 + 4|text(8) (1d8 + 4) slashing damage, or dice:1d10 + 4|text(9)\
    \ (1d10 + 4) slashing damage if used with two hands."
  "name": "Longsword"
- "desc": "Ranged Weapon Attack: dice: d20+3 (+3) to hit, range 150/600 ft.,\
    \ one target. Hit: dice:1d8 + 1|text(5) (1d8 + 1) piercing damage."
  "name": "Longbow"
"source":
- "MTF"
- "DC"
- "DIP"
- "SDW"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MTF/Sword%20Wraith%20Warrior.webp"
```
^statblock

```encounter-table
name: Sword Wraith Warrior
creatures:
 - 1: Sword Wraith Warrior
```

## Environment

grassland, swamp