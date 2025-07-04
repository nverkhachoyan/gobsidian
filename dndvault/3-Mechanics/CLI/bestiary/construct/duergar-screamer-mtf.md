---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mtf
- monster/cr/3
- monster/environment/underdark
- monster/size/medium
- monster/type/construct
aliases: ["Duergar Screamer"]
NoteIcon: monster
BestiaryType: construct
SourceType: Bestiary
BookSource: Mordenkainen's Tome of Foes p. 190
---
# [Duergar Screamer](3-Mechanics\CLI\bestiary\construct/duergar-screamer-mtf.md)
*Source: Mordenkainen's Tome of Foes p. 190*  

## Duergar Screamer

A duergar screamer is a construct that uses sonic energy to grind rock into dust. Duergar accused of spreading gossip or plotting against their superiors are trapped within one of these devices, their beards shorn and their ongoing agony channeled into psionic energy that powers the screamer.

> [!quote]- A quote from Mordenkainen  
> 
> You think dwarves are dour and unpleasant? Wait until you meet a duergar.

The cruel duergar plot not only to defeat other dwarves, but to cast down the entire dwarven pantheon in revenge for their being abandoned and left to be enslaved by mind flayers. To this end, duergar train warriors to fulfill a variety of roles.

```statblock
"name": "Duergar Screamer (MTF)"
"size": "Medium"
"type": "construct"
"alignment": "Lawful Evil"
"ac": !!int "15"
"ac_class": "natural armor"
"hp": !!int "38"
"hit_dice": "7d8 + 7"
"stats":
- !!int "18"
- !!int "7"
- !!int "12"
- !!int "5"
- !!int "5"
- !!int "5"
"speed": "20 ft."
"damage_immunities": "poison"
"condition_immunities": "[charmed](/3-Mechanics/CLI/rules/conditions.md#charmed),\
  \ [exhaustion](/3-Mechanics/CLI/rules/conditions.md#exhaustion), [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened),\
  \ [paralyzed](/3-Mechanics/CLI/rules/conditions.md#paralyzed), [petrified](/3-Mechanics/CLI/rules/conditions.md#petrified),\
  \ [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)"
"senses": "darkvision 60 ft., passive Perception 7"
"languages": "understands Dwarvish but can't speak"
"cr": "3"
"traits":
- "desc": "Once per turn, a creature that attacks the screamer can target the duergar\
    \ trapped in it. The attacker has disadvantage on the attack roll. On a hit, the\
    \ attack deals an extra dice:2d10|text(11) (2d10) damage to the screamer,\
    \ and the screamer can respond by using its Multiattack with its reaction."
  "name": "Engine of Pain"
"actions":
- "desc": "The screamer makes one drill attack and uses its Sonic Scream."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+6 (+6) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d12 + 4|text(10) (1d12 + 4) piercing damage."
  "name": "Drill"
- "desc": "The screamer emits destructive energy in a 15-foot cube. Each creature\
    \ in that area must succeed on a DC 11 Strength saving throw or take dice:2d6|text(7)\
    \ (2d6) thunder damage and be knocked [prone](/3-Mechanics/CLI/rules/conditions.md#prone)."
  "name": "Sonic Scream"
"source":
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MTF/Duergar%20Screamer.webp"
```
^statblock

```encounter-table
name: Duergar Screamer
creatures:
 - 1: Duergar Screamer
```

## Environment

underdark