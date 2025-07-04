---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mtf
- monster/cr/17
- monster/environment/coastal
- monster/environment/urban
- monster/size/medium
- monster/type/fiend/devil
aliases: ["Blue Abishai"]
NoteIcon: monster
BestiaryType: fiend (devil)
SourceType: Bestiary
BookSource: Mordenkainen's Tome of Foes p. 161
---
# [Blue Abishai](3-Mechanics\CLI\bestiary\fiend/blue-abishai-mtf.md)
*Source: Mordenkainen's Tome of Foes p. 161*  

## Blue Abishai

Seekers of forgotten lore and lost relics, blue abishais are the most cunning and learned of their kind. Their research into occult subjects gleaned from tomes and grimoires plundered from across the multiverse enables them to become accomplished spellcasters. They use their magic to devastate their mistress's enemies.

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
"name": "Blue Abishai (MTF)"
"size": "Medium"
"type": "fiend"
"subtype": "devil"
"alignment": "Lawful Evil"
"ac": !!int "19"
"ac_class": "natural armor"
"hp": !!int "195"
"hit_dice": "26d8 + 78"
"stats":
- !!int "15"
- !!int "14"
- !!int "17"
- !!int "22"
- !!int "23"
- !!int "18"
"speed": "30 ft., fly 50 ft."
"saves":
  "Wisdom": !!int "12"
  "Intelligence": !!int "12"
"skillsaves":
  "Arcana": !!int "12"
"damage_resistances": "cold; bludgeoning, piercing, slashing from nonmagical attacks\
  \ that aren't silvered"
"damage_immunities": "fire, lightning, poison"
"condition_immunities": "[poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)"
"senses": "darkvision 120 ft., passive Perception 16"
"languages": "Draconic, Infernal, telepathy 120 ft."
"cr": "17"
"traits":
- "desc": "The abishai is a 13th-level spellcaster. Its spellcasting ability is Intelligence\
    \ (spell save DC 20, dice: d20+12 (+12) to hit with spell attacks). The abishai\
    \ has the following wizard spells prepared:\n\nCantrips (at will): [friends](/3-Mechanics/CLI/spells/friends.md),\
    \ [mage hand](/3-Mechanics/CLI/spells/mage-hand.md), [message](/3-Mechanics/CLI/spells/message.md),\
    \ [minor illusion](/3-Mechanics/CLI/spells/minor-illusion.md), [shocking grasp](/3-Mechanics/CLI/spells/shocking-grasp.md)\n\
    \n1st level (4 slots): [chromatic orb](/3-Mechanics/CLI/spells/chromatic-orb.md),\
    \ [disguise self](/3-Mechanics/CLI/spells/disguise-self.md), [expeditious retreat](/3-Mechanics/CLI/spells/expeditious-retreat.md),\
    \ [magic missile](/3-Mechanics/CLI/spells/magic-missile.md), [charm person](/3-Mechanics/CLI/spells/charm-person.md),\
    \ [thunderwave](/3-Mechanics/CLI/spells/thunderwave.md)\n\n2nd level (3 slots):\
    \ [darkness](/3-Mechanics/CLI/spells/darkness.md), [mirror image](/3-Mechanics/CLI/spells/mirror-image.md),\
    \ [misty step](/3-Mechanics/CLI/spells/misty-step.md)\n\n3rd level (3 slots):\
    \ [dispel magic](/3-Mechanics/CLI/spells/dispel-magic.md), [fear](/3-Mechanics/CLI/spells/fear.md),\
    \ [lightning bolt](/3-Mechanics/CLI/spells/lightning-bolt.md)\n\n4th level (3\
    \ slots): [dimension door](/3-Mechanics/CLI/spells/dimension-door.md), [greater\
    \ invisibility](/3-Mechanics/CLI/spells/greater-invisibility.md), [ice storm](/3-Mechanics/CLI/spells/ice-storm.md)\n\
    \n5th level (2 slots): [cone of cold](/3-Mechanics/CLI/spells/cone-of-cold.md),\
    \ [wall of force](/3-Mechanics/CLI/spells/wall-of-force.md)\n\n6th level (1\
    \ slots): [chain lightning](/3-Mechanics/CLI/spells/chain-lightning.md)\n\n\
    7th level (1 slots): [teleport](/3-Mechanics/CLI/spells/teleport.md)"
  "name": "Spellcasting"
- "desc": "Magical darkness doesn't impede the abishai's darkvision."
  "name": "Devil's Sight"
- "desc": "The abishai has advantage on saving throws against spells and other magical\
    \ effects."
  "name": "Magic Resistance"
- "desc": "The abishai's weapon attacks are magical."
  "name": "Magic Weapons"
"actions":
- "desc": "The abishai makes two attacks: one with its quarterstaff and one with its\
    \ bite."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+8 (+8) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d6 + 2|text(5) (1d6 + 2) bludgeoning damage, or dice:1d8 +\
    \ 2|text(6) (1d8 + 2) bludgeoning damage if used with two hands."
  "name": "Quarterstaff"
- "desc": "Melee Weapon Attack: dice: d20+8 (+8) to hit, reach 5 ft., one target.\
    \ Hit: dice:2d10 + 2|text(13) (2d10 + 2) piercing damage plus dice:4d6|text(14)\
    \ (4d6) lightning damage."
  "name": "Bite"
"source":
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MTF/Blue%20Abishai.webp"
```
^statblock

```encounter-table
name: Blue Abishai
creatures:
 - 1: Blue Abishai
```

## Environment

coastal, urban