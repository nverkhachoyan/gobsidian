---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mtf
- monster/cr/7
- monster/environment/urban
- monster/size/medium
- monster/type/fiend/devil
aliases: ["Black Abishai"]
NoteIcon: monster
BestiaryType: fiend (devil)
SourceType: Bestiary
BookSource: Mordenkainen's Tome of Foes p. 160
---
# [Black Abishai](3-Mechanics\CLI\bestiary\fiend/black-abishai-mtf.md)
*Source: Mordenkainen's Tome of Foes p. 160*  

## Black Abishai

Expert assassins and infiltrators, black abishais can weave shadows to mask their presence, allowing them to reach a location from where they can deliver a fatal strike to their targets.

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
"name": "Black Abishai (MTF)"
"size": "Medium"
"type": "fiend"
"subtype": "devil"
"alignment": "Lawful Evil"
"ac": !!int "15"
"ac_class": "natural armor"
"hp": !!int "58"
"hit_dice": "9d8 + 18"
"stats":
- !!int "14"
- !!int "17"
- !!int "14"
- !!int "13"
- !!int "16"
- !!int "11"
"speed": "30 ft., fly 40 ft."
"saves":
  "Dexterity": !!int "6"
  "Wisdom": !!int "6"
"skillsaves":
  "Stealth": !!int "6"
  "Perception": !!int "6"
"damage_resistances": "cold; bludgeoning, piercing, slashing from nonmagical attacks\
  \ that aren't silvered"
"damage_immunities": "acid, fire, poison"
"condition_immunities": "[poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)"
"senses": "darkvision 120 ft., passive Perception 16"
"languages": "Draconic, Infernal, telepathy 120 ft."
"cr": "7"
"traits":
- "desc": "Magical darkness doesn't impede the abishai's darkvision."
  "name": "Devil's Sight"
- "desc": "The abishai has advantage on saving throws against spells and other magical\
    \ effects."
  "name": "Magic Resistance"
- "desc": "The abishai's weapon attacks are magical."
  "name": "Magic Weapons"
- "desc": "While in dim light or darkness, the abishai can take the Hide action as\
    \ a bonus action."
  "name": "Shadow Stealth"
"actions":
- "desc": "The abishai makes three attacks: two with its scimitar and one with its\
    \ bite."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+6 (+6) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d6 + 3|text(6) (1d6 + 3) slashing damage."
  "name": "Scimitar"
- "desc": "Melee Weapon Attack: dice: d20+6 (+6) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d10 + 3|text(8) (1d10 + 3) piercing damage plus dice:2d8|text(9)\
    \ (2d8) acid damage."
  "name": "Bite"
- "desc": "The abishai casts [darkness](/3-Mechanics/CLI/spells/darkness.md) at a\
    \ point within 120 feet of it, requiring no components. Wisdom is its spellcasting\
    \ ability for this spell. While the spell persists, the abishai can move the area\
    \ of darkness up to 60 feet as a bonus action."
  "name": "Creeping Darkness (Recharge 6)"
"source":
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MTF/Black%20Abishai.webp"
```
^statblock

```encounter-table
name: Black Abishai
creatures:
 - 1: Black Abishai
```

## Environment

urban