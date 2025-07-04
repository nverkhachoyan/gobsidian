---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mtf
- monster/cr/20
- monster/environment/arctic
- monster/environment/desert
- monster/environment/swamp
- monster/environment/underdark
- monster/size/huge
- monster/type/undead
aliases: ["Nightwalker"]
NoteIcon: monster
BestiaryType: undead
SourceType: Bestiary
BookSource: Mordenkainen's Tome of Foes p. 216
---
# [Nightwalker](3-Mechanics\CLI\bestiary\undead/nightwalker-mtf.md)
*Source: Mordenkainen's Tome of Foes p. 216*  

## Nightwalker

The Negative Plane is a place of darkness and death, anathema to all living things. Yet there are those who would tap into its fell power, to use its energy for sinister ends. Most often, when such individuals approach the midnight realm, they find they are unequal to the task. Those not destroyed outright are sometimes drawn inside the plane and replaced by nightwalkers, terrifying undead creatures that devour all life they encounter.

## Mighty Spawn

One can reach the Negative Plane from the Shadowfell, much in the same way that it is possible to step from the Material Plane into the Shadowfell in a place where the barrier between the planes is thin.

Stepping into the Negative Plane is tantamount to suicide, since the plane sucks the life and soul from such audacious creatures and annihilates them at once. Those few who survive the effort do so by sheer luck or by harnessing some rare form of magic that protects them against the hostile atmosphere. They soon discover, however, that they can't leave as easily as they arrived. For each creature that enters the plane, a nightwalker is released to take its place. In order for a trapped creature to escape, the released nightwalker must be lured back to the Negative Plane by offerings of life for it to devour. If the nightwalker is destroyed, the trapped creature has no hope of escape.

## Beings of Anti-Life

One can discern the nature of creatures trapped in the Negative Plane from the sites that nightwalkers frequent. Generally, a nightwalker on the Material Plane is attracted to elements of the world associated with the creature responsible for its creation. Such interest doesn't indicate a willingness to engage with the world; nightwalkers exist to make life extinct and never to serve living things.

## Undead Nature

A nightwalker doesn't require air, food, drink, or sleep.

```statblock
"name": "Nightwalker (MTF)"
"size": "Huge"
"type": "undead"
"alignment": "Chaotic Evil"
"ac": !!int "14"
"hp": !!int "297"
"hit_dice": "22d12 + 154"
"stats":
- !!int "22"
- !!int "19"
- !!int "24"
- !!int "6"
- !!int "9"
- !!int "8"
"speed": "40 ft., fly 40 ft."
"saves":
  "Constitution": !!int "13"
"damage_resistances": "acid; cold; fire; lightning; thunder; bludgeoning, piercing,\
  \ slashing from nonmagical attacks"
"damage_immunities": "necrotic, poison"
"condition_immunities": "[exhaustion](/3-Mechanics/CLI/rules/conditions.md#exhaustion),\
  \ [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened), [grappled](/3-Mechanics/CLI/rules/conditions.md#grappled),\
  \ [paralyzed](/3-Mechanics/CLI/rules/conditions.md#paralyzed), [petrified](/3-Mechanics/CLI/rules/conditions.md#petrified),\
  \ [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned), [prone](/3-Mechanics/CLI/rules/conditions.md#prone),\
  \ [restrained](/3-Mechanics/CLI/rules/conditions.md#restrained)"
"senses": "darkvision 120 ft., passive Perception 9"
"languages": ""
"cr": "20"
"traits":
- "desc": "Any creature that starts its turn within 30 feet of the nightwalker must\
    \ succeed on a DC 21 Constitution saving throw or take dice:4d6|text(14) (4d6)\
    \ necrotic damage and grant the nightwalker advantage on attack rolls against\
    \ it until the start of the creature's next turn. Undead are immune to this aura."
  "name": "Annihilating Aura"
- "desc": "A creature reduced to 0 hit points from damage dealt by the nightwalker\
    \ dies and can't be revived by any means short of a [wish](/3-Mechanics/CLI/spells/wish.md)\
    \ spell."
  "name": "Life Eater"
"actions":
- "desc": "The nightwalker uses Enervating Focus twice, or it uses Enervating Focus\
    \ and Finger of Doom, if available."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+12 (+12) to hit, reach 15 ft., one\
    \ target. Hit: dice:5d8 + 6|text(28) (5d8 + 6) necrotic damage. The target\
    \ must succeed on a DC 21 Constitution saving throw or its hit point maximum is\
    \ reduced by an amount equal to the necrotic damage taken. This reduction lasts\
    \ until the target finishes a long rest."
  "name": "Enervating Focus"
- "desc": "The nightwalker points at one creature it can see within 300 feet of it.\
    \ The target must succeed on a DC 21 Wisdom saving throw or take dice:4d12|text(26)\
    \ (4d12) necrotic damage and become [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened)\
    \ until the end of the nightwalker's next turn. While [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened)\
    \ in this way, the creature is also [paralyzed](/3-Mechanics/CLI/rules/conditions.md#paralyzed).\
    \ If a target's saving throw is successful, the target is immune to the nightwalker's\
    \ Finger of Doom for the next 24 hours."
  "name": "Finger of Doom (Recharge 6)"
"source":
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MTF/Nightwalker.webp"
```
^statblock

```encounter-table
name: Nightwalker
creatures:
 - 1: Nightwalker
```

## Environment

arctic, desert, swamp, underdark