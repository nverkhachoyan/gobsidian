---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/dip
- monster/cr/6
- monster/size/huge
- monster/type/monstrosity
aliases: ["Gorthok the Thunder Boar"]
NoteIcon: monster
BestiaryType: monstrosity
SourceType: Bestiary
BookSource: Dragon of Icespire Peak p. 58
---
# [Gorthok the Thunder Boar](3-Mechanics\CLI\bestiary\npc/gorthok-the-thunder-boar-dip.md)
*Source: Dragon of Icespire Peak p. 58*  

Gorthok is a primal nature spirit that takes the form of a boar as big as an elephant, with lightning that dances along its tusks. Gorthok serves the will of Talos, god of storms, and can be summoned during stormy weather to do the bidding of Talos's evil followers. Like its patron deity, Gorthok revels in destruction.

```statblock
"name": "Gorthok the Thunder Boar (DIP)"
"size": "Huge"
"type": "monstrosity"
"alignment": "Chaotic Evil"
"ac": !!int "15"
"ac_class": "natural armor"
"hp": !!int "73"
"hit_dice": "7d12 + 28"
"stats":
- !!int "20"
- !!int "11"
- !!int "19"
- !!int "6"
- !!int "10"
- !!int "14"
"speed": "50 ft."
"damage_resistances": "bludgeoning, piercing, slashing from nonmagical attacks"
"damage_immunities": "lightning, thunder"
"senses": "darkvision 60 ft., passive Perception 10"
"languages": ""
"cr": "6"
"traits":
- "desc": "If Gorthok takes 27 damage or less that would reduce it to 0 hit points,\
    \ it is reduced to 1 hit point instead."
  "name": "Relentless (Recharges after a Short or Long Rest)"
"actions":
- "desc": "Gorthok makes two melee attacks: one with its lightning tusks and one with\
    \ its thunder hooves."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+8 (+8) to hit, reach 10 ft., one target.\
    \ Hit: dice:2d6 + 5|text(12) (2d6 + 5) slashing damage plus dice:2d6|text(7)\
    \ (2d6) lightning damage."
  "name": "Lightning Tusks"
- "desc": "Melee Weapon Attack: dice: d20+8 (+8) to hit, reach 10 ft., one target.\
    \ Hit: dice:2d6 + 5|text(12) (2d6 + 5) bludgeoning damage plus dice:2d6|text(7)\
    \ (2d6) thunder damage."
  "name": "Thunder Hooves"
- "desc": "Gorthok shoots a bolt of lightning at one creature it can see within 120\
    \ feet of it. The target must make a DC 15 Dexterity saving throw, taking dice:4d8|text(18)\
    \ (4d8) lightning damage on a failed save, or half as much damage on a successful\
    \ one."
  "name": "Lightning Bolt (Recharge 6)"
"source":
- "DIP"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/DIP/Gorthok%20the%20Thunder%20Boar.webp"
```
^statblock

```encounter-table
name: Gorthok the Thunder Boar
creatures:
 - 1: Gorthok the Thunder Boar
```