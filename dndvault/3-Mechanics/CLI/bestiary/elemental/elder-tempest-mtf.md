---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mtf
- monster/cr/23
- monster/environment/arctic
- monster/environment/coastal
- monster/environment/grassland
- monster/environment/hill
- monster/environment/mountain
- monster/size/gargantuan
- monster/type/elemental
aliases: ["Elder Tempest"]
NoteIcon: monster
BestiaryType: elemental
SourceType: Bestiary
BookSource: Mordenkainen's Tome of Foes p. 200
---
# [Elder Tempest](3-Mechanics\CLI\bestiary\elemental/elder-tempest-mtf.md)
*Source: Mordenkainen's Tome of Foes p. 200*  

## Elder Tempest

Terrifying storms manifest in the body of the elder tempest. A being carved from clouds, wind, rain, and lightning, the elder tempest assumes the shape of a serpent that slithers through the sky. The tempest drowns the land beneath it with rain and stabs the earth with lances of lightning. Punishing winds scream around it as it flies, feeding the chaos it creates.

## Elder Elementals

On their native planes, elementals sweep across the weird and tempestuous landscape. Some possess greater power, gained by feeding on their lesser kin and adding the essence of creatures they have devoured to their own until they become something extraordinary. When summoned, these elder elementals manifest as beings of apocalyptic capability, entities whose mere existence promises destruction.

## Deadly When Summoned

The methods for summoning elder elementals remain hidden in forbidden tomes or inscribed on the walls of lost temples raised to honor the Elder Elemental Eye. Only casters of superlative skill have even the faintest chance of calling forth one of these monsters, and the spellcaster is often destroyed by the effort. Thus, only the most unhinged and nihilistic members of Elemental Evil cults attempt such a summoning, in the hope of hastening the world toward some cataclysmic end.

## Elemental Nature

An elder elemental doesn't require air, food, drink, or sleep.

```statblock
"name": "Elder Tempest (MTF)"
"size": "Gargantuan"
"type": "elemental"
"alignment": "Neutral"
"ac": !!int "19"
"hp": !!int "264"
"hit_dice": "16d20 + 96"
"stats":
- !!int "23"
- !!int "28"
- !!int "23"
- !!int "2"
- !!int "21"
- !!int "18"
"speed": "0 ft., fly 120 ft. (hover)"
"saves":
  "Charisma": !!int "11"
  "Wisdom": !!int "12"
"damage_resistances": "bludgeoning, piercing, slashing from nonmagical attacks"
"damage_immunities": "lightning, poison, thunder"
"condition_immunities": "[exhaustion](/3-Mechanics/CLI/rules/conditions.md#exhaustion),\
  \ [grappled](/3-Mechanics/CLI/rules/conditions.md#grappled), [paralyzed](/3-Mechanics/CLI/rules/conditions.md#paralyzed),\
  \ [petrified](/3-Mechanics/CLI/rules/conditions.md#petrified), [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned),\
  \ [prone](/3-Mechanics/CLI/rules/conditions.md#prone), [restrained](/3-Mechanics/CLI/rules/conditions.md#restrained),\
  \ [stunned](/3-Mechanics/CLI/rules/conditions.md#stunned)"
"senses": "darkvision 60 ft., passive Perception 15"
"languages": ""
"cr": "23"
"traits":
- "desc": "The tempest can enter a hostile creature's space and stop there. It can\
    \ move through a space as narrow as 1 inch wide without squeezing."
  "name": "Air Form"
- "desc": "The tempest doesn't provoke opportunity attacks when it flies out of an\
    \ enemy's reach."
  "name": "Flyby"
- "desc": "If the tempest fails a saving throw, it can choose to succeed instead."
  "name": "Legendary Resistance (3/Day)"
- "desc": "The tempest is always at the center of a storm dice: 1d6 + 4|avg|noform\
    \ (1d6 + 4) miles in diameter. Heavy precipitation in the form of either rain\
    \ or snow falls there, causing the area to be lightly obscured. Heavy rain also\
    \ extinguishes open flames and imposes disadvantage on Wisdom ([Perception](/3-Mechanics/CLI/rules/skills.md#Perception))\
    \ checks that rely on hearing.\n\nIn addition, strong winds swirl in the area\
    \ covered by the storm. The winds impose disadvantage on ranged attack rolls.\
    \ The winds extinguish open flames and disperse fog."
  "name": "Living Storm"
- "desc": "The tempest deals double damage to objects and structures."
  "name": "Siege Monster"
"actions":
- "desc": "The tempest makes two attacks with its thunderous slam."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+16 (+16) to hit, reach 20 ft., one\
    \ target. Hit: dice:4d6 + 9|text(23) (4d6 + 9) thunder damage."
  "name": "Thunderous Slam"
- "desc": "All other creatures within 120 feet of the tempest must each make a DC\
    \ 20 Dexterity saving throw, taking dice:6d8|text(27) (6d8) lightning damage\
    \ on a failed save, or half as much damage on a successful one. If a target's\
    \ saving throw fails by 5 or more, the creature is also [stunned](/3-Mechanics/CLI/rules/conditions.md#stunned)\
    \ until the end of its next turn."
  "name": "Lightning Storm (Recharge 6)"
"legendary_actions":
- "desc": "The tempest moves up to its speed."
  "name": "Move"
- "desc": "The tempest can cause a bolt of lightning to strike a point on the ground\
    \ anywhere under its storm. Each creature within 5 feet of that point must make\
    \ a DC 20 Dexterity saving throw, taking dice:3d10|text(16) (3d10) lightning\
    \ damage on a failed save, or half as much damage on a successful one."
  "name": "Lightning Strike (Costs 2 Actions)"
- "desc": "The tempest releases a blast of thunder and wind in a line that is 1 mile\
    \ long and 20 feet wide. Objects in that area take dice:4d10|text(22) (4d10)\
    \ thunder damage. Each creature there must succeed on a DC 21 Dexterity saving\
    \ throw or take dice:4d10|text(22) (4d10) thunder damage and be flung up to\
    \ 60 feet in a direction away from the line. If a thrown target collides with\
    \ an immovable object, such as a wall or floor, the target takes dice:1d6|text(3)\
    \ (1d6) bludgeoning damage for every 10 feet it was thrown before impact. If\
    \ the target would collide with another creature instead, that other creature\
    \ must succeed on a DC 19 Dexterity saving throw or take the same damage and be\
    \ knocked [prone](/3-Mechanics/CLI/rules/conditions.md#prone)."
  "name": "Screaming Gale (Costs 3 Actions)"
"source":
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MTF/Elder%20Tempest.webp"
```
^statblock

```encounter-table
name: Elder Tempest
creatures:
 - 1: Elder Tempest
```

## Environment

arctic, coastal, grassland, hill, mountain