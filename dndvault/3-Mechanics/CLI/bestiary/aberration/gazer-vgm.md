---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/vgm
- monster/cr/1-2
- monster/environment/underdark
- monster/size/tiny
- monster/type/aberration
aliases: ["Gazer"]
NoteIcon: monster
BestiaryType: aberration
SourceType: Bestiary
BookSource: Volo's Guide to Monsters p. 126
---
# [Gazer](3-Mechanics\CLI\bestiary\aberration/gazer-vgm.md)
*Source: Volo's Guide to Monsters p. 126*  

A gazer is a tiny manifestation of a beholder's dreams. It resembles the beholder who dreamed it into existence, but its body is only 8 inches wide, and it has only four eyestalks. It follows its creator like a devoted, aggressive puppy, and sometimes small packs of these creatures patrol their master's lair for vermin to kill and lone creatures to harass.

## Nuisance Pet

A gazer can't speak any languages but can approximate mimicking words and sentences in a high-pitched, mocking manner. Beholders find gazers amusing and tolerate their presence like spoiled pets. A gazer can't be tamed by anyone but its creator, except through the use of magic or by bonding with a spellcaster (see sidebar). Some beholders with wizard minions insist they take a gazer as a familiar because they can see through the eyes of these creatures.

## Aggressive Vermin-Eater

A wild gazer (one living separately from a beholder) is territorial, eats bugs and small animals, and is known for playing with its food. A lone gazer avoids picking fights with creatures that are Medium or larger, but a pack of them might take on larger prey. A gazer might follow humanoids in its territory, noisily mimicking their speech and generally being a nuisance, until they leave the area, but it flees if confronted by something it can't kill.

## Variant: Gazer Familiar

Spellcasters who are interested in unusual familiars find that gazers are eager to serve someone who has magical power, especially those who make a point of bullying and harassing others. The gazer behaves aggressively toward creatures smaller than itself, and it tends to randomly attack house pets, farm animals, and even children in town unless its master is very strict. A gazer serving as a familiar has the following trait.

### Familiar

The gazer can serve another creature as a familiar, forming a telepathic bond with its willing master, provided that the master is at least a 3rd-level spellcaster. While the two are bonded, the master can sense what the gazer senses as long as they are within 1 mile of each other. If its master causes it physical harm, the gazer will end its service as a familiar, breaking the telepathic bond.

```statblock
"name": "Gazer (VGM)"
"size": "Tiny"
"type": "aberration"
"alignment": "Neutral Evil"
"ac": !!int "13"
"hp": !!int "13"
"hit_dice": "3d4 + 6"
"stats":
- !!int "3"
- !!int "17"
- !!int "14"
- !!int "3"
- !!int "10"
- !!int "7"
"speed": "0 ft., fly 30 ft. (hover)"
"saves":
  "Wisdom": !!int "2"
"skillsaves":
  "Stealth": !!int "5"
  "Perception": !!int "4"
"condition_immunities": "[prone](/3-Mechanics/CLI/rules/conditions.md#prone)"
"senses": "darkvision 60 ft., passive Perception 14"
"languages": ""
"cr": "1/2"
"traits":
- "desc": "As a bonus action, the gazer can move up to its speed toward a hostile\
    \ creature that it can see."
  "name": "Aggressive"
- "desc": "The gazer can mimic simple sounds of speech it has heard, in any language.\
    \ A creature that hears the sounds can tell they are imitations with a successful\
    \ DC 10 Wisdom ([Insight](/3-Mechanics/CLI/rules/skills.md#Insight)) check."
  "name": "Mimicry"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+5 (+5) to hit, reach 5 ft., one target.\
    \ Hit: 1 piercing damage."
  "name": "Bite"
- "desc": "The gazer shoots two of the following magical eye rays at random (reroll\
    \ duplicates), choosing one or two targets it can see within 60 feet of it:\n\n\
    - 1. Dazing Ray. The targeted creature must succeed on a DC 12 Wisdom saving\
    \ throw or be [charmed](/3-Mechanics/CLI/rules/conditions.md#charmed) until the\
    \ start of the gazer's next turn. While the target is [charmed](/3-Mechanics/CLI/rules/conditions.md#charmed)\
    \ in this way, its speed is halved, and it has disadvantage on attack rolls. \
    \ \n- 2. Fear Ray. The targeted creature must succeed on a DC 12 Wisdom saving\
    \ throw or be [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened) until\
    \ the start of the gazer's next turn.  \n- 3. Frost Ray. The targeted creature\
    \ must succeed on a DC 12 Dexterity saving throw or take dice:3d6|text(10) (3d6)\
    \ cold damage.  \n- 4. Telekinetic Ray. If the target is a creature that is\
    \ Medium or smaller, it must succeed on a DC 12 Strength saving throw or be moved\
    \ up to 30 feet directly away from the gazer.  \n\n    If the target is an object\
    \ weighing 10 pounds or less that isn't being worn or carried, the gazer moves\
    \ it up to 30 feet in any direction. The gazer can also exert fine control on\
    \ objects with this ray, such as manipulating a simple tool or opening a container.\
    \  "
  "name": "Eye Rays"
"source":
- "VGM"
- "WDH"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/VGM/Gazer.webp"
```
^statblock

```encounter-table
name: Gazer
creatures:
 - 1: Gazer
```

## Environment

underdark