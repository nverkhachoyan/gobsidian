---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mtf
- monster/cr/10
- monster/environment/forest
- monster/environment/grassland
- monster/size/medium
- monster/type/fey/elf
aliases: ["Spring Eladrin"]
NoteIcon: monster
BestiaryType: fey (elf)
SourceType: Bestiary
BookSource: Mordenkainen's Tome of Foes p. 196
---
# [Spring Eladrin](3-Mechanics\CLI\bestiary\fey/spring-eladrin-mtf.md)
*Source: Mordenkainen's Tome of Foes p. 196*  

> [!quote]- A quote from Mordenkainen  
> 
> To be not just ruled by emotions, but physically changed by them? Better to have no emotions at all.

## Spring-Eladrin

Their hearts filled with joy, spring eladrin cavort through their sylvan realms, their songs and laughter filling the air. These playful eladrin beguile other creatures to fill them with the joy of spring. Their antics can lead other creatures into danger and make mischief for them.

## Eladrin

Eladrin dwell in the verdant splendor of the Feywild. They are related to the elves found on the Material Plane, and resemble them in both their love of beauty and the value they place on personal freedom. But where other elves can temper their wild impulses, eladrin are creatures ruled by emotion—and because of their unique magical nature, they undergo physical changes to match their changes in temperament.

The eladrin have spent centuries in the Feywild, and most of them have become fey creatures as a result. Some of them are still humanoid, similar in that respect to their other elven kin. The eladrin presented here are of the fey variety.

## Creatures of Passion

The magic flowing through eladrin responds to their emotional state by transforming them into different seasonal aspects, with behaviors and capabilities that change with their forms. Some eladrin might remain in a particular aspect for years, while others run through the emotional spectrum each week.

## Lovers of Beauty

Regardless of the aspect they express, eladrin love beauty and surround themselves with lovely things. Eladrin try to possess any objects they find striking. They might seek to own a painting, a statue, or a glittering jewel. When they encounter people of comely form or luminous spirit, they use their magic to delight those folk or, in the case of evil eladrin, to abduct them.

## Changeable Nature

Whenever one of the eladrin presented here finishes a long rest, it can associate itself with a different season, provided it isn't [incapacitated](/3-Mechanics/CLI/rules/conditions.md#incapacitated). When the eladrin makes this change, it uses the stat block of the new season, rather than its old stat block. Any damage the eladrin sustained in its original form applies to the new form, as do any conditions or other ongoing effects affecting it.

```statblock
"name": "Spring Eladrin (MTF)"
"size": "Medium"
"type": "fey"
"subtype": "elf"
"alignment": "Chaotic Neutral"
"ac": !!int "19"
"ac_class": "natural armor"
"hp": !!int "127"
"hit_dice": "17d8 + 51"
"stats":
- !!int "14"
- !!int "16"
- !!int "16"
- !!int "18"
- !!int "11"
- !!int "18"
"speed": "30 ft."
"skillsaves":
  "Deception": !!int "8"
  "Persuasion": !!int "8"
"damage_resistances": "bludgeoning, piercing, slashing from nonmagical attacks"
"senses": "darkvision 60 ft., passive Perception 10"
"languages": "Common, Elvish, Sylvan"
"cr": "10"
"traits":
- "desc": "The eladrin's innate spellcasting ability is Charisma (spell save DC 16).\
    \ It can innately cast the following spells, requiring no material components:\n\
    \nAt will: [charm person](/3-Mechanics/CLI/spells/charm-person.md), [Tasha's\
    \ hideous laughter](/3-Mechanics/CLI/spells/tashas-hideous-laughter.md)\n\n1/day\
    \ each: [hallucinatory terrain](/3-Mechanics/CLI/spells/hallucinatory-terrain.md),\
    \ [Otto's irresistible dance](/3-Mechanics/CLI/spells/ottos-irresistible-dance.md)\n\
    \n3/day each: [confusion](/3-Mechanics/CLI/spells/confusion.md), [enthrall](/3-Mechanics/CLI/spells/enthrall.md),\
    \ [suggestion](/3-Mechanics/CLI/spells/suggestion.md)"
  "name": "Innate Spellcasting"
- "desc": "As a bonus action, the eladrin can teleport up to 30 feet to an unoccupied\
    \ space it can see."
  "name": "Fey Step (Recharge 4-6)"
- "desc": "Any non-eladrin creature that starts its turn within 60 feet of the eladrin\
    \ must make a DC 16 Wisdom saving throw. On a failed save, the creature is [charmed](/3-Mechanics/CLI/rules/conditions.md#charmed)\
    \ for 1 minute. On a successful save, the creature becomes immune to any eladrin's\
    \ Joyful Presence for 24 hours.\n\nWhenever the eladrin deals damage to the [charmed](/3-Mechanics/CLI/rules/conditions.md#charmed)\
    \ creature, it can repeat the saving throw, ending the effect on itself on a success."
  "name": "Joyful Presence"
- "desc": "The eladrin has advantage on saving throws against spells and other magical\
    \ effects."
  "name": "Magic Resistance"
"actions":
- "desc": "The eladrin makes two weapon attacks. The eladrin can cast one spell in\
    \ place of one of these attacks."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+6 (+6) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d8 + 2|text(6) (1d8 + 2) slashing damage plus dice:1d8|text(4)\
    \ (1d8) psychic damage, or dice:1d10 + 2|text(7) (1d10 + 2) slashing damage\
    \ plus dice:1d8|text(4) (1d8) psychic damage if used with two hands."
  "name": "Longsword"
- "desc": "Ranged Weapon Attack: dice: d20+7 (+7) to hit, range 150/600 ft.,\
    \ one target. Hit: dice:1d8 + 3|text(7) (1d8 + 3) piercing damage plus dice:1d8|text(4)\
    \ (1d8) psychic damage."
  "name": "Longbow"
"source":
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MTF/Spring%20Eladrin.webp"
```
^statblock

```encounter-table
name: Spring Eladrin
creatures:
 - 1: Spring Eladrin
```

## Environment

forest, grassland