---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/14
- monster/environment/grassland
- monster/size/large
- monster/type/construct
aliases: ["Cadaver Collector"]
NoteIcon: monster
BestiaryType: construct
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 68, Mordenkainen's Tome of Foes p. 122
---
# [Cadaver Collector](3-Mechanics\CLI\bestiary\construct/cadaver-collector-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 68, Mordenkainen's Tome of Foes p. 122*  

```statblock
"name": "Cadaver Collector (MPMM)"
"size": "Large"
"type": "construct"
"alignment": "Typically  Lawful Evil"
"ac": !!int "17"
"ac_class": "natural armor"
"hp": !!int "189"
"hit_dice": "18d10 + 90"
"stats":
- !!int "21"
- !!int "14"
- !!int "20"
- !!int "5"
- !!int "11"
- !!int "8"
"speed": "30 ft."
"damage_immunities": "necrotic; poison; psychic; bludgeoning, piercing, slashing from\
  \ nonmagical attacks that aren't adamantine"
"condition_immunities": "[charmed](/3-Mechanics/CLI/rules/conditions.md#charmed),\
  \ [exhaustion](/3-Mechanics/CLI/rules/conditions.md#exhaustion), [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened),\
  \ [paralyzed](/3-Mechanics/CLI/rules/conditions.md#paralyzed), [petrified](/3-Mechanics/CLI/rules/conditions.md#petrified),\
  \ [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)"
"senses": "darkvision 60 ft., passive Perception 10"
"languages": "understands all languages but can't speak"
"cr": "14"
"traits":
- "desc": "The collector has advantage on saving throws against spells and other magical\
    \ effects."
  "name": "Magic Resistance"
- "desc": "The collector doesn't require air, food, drink, or sleep."
  "name": "Unusual Nature"
"actions":
- "desc": "The collector makes two Slam attacks."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+10 (+10) to hit, reach 5 ft., one\
    \ target. Hit: dice:3d8 + 5|text(18) (3d8 + 5) bludgeoning damage plus dice:3d10|text(16)\
    \ (3d10) necrotic damage."
  "name": "Slam"
- "desc": "The collector releases paralyzing gas in a 30-foot cone. Each creature\
    \ in that area must make a successful DC 18 Constitution saving throw or be [paralyzed](/3-Mechanics/CLI/rules/conditions.md#paralyzed)\
    \ for 1 minute. A [paralyzed](/3-Mechanics/CLI/rules/conditions.md#paralyzed)\
    \ creature repeats the saving throw at the end of each of its turns, ending the\
    \ effect on itself with a success."
  "name": "Paralyzing Breath (Recharge 5-6)"
"bonus_actions":
- "desc": "The collector calls up the enslaved spirits of those it has slain; dice:\
    \ 1d4|avg|noform (1d4) [specters](/3-Mechanics/CLI/bestiary/undead/specter.md)\
    \ (without Sunlight Sensitivity) arise in unoccupied spaces within 15 feet of\
    \ it. The specters act right after the collector on the same initiative count\
    \ and fight until they're destroyed. They disappear when the collector is destroyed."
  "name": "Summon Specters (Recharges after a Short or Long Rest)"
"source":
- "MPMM"
- "MTF"
- "AATM"
- "VEoR"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Cadaver%20Collector.webp"
```
^statblock

```encounter-table
name: Cadaver Collector
creatures:
 - 1: Cadaver Collector
```

## Environment

grassland