---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mtf
- monster/cr/10
- monster/environment/arctic
- monster/environment/forest
- monster/size/medium
- monster/type/fey/elf
aliases: ["Winter Eladrin"]
NoteIcon: monster
BestiaryType: fey (elf)
SourceType: Bestiary
BookSource: Mordenkainen's Tome of Foes p. 197
---
# [Winter Eladrin](3-Mechanics\CLI\bestiary\fey/winter-eladrin-mtf.md)
*Source: Mordenkainen's Tome of Foes p. 197*  

> [!quote]- A quote from Mordenkainen  
> 
> To be not just ruled by emotions, but physically changed by them? Better to have no emotions at all.

## Winter Eladrin

When sorrow distresses eladrin, they enter the winter season, becoming figures of melancholy and bitterness. Frozen tears drop from their cheeks, and their palpable sadness emanates from them as bitter cold.

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
"name": "Winter Eladrin (MTF)"
"size": "Medium"
"type": "fey"
"subtype": "elf"
"alignment": "Chaotic Neutral"
"ac": !!int "19"
"ac_class": "natural armor"
"hp": !!int "127"
"hit_dice": "17d8 + 51"
"stats":
- !!int "11"
- !!int "10"
- !!int "16"
- !!int "18"
- !!int "17"
- !!int "13"
"speed": "30 ft."
"damage_resistances": "cold; bludgeoning, piercing, slashing from nonmagical attacks"
"senses": "darkvision 60 ft., passive Perception 13"
"languages": "Common, Elvish, Sylvan"
"cr": "10"
"traits":
- "desc": "The eladrin's innate spellcasting ability is Intelligence (spell save DC\
    \ 16). It can innately cast the following spells, requiring no material components:\n\
    \nAt will: [fog cloud](/3-Mechanics/CLI/spells/fog-cloud.md), [gust of wind](/3-Mechanics/CLI/spells/gust-of-wind.md)\n\
    \n1/day each: [cone of cold](/3-Mechanics/CLI/spells/cone-of-cold.md), [ice\
    \ storm](/3-Mechanics/CLI/spells/ice-storm.md)"
  "name": "Innate Spellcasting"
- "desc": "As a bonus action, the eladrin can teleport up to 30 feet to an unoccupied\
    \ space it can see."
  "name": "Fey Step (Recharge 4-6)"
- "desc": "The eladrin has advantage on saving throws against spells and other magical\
    \ effects."
  "name": "Magic Resistance"
- "desc": "Any non-eladrin creature that starts its turn within 60 feet of the eladrin\
    \ must make a DC 13 Wisdom saving throw. On a failed save, the creature is [charmed](/3-Mechanics/CLI/rules/conditions.md#charmed)\
    \ for 1 minute. While [charmed](/3-Mechanics/CLI/rules/conditions.md#charmed)\
    \ in this way, the creature has disadvantage on ability checks and saving throws.\
    \ The [charmed](/3-Mechanics/CLI/rules/conditions.md#charmed) creature can repeat\
    \ the saving throw at the end of each of its turns, ending the effect on itself\
    \ on a success. If a creature's saving throw is successful or the effect ends\
    \ for it, the creature is immune to any eladrin's Sorrowful Presence for the next\
    \ 24 hours.\n\nWhenever the eladrin deals damage to the [charmed](/3-Mechanics/CLI/rules/conditions.md#charmed)\
    \ creature, it can repeat the saving throw, ending the effect on itself on a success."
  "name": "Sorrowful Presence"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+4 (+4) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d8|text(4) (1d8) slashing damage, or dice:1d10|text(5) (1d10)\
    \ slashing damage if used with two hands."
  "name": "Longsword"
- "desc": "Ranged Weapon Attack: dice: d20+4 (+4) to hit, range 150/600 ft.,\
    \ one target. Hit: dice:1d8|text(4) (1d8) piercing damage."
  "name": "Longbow"
"reactions":
- "desc": "When the eladrin takes damage from a creature the eladrin can see within\
    \ 60 feet of it, the eladrin can force that creature to succeed on a DC 16 Constitution\
    \ saving throw or take dice:2d10|text(11) (2d10) cold damage."
  "name": "Frigid Rebuke"
"source":
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MTF/Winter%20Eladrin.webp"
```
^statblock

```encounter-table
name: Winter Eladrin
creatures:
 - 1: Winter Eladrin
```

## Environment

arctic, forest