---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mtf
- monster/cr/6
- monster/environment/urban
- monster/size/medium
- monster/type/fiend/devil
aliases: ["White Abishai"]
NoteIcon: monster
BestiaryType: fiend (devil)
SourceType: Bestiary
BookSource: Mordenkainen's Tome of Foes p. 163
---
# [White Abishai](3-Mechanics\CLI\bestiary\fiend/white-abishai-mtf.md)
*Source: Mordenkainen's Tome of Foes p. 163*  

## White Abishai

Though they are the least of their kind, white abishais fight with a reckless fury, making them ideally suited for bolstering the ranks of Tiamat's armies. White abishais fight without fear, becoming whirlwinds of destruction on the battlefield.

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
"name": "White Abishai (MTF)"
"size": "Medium"
"type": "fiend"
"subtype": "devil"
"alignment": "Lawful Evil"
"ac": !!int "15"
"ac_class": "natural armor"
"hp": !!int "68"
"hit_dice": "8d8 + 32"
"stats":
- !!int "16"
- !!int "11"
- !!int "18"
- !!int "11"
- !!int "12"
- !!int "13"
"speed": "30 ft., fly 40 ft."
"saves":
  "Strength": !!int "6"
  "Constitution": !!int "7"
"damage_resistances": "bludgeoning, piercing, slashing from nonmagical attacks that\
  \ aren't silvered"
"damage_immunities": "cold, fire, poison"
"condition_immunities": "[poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)"
"senses": "darkvision 120 ft., passive Perception 11"
"languages": "Draconic, Infernal, telepathy 120 ft."
"cr": "6"
"traits":
- "desc": "Magical darkness doesn't impede the abishai's darkvision."
  "name": "Devil's Sight"
- "desc": "The abishai has advantage on saving throws against spells and other magical\
    \ effects."
  "name": "Magic Resistance"
- "desc": "The abishai's weapon attacks are magical."
  "name": "Magic Weapons"
- "desc": "At the start of its turn, the abishai can gain advantage on all melee weapon\
    \ attack rolls during that turn, but attack rolls against it have advantage until\
    \ the start of its next turn."
  "name": "Reckless"
"actions":
- "desc": "The abishai makes two attacks: one with its longsword and one with its\
    \ claw."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+6 (+6) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d8 + 3|text(7) (1d8 + 3) slashing damage, or dice:1d10 + 3|text(8)\
    \ (1d10 + 3) slashing damage if used with two hands."
  "name": "Longsword"
- "desc": "Melee Weapon Attack: dice: d20+6 (+6) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d10 + 3|text(8) (1d10 + 3) slashing damage."
  "name": "Claw"
- "desc": "Melee Weapon Attack: dice: d20+6 (+6) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d4 + 3|text(5) (1d4 + 3) piercing damage plus dice:1d6|text(3)\
    \ (1d6) cold damage."
  "name": "Bite"
"reactions":
- "desc": "In response to taking damage, the abishai makes a bite attack against a\
    \ random creature within 5 feet of it. If no creature is within reach, the abishai\
    \ moves up to half its speed toward an enemy it can see, without provoking opportunity\
    \ attacks."
  "name": "Vicious Reprisal"
"source":
- "MTF"
- "BGDIA"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MTF/White%20Abishai.webp"
```
^statblock

```encounter-table
name: White Abishai
creatures:
 - 1: White Abishai
```

## Environment

urban