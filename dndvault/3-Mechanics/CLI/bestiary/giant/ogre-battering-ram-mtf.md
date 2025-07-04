---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mtf
- monster/cr/4
- monster/environment/grassland
- monster/environment/hill
- monster/environment/mountain
- monster/size/large
- monster/type/giant
aliases: ["Ogre Battering Ram"]
NoteIcon: monster
BestiaryType: giant
SourceType: Bestiary
BookSource: Mordenkainen's Tome of Foes p. 220
---
# [Ogre Battering Ram](3-Mechanics\CLI\bestiary\giant/ogre-battering-ram-mtf.md)
*Source: Mordenkainen's Tome of Foes p. 220*  

> [!quote]- A quote from Mordenkainen  
> 
> Vile and violent though they be, ogres prove to be tractable minions for those unafraid to be brutal masters. I prefer giants, however.

## Ogre Battering Ram

An ogre battering ram carries an enormous club used primarily for bashing doors into kindling, but which also works well for smashing foes. These ogres are drilled in two simple tasks: rushing forward to shatter enemy fortifications, and using their weapons to force an advancing enemy to halt.

## Ogres

Ogres are infamously dim-witted, but with enough time and patience, some of them can be trained to carry out specialized missions in battle. The names they are given—the battering ram, the bolt launcher, the chain brute, and the howdah—reflect their particular functions. These jobs are simple, but they're tailored to take advantage of an ogre's strengths.

```statblock
"name": "Ogre Battering Ram (MTF)"
"size": "Large"
"type": "giant"
"alignment": "Chaotic Evil"
"ac": !!int "14"
"ac_class": "[ring mail](/3-Mechanics/CLI/items/ring-mail.md)"
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
"cr": "4"
"traits":
- "desc": "The ogre deals double damage to objects and structures."
  "name": "Siege Monster"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+6 (+6) to hit, reach 5 ft., one target.\
    \ Hit: dice:2d10 + 4|text(15) (2d10 + 4) bludgeoning damage, and the ogre\
    \ can push the target 5 feet away if the target is Huge or smaller."
  "name": "Bash"
- "desc": "Until the start of the ogre's next turn, attack rolls against the ogre\
    \ have disadvantage, it has advantage on the attack roll it makes for an opportunity\
    \ attack, and that attack deals an extra dice:3d10|text(16) (3d10) bludgeoning\
    \ damage on a hit. Also, each enemy that tries to move out of the ogre's reach\
    \ without teleporting must succeed on a DC 14 Strength saving throw or have its\
    \ speed reduced to 0 until the start of the ogre's next turn."
  "name": "Block the Path"
"source":
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MTF/Ogre%20Battering%20Ram.webp"
```
^statblock

```encounter-table
name: Ogre Battering Ram
creatures:
 - 1: Ogre Battering Ram
```

## Environment

grassland, hill, mountain