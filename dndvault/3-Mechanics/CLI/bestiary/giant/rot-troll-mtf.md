---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mtf
- monster/cr/9
- monster/environment/desert
- monster/environment/forest
- monster/environment/swamp
- monster/environment/underdark
- monster/size/large
- monster/type/giant
aliases: ["Rot Troll"]
NoteIcon: monster
BestiaryType: giant
SourceType: Bestiary
BookSource: Mordenkainen's Tome of Foes p. 244, Dragon of Icespire Peak, Storm Lord's Wrath
---
# [Rot Troll](3-Mechanics\CLI\bestiary\giant/rot-troll-mtf.md)
*Source: Mordenkainen's Tome of Foes p. 244, Dragon of Icespire Peak, Storm Lord's Wrath*  

## Rot Troll

A troll that is infused with waves of necrotic energy as it regenerates can develop a symbiotic relationship with that deathly power. The troll's body withers, and its flesh falls away from the body as quickly as it forms. Eventually a rot troll becomes unable to regenerate, though it still heals normally. The creature courses with necrotic energy that flows out of its withered form. Simply standing near a rot troll exposes other creatures to its lethal emanations.

## Trolls

Trolls that are nearly obliterated but survive and regenerate from mere scraps of flesh can display bizarre mutations. One of these warped trolls is especially likely to arise if the creature regenerates in the presence of magical emanations, planar energy, disease, or death on a vast scale, or if its body was damaged by elemental forces. These mutated forms can also be produced and shaped by the ritual magic of evil spellcasters.

```statblock
"name": "Rot Troll (MTF)"
"size": "Large"
"type": "giant"
"alignment": "Chaotic Evil"
"ac": !!int "16"
"ac_class": "natural armor"
"hp": !!int "138"
"hit_dice": "12d10 + 72"
"stats":
- !!int "18"
- !!int "13"
- !!int "22"
- !!int "5"
- !!int "8"
- !!int "4"
"speed": "30 ft."
"skillsaves":
  "Perception": !!int "3"
"damage_immunities": "necrotic"
"senses": "darkvision 60 ft., passive Perception 13"
"languages": "Giant"
"cr": "9"
"traits":
- "desc": "At the end of each of the troll's turns, each creature within 5 feet of\
    \ it takes dice:2d10|text(11) (2d10) necrotic damage, unless the troll has\
    \ taken acid or fire damage since the end of its last turn."
  "name": "Rancid Degeneration"
"actions":
- "desc": "The troll makes three attacks: one with its bite and two with its claws."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+8 (+8) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d6 + 4|text(7) (1d6 + 4) piercing damage plus dice:3d10|text(16)\
    \ (3d10) necrotic damage."
  "name": "Bite"
- "desc": "Melee Weapon Attack: dice: d20+8 (+8) to hit, reach 5 ft., one target.\
    \ Hit: dice:2d6 + 4|text(11) (2d6 + 4) slashing damage plus dice:1d10|text(5)\
    \ (1d10) necrotic damage."
  "name": "Claws"
"source":
- "MTF"
- "DIP"
- "SLW"
- "RtG"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MTF/Rot%20Troll.webp"
```
^statblock

```encounter-table
name: Rot Troll
creatures:
 - 1: Rot Troll
```

## Environment

desert, forest, swamp, underdark