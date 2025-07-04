---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/16
- monster/environment/desert
- monster/environment/mountain
- monster/size/gargantuan
- monster/type/elemental
aliases: ["Phoenix"]
NoteIcon: monster
BestiaryType: elemental
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 206, Mordenkainen's Tome of Foes p. 199
---
# [Phoenix](3-Mechanics\CLI\bestiary\elemental/phoenix-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 206, Mordenkainen's Tome of Foes p. 199*  

```statblock
"name": "Phoenix (MPMM)"
"size": "Gargantuan"
"type": "elemental"
"alignment": "Typically  Neutral"
"ac": !!int "18"
"hp": !!int "175"
"hit_dice": "10d20 + 70"
"stats":
- !!int "19"
- !!int "26"
- !!int "25"
- !!int "2"
- !!int "21"
- !!int "18"
"speed": "20 ft., fly 120 ft."
"saves":
  "Charisma": !!int "9"
  "Wisdom": !!int "10"
"damage_resistances": "bludgeoning, piercing, slashing from nonmagical attacks"
"damage_immunities": "fire, poison"
"condition_immunities": "[exhaustion](/3-Mechanics/CLI/rules/conditions.md#exhaustion),\
  \ [grappled](/3-Mechanics/CLI/rules/conditions.md#grappled), [paralyzed](/3-Mechanics/CLI/rules/conditions.md#paralyzed),\
  \ [petrified](/3-Mechanics/CLI/rules/conditions.md#petrified), [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned),\
  \ [prone](/3-Mechanics/CLI/rules/conditions.md#prone), [restrained](/3-Mechanics/CLI/rules/conditions.md#restrained),\
  \ [stunned](/3-Mechanics/CLI/rules/conditions.md#stunned)"
"senses": "darkvision 60 ft., passive Perception 15"
"languages": ""
"cr": "16"
"traits":
- "desc": "If the phoenix dies, it explodes. Each creature in 60-foot-radius sphere\
    \ centered on the phoenix must make a DC 20 Dexterity saving throw, taking dice:4d10|text(22)\
    \ (4d10) fire damage on a failed save, or half as much damage on a successful\
    \ one. The fire ignites flammable objects in the area that aren't being worn or\
    \ carried.\n\nThe explosion destroys the phoenix's body and leaves behind an egg-shaped\
    \ cinder, which weighs 5 pounds. The cinder deals dice:6d6|text(21) (6d6)\
    \ fire damage to any creature that touches it, though no more than once per round.\
    \ The cinder is immune to all damage, and after dice: 1d6|avg|noform (1d6)\
    \ days, it hatches a new phoenix."
  "name": "Fiery Death and Rebirth"
- "desc": "The phoenix can move through a space as narrow as 1 inch wide without squeezing.\n\
    \nAny creature that touches the phoenix or hits it with a melee attack while within\
    \ 5 feet of it takes dice:1d10|text(5) (1d10) fire damage. In addition, the\
    \ phoenix can enter a hostile creature's space and stop there. The first time\
    \ it enters a creature's space on a turn, that creature takes dice:1d10|text(5)\
    \ (1d10) fire damage.\n\nWith a touch, the phoenix can also ignite flammable\
    \ objects that aren't worn or carried (no action required)."
  "name": "Fire Form"
- "desc": "The phoenix doesn't provoke opportunity attacks when it flies out of an\
    \ enemy's reach."
  "name": "Flyby"
- "desc": "The phoenix sheds bright light in a 60-foot radius and dim light for an\
    \ additional 30 feet."
  "name": "Illumination"
- "desc": "If the phoenix fails a saving throw, it can choose to succeed instead."
  "name": "Legendary Resistance (3/Day)"
- "desc": "The phoenix deals double damage to objects and structures."
  "name": "Siege Monster"
"actions":
- "desc": "The phoenix makes two attacks: one Beak attack and one Fiery Talons attack."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+13 (+13) to hit, reach 15 ft., one\
    \ target. Hit: dice:2d6 + 8|text(15) (2d6 + 8) fire damage."
  "name": "Beak"
- "desc": "Melee Weapon Attack: dice: d20+13 (+13) to hit, reach 15 ft., one\
    \ target. Hit: dice:2d8 + 8|text(17) (2d8 + 8) fire damage."
  "name": "Fiery Talons"
"legendary_actions":
- "desc": "The phoenix moves up to its speed."
  "name": "Move"
- "desc": "The phoenix makes one beak attack."
  "name": "Peck"
- "desc": "The phoenix moves up to its speed and makes one Fiery Talons attack."
  "name": "Swoop (Costs 2 Actions)"
"source":
- "MPMM"
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Phoenix.webp"
```
^statblock

```encounter-table
name: Phoenix
creatures:
 - 1: Phoenix
```

## Environment

desert, mountain