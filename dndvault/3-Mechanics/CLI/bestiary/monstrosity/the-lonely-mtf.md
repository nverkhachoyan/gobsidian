---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mtf
- monster/cr/9
- monster/environment/coastal
- monster/environment/desert
- monster/environment/mountain
- monster/environment/underdark
- monster/environment/urban
- monster/size/medium
- monster/type/monstrosity
aliases: ["The Lonely"]
NoteIcon: monster
BestiaryType: monstrosity
SourceType: Bestiary
BookSource: Mordenkainen's Tome of Foes p. 232
---
# [The Lonely](3-Mechanics\CLI\bestiary\monstrosity/the-lonely-mtf.md)
*Source: Mordenkainen's Tome of Foes p. 232*  

## The Lonely

The sorrow of isolation afflicts many creatures that lurk in the Shadowfell, but the need for companionship is never manifested more dramatically than in the Lonely. When these sorrowsworn spot other creatures, they feel keenly the need for interaction and so they launch their harpoon-like arms to drag their victims close.

## Sorrowsworn

The Shadowfell's pervasive melancholy sometimes gives rise to strange incarnations of the plane's bleak nature. The sorrowsworn embody the forms of suffering that are inherent to the shadowy landscape, and they visit horror on those who stumble into their midst.

## Emotion Given Form

Each sorrowsworn personifies a different aspect of despair or distress. Some are manifestations of anger; others are loneliness given physical form. Their nature provides a clue both to understanding how they become more powerful and to overcoming them. Giving in to the negative emotions that the sorrowsworn represent causes these entities to grow deadlier. Fighting against these emotions can weaken them and drive them away.

```statblock
"name": "The Lonely (MTF)"
"size": "Medium"
"type": "monstrosity"
"alignment": "Lawful Evil"
"ac": !!int "16"
"ac_class": "natural armor"
"hp": !!int "112"
"hit_dice": "15d8 + 45"
"stats":
- !!int "16"
- !!int "12"
- !!int "17"
- !!int "6"
- !!int "11"
- !!int "6"
"speed": "30 ft."
"damage_resistances": "bludgeoning, piercing, slashing while in dim light or darkness"
"senses": "darkvision 60 ft., passive Perception 10"
"languages": "Common"
"cr": "9"
"traits":
- "desc": "At the start of each of the Lonely's turns, each creature within 5 feet\
    \ of it must succeed on a DC 15 Wisdom saving throw or take dice:3d6|text(10)\
    \ (3d6) psychic damage."
  "name": "Psychic Leech"
- "desc": "The Lonely has advantage on attack rolls while it is within 30 feet of\
    \ at least two other creatures. It otherwise has disadvantage on attack rolls."
  "name": "Thrives on Company"
"actions":
- "desc": "The Lonely makes one harpoon arm attack and uses Sorrowful Embrace."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+7 (+7) to hit, reach 60 ft., one target.\
    \ Hit: dice:4d8 + 3|text(21) (4d8 + 3) piercing damage, and the target is\
    \ [grappled](/3-Mechanics/CLI/rules/conditions.md#grappled) (escape DC 15) if\
    \ it is a Large or smaller creature.\n\nThe Lonely has two harpoon arms and can\
    \ grapple up to two creatures at once."
  "name": "Harpoon Arm"
- "desc": "Each creature [grappled](/3-Mechanics/CLI/rules/conditions.md#grappled)\
    \ by the Lonely must make a DC 15 Wisdom saving throw. A creature takes dice:4d8|text(18)\
    \ (4d8) psychic damage on a failed save, or half as much damage on a successful\
    \ one. In either case, the Lonely pulls each creature [grappled](/3-Mechanics/CLI/rules/conditions.md#grappled)\
    \ by it up to 30 feet straight toward it."
  "name": "Sorrowful Embrace"
"source":
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MTF/The%20Lonely.webp"
```
^statblock

```encounter-table
name: The Lonely
creatures:
 - 1: The Lonely
```

## Environment

coastal, desert, mountain, underdark, urban