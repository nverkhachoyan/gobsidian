---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mtf
- monster/cr/7
- monster/environment/underdark
- monster/size/large
- monster/type/fiend/demon
aliases: ["Armanite"]
NoteIcon: monster
BestiaryType: fiend (demon)
SourceType: Bestiary
BookSource: Mordenkainen's Tome of Foes p. 131
---
# [Armanite](3-Mechanics\CLI\bestiary\fiend/armanite-mtf.md)
*Source: Mordenkainen's Tome of Foes p. 131*  

## Armanite

Great herds of armanites race across the blasted fields of the Abyss, bent on slaughter and death, driven by unrestrained bloodlust. Whether being controlled by more powerful demons or charging into battle for the sake of it, armanites use their claws, hooves, and long, whip-like tails to tear apart their foes.

## Live for War

In the armies of the demon lords, armanites perform the role of heavy cavalry, leading the charge and tearing into their enemies' flanks. Armanites fight all the time, even among themselves if they can't find another enemy. They make ideal shock troops, courageous to the point of stupidity and utterly savage.

## Walking Arsenal

Part of what makes armanites so fearsome is the number of weapons they have at their disposal. They possess sharp hooves, claws that end in curling talons, and long, serrated tails that can flense the flesh from a victim, and they use them all to carve through their foes. When they are up against tough formations, they can call on their innate magic to loose bolts of lightning and blow holes in the enemy ranks.

```statblock
"name": "Armanite (MTF)"
"size": "Large"
"type": "fiend"
"subtype": "demon"
"alignment": "Chaotic Evil"
"ac": !!int "16"
"ac_class": "natural armor"
"hp": !!int "84"
"hit_dice": "8d10 + 40"
"stats":
- !!int "21"
- !!int "18"
- !!int "21"
- !!int "8"
- !!int "12"
- !!int "13"
"speed": "60 ft."
"damage_resistances": "cold, fire, lightning"
"damage_immunities": "poison"
"condition_immunities": "[poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)"
"senses": "darkvision 120 ft., passive Perception 11"
"languages": "Abyssal, telepathy 120 ft."
"cr": "7"
"traits":
- "desc": "The armanite has advantage on saving throws against spells and other magical\
    \ effects."
  "name": "Magic Resistance"
- "desc": "The armanite's weapon attacks are magical."
  "name": "Magic Weapons"
"actions":
- "desc": "The armanite makes three attacks: one with its hooves, one with its claws,\
    \ and one with its serrated tail."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+8 (+8) to hit, reach 5 ft., one target.\
    \ Hit: dice:2d6 + 5|text(12) (2d6 + 5) bludgeoning damage."
  "name": "Hooves"
- "desc": "Melee Weapon Attack: dice: d20+8 (+8) to hit, reach 5 ft., one target.\
    \ Hit: dice:2d4 + 5|text(10) (2d4 + 5) slashing damage."
  "name": "Claws"
- "desc": "Melee Weapon Attack: dice: d20+8 (+8) to hit, reach 10 ft., one target.\
    \ Hit: dice:2d10 + 5|text(16) (2d10 + 5) slashing damage."
  "name": "Serrated Tail"
- "desc": "The armanite looses a bolt of lightning in a line 60 feet long and 10 feet\
    \ wide. Each creature in the line must make a DC 15 Dexterity saving throw, taking\
    \ dice:6d8|text(27) (6d8) lightning damage on a failed save, or half as much\
    \ damage on a successful one."
  "name": "Lightning Lance (Recharge 5-6)"
"source":
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MTF/Armanite.webp"
```
^statblock

```encounter-table
name: Armanite
creatures:
 - 1: Armanite
```

## Environment

underdark