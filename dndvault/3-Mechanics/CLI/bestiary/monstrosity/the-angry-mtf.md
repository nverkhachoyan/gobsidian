---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mtf
- monster/cr/13
- monster/environment/underdark
- monster/environment/urban
- monster/size/medium
- monster/type/monstrosity
aliases: ["The Angry"]
NoteIcon: monster
BestiaryType: monstrosity
SourceType: Bestiary
BookSource: Mordenkainen's Tome of Foes p. 231
---
# [The Angry](3-Mechanics\CLI\bestiary\monstrosity/the-angry-mtf.md)
*Source: Mordenkainen's Tome of Foes p. 231*  

## The Angry

Relying on violence to sustain their existence, the Angry grow more powerful when their foes fight back. If a creature opts not to attack, though, the Angry becomes confused, and its attacks weaken. Each of the Angry has two heads, which bicker with each other until they find something else on which they can vent their wrath.

## Sorrowsworn

The Shadowfell's pervasive melancholy sometimes gives rise to strange incarnations of the plane's bleak nature. The sorrowsworn embody the forms of suffering that are inherent to the shadowy landscape, and they visit horror on those who stumble into their midst.

## Emotion Given Form

Each sorrowsworn personifies a different aspect of despair or distress. Some are manifestations of anger; others are loneliness given physical form. Their nature provides a clue both to understanding how they become more powerful and to overcoming them. Giving in to the negative emotions that the sorrowsworn represent causes these entities to grow deadlier. Fighting against these emotions can weaken them and drive them away.

```statblock
"name": "The Angry (MTF)"
"size": "Medium"
"type": "monstrosity"
"alignment": "Lawful Evil"
"ac": !!int "18"
"ac_class": "natural armor"
"hp": !!int "255"
"hit_dice": "30d8 + 120"
"stats":
- !!int "17"
- !!int "10"
- !!int "19"
- !!int "8"
- !!int "13"
- !!int "6"
"speed": "30 ft."
"skillsaves":
  "Perception": !!int "6"
"damage_resistances": "bludgeoning, piercing, slashing while in dim light or darkness"
"senses": "darkvision 60 ft., passive Perception 16"
"languages": "Common"
"cr": "13"
"traits":
- "desc": "The Angry has advantage on Wisdom ([Perception](/3-Mechanics/CLI/rules/skills.md#Perception))\
    \ checks and on saving throws against being [blinded](/3-Mechanics/CLI/rules/conditions.md#blinded),\
    \ [charmed](/3-Mechanics/CLI/rules/conditions.md#charmed), [deafened](/3-Mechanics/CLI/rules/conditions.md#deafened),\
    \ [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened), [stunned](/3-Mechanics/CLI/rules/conditions.md#stunned),\
    \ or knocked [unconscious](/3-Mechanics/CLI/rules/conditions.md#unconscious)."
  "name": "Two Heads"
- "desc": "If another creature deals damage to the Angry, the Angry's attack rolls\
    \ have advantage until the end of its next turn, and the first time it hits with\
    \ a hook attack on its next turn, the attack's target takes an extra dice:3d12|text(19)\
    \ (3d12) psychic damage.\n\nOn its turn, the Angry has disadvantage on attack\
    \ rolls if no other creature has dealt damage to it since the end of its last\
    \ turn."
  "name": "Rising Anger"
"actions":
- "desc": "The Angry makes two hook attacks."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+8 (+8) to hit, reach 5 ft., one target.\
    \ Hit: dice:2d12 + 3|text(16) (2d12 + 3) piercing damage."
  "name": "Hook"
"source":
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MTF/The%20Angry.webp"
```
^statblock

```encounter-table
name: The Angry
creatures:
 - 1: The Angry
```

## Environment

underdark, urban