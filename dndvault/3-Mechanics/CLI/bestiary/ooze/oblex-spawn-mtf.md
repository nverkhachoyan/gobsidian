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
- monster/size/tiny
- monster/type/ooze
aliases: ["Oblex Spawn"]
NoteIcon: monster
BestiaryType: ooze
SourceType: Bestiary
BookSource: Mordenkainen's Tome of Foes p. 217
---
# [Oblex Spawn](3-Mechanics\CLI\bestiary\ooze/oblex-spawn-mtf.md)
*Source: Mordenkainen's Tome of Foes p. 217*  

> [!quote]- A quote from Mordenkainen  
> 
> Mind flayers unleash all manner of foul experiments upon the planes with little thought for the consequences. Here, though, I suspect another influence: Juiblex.

## Oblex Spawn

An oblex devours memories not only to sustain its existence, but also to spawn new oblexes. Each time it fully drains the memories of a victim, it gains the creature's personality—now twisted by the oblex's foul nature. The more memories an oblex steals, the larger it becomes, until it must shed a personality it has absorbed or else go insane. This act spawns a new oblex.

## Oblex

By experimenting on the slimes, jellies, and puddings that infest the depths of the Underdark, mind flayers created a special breed of ooze, the oblex—a slime capable of assaulting the minds of other creatures. Cunning hunters, these pools of jelly stalk prey, searching for the memories they so desperately crave. When oblexes feed on those thoughts, sometimes killing their victims, they can form weird copies of their prey, which help them to harvest even more victims for their dark masters.

## Memory Eaters

Oblexes feed on thoughts and memories. The sharper the mind, the better the meal, so oblexes hunt obviously intelligent targets such as wizards and other spellcasters. When suitable fare comes within reach, an oblex draws its body up to engulf its victim. As it withdraws, it plunders the creature's mind, leaving its prey befuddled and confused.

## Ooze Nature

An oblex doesn't require sleep.

> [!quote]- A quote from Mordenkainen  
> 
> An oblex wants memories, but not to serve any end of its own making. Oblexes are hungry for memories and personalities because they are empty without such nourishment. In this way they serve their creators, the illithids. An oblex in the range of an elder brain's powers provides everything necessary for the mind flayers to find choice victims.


```statblock
"name": "Oblex Spawn (MTF)"
"size": "Tiny"
"type": "ooze"
"alignment": "Lawful Evil"
"ac": !!int "13"
"hp": !!int "18"
"hit_dice": "4d4 + 8"
"stats":
- !!int "8"
- !!int "16"
- !!int "15"
- !!int "14"
- !!int "11"
- !!int "10"
"speed": "20 ft."
"saves":
  "Charisma": !!int "2"
  "Intelligence": !!int "4"
"condition_immunities": "[blinded](/3-Mechanics/CLI/rules/conditions.md#blinded),\
  \ [charmed](/3-Mechanics/CLI/rules/conditions.md#charmed), [deafened](/3-Mechanics/CLI/rules/conditions.md#deafened),\
  \ [exhaustion](/3-Mechanics/CLI/rules/conditions.md#exhaustion), [prone](/3-Mechanics/CLI/rules/conditions.md#prone)"
"senses": "blindsight 60 ft. (blind beyond this distance), passive Perception 12"
"languages": ""
"cr": "1/4"
"traits":
- "desc": "The oblex can move through a space as narrow as 1 inch wide without squeezing."
  "name": "Amorphous"
- "desc": "If the oblex takes fire damage, it has disadvantage on attack rolls and\
    \ ability checks until the end of its next turn."
  "name": "Aversion to Fire"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+5 (+5) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d4 + 3|text(5) (1d4 + 3) bludgeoning damage plus dice:1d4|text(2)\
    \ (1d4) psychic damage."
  "name": "Pseudopod"
"source":
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MTF/Oblex%20Spawn.webp"
```
^statblock

```encounter-table
name: Oblex Spawn
creatures:
 - 1: Oblex Spawn
```

## Environment

swamp, underdark, urban