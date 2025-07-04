---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/tftyp
- monster/cr/5
- monster/size/medium
- monster/type/undead
aliases: ["Greater Zombie"]
NoteIcon: monster
BestiaryType: undead
SourceType: Bestiary
BookSource: Tales from the Yawning Portal p. 237, Divine Contention, Sleeping Dragon's Wake
---
# [Greater Zombie](3-Mechanics\CLI\bestiary\undead/greater-zombie-tftyp.md)
*Source: Tales from the Yawning Portal p. 237, Divine Contention, Sleeping Dragon's Wake*  

```statblock
"name": "Greater Zombie (TftYP)"
"size": "Medium"
"type": "undead"
"alignment": "Neutral Evil"
"ac": !!int "15"
"ac_class": "natural armor"
"hp": !!int "97"
"hit_dice": "13d8 + 39"
"stats":
- !!int "18"
- !!int "10"
- !!int "17"
- !!int "4"
- !!int "6"
- !!int "6"
"speed": "30 ft."
"saves":
  "Wisdom": !!int "1"
"damage_resistances": "cold, necrotic"
"damage_immunities": "poison"
"condition_immunities": "[charmed](/3-Mechanics/CLI/rules/conditions.md#charmed),\
  \ [exhaustion](/3-Mechanics/CLI/rules/conditions.md#exhaustion), [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened),\
  \ [paralyzed](/3-Mechanics/CLI/rules/conditions.md#paralyzed), [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)"
"senses": "darkvision 60 ft., passive Perception 8"
"languages": "understands the languages it knew in life but can't speak"
"cr": "5"
"traits":
- "desc": "The zombie has advantage on saving throws against any effect that turns\
    \ undead."
  "name": "Turn Resistance"
- "desc": "If damage reduces the zombie to 0 hit points, it must make a Constitution\
    \ saving throw with a DC of 5 + the damage taken, unless the damage is radiant\
    \ or from a critical hit. On a success, the zombie drops to 1 hit point instead."
  "name": "Undead Fortitude"
"actions":
- "desc": "The zombie makes two melee attacks."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+7 (+7) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d6 + 4|text(7) (1d6 + 4) bludgeoning damage and dice:2d6|text(7)\
    \ (2d6) necrotic damage."
  "name": "Empowered Slam"
"source":
- "TftYP"
- "DC"
- "SDW"
- "IMR"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/TftYP/Greater%20Zombie.webp"
```
^statblock

```encounter-table
name: Greater Zombie
creatures:
 - 1: Greater Zombie
```