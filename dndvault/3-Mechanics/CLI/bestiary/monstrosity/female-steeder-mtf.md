---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mtf
- monster/cr/1
- monster/environment/underdark
- monster/size/large
- monster/type/monstrosity
aliases: ["Female Steeder"]
NoteIcon: monster
BestiaryType: monstrosity
SourceType: Bestiary
BookSource: Mordenkainen's Tome of Foes p. 238
---
# [Female Steeder](3-Mechanics\CLI\bestiary\monstrosity/female-steeder-mtf.md)
*Source: Mordenkainen's Tome of Foes p. 238*  

> [!quote]- A quote from Mordenkainen  
> 
> Steeders resemble spiders as much as worgs resemble wolves. The creatures may appear similar, but steeders are more than mere vermin.

## Steeders

Giant hunting spiders, steeders prowl the depths of the Underdark. Most steeders are encountered in the company of duergar.

## Female Dominance

Female steeders grow larger and stronger than males, and the female often devours the male after breeding. In captivity, males are used as draft animals, while females serve as steeds in battle.

## Lone Predators

Steeders consider other steeders as enemies and attempt to tear apart perceived threats. Their duergar handlers must stable steeders separate from one another and place blinders on them when they're put to work to keep them from attacking each other.

## Low Cunning

Steeders are intelligent enough to learn simple hand signals and vocal commands, but even a domesticated steeder can turn against its handler. Training these beasts requires a rider to bond with the steeder, a process that begins shortly after the creature hatches. The rider stays with the steeder as it grows to full size, working throughout that time to channel the beast's predatory instincts.

## Deadly Hunters

Rather than spinning webs, steeders excrete a viscous substance from their legs. This goo allows them to creep along walls and ceilings and to grapple prey.

```statblock
"name": "Female Steeder (MTF)"
"size": "Large"
"type": "monstrosity"
"alignment": "Unaligned"
"ac": !!int "14"
"ac_class": "natural armor"
"hp": !!int "30"
"hit_dice": "4d10 + 8"
"stats":
- !!int "15"
- !!int "16"
- !!int "14"
- !!int "2"
- !!int "10"
- !!int "3"
"speed": "30 ft., climb 30 ft."
"skillsaves":
  "Stealth": !!int "7"
  "Perception": !!int "4"
"senses": "darkvision 120 ft., passive Perception 14"
"languages": ""
"cr": "1"
"traits":
- "desc": "The steeder can climb difficult surfaces, including upside down on ceilings,\
    \ without needing to make an ability check."
  "name": "Spider Climb"
- "desc": "The distance of the steeder's long jumps is tripled; every foot of its\
    \ walking speed that it spends on the jump allows it to move 3 feet."
  "name": "Extraordinary Leap"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+5 (+5) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d8 + 3|text(7) (1d8 + 3) piercing damage plus dice:2d8|text(9)\
    \ (2d8) poison damage."
  "name": "Bite"
- "desc": "Melee Weapon Attack: dice: d20+5 (+5) to hit, reach 5 ft., one Medium\
    \ or smaller creature. Hit: The target is stuck to the steeder's leg and is\
    \ [grappled](/3-Mechanics/CLI/rules/conditions.md#grappled) until it escapes (escape\
    \ DC 12). The steeder can have only one creature [grappled](/3-Mechanics/CLI/rules/conditions.md#grappled)\
    \ at a time."
  "name": "Sticky Leg"
"source":
- "MTF"
- "OotA"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MTF/Female%20Steeder.webp"
```
^statblock

```encounter-table
name: Female Steeder
creatures:
 - 1: Female Steeder
```

## Environment

underdark