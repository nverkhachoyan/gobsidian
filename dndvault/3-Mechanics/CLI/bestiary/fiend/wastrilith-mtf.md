---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mtf
- monster/cr/13
- monster/environment/coastal
- monster/environment/swamp
- monster/environment/underdark
- monster/environment/underwater
- monster/size/large
- monster/type/fiend/demon
aliases: ["Wastrilith"]
NoteIcon: monster
BestiaryType: fiend (demon)
SourceType: Bestiary
BookSource: Mordenkainen's Tome of Foes p. 139
---
# [Wastrilith](3-Mechanics\CLI\bestiary\fiend/wastrilith-mtf.md)
*Source: Mordenkainen's Tome of Foes p. 139*  

> [!quote]- A quote from Mordenkainen  
> 
> The corruption left behind when a wastrilith visits the world can persist for decades. If left unchecked, it can become a bridge to the Abyss.

## Wastrilith

Found in the waters of the Abyss and other bodies of water contaminated by the plane's fell influence, wastriliths establish themselves as lords of the deep and rule their dominions with cruelty.

## Despoilers

A wastrilith contaminates the waters around it. Its noxious presence even affects nearby sources of water when the demon travels on land. The corrupted water, which contains a measure of the demon's essence, responds to its commands—perhaps hardening to prevent foes from escaping, or erupting in a surge that drags would-be victims into its reach.

## Silent Corrupters

Creatures that ingest water that has been corrupted by a wastrilith risk their very souls. Those who drink the poisonous liquid might wither away until they finally die, or remain alive only to become a thrall of chaos and evil.

To represent this defilement, you can use the optional rule on planar effects||abyssal corruption from  in chapter 2 of the "Dungeon Master's Guide", causing the [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned) creature to be corrupted.

```statblock
"name": "Wastrilith (MTF)"
"size": "Large"
"type": "fiend"
"subtype": "demon"
"alignment": "Chaotic Evil"
"ac": !!int "18"
"ac_class": "natural armor"
"hp": !!int "157"
"hit_dice": "15d10 + 75"
"stats":
- !!int "19"
- !!int "18"
- !!int "21"
- !!int "19"
- !!int "12"
- !!int "14"
"speed": "30 ft., swim 80 ft."
"saves":
  "Strength": !!int "9"
  "Constitution": !!int "10"
"damage_resistances": "cold; fire; lightning; bludgeoning, piercing, slashing from\
  \ nonmagical attacks"
"damage_immunities": "poison"
"condition_immunities": "[poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)"
"senses": "darkvision 120 ft., passive Perception 11"
"languages": "Abyssal, telepathy 120 ft."
"cr": "13"
"traits":
- "desc": "The wastrilith can breathe air and water."
  "name": "Amphibious"
- "desc": "At the start of each of the wastrilith's turns, exposed water within 30\
    \ feet of it is befouled. Underwater, this effect lightly obscures the area until\
    \ a current clears it away. Water in containers remains corrupted until it evaporates.\n\
    \nA creature that consumes this foul water or swims in it must make a DC 18 Constitution\
    \ saving throw. On a successful save, the creature is immune to the foul water\
    \ for 24 hours. On a failed save, the creature takes dice:4d6|text(14) (4d6)\
    \ poison damage and is [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)\
    \ for 1 minute. At the end of this time, the [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)\
    \ creature must repeat the saving throw. On a failure, the creature takes dice:4d8|text(18)\
    \ (4d8) poison damage and is [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)\
    \ until it finishes a long rest.\n\nIf another demon drinks the foul water as\
    \ an action, it gains dice:2d10|text(11) (2d10) temporary hit points."
  "name": "Corrupt Water"
- "desc": "The wastrilith has advantage on saving throws against spells and other\
    \ magical effects."
  "name": "Magic Resistance"
- "desc": "As a bonus action when the wastrilith is underwater, it can cause all water\
    \ within 60 feet of it to be difficult terrain for other creatures until the start\
    \ of its next turn."
  "name": "Undertow"
"actions":
- "desc": "The wastrilith uses Grasping Spout and makes three attacks: one with its\
    \ bite and two with its claws."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+9 (+9) to hit, reach 10 ft., one target.\
    \ Hit: dice:4d12 + 4|text(30) (4d12 + 4) piercing damage."
  "name": "Bite"
- "desc": "Melee Weapon Attack: dice: d20+9 (+9) to hit, reach 10 ft., one target.\
    \ Hit: dice:4d6 + 4|text(18) (4d6 + 4) slashing damage."
  "name": "Claws"
- "desc": "The wastrilith magically launches a spout of water at one creature it can\
    \ see within 60 feet of it. The target must make a DC 17 Strength saving throw,\
    \ and it has disadvantage if it's underwater. On a failed save, it takes dice:4d8\
    \ + 4|text(22) (4d8 + 4) acid damage and is pulled up to 60 feet toward the\
    \ wastrilith. On a successful save, it takes half as much damage and isn't pulled."
  "name": "Grasping Spout"
"source":
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MTF/Wastrilith.webp"
```
^statblock

```encounter-table
name: Wastrilith
creatures:
 - 1: Wastrilith
```

## Environment

coastal, swamp, underdark, underwater