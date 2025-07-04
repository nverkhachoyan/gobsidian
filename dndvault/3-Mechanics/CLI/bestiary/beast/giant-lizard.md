---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mm
- monster/cr/1-4
- monster/environment/coastal
- monster/environment/desert
- monster/environment/forest
- monster/environment/swamp
- monster/environment/underdark
- monster/size/large
- monster/type/beast
aliases: ["Giant Lizard"]
NoteIcon: monster
BestiaryType: beast
SourceType: Bestiary
BookSource: Monster Manual p. 326. Available in the SRD and the Basic Rules.
---
# [Giant Lizard](3-Mechanics\CLI\bestiary\beast/giant-lizard.md)
*Source: Monster Manual p. 326. Available in the SRD and the Basic Rules.*  

A giant lizard can be ridden or used as a draft animal. Lizardfolk also keep them as pets, and subterranean giant lizards are used as mounts and pack animals by drow, duergar, and other Underdark dwellers.

> [!note] Variant: Giant Lizard Traits
> 
> Some giant lizards have one or both of the following traits.
> 
> **Hold Breath.** The lizard can hold its breath for 15 minutes. (A lizard that has this trait also has a swimming speed of 30 feet.)
> 
> **Spider Climb.** The lizard can climb difficult surfaces, including upside down on ceilings, without needing to make an ability check.
^variant-giant-lizard-traits

```statblock
"name": "Giant Lizard"
"size": "Large"
"type": "beast"
"alignment": "Unaligned"
"ac": !!int "12"
"ac_class": "natural armor"
"hp": !!int "19"
"hit_dice": "3d10 + 3"
"stats":
- !!int "15"
- !!int "12"
- !!int "13"
- !!int "2"
- !!int "10"
- !!int "5"
"speed": "30 ft., climb 30 ft."
"senses": "darkvision 30 ft., passive Perception 10"
"languages": ""
"cr": "1/4"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+4 (+4) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d8 + 2|text(6) (1d8 + 2) piercing damage."
  "name": "Bite"
"source":
- "MM"
- "HotDQ"
- "TftYP"
- "ToA"
- "WDMM"
- "GoS"
- "IDRotF"
- "JttRC"
- "KftGV"
- "PSA"
- "VEoR"
- "QftIS"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MM/Giant%20Lizard.webp"
```
^statblock

```encounter-table
name: Giant Lizard
creatures:
 - 1: Giant Lizard
```

## Environment

underdark, forest, swamp, desert, coastal