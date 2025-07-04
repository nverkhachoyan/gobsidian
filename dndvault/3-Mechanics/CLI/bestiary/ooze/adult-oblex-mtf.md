---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mtf
- monster/cr/5
- monster/environment/swamp
- monster/environment/underdark
- monster/environment/urban
- monster/size/medium
- monster/type/ooze
aliases: ["Adult Oblex"]
NoteIcon: monster
BestiaryType: ooze
SourceType: Bestiary
BookSource: Mordenkainen's Tome of Foes p. 218
---
# [Adult Oblex](3-Mechanics\CLI\bestiary\ooze/adult-oblex-mtf.md)
*Source: Mordenkainen's Tome of Foes p. 218*  

> [!quote]- A quote from Mordenkainen  
> 
> Mind flayers unleash all manner of foul experiments upon the planes with little thought for the consequences. Here, though, I suspect another influence: Juiblex.

## Adult and Elder Oblexes

Newly formed oblexes lack the capabilities of their older kin. They seek only to feed on memories and grow until they can impersonate their victims.

Older oblexes, called adults and elders, have eaten so many memories that they can form duplicates of the creatures they have devoured from the substance of their bodies, sending them off to lure prey into their clutches, while remaining tethered to the slime by long tendrils of goo. These duplicated creatures are indistinguishable from their victims except for a faint sulfurous smell. Oblexes use these duplicates to lure prey into danger or to infiltrate settlements so they can feed on superior victims.

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
"name": "Adult Oblex (MTF)"
"size": "Medium"
"type": "ooze"
"alignment": "Lawful Evil"
"ac": !!int "14"
"hp": !!int "75"
"hit_dice": "10d8 + 30"
"stats":
- !!int "8"
- !!int "19"
- !!int "16"
- !!int "19"
- !!int "12"
- !!int "15"
"speed": "20 ft."
"saves":
  "Charisma": !!int "5"
  "Intelligence": !!int "7"
"skillsaves":
  "Nature": !!int "7"
  "Deception": !!int "5"
  "Religion": !!int "7"
  "Perception": !!int "4"
  "History": !!int "7"
  "Arcana": !!int "7"
"condition_immunities": "[blinded](/3-Mechanics/CLI/rules/conditions.md#blinded),\
  \ [charmed](/3-Mechanics/CLI/rules/conditions.md#charmed), [deafened](/3-Mechanics/CLI/rules/conditions.md#deafened),\
  \ [exhaustion](/3-Mechanics/CLI/rules/conditions.md#exhaustion), [prone](/3-Mechanics/CLI/rules/conditions.md#prone)"
"senses": "blindsight 60 ft. (blind beyond this distance), passive Perception 14"
"languages": "Common plus two more languages"
"cr": "5"
"traits":
- "desc": "The oblex's innate spellcasting ability is Intelligence (spell save DC\
    \ 15). It can innately cast the following spells, requiring no components:\n\n\
    3/day each: [charm person](/3-Mechanics/CLI/spells/charm-person.md) (as 5th-level\
    \ spell), [color spray](/3-Mechanics/CLI/spells/color-spray.md), [detect thoughts](/3-Mechanics/CLI/spells/detect-thoughts.md),\
    \ [hold person](/3-Mechanics/CLI/spells/hold-person.md) (as 3rd-level spell)"
  "name": "Innate Spellcasting"
- "desc": "The oblex can move through a space as narrow as 1 inch wide without squeezing."
  "name": "Amorphous"
- "desc": "If the oblex takes fire damage, it has disadvantage on attack rolls and\
    \ ability checks until the end of its next turn."
  "name": "Aversion to Fire"
- "desc": "As a bonus action, the oblex can extrude a piece of itself that assumes\
    \ the appearance of one Medium or smaller creature whose memories it has stolen.\
    \ This simulacrum appears, feels, and sounds exactly like the creature it impersonates,\
    \ though it smells faintly of sulfur. The oblex can impersonate dice: 1d4 + 1|avg|noform\
    \ (1d4 + 1) different creatures, each one tethered to its body by a strand of\
    \ slime that can extend up to 120 feet away. For all practical purposes, the simulacrum\
    \ is the oblex, meaning that the oblex occupies its space and the simulacrum's\
    \ space simultaneously. The slimy tether is immune to damage, but it is severed\
    \ if there is no opening at least 1 inch wide between the oblex's main body and\
    \ the simulacrum. The simulacrum disappears if the tether is severed."
  "name": "Sulfurous Impersonation"
"actions":
- "desc": "The oblex makes one pseudopod attack and uses Eat Memories."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+7 (+7) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d6 + 4|text(7) (1d6 + 4) bludgeoning damage plus dice:2d4|text(5)\
    \ (2d4) psychic damage."
  "name": "Pseudopod"
- "desc": "The oblex targets one creature it can see within 5 feet of it. The target\
    \ must succeed on a DC 15 Wisdom saving throw or take dice:4d8|text(18) (4d8)\
    \ psychic damage and become memory drained until it finishes a short or long rest\
    \ or until it benefits from the [greater restoration](/3-Mechanics/CLI/spells/greater-restoration.md)\
    \ or [heal](/3-Mechanics/CLI/spells/heal.md) spell. Constructs, oozes, plants,\
    \ and undead succeed on the save automatically.\n\nWhile memory drained, the target\
    \ must roll a dice: d4|avg|noform (d4) and subtract the number rolled from\
    \ any ability check or attack roll it makes. Each time the target is memory drained\
    \ beyond the first, the die size increases by one: the dice: d4|avg|noform (d4)\
    \ becomes a dice: d6|avg|noform (d6), the dice: d6|avg|noform (d6) becomes\
    \ a dice: d8|avg|noform (d8), and so on until the die becomes a dice: d20|avg|noform\
    \ (d20), at which point the target becomes [unconscious](/3-Mechanics/CLI/rules/conditions.md#unconscious)\
    \ for 1 hour. The effect then ends.\n\nWhen an oblex causes a target to become\
    \ memory drained, the oblex learns all the languages the target knows and gains\
    \ all its proficiencies, except for any saving throw proficiencies."
  "name": "Eat Memories"
"source":
- "MTF"
- "IMR"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MTF/Adult%20Oblex.webp"
```
^statblock

```encounter-table
name: Adult Oblex
creatures:
 - 1: Adult Oblex
```

## Environment

swamp, underdark, urban