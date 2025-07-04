---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mtf
- monster/cr/16
- monster/environment/desert
- monster/environment/mountain
- monster/size/gargantuan
- monster/type/elemental
aliases: ["Phoenix"]
NoteIcon: monster
BestiaryType: elemental
SourceType: Bestiary
BookSource: Mordenkainen's Tome of Foes p. 199
---
# [Phoenix](3-Mechanics\CLI\bestiary\elemental/phoenix-mtf.md)
*Source: Mordenkainen's Tome of Foes p. 199*  

> [!quote]- A quote from Mordenkainen  
> 
> To rise like a phoenix from the ashes—so many use that quaint colloquialism. Little do they know about the true horror of such a rebirth.

## Phoenix

Releasing a phoenix from the Inner Planes creates an explosion of fire that spreads across the sky. An enormous fiery bird forms in the center of the flames and smoke—an elder elemental possessed by a need to burn everything to ash. The phoenix rarely stays in one place for long as it strives to transform the world into an inferno.

## Elder Elementals

On their native planes, elementals sweep across the weird and tempestuous landscape. Some possess greater power, gained by feeding on their lesser kin and adding the essence of creatures they have devoured to their own until they become something extraordinary. When summoned, these elder elementals manifest as beings of apocalyptic capability, entities whose mere existence promises destruction.

## Deadly When Summoned

The methods for summoning elder elementals remain hidden in forbidden tomes or inscribed on the walls of lost temples raised to honor the Elder Elemental Eye. Only casters of superlative skill have even the faintest chance of calling forth one of these monsters, and the spellcaster is often destroyed by the effort. Thus, only the most unhinged and nihilistic members of Elemental Evil cults attempt such a summoning, in the hope of hastening the world toward some cataclysmic end.

## Elemental Nature

An elder elemental doesn't require air, food, drink, or sleep.

```statblock
"name": "Phoenix (MTF)"
"size": "Gargantuan"
"type": "elemental"
"alignment": "Neutral"
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
- "desc": "When the phoenix dies, it explodes. Each creature within 60-feet of it\
    \ must make a DC 20 Dexterity saving throw, taking dice:4d10|text(22) (4d10)\
    \ fire damage on a failed save, or half as much damage on a successful one. The\
    \ fire ignites flammable objects in the area that aren't worn or carried.\n\n\
    The explosion destroys the phoenix's body and leaves behind an egg-shaped cinder\
    \ that weighs 5 pounds. The cinder is blazing hot, dealing dice:6d6|text(21)\
    \ (6d6) fire damage to any creature that touches it, though no more than once\
    \ per round. The cinder is immune to all damage, and after dice: 1d6|avg|noform\
    \ (1d6) days, it hatches a new phoenix."
  "name": "Fiery Death and Rebirth"
- "desc": "The phoenix can move through a space as narrow as 1 inch wide without squeezing.\
    \ Any creature that touches the phoenix or hits it with a melee attack while within\
    \ 5 feet of it takes dice:1d10|text(5) (1d10) fire damage. In addition, the\
    \ phoenix can enter a hostile creature's space and stop there. The first time\
    \ it enters a creature's space on a turn, that creature takes dice:1d10|text(5)\
    \ (1d10) fire damage. With a touch, the phoenix can also ignite flammable objects\
    \ that aren't worn or carried (no action required)."
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
- "desc": "The phoenix makes two attacks: one with its beak and one with its fiery\
    \ talons."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+13 (+13) to hit, reach 15 ft., one\
    \ target. Hit: dice:2d6 + 8|text(15) (2d6 + 8) fire damage. If the target\
    \ is a creature or a flammable object, it ignites. Until a creature takes an action\
    \ to douse the fire, the target takes dice:1d10|text(5) (1d10) fire damage\
    \ at the start of each of its turns."
  "name": "Beak"
- "desc": "Melee Weapon Attack: dice: d20+13 (+13) to hit, reach 15 ft., one\
    \ target. Hit: dice:2d8 + 8|text(17) (2d8 + 8) fire damage."
  "name": "Fiery Talons"
"legendary_actions":
- "desc": "The phoenix makes one beak attack."
  "name": "Peck"
- "desc": "The phoenix moves up to its speed."
  "name": "Move"
- "desc": "The phoenix moves up to its speed and attacks with its fiery talons."
  "name": "Swoop (Costs 2 Actions)"
"source":
- "MTF"
- "MOT"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MTF/Phoenix.webp"
```
^statblock

```encounter-table
name: Phoenix
creatures:
 - 1: Phoenix
```

## Environment

desert, mountain