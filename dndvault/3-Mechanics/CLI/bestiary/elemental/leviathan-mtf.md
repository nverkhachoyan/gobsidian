---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mtf
- monster/cr/20
- monster/environment/coastal
- monster/environment/underwater
- monster/size/gargantuan
- monster/type/elemental
aliases: ["Leviathan"]
NoteIcon: monster
BestiaryType: elemental
SourceType: Bestiary
BookSource: Mordenkainen's Tome of Foes p. 198
---
# [Leviathan](3-Mechanics\CLI\bestiary\elemental/leviathan-mtf.md)
*Source: Mordenkainen's Tome of Foes p. 198*  

## Leviathan

A towering wall of water that drags ships down to the ocean's depths and washes away coastal settlements—that phenomenon typifies the destruction a leviathan can unleash on the world. When called forth, a leviathan arises from a large body of water to form an immense serpent-shaped creature.

## Elder Elementals

On their native planes, elementals sweep across the weird and tempestuous landscape. Some possess greater power, gained by feeding on their lesser kin and adding the essence of creatures they have devoured to their own until they become something extraordinary. When summoned, these elder elementals manifest as beings of apocalyptic capability, entities whose mere existence promises destruction.

## Deadly When Summoned

The methods for summoning elder elementals remain hidden in forbidden tomes or inscribed on the walls of lost temples raised to honor the Elder Elemental Eye. Only casters of superlative skill have even the faintest chance of calling forth one of these monsters, and the spellcaster is often destroyed by the effort. Thus, only the most unhinged and nihilistic members of Elemental Evil cults attempt such a summoning, in the hope of hastening the world toward some cataclysmic end.

## Elemental Nature

An elder elemental doesn't require air, food, drink, or sleep.

```statblock
"name": "Leviathan (MTF)"
"size": "Gargantuan"
"type": "elemental"
"alignment": "Neutral"
"ac": !!int "17"
"hp": !!int "328"
"hit_dice": "16d20 + 160"
"stats":
- !!int "30"
- !!int "24"
- !!int "30"
- !!int "2"
- !!int "18"
- !!int "17"
"speed": "40 ft., swim 120 ft."
"saves":
  "Charisma": !!int "9"
  "Wisdom": !!int "10"
"damage_resistances": "bludgeoning, piercing, slashing from nonmagical attacks"
"damage_immunities": "acid, poison"
"condition_immunities": "[exhaustion](/3-Mechanics/CLI/rules/conditions.md#exhaustion),\
  \ [grappled](/3-Mechanics/CLI/rules/conditions.md#grappled), [paralyzed](/3-Mechanics/CLI/rules/conditions.md#paralyzed),\
  \ [petrified](/3-Mechanics/CLI/rules/conditions.md#petrified), [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned),\
  \ [prone](/3-Mechanics/CLI/rules/conditions.md#prone), [restrained](/3-Mechanics/CLI/rules/conditions.md#restrained),\
  \ [stunned](/3-Mechanics/CLI/rules/conditions.md#stunned)"
"senses": "darkvision 60 ft., passive Perception 14"
"languages": ""
"cr": "20"
"traits":
- "desc": "If the leviathan fails a saving throw, it can choose to succeed instead."
  "name": "Legendary Resistance (3/Day)"
- "desc": "If the leviathan takes 50 cold damage or more during a single turn, the\
    \ leviathan partially freezes; until the end of its next turn, its speeds are\
    \ reduced to 20 feet, and it makes attack rolls with disadvantage."
  "name": "Partial Freeze"
- "desc": "The leviathan deals double damage to objects and structures (included in\
    \ Tidal Wave)."
  "name": "Siege Monster"
- "desc": "The leviathan can enter a hostile creature's space and stop there. It can\
    \ move through a space as narrow as 1 inch wide without squeezing."
  "name": "Water Form"
"actions":
- "desc": "The leviathan makes two attacks: one with its slam and one with its tail."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+16 (+16) to hit, reach 20 ft., one\
    \ target. Hit: dice:1d10 + 10|text(15) (1d10 + 10) bludgeoning damage plus\
    \ dice:1d10|text(5) (1d10) acid damage."
  "name": "Slam"
- "desc": "Melee Weapon Attack: dice: d20+16 (+16) to hit, reach 20 ft., one\
    \ target. Hit: dice:1d12 + 10|text(16) (1d12 + 10) bludgeoning damage plus\
    \ dice:1d12|text(6) (1d12) acid damage."
  "name": "Tail"
- "desc": "While submerged, the leviathan magically creates a wall of water centered\
    \ on itself. The wall is up 250 feet long, up to 250 feet high, and up to 50 feet\
    \ thick. When the wall appears, all other creatures within its area must each\
    \ make a DC 24 Strength saving throw. A creature takes dice:6d10|text(33) (6d10)\
    \ bludgeoning damage on failed save, or half as much damage on a successful one.\n\
    \nAt the start of each of the leviathan's turns after the wall appears, the wall,\
    \ along with any other creatures in it, moves 50 feet away from the leviathan.\
    \ Any Huge or smaller creature inside the wall or whose space the wall enters\
    \ when it moves must succeed on a DC 24 Strength saving throw or take dice:5d10|text(27)\
    \ (5d10) bludgeoning damage. A creature takes this damage no more than once\
    \ on a turn. At the end of each turn the wall moves, the wall's height is reduced\
    \ by 50 feet, and the damage creatures take from the wall on subsequent rounds\
    \ is reduced by dice: 1d10|avg|noform (1d10). When the wall reaches 0 feet\
    \ in height, the effect ends.\n\nA creature caught in the wall can move by swimming.\
    \ Because of the force of the wave, though, the creature must make a successful\
    \ DC 24 Strength ([Athletics](/3-Mechanics/CLI/rules/skills.md#Athletics)) check\
    \ to swim at all during that turn."
  "name": "Tidal Wave (Recharge 6)"
"legendary_actions":
- "desc": "The leviathan makes one slam attack."
  "name": "Slam (Costs 2 Actions)"
- "desc": "The leviathan moves up to its speed."
  "name": "Move"
"source":
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MTF/Leviathan.webp"
```
^statblock

```encounter-table
name: Leviathan
creatures:
 - 1: Leviathan
```

## Environment

coastal, underwater