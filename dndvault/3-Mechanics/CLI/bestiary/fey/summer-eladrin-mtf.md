---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mtf
- monster/cr/10
- monster/environment/desert
- monster/environment/forest
- monster/size/medium
- monster/type/fey/elf
aliases: ["Summer Eladrin"]
NoteIcon: monster
BestiaryType: fey (elf)
SourceType: Bestiary
BookSource: Mordenkainen's Tome of Foes p. 196
---
# [Summer Eladrin](3-Mechanics\CLI\bestiary\fey/summer-eladrin-mtf.md)
*Source: Mordenkainen's Tome of Foes p. 196*  

> [!quote]- A quote from Mordenkainen  
> 
> To be not just ruled by emotions, but physically changed by them? Better to have no emotions at all.

## Summer Eladrin

When they are angered, eladrin enter the season of summer, a burning, tempestuous state that transforms them into aggressive warriors eager to vent their wrath. Their magic responds to their fury and amplifies their fighting ability, which helps them move with astonishing quickness and strike with terrible force.

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
"name": "Summer Eladrin (MTF)"
"size": "Medium"
"type": "fey"
"subtype": "elf"
"alignment": "Chaotic Neutral"
"ac": !!int "19"
"ac_class": "natural armor"
"hp": !!int "127"
"hit_dice": "17d8 + 51"
"stats":
- !!int "19"
- !!int "21"
- !!int "16"
- !!int "14"
- !!int "12"
- !!int "18"
"speed": "50 ft."
"skillsaves":
  "Intimidation": !!int "8"
  "Athletics": !!int "8"
"damage_resistances": "bludgeoning, piercing, slashing from nonmagical attacks"
"senses": "darkvision 60 ft., passive Perception 11"
"languages": "Common, Elvish, Sylvan"
"cr": "10"
"traits":
- "desc": "Any non-eladrin creature that starts its turn within 60 feet of the eladrin\
    \ must make a DC 16 Wisdom saving throw. On a failed save, the creature becomes\
    \ [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened) for 1 minute.\
    \ A creature can repeat the saving throw at the end of each of its turns, ending\
    \ the effect on itself on a success. If a creature's saving throw is successful\
    \ or the effect ends for it, the creature is immune to any eladrin's Fearsome\
    \ Presence for the next 24 hours."
  "name": "Fearsome Presence"
- "desc": "As a bonus action, the eladrin can teleport up to 30 feet to an unoccupied\
    \ space it can see."
  "name": "Fey Step (Recharge 4-6)"
- "desc": "The eladrin has advantage on saving throws against spells and other magical\
    \ effects."
  "name": "Magic Resistance"
"actions":
- "desc": "The eladrin makes two weapon attacks."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+8 (+8) to hit, reach 5 ft., one target.\
    \ Hit: dice:2d8 + 4|text(13) (2d8 + 4) slashing damage plus dice:1d8|text(4)\
    \ (1d8) fire damage, or dice:2d10 + 4|text(15) (2d10 + 4) slashing damage\
    \ plus dice:1d8|text(4) (1d8) fire damage if used with two hands."
  "name": "Longsword"
- "desc": "Ranged Weapon Attack: dice: d20+9 (+9) to hit, range 150/600 ft.,\
    \ one target. Hit: dice:2d8 + 5|text(14) (2d8 + 5) piercing damage plus\
    \ dice:1d8|text(4) (1d8) fire damage."
  "name": "Longbow"
"reactions":
- "desc": "The eladrin adds 3 to its AC against one melee attack that would hit it.\
    \ To do so, the eladrin must see the attacker and be wielding a melee weapon."
  "name": "Parry"
"source":
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MTF/Summer%20Eladrin.webp"
```
^statblock

```encounter-table
name: Summer Eladrin
creatures:
 - 1: Summer Eladrin
```

## Environment

desert, forest