---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/3
- monster/environment/desert
- monster/environment/grassland
- monster/size/large
- monster/type/monstrosity
aliases: ["Leucrotta"]
NoteIcon: monster
BestiaryType: monstrosity
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 170, Volo's Guide to Monsters p. 169
---
# [Leucrotta](3-Mechanics\CLI\bestiary\monstrosity/leucrotta-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 170, Volo's Guide to Monsters p. 169*  

```statblock
"name": "Leucrotta (MPMM)"
"size": "Large"
"type": "monstrosity"
"alignment": "Typically  Chaotic Evil"
"ac": !!int "14"
"ac_class": "natural armor"
"hp": !!int "67"
"hit_dice": "9d10 + 18"
"stats":
- !!int "18"
- !!int "14"
- !!int "15"
- !!int "9"
- !!int "12"
- !!int "6"
"speed": "50 ft."
"skillsaves":
  "Deception": !!int "2"
  "Perception": !!int "5"
"senses": "darkvision 60 ft., passive Perception 15"
"languages": "Abyssal, Gnoll"
"cr": "3"
"traits":
- "desc": "The leucrotta can mimic Beast sounds and Humanoid voices. A creature that\
    \ hears the sounds can tell they are imitations only with a successful DC 14 Wisdom\
    \ ([Insight](/3-Mechanics/CLI/rules/skills.md#Insight)) check."
  "name": "Mimicry"
- "desc": "Any creature other than a leucrotta or gnoll that starts its turn within\
    \ 5 feet of the leucrotta must succeed on a DC 12 Constitution saving throw or\
    \ be [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned) until the start\
    \ of the creature's next turn. On a successful saving throw, the creature is immune\
    \ to the Stench of all leucrottas for 1 hour."
  "name": "Stench"
"actions":
- "desc": "The leucrotta makes one Bite attack and one Hooves attack."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+6 (+6) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d8 + 4|text(8) (1d8 + 4) piercing damage. If the leucrotta\
    \ scores a critical hit, it rolls the damage dice three times, instead of twice."
  "name": "Bite"
- "desc": "Melee Weapon Attack: dice: d20+6 (+6) to hit, reach 5 ft., one target.\
    \ Hit: dice:2d6 + 4|text(11) (2d6 + 4) bludgeoning damage."
  "name": "Hooves"
"bonus_actions":
- "desc": "Immediately after the leucrotta makes a Hooves attack, it takes the [Disengage](/3-Mechanics/CLI/rules/actions.md#Disengage)\
    \ action."
  "name": "Kicking Retreat"
"source":
- "MPMM"
- "VGM"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Leucrotta.webp"
```
^statblock

```encounter-table
name: Leucrotta
creatures:
 - 1: Leucrotta
```

## Environment

desert, grassland