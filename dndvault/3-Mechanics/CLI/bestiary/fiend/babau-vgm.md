---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/vgm
- monster/cr/4
- monster/environment/underdark
- monster/environment/urban
- monster/size/medium
- monster/type/fiend/demon
aliases: ["Babau"]
NoteIcon: monster
BestiaryType: fiend (demon)
SourceType: Bestiary
BookSource: Volo's Guide to Monsters p. 136
---
# [Babau](3-Mechanics\CLI\bestiary\fiend/babau-vgm.md)
*Source: Volo's Guide to Monsters p. 136*  

Demon lords create lesser demons for the purpose of spreading chaos and terror throughout the multiverse.

## Babau

Demons and devils clash endlessly for control of the Lower Planes. One of these battles pitted the legions of the archdevil Glasya against the screaming hordes of the demon lord Graz'zt. It is said that Glasya wounded Graz'zt with her sword, and the first babaus arose where his blood struck the ground. Their sudden appearance helped rout Glasya and secured Graz'zt's place as one of the preeminent demon lords of the Abyss.

A babau demon possesses the cunning of a devil and the bloodthirstiness of a demon. It has leathery black skin pulled tight over its gaunt frame, and a curved horn protruding from the back of its elongated skull. A babau's baleful glare can weaken a creature.

```statblock
"name": "Babau (VGM)"
"size": "Medium"
"type": "fiend"
"subtype": "demon"
"alignment": "Chaotic Evil"
"ac": !!int "16"
"ac_class": "natural armor"
"hp": !!int "82"
"hit_dice": "11d8 + 33"
"stats":
- !!int "19"
- !!int "16"
- !!int "16"
- !!int "11"
- !!int "12"
- !!int "13"
"speed": "40 ft."
"skillsaves":
  "Stealth": !!int "5"
  "Perception": !!int "5"
"damage_resistances": "cold; fire; lightning; bludgeoning, piercing, slashing from\
  \ nonmagical attacks"
"damage_immunities": "poison"
"condition_immunities": "[poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)"
"senses": "darkvision 120 ft., passive Perception 15"
"languages": "Abyssal"
"cr": "4"
"traits":
- "desc": "The babau's innate spellcasting ability is Wisdom (spell save DC 11). The\
    \ babau can innately cast the following spells, requiring no material components:\n\
    \nAt will: [darkness](/3-Mechanics/CLI/spells/darkness.md), [dispel magic](/3-Mechanics/CLI/spells/dispel-magic.md),\
    \ [fear](/3-Mechanics/CLI/spells/fear.md), [heat metal](/3-Mechanics/CLI/spells/heat-metal.md),\
    \ [levitate](/3-Mechanics/CLI/spells/levitate.md)"
  "name": "Innate Spellcasting"
"actions":
- "desc": "The babau makes two melee attacks. It can also use Weakening Gaze before\
    \ or after making these attacks."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+6 (+6) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d8 + 4|text(8) (1d8 + 4) slashing damage."
  "name": "Claw"
- "desc": "Melee or Ranged Weapon Attack: dice: d20+6 (+6) to hit, reach 5 ft.\
    \ or range 20/60 ft., one target. Hit: dice:1d6 + 4|text(7) (1d6 + 4) piercing\
    \ damage, or dice:1d8 + 4|text(8) (1d8 + 4) piercing damage when used with\
    \ two hands to make a melee attack."
  "name": "Spear"
- "desc": "The babau targets one creature that it can see within 20 feet of it. The\
    \ target must make a DC 13 Constitution saving throw. On a failed save, the target\
    \ deals only half damage with weapon attacks that use Strength for 1 minute. The\
    \ target can repeat the saving throw at the end of each of its turns, ending the\
    \ effect on itself on a success."
  "name": "Weakening Gaze"
"source":
- "VGM"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/VGM/Babau.webp"
```
^statblock

```encounter-table
name: Babau
creatures:
 - 1: Babau
```

## Environment

underdark, urban