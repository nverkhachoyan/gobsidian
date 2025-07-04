---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mm
- monster/cr/2
- monster/environment/underdark
- monster/size/tiny
- monster/type/aberration
aliases: ["Intellect Devourer"]
NoteIcon: monster
BestiaryType: aberration
SourceType: Bestiary
BookSource: Monster Manual p. 191
---
# [Intellect Devourer](3-Mechanics\CLI\bestiary\aberration/intellect-devourer.md)
*Source: Monster Manual p. 191*  

An intellect devourer resembles a walking brain protected by a crusty covering and set on bestial clawed legs. This foul aberration feeds on the intelligence of sentient creatures, taking over a victim's body on behalf of its mind flayer masters.

## Illithid Creations

Mind flayers breed intellect devourers to serve as roaming hunters of the Underdark, creating an intellect devourer by taking the brain of a thrall and subjecting it to a horrible ritual. As it sprouts legs, the brain becomes an intelligent predator as twisted and evil as its masters.

## Deadly Puppet Masters

An intellect devourer consumes a creature's mind and memories, then turns the host body into a puppet under its control. An intellect devourer typically uses its puppet host to lure others into the domain of the mind flayers to be enthralled or consumed.

> [!quote]- A quote from Qorik el-Slurrk, mind flayer  
> 
> Don't cry. We have no intention of eating your brain. In fact, your brain is going on a wonderful journey!


```statblock
"name": "Intellect Devourer"
"size": "Tiny"
"type": "aberration"
"alignment": "Lawful Evil"
"ac": !!int "12"
"hp": !!int "21"
"hit_dice": "6d4 + 6"
"stats":
- !!int "6"
- !!int "14"
- !!int "13"
- !!int "12"
- !!int "11"
- !!int "10"
"speed": "40 ft."
"skillsaves":
  "Stealth": !!int "4"
  "Perception": !!int "2"
"damage_resistances": "bludgeoning, piercing, slashing from nonmagical attacks"
"condition_immunities": "[blinded](/3-Mechanics/CLI/rules/conditions.md#blinded)"
"senses": "blindsight 60 ft. (blind beyond this radius), passive Perception 12"
"languages": "understands Deep Speech but can't speak, telepathy 60 ft."
"cr": "2"
"traits":
- "desc": "The intellect devourer can sense the presence and location of any creature\
    \ within 300 feet of it that has an Intelligence of 3 or higher, regardless of\
    \ interposing barriers, unless the creature is protected by a [mind blank](/3-Mechanics/CLI/spells/mind-blank.md)\
    \ spell."
  "name": "Detect Sentience"
"actions":
- "desc": "The intellect devourer makes one attack with its claws and uses Devour\
    \ Intellect."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+4 (+4) to hit, reach 5 ft., one target.\
    \ Hit: dice:2d4 + 2|text(7) (2d4 + 2) slashing damage."
  "name": "Claws"
- "desc": "The intellect devourer targets one creature it can see within 10 feet of\
    \ it that has a brain. The target must succeed on a DC 12 Intelligence saving\
    \ throw against this magic or take dice:2d10|text(11) (2d10) psychic damage.\
    \ Also on a failure, roll dice: 3d6|avg|noform (3d6): If the total equals\
    \ or exceeds the target's Intelligence score, that score is reduced to 0. The\
    \ target is [stunned](/3-Mechanics/CLI/rules/conditions.md#stunned) until it regains\
    \ at least one point of Intelligence."
  "name": "Devour Intellect"
- "desc": "The intellect devourer initiates an Intelligence contest with an [incapacitated](/3-Mechanics/CLI/rules/conditions.md#incapacitated)\
    \ humanoid within 5 feet of it that isn't protected by [protection from evil and\
    \ good](/3-Mechanics/CLI/spells/protection-from-evil-and-good.md). If it wins\
    \ the contest, the intellect devourer magically consumes the target's brain, teleports\
    \ into the target's skull, and takes control of the target's body. While inside\
    \ a creature, the intellect devourer has total cover against attacks and other\
    \ effects originating outside its host. The intellect devourer retains its Intelligence,\
    \ Wisdom, and Charisma scores, as well as its understanding of Deep Speech, its\
    \ telepathy, and its traits. It otherwise adopts the target's statistics. It knows\
    \ everything the creature knew, including spells and languages.\n\nIf the host\
    \ body dies, the intellect devourer must leave it. A [protection from evil and\
    \ good](/3-Mechanics/CLI/spells/protection-from-evil-and-good.md) spell cast on\
    \ the body drives the intellect devourer out. The intellect devourer is also forced\
    \ out if the target regains its devoured brain by means of a [wish](/3-Mechanics/CLI/spells/wish.md).\
    \ By spending 5 feet of its movement, the intellect devourer can voluntarily leave\
    \ the body, teleporting to the nearest unoccupied space within 5 feet of it. The\
    \ body then dies, unless its brain is restored within 1 round."
  "name": "Body Thief"
"source":
- "MM"
- "WDH"
- "WDMM"
- "ERLW"
- "PaBTSO"
- "QftIS"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MM/Intellect%20Devourer.webp"
```
^statblock

```encounter-table
name: Intellect Devourer
creatures:
 - 1: Intellect Devourer
```

## Environment

underdark