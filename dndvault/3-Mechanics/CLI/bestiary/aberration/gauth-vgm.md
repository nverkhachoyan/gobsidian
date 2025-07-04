---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/vgm
- monster/cr/6
- monster/environment/underdark
- monster/size/medium
- monster/type/aberration
aliases: ["Gauth"]
NoteIcon: monster
BestiaryType: aberration
SourceType: Bestiary
BookSource: Volo's Guide to Monsters p. 125
---
# [Gauth](3-Mechanics\CLI\bestiary\aberration/gauth-vgm.md)
*Source: Volo's Guide to Monsters p. 125*  

A gauth is a hungry, tyrannical beholder-like creature that eats magic and tries to exact tribute from anything weaker than itself. Its body is about 4 feet in diameter, with six eyestalks, a central eye (sometimes surrounded by multiple smaller eyes), and four small grasping tentacles near its mouth. It has color and texture variations similar to a true beholder.

## Magical Metabolism

A gauth can survive on meat but prefers to sustain itself with power drained from magic objects. If starved of magic for several weeks, it is forced back to its home plane, so it constantly seeks new items to drain. A gauth might employ creatures to serve it by bringing it items that provide it with sustenance.

## Accidental Summoning

When the ritual to summon a spectator goes wrong, a gauth might push itself through the flawed connection, arriving immediately or several minutes later. It might present itself as a beholder to ignorant creatures in an attempt to intimidate them, or as a spectator to its summoner in order to drain magic items it is expected to guard.

## Inferior Tyrant

A beholder usually drives away or kills any gauths that enter its territory, but it might choose to enslave them and use them as lieutenants. Gauths are less xenophobic than beholders, so they might form small clusters and work together, though they're just as likely to ignore each other entirely.

```statblock
"name": "Gauth (VGM)"
"size": "Medium"
"type": "aberration"
"alignment": "Lawful Evil"
"ac": !!int "15"
"ac_class": "natural armor"
"hp": !!int "67"
"hit_dice": "9d8 + 27"
"stats":
- !!int "10"
- !!int "14"
- !!int "16"
- !!int "15"
- !!int "15"
- !!int "13"
"speed": "0 ft., fly 20 ft. (hover)"
"saves":
  "Charisma": !!int "4"
  "Wisdom": !!int "5"
  "Intelligence": !!int "5"
"skillsaves":
  "Perception": !!int "5"
"condition_immunities": "[prone](/3-Mechanics/CLI/rules/conditions.md#prone)"
"senses": "darkvision 120 ft., passive Perception 15"
"languages": "Deep Speech, Undercommon"
"cr": "6"
"traits":
- "desc": "When a creature that can see the gauth's central eye starts its turn within\
    \ 30 feet of the gauth, the gauth can force it to make a DC 14 Wisdom saving throw\
    \ if the gauth isn't [incapacitated](/3-Mechanics/CLI/rules/conditions.md#incapacitated)\
    \ and can see the creature. A creature that fails the save is [stunned](/3-Mechanics/CLI/rules/conditions.md#stunned)\
    \ until the start of its next turn, when it can avert its eyes again. If the creature\
    \ looks at the gauth in the meantime, it must immediately make the save."
  "name": "Stunning Gaze"
- "desc": "When the gauth dies, the magical energy within it explodes, and each creature\
    \ within 10 feet of it must make a DC 14 Dexterity saving throw, taking dice:3d8|text(13)\
    \ (3d8) force damage on a failed save, or half as much damage on a successful\
    \ one."
  "name": "Death Throes"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+8 (+8) to hit, reach 5 ft., one target.\
    \ Hit: dice:2d8|text(9) (2d8) piercing damage."
  "name": "Bite"
- "desc": "The gauth shoots three of the following magical eye rays at random (reroll\
    \ duplicates), choosing one to three targets it can see within 120 feet of it:\n\
    \n- 1. Devour Magic Ray. The targeted creature must succeed on a DC 14 Dexterity\
    \ saving throw or have one of its magic items lose all magical properties until\
    \ the start of the gauth's next turn. If the object is a charged item, it also\
    \ loses dice: 1d4|avg|noform (1d4) charges. Determine the affected item randomly,\
    \ ignoring single-use items such as potions and scrolls.  \n- 2. Enervation\
    \ Ray. The targeted creature must make a DC 14 Constitution saving throw, taking\
    \ dice:4d8|text(18) (4d8) necrotic damage on a failed save, or half as much\
    \ damage on a successful one.  \n- 3. Pushing Ray. The targeted creature must\
    \ succeed on a DC 14 Strength saving throw or be pushed up to 15 feet directly\
    \ away from the gauth and have its speed halved until the start of the gauth's\
    \ next turn.  \n- 4. Fire Ray. The targeted creature must succeed on a DC\
    \ 14 Dexterity saving throw or take dice:4d10|text(22) (4d10) fire damage.\
    \  \n- 5. Paralyzing Ray. The targeted creature must succeed on a DC 14 Constitution\
    \ saving throw or be [paralyzed](/3-Mechanics/CLI/rules/conditions.md#paralyzed)\
    \ for 1 minute. The target can repeat the saving throw at the end of each of its\
    \ turns, ending the effect on itself on a success.  \n- 6. Sleep Ray. The\
    \ targeted creature must succeed on a DC 14 Wisdom saving throw or fall asleep\
    \ and remain [unconscious](/3-Mechanics/CLI/rules/conditions.md#unconscious) for\
    \ 1 minute. The target awakens if it takes damage or another creature takes an\
    \ action to wake it. This ray has no effect on constructs and undead.  "
  "name": "Eye Rays"
"source":
- "VGM"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/VGM/Gauth.webp"
```
^statblock

```encounter-table
name: Gauth
creatures:
 - 1: Gauth
```

## Environment

underdark