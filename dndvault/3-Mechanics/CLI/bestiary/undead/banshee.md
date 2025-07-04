---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mm
- monster/cr/4
- monster/environment/coastal
- monster/environment/forest
- monster/size/medium
- monster/type/undead
aliases: ["Banshee"]
NoteIcon: monster
BestiaryType: undead
SourceType: Bestiary
BookSource: Monster Manual p. 23, Dragon of Icespire Peak, Tasha's Cauldron of Everything. Available in the Basic Rules.
---
# [Banshee](3-Mechanics\CLI\bestiary\undead/banshee.md)
*Source: Monster Manual p. 23, Dragon of Icespire Peak, Tasha's Cauldron of Everything. Available in the Basic Rules.*  

When night falls, unlucky travelers hear the faint cries of the forlorn dead. This woeful spirit is a banshee, a spiteful creature formed from the spirit of a female elf. Banshees appear as luminous, wispy forms that vaguely recall their mortal features. A banshee's face is wreathed in a wild tangle of hair, its body clad in wispy rags that flutter and stream around it.

## Divine Wrath

Banshees are the undead remnants of elves who, blessed with great beauty, failed to use their gift to bring joy to the world. Instead, they used their beauty to corrupt and control others. Elves afflicted by the banshee's curse experience no gladness, feeling only distress in the presence of the living. As the curse takes its toll, their minds and bodies decay, until death completes their transformation into undead monsters.

## Sorrow Bound

A banshee becomes forever bound to the place of its demise, unable to venture more than five miles from there. It is forced to relive every moment of its life with perfect recall, yet always refuses to accept responsibility for its doom.

## Beauty Hoarders

The vanity that inspired the banshee's cursed creation persists in undeath. These creatures covet beautiful objects: fine jewelery, paintings, statues, and other objects of art. At the same time, a banshee abhors any mirrored surface, for it can't bear to see the horror of its own existence. A single glimpse of itself is enough to send a banshee into a rage.

## Undead Nature

A banshee doesn't require air, food, drink, or sleep.

```statblock
"name": "Banshee"
"size": "Medium"
"type": "undead"
"alignment": "Chaotic Evil"
"ac": !!int "12"
"hp": !!int "58"
"hit_dice": "13d8"
"stats":
- !!int "1"
- !!int "14"
- !!int "10"
- !!int "12"
- !!int "11"
- !!int "17"
"speed": "0 ft., fly 40 ft. (hover)"
"saves":
  "Charisma": !!int "5"
  "Wisdom": !!int "2"
"damage_resistances": "acid; fire; lightning; thunder; bludgeoning, piercing, slashing\
  \ from nonmagical attacks"
"damage_immunities": "cold, necrotic, poison"
"condition_immunities": "[charmed](/3-Mechanics/CLI/rules/conditions.md#charmed),\
  \ [exhaustion](/3-Mechanics/CLI/rules/conditions.md#exhaustion), [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened),\
  \ [grappled](/3-Mechanics/CLI/rules/conditions.md#grappled), [paralyzed](/3-Mechanics/CLI/rules/conditions.md#paralyzed),\
  \ [petrified](/3-Mechanics/CLI/rules/conditions.md#petrified), [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned),\
  \ [prone](/3-Mechanics/CLI/rules/conditions.md#prone), [restrained](/3-Mechanics/CLI/rules/conditions.md#restrained)"
"senses": "darkvision 60 ft., passive Perception 10"
"languages": "Common, Elvish"
"cr": "4"
"traits":
- "desc": "The banshee can magically sense the presence of living creatures up to\
    \ 5 miles away that aren't undead or constructs. She knows the general direction\
    \ they're in but not their exact locations."
  "name": "Detect Life"
- "desc": "The banshee can move through other creatures and objects as if they were\
    \ difficult terrain. She takes dice:1d10|text(5) (1d10) force damage if she\
    \ ends her turn inside an object."
  "name": "Incorporeal Movement"
"actions":
- "desc": "Melee Spell Attack: dice: d20+4 (+4) to hit, reach 5 ft., one target.\
    \ Hit: dice:3d6 + 2|text(12) (3d6 + 2) necrotic damage."
  "name": "Corrupting Touch"
- "desc": "Each non-undead creature within 60 feet of the banshee that can see her\
    \ must succeed on a DC 13 Wisdom saving throw or be [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened)\
    \ for 1 minute. A [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened)\
    \ target can repeat the saving throw at the end of each of its turns, with disadvantage\
    \ if the banshee is within line of sight, ending the effect on itself on a success.\
    \ If a target's saving throw is successful or the effect ends for it, the target\
    \ is immune to the banshee's Horrifying Visage for the next 24 hours."
  "name": "Horrifying Visage"
- "desc": "The banshee releases a mournful wail, provided that she isn't in sunlight.\
    \ This wail has no effect on constructs and undead. All other creatures within\
    \ 30 feet of her that can hear her must make a DC 13 Constitution saving throw.\
    \ On a failure, a creature drops to 0 hit points. On a success, a creature takes\
    \ dice:3d6|text(10) (3d6) psychic damage."
  "name": "Wail (1/Day)"
"source":
- "MM"
- "CoS"
- "WDMM"
- "GoS"
- "DIP"
- "BGDIA"
- "EGW"
- "TCE"
- "WBtW"
- "JttRC"
- "PaBTSO"
- "BMT"
- "QftIS"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MM/Banshee.webp"
```
^statblock

```encounter-table
name: Banshee
creatures:
 - 1: Banshee
```

## Environment

forest, coastal