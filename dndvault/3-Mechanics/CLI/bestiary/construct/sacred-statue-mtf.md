---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mtf
- monster/cr/
- monster/size/large
- monster/type/construct
aliases: ["Sacred Statue"]
NoteIcon: monster
BestiaryType: construct
SourceType: Bestiary
BookSource: Mordenkainen's Tome of Foes p. 194, Divine Contention, Dragon of Icespire Peak
---
# [Sacred Statue](3-Mechanics\CLI\bestiary\construct/sacred-statue-mtf.md)
*Source: Mordenkainen's Tome of Foes p. 194, Divine Contention, Dragon of Icespire Peak*  

> [!quote]- A quote from Mordenkainen  
> 
> It's not just gods that have the power to bind spirits to their idols. Beings such as archdevils can do it with the souls of their cultists. Moloch, for instance.

## Sacred Statue

See [Eidolon](/3-Mechanics/CLI/bestiary/undead/eidolon-mtf.md) for details

```statblock
"name": "Sacred Statue (MTF)"
"size": "Large"
"type": "construct"
"alignment": "as the eidolon's alignment"
"ac": !!int "19"
"ac_class": "natural armor"
"hp": !!int "95"
"hit_dice": "10d10 + 40"
"stats":
- !!int "19"
- !!int "8"
- !!int "19"
- !!int "14"
- !!int "19"
- !!int "16"
"speed": "25 ft."
"saves":
  "Wisdom": !!int "8"
"damage_resistances": "acid; fire; lightning; bludgeoning, piercing, slashing from\
  \ nonmagical attacks"
"damage_immunities": "cold, necrotic, poison"
"condition_immunities": "[charmed](/3-Mechanics/CLI/rules/conditions.md#charmed),\
  \ [exhaustion](/3-Mechanics/CLI/rules/conditions.md#exhaustion), [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened),\
  \ [paralyzed](/3-Mechanics/CLI/rules/conditions.md#paralyzed), [petrified](/3-Mechanics/CLI/rules/conditions.md#petrified),\
  \ [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)"
"senses": "darkvision 60 ft., passive Perception 14"
"languages": "the languages the eidolon knew in life"
"traits":
- "desc": "While the statue remains motionless, it is indistinguishable from a normal\
    \ statue."
  "name": "False Appearance"
- "desc": "The [eidolon](/3-Mechanics/CLI/bestiary/undead/eidolon-mtf.md) that enters\
    \ the sacred statue remains inside it until the statue drops to 0 hit points,\
    \ the eidolon uses a bonus action to move out of the statue, or the eidolon is\
    \ turned or forced out by an effect such as the [dispel evil and good](/3-Mechanics/CLI/spells/dispel-evil-and-good.md)\
    \ spell. When the eidolon leaves the statue, it appears in an unoccupied space\
    \ within 5 feet of the statue."
  "name": "Ghostly Inhabitant"
- "desc": "When not inhabited by an [eidolon](/3-Mechanics/CLI/bestiary/undead/eidolon-mtf.md),\
    \ the statue is an object."
  "name": "Inert"
"actions":
- "desc": "The statue makes two slam attacks."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+8 (+8) to hit, reach 10 ft., one target.\
    \ Hit: dice:6d12 + 4|text(43) (6d12 + 4) bludgeoning damage."
  "name": "Slam"
- "desc": "Ranged Weapon Attack: dice: d20+8 (+8) to hit, range 60/240 ft.,\
    \ one target. Hit: dice:6d10 + 4|text(37) (6d10 + 4) bludgeoning damage."
  "name": "Rock"
"source":
- "MTF"
- "DC"
- "DIP"
- "IMR"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MTF/Sacred%20Statue.webp"
```
^statblock

```encounter-table
name: Sacred Statue
creatures:
 - 1: Sacred Statue
```