---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/6
- monster/environment/hill
- monster/environment/mountain
- monster/size/large
- monster/type/fey
aliases: ["Annis Hag"]
NoteIcon: monster
BestiaryType: fey
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 47, Volo's Guide to Monsters p. 159
---
# [Annis Hag](3-Mechanics\CLI\bestiary\fey/annis-hag-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 47, Volo's Guide to Monsters p. 159*  

```statblock
"name": "Annis Hag (MPMM)"
"size": "Large"
"type": "fey"
"alignment": "Typically  Chaotic Evil"
"ac": !!int "17"
"ac_class": "natural armor"
"hp": !!int "90"
"hit_dice": "12d10 + 24"
"stats":
- !!int "21"
- !!int "12"
- !!int "14"
- !!int "13"
- !!int "14"
- !!int "15"
"speed": "40 ft."
"saves":
  "Constitution": !!int "5"
"skillsaves":
  "Deception": !!int "5"
  "Perception": !!int "5"
"damage_resistances": "cold"
"senses": "darkvision 60 ft., passive Perception 15"
"languages": "Common, Giant, Sylvan"
"cr": "6"
"traits":
- "desc": "The hag casts one of the following spells, using Charisma as the spellcasting\
    \ ability (spell save DC 13):\n\n3/day each: [disguise self](/3-Mechanics/CLI/spells/disguise-self.md)\
    \ (including the form of a Medium Humanoid), [Fog cloud](/3-Mechanics/CLI/spells/fog-cloud.md)"
  "name": "Spellcasting"
"actions":
- "desc": "The annis makes one Bite attack and two Claw attacks."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+8 (+8) to hit, reach 5 ft., one target.\
    \ Hit: dice:3d6 + 5|text(15) (3d6 + 5) piercing damage."
  "name": "Bite"
- "desc": "Melee Weapon Attack: dice: d20+8 (+8) to hit, reach 5 ft., one target.\
    \ Hit: dice:3d6 + 5|text(15) (3d6 + 5) slashing damage."
  "name": "Claw"
- "desc": "Melee Weapon Attack: dice: d20+8 (+8) to hit, reach 5 ft., one target.\
    \ Hit: dice:9d6 + 5|text(36) (9d6 + 5) bludgeoning damage, and the target\
    \ is [grappled](/3-Mechanics/CLI/rules/conditions.md#grappled) (escape DC 15)\
    \ if it is a Large or smaller creature. Until the grapple ends, the target takes\
    \ dice:9d6 + 5|text(36) (9d6 + 5) bludgeoning damage at the start of each\
    \ of the hag's turns. The hag can't make attacks while grappling a creature in\
    \ this way."
  "name": "Crushing Hug"
"source":
- "MPMM"
- "VGM"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Annis%20Hag.webp"
```
^statblock

```encounter-table
name: Annis Hag
creatures:
 - 1: Annis Hag
```

## Environment

hill, mountain