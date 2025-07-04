---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/2
- monster/environment/coastal
- monster/environment/hill
- monster/environment/mountain
- monster/size/huge
- monster/type/beast/dinosaur
aliases: ["Quetzalcoatlus"]
NoteIcon: monster
BestiaryType: beast (dinosaur)
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 96, Volo's Guide to Monsters p. 140
---
# [Quetzalcoatlus](3-Mechanics\CLI\bestiary\beast/quetzalcoatlus-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 96, Volo's Guide to Monsters p. 140*  

```statblock
"name": "Quetzalcoatlus (MPMM)"
"size": "Huge"
"type": "beast"
"subtype": "dinosaur"
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
- "desc": "The quetzalcoatlus doesn't provoke an [opportunity attack](/3-Mechanics/CLI/rules/actions.md#opportunity%20attack)\
    \ when it flies out of an enemy's reach."
  "name": "Flyby"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+4 (+4) to hit, reach 10 ft., one creature.\
    \ Hit: dice:3d6 + 2|text(12) (3d6 + 2) piercing damage. If the quetzalcoatlus\
    \ flew least 30 feet toward the target immediately before the hit, the target\
    \ takes an extra dice:3d6|text(10) (3d6) piercing damage."
  "name": "Bite"
"source":
- "MPMM"
- "VGM"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Quetzalcoatlus.webp"
```
^statblock

```encounter-table
name: Quetzalcoatlus
creatures:
 - 1: Quetzalcoatlus
```

## Environment

coastal, hill, mountain