---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mtf
- monster/cr/16
- monster/environment/urban
- monster/size/large
- monster/type/construct
aliases: ["Steel Predator"]
NoteIcon: monster
BestiaryType: construct
SourceType: Bestiary
BookSource: Mordenkainen's Tome of Foes p. 239
---
# [Steel Predator](3-Mechanics\CLI\bestiary\construct/steel-predator-mtf.md)
*Source: Mordenkainen's Tome of Foes p. 239*  

## Steel Predator

A steel predator is a merciless machine with one purpose: to locate and kill its target regardless of distance and obstacles.

## Modron Engineering

Steel predators are created by a particular hexton modron, using a machine located in the city of Sigil. It wasn't always headquartered in the City of Doors, however. On its original home, the plane of Mechanus, the ingenious hexton was lauded for its invention—until it turned its creations against its superiors. Steel predators wreaked havoc across the modron hierarchy until the rogue hexton was trapped and exiled. Now it operates a shop in Sigil where, for a steep price, anyone can commission the manufacture of a steel predator.

## Assassins on Demand

To create a steel predator, the hexton's machine must be fed something that identifies the predator's target, such as a lock of hair, a well-worn glove, or a much-used weapon. The moment the newly manufactured steel predator emerges, it bounds away in search of its prey. It senses the location of its target across planar boundaries, but such detection is accurate only to within a thousand yards; to close the remaining distance, the steel predator locates its prey by sight and smell.

Once battle is joined, the predator ignores every other threat to attack its target, unless other creatures prevent it from reaching the target. In that case, it does what it must to fulfill its mission.

## Rogue Killers

If all goes according to plan, a steel predator slays its target and then voluntarily returns to Sigil, where it's broken down into parts that can be used in another steel predator. Battle damage can cause this instinct to fail, however, causing the steel predator to linger in the area, hunting and killing other creatures that resemble its target, that fit the target's general description, or that simply live nearby. Such rogues become the most dangerous of predators.

## Constructed Nature

A steel predator doesn't require air, food, drink, or sleep.

```statblock
"name": "Steel Predator (MTF)"
"size": "Large"
"type": "construct"
"alignment": "Lawful Evil"
"ac": !!int "20"
"ac_class": "natural armor"
"hp": !!int "207"
"hit_dice": "18d10 + 108"
"stats":
- !!int "24"
- !!int "17"
- !!int "22"
- !!int "4"
- !!int "14"
- !!int "6"
"speed": "40 ft."
"skillsaves":
  "Stealth": !!int "8"
  "Perception": !!int "7"
  "Survival": !!int "7"
"damage_resistances": "cold, lightning, necrotic, thunder"
"damage_immunities": "poison; psychic; bludgeoning, piercing, slashing from nonmagical\
  \ attacks"
"condition_immunities": "[charmed](/3-Mechanics/CLI/rules/conditions.md#charmed),\
  \ [exhaustion](/3-Mechanics/CLI/rules/conditions.md#exhaustion), [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened),\
  \ [paralyzed](/3-Mechanics/CLI/rules/conditions.md#paralyzed), [petrified](/3-Mechanics/CLI/rules/conditions.md#petrified),\
  \ [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned), [stunned](/3-Mechanics/CLI/rules/conditions.md#stunned)"
"senses": "blindsight 30 ft., darkvision 60 ft., passive Perception 17"
"languages": "understands Modron and the language of its owner but can't speak"
"cr": "16"
"traits":
- "desc": "The steel predator's innate spellcasting ability is Wisdom. The steel predator\
    \ can innately cast the following spells, requiring no components:\n\n3/day\
    \ each: [dimension door](/3-Mechanics/CLI/spells/dimension-door.md) (self only),\
    \ [plane shift](/3-Mechanics/CLI/spells/plane-shift.md) (self only)"
  "name": "Innate Spellcasting"
- "desc": "The steel predator has advantage on saving throws against spells and other\
    \ magical effects."
  "name": "Magic Resistance"
- "desc": "The steel predator's weapon attacks are magical."
  "name": "Magic Weapons"
"actions":
- "desc": "The steel predator makes three attacks: one with its bite and two with\
    \ its claw."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+12 (+12) to hit, reach 5 ft., one\
    \ target. Hit: dice:2d6 + 7|text(14) (2d6 + 7) piercing damage."
  "name": "Bite"
- "desc": "Melee Weapon Attack: dice: d20+12 (+12) to hit, reach 5 ft., one\
    \ target. Hit: dice:2d8 + 7|text(16) (2d8 + 7) slashing damage."
  "name": "Claw"
- "desc": "The steel predator emits a roar in a 60-foot cone. Each creature in that\
    \ area must make a DC 19 Constitution saving throw. On a failed save, a creature\
    \ takes dice:5d10|text(27) (5d10) thunder damage, drops everything it's holding,\
    \ and is [stunned](/3-Mechanics/CLI/rules/conditions.md#stunned) for 1 minute.\
    \ On a successful save, a creature takes half as much damage. The [stunned](/3-Mechanics/CLI/rules/conditions.md#stunned)\
    \ creature can repeat the saving throw at the end of each of its turns, ending\
    \ the effect on itself on a success."
  "name": "Stunning Roar (Recharge 5-6)"
"source":
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MTF/Steel%20Predator.webp"
```
^statblock

```encounter-table
name: Steel Predator
creatures:
 - 1: Steel Predator
```

## Environment

urban