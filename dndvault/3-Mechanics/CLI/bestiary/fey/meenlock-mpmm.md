---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/2
- monster/environment/forest
- monster/environment/swamp
- monster/environment/urban
- monster/size/small
- monster/type/fey
aliases: ["Meenlock"]
NoteIcon: monster
BestiaryType: fey
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 178, Volo's Guide to Monsters p. 170
---
# [Meenlock](3-Mechanics\CLI\bestiary\fey/meenlock-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 178, Volo's Guide to Monsters p. 170*  

```statblock
"name": "Meenlock (MPMM)"
"size": "Small"
"type": "fey"
"alignment": "Typically  Neutral Evil"
"ac": !!int "15"
"ac_class": "natural armor"
"hp": !!int "31"
"hit_dice": "7d6 + 7"
"stats":
- !!int "7"
- !!int "15"
- !!int "12"
- !!int "11"
- !!int "10"
- !!int "8"
"speed": "30 ft."
"skillsaves":
  "Stealth": !!int "6"
  "Perception": !!int "4"
  "Survival": !!int "2"
"condition_immunities": "[frightened](/3-Mechanics/CLI/rules/conditions.md#frightened)"
"senses": "darkvision 120 ft., passive Perception 14"
"languages": "telepathy 120 ft."
"cr": "2"
"traits":
- "desc": "Any Beast or Humanoid that starts its turn within 10 feet of the meenlock\
    \ must succeed on a DC 11 Wisdom saving throw or be [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened)\
    \ until the start of the creature's next turn."
  "name": "Fear Aura"
- "desc": "While in bright light, the meenlock has disadvantage on attack rolls, as\
    \ well as on Wisdom ([Perception](/3-Mechanics/CLI/rules/skills.md#Perception))\
    \ checks that rely on sight."
  "name": "Light Sensitivity"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+4 (+4) to hit, reach 5 ft., one target.\
    \ Hit: dice:2d4 + 2|text(7) (2d4 + 2) slashing damage, and the target must\
    \ succeed on a DC 11 Constitution saving throw or be [paralyzed](/3-Mechanics/CLI/rules/conditions.md#paralyzed)\
    \ for 1 minute. The target can repeat the saving throw at the end of each of its\
    \ turns, ending the effect on itself on a success."
  "name": "Claw"
"bonus_actions":
- "desc": "The meenlock teleports to an unoccupied space within 30 feet of it, provided\
    \ that both the space it's teleporting from and its destination are in dim light\
    \ or darkness. The destination need not be within line of sight."
  "name": "Shadow Teleport (Recharge 5-6)"
"source":
- "MPMM"
- "VGM"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Meenlock.webp"
```
^statblock

```encounter-table
name: Meenlock
creatures:
 - 1: Meenlock
```

## Environment

forest, swamp, urban