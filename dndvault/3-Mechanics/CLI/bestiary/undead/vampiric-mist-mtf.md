---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mtf
- monster/cr/3
- monster/environment/arctic
- monster/environment/coastal
- monster/environment/forest
- monster/environment/grassland
- monster/environment/mountain
- monster/environment/swamp
- monster/environment/underdark
- monster/environment/urban
- monster/size/medium
- monster/type/undead
aliases: ["Vampiric Mist"]
NoteIcon: monster
BestiaryType: undead
SourceType: Bestiary
BookSource: Mordenkainen's Tome of Foes p. 246
---
# [Vampiric Mist](3-Mechanics\CLI\bestiary\undead/vampiric-mist-mtf.md)
*Source: Mordenkainen's Tome of Foes p. 246*  

## Vampiric Mist

In billowing clouds of fog lurk vampiric mists, the wretched remnants of vampires that were prevented from finding rest. Indistinguishable from the mists they lurk within, they strike unseen and undetected to bleed their victims dry.

## Former Vampires

Vampiric mists, sometimes called crimson mists, are all that remain of vampires who couldn't return to their burial places after being defeated or suffering some mishap. Denied the restorative power of these places, the vampires' bodies dissolve into mist. The transformation strips the intelligence and personality from them until only an unholy, insatiable thirst for blood remains.

## Blood Thief

Indistinguishable from fog aside from the charnel reek it exudes, a vampiric mist descends on a creature and causes the blood in its body to ooze through the creature's pores or spill out from its eyes, nose, and mouth. This blood wafts out from the victim like crimson smoke, which the mist then consumes. The feeding causes no pain or discomfort to the victim, so vampiric mists can feed on sleepers without waking them. The more a mist feeds, the redder it gets, such that it turns pink, then red, and finally a deep scarlet hue that rains blood droplets wherever it goes.

## Attracted to Blood

Like sharks in water, vampiric mists can scent blood from up to a mile away. Any injury, no matter how small, might catch their attention and draw them toward their victims. In battle, a mist focuses its attacks on injured targets, since open wounds are a more ready source of blood.

## Undead Nature

A vampiric mist doesn't require air or sleep.

```statblock
"name": "Vampiric Mist (MTF)"
"size": "Medium"
"type": "undead"
"alignment": "Chaotic Evil"
"ac": !!int "13"
"hp": !!int "30"
"hit_dice": "4d8 + 12"
"stats":
- !!int "6"
- !!int "16"
- !!int "16"
- !!int "6"
- !!int "12"
- !!int "7"
"speed": "0 ft., fly 30 ft. (hover)"
"saves":
  "Wisdom": !!int "3"
"damage_resistances": "acid; cold; lightning; necrotic; thunder; bludgeoning, piercing,\
  \ slashing from nonmagical attacks"
"damage_immunities": "poison"
"condition_immunities": "[charmed](/3-Mechanics/CLI/rules/conditions.md#charmed),\
  \ [exhaustion](/3-Mechanics/CLI/rules/conditions.md#exhaustion), [grappled](/3-Mechanics/CLI/rules/conditions.md#grappled),\
  \ [paralyzed](/3-Mechanics/CLI/rules/conditions.md#paralyzed), [petrified](/3-Mechanics/CLI/rules/conditions.md#petrified),\
  \ [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned), [prone](/3-Mechanics/CLI/rules/conditions.md#prone),\
  \ [restrained](/3-Mechanics/CLI/rules/conditions.md#restrained)"
"senses": "darkvision 60 ft., passive Perception 11"
"languages": ""
"cr": "3"
"traits":
- "desc": "The mist can sense the location of any creature within 60 feet of it, unless\
    \ that creature's type is construct or undead."
  "name": "Life Sense"
- "desc": "The mist can't enter a residence without an invitation from one of the\
    \ occupants."
  "name": "Forbiddance"
- "desc": "The mist can occupy another creature's space and vice versa. In addition,\
    \ if air can pass through a space, the mist can pass through it without squeezing.\
    \ Each foot of movement in water costs it 2 extra feet, rather than 1 extra foot.\
    \ The mist can't manipulate objects in any way that requires fingers or manual\
    \ dexterity."
  "name": "Misty Form"
- "desc": "The mist takes 10 radiant damage whenever it starts its turn in sunlight.\
    \ While in sunlight, the mist has disadvantage on attack rolls and ability checks."
  "name": "Sunlight Hypersensitivity"
"actions":
- "desc": "The mist touches one creature in its space. The target must succeed on\
    \ a DC 13 Constitution saving throw (undead and constructs automatically succeed),\
    \ or it takes dice:2d6 + 3|text(10) (2d6 + 3) necrotic damage, the mist regains\
    \ 10 hit points, and the target's hit point maximum is reduced by an amount equal\
    \ to the necrotic damage taken. This reduction lasts until the target finishes\
    \ a long rest. The target dies if its hit point maximum is reduced to 0."
  "name": "Life Drain"
"source":
- "MTF"
- "TftYP"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MTF/Vampiric%20Mist.webp"
```
^statblock

```encounter-table
name: Vampiric Mist
creatures:
 - 1: Vampiric Mist
```

## Environment

arctic, coastal, forest, grassland, mountain, swamp, underdark, urban