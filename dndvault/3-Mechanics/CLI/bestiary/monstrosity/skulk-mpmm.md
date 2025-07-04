---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/1-2
- monster/environment/coastal
- monster/environment/forest
- monster/environment/swamp
- monster/environment/underdark
- monster/environment/urban
- monster/size/medium
- monster/type/monstrosity
aliases: ["Skulk"]
NoteIcon: monster
BestiaryType: monstrosity
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 219, Mordenkainen's Tome of Foes p. 227
---
# [Skulk](3-Mechanics\CLI\bestiary\monstrosity/skulk-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 219, Mordenkainen's Tome of Foes p. 227*  

```statblock
"name": "Skulk (MPMM)"
"size": "Medium"
"type": "monstrosity"
"alignment": "Typically  Chaotic Neutral"
"ac": !!int "14"
"hp": !!int "18"
"hit_dice": "4d8"
"stats":
- !!int "6"
- !!int "19"
- !!int "10"
- !!int "10"
- !!int "7"
- !!int "1"
"speed": "30 ft."
"saves":
  "Constitution": !!int "2"
"skillsaves":
  "Stealth": !!int "8"
"condition_immunities": "[blinded](/3-Mechanics/CLI/rules/conditions.md#blinded)"
"senses": "darkvision 120 ft., passive Perception 8"
"languages": "understands Common but can't speak"
"cr": "1/2"
"traits":
- "desc": "The skulk is [invisible](/3-Mechanics/CLI/rules/conditions.md#invisible).\
    \ This invisibility can be circumvented by three things:\n\n- Charnel Candles.\
    \ The skulk appears as a dim, translucent form in the light of a candle made of\
    \ fat rendered from a corpse whose identity is unknown.  \n- Children. Humanoid\
    \ children, aged 10 and under, can see through this invisibility.  \n- Reflective\
    \ Surfaces. The skulk appears as a drab, smoothskinned biped if its reflection\
    \ can be seen in a mirror or on another surface.  "
  "name": "Fallible Invisibility"
- "desc": "The skulk leaves no tracks to indicate where it has been or where it's\
    \ headed."
  "name": "Trackless"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+6 (+6) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d4 + 4|text(6) (1d4 + 4) slashing damage plus dice:1d6|text(3)\
    \ (1d6) necrotic damage."
  "name": "Claw"
"source":
- "MPMM"
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Skulk.webp"
```
^statblock

```encounter-table
name: Skulk
creatures:
 - 1: Skulk
```

## Environment

coastal, forest, swamp, underdark, urban