---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mtf
- monster/cr/15
- monster/environment/swamp
- monster/environment/underdark
- monster/environment/urban
- monster/size/medium
- monster/type/fiend/demon
aliases: ["Nabassu"]
NoteIcon: monster
BestiaryType: fiend (demon)
SourceType: Bestiary
BookSource: Mordenkainen's Tome of Foes p. 135
---
# [Nabassu](3-Mechanics\CLI\bestiary\fiend/nabassu-mtf.md)
*Source: Mordenkainen's Tome of Foes p. 135*  

## Nabassu

The insatiable nabassus prowl the multiverse in search of souls to devour. If they think they can kill a creature and consume its soul, they attack—even if that other creature is a demon, including another nabassu.

## Hated Outcasts

Demons have few rules, and the murder of other demons hardly raises an eyebrow among these fiends. The act of devouring souls is something else. For this reason, most demons shun nabassus and force them to live on the fringes of the Abyss. There, nabassus pick off weaker demons or, if the situation warrants, gather in packs to take down larger prey. Some especially powerful nabassus even search for demon lords' amulets.

## Demonic Infiltrators

Whenever magic pulls demons from the Abyss to the Material Plane, nabassus try to get summoned so that they can embark on a feast of souls there. If a nabassu is summoned, it tries to break free so that it can devour the soul of its summoner and then set out to feed on the souls of whatever creatures it can catch. One way a summoner can avoid this fate is by providing a steady supply of souls to the nabassu, which can cause the demon to be cooperative—for as long as the supply lasts.

```statblock
"name": "Nabassu (MTF)"
"size": "Medium"
"type": "fiend"
"subtype": "demon"
"alignment": "Chaotic Evil"
"ac": !!int "18"
"ac_class": "natural armor"
"hp": !!int "190"
"hit_dice": "20d8 + 100"
"stats":
- !!int "22"
- !!int "14"
- !!int "21"
- !!int "14"
- !!int "15"
- !!int "17"
"speed": "40 ft., fly 60 ft."
"saves":
  "Dexterity": !!int "7"
  "Strength": !!int "11"
"skillsaves":
  "Perception": !!int "7"
"damage_resistances": "cold; fire; lightning; bludgeoning, piercing, slashing from\
  \ nonmagical attacks"
"damage_immunities": "poison"
"condition_immunities": "[poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)"
"senses": "darkvision 60 ft., passive Perception 17"
"languages": "Abyssal, telepathy 120 ft."
"cr": "15"
"traits":
- "desc": "The nabassu darkens the area around its body in a 10-foot radius. Nonmagical\
    \ light can't illuminate this area of dim light."
  "name": "Demonic Shadows"
- "desc": "A nabassu can eat the soul of a creature it has killed within the last\
    \ hour, provided that creature is neither a construct nor an undead. The devouring\
    \ requires the nabassu to be within 5 feet of the corpse for at least 10 minutes,\
    \ after which it gains a number of Hit Dice (dice: d8|avg|noform (d8)s) equal\
    \ to half the creature's number of Hit Dice. Roll those dice, and increase the\
    \ nabassu's hit points by the numbers rolled. For every 4 Hit Dice the nabassu\
    \ gains in this way, its attacks deal an extra dice:1d6|text(3) (1d6) damage\
    \ on a hit. The nabassu retains these benefits for 6 days. A creature devoured\
    \ by a nabassu can be restored to life only by a [wish](/3-Mechanics/CLI/spells/wish.md)\
    \ spell."
  "name": "Devour Soul"
- "desc": "The nabassu has advantage on saving throws against spells and other magical\
    \ effects."
  "name": "Magic Resistance"
- "desc": "The nabassu's weapon attacks are magical."
  "name": "Magic Weapons"
"actions":
- "desc": "The nabassu uses its Soul-Stealing Gaze and makes two attacks: one with\
    \ its claws and one with its bite."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+11 (+11) to hit, reach 5 ft., one\
    \ target. Hit: dice:2d10 + 6|text(17) (2d10 + 6) slashing damage."
  "name": "Claws"
- "desc": "Melee Weapon Attack: dice: d20+11 (+11) to hit, reach 5 ft., one\
    \ target. Hit: dice:4d12 + 6|text(32) (4d12 + 6) piercing damage."
  "name": "Bite"
- "desc": "The nabassu targets one creature it can see within 30 feet of it. If the\
    \ target can see the nabassu and isn't a construct or an undead, it must succeed\
    \ on a DC 16 Charisma saving throw or reduce its hit point maximum by dice:2d12|text(13)\
    \ (2d12) damage and give the nabassu an equal number of temporary hit points.\
    \ This reduction lasts until the target finishes a short or long rest. The target\
    \ dies if its hit point maximum is reduced to 0, and if the target is a humanoid,\
    \ it immediately rises as a [ghoul](/3-Mechanics/CLI/bestiary/undead/ghoul.md)\
    \ under the nabassu's control."
  "name": "Soul-Stealing Gaze"
"source":
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MTF/Nabassu.webp"
```
^statblock

```encounter-table
name: Nabassu
creatures:
 - 1: Nabassu
```

## Environment

swamp, underdark, urban