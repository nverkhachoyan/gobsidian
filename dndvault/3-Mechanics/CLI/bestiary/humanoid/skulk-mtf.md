---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mtf
- monster/cr/1-2
- monster/environment/coastal
- monster/environment/forest
- monster/environment/swamp
- monster/environment/underdark
- monster/environment/urban
- monster/size/medium
- monster/type/humanoid
aliases: ["Skulk"]
NoteIcon: monster
BestiaryType: humanoid
SourceType: Bestiary
BookSource: Mordenkainen's Tome of Foes p. 227
---
# [Skulk](3-Mechanics\CLI\bestiary\humanoid/skulk-mtf.md)
*Source: Mordenkainen's Tome of Foes p. 227*  

> [!quote]- A quote from Mordenkainen  
> 
> Some children have imaginary friends that their parents can't see. Sometimes those invisible friends aren't imaginary.

## Skulk

Skulks are the soulless shells of travelers who became lost in the Shadowfell, wandering its gray wastes until they lost all sense of self. They are so devoid of identity that they have become permanently invisible. Only children can see a skulk without the help of a mirror or a special candle. On the rare occasions when a skulk is visible, it appears as a drab, featureless, hairless humanoid.

## Summoned Servants

A skulk can be summoned from the Shadowfell by performing a ritual. If the creature is given a portion of the summoner's identity, the skulk is bound to obey the summoner's commands for 30 days. If a skulk is visible, an astute observer might deduce who summoned it, because a skulk assumes a vague likeness of its master.

Cruel and chaotic, skulks carry out their orders in the most violent manner possible. A summoned skulk can't return to the Shadowfell until it dies, so it has every motivation to throw itself into creating bloodshed and mayhem.

## Hollow Lives

After killing a person in the material world, a skulk sometimes takes up a silent imitation of that person's life. In extreme cases, skulks have invaded villages, killed all the occupants, and turned the places into seeming ghost towns, where flavorless food is prepared daily, colorless clothes are hung up to dry, and livestock is shifted from pen to pen until it starves.

```statblock
"name": "Skulk (MTF)"
"size": "Medium"
"type": "humanoid"
"alignment": "Chaotic Neutral"
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
"damage_immunities": "radiant"
"condition_immunities": "[blinded](/3-Mechanics/CLI/rules/conditions.md#blinded)"
"senses": "darkvision 120 ft., passive Perception 8"
"languages": "understands Common but can't speak"
"cr": "1/2"
"traits":
- "desc": "The skulk is [invisible](/3-Mechanics/CLI/rules/conditions.md#invisible).\
    \ This invisibility can be circumvented by three things:\n\n- The skulk appears\
    \ as a drab, smooth-skinned humanoid if its reflection can be seen in a mirror\
    \ or on another surface.  \n- The skulk appears as a dim, translucent form in\
    \ the light of a candle made of fat rendered from a corpse whose identity is unknown.\
    \  \n- Humanoid children, aged 10 and under, can see through this invisibility.\
    \  "
  "name": "Fallible Invisibility"
- "desc": "The skulk leaves no tracks to indicate where it has been or where it's\
    \ headed."
  "name": "Trackless"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+6 (+6) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d4 + 4|text(6) (1d4 + 4) slashing damage. If the skulk has\
    \ advantage on the attack roll, the target also takes dice:2d6|text(7) (2d6)\
    \ necrotic damage."
  "name": "Claws"
"source":
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MTF/Skulk.webp"
```
^statblock

```encounter-table
name: Skulk
creatures:
 - 1: Skulk
```

## Environment

coastal, forest, swamp, underdark, urban