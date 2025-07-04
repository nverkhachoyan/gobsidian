---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mtf
- monster/cr/2
- monster/environment/mountain
- monster/environment/underdark
- monster/size/medium
- monster/type/construct
aliases: ["Duergar Hammerer"]
NoteIcon: monster
BestiaryType: construct
SourceType: Bestiary
BookSource: Mordenkainen's Tome of Foes p. 188
---
# [Duergar Hammerer](3-Mechanics\CLI\bestiary\construct/duergar-hammerer-mtf.md)
*Source: Mordenkainen's Tome of Foes p. 188*  

## Duergar Hammerer

The hammerer is a digging machine with a duergar strapped inside it—typically a punishment for those whose work ethic wavers. The machine's mechanism transforms the captive duergar's pain into energy that powers the device, which is typically used to dig tunnels and repel invaders.

> [!quote]- A quote from Mordenkainen  
> 
> You think dwarves are dour and unpleasant? Wait until you meet a duergar.

The cruel duergar plot not only to defeat other dwarves, but to cast down the entire dwarven pantheon in revenge for their being abandoned and left to be enslaved by mind flayers. To this end, duergar train warriors to fulfill a variety of roles.

```statblock
"name": "Duergar Hammerer (MTF)"
"size": "Medium"
"type": "construct"
"alignment": "Lawful Evil"
"ac": !!int "17"
"ac_class": "natural armor"
"hp": !!int "33"
"hit_dice": "6d8 + 6"
"stats":
- !!int "17"
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
"cr": "2"
"traits":
- "desc": "Once per turn, a creature that attacks the hammerer can target the duergar\
    \ trapped in it. The attacker has disadvantage on the attack roll. On a hit, the\
    \ attack deals an extra dice:1d10|text(5) (1d10) damage to the hammerer, and\
    \ the hammerer can respond by using its Multiattack with its reaction."
  "name": "Engine of Pain"
- "desc": "The hammerer deals double damage to objects and structures."
  "name": "Siege Monster"
"actions":
- "desc": "The hammerer makes two attacks: one with its claw and one with its hammer."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+5 (+5) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d6 + 3|text(6) (1d6 + 3) bludgeoning damage."
  "name": "Claw"
- "desc": "Melee Weapon Attack: dice: d20+5 (+5) to hit, reach 5 ft., one target.\
    \ Hit: dice:2d6 + 3|text(10) (2d6 + 3) bludgeoning damage."
  "name": "Hammer"
"source":
- "MTF"
- "IDRotF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MTF/Duergar%20Hammerer.webp"
```
^statblock

```encounter-table
name: Duergar Hammerer
creatures:
 - 1: Duergar Hammerer
```

## Environment

mountain, underdark