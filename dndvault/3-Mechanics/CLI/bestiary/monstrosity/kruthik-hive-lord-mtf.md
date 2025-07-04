---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mtf
- monster/cr/5
- monster/environment/desert
- monster/environment/mountain
- monster/environment/underdark
- monster/size/large
- monster/type/monstrosity
aliases: ["Kruthik Hive Lord"]
NoteIcon: monster
BestiaryType: monstrosity
SourceType: Bestiary
BookSource: Mordenkainen's Tome of Foes p. 212
---
# [Kruthik Hive Lord](3-Mechanics\CLI\bestiary\monstrosity/kruthik-hive-lord-mtf.md)
*Source: Mordenkainen's Tome of Foes p. 212*  

> [!quote]- A quote from Mordenkainen  
> 
> Imagine a hive of ants the size of horses, but the ants are wearing armor. Then exterminate them.

## Kruthik Hive Lord

A hive lord rules each kruthik hive. When the hive lord dies, the surviving members of the hive abandon their lair and search for a new one. When a suitable location is found, the largest kruthik in the hive undergoes a metamorphosis, forming a cocoon around itself and emerging several weeks later as a hive lord—a bigger and smarter kruthik with the ability to spray digestive acid from its maw. The hive lord claims the largest chamber of the lair and keeps several adult kruthiks nearby as bodyguards.

## Kruthiks

Kruthiks are chitin-covered reptiles that hunt in packs and nest in sprawling subterranean warrens. They are attracted to sources of heat, such as dwarven forges and pools of molten lava, and carve out lairs as close to such locations as possible. As they burrow through the earth, they leave behind tunnels—evidence that is often the first clue to the nearby presence of a kruthik hive. Kruthiks also make use of preexisting underground chambers, incorporating them into their lairs when they can.

Kruthiks communicate with one another through a series of hisses and chittering noises. These sounds can often be heard in advance of a kruthik attack. Whenever their lair is invaded, kruthik guards send out an alarm by rapidly tapping the stone floor with their sharp legs.

## Sharp Senses

In addition to having an acute sense of smell, kruthiks can see in the dark and can detect vibrations in the earth around them. They take the scent of their own dead as a warning and avoid areas where many other kruthiks have died. Slaying a sufficient number of kruthiks in one area might cause the remaining hive members to move elsewhere.

## Sharper Weapons

Although they can feed on carrion, kruthiks prefer live prey. They kill enemies by impaling them on their spiked limbs, then grind up the flesh and bones with mandibles strong enough to chew rock. When several kruthiks gang up on a single foe, they become frenzied and even more lethal.

## Shared Lair

Kruthiks abide the presence of constructs, elementals, oozes, and undead, and use such creatures to help guard their hive. Kruthiks are smart enough to barricade some tunnels and dig new ones that keep their neighbors away from their eggs.

> [!quote]- A quote from Mordenkainen  
> 
> Other creatures that abide in hives serve a purpose in the natural world. Bees pollinate flowers. Termites make earth out of wood. Kruthiks, by contrast, slay societies. Perhaps that function is just as necessary.


```statblock
"name": "Kruthik Hive Lord (MTF)"
"size": "Large"
"type": "monstrosity"
"alignment": "Unaligned"
"ac": !!int "20"
"ac_class": "natural armor"
"hp": !!int "102"
"hit_dice": "12d10 + 36"
"stats":
- !!int "19"
- !!int "16"
- !!int "17"
- !!int "10"
- !!int "14"
- !!int "10"
"speed": "40 ft., burrow 20 ft., climb 40 ft."
"senses": "darkvision 60 ft., tremorsense 60 ft., passive Perception 12"
"languages": "Kruthik"
"cr": "5"
"traits":
- "desc": "The kruthik has advantage on Wisdom ([Perception](/3-Mechanics/CLI/rules/skills.md#Perception))\
    \ checks that rely on smell."
  "name": "Keen Smell"
- "desc": "The kruthik has advantage on an attack roll against a creature if at least\
    \ one of the kruthik's allies is within 5 feet of the creature and the ally isn't\
    \ [incapacitated](/3-Mechanics/CLI/rules/conditions.md#incapacitated)."
  "name": "Pack Tactics"
- "desc": "The kruthik can burrow through solid rock at half its burrowing speed and\
    \ leaves a 10-foot-diameter tunnel in its wake."
  "name": "Tunneler"
"actions":
- "desc": "The kruthik makes two stab attacks or two spike attacks."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+7 (+7) to hit, reach 10 ft., one target.\
    \ Hit: dice:1d10 + 4|text(9) (1d10 + 4) piercing damage."
  "name": "Stab"
- "desc": "Ranged Weapon Attack: dice: d20+6 (+6) to hit, range 30/120 ft.,\
    \ one target. Hit: dice:1d6 + 4|text(7) (1d6 + 4) piercing damage."
  "name": "Spike"
- "desc": "The kruthik sprays acid in a 15-foot cone. Each creature in that area must\
    \ make a DC 14 Dexterity saving throw, taking dice:4d10|text(22) (4d10) acid\
    \ damage on a failed save, or half as much damage on a successful one."
  "name": "Acid Spray (Recharge 5-6)"
"source":
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MTF/Kruthik%20Hive%20Lord.webp"
```
^statblock

```encounter-table
name: Kruthik Hive Lord
creatures:
 - 1: Kruthik Hive Lord
```

## Environment

desert, mountain, underdark