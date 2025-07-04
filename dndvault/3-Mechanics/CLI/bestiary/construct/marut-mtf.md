---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mtf
- monster/cr/25
- monster/size/large
- monster/type/construct/inevitable
aliases: ["Marut"]
NoteIcon: monster
BestiaryType: construct (inevitable)
SourceType: Bestiary
BookSource: Mordenkainen's Tome of Foes p. 213
---
# [Marut](3-Mechanics\CLI\bestiary\construct/marut-mtf.md)
*Source: Mordenkainen's Tome of Foes p. 213*  

## Marut

The nigh-unstoppable inevitables serve a singular purpose: they enforce contracts forged in the Hall of Concordance in the city of Sigil. Primus, the leader of the modrons, created maruts and other inevitables to bring order to dealings between planar folk. Many creatures, including yugoloths, will enter into a contract with inevitables if asked.

## Cosmic Enforcers

The Hall of Concordance is an embassy of pure law in Sigil, the City of Doors. In the hall, two parties who agree to mutual terms—and who pay the requisite gold to the Kolyarut, a mechanical engine of absolute jurisprudence—can have their contract chiseled onto a sheet of gold that is placed in the chest of a marut. From that moment until the contract is fulfilled, the marut is bound to enforce its terms and to punish any party who breaks them. A marut resorts to lethal force only when a contract calls for it, when the contract is fully broken, or when the marut is attacked.

## Word Is Law

Inevitables care nothing for the spirit of an agreement, only the letter. A marut enforces what is written, not what was meant by or supposed to be understood from the writing. The Kolyarut rejects contracts that contain vague, contradictory, or unenforceable terms. Beyond that, it doesn't care whether both parties understand what they're agreeing to. A small army of solicitors waits outside the Hall of Concordance, eager to sell their expertise in the crafting or vetting of contracts.

## Constructed Nature

A marut doesn't require air, food, drink, or sleep.

```statblock
"name": "Marut (MTF)"
"size": "Large"
"type": "construct"
"subtype": "inevitable"
"alignment": "Lawful Neutral"
"ac": !!int "22"
"ac_class": "natural armor"
"hp": !!int "432"
"hit_dice": "32d10 + 256"
"stats":
- !!int "28"
- !!int "12"
- !!int "26"
- !!int "19"
- !!int "15"
- !!int "18"
"speed": "40 ft., fly 30 ft. (hover)"
"saves":
  "Charisma": !!int "12"
  "Wisdom": !!int "10"
  "Intelligence": !!int "12"
"skillsaves":
  "Intimidation": !!int "12"
  "Insight": !!int "10"
  "Perception": !!int "10"
"damage_resistances": "thunder; bludgeoning, piercing, slashing from nonmagical attacks"
"damage_immunities": "poison"
"condition_immunities": "[charmed](/3-Mechanics/CLI/rules/conditions.md#charmed),\
  \ [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened), [paralyzed](/3-Mechanics/CLI/rules/conditions.md#paralyzed),\
  \ [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned), [unconscious](/3-Mechanics/CLI/rules/conditions.md#unconscious)"
"senses": "darkvision 60 ft., passive Perception 20"
"languages": "all but rarely speaks"
"cr": "25"
"traits":
- "desc": "The marut's innate spellcasting ability is Intelligence (spell save DC\
    \ 20). The marut can innately cast the following spell, requiring no material\
    \ components.\n\nAt will: [plane shift](/3-Mechanics/CLI/spells/plane-shift.md)\
    \ (self only)"
  "name": "Innate Spellcasting"
- "desc": "The marut is immune to any spell or effect that would alter its form."
  "name": "Immutable Form"
- "desc": "If the marut fails a saving throw, it can choose to succeed instead."
  "name": "Legendary Resistance (3/Day)"
- "desc": "The marut has advantage on saving throws against spells and other magical\
    \ effects."
  "name": "Magic Resistance"
"actions":
- "desc": "The marut makes two slam attacks."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: automatic hit, reach 5 ft., one target. Hit: 60\
    \ force damage, and the target is pushed up to 5 feet away from the marut if it\
    \ is Huge or smaller."
  "name": "Unerring Slam"
- "desc": "Arcane energy emanates from the marut's chest in a 60-foot cube. Every\
    \ creature in that area takes 45 radiant damage. Each creature that takes any\
    \ of this damage must succeed on a DC 20 Wisdom saving throw or be [stunned](/3-Mechanics/CLI/rules/conditions.md#stunned)\
    \ until the end of the marut's next turn."
  "name": "Blazing Edict (Recharge 5-6)"
- "desc": "The marut targets up to two creatures it can see within 60 feet of it.\
    \ Each target must succeed on a DC 20 Charisma saving throw or be teleported to\
    \ a teleportation circle in the Hall of Concordance in Sigil. A target fails automatically\
    \ if it is [incapacitated](/3-Mechanics/CLI/rules/conditions.md#incapacitated).\
    \ If either target is teleported in this way, the marut teleports with it to the\
    \ circle.\n\nAfter teleporting in this way, the marut can't use this action again\
    \ until it finishes a short or long rest."
  "name": "Justify"
"source":
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MTF/Marut.webp"
```
^statblock

```encounter-table
name: Marut
creatures:
 - 1: Marut
```