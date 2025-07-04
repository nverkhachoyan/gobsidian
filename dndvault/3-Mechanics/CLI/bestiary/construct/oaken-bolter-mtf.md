---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mtf
- monster/cr/5
- monster/environment/forest
- monster/environment/grassland
- monster/environment/hill
- monster/environment/mountain
- monster/size/medium
- monster/type/construct
aliases: ["Oaken Bolter"]
NoteIcon: monster
BestiaryType: construct
SourceType: Bestiary
BookSource: Mordenkainen's Tome of Foes p. 126
---
# [Oaken Bolter](3-Mechanics\CLI\bestiary\construct/oaken-bolter-mtf.md)
*Source: Mordenkainen's Tome of Foes p. 126*  

> [!quote]- A quote from Mordenkainen  
> 
> Never depend on something built by a gnome. You can always rely on a gnome to take a good idea and make it impractical.

## Oaken Bolter

No ordinary ballista, an oaken bolter is a construct capable of striking at long distances. The bolts it launches can rend flesh, destroy armor, or drag enemies toward traps or melee-oriented clockworks—and at shorter ranges, burst with explosive force.

The gnomes' efforts to invent and tinker with magic and mechanical devices produce many failed constructs, but also result in genuine advances, such as clockworks. Since their discovery, the methods used to craft clockworks have passed from one community of gnomes to another and down the generations.

## Constructed Nature

A clockwork doesn't require air, food, drink, or sleep.

## Individual Designs

A gnome artisan values an individualized clockwork more highly than a perfectly functioning one that copies too much from another creation. For that reason, even clockworks that fit established designs, such as those described here, are seldom identical.

A clockwork can be customized by adding one of the following enhancements and one potential malfunction to its stat block. You can select randomly or choose a pair of modifications that fit the temperament of the clockwork's builder.

**Clockwork Enhancements**

`dice: [](oaken-bolter-mtf.md#^clockwork-enhancements)`

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

`dice: [](oaken-bolter-mtf.md#^clockwork-malfunctions)`

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
"name": "Oaken Bolter (MTF)"
"size": "Medium"
"type": "construct"
"alignment": "Unaligned"
"ac": !!int "16"
"ac_class": "natural armor"
"hp": !!int "58"
"hit_dice": "9d8 + 18"
"stats":
- !!int "12"
- !!int "18"
- !!int "15"
- !!int "3"
- !!int "10"
- !!int "1"
"speed": "30 ft."
"damage_immunities": "poison; bludgeoning, piercing, slashing from nonmagical attacks\
  \ that aren't adamantine"
"condition_immunities": "[charmed](/3-Mechanics/CLI/rules/conditions.md#charmed),\
  \ [exhaustion](/3-Mechanics/CLI/rules/conditions.md#exhaustion), [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened),\
  \ [paralyzed](/3-Mechanics/CLI/rules/conditions.md#paralyzed), [petrified](/3-Mechanics/CLI/rules/conditions.md#petrified),\
  \ [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)"
"senses": "darkvision 60 ft., passive Perception 10"
"languages": "understands one language of its creator but can't speak"
"cr": "5"
"traits":
- "desc": "The oaken bolter has advantage on saving throws against spells and other\
    \ magical effects."
  "name": "Magic Resistance"
"actions":
- "desc": "The oaken bolter makes two lancing bolt attacks or one lancing bolt attack\
    \ and one harpoon attack."
  "name": "Multiattack"
- "desc": "Melee or Ranged Weapon Attack: dice: d20+7 (+7) to hit, reach 5 ft.\
    \ or range 100/400 ft., one target. Hit: dice:2d10 + 4|text(15) (2d10 + 4)\
    \ piercing damage."
  "name": "Lancing Bolt"
- "desc": "Ranged Weapon Attack: dice: d20+7 (+7) to hit, range 50/200 ft.,\
    \ one target. Hit: dice:1d10 + 4|text(9) (1d10 + 4) piercing damage, and\
    \ the target is [grappled](/3-Mechanics/CLI/rules/conditions.md#grappled) (escape\
    \ DC 12). While [grappled](/3-Mechanics/CLI/rules/conditions.md#grappled) in this\
    \ way, a creature's speed isn't reduced, but it can move only in directions that\
    \ bring it closer to the oaken bolter. A creature takes dice:1d10|text(5) (1d10)\
    \ slashing damage if it escapes from the grapple or if it tries and fails. As\
    \ a bonus action, the oaken bolter can pull a creature [grappled](/3-Mechanics/CLI/rules/conditions.md#grappled)\
    \ by it 20 feet closer. The oaken bolter can grapple only one creature at a time."
  "name": "Harpoon"
- "desc": "The oaken bolter launches an explosive charge at a point within 120 feet.\
    \ Each creature within 20 feet of that point must make a DC 15 Dexterity saving\
    \ throw, taking dice:5d6|text(17) (5d6) fire damage on a failed save, or half\
    \ as much damage on a successful one."
  "name": "Explosive Bolt (Recharge 5-6)"
"source":
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MTF/Oaken%20Bolter.webp"
```
^statblock

```encounter-table
name: Oaken Bolter
creatures:
 - 1: Oaken Bolter
```

## Environment

forest, grassland, hill, mountain