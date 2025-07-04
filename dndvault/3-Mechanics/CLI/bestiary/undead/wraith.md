---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mm
- monster/cr/5
- monster/environment/underdark
- monster/size/medium
- monster/type/undead
aliases: ["Wraith"]
NoteIcon: monster
BestiaryType: undead
SourceType: Bestiary
BookSource: Monster Manual p. 302, Divine Contention, Dragon of Icespire Peak, Storm Lord's Wrath, Sleeping Dragon's Wake, Tasha's Cauldron of Everything. Available in the SRD.
---
# [Wraith](3-Mechanics\CLI\bestiary\undead/wraith.md)
*Source: Monster Manual p. 302, Divine Contention, Dragon of Icespire Peak, Storm Lord's Wrath, Sleeping Dragon's Wake, Tasha's Cauldron of Everything. Available in the SRD.*  

A wraith is malice incarnate, concentrated into an incorporeal form that seeks to quench all life. The creature is suffused with negative energy, and its mere passage through the world leaves nearby plants blackened and withered. Animals flee from its presence. Even small fires can be extinguished by the sucking oblivion of the wraith's horrifying existence.

## Vile Oblivion

When a mortal humanoid lives a debased life or enters into a fiendish pact, it consigns its soul to eternal damnation in the Lower Planes. However, sometimes the soul becomes so suffused with negative energy that it collapses in on itself and ceases to exist the instant before it can shuffle off to some horrible afterlife. When this occurs, the spirit becomes a soulless wraith-a malevolent void trapped on the plane where it died. Almost nothing of the wraith's former existence is preserved; in this new form, it exists only to annihilate other life.

## Bereft of Body

A wraith can move through solid creatures and objects as easily as a mortal creature moves through fog.

A wraith might retain a few memories of its mortal life as shadowy echoes. However, even the strongest events and emotions become little more than faint impressions, fleeting as half-remembered dreams. A wraith might pause to stare at something that fascinated it in life, or it might curb its wrath in acknowledgment of a past friendship. Such moments come rarely, however, because most wraiths despise what they were as a reminder of what they have become.

## Undead Commanders

A wraith can make an undead servant from the spirit of a humanoid creature that has recently suffered a violent death. Such a fragment of woe becomes a specter, spiteful of all that lives.

Wraiths sometimes rule the legions of the dead, plotting the doom of living creatures. When they emerge from their tombs to do battle, life and hope shrivel before them. Even if a wraith's armies are forced to retreat, the lands its forces occupied are so blasted and withered that those who live there often starve and die.

## Undead Nature

A wraith doesn't require air, food, drink, or sleep.

```statblock
"name": "Wraith"
"size": "Medium"
"type": "undead"
"alignment": "Neutral Evil"
"ac": !!int "13"
"hp": !!int "67"
"hit_dice": "9d8 + 27"
"stats":
- !!int "6"
- !!int "16"
- !!int "16"
- !!int "12"
- !!int "14"
- !!int "15"
"speed": "0 ft., fly 60 ft. (hover)"
"damage_resistances": "acid; cold; fire; lightning; thunder; bludgeoning, piercing,\
  \ slashing from nonmagical attacks that aren't silvered"
"damage_immunities": "necrotic, poison"
"condition_immunities": "[charmed](/3-Mechanics/CLI/rules/conditions.md#charmed),\
  \ [exhaustion](/3-Mechanics/CLI/rules/conditions.md#exhaustion), [grappled](/3-Mechanics/CLI/rules/conditions.md#grappled),\
  \ [paralyzed](/3-Mechanics/CLI/rules/conditions.md#paralyzed), [petrified](/3-Mechanics/CLI/rules/conditions.md#petrified),\
  \ [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned), [prone](/3-Mechanics/CLI/rules/conditions.md#prone),\
  \ [restrained](/3-Mechanics/CLI/rules/conditions.md#restrained)"
"senses": "darkvision 60 ft., passive Perception 12"
"languages": "the languages it knew in life"
"cr": "5"
"traits":
- "desc": "The wraith can move through other creatures and objects as if they were\
    \ difficult terrain. It takes dice:1d10|text(5) (1d10) force damage if it\
    \ ends its turn inside an object."
  "name": "Incorporeal Movement"
- "desc": "While in sunlight, the wraith has disadvantage on attack rolls, as well\
    \ as on Wisdom ([Perception](/3-Mechanics/CLI/rules/skills.md#Perception)) checks\
    \ that rely on sight."
  "name": "Sunlight Sensitivity"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+6 (+6) to hit, reach 5 ft., one creature.\
    \ Hit: dice:4d8 + 3|text(21) (4d8 + 3) necrotic damage. The target must\
    \ succeed on a DC 14 Constitution saving throw or its hit point maximum is reduced\
    \ by an amount equal to the damage taken. This reduction lasts until the target\
    \ finishes a long rest. The target dies if this effect reduces its hit point maximum\
    \ to 0."
  "name": "Life Drain"
- "desc": "The wraith targets a humanoid within 10 feet of it that has been dead for\
    \ no longer than 1 minute and died violently. The target's spirit rises as a [specter](/3-Mechanics/CLI/bestiary/undead/specter.md)\
    \ in the space of its corpse or in the nearest unoccupied space. The [specter](/3-Mechanics/CLI/bestiary/undead/specter.md)\
    \ is under the wraith's control. The wraith can have no more than seven specters\
    \ under its control at one time."
  "name": "Create Specter"
"source":
- "MM"
- "CoS"
- "RoT"
- "TftYP"
- "ToA"
- "WDMM"
- "GoS"
- "DC"
- "DIP"
- "SLW"
- "SDW"
- "BGDIA"
- "EGW"
- "MOT"
- "IDRotF"
- "TCE"
- "CM"
- "JttRC"
- "DSotDQ"
- "KftGV"
- "PSI"
- "PaBTSO"
- "AATM"
- "SatO"
- "BMT"
- "GHLoE"
- "DoDk"
- "VEoR"
- "QftIS"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MM/Wraith.webp"
```
^statblock

```encounter-table
name: Wraith
creatures:
 - 1: Wraith
```

## Environment

underdark