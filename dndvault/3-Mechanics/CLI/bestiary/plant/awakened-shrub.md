---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mm
- monster/cr/0
- monster/environment/forest
- monster/size/small
- monster/type/plant
aliases: ["Awakened Shrub"]
NoteIcon: monster
BestiaryType: plant
SourceType: Bestiary
BookSource: Monster Manual p. 317. Available in the SRD and the Basic Rules.
---
# [Awakened Shrub](3-Mechanics\CLI\bestiary\plant/awakened-shrub.md)
*Source: Monster Manual p. 317. Available in the SRD and the Basic Rules.*  

An awakened shrub is an ordinary shrub given sentience and mobility by the [awaken](/3-Mechanics/CLI/spells/awaken.md) spell or similar magic.

```statblock
"name": "Awakened Shrub"
"size": "Small"
"type": "plant"
"alignment": "Unaligned"
"ac": !!int "9"
"hp": !!int "10"
"hit_dice": "3d6"
"stats":
- !!int "3"
- !!int "8"
- !!int "11"
- !!int "10"
- !!int "10"
- !!int "6"
"speed": "20 ft."
"damage_vulnerabilities": "fire"
"damage_resistances": "piercing"
"senses": "passive Perception 10"
"languages": "one language known by its creator"
"cr": "0"
"traits":
- "desc": "While the shrub remains motionless, it is indistinguishable from a normal\
    \ shrub."
  "name": "False Appearance"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+1 (+1) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d4 - 1|text(1) (1d4 - 1) slashing damage."
  "name": "Rake"
"source":
- "MM"
- "SKT"
- "WDMM"
- "IMR"
- "MOT"
- "IDRotF"
- "WBtW"
- "PSI"
- "DoDk"
- "QftIS"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MM/Awakened%20Shrub.webp"
```
^statblock

```encounter-table
name: Awakened Shrub
creatures:
 - 1: Awakened Shrub
```

## Environment

forest