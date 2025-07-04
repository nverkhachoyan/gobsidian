---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mtf
- monster/cr/1-2
- monster/size/medium
- monster/type/fiend/devil
aliases: ["Nupperibo"]
NoteIcon: monster
BestiaryType: fiend (devil)
SourceType: Bestiary
BookSource: Mordenkainen's Tome of Foes p. 168
---
# [Nupperibo](3-Mechanics\CLI\bestiary\fiend/nupperibo-mtf.md)
*Source: Mordenkainen's Tome of Foes p. 168*  

> [!quote]- A quote from Mordenkainen  
> 
> A lemure emerges from the Styx wiped of memory, yet the patterns of evil it performed in life remain indelibly inscribed upon its soul. Those who lacked ambition cannot climb the hierarchical ladder of the Hells. They instead step down, becoming nupperibos.

## Nupperibo

No soul is turned away from the Nine Hells, but the truly worthless—those whose evil acts in life arose from carelessness and sloth more than anything else—are suitable only to become nupperibos. These pitiful creatures shuffle mindlessly across the landscape: blind, bloated from unquenchable hunger, and groping for whatever scraps of fetid matter or swarming vermin they can scoop into their groaning mouths.

## Nauseating Bulk

Individually, nupperibos are pathetic, but they're rarely alone and can be dangerous when gathered into packs. They herd together into throngs that can clog a vital passage or an entire valley. Clouds of stinging insects, stirges, and other vermin surround them in a terrifying, reeking sheath that torments any non-devil that draws near.

## Hunger Unending

A nupperibo knows nothing but the hunger that propels it on a blind quest for anything to devour. Once it senses a potential meal, it pursues that prey tirelessly until the food is consumed, the nupperibo is slain, or some other morsel crosses the fiend's path and distracts it.

## Slavish Obedience

With no interest of its own beyond the need to consume, a nupperibo obeys unthinkingly any command it receives telepathically from another devil. This blind loyalty makes them the easiest of infernal troops to lead into battle, but their presence in a legion does nothing to elevate its general's status.

```statblock
"name": "Nupperibo (MTF)"
"size": "Medium"
"type": "fiend"
"subtype": "devil"
"alignment": "Lawful Evil"
"ac": !!int "13"
"ac_class": "natural armor"
"hp": !!int "11"
"hit_dice": "2d8 + 2"
"stats":
- !!int "16"
- !!int "11"
- !!int "13"
- !!int "3"
- !!int "8"
- !!int "1"
"speed": "20 ft."
"skillsaves":
  "Perception": !!int "1"
"damage_resistances": "acid; cold; bludgeoning, piercing, slashing from nonmagical\
  \ attacks that aren't silvered"
"damage_immunities": "fire, poison"
"condition_immunities": "[blinded](/3-Mechanics/CLI/rules/conditions.md#blinded),\
  \ [charmed](/3-Mechanics/CLI/rules/conditions.md#charmed), [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened),\
  \ [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)"
"senses": "blindsight 10 ft. (blind beyond this radius), passive Perception 11"
"languages": "understands Infernal but can't speak"
"cr": "1/2"
"traits":
- "desc": "Any creature, other than a devil, that starts its turn within 20 feet of\
    \ the nupperibo must make a DC 11 Constitution saving throw. A creature within\
    \ the areas of two or more nupperibos makes the saving throw with disadvantage.\
    \ On a failure, the creature takes dice:1d4|text(2) (1d4) piercing damage."
  "name": "Cloud of Vermin"
- "desc": "In the Nine Hells, the nupperibos can flawlessly track any creature that\
    \ has taken damage from any nupperibo's Cloud of Vermin within the previous 24\
    \ hours."
  "name": "Hunger-Driven"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+5 (+5) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d6 + 3|text(6) (1d6 + 3) piercing damage."
  "name": "Bite"
"source":
- "MTF"
- "BGDIA"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MTF/Nupperibo.webp"
```
^statblock

```encounter-table
name: Nupperibo
creatures:
 - 1: Nupperibo
```