---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/1
- monster/environment/forest
- monster/size/small
- monster/type/humanoid
aliases: ["Grung Wildling"]
NoteIcon: monster
BestiaryType: humanoid
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 150, Volo's Guide to Monsters p. 157
---
# [Grung Wildling](3-Mechanics\CLI\bestiary\humanoid/grung-wildling-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 150, Volo's Guide to Monsters p. 157*  

```statblock
"name": "Grung Wildling (MPMM)"
"size": "Small"
"type": "humanoid"
"alignment": "Any alignment"
"ac": !!int "16"
"ac_class": "natural armor"
"hp": !!int "27"
"hit_dice": "5d6 + 10"
"stats":
- !!int "7"
- !!int "16"
- !!int "15"
- !!int "10"
- !!int "15"
- !!int "11"
"speed": "25 ft., climb 25 ft."
"saves":
  "Dexterity": !!int "5"
"skillsaves":
  "Athletics": !!int "2"
  "Stealth": !!int "5"
  "Perception": !!int "4"
  "Survival": !!int "4"
"damage_immunities": "poison"
"condition_immunities": "[poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)"
"senses": "passive Perception 14"
"languages": "Grung"
"cr": "1"
"traits":
- "desc": "The grung casts one of the following spells, using Wisdom as the spellcasting\
    \ ability (spell save DC 12):\n\nAt will: [druidcraft](/3-Mechanics/CLI/spells/druidcraft.md)\n\
    \n2/day: [plant growth](/3-Mechanics/CLI/spells/plant-growth.md)\n\n3/day\
    \ each: [cure wounds](/3-Mechanics/CLI/spells/cure-wounds.md), [spike growth](/3-Mechanics/CLI/spells/spike-growth.md)"
  "name": "Spellcasting"
- "desc": "The grung can breathe air and water."
  "name": "Amphibious"
- "desc": "Any creature that grapples the grung or otherwise comes into direct contact\
    \ with the grung's skin must succeed on a DC 12 Constitution saving throw or become\
    \ [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned) for 1 minute. A [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)\
    \ creature no longer in direct contact with the grung can repeat the saving throw\
    \ at the end of each of its turns, ending the effect on itself on a success."
  "name": "Poisonous Skin"
- "desc": "The grung's long jump is up to 25 feet and its high jump is up to 15 feet,\
    \ with or without a running start."
  "name": "Standing Leap"
- "desc": "If the grung isn't immersed in water for at least 1 hour during a day,\
    \ it suffers 1 level of [exhaustion](/3-Mechanics/CLI/rules/conditions.md#exhaustion)\
    \ at the end of that day. The grung can recover from this [exhaustion](/3-Mechanics/CLI/rules/conditions.md#exhaustion)\
    \ only through magic or by immersing itself in water for at least 1 hour."
  "name": "Water Dependency"
"actions":
- "desc": "Melee or Ranged Weapon Attack: dice: d20+5 (+5) to hit, reach 5 ft.\
    \ or range 20/60 ft., one target. Hit: dice:1d4 + 3|text(5) (1d4 + 3) piercing\
    \ damage plus dice:2d4|text(5) (2d4) poison damage."
  "name": "Dagger"
- "desc": "Ranged Weapon Attack: dice: d20+5 (+5) to hit, range 80/320 ft.,\
    \ one target. Hit: dice:1d6 + 3|text(6) (1d6 + 3) piercing damage plus dice:2d4|text(5)\
    \ (2d4) poison damage."
  "name": "Shortbow"
"source":
- "MPMM"
- "VGM"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Grung%20Wildling.webp"
```
^statblock

```encounter-table
name: Grung Wildling
creatures:
 - 1: Grung Wildling
```

## Environment

forest