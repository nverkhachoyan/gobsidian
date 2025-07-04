---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mtf
- monster/cr/18
- monster/environment/underdark
- monster/size/huge
- monster/type/fiend/demon
aliases: ["Sibriex"]
NoteIcon: monster
BestiaryType: fiend (demon)
SourceType: Bestiary
BookSource: Mordenkainen's Tome of Foes p. 137
---
# [Sibriex](3-Mechanics\CLI\bestiary\fiend/sibriex-mtf.md)
*Source: Mordenkainen's Tome of Foes p. 137*  

> [!quote]- A quote from Mordenkainen  
> 
> No creature embodies the chaotic nature of the Abyss so well as the sibriex. Although the realm of the demons is already a place of infinite horrors, sibriexes for some reason make even more of these monstrosities.

## Sibriex

Thought to be as old as the Abyss itself, sibriexes haunt remote parts of the plane, where they use their vile abilities to breed new horrors and apprehend forbidden lore. Rivulets of blood and bile cascade from a sibriex's body. Where these noxious fluids hit the ground, the land scape becomes polluted.

## Keepers of Forbidden Lore

Sibriexes have spent eons amassing information from across the planes, hoarding knowledge for when it might be useful. Such is their incredible intellect that many seek them out, including demon lords. Some sibriexes act as advisors and oracles, manipulating demons into serving their ends, while other sibriexes cling to their secrets, parceling out lore only when doing so would advance their plans.

## Demon Crafters

Sibriexes can channel the power of the Abyss to create new demons from other creatures. Over the course of days, they can create vast numbers of rutterkins to protect their lands and to ensure that the plane teems with destructive monsters. Some demons petition sibriexes for physical gifts, and if they are moved to do so, sibriexes can graft on new body parts to give the demons greater strength, vision, or stamina. Sibriexes never give aid freely; they demand a service or a treasure in return for the flesh-shaping they provide.

```statblock
"name": "Sibriex (MTF)"
"size": "Huge"
"type": "fiend"
"subtype": "demon"
"alignment": "Chaotic Evil"
"ac": !!int "19"
"ac_class": "natural armor"
"hp": !!int "150"
"hit_dice": "12d12 + 72"
"stats":
- !!int "10"
- !!int "3"
- !!int "23"
- !!int "25"
- !!int "24"
- !!int "25"
"speed": "0 ft., fly 20 ft. (hover)"
"saves":
  "Charisma": !!int "13"
  "Intelligence": !!int "13"
"skillsaves":
  "Perception": !!int "13"
  "History": !!int "13"
  "Arcana": !!int "13"
"damage_resistances": "cold; fire; lightning; bludgeoning, piercing, slashing from\
  \ nonmagical attacks"
"damage_immunities": "poison"
"condition_immunities": "[poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)"
"senses": "truesight 120 ft., passive Perception 23"
"languages": "all, telepathy 120 ft."
"cr": "18"
"traits":
- "desc": "The sibriex's innate spellcasting ability is Charisma (spell save DC 21).\
    \ It can innately cast the following spells, requiring no material components:\n\
    \nAt will: [charm person](/3-Mechanics/CLI/spells/charm-person.md), [command](/3-Mechanics/CLI/spells/command.md),\
    \ [dispel magic](/3-Mechanics/CLI/spells/dispel-magic.md), [hold monster](/3-Mechanics/CLI/spells/hold-monster.md)\n\
    \n3/day: [feeblemind](/3-Mechanics/CLI/spells/feeblemind.md)"
  "name": "Innate Spellcasting"
- "desc": "The sibriex emits an aura of corruption 30 feet in every direction. Plants\
    \ that aren't creatures wither in the aura, and the ground in it is difficult\
    \ terrain for other creatures. Any creature that starts its turn in the aura must\
    \ succeed on a DC 20 Constitution saving throw or take dice:4d6|text(14) (4d6)\
    \ poison damage. A creature that succeeds on the save is immune to this sibriex's\
    \ Contamination for 24 hours."
  "name": "Contamination"
- "desc": "If the sibriex fails a saving throw, it can choose to succeed instead."
  "name": "Legendary Resistance (3/Day)"
- "desc": "The sibriex has advantage on saving throws against spells and other magical\
    \ effects."
  "name": "Magic Resistance"
"actions":
- "desc": "The sibriex uses Squirt Bile once and makes three attacks using its chain,\
    \ bite, or both."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+6 (+6) to hit, reach 15 ft., one target.\
    \ Hit: dice:2d12 + 7|text(20) (2d12 + 7) piercing damage."
  "name": "Chain"
- "desc": "Melee Weapon Attack: dice: d20+6 (+6) to hit, reach 5 ft., one target.\
    \ Hit: dice:2d8|text(9) (2d8) piercing damage plus dice:2d8|text(9) (2d8)\
    \ acid damage."
  "name": "Bite"
- "desc": "The sibriex targets one creature it can see within 120 feet of it. The\
    \ target must succeed on a DC 20 Dexterity saving throw or take dice:10d6|text(35)\
    \ (10d6) acid damage."
  "name": "Squirt Bile"
- "desc": "The sibriex targets up to three creatures it can see within 120 feet of\
    \ it. Each target must make a DC 20 Constitution saving throw. On a successful\
    \ save, a creature becomes immune to this sibriex's Warp Creature. On a failed\
    \ save, the target is [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned),\
    \ which causes it to also gain 1 level of [exhaustion](/3-Mechanics/CLI/rules/conditions.md#exhaustion).\
    \ While [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned) in this way,\
    \ the target must repeat the saving throw at the start of each of its turns. Three\
    \ successful saves against the poison end it, and ending the poison removes any\
    \ levels of [exhaustion](/3-Mechanics/CLI/rules/conditions.md#exhaustion) caused\
    \ by it. Each failed save causes the target to suffer another level of [exhaustion](/3-Mechanics/CLI/rules/conditions.md#exhaustion).\
    \ Once the target reaches 6 levels of [exhaustion](/3-Mechanics/CLI/rules/conditions.md#exhaustion),\
    \ it dies and instantly transforms into a living [abyssal wretch](/3-Mechanics/CLI/bestiary/fiend/abyssal-wretch-mtf.md)\
    \ under the sibriex's control. The transformation of the body can be undone only\
    \ by a [wish](/3-Mechanics/CLI/spells/wish.md) spell."
  "name": "Warp Creature"
"legendary_actions":
- "desc": "The sibriex casts a spell."
  "name": "Cast a Spell"
- "desc": "The sibriex uses Squirt Bile."
  "name": "Spray Bile"
- "desc": "The sibriex uses Warp Creature."
  "name": "Warp (Costs 2 Actions)"
"source":
- "MTF"
- "BGDIA"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MTF/Sibriex.webp"
```
^statblock

```encounter-table
name: Sibriex
creatures:
 - 1: Sibriex
```

## Environment

underdark