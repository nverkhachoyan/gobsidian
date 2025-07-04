---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mtf
- monster/cr/12
- monster/environment/coastal
- monster/environment/desert
- monster/environment/forest
- monster/environment/grassland
- monster/environment/mountain
- monster/environment/urban
- monster/size/medium
- monster/type/undead
aliases: ["Eidolon"]
NoteIcon: monster
BestiaryType: undead
SourceType: Bestiary
BookSource: Mordenkainen's Tome of Foes p. 194, Divine Contention, Dragon of Icespire Peak
---
# [Eidolon](3-Mechanics\CLI\bestiary\undead/eidolon-mtf.md)
*Source: Mordenkainen's Tome of Foes p. 194, Divine Contention, Dragon of Icespire Peak*  

> [!quote]- A quote from Mordenkainen  
> 
> It's not just gods that have the power to bind spirits to their idols. Beings such as archdevils can do it with the souls of their cultists. Moloch, for instance.

## Eidolon

The gods have many methods for protecting sites they deem holy. One servant they rely on often to do so is the eidolon, a ghostly spirit bound by a sacred oath to safeguard a place of import to the divine. Forged from the souls of those who had proven their unwavering devotion, eidolons stalk temples and vaults, places where miracles have been witnessed and relics enshrined, to ensure that no enemy can gain a foothold against the gods' cause through defilement or violence within these sites. If an enemy with such intent sets foot inside a warded location, the eidolon plunges into a graven vessel, a statue specially prepared to house the souls of these protectors. The eidolon then animates the effigy and uses the borrowed body to drive out the intruders bent on plundering the relics it is charged with guarding.

## Sacred Guardians

Creating an eidolon requires a spirit of fanatical devotion—that of an individual who, in life, served with unwavering faithfulness. Upon death, a god might reward such a follower with everlasting service in the protection of a holy site. An eidolon has no purpose beyond guarding the place it was assigned to and never leaves.

## Animated Statues

An eidolon has few methods for protecting itself beyond its ability to awaken its sacred vessels. When a foe enters, the eidolon leaps into action by merging its body with one of several statues at the site. After doing so, the eidolon controls the construct as if it was its own body and uses its fists to drive back intruders, smashing and crushing anything it can reach.

## Undead Nature

An eidolon doesn't require air, food, drink, or sleep.

```statblock
"name": "Eidolon (MTF)"
"size": "Medium"
"type": "undead"
"alignment": "Any alignment"
"ac": !!int "9"
"hp": !!int "63"
"hit_dice": "18d8 - 18"
"stats":
- !!int "7"
- !!int "8"
- !!int "9"
- !!int "14"
- !!int "19"
- !!int "16"
"speed": "0 ft., fly 40 ft. (hover)"
"saves":
  "Wisdom": !!int "8"
"skillsaves":
  "Perception": !!int "8"
"damage_resistances": "acid; fire; lightning; thunder; bludgeoning, piercing, slashing\
  \ from nonmagical attacks"
"damage_immunities": "cold, necrotic, poison"
"condition_immunities": "[charmed](/3-Mechanics/CLI/rules/conditions.md#charmed),\
  \ [exhaustion](/3-Mechanics/CLI/rules/conditions.md#exhaustion), [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened),\
  \ [grappled](/3-Mechanics/CLI/rules/conditions.md#grappled), [paralyzed](/3-Mechanics/CLI/rules/conditions.md#paralyzed),\
  \ [petrified](/3-Mechanics/CLI/rules/conditions.md#petrified), [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned),\
  \ [prone](/3-Mechanics/CLI/rules/conditions.md#prone), [restrained](/3-Mechanics/CLI/rules/conditions.md#restrained)"
"senses": "darkvision 60 ft., passive Perception 18"
"languages": "the languages it knew in life"
"cr": "12"
"traits":
- "desc": "The eidolon can move through other creatures and objects as if they were\
    \ difficult terrain. It takes dice:1d10|text(5) (1d10) force damage if it\
    \ ends its turn inside an object other than a sacred statue."
  "name": "Incorporeal Movement"
- "desc": "When the eidolon moves into a space occupied by a sacred statue, the eidolon\
    \ can disappear, causing the statue to become a creature under the eidolon's control.\
    \ The eidolon uses the [sacred statue](/3-Mechanics/CLI/bestiary/construct/sacred-statue-mtf.md)'s\
    \ statistics in place of its own."
  "name": "Sacred Animation (Recharge 5-6)"
- "desc": "The eidolon has advantage on saving throws against any effect that turns\
    \ undead."
  "name": "Turn Resistance"
"actions":
- "desc": "Each creature within 60 feet of the eidolon that can see it must succeed\
    \ on a DC 15 Wisdom saving throw or be [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened)\
    \ for 1 minute. While [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened)\
    \ in this way, the creature must take the Dash action and move away from the eidolon\
    \ by the safest available route at the start of each of its turns, unless there\
    \ is nowhere for it to move, in which case the creature also becomes [stunned](/3-Mechanics/CLI/rules/conditions.md#stunned)\
    \ until it can move again. A [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened)\
    \ target can repeat the saving throw at the end of each of its turns, ending the\
    \ effect on itself on a success. If a target's saving throw is successful or the\
    \ effect ends for it, the target is immune to any eidolon's Divine Dread for the\
    \ next 24 hours."
  "name": "Divine Dread"
"source":
- "MTF"
- "DC"
- "DIP"
- "IMR"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MTF/Eidolon.webp"
```
^statblock

```encounter-table
name: Eidolon
creatures:
 - 1: Eidolon
```

## Environment

coastal, desert, forest, grassland, mountain, urban