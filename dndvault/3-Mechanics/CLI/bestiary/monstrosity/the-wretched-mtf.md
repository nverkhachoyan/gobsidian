---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mtf
- monster/cr/1-4
- monster/environment/swamp
- monster/environment/underdark
- monster/environment/urban
- monster/size/small
- monster/type/monstrosity
aliases: ["The Wretched"]
NoteIcon: monster
BestiaryType: monstrosity
SourceType: Bestiary
BookSource: Mordenkainen's Tome of Foes p. 233
---
# [The Wretched](3-Mechanics\CLI\bestiary\monstrosity/the-wretched-mtf.md)
*Source: Mordenkainen's Tome of Foes p. 233*  

## The Wretched

Horrid little monsters, the Wretched gather in large packs to scour the Shadowfell for prey. These pitiful entities subsist on life force, so when they find a creature, they surge forward to sink their fangs into their victims and drink deep of their life energy and their fear.

## Sorrowsworn

The Shadowfell's pervasive melancholy sometimes gives rise to strange incarnations of the plane's bleak nature. The sorrowsworn embody the forms of suffering that are inherent to the shadowy landscape, and they visit horror on those who stumble into their midst.

## Emotion Given Form

Each sorrowsworn personifies a different aspect of despair or distress. Some are manifestations of anger; others are loneliness given physical form. Their nature provides a clue both to understanding how they become more powerful and to overcoming them. Giving in to the negative emotions that the sorrowsworn represent causes these entities to grow deadlier. Fighting against these emotions can weaken them and drive them away.

```statblock
"name": "The Wretched (MTF)"
"size": "Small"
"type": "monstrosity"
"alignment": "Lawful Evil"
"ac": !!int "15"
"ac_class": "natural armor"
"hp": !!int "10"
"hit_dice": "4d6 - 4"
"stats":
- !!int "7"
- !!int "12"
- !!int "9"
- !!int "5"
- !!int "6"
- !!int "5"
"speed": "40 ft."
"damage_resistances": "bludgeoning, piercing, slashing while in dim light or darkness"
"senses": "darkvision 60 ft., passive Perception 8"
"languages": ""
"cr": "1/4"
"traits":
- "desc": "The Wretched has advantage on an attack roll against a creature if at least\
    \ one of the Wretched's allies is within 5 feet of the creature and the ally isn't\
    \ [incapacitated](/3-Mechanics/CLI/rules/conditions.md#incapacitated). The Wretched\
    \ otherwise has disadvantage on attack rolls."
  "name": "Wretched Pack Tactics"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+3 (+3) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d10 + 1|text(6) (1d10 + 1) piercing damage, and the Wretched\
    \ attaches to the target. While attached, the Wretched can't attack, and at the\
    \ start of each of the Wretched's turns, the target takes dice:1d10 + 1|text(6)\
    \ (1d10 + 1) necrotic damage.\n\nThe attached Wretched moves with the target\
    \ whenever the target moves, requiring none of the Wretched's movement. The Wretched\
    \ can detach itself by spending 5 feet of its movement on its turn. A creature,\
    \ including the target, can use its action to detach a Wretched."
  "name": "Bite"
"source":
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MTF/The%20Wretched.webp"
```
^statblock

```encounter-table
name: The Wretched
creatures:
 - 1: The Wretched
```

## Environment

swamp, underdark, urban