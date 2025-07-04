---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/3
- monster/environment/forest
- monster/environment/hill
- monster/environment/swamp
- monster/size/medium
- monster/type/monstrosity
aliases: ["Shadow Mastiff Alpha"]
NoteIcon: monster
BestiaryType: monstrosity
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 215, Volo's Guide to Monsters p. 190
---
# [Shadow Mastiff Alpha](3-Mechanics\CLI\bestiary\monstrosity/shadow-mastiff-alpha-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 215, Volo's Guide to Monsters p. 190*  

```statblock
"name": "Shadow Mastiff Alpha (MPMM)"
"size": "Medium"
"type": "monstrosity"
"alignment": "Typically  Neutral Evil"
"ac": !!int "12"
"hp": !!int "44"
"hit_dice": "8d8 + 8"
"stats":
- !!int "16"
- !!int "14"
- !!int "13"
- !!int "6"
- !!int "12"
- !!int "5"
"speed": "40 ft."
"skillsaves":
  "Stealth": !!int "6"
  "Perception": !!int "5"
"damage_resistances": "bludgeoning, piercing, slashing from nonmagical attacks while\
  \ in dim light or darkness"
"senses": "darkvision 60 ft., passive Perception 15"
"languages": ""
"cr": "3"
"traits":
- "desc": "The shadow mastiff can see ethereal creatures and objects."
  "name": "Ethereal Awareness"
- "desc": "While in bright light created by sunlight, the shadow mastiff has disadvantage\
    \ on attack rolls, ability checks, and saving throws."
  "name": "Sunlight Weakness"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+5 (+5) to hit, reach 5 ft., one target.\
    \ Hit: dice:2d6 + 3|text(10) (2d6 + 3) piercing damage. If the target is\
    \ a creature, it must succeed on a DC 13 Strength saving throw or be knocked [prone](/3-Mechanics/CLI/rules/conditions.md#prone)."
  "name": "Bite"
- "desc": "The shadow mastiff howls. Any Beast or Humanoid within 300 feet of it must\
    \ succeed on a DC 11 Wisdom saving throw or be [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened)\
    \ of it for 1 minute. A [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened)\
    \ target can repeat the saving throw at the end of each of its turns, ending the\
    \ effect on itself on a success. If a target's save is successful or the effect\
    \ ends for it, the target is immune to any shadow mastiff's Terrifying Howl for\
    \ the next 24 hours."
  "name": "Terrifying Howl (Recharge 6)"
"bonus_actions":
- "desc": "While in dim light or darkness, the shadow mastiff becomes [invisible](/3-Mechanics/CLI/rules/conditions.md#invisible),\
    \ along with anything it is wearing or carrying. The invisibility lasts until\
    \ the shadow mastiff uses a bonus action to end it or until the shadow mastiff\
    \ attacks, is in bright light, or is [incapacitated](/3-Mechanics/CLI/rules/conditions.md#incapacitated)."
  "name": "Shadow Blend"
"source":
- "MPMM"
- "VGM"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Shadow%20Mastiff%20Alpha.webp"
```
^statblock

```encounter-table
name: Shadow Mastiff Alpha
creatures:
 - 1: Shadow Mastiff Alpha
```

## Environment

forest, hill, swamp