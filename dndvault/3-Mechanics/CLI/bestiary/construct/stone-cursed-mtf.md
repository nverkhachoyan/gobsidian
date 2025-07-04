---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mtf
- monster/cr/1
- monster/environment/desert
- monster/environment/mountain
- monster/environment/urban
- monster/size/medium
- monster/type/construct
aliases: ["Stone Cursed"]
NoteIcon: monster
BestiaryType: construct
SourceType: Bestiary
BookSource: Mordenkainen's Tome of Foes p. 240
---
# [Stone Cursed](3-Mechanics\CLI\bestiary\construct/stone-cursed-mtf.md)
*Source: Mordenkainen's Tome of Foes p. 240*  

## Stone Cursed

The stone cursed are spawned through a foul alchemical ritual performed on a humanoid that has been turned to stone. The ritual, which requires a mixture of basilisk blood and the ashes from the burned feathers of a cockatrice, awakens a dim echo of the petrified victim's spirit, animating the statue and turning it into a useful guardian.

## Lingering Spirits

The stone cursed possess a malevolent drive to slay the living, yet they are utterly loyal to whoever performed the ritual to animate them, and they obey that being's orders to the best of their ability. In combat, stony claws that drip with thick, gray sludge emerge from a stone cursed's fingers. This alchemical sludge transforms those slashed by the claws into statues.

## A Strange Harvest

As part of the ritual used to create a stone cursed, a fist-sized obsidian skull forms within the creature's torso. The skull isn't visible while the stone cursed is active, but when it is slain, the statue shatters and the skull clatters to the ground. Because it is the nexus for the alchemy used to create these horrors, a dim echo of the original victim's memories resonates within the skull. A skilled magic-wielder can attempt to extract memories from it to gain insight into the victim's past or find lore that otherwise would be lost.

## Constructed Nature

A stone cursed doesn't require air, food, drink, or sleep.

## Cryptic Whispers

Even though creatures transformed into stone cursed are long dead, a vague whisper of their memories lives on in the obsidian skull embedded within the stone cursed's body. At the end of a short rest, a character can make a DC 20 Intelligence ([Arcana](/3-Mechanics/CLI/rules/skills.md#Arcana)) check to attempt to extract a memory from the skull—a memory that is a response to a verbal question posed by the character to the skull. Once this check is made, whether it succeeds or fails, the skull can't be used in this manner again.

```statblock
"name": "Stone Cursed (MTF)"
"size": "Medium"
"type": "construct"
"alignment": "Lawful Evil"
"ac": !!int "17"
"ac_class": "natural armor"
"hp": !!int "19"
"hit_dice": "3d8 + 4"
"stats":
- !!int "16"
- !!int "5"
- !!int "14"
- !!int "5"
- !!int "8"
- !!int "7"
"speed": "10 ft."
"damage_vulnerabilities": "bludgeoning"
"damage_immunities": "poison"
"condition_immunities": "[charmed](/3-Mechanics/CLI/rules/conditions.md#charmed),\
  \ [exhaustion](/3-Mechanics/CLI/rules/conditions.md#exhaustion), [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened),\
  \ [petrified](/3-Mechanics/CLI/rules/conditions.md#petrified), [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)"
"senses": "passive Perception 9"
"languages": "the languages it knew in life"
"cr": "1"
"traits":
- "desc": "The stone cursed has advantage on the attack rolls of opportunity attacks."
  "name": "Cunning Opportunist"
- "desc": "While the stone cursed remains motionless, it is indistinguishable from\
    \ a normal statue."
  "name": "False Appearance"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+5 (+5) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d10 + 3|text(8) (1d10 + 3) slashing damage, or dice:2d10 +\
    \ 3|text(14) (2d10 + 3) slashing damage if the attack roll had advantage. If\
    \ the target is a creature, it must succeed on a DC 12 Constitution saving throw,\
    \ or it begins to turn to stone and is [restrained](/3-Mechanics/CLI/rules/conditions.md#restrained)\
    \ until the end of its next turn, when it must repeat the saving throw. The effect\
    \ ends if the second save is successful; otherwise the target is [petrified](/3-Mechanics/CLI/rules/conditions.md#petrified)\
    \ for 24 hours."
  "name": "Petrifying Claws"
"source":
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MTF/Stone%20Cursed.webp"
```
^statblock

```encounter-table
name: Stone Cursed
creatures:
 - 1: Stone Cursed
```

## Environment

desert, mountain, urban