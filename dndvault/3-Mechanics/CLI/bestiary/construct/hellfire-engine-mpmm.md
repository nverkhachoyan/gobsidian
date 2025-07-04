---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/16
- monster/size/huge
- monster/type/construct
aliases: ["Hellfire Engine"]
NoteIcon: monster
BestiaryType: construct
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 152, Mordenkainen's Tome of Foes p. 165
---
# [Hellfire Engine](3-Mechanics\CLI\bestiary\construct/hellfire-engine-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 152, Mordenkainen's Tome of Foes p. 165*  

```statblock
"name": "Hellfire Engine (MPMM)"
"size": "Huge"
"type": "construct"
"alignment": "Typically  Lawful Evil"
"ac": !!int "18"
"ac_class": "natural armor"
"hp": !!int "216"
"hit_dice": "16d12 + 112"
"stats":
- !!int "20"
- !!int "16"
- !!int "24"
- !!int "2"
- !!int "10"
- !!int "1"
"speed": "40 ft."
"saves":
  "Charisma": !!int "0"
  "Dexterity": !!int "8"
  "Wisdom": !!int "5"
"damage_resistances": "cold; psychic; bludgeoning, piercing, slashing from nonmagical\
  \ attacks that aren't silvered"
"damage_immunities": "fire, poison"
"condition_immunities": "[charmed](/3-Mechanics/CLI/rules/conditions.md#charmed),\
  \ [deafened](/3-Mechanics/CLI/rules/conditions.md#deafened), [exhaustion](/3-Mechanics/CLI/rules/conditions.md#exhaustion),\
  \ [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened), [paralyzed](/3-Mechanics/CLI/rules/conditions.md#paralyzed),\
  \ [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned), [unconscious](/3-Mechanics/CLI/rules/conditions.md#unconscious)"
"senses": "darkvision 120 ft., passive Perception 10"
"languages": "understands Infernal but can't speak"
"cr": "16"
"traits":
- "desc": "The hellfire engine is immune to any spell or effect that would alter its\
    \ form."
  "name": "Immutable Form"
- "desc": "The hellfire engine has advantage on saving throws against spells and other\
    \ magical effects."
  "name": "Magic Resistance"
- "desc": "The hellfire engine doesn't require air, food, drink, or sleep."
  "name": "Unusual Nature"
"actions":
- "desc": "The hellfire engine moves up to its speed in a straight line. During this\
    \ move, it can enter Large or smaller creatures' spaces. A creature whose space\
    \ the hellfire engine enters must make a DC 18 Dexterity saving throw. On a successful\
    \ save, the creature is pushed to the nearest space out of the hellfire engine's\
    \ path. On a failed save, the creature falls [prone](/3-Mechanics/CLI/rules/conditions.md#prone)\
    \ and takes dice:8d6|text(28) (8d6) bludgeoning damage.\n\nIf the hellfire\
    \ engine remains in the [prone](/3-Mechanics/CLI/rules/conditions.md#prone) creature's\
    \ space, the creature is also [restrained](/3-Mechanics/CLI/rules/conditions.md#restrained)\
    \ until it's no longer in the same space as the hellfire engine. While [restrained](/3-Mechanics/CLI/rules/conditions.md#restrained)\
    \ in this way, the creature, or another creature within 5 feet of it, can make\
    \ a DC 18 Strength check. On a success, the creature is shunted to an unoccupied\
    \ space of its choice within 5 feet of the hellfire engine and is no longer [restrained](/3-Mechanics/CLI/rules/conditions.md#restrained)."
  "name": "Flesh-Crushing Stride"
- "desc": "The hellfire engine uses one of the following options (choose one or roll\
    \ a dice: d6|avg|noform (d6)):\n\n- 1–2 Bonemelt Sprayer. The hellfire\
    \ engine spews acidic flame in a 60-foot cone. Each creature in the cone must\
    \ make a DC 20 Dexterity saving throw, taking dice:2d10|text(11) (2d10) fire\
    \ damage plus dice:4d8|text(18) (4d8) acid damage on a failed save, or half\
    \ as much damage on a successful one. Creatures that fail the saving throw are\
    \ drenched in burning acid and take dice:1d10|text(5) (1d10) fire damage plus\
    \ dice:2d8|text(9) (2d8) acid damage at the end of their turns. An affected\
    \ creature or another creature within 5 feet of it can take an action to scrape\
    \ off the burning fuel.  \n- 3–4 Lightning Flail. Melee Weapon Attack: dice:\
    \ d20+10 (+10) to hit, reach 15 ft., one creature. Hit: dice:3d8 + 5|text(18)\
    \ (3d8 + 5) bludgeoning damage plus dice:5d8|text(22) (5d8) lightning damage.\
    \ Up to three other creatures of the hellfire engine's choice that it can see\
    \ within 30 feet of the target must each make a DC 20 Dexterity saving throw,\
    \ taking dice:5d8|text(22) (5d8) lightning damage on a failed save, or half\
    \ as much damage on a successful one.  \n- 5–6 Thunder Cannon. The hellfire\
    \ engine targets a point within 120 feet of it that it can see. Each creature\
    \ within 30 feet of that point must make a DC 20 Dexterity saving throw, taking\
    \ dice:5d10|text(27) (5d10) bludgeoning damage plus dice:3d12|text(19) (3d12)\
    \ thunder damage on a failed save, or half as much damage on a successful one.\
    \  \n\n    If the chosen option kills a creature, the creature's soul rises from\
    \ the River Styx as a [lemure](/3-Mechanics/CLI/bestiary/fiend/lemure.md) in Avernus\
    \ in dice: 1d4|avg|noform (1d4) hours. If the creature isn't revived before\
    \ then, only a [wish](/3-Mechanics/CLI/spells/wish.md) spell or killing the [lemure](/3-Mechanics/CLI/bestiary/fiend/lemure.md)\
    \ and casting true resurrection on the creature's original body can restore it\
    \ to life. Constructs and devils are immune to this effect.  "
  "name": "Hellfire Weapons"
"source":
- "MPMM"
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Hellfire%20Engine.webp"
```
^statblock

```encounter-table
name: Hellfire Engine
creatures:
 - 1: Hellfire Engine
```