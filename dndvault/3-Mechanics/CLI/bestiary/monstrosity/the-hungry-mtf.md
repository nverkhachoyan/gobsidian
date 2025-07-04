---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mtf
- monster/cr/11
- monster/environment/forest
- monster/environment/underdark
- monster/environment/urban
- monster/size/medium
- monster/type/monstrosity
aliases: ["The Hungry"]
NoteIcon: monster
BestiaryType: monstrosity
SourceType: Bestiary
BookSource: Mordenkainen's Tome of Foes p. 232
---
# [The Hungry](3-Mechanics\CLI\bestiary\monstrosity/the-hungry-mtf.md)
*Source: Mordenkainen's Tome of Foes p. 232*  

## The Hungry

Horrid beasts with grasping claws and yawning mouths, the Hungry do whatever is necessary to sate their appetites. These greedy devourers consume all life and energy they encounter, stuffing their maws with flesh and drinking in their victims' screams. When they finish, they lurch away while their bright eyes resume the search for something else to consume.

## Sorrowsworn

The Shadowfell's pervasive melancholy sometimes gives rise to strange incarnations of the plane's bleak nature. The sorrowsworn embody the forms of suffering that are inherent to the shadowy landscape, and they visit horror on those who stumble into their midst.

## Emotion Given Form

Each sorrowsworn personifies a different aspect of despair or distress. Some are manifestations of anger; others are loneliness given physical form. Their nature provides a clue both to understanding how they become more powerful and to overcoming them. Giving in to the negative emotions that the sorrowsworn represent causes these entities to grow deadlier. Fighting against these emotions can weaken them and drive them away.

```statblock
"name": "The Hungry (MTF)"
"size": "Medium"
"type": "monstrosity"
"alignment": "Lawful Evil"
"ac": !!int "17"
"ac_class": "natural armor"
"hp": !!int "225"
"hit_dice": "30d8 + 90"
"stats":
- !!int "19"
- !!int "10"
- !!int "17"
- !!int "6"
- !!int "11"
- !!int "6"
"speed": "30 ft."
"damage_resistances": "bludgeoning, piercing, slashing while in dim light or darkness"
"senses": "darkvision 60 ft., passive Perception 10"
"languages": "Common"
"cr": "11"
"traits":
- "desc": "If a creature the Hungry can see regains hit points, the Hungry gains two\
    \ benefits until the end of its next turn: it has advantage on attack rolls, and\
    \ its bite deals an extra dice:4d10|text(22) (4d10) necrotic damage on a hit."
  "name": "Life Hunger"
"actions":
- "desc": "The Hungry makes two attacks: one with its bite and one with its claws."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+8 (+8) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d8 + 4|text(8) (1d8 + 4) piercing damage plus dice:3d8|text(13)\
    \ (3d8) necrotic damage."
  "name": "Bite"
- "desc": "Melee Weapon Attack: dice: d20+8 (+8) to hit, reach 10 ft., one target.\
    \ Hit: dice:4d6 + 4|text(18) (4d6 + 4) slashing damage. If the target is\
    \ Medium or smaller, it is [grappled](/3-Mechanics/CLI/rules/conditions.md#grappled)\
    \ (escape DC 16) and is [restrained](/3-Mechanics/CLI/rules/conditions.md#restrained)\
    \ until the grapple ends. While grappling a creature, the Hungry can't attack\
    \ with its claws."
  "name": "Claws"
"source":
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MTF/The%20Hungry.webp"
```
^statblock

```encounter-table
name: The Hungry
creatures:
 - 1: The Hungry
```

## Environment

forest, underdark, urban