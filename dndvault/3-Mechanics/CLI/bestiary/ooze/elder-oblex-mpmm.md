---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/10
- monster/environment/swamp
- monster/environment/underdark
- monster/environment/urban
- monster/size/huge
- monster/type/ooze
aliases: ["Elder Oblex"]
NoteIcon: monster
BestiaryType: ooze
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 199, Mordenkainen's Tome of Foes p. 219
---
# [Elder Oblex](3-Mechanics\CLI\bestiary\ooze/elder-oblex-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 199, Mordenkainen's Tome of Foes p. 219*  

```statblock
"name": "Elder Oblex (MPMM)"
"size": "Huge"
"type": "ooze"
"alignment": "Typically  Lawful Evil"
"ac": !!int "16"
"hp": !!int "115"
"hit_dice": "10d12 + 50"
"stats":
- !!int "15"
- !!int "16"
- !!int "21"
- !!int "22"
- !!int "13"
- !!int "18"
"speed": "20 ft."
"saves":
  "Charisma": !!int "8"
  "Intelligence": !!int "10"
"skillsaves":
  "Nature": !!int "10"
  "Deception": !!int "8"
  "Religion": !!int "10"
  "Perception": !!int "5"
  "History": !!int "10"
  "Arcana": !!int "10"
"condition_immunities": "[blinded](/3-Mechanics/CLI/rules/conditions.md#blinded),\
  \ [charmed](/3-Mechanics/CLI/rules/conditions.md#charmed), [deafened](/3-Mechanics/CLI/rules/conditions.md#deafened),\
  \ [exhaustion](/3-Mechanics/CLI/rules/conditions.md#exhaustion), [prone](/3-Mechanics/CLI/rules/conditions.md#prone)"
"senses": "blindsight 60 ft. (blind beyond this distance), passive Perception 15"
"languages": "Common plus six more languages"
"cr": "10"
"traits":
- "desc": "The oblex casts one of the following spells, requiring no spell components\
    \ and using Intelligence as the spellcasting ability (spell save DC 18):\n\nAt\
    \ will: [charm person](/3-Mechanics/CLI/spells/charm-person.md) (as 5th-level\
    \ spell), [detect thoughts](/3-Mechanics/CLI/spells/detect-thoughts.md)\n\n3/day\
    \ each: [dimension door](/3-Mechanics/CLI/spells/dimension-door.md), [dominate\
    \ person](/3-Mechanics/CLI/spells/dominate-person.md), [hypnotic pattern](/3-Mechanics/CLI/spells/hypnotic-pattern.md),\
    \ [telekinesis](/3-Mechanics/CLI/spells/telekinesis.md)"
  "name": "Spellcasting (Psionics)"
- "desc": "The oblex can move through a space as narrow as 1 inch wide without squeezing."
  "name": "Amorphous"
- "desc": "If the oblex takes fire damage, it has disadvantage on attack rolls and\
    \ ability checks until the end of its next turn."
  "name": "Aversion to Fire"
- "desc": "The oblex doesn't require sleep."
  "name": "Unusual Nature"
"actions":
- "desc": "The elder oblex makes two Pseudopod attacks, and it uses Eat Memories."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+7 (+7) to hit, reach 10 ft., one target.\
    \ Hit: dice:4d6 + 3|text(17) (4d6 + 3) bludgeoning damage plus dice:4d6|text(14)\
    \ (4d6) psychic damage."
  "name": "Pseudopod"
- "desc": "The oblex targets one creature it can see within 5 feet of it. The target\
    \ must succeed on a DC 18 Wisdom saving throw or take dice:8d10|text(44) (8d10)\
    \ psychic damage and become memory drained until it finishes a short or long rest\
    \ or until it benefits from the [greater restoration](/3-Mechanics/CLI/spells/greater-restoration.md)\
    \ or [heal](/3-Mechanics/CLI/spells/heal.md) spell. Constructs, Oozes, Plants,\
    \ and Undead succeed on the save automatically.\n\nWhile memory drained, the target\
    \ must roll a dice: d4|avg|noform (d4) and subtract the number rolled from\
    \ any ability check or attack roll it makes. Each time the target is memory drained\
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
    \ faintly of sulfur. The oblex can impersonate dice: 2d6 + 1|avg|noform (2d6\
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
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Elder%20Oblex.webp"
```
^statblock

```encounter-table
name: Elder Oblex
creatures:
 - 1: Elder Oblex
```

## Environment

swamp, underdark, urban