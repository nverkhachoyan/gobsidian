---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mm
- monster/cr/2
- monster/environment/forest
- monster/environment/underdark
- monster/size/medium
- monster/type/monstrosity
aliases: ["Grick"]
NoteIcon: monster
BestiaryType: monstrosity
SourceType: Bestiary
BookSource: Monster Manual p. 173. Available in the SRD and the Basic Rules.
---
# [Grick](3-Mechanics\CLI\bestiary\monstrosity/grick.md)
*Source: Monster Manual p. 173. Available in the SRD and the Basic Rules.*  

The wormlike grick waits unseen, blending in with the rock of the caves and caverns it haunts. Only when prey comes near does it rear up, its four barbed tentacles unfurling to reveal its hungry, snapping beak.

## Passive Predators

Gricks rarely hunt. Instead, they drag their rubbery bodies to places where creatures regularly pass, lurking out of sight amid rocky rubble and debris, squeezing into burrows, holes, or crevices, climbing up to ledges, or coiling around stalactites to drop on unwary prey. A grick consumes virtually anything that moves except for other gricks. It targets the nearest prey, grabbing a fallen creature with its tentacles and dragging it off to eat alone.

## Roving Ambushers

Gricks remain in an area until the food supply dwindles, often because sentient creatures become aware of their presence and plot alternate routes around their lairs. When prey is scarce in the Underdark, gricks venture aboveground to hunt in the wilderness, lurking in trees or on cliff-side ledges. A grick pack is often led by a single well-fed, oversized alpha around which the others congregate.

## Spoils of Slaughter

Over time, grick lairs accumulate the cast-off possessions of intelligent prey, and expert guides know to look out for these tell-tale signs. Underdark explorers sometimes seal off the routes leading to and from a grick lair to starve them, then claim the wealth of the foul creatures' victims.

```statblock
"name": "Grick"
"size": "Medium"
"type": "monstrosity"
"alignment": "Neutral"
"ac": !!int "14"
"ac_class": "natural armor"
"hp": !!int "27"
"hit_dice": "6d8"
"stats":
- !!int "14"
- !!int "14"
- !!int "11"
- !!int "3"
- !!int "14"
- !!int "5"
"speed": "30 ft., climb 30 ft."
"damage_resistances": "bludgeoning, piercing, slashing from nonmagical attacks"
"senses": "darkvision 60 ft., passive Perception 12"
"languages": ""
"cr": "2"
"traits":
- "desc": "The grick has advantage on Dexterity ([Stealth](/3-Mechanics/CLI/rules/skills.md#Stealth))\
    \ checks made to hide in rocky terrain."
  "name": "Stone Camouflage"
"actions":
- "desc": "The grick makes one attack with its tentacles. If that attack hits, the\
    \ grick can make one beak attack against the same target."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+4 (+4) to hit, reach 5 ft., one target.\
    \ Hit: dice:2d6 + 2|text(9) (2d6 + 2) slashing damage."
  "name": "Tentacles"
- "desc": "Melee Weapon Attack: dice: d20+4 (+4) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d6 + 2|text(5) (1d6 + 2) piercing damage."
  "name": "Beak"
"source":
- "MM"
- "CoS"
- "LMoP"
- "PotA"
- "TftYP"
- "WDMM"
- "CM"
- "DSotDQ"
- "KftGV"
- "PaBTSO"
- "QftIS"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MM/Grick.webp"
```
^statblock

```encounter-table
name: Grick
creatures:
 - 1: Grick
```

## Environment

underdark, forest