---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/7
- monster/environment/underdark
- monster/size/large
- monster/type/fiend/demon
aliases: ["Armanite"]
NoteIcon: monster
BestiaryType: fiend (demon)
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 50, Mordenkainen's Tome of Foes p. 131
---
# [Armanite](3-Mechanics\CLI\bestiary\fiend/armanite-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 50, Mordenkainen's Tome of Foes p. 131*  

```statblock
"name": "Armanite (MPMM)"
"size": "Large"
"type": "fiend"
"subtype": "demon"
"alignment": "Typically  Chaotic Evil"
"ac": !!int "16"
"ac_class": "natural armor"
"hp": !!int "94"
"hit_dice": "9d10 + 45"
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
"actions":
- "desc": "The armanite makes one Claw attack, one Hooves attack, and one Serrated\
    \ Tail attack."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+8 (+8) to hit, reach 5 ft., one target.\
    \ Hit: dice:2d4 + 5|text(10) (2d4 + 5) slashing damage plus dice:2d8|text(9)\
    \ (2d8) lightning damage."
  "name": "Claw"
- "desc": "Melee Weapon Attack: dice: d20+8 (+8) to hit, reach 5 ft., one target.\
    \ Hit: dice:2d6 + 5|text(12) (2d6 + 5) bludgeoning damage. If the target\
    \ is a Large or smaller creature, it must succeed on a DC 16 Strength saving throw\
    \ or be knocked [prone](/3-Mechanics/CLI/rules/conditions.md#prone)."
  "name": "Hooves"
- "desc": "Melee Weapon Attack: dice: d20+8 (+8) to hit, reach 10 ft., one target.\
    \ Hit: dice:2d10 + 5|text(16) (2d10 + 5) slashing damage."
  "name": "Serrated Tail"
- "desc": "The armanite looses a bolt of lightning in a line that is 60 feet long\
    \ and 10 feet wide. Each creature in the line must make a DC 15 Dexterity saving\
    \ throw, taking dice:8d8|text(36) (8d8) lightning damage on a failed save,\
    \ or half as much damage on a successful one."
  "name": "Lightning Lance (Recharge 5-6)"
"source":
- "MPMM"
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Armanite.webp"
```
^statblock

```encounter-table
name: Armanite
creatures:
 - 1: Armanite
```

## Environment

underdark