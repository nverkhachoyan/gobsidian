---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/vgm
- monster/cr/2
- monster/environment/desert
- monster/environment/urban
- monster/size/medium
- monster/type/dragon
aliases: ["Blue Guard Drake"]
NoteIcon: monster
BestiaryType: dragon
SourceType: Bestiary
BookSource: Volo's Guide to Monsters p. 158
---
# [Blue Guard Drake](3-Mechanics\CLI\bestiary\dragon/blue-guard-drake-vgm.md)
*Source: Volo's Guide to Monsters p. 158*  

A guard drake is a reptilian creature created out of dragon scales by means of a bizarre and grisly ritual. When trained properly, a drake is obedient, loyal, and territorial, which makes it an excellent watch-beast that can follow simple commands.

## Gifts from Dragons

The ritual to create a guard drake was originally devised by the cult of Tiamat, but has spread to other groups that are skilled in arcana and associated with dragons. The cooperation of a dragon is necessary for the ritual to succeed, and a dragon typically provides its help when it wants to reward its allies or worshipers with a valuable servant.

The ritual, which takes several days, requires 10 pounds of fresh dragon scales (donated by the dragon allied with the group), a large amount of fresh meat, and an iron cauldron. When the process is complete, a halfling-sized egg emerges from the cauldron and is ready to hatch within a few hours.

A guard drake resembles the type of dragon it was created from, but with a wingless, squat, muscular build. A drake can't reproduce, nor can its scales be used to make other guard drakes.

## Eager to Learn

A newly hatched guard drake imprints upon the first creature that feeds it (usually the one planning to train it), establishing an aggressive but trusting bond with that individual. A guard drake is fully grown within two to three weeks and can be trained in the same length of time. One is the equivalent of a guard dog in terms of what it can be trained to do.

## Variant: Chromatic Guard Drakes

Each type of chromatic dragon's scales and blood creates a guard drake that resembles a wingless, stunted version of that type of dragon, with unique abilities related to that type. Each has the special features described below.

### Black Guard Drake

A black guard drake is amphibious (it can breathe air or water), has a swimming speed of 30 feet, and has resistance to acid damage.

### Blue Guard Drake

A blue guard drake has a burrowing speed of 20 feet and resistance to lightning damage.

### Green Guard Drake

A green guard drake is amphibious (it can breathe air or water), has a swimming speed of 30 feet, and has resistance to poison damage.

### Red Guard Drake

A red guard drake has climbing speed of 30 feet and resistance to fire damage.

### White Guard Drake

A white guard drake has a burrowing speed of 20 feet, a climbing speed of 30 feet, and resistance to cold damage.

```statblock
"name": "Blue Guard Drake (VGM)"
"size": "Medium"
"type": "dragon"
"alignment": "Unaligned"
"ac": !!int "14"
"ac_class": "natural armor"
"hp": !!int "52"
"hit_dice": "7d8 + 21"
"stats":
- !!int "16"
- !!int "11"
- !!int "16"
- !!int "4"
- !!int "10"
- !!int "7"
"speed": "30 ft., burrow 20 ft."
"skillsaves":
  "Perception": !!int "2"
"damage_resistances": "lightning"
"senses": "darkvision 60 ft., passive Perception 12"
"languages": "understands Draconic but can't speak it"
"cr": "2"
"actions":
- "desc": "The drake attacks twice, once with its bite and once with its tail."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+5 (+5) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d8 + 3|text(7) (1d8 + 3) piercing damage."
  "name": "Bite"
- "desc": "Melee Weapon Attack: dice: d20+5 (+5) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d6 + 3|text(6) (1d6 + 3) bludgeoning damage."
  "name": "Tail"
"source":
- "VGM"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/VGM/Blue%20Guard%20Drake.webp"
```
^statblock

```encounter-table
name: Blue Guard Drake
creatures:
 - 1: Blue Guard Drake
```

## Environment

desert, urban