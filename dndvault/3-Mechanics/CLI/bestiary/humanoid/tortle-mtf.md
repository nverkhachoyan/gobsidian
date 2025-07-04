---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mtf
- monster/cr/1-4
- monster/environment/coastal
- monster/size/medium
- monster/type/humanoid/tortle
aliases: ["Tortle"]
NoteIcon: monster
BestiaryType: humanoid (tortle)
SourceType: Bestiary
BookSource: Mordenkainen's Tome of Foes p. 242
---
# [Tortle](3-Mechanics\CLI\bestiary\humanoid/tortle-mtf.md)
*Source: Mordenkainen's Tome of Foes p. 242*  

## Tortles

Tortles are omnivorous, turtle-like humanoids with shells that cover most of their bodies. Tortles have a saying: "We wear our homes on our backs." Consequently, tortles feel little need to stay put for long.

An adult tortle stands about 6 feet tall and weighs between 450 and 500 pounds. Males and females are nearly identical in size and appearance.

## Temporary Towns

A tortle settlement is primarily used as a kind of moot, where tortles can socialize with one another and trade with strangers. Tortles don't regard these settlements as places worth defending with their lives, and they abandon a settlement when it no longer serves their needs.

## A Life of Wandering

Most tortles like to see how other creatures live and discover new customs. The urge to procreate doesn't kick in until the end of a tortle's life, and a tortle can spend decades away from its native land without feeling homesick.

```statblock
"name": "Tortle (MTF)"
"size": "Medium"
"type": "humanoid"
"subtype": "tortle"
"alignment": "Lawful Good"
"ac": !!int "17"
"ac_class": "natural armor"
"hp": !!int "22"
"hit_dice": "4d8 + 4"
"stats":
- !!int "15"
- !!int "10"
- !!int "12"
- !!int "11"
- !!int "13"
- !!int "12"
"speed": "30 ft."
"skillsaves":
  "Athletics": !!int "4"
  "Survival": !!int "3"
"senses": "passive Perception 11"
"languages": "Aquan, Common"
"cr": "1/4"
"traits":
- "desc": "The tortle can hold its breath for 1 hour."
  "name": "Hold Breath"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+4 (+4) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d4 + 2|text(4) (1d4 + 2) slashing damage."
  "name": "Claws"
- "desc": "Melee Weapon Attack: dice: d20+4 (+4) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d6 + 2|text(5) (1d6 + 2) bludgeoning damage, or dice:1d8 +\
    \ 2|text(6) (1d8 + 2) bludgeoning damage if used with two hands."
  "name": "Quarterstaff"
- "desc": "Ranged Weapon Attack: dice: d20+2 (+2) to hit, range 80/320 ft.,\
    \ one target. Hit: dice:1d8|text(4) (1d8) piercing damage."
  "name": "Light Crossbow"
- "desc": "The tortle withdraws into its shell. Until it emerges, it gains a +4 bonus\
    \ to AC and has advantage on Strength and Constitution saving throws. While in\
    \ its shell, the tortle is [prone](/3-Mechanics/CLI/rules/conditions.md#prone),\
    \ its speed is 0 and can't increase, it has disadvantage on Dexterity saving throws,\
    \ it can't take reactions, and the only action it can take is a bonus action to\
    \ emerge."
  "name": "Shell Defense"
"source":
- "MTF"
- "TTP"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MTF/Tortle.webp"
```
^statblock

```encounter-table
name: Tortle
creatures:
 - 1: Tortle
```

## Environment

coastal