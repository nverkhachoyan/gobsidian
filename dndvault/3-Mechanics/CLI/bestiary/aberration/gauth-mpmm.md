---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/6
- monster/environment/underdark
- monster/size/medium
- monster/type/aberration/beholder
aliases: ["Gauth"]
NoteIcon: monster
BestiaryType: aberration (beholder)
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 133, Volo's Guide to Monsters p. 125
---
# [Gauth](3-Mechanics\CLI\bestiary\aberration/gauth-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 133, Volo's Guide to Monsters p. 125*  

```statblock
"name": "Gauth (MPMM)"
"size": "Medium"
"type": "aberration"
"subtype": "beholder"
"alignment": "Typically  Lawful Evil"
"ac": !!int "15"
"ac_class": "natural armor"
"hp": !!int "52"
"hit_dice": "7d8 + 21"
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
    \ until the start of its next turn.\n\nUnless [surprised](/3-Mechanics/CLI/rules/conditions.md#surprised),\
    \ a creature can avert its eyes at the start of its turn to avoid the saving throw.\
    \ If the creature does so, it can't see the gauth until the start of its next\
    \ turn, when it can avert its eyes again. If the creature looks at the gauth in\
    \ the meantime, it must immediately make the save."
  "name": "Stunning Gaze"
- "desc": "When the gauth dies, the magical energy within it explodes, and each creature\
    \ within 10 feet of it must make a DC 14 Dexterity saving throw, taking dice:3d8|text(13)\
    \ (3d8) force damage on a failed save, or half as much damage on a successful\
    \ one."
  "name": "Death Throes"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+3 (+3) to hit, reach 5 ft., one target.\
    \ Hit: dice:2d8|text(9) (2d8) piercing damage."
  "name": "Bite"
- "desc": "The gauth shoots three of the following magical eye rays at random (roll\
    \ three dice: d6|avg|noform (d6)s, and reroll duplicates), targeting one to\
    \ three creatures it can see within 120 feet of it:\n\n- 1 Devour Magic Ray.\
    \ The target must succeed on a DC 14 Dexterity saving throw or have one of its\
    \ magic items lose all magical properties until the start of the gauth's next\
    \ turn. If the object is a charged item, it also loses dice: 1d4|avg|noform\
    \ (1d4) charges. Determine the affected item randomly, ignoring single-use items\
    \ such as potions and scrolls.  \n- 2 Enervation Ray. The target must make\
    \ a DC 14 Constitution saving throw, taking dice:4d8|text(18) (4d8) necrotic\
    \ damage on a failed save, or half as much damage on a successful one.  \n- 3\
    \ Fire Ray. The target must succeed on a DC 14 Dexterity saving throw or take\
    \ dice:4d10|text(22) (4d10) fire damage.  \n- 4 Paralyzing Ray. The target\
    \ must succeed on a DC 14 Constitution saving throw or be [paralyzed](/3-Mechanics/CLI/rules/conditions.md#paralyzed)\
    \ for 1 minute. The target can repeat the saving throw at the end of each of its\
    \ turns, ending the effect on itself on a success.  \n- 5 Pushing Ray. The\
    \ target must succeed on a DC 14 Strength saving throw or be pushed up to 15 feet\
    \ away from the gauth and have its speed halved until the start of the gauth's\
    \ next turn.  \n- 6 Sleep Ray. The target must succeed on a DC 14 Wisdom saving\
    \ throw or fall asleep and remain [unconscious](/3-Mechanics/CLI/rules/conditions.md#unconscious)\
    \ for 1 minute. The target awakens if it takes damage or another creature takes\
    \ an action to wake it. This ray has no effect on Constructs and Undead.  "
  "name": "Eye Rays"
"source":
- "MPMM"
- "VGM"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Gauth.webp"
```
^statblock

```encounter-table
name: Gauth
creatures:
 - 1: Gauth
```

## Environment

underdark