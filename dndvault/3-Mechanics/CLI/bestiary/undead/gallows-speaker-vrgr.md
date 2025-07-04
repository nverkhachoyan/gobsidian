---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/vrgr
- monster/cr/6
- monster/size/medium
- monster/type/undead
aliases: ["Gallows Speaker"]
NoteIcon: monster
BestiaryType: undead
SourceType: Bestiary
BookSource: Van Richten's Guide to Ravenloft p. 234
---
# [Gallows Speaker](3-Mechanics\CLI\bestiary\undead/gallows-speaker-vrgr.md)
*Source: Van Richten's Guide to Ravenloft p. 234*  

Gallows speakers arise from places of mass death or sites where creatures regularly meet their doom. Over time, pain-wracked phantoms and lingering souls combine into an entity that knows death in myriad forms. Such amalgamated spirits are tormented by their collective pain, endlessly moaning disjointed final thoughts as they lash out at the living. Having known untold deaths, gallows speakers can predict suffering, foreseeing dooms leveled against them and overwhelming their foes with visions of innumerable violent deaths.

Gallows speakers rarely speak coherently or communicate with the living, instead being entirely obsessed with their memories of death. These undead endlessly mutter to themselves, giving voice to final curses, regrets, pleas, and apologies. Those who linger and listen to a gallows speaker might gain insight into any of its many deaths.

```statblock
"name": "Gallows Speaker (VRGR)"
"size": "Medium"
"type": "undead"
"alignment": "Unaligned"
"ac": !!int "12"
"hp": !!int "85"
"hit_dice": "19d8"
"stats":
- !!int "8"
- !!int "14"
- !!int "10"
- !!int "10"
- !!int "12"
- !!int "18"
"speed": "0 ft., fly 40 ft. (hover)"
"saves":
  "Wisdom": !!int "4"
"skillsaves":
  "Perception": !!int "7"
"damage_resistances": "bludgeoning, piercing, slashing from nonmagical attacks"
"damage_immunities": "necrotic, poison"
"condition_immunities": "[charmed](/3-Mechanics/CLI/rules/conditions.md#charmed),\
  \ [exhaustion](/3-Mechanics/CLI/rules/conditions.md#exhaustion), [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened),\
  \ [grappled](/3-Mechanics/CLI/rules/conditions.md#grappled), [paralyzed](/3-Mechanics/CLI/rules/conditions.md#paralyzed),\
  \ [petrified](/3-Mechanics/CLI/rules/conditions.md#petrified), [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned),\
  \ [prone](/3-Mechanics/CLI/rules/conditions.md#prone), [restrained](/3-Mechanics/CLI/rules/conditions.md#restrained)"
"senses": "truesight 60 ft., passive Perception 17"
"languages": "any languages its component spirits knew in life"
"cr": "6"
"traits":
- "desc": "The gallows speaker can see 60 feet into the Ethereal Plane when it is\
    \ on the Material Plane and vice versa."
  "name": "Divination Senses"
- "desc": "The gallows speaker can move through other creatures and objects as if\
    \ they were difficult terrain. It takes dice:1d10|text(5) (1d10) force damage\
    \ if it ends it turn inside an object."
  "name": "Incorporeal Movement"
- "desc": "The gallows speaker doesn't require air, food, drink, or sleep."
  "name": "Unusual Nature"
"actions":
- "desc": "Melee Spell Attack: dice: d20+7 (+7) to hit, reach 5 ft., one creature.\
    \ Hit: dice:2d10 + 4|text(15) (2d10 + 4) psychic damage, and the target\
    \ must roll a dice: d4|avg|noform (d4) and subtract the number rolled from\
    \ the next attack roll or saving throw it makes before the start of the gallows\
    \ speaker's next turn."
  "name": "Foretelling Touch"
- "desc": "The gallows speaker targets a creature it can see within 30 feet of it.\
    \ The target must make a DC 15 Wisdom saving throw. On a failed save, the target\
    \ takes dice:3d12|text(19) (3d12) psychic damage, and waves of painful memories\
    \ leap from the target to up to three other creatures of the gallows speaker's\
    \ choice that are within 30 feet of the target, each of which takes dice:3d8|text(13)\
    \ (3d8) psychic damage."
  "name": "Suffering Echoes"
"source":
- "VRGR"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/VRGR/Gallows%20Speaker.webp"
```
^statblock

```encounter-table
name: Gallows Speaker
creatures:
 - 1: Gallows Speaker
```