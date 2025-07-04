---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mtf
- monster/cr/7
- monster/environment/underdark
- monster/size/medium
- monster/type/fiend/yugoloth
aliases: ["Dhergoloth"]
NoteIcon: monster
BestiaryType: fiend (yugoloth)
SourceType: Bestiary
BookSource: Mordenkainen's Tome of Foes p. 248
---
# [Dhergoloth](3-Mechanics\CLI\bestiary\fiend/dhergoloth-mtf.md)
*Source: Mordenkainen's Tome of Foes p. 248*  

> [!quote]- A quote from Mordenkainen  
> 
> A dhergoloth's head doesn't turn along with its furiously spinning torso, and its torso can spin a different direction from its dancing legs. I'd like to vivisect one at some point to find out how this can be.

## Dhergoloth

Dhergoloths rush into battle like whirlwinds of destruction, lashing out with five sets of claws, which extend from their squat, barrel-shaped bodies. They take contracts to put down uprisings, clear out rabble, and eliminate scouts and skirmishers, and they revel in the butchery they create, their unhinged laughter rising above their victims' screams.

Since dhergoloths are little more than dumb brutes, employers must use caution when instructing these fiends. They can handle simple orders that don't take a lot of time to resolve. When given anything complex to do, they either forget what they're told or don't listen in the first place, and then bungle the task that was set for them.

## Yugoloths

Mercenaries that ply their trade throughout the Lower Planes and in other realms, yugoloths have a reputation for effectiveness that is matched only by their desire for ever more wealth. Although yugoloths aren't especially loyal and typically try to exploit every potential loophole in a contract, they undertake any task for which they are hired, no matter how despicable. Yugoloths come in a wide variety of forms, including those described in the Monster Manual and the six creatures presented here.

```statblock
"name": "Dhergoloth (MTF)"
"size": "Medium"
"type": "fiend"
"subtype": "yugoloth"
"alignment": "Neutral Evil"
"ac": !!int "15"
"ac_class": "natural armor"
"hp": !!int "119"
"hit_dice": "14d8 + 56"
"stats":
- !!int "17"
- !!int "10"
- !!int "19"
- !!int "7"
- !!int "10"
- !!int "9"
"speed": "30 ft."
"saves":
  "Strength": !!int "6"
"damage_resistances": "cold; fire; lightning; bludgeoning, piercing, slashing from\
  \ nonmagical attacks"
"damage_immunities": "acid, poison"
"condition_immunities": "[poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)"
"senses": "blindsight 60 ft., darkvision 60 ft., passive Perception 10"
"languages": "Abyssal, Infernal, telepathy 60 ft."
"cr": "7"
"traits":
- "desc": "The dhergoloth's innate spellcasting ability is Charisma (spell save DC\
    \ 10). It can innately cast the following spells, requiring no material components:\n\
    \nAt will: [darkness](/3-Mechanics/CLI/spells/darkness.md), [fear](/3-Mechanics/CLI/spells/fear.md)\n\
    \n3/day: [sleep](/3-Mechanics/CLI/spells/sleep.md)"
  "name": "Innate Spellcasting"
- "desc": "The dhergoloth has advantage on saving throws against spells and other\
    \ magical effects."
  "name": "Magic Resistance"
- "desc": "The dhergoloth's weapon attacks are magical."
  "name": "Magic Weapons"
"actions":
- "desc": "The dhergoloth makes two claw attacks."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+6 (+6) to hit, reach 5 ft., one target.\
    \ Hit: dice:2d8 + 3|text(12) (2d8 + 3) slashing damage."
  "name": "Claw"
- "desc": "The dhergoloth moves up to its walking speed in a straight line and targets\
    \ each creature within 5 feet of it during its movement. Each target must succeed\
    \ on a DC 14 Dexterity saving throw or take dice:3d12 + 3|text(22) (3d12 +\
    \ 3) slashing damage."
  "name": "Flailing Claws (Recharge 5-6)"
- "desc": "The dhergoloth magically teleports, along with any equipment it is wearing\
    \ or carrying, up to 60 feet to an unoccupied space it can see."
  "name": "Teleport"
"source":
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MTF/Dhergoloth.webp"
```
^statblock

```encounter-table
name: Dhergoloth
creatures:
 - 1: Dhergoloth
```

## Environment

underdark