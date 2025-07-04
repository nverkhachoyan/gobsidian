---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mm
- monster/cr/1-8
- monster/environment/desert
- monster/environment/forest
- monster/environment/grassland
- monster/environment/urban
- monster/size/tiny
- monster/type/beast
aliases: ["Flying Snake"]
NoteIcon: monster
BestiaryType: beast
SourceType: Bestiary
BookSource: Monster Manual p. 322. Available in the SRD and the Basic Rules.
---
# [Flying Snake](3-Mechanics\CLI\bestiary\beast/flying-snake.md)
*Source: Monster Manual p. 322. Available in the SRD and the Basic Rules.*  

A flying snake is a brightly colored, winged serpent found in remote jungles. Tribespeople and cultists sometimes domesticate flying snakes to serve as messengers that deliver scrolls wrapped in their coils.

```statblock
"name": "Flying Snake"
"size": "Tiny"
"type": "beast"
"alignment": "Unaligned"
"ac": !!int "14"
"hp": !!int "5"
"hit_dice": "2d4"
"stats":
- !!int "4"
- !!int "18"
- !!int "11"
- !!int "2"
- !!int "12"
- !!int "5"
"speed": "30 ft., fly 60 ft., swim 30 ft."
"senses": "blindsight 10 ft., passive Perception 11"
"languages": ""
"cr": "1/8"
"traits":
- "desc": "The snake doesn't provoke opportunity attacks when it flies out of an enemy's\
    \ reach."
  "name": "Flyby"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+6 (+6) to hit, reach 5 ft., one target.\
    \ Hit: 1 piercing damage plus dice:3d4|text(7) (3d4) poison damage."
  "name": "Bite"
"source":
- "MM"
- "SKT"
- "ToA"
- "WDH"
- "IDRotF"
- "KftGV"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MM/Flying%20Snake.webp"
```
^statblock

```encounter-table
name: Flying Snake
creatures:
 - 1: Flying Snake
```

## Environment

grassland, forest, urban, desert