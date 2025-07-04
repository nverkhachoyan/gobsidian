---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mtf
- monster/cr/2
- monster/environment/grassland
- monster/environment/hill
- monster/size/large
- monster/type/giant
aliases: ["Ogre Howdah"]
NoteIcon: monster
BestiaryType: giant
SourceType: Bestiary
BookSource: Mordenkainen's Tome of Foes p. 221
---
# [Ogre Howdah](3-Mechanics\CLI\bestiary\giant/ogre-howdah-mtf.md)
*Source: Mordenkainen's Tome of Foes p. 221*  

> [!quote]- A quote from Mordenkainen  
> 
> Vile and violent though they be, ogres prove to be tractable minions for those unafraid to be brutal masters. I prefer giants, however.

## Ogre Howdah

The most unusual of the specialized ogres, the howdah carries a palisaded wooden fort on its back. The fort is big enough to serve as a fighting platform for up to four small humanoids. Ogre howdahs are most often seen bearing goblins equipped with bows and spears into battle, but they could just as easily transport kobolds, deep gnomes, or other humanoids of similar size.

## Ogres

Ogres are infamously dim-witted, but with enough time and patience, some of them can be trained to carry out specialized missions in battle. The names they are given—the battering ram, the bolt launcher, the chain brute, and the howdah—reflect their particular functions. These jobs are simple, but they're tailored to take advantage of an ogre's strengths.

```statblock
"name": "Ogre Howdah (MTF)"
"size": "Large"
"type": "giant"
"alignment": "Chaotic Evil"
"ac": !!int "13"
"ac_class": "[breastplate](/3-Mechanics/CLI/items/breastplate.md)"
"hp": !!int "59"
"hit_dice": "7d10 + 21"
"stats":
- !!int "19"
- !!int "8"
- !!int "16"
- !!int "5"
- !!int "7"
- !!int "7"
"speed": "40 ft."
"senses": "darkvision 60 ft., passive Perception 8"
"languages": "Common, Giant"
"cr": "2"
"traits":
- "desc": "The ogre carries a compact fort on its back. Up to four Small creatures\
    \ can ride in the fort without squeezing. To make a melee attack against a target\
    \ within 5 feet of the ogre, they must use spears or weapons with reach. Creatures\
    \ in the fort have three-quarters cover against attacks and effects from outside\
    \ it. If the ogre dies, creatures in the fort are placed in unoccupied spaces\
    \ within 5 feet of the ogre."
  "name": "Howdah"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+6 (+6) to hit, reach 5 ft., one target.\
    \ Hit: dice:2d6 + 4|text(11) (2d6 + 4) bludgeoning damage."
  "name": "Mace"
"source":
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MTF/Ogre%20Howdah.webp"
```
^statblock

```encounter-table
name: Ogre Howdah
creatures:
 - 1: Ogre Howdah
```

## Environment

grassland, hill