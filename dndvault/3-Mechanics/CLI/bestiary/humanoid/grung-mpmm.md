---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/1-4
- monster/environment/forest
- monster/size/small
- monster/type/humanoid
aliases: ["Grung"]
NoteIcon: monster
BestiaryType: humanoid
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 149, Volo's Guide to Monsters p. 156
---
# [Grung](3-Mechanics\CLI\bestiary\humanoid/grung-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 149, Volo's Guide to Monsters p. 156*  

```statblock
"name": "Grung (MPMM)"
"size": "Small"
"type": "humanoid"
"alignment": "Any alignment"
"ac": !!int "12"
"hp": !!int "11"
"hit_dice": "2d6 + 4"
"stats":
- !!int "7"
- !!int "14"
- !!int "15"
- !!int "10"
- !!int "11"
- !!int "10"
"speed": "25 ft., climb 25 ft."
"saves":
  "Dexterity": !!int "4"
"skillsaves":
  "Athletics": !!int "2"
  "Stealth": !!int "4"
  "Perception": !!int "2"
  "Survival": !!int "2"
"damage_immunities": "poison"
"condition_immunities": "[poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)"
"senses": "passive Perception 12"
"languages": "Grung"
"cr": "1/4"
"traits":
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
- "desc": "Melee or Ranged Weapon Attack: dice: d20+4 (+4) to hit, reach 5 ft.\
    \ or range 20/60 ft., one target. Hit: dice:1d4 + 2|text(4) (1d4 + 2) piercing\
    \ damage plus dice:2d4|text(5) (2d4) poison damage."
  "name": "Dagger"
"source":
- "MPMM"
- "VGM"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Grung.webp"
```
^statblock

```encounter-table
name: Grung
creatures:
 - 1: Grung
```

## Environment

forest