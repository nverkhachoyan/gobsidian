---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mm
- monster/cr/11
- monster/environment/arctic
- monster/size/huge
- monster/type/monstrosity
aliases: ["Remorhaz"]
NoteIcon: monster
BestiaryType: monstrosity
SourceType: Bestiary
BookSource: Monster Manual p. 258. Available in the SRD.
---
# [Remorhaz](3-Mechanics\CLI\bestiary\monstrosity/remorhaz.md)
*Source: Monster Manual p. 258. Available in the SRD.*  

From beneath the snow and ice bursts a remorhaz in a cloud of steam, its body pulsing with internal fire. Wing like fins flare from the back of the creature's head, and its wide mouth brims with jagged teeth.

## Arctic Predators

Remorhazes live in arctic climes, preying on elk, polar bears, and other creatures sharing their territory. They can't tolerate warm weather, having adapted to the cold by generating a furnace-like heat within their bodies. When hunting, a remorhaz burrows deep below the snow and ice and lies in wait for the faint vibrations created by a creature moving above it. While hidden under the ice and snow, it can lower its body temperature so that it doesn't melt its cover.

## Young Ones

Frost giant hunters scour the icy wastes for remorhaz nests and eggs. The giants prize young remorhazes, which can be trained from hatching to obey commands and guard the giants' icy citadels. Unlike fully grown specimens, young remorhazes gnaw on their victims instead of swallowing them whole.

```statblock
"name": "Remorhaz"
"size": "Huge"
"type": "monstrosity"
"alignment": "Unaligned"
"ac": !!int "17"
"ac_class": "natural armor"
"hp": !!int "195"
"hit_dice": "17d12 + 85"
"stats":
- !!int "24"
- !!int "13"
- !!int "21"
- !!int "4"
- !!int "10"
- !!int "5"
"speed": "30 ft., burrow 20 ft."
"damage_immunities": "cold, fire"
"senses": "darkvision 60 ft., tremorsense 60 ft., passive Perception 10"
"languages": ""
"cr": "11"
"traits":
- "desc": "A creature that touches the remorhaz or hits it with a melee attack while\
    \ within 5 feet of it takes dice:3d6|text(10) (3d6) fire damage."
  "name": "Heated Body"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+11 (+11) to hit, reach 10 ft., one\
    \ target. Hit: dice:6d10 + 7|text(40) (6d10 + 7) piercing damage plus dice:3d6|text(10)\
    \ (3d6) fire damage. If the target is a creature, it is [grappled](/3-Mechanics/CLI/rules/conditions.md#grappled)\
    \ (escape DC 17). Until this grapple ends, the target is [restrained](/3-Mechanics/CLI/rules/conditions.md#restrained),\
    \ and the remorhaz can't bite another target."
  "name": "Bite"
- "desc": "The remorhaz makes one bite attack against a Medium or smaller creature\
    \ it is grappling. If the attack hits, that creature takes the bite's damage and\
    \ is swallowed, and the grapple ends. While swallowed, the creature is [blinded](/3-Mechanics/CLI/rules/conditions.md#blinded)\
    \ and [restrained](/3-Mechanics/CLI/rules/conditions.md#restrained), it has total\
    \ cover against attacks and other effects outside the remorhaz, and it takes dice:6d6|text(21)\
    \ (6d6) acid damage at the start of each of the remorhaz's turns.\n\nIf the\
    \ remorhaz takes 30 damage or more on a single turn from a creature inside it,\
    \ the remorhaz must succeed on a DC 15 Constitution saving throw at the end of\
    \ that turn or regurgitate all swallowed creatures, which fall [prone](/3-Mechanics/CLI/rules/conditions.md#prone)\
    \ in a space within 10 feet of the remorhaz. If the remorhaz dies, a swallowed\
    \ creature is no longer [restrained](/3-Mechanics/CLI/rules/conditions.md#restrained)\
    \ by it and can escape from the corpse using 15 feet of movement, exiting [prone](/3-Mechanics/CLI/rules/conditions.md#prone)."
  "name": "Swallow"
"source":
- "MM"
- "SKT"
- "TftYP"
- "WDMM"
- "BGDIA"
- "IMR"
- "EGW"
- "IDRotF"
- "LoX"
- "PaBTSO"
- "BMT"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MM/Remorhaz.webp"
```
^statblock

```encounter-table
name: Remorhaz
creatures:
 - 1: Remorhaz
```

## Environment

arctic