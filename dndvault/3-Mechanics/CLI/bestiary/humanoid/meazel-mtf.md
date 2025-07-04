---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mtf
- monster/cr/1
- monster/environment/desert
- monster/environment/forest
- monster/environment/grassland
- monster/environment/hill
- monster/environment/mountain
- monster/environment/swamp
- monster/environment/underdark
- monster/environment/urban
- monster/size/medium
- monster/type/humanoid/meazel
aliases: ["Meazel"]
NoteIcon: monster
BestiaryType: humanoid (meazel)
SourceType: Bestiary
BookSource: Mordenkainen's Tome of Foes p. 214
---
# [Meazel](3-Mechanics\CLI\bestiary\humanoid/meazel-mtf.md)
*Source: Mordenkainen's Tome of Foes p. 214*  

## Meazel

In places where the Shadowfell washes against the shores of the Material Plane dwell meazels, hateful hermits who left behind their old lives to contemplate their misery in shadow. Now evil burns in their hearts, and they resent any intrusion into their suffering.

## Hateful Hermit

Meazels are all that remain of people who fled into the Shadowfell to escape their mortal existence. There the darkness transformed them, and their bitterness made them twisted and cruel. Now, they loiter near Shadowfell crossings to waylay travelers who venture too close to their lairs.

## Divide and Conquer

The stain of darkness responsible for the existence of meazels imparts to them magical powers that allow them to move through shadows with ease. Merely stepping into one pool of darkness allows a meazel to move to another one. They use this talent to ambush creatures, snatching them around the throat with their strangling cords and then stepping away. Meazels also use this ability to ferry their victims to isolated spots and then leave the hapless souls to the designs of whatever horrors lurk there.

Creatures that are drawn through the shadows by meazels are cursed by the meazels' baleful magic. The curse acts as a beacon; sorrowsworn, undead, and other terrors sense where they are located and descend on the stranded victims to tear them apart.

```statblock
"name": "Meazel (MTF)"
"size": "Medium"
"type": "humanoid"
"subtype": "meazel"
"alignment": "Lawful Evil"
"ac": !!int "13"
"hp": !!int "35"
"hit_dice": "10d8 - 10"
"stats":
- !!int "8"
- !!int "17"
- !!int "9"
- !!int "14"
- !!int "13"
- !!int "10"
"speed": "30 ft."
"skillsaves":
  "Stealth": !!int "5"
  "Perception": !!int "3"
"senses": "darkvision 120 ft., passive Perception 13"
"languages": "Common"
"cr": "1"
"traits":
- "desc": "While in dim light or darkness, the meazel can take the Hide action as\
    \ a bonus action."
  "name": "Shadow Stealth"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+5 (+5) to hit, reach 5 ft., one target\
    \ of the meazel's size or smaller. Hit: dice:1d6 + 3|text(6) (1d6 + 3) bludgeoning\
    \ damage, and the target is [grappled](/3-Mechanics/CLI/rules/conditions.md#grappled)\
    \ (escape DC 13 with disadvantage). Until the grapple ends, the target takes dice:2d6\
    \ + 3|text(10) (2d6 + 3) bludgeoning damage at the start of each of the meazel's\
    \ turns. The meazel can't make weapon attacks while grappling a creature in this\
    \ way."
  "name": "Garrote"
- "desc": "Melee Weapon Attack: dice: d20+5 (+5) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d6 + 3|text(6) (1d6 + 3) piercing damage, plus dice:1d6|text(3)\
    \ (1d6) necrotic damage."
  "name": "Shortsword"
- "desc": "The meazel, any equipment it is wearing or carrying, and any creature it\
    \ is grappling teleport to an unoccupied space within 500 feet of it, provided\
    \ that the starting space and the destination are in dim light or darkness. The\
    \ destination must be a place the meazel has seen before, but it need not be within\
    \ line of sight. If the destination space is occupied, the teleportation leads\
    \ to the nearest unoccupied space.\n\nAny other creature the meazel teleports\
    \ becomes cursed by shadow for 1 hour. Until this curse ends, every undead and\
    \ every creature native to the Shadowfell within 300 feet of the cursed creature\
    \ can sense it, which prevents that creature from hiding from them."
  "name": "Shadow Teleport (Recharge 5-6)"
"source":
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MTF/Meazel.webp"
```
^statblock

```encounter-table
name: Meazel
creatures:
 - 1: Meazel
```

## Environment

desert, forest, grassland, hill, mountain, swamp, underdark, urban