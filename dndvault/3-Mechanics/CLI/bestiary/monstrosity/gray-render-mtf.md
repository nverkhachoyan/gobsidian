---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mtf
- monster/cr/12
- monster/environment/forest
- monster/environment/hill
- monster/size/large
- monster/type/monstrosity
aliases: ["Gray Render"]
NoteIcon: monster
BestiaryType: monstrosity
SourceType: Bestiary
BookSource: Mordenkainen's Tome of Foes p. 209
---
# [Gray Render](3-Mechanics\CLI\bestiary\monstrosity/gray-render-mtf.md)
*Source: Mordenkainen's Tome of Foes p. 209*  

> [!quote]- A quote from Mordenkainen  
> 
> I suspect gray renders owe their origin to the neogi, since they are often in their company. Those that appear in the wilderness are likely castaways from frustrated neogi masters.

## Gray Render

A curious impulse drives the gray render. Despite its hulking form and terrible appetite, it wants most of all to bond with an intelligent creature and, once bonded, give its life to protect that creature. Great strength and a savage nature enable gray renders to be fierce guardians, but they lack even a shred of cunning.

## A Spreading Plague

Gray renders reproduce by forming nodules on their bodies that, upon reaching maturity, break off to begin life as young gray renders. These monstrosities feel no obligation to their young, and they have no inclination to gather with others of their kind.

## Chaotic Allies

As a side effect of its breeding, each gray render has an overpowering need to bond with an intelligent creature. When it encounters a suitable master, the render sings to it—a weird, warbling cry accompanied by scratching at the earth and a show of deference. Once it forms the bond, the render serves its master in all things.

Although this bond can be a great benefit, renders are inherently chaotic. In a battle, a render fights with all the savagery it can muster and never willingly harms its master, but outside battle, a gray render might present considerable difficulty for its master's associates. It might follow its master even after being told to stay put, destroy its master's house, burrow holes in the side of a ship, kill horses, attack when it feels jealous, and more. A gray render might be a boon companion, but it is always an unpredictable one.

The Gray Render Quirks table presents possible quirks for gray renders that can be generated randomly or selected as desired.

**Grey Render Quirk**

`dice: [](gray-render-mtf.md#^grey-render-quirk)`

| dice: d12 | Quirk |
|-----------|-------|
| 1 | Hates horses and other mounts |
| 2 | Roars loudly when its bonded creature is touched by another creature |
| 3 | Likes to snuggle |
| 4 | Uproots and chews on trees |
| 5 | Has terrific and eye-watering flatulence |
| 6 | Brings offerings of meat to its bonded creature |
| 7 | Compulsively digs up the ground |
| 8 | Attacks carts and wagons as if they were terrible monsters |
| 9 | Howls when it rains |
| 10 | Whines piteously in the dark |
| 11 | Buries treasure it finds |
| 12 | Chases birds, leaping into the air to catch them, heedless of the destruction it causes |
^grey-render-quirk

```statblock
"name": "Gray Render (MTF)"
"size": "Large"
"type": "monstrosity"
"alignment": "Chaotic Neutral"
"ac": !!int "19"
"ac_class": "natural armor"
"hp": !!int "189"
"hit_dice": "18d10 + 90"
"stats":
- !!int "19"
- !!int "13"
- !!int "20"
- !!int "3"
- !!int "6"
- !!int "8"
"speed": "30 ft."
"saves":
  "Strength": !!int "8"
  "Constitution": !!int "9"
"skillsaves":
  "Perception": !!int "2"
"senses": "darkvision 60 ft., passive Perception 12"
"languages": ""
"cr": "12"
"actions":
- "desc": "The gray render makes three attacks: one with its bite and two with its\
    \ claws."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+8 (+8) to hit, reach 5 ft., one target.\
    \ Hit: dice:2d12 + 4|text(17) (2d12 + 4) piercing damage. If the target\
    \ is Medium or smaller, the target must succeed on a DC 16 Strength saving throw\
    \ or be knocked [prone](/3-Mechanics/CLI/rules/conditions.md#prone)."
  "name": "Bite"
- "desc": "Melee Weapon Attack: dice: d20+8 (+8) to hit, reach 10 ft., one target.\
    \ Hit: dice:2d8 + 4|text(13) (2d8 + 4) slashing damage. If the target is\
    \ [prone](/3-Mechanics/CLI/rules/conditions.md#prone) an additional dice:2d6|text(7)\
    \ (2d6) bludgeoning damage is dealt to the target."
  "name": "Claws"
"reactions":
- "desc": "When the gray render takes damage, it makes one attack with its claws against\
    \ a random creature within its reach, other than its master."
  "name": "Bloody Rampage"
"source":
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MTF/Gray%20Render.webp"
```
^statblock

```encounter-table
name: Gray Render
creatures:
 - 1: Gray Render
```

## Environment

forest, hill