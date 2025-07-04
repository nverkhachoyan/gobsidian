---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mm
- monster/cr/2
- monster/environment/forest
- monster/size/huge
- monster/type/plant
aliases: ["Awakened Tree"]
NoteIcon: monster
BestiaryType: plant
SourceType: Bestiary
BookSource: Monster Manual p. 317. Available in the SRD and the Basic Rules.
---
# [Awakened Tree](3-Mechanics\CLI\bestiary\plant/awakened-tree.md)
*Source: Monster Manual p. 317. Available in the SRD and the Basic Rules.*  

An awakened tree is an ordinary tree given sentience and mobility by the [awaken](/3-Mechanics/CLI/spells/awaken.md) spell or similar magic.

```statblock
"name": "Awakened Tree"
"size": "Huge"
"type": "plant"
"alignment": "Unaligned"
"ac": !!int "13"
"ac_class": "natural armor"
"hp": !!int "59"
"hit_dice": "7d12 + 14"
"stats":
- !!int "19"
- !!int "6"
- !!int "15"
- !!int "10"
- !!int "10"
- !!int "7"
"speed": "20 ft."
"damage_vulnerabilities": "fire"
"damage_resistances": "bludgeoning, piercing"
"senses": "passive Perception 10"
"languages": "one language known by its creator"
"cr": "2"
"traits":
- "desc": "While the tree remains motionless, it is indistinguishable from a normal\
    \ tree."
  "name": "False Appearance"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+6 (+6) to hit, reach 10 ft., one target.\
    \ Hit: dice:3d6 + 4|text(14) (3d6 + 4) bludgeoning damage."
  "name": "Slam"
"source":
- "MM"
- "PotA"
- "RoT"
- "SKT"
- "WDH"
- "WDMM"
- "GoS"
- "MOT"
- "IDRotF"
- "CM"
- "WBtW"
- "PSI"
- "QftIS"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MM/Awakened%20Tree.webp"
```
^statblock

```encounter-table
name: Awakened Tree
creatures:
 - 1: Awakened Tree
```

## Environment

forest