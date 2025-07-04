---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mm
- monster/cr/13
- monster/environment/underdark
- monster/size/large
- monster/type/aberration
aliases: ["Beholder"]
NoteIcon: monster
BestiaryType: aberration
SourceType: Bestiary
BookSource: Monster Manual p. 28, Tasha's Cauldron of Everything
---
# [Beholder](3-Mechanics\CLI\bestiary\aberration/beholder.md)
*Source: Monster Manual p. 28, Tasha's Cauldron of Everything*  

One glance at a beholder is enough to assess its foul and otherworldly nature. Aggressive, hateful, and greedy, these aberrations dismiss all other creatures as lesser beings, toying with them or destroying them as they choose.

A beholder's spheroid body levitates at all times, and its great bulging eye sits above a wide, toothy maw, while the smaller eyestalks that crown its body twist and turn to keep its foes in sight. When a beholder sleeps, it closes its central eye but leaves its smaller eyes open and alert.

## Xenophobic Isolationists

Enemies abound, or so every beholder believes. Beholders are convinced that other creatures resent them for their brilliance and magical power, even as they dismiss those lesser creatures as crude and disgusting. Beholders always suspect others of plotting against them, even when no other creatures are around.

The disdain a beholder has for other creatures extends to other beholders. Each beholder believes its form to be an ideal, and that any deviation from that form is a flaw in the racial purity of its kind. Beholders vary greatly in their physical forms, making conflict between them inevitable. Some beholders are protected by overlapping chitinous plates. Some have smooth hides. Some have eyestalks that writhe like tentacles, while others' stalks bear crustacean-like joints. Even slight differences of coloration in hide can turn two beholders into lifelong enemies.

## Eye Tyrant

Some beholders manage to channel their xenophobic tendencies into a terrible despotism. Rather than live in isolation, the aptly named eye tyrants enslave those other creatures, founding and controlling vast empires. An eye tyrant sometimes carves out a domain within or under a major city, commanding networks of agents that operate on their master's behalf.

## Alien Lairs

Because they refuse to share territory with others, most beholders withdraw to frigid hills, abandoned ruins, and deep caverns to scheme. A beholder's lair is carved out by its disintegration eye ray, emphasizing vertical passages connecting chambers stacked on top of each other. Such an environment allows a beholder to move freely, even as it prevents intruders from easily creeping about. When intruders do break in, the height of its open ceilings allows a beholder to float up and harry foes on the floor.

As alien as their creator, the rooms in a beholder's lair reflect the creature's arrogance. It festoons its chambers with trophies from the battles it has won, including [petrified](/3-Mechanics/CLI/rules/conditions.md#petrified) adventurers standing frozen in their horrified final moments, pieces of other beholders, and magic items wrested from powerful foes. A beholder judges its own worth by its acquisitions, and it never willingly parts with its treasures.

> [!quote]- A quote from Valkara Ironfeel, dwarf sage  
> 
> Every beholder thinks it is the epitome of beholderkind, and the only thing it fears is that it might be wrong.

## A Beholder's Lair

A beholder's central lair is typically a large, spacious cavern with high ceilings, where it can attack without fear of closing to melee range. A beholder encountered in its lair has a challenge rating of 14 (11,500 XP).

```statblock
"name": "Beholder"
"size": "Large"
"type": "aberration"
"alignment": "Lawful Evil"
"ac": !!int "18"
"ac_class": "natural armor"
"hp": !!int "180"
"hit_dice": "19d10 + 76"
"stats":
- !!int "10"
- !!int "14"
- !!int "18"
- !!int "17"
- !!int "15"
- !!int "17"
"speed": "0 ft., fly 20 ft. (hover)"
"saves":
  "Charisma": !!int "8"
  "Wisdom": !!int "7"
  "Intelligence": !!int "8"
"skillsaves":
  "Perception": !!int "12"
"condition_immunities": "[prone](/3-Mechanics/CLI/rules/conditions.md#prone)"
"senses": "darkvision 120 ft., passive Perception 22"
"languages": "Deep Speech, Undercommon"
"cr": "13"
"traits":
- "desc": "The beholder's central eye creates an area of antimagic, as in the [antimagic\
    \ field](/3-Mechanics/CLI/spells/antimagic-field.md) spell, in a 150-foot cone.\
    \ At the start of each of its turns, the beholder decides which way the cone faces\
    \ and whether the cone is active. The area works against the beholder's own eye\
    \ rays."
  "name": "Antimagic Cone"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+5 (+5) to hit, reach 5 ft., one target.\
    \ Hit: dice:4d6|text(14) (4d6) piercing damage."
  "name": "Bite"
- "desc": "The beholder shoots three of the following magical eye rays at random (reroll\
    \ duplicates), choosing one to three targets it can see within 120 feet of it:\n\
    \n- 1. Charm Ray. The targeted creature must succeed on a DC 16 Wisdom saving\
    \ throw or be [charmed](/3-Mechanics/CLI/rules/conditions.md#charmed) by the beholder\
    \ for 1 hour, or until the beholder harms the creature.  \n- 2. Paralyzing Ray.\
    \ The targeted creature must succeed on a DC 16 Constitution saving throw or be\
    \ [paralyzed](/3-Mechanics/CLI/rules/conditions.md#paralyzed) for 1 minute. The\
    \ target can repeat the saving throw at the end of each of its turns, ending the\
    \ effect on itself on a success.  \n- 3. Fear Ray. The targeted creature must\
    \ succeed on a DC 16 Wisdom saving throw or be [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened)\
    \ for 1 minute. The target can repeat the saving throw at the end of each of its\
    \ turns, ending the effect on itself on a success.  \n- 4. Slowing Ray. The\
    \ targeted creature must succeed on a DC 16 Dexterity saving throw. On a failed\
    \ save, the target's speed is halved for 1 minute. In addition, the creature can't\
    \ take reactions, and it can take either an action or a bonus action on its turn,\
    \ not both. The creature can repeat the saving throw at the end of each of its\
    \ turns, ending the effect on itself on a success.  \n- 5. Enervation Ray.\
    \ The targeted creature must make a DC 16 Constitution saving throw, taking dice:8d8|text(36)\
    \ (8d8) necrotic damage on a failed save, or half as much damage on a successful\
    \ one.  \n- 6. Telekinetic Ray. If the target is a creature, it must succeed\
    \ on a DC 16 Strength saving throw or the beholder moves it up to 30 feet in any\
    \ direction. It is [restrained](/3-Mechanics/CLI/rules/conditions.md#restrained)\
    \ by the ray's telekinetic grip until the start of the beholder's next turn or\
    \ until the beholder is [incapacitated](/3-Mechanics/CLI/rules/conditions.md#incapacitated).\
    \  \n\n    If the target is an object weighing 300 pounds or less that isn't being\
    \ worn or carried, it is moved up to 30 feet in any direction. The beholder can\
    \ also exert fine control on objects with this ray, such as manipulating a simple\
    \ tool or opening a door or a container.  \n- 7. Sleep Ray. The targeted creature\
    \ must succeed on a DC 16 Wisdom saving throw or fall asleep and remain [unconscious](/3-Mechanics/CLI/rules/conditions.md#unconscious)\
    \ for 1 minute. The target awakens if it takes damage or another creature takes\
    \ an action to wake it. This ray has no effect on constructs and undead.  \n-\
    \ 8. Petrification Ray. The targeted creature must make a DC 16 Dexterity\
    \ saving throw. On a failed save, the creature begins to turn to stone and is\
    \ [restrained](/3-Mechanics/CLI/rules/conditions.md#restrained). It must repeat\
    \ the saving throw at the end of its next turn. On a success, the effect ends.\
    \ On a failure, the creature is [petrified](/3-Mechanics/CLI/rules/conditions.md#petrified)\
    \ until freed by the  [greater restoration](/3-Mechanics/CLI/spells/greater-restoration.md)\
    \ spell or other magic.  \n- 9. Disintegration Ray. If the target is a creature,\
    \ it must succeed on a DC 16 Dexterity saving throw or take dice:10d8|text(45)\
    \ (10d8) force damage. If this damage reduces the creature to 0 hit points,\
    \ its body becomes a pile of fine gray dust.  \n\n    If the target is a Large\
    \ or smaller nonmagical object or creation of magical force, it is disintegrated\
    \ without a saving throw. If the target is a Huge or larger object or creation\
    \ of magical force, this ray disintegrates a 10-foot cube of it.  \n- 10. Death\
    \ Ray. The targeted creature must succeed on a DC 16 Dexterity saving throw\
    \ or take dice:10d10|text(55) (10d10) necrotic damage. The target dies if\
    \ the ray reduces it to 0 hit points.  "
  "name": "Eye Rays"
"legendary_actions":
- "desc": "The beholder can take 3 legendary actions, using the Eye Ray option below.\
    \ It can take only one legendary action at a time and only at the end of another\
    \ creature's turn. The beholder regains spent legendary actions at the start of\
    \ its turn."
  "name": ""
- "desc": "The beholder uses one random eye ray."
  "name": "Eye Ray"
"lair_actions":
- "desc": "When fighting inside its lair, a beholder can invoke the ambient magic\
    \ to take lair actions. On initiative count 20 (losing initiative ties), the beholder\
    \ can take one lair action to cause one of the following effects:"
  "name": ""
- "desc": "- A 50-foot-square area of ground within 120 feet of the beholder becomes\
    \ slimy; that area is difficult terrain until initiative count 20 on the next\
    \ round.  \n- Walls within 120 feet of the beholder sprout grasping appendages\
    \ until initiative count 20 on the round after next. Each creature of the beholder's\
    \ choice that starts its turn within 10 feet of such a wall must succeed on a\
    \ DC 15 Dexterity saving throw or be [grappled](/3-Mechanics/CLI/rules/conditions.md#grappled).\
    \ Escaping requires a successful DC 15 Strength ([Athletics](/3-Mechanics/CLI/rules/skills.md#Athletics))\
    \ or Dexterity ([Acrobatics](/3-Mechanics/CLI/rules/skills.md#Acrobatics)) check.\
    \  \n- An eye opens on a solid surface within 60 feet of the beholder. One random\
    \ eye ray of the beholder shoots from that eye at a target of the beholder's choice\
    \ that it can see. The eye then closes and disappears.  "
  "name": ""
- "desc": "The beholder can't repeat an effect until they have all been used, and\
    \ it can't use the same effect two rounds in a row."
  "name": ""
"regional_effects":
- "desc": "A region containing a beholder's lair is warped by the creature's unnatural\
    \ presence, which creates one or more of the following effects:"
  "name": ""
- "desc": "- Creatures within 1 mile of the beholder's lair sometimes feel as if they're\
    \ being watched when they aren't.  \n- When the beholder sleeps, minor warps in\
    \ reality occur within 1 mile of its lair and then vanish 24 hours later. Marks\
    \ on cave walls might change subtly, an eerie trinket might appear where none\
    \ existed before, harmless slime might coat a statue, and so on. These effects\
    \ apply only to natural surfaces and to nonmagical objects that aren't on anyone's\
    \ person.  "
  "name": ""
- "desc": "If the beholder dies, these effects fade over the course of dice: 1d10|avg|noform\
    \ (1d10) days."
  "name": ""
"source":
- "MM"
- "ToA"
- "WDMM"
- "GoS"
- "ERLW"
- "EGW"
- "TCE"
- "CM"
- "JttRC"
- "LoX"
- "KftGV"
- "PaBTSO"
- "SatO"
- "ToFW"
- "BMT"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MM/Beholder.webp"
```
^statblock

```encounter-table
name: Beholder
creatures:
 - 1: Beholder
```

## Environment

underdark