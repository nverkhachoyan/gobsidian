---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mtf
- monster/cr/15
- monster/environment/urban
- monster/size/medium
- monster/type/fiend/devil
aliases: ["Green Abishai"]
NoteIcon: monster
BestiaryType: fiend (devil)
SourceType: Bestiary
BookSource: Mordenkainen's Tome of Foes p. 162
---
# [Green Abishai](3-Mechanics\CLI\bestiary\fiend/green-abishai-mtf.md)
*Source: Mordenkainen's Tome of Foes p. 162*  

## Green Abishai

The envoys of Tiamat's armies, green abishais represent the god's interests in the Nine Hells and beyond. Their keen senses make them adept at discovering secrets and other sensitive information, while their diplomatic skills and their magic ensure that they can manipulate even the shrewdest opponents.

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
"name": "Green Abishai (MTF)"
"size": "Medium"
"type": "fiend"
"subtype": "devil"
"alignment": "Lawful Evil"
"ac": !!int "18"
"ac_class": "natural armor"
"hp": !!int "187"
"hit_dice": "25d8 + 75"
"stats":
- !!int "12"
- !!int "17"
- !!int "16"
- !!int "17"
- !!int "12"
- !!int "19"
"speed": "30 ft., fly 40 ft."
"saves":
  "Charisma": !!int "9"
  "Intelligence": !!int "8"
"skillsaves":
  "Deception": !!int "9"
  "Insight": !!int "6"
  "Perception": !!int "6"
  "Persuasion": !!int "9"
"damage_resistances": "cold; bludgeoning, piercing, slashing from nonmagical attacks\
  \ that aren't silvered"
"damage_immunities": "fire, poison"
"condition_immunities": "[poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)"
"senses": "darkvision 120 ft., passive Perception 16"
"languages": "Draconic, Infernal, telepathy 120 ft."
"cr": "15"
"traits":
- "desc": "The abishai's innate spellcasting ability is Charisma (spell save DC 17).\
    \ It can innately cast the following spells, requiring no material components:\n\
    \nAt will: [alter self](/3-Mechanics/CLI/spells/alter-self.md), [major image](/3-Mechanics/CLI/spells/major-image.md)\n\
    \n1/day each: [confusion](/3-Mechanics/CLI/spells/confusion.md), [dominate\
    \ person](/3-Mechanics/CLI/spells/dominate-person.md), [mass suggestion](/3-Mechanics/CLI/spells/mass-suggestion.md)\n\
    \n3/day each: [charm person](/3-Mechanics/CLI/spells/charm-person.md), [detect\
    \ thoughts](/3-Mechanics/CLI/spells/detect-thoughts.md), [fear](/3-Mechanics/CLI/spells/fear.md)"
  "name": "Innate Spellcasting"
- "desc": "Magical darkness doesn't impede the abishai's darkvision."
  "name": "Devil's Sight"
- "desc": "The abishai has advantage on saving throws against spells and other magical\
    \ effects."
  "name": "Magic Resistance"
- "desc": "The abishai's weapon attacks are magical."
  "name": "Magic Weapons"
"actions":
- "desc": "The abishai makes two attacks, one with its claws and one with its longsword,\
    \ or it casts one spell from its Innate Spellcasting trait and makes one claw\
    \ attack."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+6 (+6) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d8 + 1|text(5) (1d8 + 1) slashing damage, or dice:1d10 + 1|text(6)\
    \ (1d10 + 1) slashing damage if used with two hands."
  "name": "Longsword"
- "desc": "Melee Weapon Attack: dice: d20+8 (+8) to hit, reach 5 ft., one target.\
    \ Hit: dice:2d8 + 3|text(12) (2d8 + 3) piercing damage. If the target is\
    \ a creature, it must succeed on a DC 16 Constitution saving throw or take dice:2d10|text(11)\
    \ (2d10) poison damage and become [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)\
    \ for 1 minute. The [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)\
    \ target can repeat the saving throw at the end of each of its turns, ending the\
    \ effect on itself on a success."
  "name": "Claws"
"source":
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MTF/Green%20Abishai.webp"
```
^statblock

```encounter-table
name: Green Abishai
creatures:
 - 1: Green Abishai
```

## Environment

urban