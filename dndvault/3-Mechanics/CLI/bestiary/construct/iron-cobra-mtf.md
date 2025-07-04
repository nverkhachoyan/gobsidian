---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mtf
- monster/cr/4
- monster/environment/forest
- monster/environment/grassland
- monster/environment/hill
- monster/environment/mountain
- monster/size/medium
- monster/type/construct
aliases: ["Iron Cobra"]
NoteIcon: monster
BestiaryType: construct
SourceType: Bestiary
BookSource: Mordenkainen's Tome of Foes p. 125
---
# [Iron Cobra](3-Mechanics\CLI\bestiary\construct/iron-cobra-mtf.md)
*Source: Mordenkainen's Tome of Foes p. 125*  

> [!quote]- A quote from Mordenkainen  
> 
> Never depend on something built by a gnome. You can always rely on a gnome to take a good idea and make it impractical.

## Iron Cobra

An iron cobra is exactly what its name implies: a metal snake with a poisonous bite. In addition to standard poisons, gnomes load this clockwork with alchemical concoctions that can paralyze creatures and cloud their minds with paranoia.

The gnomes' efforts to invent and tinker with magic and mechanical devices produce many failed constructs, but also result in genuine advances, such as clockworks. Since their discovery, the methods used to craft clockworks have passed from one community of gnomes to another and down the generations.

## Constructed Nature

A clockwork doesn't require air, food, drink, or sleep.

## Individual Designs

A gnome artisan values an individualized clockwork more highly than a perfectly functioning one that copies too much from another creation. For that reason, even clockworks that fit established designs, such as those described here, are seldom identical.

A clockwork can be customized by adding one of the following enhancements and one potential malfunction to its stat block. You can select randomly or choose a pair of modifications that fit the temperament of the clockwork's builder.

**Clockwork Enhancements**

`dice: [](iron-cobra-mtf.md#^clockwork-enhancements)`

| dice: d10 | Enhancement |
|-----------|-------------|
| 1 | **Camouflaged.** The clockwork gains proficiency in [Stealth](/3-Mechanics/CLI/rules/skills.md#Stealth) if it doesn't already have it. While motionless, it is indistinguishable from a stopped machine. |
| 2 | **Sensors.** The range of the clockwork's [darkvision](/3-Mechanics/CLI/rules/senses.md#darkvision) becomes 120 feet, unless it is higher, and it gains proficiency in [Perception](/3-Mechanics/CLI/rules/skills.md#Perception) if it doesn't already have it. |
| 3 | **Improved Armor.** The clockwork's AC increases by 2. |
| 4 | **Increased Speed.** The clockwork's speed increases by 10 feet. |
| 5 | **Reinforced Construction.** The clockwork has resistance to force, lightning, and thunder damage. |
| 6 | **Self-Repairing.** If the clockwork starts its turn with at least 1 hit point, it regains 5 hit points. If it takes lightning damage, this ability doesn't function at the start of its next turn. |
| 7 | **Sturdy Frame.** The clockwork's hit point maximum increases by an amount equal to its number of Hit Dice. |
| 8 | **Suction.** The clockwork gains a climbing speed of 30 feet. |
| 9 | **Vocal Resonator.** The clockwork gains the ability to speak rudimentary Common or Gnomish (creator's choice). |
| 10 | **Water Propulsion.** The clockwork gains a swimming speed of 30 feet. |
^clockwork-enhancements

**Clockwork Malfunctions**

`dice: [](iron-cobra-mtf.md#^clockwork-malfunctions)`

| dice: d10 | Malfunction |
|-----------|-------------|
| 1 | **Faulty Sensors.** Roll a `dice: d6\|avg\|noform` (`d6`) at the start of the clockwork's turn. If you roll a 1, the clockwork is [blinded](/3-Mechanics/CLI/rules/conditions.md#blinded) until the end of its turn. |
| 2 | **Flawed Targeting.** Roll a `dice: d6\|avg\|noform` (`d6`) at the start of the clockwork's turn. If you roll a 1, the clockwork makes attack rolls with disadvantage until the end of its turn. |
| 3 | **Ground Fault.** The clockwork has vulnerability to lightning damage. |
| 4 | **Imprinting Loop.** Roll a `dice: d6\|avg\|noform` (`d6`) at the start of the clockwork's turn. If you roll a 1, the clockwork mistakes one creature it can see within 30 feet for its creator. The clockwork won't willingly harm that creature for 1 minute or until that creature attacks it or deals damage to it. |
| 5 | **Leaking Lubricant.** Roll a `dice: d6\|avg\|noform` (`d6`) at the start of the clockwork's turn. If you roll a 1, the clockwork gains 1 level of [exhaustion](/3-Mechanics/CLI/rules/conditions.md#exhaustion) that it isn't immune to. |
| 6 | **Limited Steering.** The clockwork must move in a straight line. It can turn up to 90 degrees before moving and again at the midpoint of its movement. It can rotate freely if it doesn't use any of its speeds on its turn. |
| 7 | **Overactive Sense of Self-Preservation.** If the clockwork has half its hit points or fewer at the start of its turn in combat, roll a `dice: d6\|avg\|noform` (`d6`). If you roll a 1, it retreats from combat. If retreat isn't possible, it continues fighting. |
| 8 | **Overheats.** Roll a `dice: d6\|avg\|noform` (`d6`) at the start of the clockwork's turn. If you roll a 1, the clockwork is [incapacitated](/3-Mechanics/CLI/rules/conditions.md#incapacitated) until the end of its turn. |
| 9 | **Rusty Gears.** The clockwork has disadvantage on initiative rolls, and its speed decreases by 10 feet. |
| 10 | **Weak Armor.** The clockwork isn't immune to bludgeoning, piercing, and slashing damage from nonmagical attacks that aren't adamantine. |
^clockwork-malfunctions

```statblock
"name": "Iron Cobra (MTF)"
"size": "Medium"
"type": "construct"
"alignment": "Unaligned"
"ac": !!int "13"
"hp": !!int "45"
"hit_dice": "7d8 + 14"
"stats":
- !!int "12"
- !!int "16"
- !!int "14"
- !!int "3"
- !!int "10"
- !!int "1"
"speed": "30 ft."
"skillsaves":
  "Stealth": !!int "7"
"damage_immunities": "poison; bludgeoning, piercing, slashing from nonmagical attacks\
  \ that aren't adamantine"
"condition_immunities": "[charmed](/3-Mechanics/CLI/rules/conditions.md#charmed),\
  \ [exhaustion](/3-Mechanics/CLI/rules/conditions.md#exhaustion), [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened),\
  \ [paralyzed](/3-Mechanics/CLI/rules/conditions.md#paralyzed), [petrified](/3-Mechanics/CLI/rules/conditions.md#petrified),\
  \ [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)"
"senses": "darkvision 60 ft., passive Perception 10"
"languages": "understands one language of its creator but can't speak"
"cr": "4"
"traits":
- "desc": "The iron cobra has advantage on saving throws against spells and other\
    \ magical effects."
  "name": "Magic Resistance"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+5 (+5) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d6 + 3|text(6) (1d6 + 3) piercing damage. If the target is\
    \ a creature, it must succeed on a DC 13 Constitution saving throw or suffer one\
    \ random poison effect:\n\n1. Poison Damage: The target takes dice:3d8|text(13)\
    \ (3d8) poison damage.\n\n2. Confusion: On its next turn, the target must use\
    \ its action to make one weapon attack against a random creature it can see within\
    \ 30 feet of it, using whatever weapon it has in hand and moving beforehand if\
    \ necessary to get in range. If it's holding no weapon, it makes an unarmed strike.\
    \ If no creature is visible within 30 feet, it takes the Dash action, moving toward\
    \ the nearest creature.\n\n3. Paralysis: The target is [paralyzed](/3-Mechanics/CLI/rules/conditions.md#paralyzed)\
    \ until the end of its next turn."
  "name": "Bite"
"source":
- "MTF"
- "GoS"
- "RtG"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MTF/Iron%20Cobra.webp"
```
^statblock

```encounter-table
name: Iron Cobra
creatures:
 - 1: Iron Cobra
```

## Environment

forest, grassland, hill, mountain