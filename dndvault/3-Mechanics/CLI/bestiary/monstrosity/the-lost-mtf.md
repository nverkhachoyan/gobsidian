---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mtf
- monster/cr/7
- monster/environment/arctic
- monster/environment/desert
- monster/environment/forest
- monster/environment/mountain
- monster/environment/swamp
- monster/environment/underdark
- monster/environment/urban
- monster/size/medium
- monster/type/monstrosity
aliases: ["The Lost"]
NoteIcon: monster
BestiaryType: monstrosity
SourceType: Bestiary
BookSource: Mordenkainen's Tome of Foes p. 233
---
# [The Lost](3-Mechanics\CLI\bestiary\monstrosity/the-lost-mtf.md)
*Source: Mordenkainen's Tome of Foes p. 233*  

## The Lost

The Shadowfell turns visitors around until they become marooned in its twisted landscape. The Lost are representations of the anxiety and fear that people experience when they can't find their way. These sorrowsworn appear as desperate and panicked things.

The Lost try to embrace any creatures they can reach, attempting to find solace in the contact. Aside from the horror of being embraced by such a thing, the victim experiences a flood of fear and panic as its mind buckles under the fury of this assault. The harder the victims' allies fight for release, the worse the experience becomes.

## Sorrowsworn

The Shadowfell's pervasive melancholy sometimes gives rise to strange incarnations of the plane's bleak nature. The sorrowsworn embody the forms of suffering that are inherent to the shadowy landscape, and they visit horror on those who stumble into their midst.

## Emotion Given Form

Each sorrowsworn personifies a different aspect of despair or distress. Some are manifestations of anger; others are loneliness given physical form. Their nature provides a clue both to understanding how they become more powerful and to overcoming them. Giving in to the negative emotions that the sorrowsworn represent causes these entities to grow deadlier. Fighting against these emotions can weaken them and drive them away.

```statblock
"name": "The Lost (MTF)"
"size": "Medium"
"type": "monstrosity"
"alignment": "Lawful Evil"
"ac": !!int "15"
"ac_class": "natural armor"
"hp": !!int "78"
"hit_dice": "12d8 + 24"
"stats":
- !!int "17"
- !!int "12"
- !!int "15"
- !!int "6"
- !!int "7"
- !!int "5"
"speed": "30 ft."
"skillsaves":
  "Athletics": !!int "6"
"damage_resistances": "bludgeoning, piercing, slashing while in dim light or darkness"
"senses": "darkvision 60 ft., passive Perception 8"
"languages": "Common"
"cr": "7"
"actions":
- "desc": "The Lost makes two arm spike attacks."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+6 (+6) to hit, reach 5 ft., one target.\
    \ Hit: dice:2d10 + 3|text(14) (2d10 + 3) piercing damage."
  "name": "Arm Spike"
- "desc": "Melee Weapon Attack: dice: d20+6 (+6) to hit, reach 5 ft., one target.\
    \ Hit: dice:4d10 + 3|text(25) (4d10 + 3) piercing damage, and the target\
    \ is [grappled](/3-Mechanics/CLI/rules/conditions.md#grappled) (escape DC 14)\
    \ if it is a Medium or smaller creature. Until the grapple ends, the target is\
    \ [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened), and it takes\
    \ dice:6d8|text(27) (6d8) psychic damage at the end of each of its turns.\
    \ The Lost can embrace only one creature at a time."
  "name": "Embrace"
"reactions":
- "desc": "If the Lost takes damage while it has a creature [grappled](/3-Mechanics/CLI/rules/conditions.md#grappled),\
    \ that creature takes dice:4d8|text(18) (4d8) psychic damage."
  "name": "Tightening Embrace"
"source":
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MTF/The%20Lost.webp"
```
^statblock

```encounter-table
name: The Lost
creatures:
 - 1: The Lost
```

## Environment

arctic, desert, forest, mountain, swamp, underdark, urban