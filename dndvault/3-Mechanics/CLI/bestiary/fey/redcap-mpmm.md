---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/3
- monster/environment/forest
- monster/environment/hill
- monster/environment/swamp
- monster/size/small
- monster/type/fey
aliases: ["Redcap"]
NoteIcon: monster
BestiaryType: fey
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 208, Volo's Guide to Monsters p. 188
---
# [Redcap](3-Mechanics\CLI\bestiary\fey/redcap-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 208, Volo's Guide to Monsters p. 188*  

```statblock
"name": "Redcap (MPMM)"
"size": "Small"
"type": "fey"
"alignment": "Typically  Chaotic Evil"
"ac": !!int "14"
"ac_class": "natural armor"
"hp": !!int "45"
"hit_dice": "6d6 + 24"
"stats":
- !!int "18"
- !!int "13"
- !!int "18"
- !!int "10"
- !!int "12"
- !!int "9"
"speed": "25 ft."
"skillsaves":
  "Athletics": !!int "6"
  "Perception": !!int "3"
"senses": "darkvision 60 ft., passive Perception 13"
"languages": "Common, Sylvan"
"cr": "3"
"traits":
- "desc": "The redcap has disadvantage on Dexterity ([Stealth](/3-Mechanics/CLI/rules/skills.md#Stealth))\
    \ checks."
  "name": "Iron Boots"
- "desc": "While grappling, the redcap is considered to be Medium. Also, wielding\
    \ a heavy weapon doesn't impose disadvantage on its attack rolls."
  "name": "Outsize Strength"
"actions":
- "desc": "The redcap makes three Wicked Sickle attacks."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+6 (+6) to hit, reach 5 ft., one target.\
    \ Hit: dice:2d4 + 4|text(9) (2d4 + 4) slashing damage."
  "name": "Wicked Sickle"
- "desc": "The redcap moves up to its speed to a creature it can see and kicks with\
    \ its iron boots. The target must succeed on a DC 14 Dexterity saving throw or\
    \ take dice:3d10 + 4|text(20) (3d10 + 4) bludgeoning damage and be knocked\
    \ [prone](/3-Mechanics/CLI/rules/conditions.md#prone)."
  "name": "Ironbound Pursuit"
"source":
- "MPMM"
- "VGM"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Redcap.webp"
```
^statblock

```encounter-table
name: Redcap
creatures:
 - 1: Redcap
```

## Environment

forest, hill, swamp