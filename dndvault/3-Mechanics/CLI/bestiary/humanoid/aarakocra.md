---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mm
- monster/cr/1-4
- monster/environment/mountain
- monster/size/medium
- monster/type/humanoid/aarakocra
aliases: ["Aarakocra"]
NoteIcon: monster
BestiaryType: humanoid (aarakocra)
SourceType: Bestiary
BookSource: Monster Manual p. 12
---
# [Aarakocra](3-Mechanics\CLI\bestiary\humanoid/aarakocra.md)
*Source: Monster Manual p. 12*  

Aarakocra range the Howling Gyre, an endless storm of mighty winds and lashing rains that surrounds the tranquil realm of Aaqa in the Elemental Plane of Air. Making aerial patrols, these birdlike humanoids guard the windy borders of their home against invaders from the Elemental Plane of Earth, such as gargoyles, their sworn enemies.

## Enemies of Elemental Evil

In service to the Wind Dukes of Aaqa, aarakocra scout the planes in search of temples of Elemental Evil. They spy on malign elemental creatures and then either take the fight to those creatures or report back to the Wind Dukes.

On the Material Plane, aarakocra create aeries atop the highest mountains, especially peaks near portals to the Elemental Plane of Air. From such heights, aarakocra watch for signs of elemental incursions, as well as for nascent threats to their home plane. Aarakocra prefer to live their lives like the wind-unburdened and ever moving-yet they watch over a region for years if that's what it takes to guard against the incursions of Elemental Evil.

Aarakocra have no concept of political borders or property ownership, and the value of gems, gold, and other precious materials means little to aarakocra. In their eyes, a creature should use what is necessary and then cast what is left on the wind for others to use.

### Search for the Seven Shards

The Wind Dukes of Aaqa come from a race of elemental beings called the vaati, which once ruled many worlds. A creature known as the Queen of Chaos arose and initiated an interplanar war against vaati rule. To combat the threat, seven vaati heroes combined their powers to create the mighty *Rod of Law*. In a battle against the queen's greatest general, Mishka the Wolf Spider, a vaati killed Mishka by thrusting the rod into him like a spear. The rod shattered into seven shards that scattered across the multiverse. Aaracokra seek signs of the pieces' locations in order to rebuild what is now known as the *Rod of Seven Parts*.

> [!note] Summoning Air Elementals
> 
> Five aarakocra within 30 feet of each other can magically summon an air elemental. Each of the five must use its action and movement on three consecutive turns to perform an aerial dance and must maintain concentration while doing so (as if concentrating on a spell). When all five have finished their third turn of the dance, the elemental appears in an unoccupied space within 60 feet of them. It is friendly toward them and obeys their spoken commands. It remains for 1 hour, until it or all its summoners die, or until any of its summoners dismisses it as a bonus action. A summoner can't perform the dance again until it finishes a short rest. When the elemental returns to the Elemental Plane of Air, any aarakocra within 5 feet of it can return with it.
^summoning-air-elementals

```statblock
"name": "Aarakocra"
"size": "Medium"
"type": "humanoid"
"subtype": "aarakocra"
"alignment": "Neutral Good"
"ac": !!int "12"
"hp": !!int "13"
"hit_dice": "3d8"
"stats":
- !!int "10"
- !!int "14"
- !!int "10"
- !!int "11"
- !!int "12"
- !!int "11"
"speed": "20 ft., fly 50 ft."
"skillsaves":
  "Perception": !!int "5"
"senses": "passive Perception 15"
"languages": "Auran, Aarakocra"
"cr": "1/4"
"traits":
- "desc": "If the aarakocra is flying and dives at least 30 feet straight toward a\
    \ target and then hits it with a melee weapon attack, the attack deals an extra\
    \ dice:1d6|text(3) (1d6) damage to the target."
  "name": "Dive Attack"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+4 (+4) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d4 + 2|text(4) (1d4 + 2) slashing damage."
  "name": "Talon"
- "desc": "Melee or Ranged Weapon Attack: dice: d20+4 (+4) to hit, reach 5 ft.\
    \ or range 30/120 ft., one target. Hit: dice:1d6 + 2|text(5) (1d6 + 2) piercing\
    \ damage."
  "name": "Javelin"
- "desc": "Five aarakocra within 30 feet of each other can magically summon an [air\
    \ elemental](/3-Mechanics/CLI/bestiary/elemental/air-elemental.md). Each of the\
    \ five must use its action and movement on three consecutive turns to perform\
    \ an aerial dance and must maintain [concentration](/3-Mechanics/CLI/rules/conditions.md#concentration)\
    \ while doing so (as if concentrating on a spell). When all five have finished\
    \ their third turn of the dance, the elemental appears in an unoccupied space\
    \ within 60 feet of them. It is friendly toward them and obeys their spoken commands.\
    \ It remains for 1 hour, until it or all its summoners die, or until any of its\
    \ summoners dismisses it as a bonus action. A summoner can't perform the dance\
    \ again until it finishes a short rest. When the elemental returns to the Elemental\
    \ Plane of Air, any aarakocra within 5 feet of it can return with it."
  "name": "Summon Air Elemental"
"source":
- "MM"
- "PotA"
- "SKT"
- "ToA"
- "EGW"
- "IDRotF"
- "CRCotN"
- "LoX"
- "QftIS"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MM/Aarakocra.webp"
```
^statblock

```encounter-table
name: Aarakocra
creatures:
 - 1: Aarakocra
```

## Environment

mountain