---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mm
- monster/cr/6
- monster/environment/coastal
- monster/environment/desert
- monster/environment/grassland
- monster/environment/hill
- monster/environment/mountain
- monster/environment/underdark
- monster/size/huge
- monster/type/giant
aliases: ["Cyclops"]
NoteIcon: monster
BestiaryType: giant
SourceType: Bestiary
BookSource: Monster Manual p. 45. Available in the Basic Rules.
---
# [Cyclops](3-Mechanics\CLI\bestiary\giant/cyclops.md)
*Source: Monster Manual p. 45. Available in the Basic Rules.*  

Cyclopes are one-eyed giants that eke out a meager existence in wild lands. Isolationists by nature, they avoid contact with other races and try to drive away strangers in their territory.

## Nonreligious

Legends claim that the cyclopes are the spawn of one of the gods of the giants, but these creatures pay little heed to any deities. They see little benefit in prayer and dislike ritual, which they perceive as complex and foreign. However, a cyclops that gains direct benefit from some site of divine power, or which is threatened by a supernatural force or creature, will pay homage as long as the benefit or threat remains.

## Unsophisticated

Though they are reasonably intelligent, cyclopes live simple, reclusive lives, keeping herds of animals for food. They prefer to dwell alone or in small family groups, lairing in caves, ruins, or rough structures of dry stone construction they build themselves. A cyclops keeps its herd animals with it at night, sealing the entrance to its home with boulders to let it serve double duty as a barn.

A cyclops lairs within a day's journey of other cyclopes, so that they can meet to trade goods or seek mates. They craft weapons and tools of wood and stone, but will use metal when they can find it. Although cyclopes understand the Giant tongue, they write nothing and speak little, using grunts and gestures for their interactions with each other. Cyclopes don't use money for trade, but they value gold, shells, and other glittering and colorful objects as jewelry. A cyclops might wear a necklace strung with feathers and silver coins, but also with pewter goblets, cutlery, and other bits of ruined metal.

## Unwise

Cyclopes aren't great thinkers or strategists. Slow to learn and bound to their traditional ways, they find innovation difficult. Although they are a terrifying threat in combat due to their size and strength, they can often be tricked by clever foes. Cyclopes can be cowed and awed by obvious displays of magic. Rustics with little exposure to magic, they can be deceived into mistaking a warlock, cleric, or other caster for a powerful divine figure. However, their sense of pride causes them to react with vengeful, bloodthirsty violence once they learn that the individual they assumed was a "god" is a mere mortal.

```statblock
"name": "Cyclops"
"size": "Huge"
"type": "giant"
"alignment": "Chaotic Neutral"
"ac": !!int "14"
"ac_class": "natural armor"
"hp": !!int "138"
"hit_dice": "12d12 + 60"
"stats":
- !!int "22"
- !!int "11"
- !!int "20"
- !!int "8"
- !!int "6"
- !!int "10"
"speed": "30 ft."
"senses": "passive Perception 8"
"languages": "Giant"
"cr": "6"
"traits":
- "desc": "The cyclops has disadvantage on any attack roll against a target more than\
    \ 30 feet away."
  "name": "Poor Depth Perception"
"actions":
- "desc": "The cyclops makes two greatclub attacks."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+9 (+9) to hit, reach 10 ft., one target.\
    \ Hit: dice:3d8 + 6|text(19) (3d8 + 6) bludgeoning damage."
  "name": "Greatclub"
- "desc": "Ranged Weapon Attack: dice: d20+9 (+9) to hit, range 30/120 ft.,\
    \ one target. Hit: dice:4d10 + 6|text(28) (4d10 + 6) bludgeoning damage."
  "name": "Rock"
"source":
- "MM"
- "RoT"
- "ToA"
- "GoS"
- "EGW"
- "MOT"
- "CM"
- "WBtW"
- "DSotDQ"
- "SatO"
- "VEoR"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MM/Cyclops.webp"
```
^statblock

```encounter-table
name: Cyclops
creatures:
 - 1: Cyclops
```

## Environment

underdark, mountain, grassland, hill, desert, coastal