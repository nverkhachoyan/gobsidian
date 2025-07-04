---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mtf
- monster/cr/4
- monster/size/medium
- monster/type/fiend/devil
aliases: ["Merregon"]
NoteIcon: monster
BestiaryType: fiend (devil)
SourceType: Bestiary
BookSource: Mordenkainen's Tome of Foes p. 166
---
# [Merregon](3-Mechanics\CLI\bestiary\fiend/merregon-mtf.md)
*Source: Mordenkainen's Tome of Foes p. 166*  

> [!quote]- A quote from Mordenkainen  
> 
> Merregons can't speak, and their telepathy is one-to-one. Their orderly ranks are easily confused if you slay their shouting masters.

## Merregon

The souls of fallen soldiers, mercenaries, and bodyguards who served evil without reservation often find everlasting servitude in the Nine Hells as merregons. These faceless foot soldiers are the hells' legionnaires, tasked with protecting the realm and its rulers against intruders.

## Masks of Uniformity

Merregons have no individuality, and hence no need for faces. Every merregon legionnaire has a metal mask bolted to its head. Markings on the mask indicate the only elements of the wearer's identity that matter: its commander and the layer of the Nine Hells it serves.

## Fearless Obedience

Because of their unshakable loyalty, merregons form the backbone of many devils' protective retinues. They shrink from no task, no matter how dangerous. Unless ordered to fall back, they retreat from no fight.

```statblock
"name": "Merregon (MTF)"
"size": "Medium"
"type": "fiend"
"subtype": "devil"
"alignment": "Lawful Evil"
"ac": !!int "16"
"ac_class": "natural armor"
"hp": !!int "45"
"hit_dice": "6d8 + 18"
"stats":
- !!int "18"
- !!int "14"
- !!int "17"
- !!int "6"
- !!int "12"
- !!int "8"
"speed": "30 ft."
"damage_resistances": "cold; bludgeoning, piercing, slashing from nonmagical attacks\
  \ that aren't silvered"
"damage_immunities": "fire, poison"
"condition_immunities": "[frightened](/3-Mechanics/CLI/rules/conditions.md#frightened),\
  \ [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)"
"senses": "darkvision 60 ft., passive Perception 11"
"languages": "understands Infernal but can't speak, telepathy 120 ft."
"cr": "4"
"traits":
- "desc": "Magical darkness doesn't impede the merregon's darkvision."
  "name": "Devil's Sight"
- "desc": "The merregon has advantage on saving throws against spells and other magical\
    \ effects."
  "name": "Magic Resistance"
"actions":
- "desc": "The merregon makes two halberd attacks, or if an allied fiend of challenge\
    \ rating 6 or higher is within 60 feet of it, the merregon makes three halberd\
    \ attacks."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+6 (+6) to hit, reach 10 ft., one target.\
    \ Hit: dice:1d10 + 4|text(9) (1d10 + 4) slashing damage."
  "name": "Halberd"
- "desc": "Ranged Weapon Attack: dice: d20+4 (+4) to hit, range 100/400 ft.,\
    \ one target. Hit: dice:1d10 + 2|text(7) (1d10 + 2) piercing damage."
  "name": "Heavy Crossbow"
"reactions":
- "desc": "When another fiend within 5 feet of the merregon is hit by an attack, the\
    \ merregon causes itself to be hit instead."
  "name": "Loyal Bodyguard"
"source":
- "MTF"
- "BGDIA"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MTF/Merregon.webp"
```
^statblock

```encounter-table
name: Merregon
creatures:
 - 1: Merregon
```