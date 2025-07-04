---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/vgm
- monster/cr/2
- monster/environment/coastal
- monster/environment/hill
- monster/environment/mountain
- monster/size/huge
- monster/type/beast
aliases: ["Quetzalcoatlus"]
NoteIcon: monster
BestiaryType: beast
SourceType: Bestiary
BookSource: Volo's Guide to Monsters p. 140
---
# [Quetzalcoatlus](3-Mechanics\CLI\bestiary\beast/quetzalcoatlus-vgm.md)
*Source: Volo's Guide to Monsters p. 140*  

This giant relative of the pteranodon has a wingspan exceeding 30 feet. Although it can move on the ground like a quadruped, it is more comfortable in the air.

```statblock
"name": "Quetzalcoatlus (VGM)"
"size": "Huge"
"type": "beast"
"alignment": "Unaligned"
"ac": !!int "13"
"ac_class": "natural armor"
"hp": !!int "30"
"hit_dice": "4d12 + 4"
"stats":
- !!int "15"
- !!int "13"
- !!int "13"
- !!int "2"
- !!int "10"
- !!int "5"
"speed": "10 ft., fly 80 ft."
"skillsaves":
  "Perception": !!int "2"
"senses": "passive Perception 12"
"languages": ""
"cr": "2"
"traits":
- "desc": "If the quetzalcoatlus is flying and dives at least 30 feet toward a target\
    \ and then hits with a bite attack, the attack deals an extra dice:3d6|text(10)\
    \ (3d6) damage to the target."
  "name": "Dive Attack"
- "desc": "The quetzalcoatlus doesn't provoke an opportunity attack when it flies\
    \ out of an enemy's reach."
  "name": "Flyby"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+4 (+4) to hit, reach 10 ft., one creature.\
    \ Hit: dice:3d6 + 2|text(12) (3d6 + 2) piercing damage."
  "name": "Bite"
"source":
- "VGM"
- "ToA"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/VGM/Quetzalcoatlus.webp"
```
^statblock

```encounter-table
name: Quetzalcoatlus
creatures:
 - 1: Quetzalcoatlus
```

## Environment

mountain, hill, coastal