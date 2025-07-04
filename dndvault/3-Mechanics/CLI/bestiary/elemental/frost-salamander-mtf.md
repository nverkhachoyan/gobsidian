---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mtf
- monster/cr/9
- monster/environment/arctic
- monster/size/huge
- monster/type/elemental
aliases: ["Frost Salamander"]
NoteIcon: monster
BestiaryType: elemental
SourceType: Bestiary
BookSource: Mordenkainen's Tome of Foes p. 223
---
# [Frost Salamander](3-Mechanics\CLI\bestiary\elemental/frost-salamander-mtf.md)
*Source: Mordenkainen's Tome of Foes p. 223*  

## Salamander, Frost

Frost salamanders are natives of the Plane of Ice, also called the Frostfell, which rests between the Plane of Air and the Plane of Water. Frost salamanders especially like to hunt warm-blooded creatures. They sometimes travel to frigid climes on the Material Plane by wandering through planar gates.

## Devourers of Heat

The frost salamanders' aggression appetite for any heat source leads them to attack settlements they come across. They might mistake the fire of a forge or a campfire for a large, tasty meal, drawing them to attack expeditions and settlements that other predators would avoid. Azers sometimes venture into the Frostfell, where they use large fires to lure frost salamanders into traps to kill them and collect their hides and fangs for use in crafting weapons and armor.

## False Refuge

Although frost salamanders can burrow their way through loose soil, they prefer to dig into the ice. They roll around in piles of broken chunks of ice, allowing it to scratch their backs as they grind it down. This habit leads them to create extensive networks of ice caves, becoming ever larger as they claw fresh chunks of ice from the walls of their lairs.

A frost salamander that dwells in a lair for a while carves out enough space to allow a small army to camp within. Inexperienced travelers who come across these caves see them as a welcome shelter, though they are anything but. Frost salamanders greedily devour any prey foolhardy enough to try sleeping in their lairs.

On rare occasions, [frost giants](/3-Mechanics/CLI/bestiary/giant/frost-giant.md) capture and tame these creatures, using them to burrow into the ice to help create outposts and fortresses.

```statblock
"name": "Frost Salamander (MTF)"
"size": "Huge"
"type": "elemental"
"alignment": "Unaligned"
"ac": !!int "17"
"ac_class": "natural armor"
"hp": !!int "168"
"hit_dice": "16d12 + 64"
"stats":
- !!int "20"
- !!int "12"
- !!int "18"
- !!int "7"
- !!int "11"
- !!int "7"
"speed": "60 ft., burrow 40 ft., climb 40 ft."
"saves":
  "Wisdom": !!int "4"
  "Constitution": !!int "8"
"skillsaves":
  "Perception": !!int "4"
"damage_vulnerabilities": "fire"
"damage_immunities": "cold"
"senses": "darkvision 60 ft., tremorsense 60 ft., passive Perception 14"
"languages": "Primordial"
"cr": "9"
"traits":
- "desc": "When the salamander takes fire damage, its Freezing Breath automatically\
    \ recharges."
  "name": "Burning Fury"
"actions":
- "desc": "The salamander makes five attacks: four with its claws and one with its\
    \ bite."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+9 (+9) to hit, reach 10 ft., one target.\
    \ Hit: dice:1d6 + 5|text(8) (1d6 + 5) piercing damage."
  "name": "Claws"
- "desc": "Melee Weapon Attack: dice: d20+9 (+9) to hit, reach 15 ft., one target.\
    \ Hit: dice:1d8 + 5|text(9) (1d8 + 5) piercing damage and dice:1d10|text(5)\
    \ (1d10) cold damage."
  "name": "Bite"
- "desc": "The salamander exhales chill wind in a 60-foot cone. Each creature in that\
    \ area must make a DC 17 Constitution saving throw, taking dice:8d10|text(44)\
    \ (8d10) cold damage on a failed save, or half as much damage on a successful\
    \ one."
  "name": "Freezing Breath (Recharge 6)"
"source":
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MTF/Frost%20Salamander.webp"
```
^statblock

```encounter-table
name: Frost Salamander
creatures:
 - 1: Frost Salamander
```

## Environment

arctic