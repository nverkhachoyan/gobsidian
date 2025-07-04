---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mtf
- monster/cr/14
- monster/environment/grassland
- monster/size/large
- monster/type/construct
aliases: ["Cadaver Collector"]
NoteIcon: monster
BestiaryType: construct
SourceType: Bestiary
BookSource: Mordenkainen's Tome of Foes p. 122
---
# [Cadaver Collector](3-Mechanics\CLI\bestiary\construct/cadaver-collector-mtf.md)
*Source: Mordenkainen's Tome of Foes p. 122*  

> [!quote]- A quote from Mordenkainen  
> 
> Cadaver collectors cannot summon specters of those they have not slain, nor will slain specters return to one. Starve it of victims, and then do not attack until you can be sure you'll destroy it. Meteor swarm works well.

## Cadaver Collector

The ancient war machines known as cadaver collectors lumber aimlessly across the blasted plains of Acheron until they are called upon by a necromancer, hobgoblin general, or other evil warlord to bolster the ranks of a conquering army. These fearsome constructs obey their summoners until being dismissed back to Acheron, but if a summoner comes to a bad end, a cadaver collector might wander the Material Plane for centuries, collecting corpses while searching for a way to return home.

## Sweeping the Dead

Cadaver collectors respond to a summons from a mortal only when they are called to the scene of a great battle—either where one is in progress, where one is imminent, or where one once took place. They encase themselves in the armor and weapons of fallen warriors and impale the corpses of those warriors on the lances and other weapons embedded in their salvaged armor.

## Conjured Berserkers

Corpses that accumulate on the construct's shell aren't just grisly battle trophies. A cadaver collector can summon the spirits of these cadavers to join battle with its enemies and to paralyze more creatures for eventual impalement. Although these specters are individually weak, a cadaver collector can call up an almost endless supply of them, if given time.

## Constructed Nature

A cadaver collector doesn't require air, food, drink, or sleep.

```statblock
"name": "Cadaver Collector (MTF)"
"size": "Large"
"type": "construct"
"alignment": "Lawful Evil"
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
- "desc": "The cadaver collector has advantage on saving throws against spells and\
    \ other magical effects."
  "name": "Magic Resistance"
- "desc": "As a bonus action, the cadaver collector calls up the enslaved spirits\
    \ of those it has slain; dice: 1d6|avg|noform (1d6) [specters](/3-Mechanics/CLI/bestiary/undead/specter.md)\
    \ (without Sunlight Sensitivity) arise in unoccupied spaces within 15 feet of\
    \ the cadaver collector. The specters act right after the cadaver collector on\
    \ the same initiative count and fight until they're destroyed. They disappear\
    \ when the cadaver collector is destroyed."
  "name": "Summon Specters (Recharges after a Short or Long Rest)"
"actions":
- "desc": "The cadaver collector makes two slam attacks."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+10 (+10) to hit, reach 5 ft., one\
    \ target. Hit: dice:3d8 + 5|text(18) (3d8 + 5) bludgeoning damage plus dice:3d10|text(16)\
    \ (3d10) necrotic damage."
  "name": "Slam"
- "desc": "The cadaver collector releases paralyzing gas in a 30-foot cone. Each creature\
    \ in that area must make a successful DC 18 Constitution saving throw or be [paralyzed](/3-Mechanics/CLI/rules/conditions.md#paralyzed)\
    \ for 1 minute. A [paralyzed](/3-Mechanics/CLI/rules/conditions.md#paralyzed)\
    \ creature repeats the saving throw at the end of each of its turns, ending the\
    \ effect on itself with a success."
  "name": "Paralyzing Breath (Recharge 5-6)"
"source":
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MTF/Cadaver%20Collector.webp"
```
^statblock

```encounter-table
name: Cadaver Collector
creatures:
 - 1: Cadaver Collector
```

## Environment

grassland