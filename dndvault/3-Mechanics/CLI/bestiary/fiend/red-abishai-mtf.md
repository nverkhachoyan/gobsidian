---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mtf
- monster/cr/19
- monster/environment/mountain
- monster/environment/urban
- monster/size/medium
- monster/type/fiend/devil
aliases: ["Red Abishai"]
NoteIcon: monster
BestiaryType: fiend (devil)
SourceType: Bestiary
BookSource: Mordenkainen's Tome of Foes p. 160
---
# [Red Abishai](3-Mechanics\CLI\bestiary\fiend/red-abishai-mtf.md)
*Source: Mordenkainen's Tome of Foes p. 160*  

## Red Abishai

Red abishais have no equals among the abishais when it comes to leadership ability and raw power. They can invoke Tiamat's authority to bend even dragons to their will. Red abishais lead other devils into battle or take charge of troublesome cults to ensure that they continue to carry out Tiamat's commands. A red abishai cuts a fearsome figure, and that sight can be inspiring to the abishai's allies, filling them with a fanatical willingness to fight.

## Abishai

Each abishai was once a mortal who somehow won Tiamat's favor before death and, as a reward, found its soul transformed into a hideous devil to serve at her pleasure in the Nine Hells.

## Emissaries of Doom

Tiamat deploys abishais as emissaries, sending them to represent her interests in the Hells and across the multiverse. Some have simple tasks, such as delivering a message to cultists or taking charge of worshipers to carry out a sensitive mission. Others have greater responsibilities, such as leading large groups, assassinating targets, and serving in armies. In all cases, abishais are fanatically loyal to Tiamat, ready to lay down their lives if needed.

## Outsiders in Hell

Abishais stand outside the normal hierarchy of the Nine Hells, having their own chain of command and ultimately answering to Tiamat (and Asmodeus, when the dark lord chooses to use them). Other archdevils can command abishais to work for them, but most archdevils do so rarely, since it is never clear whether an abishai follows Tiamat's orders or Asmodeus's.

There is inherent risk in countermanding an order given by Tiamat, but interfering with Asmodeus's plans invites certain destruction.

> [!quote]- A quote from Mordenkainen  
> 
> Tiamat is a force of Chaos bound to a place of Law. Are abishai her servants or her jailers?


```statblock
"name": "Red Abishai (MTF)"
"size": "Medium"
"type": "fiend"
"subtype": "devil"
"alignment": "Lawful Evil"
"ac": !!int "22"
"ac_class": "natural armor"
"hp": !!int "255"
"hit_dice": "30d8 + 120"
"stats":
- !!int "23"
- !!int "16"
- !!int "19"
- !!int "14"
- !!int "15"
- !!int "19"
"speed": "30 ft., fly 50 ft."
"saves":
  "Wisdom": !!int "8"
  "Strength": !!int "12"
  "Constitution": !!int "10"
"skillsaves":
  "Intimidation": !!int "10"
  "Perception": !!int "8"
"damage_resistances": "cold; bludgeoning, piercing, slashing from nonmagical attacks\
  \ that aren't silvered"
"damage_immunities": "fire, poison"
"condition_immunities": "[poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)"
"senses": "darkvision 120 ft., passive Perception 18"
"languages": "Draconic, Infernal, telepathy 120 ft."
"cr": "19"
"traits":
- "desc": "Magical darkness doesn't impede the abishai's darkvision."
  "name": "Devil's Sight"
- "desc": "The abishai has advantage on saving throws against spells and other magical\
    \ effects."
  "name": "Magic Resistance"
- "desc": "The abishai's weapon attacks are magical."
  "name": "Magic Weapons"
"actions":
- "desc": "The abishai can use its Frightful Presence. It also makes three attacks:\
    \ one with its morningstar, one with its claw, and one with its bite."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+12 (+12) to hit, reach 5 ft., one\
    \ target. Hit: dice:1d8 + 6|text(10) (1d8 + 6) piercing damage."
  "name": "Morningstar"
- "desc": "Melee Weapon Attack: dice: d20+12 (+12) to hit, reach 5 ft., one\
    \ target. Hit: dice:2d10 + 6|text(17) (2d10 + 6) slashing damage."
  "name": "Claw"
- "desc": "Melee Weapon Attack: dice: d20+12 (+12) to hit, reach 5 ft., one\
    \ target. Hit: dice:3d10 + 6|text(22) (3d10 + 6) piercing damage plus dice:7d10|text(38)\
    \ (7d10) fire damage."
  "name": "Bite"
- "desc": "Each creature of the abishai's choice that is within 120 feet and aware\
    \ of it must succeed on a DC 18 Wisdom saving throw or become [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened)\
    \ for 1 minute. A creature can repeat the saving throw at the end of each of its\
    \ turns, ending the effect on itself on a success. If a creature's saving throw\
    \ is successful or the effect ends for it, the creature is immune to the abishai's\
    \ Frightful Presence for the next 24 hours."
  "name": "Frightful Presence"
- "desc": "The abishai chooses up to four of its allies within 60 feet of it that\
    \ can see it. For 1 minute, each of those allies makes attack rolls with advantage\
    \ and can't be [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened)."
  "name": "Incite Fanaticism"
- "desc": "The abishai targets one dragon it can see within 120 feet of it. The dragon\
    \ must make a DC 18 Charisma saving throw. A chromatic dragon makes this save\
    \ with disadvantage. On a successful save, the target is immune to the abishai's\
    \ Power of the Dragon Queen for 1 hour. On a failed save, the target is [charmed](/3-Mechanics/CLI/rules/conditions.md#charmed)\
    \ by the abishai for 1 hour. While [charmed](/3-Mechanics/CLI/rules/conditions.md#charmed)\
    \ in this way, the target regards the abishai as a trusted friend to be heeded\
    \ and protected. This effect ends if the abishai or its companions deal damage\
    \ to the target."
  "name": "Power of the Dragon Queen"
"source":
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MTF/Red%20Abishai.webp"
```
^statblock

```encounter-table
name: Red Abishai
creatures:
 - 1: Red Abishai
```

## Environment

mountain, urban