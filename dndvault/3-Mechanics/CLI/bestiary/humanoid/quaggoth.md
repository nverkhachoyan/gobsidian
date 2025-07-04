---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mm
- monster/cr/2
- monster/environment/underdark
- monster/size/medium
- monster/type/humanoid/quaggoth
aliases: ["Quaggoth"]
NoteIcon: monster
BestiaryType: humanoid (quaggoth)
SourceType: Bestiary
BookSource: Monster Manual p. 256
---
# [Quaggoth](3-Mechanics\CLI\bestiary\humanoid/quaggoth.md)
*Source: Monster Manual p. 256*  

Savage and territorial, quaggoths climb the chasms of the Underdark. They maul their foes in a frenzy, becoming even more murderous in the face of death.

## Quaggoth Origins

Quaggoths were never an enlightened species, but they were not always the brutal Underdark denizens they are today. In a distant age, quaggoth tribes dwelled upon the surface as nocturnal arboreal hunters, possessing their own language and culture. When elves appeared in the mortal realm, they clashed with the quaggoths, eventually driving them to near extinction. Only by fleeing deep into the Underdark did the quaggoths survive.

As they passed the ages deep beneath the world, the quaggoths' fur lost its color and their vision adapted to the darkness, even as the constant danger and weird magic of their new realm transformed them. Turning increasingly brutal and savage, they ate whatever food they could find-and when they could not find it, they preyed on each other. As cannibalism became part of their culture, their past was abandoned.

## Servants of the Drow

The ancient enmity between quaggoths and surface elves makes them easy converts to the dark elf cause. In recent years, the drow have taken an interest in breeding quaggoths, encouraging their ferocity while strengthening their obedience. Wealthy drow houses have legions of quaggoths at their command. Even worse, the drow cultivate the quaggoths' hatred of the elves by leading them on surface raids against known elven enclaves.

## Thonots

Some quaggoths absorb psionic energy that suffuses certain parts of the Underdark. When a tribe discovers that one of its own has inherited such powers, they press it into the role of tribal shaman, or thonot.

A thonots keep a tribe's lore and ensures its superiority against enemies. A thonot that fails the tribe is slain and devoured in a cannibalistic ritual, in the hope that its power passes to another more worthy quaggoth.

## Poison Immunity

Generations of hunting venomous subterranean creatures and perpetual exposure to the molds and fungi that grow in the depths have forced quaggoths to adapt immunities to poisons of all kinds.

```statblock
"name": "Quaggoth"
"size": "Medium"
"type": "humanoid"
"subtype": "quaggoth"
"alignment": "Chaotic Neutral"
"ac": !!int "13"
"ac_class": "natural armor"
"hp": !!int "45"
"hit_dice": "6d8 + 18"
"stats":
- !!int "17"
- !!int "12"
- !!int "16"
- !!int "6"
- !!int "12"
- !!int "7"
"speed": "30 ft., climb 30 ft."
"skillsaves":
  "Athletics": !!int "5"
"damage_immunities": "poison"
"condition_immunities": "[poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)"
"senses": "darkvision 120 ft., passive Perception 11"
"languages": "Undercommon"
"cr": "2"
"traits":
- "desc": "While it has 10 hit points or fewer, the quaggoth has advantage on attack\
    \ rolls. In addition, it deals an extra dice:2d6|text(7) (2d6) damage to any\
    \ target it hits with a melee attack."
  "name": "Wounded Fury"
"actions":
- "desc": "The quaggoth makes two claw attacks."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+5 (+5) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d6 + 3|text(6) (1d6 + 3) slashing damage."
  "name": "Claw"
"source":
- "MM"
- "WDMM"
- "IDRotF"
- "LoX"
- "PaBTSO"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MM/Quaggoth.webp"
```
^statblock

```encounter-table
name: Quaggoth
creatures:
 - 1: Quaggoth
```

## Environment

underdark