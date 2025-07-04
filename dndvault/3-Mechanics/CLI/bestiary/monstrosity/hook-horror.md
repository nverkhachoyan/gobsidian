---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mm
- monster/cr/3
- monster/environment/underdark
- monster/size/large
- monster/type/monstrosity
aliases: ["Hook Horror"]
NoteIcon: monster
BestiaryType: monstrosity
SourceType: Bestiary
BookSource: Monster Manual p. 189
---
# [Hook Horror](3-Mechanics\CLI\bestiary\monstrosity/hook-horror.md)
*Source: Monster Manual p. 189*  

A fierce predator of the Underdark, the hook horror aggressively defends its hunting grounds. The subterranean caverns where these creatures dwell echo with the constant clacking and scraping of their hooks as they wend their way up cliffs and along cavern walls. The monstrous hook horror has a head resembling a vulture's and the torso of an enormous beetle, with an exoskeleton studded by sharp, bony protuberances. It gains its name from its long, powerfully built arms and legs, which end in wickedly curved hooked claws.

## Echoes in the Dark

Hook horrors communicate by striking their hooks against their exoskeletons or the stone surfaces around them. What sounds to others like random clacking noise is actually a complex language that only hook horrors understand, and which carries for miles through the echoing Underdark.

## Pack Predators

The omnivorous hook horrors eat lichens, fungi, plants, and any creature they can catch. A hook horror's hooked limbs give it excellent purchase on rock surfaces, and these creatures use their climbing skills to ambush prey from above. They hunt in packs, working together against the largest and most dangerous opponents. If a battle goes poorly, a hook horror quickly climbs a cavern wall to flee.

## Dedicated Clans

Hook horrors live in extended family groups or clans. Each clan is ruled by the eldest female, who typically places her mate in charge of the clan's hunters. Hook horrors lay eggs, which are clustered in a central, well-defended area of a clan's home caverns.

```statblock
"name": "Hook Horror"
"size": "Large"
"type": "monstrosity"
"alignment": "Neutral"
"ac": !!int "15"
"ac_class": "natural armor"
"hp": !!int "75"
"hit_dice": "10d10 + 20"
"stats":
- !!int "18"
- !!int "10"
- !!int "15"
- !!int "6"
- !!int "12"
- !!int "7"
"speed": "30 ft., climb 30 ft."
"skillsaves":
  "Perception": !!int "3"
"senses": "blindsight 60 ft., darkvision 120 ft., passive Perception 13"
"languages": "Hook Horror"
"cr": "3"
"traits":
- "desc": "The hook horror can't use its blindsight while [deafened](/3-Mechanics/CLI/rules/conditions.md#deafened)."
  "name": "Echolocation"
- "desc": "The hook horror has advantage on Wisdom ([Perception](/3-Mechanics/CLI/rules/skills.md#Perception))\
    \ checks that rely on hearing."
  "name": "Keen Hearing"
"actions":
- "desc": "The hook horror makes two hook attacks."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+6 (+6) to hit, reach 10 ft., one target.\
    \ Hit: dice:2d6 + 4|text(11) (2d6 + 4) piercing damage."
  "name": "Hook"
"source":
- "MM"
- "PotA"
- "TftYP"
- "WDMM"
- "LoX"
- "KftGV"
- "PaBTSO"
- "QftIS"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MM/Hook%20Horror.webp"
```
^statblock

```encounter-table
name: Hook Horror
creatures:
 - 1: Hook Horror
```

## Environment

underdark