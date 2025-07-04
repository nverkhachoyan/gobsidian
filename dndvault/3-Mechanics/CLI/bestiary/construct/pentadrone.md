---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mm
- monster/cr/2
- monster/size/large
- monster/type/construct
aliases: ["Pentadrone"]
NoteIcon: monster
BestiaryType: construct
SourceType: Bestiary
BookSource: Monster Manual p. 226
---
# [Pentadrone](3-Mechanics\CLI\bestiary\construct/pentadrone.md)
*Source: Monster Manual p. 226*  

Pentadrones oversee Mechanus's worker populace and can improvise in response to new situations.

## Modrons

Modrons are beings of absolute law that adhere to a hive-like hierarchy. They inhabit the plane of Mechanus and tend its eternally revolving gears, their existence a clockwork routine of perfect order.

### Absolute Law and Order

Under the direction of their leader, Primus, modrons increase order in the multiverse in accordance with laws beyond the comprehension of mortal minds. Their own minds are networked in a hierarchal pyramid, in which each modron receives commands from superiors and delegates orders to underlings. A modron carries out commands with total obedience, utmost efficiency, and an absence of morality or ego. Modrons have no sense of self beyond what is necessary to fulfill their duties. They exist as a unified collective, divided by ranks, yet they always refer to themselves collectively. To a modron, there is no "I," but only "we" or "us."

> [!note] Variant: Rogue Modrons
> 
> A modron unit sometimes becomes defective, either through natural decay or exposure to chaotic forces. Rogue modrons don't act in accordance with Primus's wishes and directives, breaking laws, disobeying orders, and even engaging in violence. Other modrons hunt down such rogues.
> 
> A rogue modron loses the Axiomatic Mind trait and can have any alignment other than lawful neutral. Otherwise, it has the same statistics as a regular modron of its rank.
^variant-rogue-modrons

### Absolute Hierarchy

Modrons communicate only with their own rank and the ranks immediately above and below them. Modrons more than one rank away are either too advanced or too simple to understand.

### Cogs of the Great Machine

If a modron is destroyed, its remains disintegrate. A replacement from the next lowest rank then transforms in a flash of light, gaining the physical form of its new rank. The promoted modron is replaced by one of its underlings in the same manner, all the way to the lowest levels of the hierarchy. There, a new modron is created by Primus, with a steady stream of monodrones leaving the Great Modron Cathedral on Mechanus as a result.

### The Great Modron March

When the gears of Mechanus complete seventeen cycles once every 289 years, Primus sends a vast army of modrons across the Outer Planes, ostensibly on a reconnaissance mission. The march is long and dangerous, and only a small number of modrons returns to Mechanus.

> [!quote]- A quote from Kwint Stormbellow, rock gnome adventurer  
> 
> Every 289 years, the entire multiverse goes mad. Like clockwork.


```statblock
"name": "Pentadrone"
"size": "Large"
"type": "construct"
"alignment": "Lawful Neutral"
"ac": !!int "16"
"ac_class": "natural armor"
"hp": !!int "32"
"hit_dice": "5d10 + 5"
"stats":
- !!int "15"
- !!int "14"
- !!int "12"
- !!int "10"
- !!int "10"
- !!int "13"
"speed": "40 ft."
"skillsaves":
  "Perception": !!int "4"
"senses": "truesight 120 ft., passive Perception 14"
"languages": "Modron"
"cr": "2"
"traits":
- "desc": "The pentadrone can't be compelled to act in a manner contrary to its nature\
    \ or its instructions."
  "name": "Axiomatic Mind"
- "desc": "If the pentadrone dies, its body disintegrates into dust, leaving behind\
    \ its weapons and anything else it was carrying."
  "name": "Disintegration"
"actions":
- "desc": "The pentadrone makes five arm attacks."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+4 (+4) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d6 + 2|text(5) (1d6 + 2) bludgeoning damage."
  "name": "Arm"
- "desc": "The pentadrone exhales a 30-foot cone of gas. Each creature in that area\
    \ must succeed on a DC 11 Constitution saving throw or be [paralyzed](/3-Mechanics/CLI/rules/conditions.md#paralyzed)\
    \ for 1 minute. A creature can repeat the saving throw at the end of each of its\
    \ turns, ending the effect on itself on a success."
  "name": "Paralysis Gas (Recharge 5-6)"
"source":
- "MM"
- "ToA"
- "KftGV"
- "SatO"
- "ToFW"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MM/Pentadrone.webp"
```
^statblock

```encounter-table
name: Pentadrone
creatures:
 - 1: Pentadrone
```