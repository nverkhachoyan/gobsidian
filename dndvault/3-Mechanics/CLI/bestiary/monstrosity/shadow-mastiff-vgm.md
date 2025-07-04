---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/vgm
- monster/cr/2
- monster/environment/forest
- monster/environment/hill
- monster/environment/swamp
- monster/size/medium
- monster/type/monstrosity
aliases: ["Shadow Mastiff"]
NoteIcon: monster
BestiaryType: monstrosity
SourceType: Bestiary
BookSource: Volo's Guide to Monsters p. 190
---
# [Shadow Mastiff](3-Mechanics\CLI\bestiary\monstrosity/shadow-mastiff-vgm.md)
*Source: Volo's Guide to Monsters p. 190*  

These black hounds of the Shadowfell move invisibly through the shadows, always on the hunt. In gloomy places where the veil between the Shadowfell and the Material Plane is thinnest, they can cross over into the dark realms of the world.

## Ravenous Lurkers

Shadow mastiffs hunt in packs on the Shadowfell, so when one of them enters a rift between the planes, several more are sure to follow. Each pack is led by an alpha (male or female) that is the smartest and toughest one of the group. The alpha must remain sharp to keep the rest of the pack in line, lest it be killed and replaced.

When a shadow mastiff pack is hungry and senses prey nearby, the alpha lets loose a howl that strikes fear into the hearts of nearby beasts and humanoids. Its howl is also a signal to the rest of the pack to move in for the kill. Gloom provides a shadow mastiff with supernatural protection, granting it resistance to nonmagical weapons while in dim light or darkness. Shadow mastiffs can tolerate bright light, but they shun sunlight.

## Summoned for Service

Some faiths devoted to deities of gloom and night, such as Shar in the Forgotten Realms, perform unholy rites to summon shadow mastiffs from the Shadowfell and then put them to work as temple sentinels, bodyguards, and punishers of nonbelievers, heretics, and apostates. The method for bringing shadow mastiffs into the world is also known by other strong-willed and evil-minded individuals, who find use for the hounds as guards in their strongholds.

## Ethereal Sight

In addition to its other capabilities, a shadow mastiff can see creatures and objects on the Ethereal Plane. This extraplanar perception makes a mastiff an especially skilled guardian, especially in situations when magical or spiritual incursion is likely.

```statblock
"name": "Shadow Mastiff (VGM)"
"size": "Medium"
"type": "monstrosity"
"alignment": "Neutral Evil"
"ac": !!int "12"
"hp": !!int "33"
"hit_dice": "6d8 + 6"
"stats":
- !!int "16"
- !!int "14"
- !!int "13"
- !!int "5"
- !!int "12"
- !!int "5"
"speed": "40 ft."
"skillsaves":
  "Stealth": !!int "6"
  "Perception": !!int "3"
"damage_resistances": "bludgeoning, piercing, slashing from nonmagical attacks while\
  \ in dim light or darkness"
"senses": "darkvision 60 ft., passive Perception 13"
"languages": ""
"cr": "2"
"traits":
- "desc": "The shadow mastiff can see ethereal creatures and objects."
  "name": "Ethereal Awareness"
- "desc": "The shadow mastiff has advantage on Wisdom ([Perception](/3-Mechanics/CLI/rules/skills.md#Perception))\
    \ checks that rely on hearing or smell."
  "name": "Keen Hearing and Smell"
- "desc": "While in dim light or darkness, the shadow mastiff can use a bonus action\
    \ to become [invisible](/3-Mechanics/CLI/rules/conditions.md#invisible), along\
    \ with anything it is wearing or carrying. The invisibility lasts until the shadow\
    \ mastiff uses a bonus action to end it or until the shadow mastiff attacks, is\
    \ in bright light, or is [incapacitated](/3-Mechanics/CLI/rules/conditions.md#incapacitated)."
  "name": "Shadow Blend"
- "desc": "While in bright light created by sunlight, the shadow mastiff has disadvantage\
    \ on attack rolls, ability checks, and saving throws."
  "name": "Sunlight Weakness"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+5 (+5) to hit, reach 5 ft., one target.\
    \ Hit: dice:2d6 + 3|text(10) (2d6 + 3) piercing damage. If the target is\
    \ a creature, it must succeed on a DC 13 Strength saving throw or be knocked [prone](/3-Mechanics/CLI/rules/conditions.md#prone)."
  "name": "Bite"
"source":
- "VGM"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/VGM/Shadow%20Mastiff.webp"
```
^statblock

```encounter-table
name: Shadow Mastiff
creatures:
 - 1: Shadow Mastiff
```

## Environment

forest, swamp, hill