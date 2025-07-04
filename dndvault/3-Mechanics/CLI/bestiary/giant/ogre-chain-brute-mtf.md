---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mtf
- monster/cr/3
- monster/environment/grassland
- monster/environment/hill
- monster/environment/mountain
- monster/size/large
- monster/type/giant
aliases: ["Ogre Chain Brute"]
NoteIcon: monster
BestiaryType: giant
SourceType: Bestiary
BookSource: Mordenkainen's Tome of Foes p. 221
---
# [Ogre Chain Brute](3-Mechanics\CLI\bestiary\giant/ogre-chain-brute-mtf.md)
*Source: Mordenkainen's Tome of Foes p. 221*  

> [!quote]- A quote from Mordenkainen  
> 
> Vile and violent though they be, ogres prove to be tractable minions for those unafraid to be brutal masters. I prefer giants, however.

## Ogre Chain Brute

An ogre chain brute wields a great spiked chain. It swings this chain with both hands in a wide circle around itself to knock foes off their feet. Alternatively, it can swing the chain in a crushing overhead smash that's nearly impossible to block or deflect.

## Ogres

Ogres are infamously dim-witted, but with enough time and patience, some of them can be trained to carry out specialized missions in battle. The names they are given—the battering ram, the bolt launcher, the chain brute, and the howdah—reflect their particular functions. These jobs are simple, but they're tailored to take advantage of an ogre's strengths.

```statblock
"name": "Ogre Chain Brute (MTF)"
"size": "Large"
"type": "giant"
"alignment": "Chaotic Evil"
"ac": !!int "11"
"ac_class": "[hide armor](/3-Mechanics/CLI/items/hide-armor.md)"
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
"cr": "3"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+6 (+6) to hit, reach 5 ft., one target.\
    \ Hit: dice:2d4 + 4|text(9) (2d4 + 4) bludgeoning damage."
  "name": "Fist"
- "desc": "The ogre swings its chain, and every creature within 10 feet of it must\
    \ make a DC 14 Dexterity saving throw. On a failed saving throw, a creature takes\
    \ dice:1d8 + 4|text(8) (1d8 + 4) bludgeoning damage and is knocked [prone](/3-Mechanics/CLI/rules/conditions.md#prone).\
    \ On a successful save, the creature takes half as much damage and isn't knocked\
    \ [prone](/3-Mechanics/CLI/rules/conditions.md#prone)."
  "name": "Chain Sweep"
- "desc": "Melee Weapon Attack: dice: d20+6 (+6) to hit, reach 10 ft., one target.\
    \ Hit: dice:2d8 + 4|text(13) (2d8 + 4) bludgeoning damage, and the target\
    \ must make a DC 14 Constitution saving throw. On a failure the target is [unconscious](/3-Mechanics/CLI/rules/conditions.md#unconscious)\
    \ for 1 minute. The [unconscious](/3-Mechanics/CLI/rules/conditions.md#unconscious)\
    \ target repeats the saving throw if it takes damage and at the end of each of\
    \ its turns, ending the effect on itself on a success."
  "name": "Chain Smash (Recharge 6)"
"source":
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MTF/Ogre%20Chain%20Brute.webp"
```
^statblock

```encounter-table
name: Ogre Chain Brute
creatures:
 - 1: Ogre Chain Brute
```

## Environment

grassland, hill, mountain