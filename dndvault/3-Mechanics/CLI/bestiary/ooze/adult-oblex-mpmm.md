---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
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
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 198, Mordenkainen's Tome of Foes p. 218
---
# [Adult Oblex](3-Mechanics\CLI\bestiary\ooze/adult-oblex-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 198, Mordenkainen's Tome of Foes p. 218*  

```statblock
"name": "Adult Oblex (MPMM)"
"size": "Medium"
"type": "ooze"
"alignment": "Typically  Lawful Evil"
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
- "desc": "The oblex casts one of the following spells, requiring no spell components\
    \ and using Intelligence as the spellcasting ability (spell save DC 15):\n\n3/day\
    \ each: [charm person](/3-Mechanics/CLI/spells/charm-person.md) (as 5th-level\
    \ spell), [detect thoughts](/3-Mechanics/CLI/spells/detect-thoughts.md), [hypnotic\
    \ pattern](/3-Mechanics/CLI/spells/hypnotic-pattern.md)"
  "name": "Spellcasting (Psionics)"
- "desc": "The oblex can move through a space as narrow as 1 inch wide without squeezing."
  "name": "Amorphous"
- "desc": "If the oblex takes fire damage, it has disadvantage on attack rolls and\
    \ ability checks until the end of its next turn."
  "name": "Aversion to Fire"
- "desc": "The oblex doesn't require sleep."
  "name": "Unusual Nature"
"actions":
- "desc": "The oblex makes two pseudopod attacks, and it uses Eat Memories."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+7 (+7) to hit, reach 5 ft., one target.\
    \ Hit: dice:2d6 + 4|text(11) (2d6 + 4) bludgeoning damage plus dice:2d6|text(7)\
    \ (2d6) psychic damage."
  "name": "Pseudopod"
- "desc": "The oblex targets one creature it can see within 5 feet of it. The target\
    \ must succeed on a DC 15 Wisdom saving throw or take dice:4d8|text(18) (4d8)\
    \ psychic damage and become memory drained until it finishes a short or long rest\
    \ or until it benefits from the [greater restoration](/3-Mechanics/CLI/spells/greater-restoration.md)\
    \ or [heal](/3-Mechanics/CLI/spells/heal.md) spell. Constructs, Oozes, Plants,\
    \ and Undead succeed on the save automatically.\n\nWhile memory drained, the target\
    \ must roll a dice: d4|avg|noform (d4) and subtract the number rolled from\
    \ its ability checks and attack rolls. Each time the target is memory drained\
    \ beyond the first, the die size increases by one: the dice: d4|avg|noform (d4)\
    \ becomes a dice: d6|avg|noform (d6), the dice: d6|avg|noform (d6) becomes\
    \ a dice: d8|avg|noform (d8), and so on until the die becomes a dice: d20|avg|noform\
    \ (d20), at which point the target becomes [unconscious](/3-Mechanics/CLI/rules/conditions.md#unconscious)\
    \ for 1 hour. The effect then ends.\n\nThe oblex learns all the languages a memory-drained\
    \ target knows and gains all its skill proficiencies."
  "name": "Eat Memories"
"bonus_actions":
- "desc": "The oblex extrudes a piece of itself that assumes the appearance of one\
    \ Medium or smaller creature whose memories it has stolen. This simulacrum appears,\
    \ feels, and sounds exactly like the creature it impersonates, though it smells\
    \ faintly of sulfur. The oblex can impersonate dice: 1d4 + 1|avg|noform (1d4\
    \ + 1) different creatures, each one tethered to its body by a strand of slime\
    \ that can extend up to 120 feet away. The simulacrum is an extension of the oblex,\
    \ meaning that the oblex occupies its space and the simulacrum's space simultaneously.\
    \ The tether is immune to damage, but it is severed if there is no opening at\
    \ least 1 inch wide between the oblex and the simulacrum. The simulacrum disappears\
    \ if the tether is severed."
  "name": "Sulfurous Impersonation"
"source":
- "MPMM"
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Adult%20Oblex.webp"
```
^statblock

```encounter-table
name: Adult Oblex
creatures:
 - 1: Adult Oblex
```

## Environment

swamp, underdark, urban