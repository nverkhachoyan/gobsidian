---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mm
- monster/cr/0
- monster/size/tiny
- monster/type/undead
aliases: ["Crawling Claw"]
NoteIcon: monster
BestiaryType: undead
SourceType: Bestiary
BookSource: Monster Manual p. 44
---
# [Crawling Claw](3-Mechanics\CLI\bestiary\undead/crawling-claw.md)
*Source: Monster Manual p. 44*  

Crawling claws are the severed hands of murderers animated by dark magic so that they can go on killing. Wizards and warlocks of a dark bent use crawling claws as extra hands in their labors.

## Magical Origins

Through dark necromantic rituals, the life force of a murderer is bound to its severed hand, haunting and animating it. If a dead murderer's spirit already manifests as another undead creature, if the murderer is raised from death, or if the spirit has long passed on to another plane, the ritual fails.

The ritual invoked to create a crawling claw works best with a hand recently severed from a murderer. To this end, ritualists and their servants frequent public executions to gain possession of suitable hands, or make bargains with assassins and torturers.

## Creator's Control

A crawling claw can't be turned, nor can it be controlled by spells that control undead. These foul monsters are entirely bound to the will of their creator, which can concentrate on a claw in sight to mentally command its every action. If the crawling claw's creator doesn't command it, the claw follows its last command to the best of its ability.

Commands given to a crawling claw must be simple. A claw can't be tasked with finding and killing a particular person, because its limited senses and intelligence prevent it from tracking and picking out specific individuals. However, a command to kill all creatures in a particular locale works. A crawling claw can easily feel out the contours of keys and doorknobs, crawling from room to room on a blind killing spree.

## Malign Intelligence

A crawling claw possesses little of the intellect and memories of the individual of which it was once a living part. The hate, jealousy, or greed that drove that person to murder lingers on, however, amplified by the claw's torturous fragmented state. Left to its own devices, a crawling claw imitates and recreates the same murderous acts it committed in life.

## Living Claws

If a crawling claw is animated from the severed hand of a still-living murderer, the ritual binds the claw to the murderer's soul. The disembodied hand can then return to its former limb, its undead flesh knitting to the living arm from which it was severed.

Made whole again, the murderer acts as though the hand had never been severed and the ritual had never taken place. When the crawling claw separates again, the living body falls into a coma. Destroying the crawling claw while it is away from the body kills the murderer. However, killing the murderer has no effect on the crawling claw.

## Undead Nature

A crawling claw doesn't require air, food, drink, or sleep.

> [!quote]- A quote from Evangeliza Lavain, necromancer  
> 
> Makes you wonder what can be done with all those other murderer parts, doesn't it?


```statblock
"name": "Crawling Claw"
"size": "Tiny"
"type": "undead"
"alignment": "Neutral Evil"
"ac": !!int "12"
"hp": !!int "2"
"hit_dice": "1d4"
"stats":
- !!int "13"
- !!int "14"
- !!int "11"
- !!int "5"
- !!int "10"
- !!int "4"
"speed": "20 ft., climb 20 ft."
"damage_immunities": "poison"
"condition_immunities": "[charmed](/3-Mechanics/CLI/rules/conditions.md#charmed),\
  \ [exhaustion](/3-Mechanics/CLI/rules/conditions.md#exhaustion), [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)"
"senses": "blindsight 30 ft. (blind beyond this radius), passive Perception 10"
"languages": "understands Common but can't speak"
"cr": "0"
"traits":
- "desc": "The claw is immune to effects that turn undead."
  "name": "Turn Immunity"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+3 (+3) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d4 + 1|text(3) (1d4 + 1) bludgeoning or slashing damage (claw's\
    \ choice)."
  "name": "Claw"
"source":
- "MM"
- "CoS"
- "PotA"
- "ToA"
- "WDH"
- "BGDIA"
- "IDRotF"
- "CM"
- "JttRC"
- "LoX"
- "KftGV"
- "BMT"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MM/Crawling%20Claw.webp"
```
^statblock

```encounter-table
name: Crawling Claw
creatures:
 - 1: Crawling Claw
```