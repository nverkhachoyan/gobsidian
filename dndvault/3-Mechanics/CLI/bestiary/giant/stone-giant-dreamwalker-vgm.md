---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/vgm
- monster/cr/10
- monster/environment/coastal
- monster/environment/hill
- monster/environment/mountain
- monster/size/huge
- monster/type/giant/stone-giant
aliases: ["Stone Giant Dreamwalker"]
NoteIcon: monster
BestiaryType: giant (stone giant)
SourceType: Bestiary
BookSource: Volo's Guide to Monsters p. 150
---
# [Stone Giant Dreamwalker](3-Mechanics\CLI\bestiary\giant/stone-giant-dreamwalker-vgm.md)
*Source: Volo's Guide to Monsters p. 150*  

The surface of the world is an alien realm to stone giants: fluctuating, temporary, exposed to gusting wind and sudden rain. It is as wildly changeable as a dream, and that's how they regard it-as a dream. Nothing there is permanent, so nothing there is real. What happens on the surface doesn't matter. Promises and bargains made there needn't be honored. Life and even art hold less value there.

## Dream Dwellers

Stone giants sometimes go on dream quests in the surface world, seeking inspiration for their art, to break a decades-long ennui, or out of simple curiosity. Some who go on these quests let themselves become lost in the dream. Other stone giants are banished to the surface as punishment. Regardless of the reason, if they don't take shelter under stone, such individuals can become dreamwalkers.

Dreamwalkers occupy an odd place of respect outside of stone giant ordning. They are considered outcasts, but their familiarity with the surface world makes them valuable guides, and their insights can help other stone giants grasp the dangers of living in a dream.

## Mad Wanderers

Dreamwalkers are driven mad by isolation, shame, and their unendingly alien surroundings, and this madness leeches out into the world around them, affecting other creatures that get too close. Believing that they're living in a dream and that their actions have no real consequences, dreamwalkers act as they please, becoming forces of chaos. As they travel the world, they collect objects and creatures that seem especially significant in their mad minds. Over time, the collected things accrete to their bodies, becoming encased in stone.

```statblock
"name": "Stone Giant Dreamwalker (VGM)"
"size": "Huge"
"type": "giant"
"subtype": "stone giant"
"alignment": "Chaotic Neutral"
"ac": !!int "18"
"ac_class": "natural armor"
"hp": !!int "161"
"hit_dice": "14d12 + 70"
"stats":
- !!int "23"
- !!int "14"
- !!int "21"
- !!int "10"
- !!int "8"
- !!int "12"
"speed": "40 ft."
"saves":
  "Dexterity": !!int "6"
  "Wisdom": !!int "3"
  "Constitution": !!int "9"
"skillsaves":
  "Athletics": !!int "14"
  "Perception": !!int "3"
"condition_immunities": "[charmed](/3-Mechanics/CLI/rules/conditions.md#charmed),\
  \ [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened)"
"senses": "darkvision 60 ft., passive Perception 13"
"languages": "Common, Giant"
"cr": "10"
"traits":
- "desc": "An enemy that starts its turn within 30 feet of the giant must make a DC\
    \ 13 Charisma saving throw, provided that the giant isn't [incapacitated](/3-Mechanics/CLI/rules/conditions.md#incapacitated).\
    \ On a failed save, the creature is [charmed](/3-Mechanics/CLI/rules/conditions.md#charmed)\
    \ by the giant. A creature [charmed](/3-Mechanics/CLI/rules/conditions.md#charmed)\
    \ in this way can repeat the saving throw at the end of each of its turns, ending\
    \ the effect on itself on a success. Once it succeeds on the saving throw, the\
    \ creature is immune to this giant's Dreamwalker's Charm for 24 hours."
  "name": "Dreamwalker's Charm"
"actions":
- "desc": "The giant makes two attacks with its greatclub."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+10 (+10) to hit, reach 15 ft., one\
    \ target. Hit: dice:3d8 + 6|text(19) (3d8 + 6) bludgeoning damage."
  "name": "Greatclub"
- "desc": "The giant touches one Medium or smaller creature within 10 feet of it that\
    \ is [charmed](/3-Mechanics/CLI/rules/conditions.md#charmed) by it. The target\
    \ must make a DC 17 Constitution saving throw. On a failed save, the target becomes\
    \ [petrified](/3-Mechanics/CLI/rules/conditions.md#petrified), and the giant can\
    \ adhere the target to its stony body. [Greater restoration](/3-Mechanics/CLI/spells/greater-restoration.md)\
    \ spells and other magic that can undo petrification have no effect on a [petrified](/3-Mechanics/CLI/rules/conditions.md#petrified)\
    \ creature on the giant unless the giant is dead, in which case the magic works\
    \ normally, freeing the [petrified](/3-Mechanics/CLI/rules/conditions.md#petrified)\
    \ creature as well as ending the [petrified](/3-Mechanics/CLI/rules/conditions.md#petrified)\
    \ condition on it."
  "name": "Petrifying Touch"
- "desc": "Ranged Weapon Attack: dice: d20+10 (+10) to hit, range 60/240 ft.,\
    \ one target. Hit: dice:4d10 + 6|text(28) (4d10 + 6) bludgeoning damage.\
    \ If the target is a creature, it must succeed on a DC 17 Strength saving throw\
    \ or be knocked [prone](/3-Mechanics/CLI/rules/conditions.md#prone)."
  "name": "Rock"
"source":
- "VGM"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/VGM/Stone%20Giant%20Dreamwalker.webp"
```
^statblock

```encounter-table
name: Stone Giant Dreamwalker
creatures:
 - 1: Stone Giant Dreamwalker
```

## Environment

mountain, hill, coastal