---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mm
- monster/cr/1
- monster/environment/underdark
- monster/environment/urban
- monster/size/medium
- monster/type/undead
aliases: ["Specter"]
NoteIcon: monster
BestiaryType: undead
SourceType: Bestiary
BookSource: Monster Manual p. 279, Dragon of Icespire Peak, Storm Lord's Wrath, Sleeping Dragon's Wake, Tasha's Cauldron of Everything. Available in the SRD.
---
# [Specter](3-Mechanics\CLI\bestiary\undead/specter.md)
*Source: Monster Manual p. 279, Dragon of Icespire Peak, Storm Lord's Wrath, Sleeping Dragon's Wake, Tasha's Cauldron of Everything. Available in the SRD.*  

A specter is the angry, unfettered spirit of a humanoid that has been prevented from passing to the afterlife. Specters no longer possess connections to who or what they were, yet are condemned to walk the world forever. Some are spawned when dark magic or the touch of a wraith rips a soul from a living body.

## Beyond Redemption

When a ghost's unfinished business is completed, it can rest at last. No such rest or redemption awaits a specter. It is doomed to the Material Plane, its only end the oblivion that comes with the destruction of its soul. Until then, it bears out its lonely life in forlorn places, carrying on forgotten through the ages of the world.

## Undying Hatred

Living creatures remind the specter that life is beyond its grasp. The mere sight of the living overwhelms a specter with sorrow and wrath, which can be abated only by destroying said life. A specter kills quickly and mercilessly, for only by depriving others of life can it gain the slightest satisfaction. However, no matter how many lives it extinguishes, a specter always succumbs to its hatred and sorrow.

### Dwellers in Darkness

Sunlight represents a source of life that no specter can ever hope to douse, and it pains them. When night falls, they leave their final resting places in search of living creatures to slay, knowing that few weapons can harm them in return.

At the first light of dawn, they retreat back into the darkness, where they remain until night falls again.

## Undead Nature

A specter doesn't require air, food, drink, or sleep.

```statblock
"name": "Specter"
"size": "Medium"
"type": "undead"
"alignment": "Chaotic Evil"
"ac": !!int "12"
"hp": !!int "22"
"hit_dice": "5d8"
"stats":
- !!int "1"
- !!int "14"
- !!int "11"
- !!int "10"
- !!int "10"
- !!int "11"
"speed": "0 ft., fly 50 ft. (hover)"
"damage_resistances": "acid; cold; fire; lightning; thunder; bludgeoning, piercing,\
  \ slashing from nonmagical attacks"
"damage_immunities": "necrotic, poison"
"condition_immunities": "[charmed](/3-Mechanics/CLI/rules/conditions.md#charmed),\
  \ [exhaustion](/3-Mechanics/CLI/rules/conditions.md#exhaustion), [grappled](/3-Mechanics/CLI/rules/conditions.md#grappled),\
  \ [paralyzed](/3-Mechanics/CLI/rules/conditions.md#paralyzed), [petrified](/3-Mechanics/CLI/rules/conditions.md#petrified),\
  \ [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned), [prone](/3-Mechanics/CLI/rules/conditions.md#prone),\
  \ [restrained](/3-Mechanics/CLI/rules/conditions.md#restrained), [unconscious](/3-Mechanics/CLI/rules/conditions.md#unconscious)"
"senses": "darkvision 60 ft., passive Perception 10"
"languages": "understands all languages it knew in life but can't speak"
"cr": "1"
"traits":
- "desc": "The specter can move through other creatures and objects as if they were\
    \ difficult terrain. It takes dice:1d10|text(5) (1d10) force damage if it\
    \ ends its turn inside an object."
  "name": "Incorporeal Movement"
- "desc": "While in sunlight, the specter has disadvantage on attack rolls, as well\
    \ as on Wisdom ([Perception](/3-Mechanics/CLI/rules/skills.md#Perception)) checks\
    \ that rely on sight."
  "name": "Sunlight Sensitivity"
"actions":
- "desc": "Melee Spell Attack: dice: d20+4 (+4) to hit, reach 5 ft., one creature.\
    \ Hit: dice:3d6|text(10) (3d6) necrotic damage. The target must succeed\
    \ on a DC 10 Constitution saving throw or its hit point maximum is reduced by\
    \ an amount equal to the damage taken. This reduction lasts until the creature\
    \ finishes a long rest. The target dies if this effect reduces its hit point maximum\
    \ to 0."
  "name": "Life Drain"
"source":
- "MM"
- "CoS"
- "HotDQ"
- "PotA"
- "RoT"
- "SKT"
- "TftYP"
- "ToA"
- "WDH"
- "WDMM"
- "GoS"
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
- "KftGV"
- "PSI"
- "PaBTSO"
- "AATM"
- "SatO"
- "ToFW"
- "BMT"
- "VEoR"
- "QftIS"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MM/Specter.webp"
```
^statblock

```encounter-table
name: Specter
creatures:
 - 1: Specter
```

## Environment

underdark, urban