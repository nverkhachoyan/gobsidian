---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/12
- monster/environment/desert
- monster/environment/mountain
- monster/environment/urban
- monster/size/medium
- monster/type/humanoid/gith
aliases: ["Githyanki Kith'rak"]
NoteIcon: monster
BestiaryType: humanoid (gith)
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 140, Mordenkainen's Tome of Foes p. 205
---
# [Githyanki Kith'rak](3-Mechanics\CLI\bestiary\humanoid/githyanki-kithrak-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 140, Mordenkainen's Tome of Foes p. 205*  

```statblock
"name": "Githyanki Kith'rak (MPMM)"
"size": "Medium"
"type": "humanoid"
"subtype": "gith"
"alignment": "Any alignment"
"ac": !!int "18"
"ac_class": "[plate](/3-Mechanics/CLI/items/plate-armor.md)"
"hp": !!int "180"
"hit_dice": "24d8 + 72"
"stats":
- !!int "18"
- !!int "16"
- !!int "17"
- !!int "16"
- !!int "15"
- !!int "17"
"speed": "30 ft."
"saves":
  "Wisdom": !!int "6"
  "Intelligence": !!int "7"
  "Constitution": !!int "7"
"skillsaves":
  "Intimidation": !!int "7"
  "Perception": !!int "6"
"senses": "passive Perception 16"
"languages": "Gith"
"cr": "12"
"traits":
- "desc": "The githyanki casts one of the following spells, requiring no spell components\
    \ and using Intelligence as the spellcasting ability (spell save DC 15):\n\nAt\
    \ will: [mage hand](/3-Mechanics/CLI/spells/mage-hand.md) (the hand is invisible)\n\
    \n1/day each: [plane shift](/3-Mechanics/CLI/spells/plane-shift.md), [telekinesis](/3-Mechanics/CLI/spells/telekinesis.md)\n\
    \n3/day each: [blur](/3-Mechanics/CLI/spells/blur.md), [nondetection](/3-Mechanics/CLI/spells/nondetection.md)\
    \ (self only)"
  "name": "Spellcasting (Psionics)"
"actions":
- "desc": "The githyanki makes three Greatsword attacks."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+8 (+8) to hit, reach 5 ft., one target.\
    \ Hit: dice:2d6 + 4|text(11) (2d6 + 4) slashing damage plus dice:5d6|text(17)\
    \ (5d6) psychic damage."
  "name": "Greatsword"
"bonus_actions":
- "desc": "The githyanki teleports, along with any equipment it is wearing or carrying,\
    \ up to 30 feet to an unoccupied space it can see."
  "name": "Astral Step (Recharge 4-6)"
- "desc": "The githyanki magically ends the [charmed](/3-Mechanics/CLI/rules/conditions.md#charmed)\
    \ and [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened) conditions\
    \ on itself and each creature of its choice that it can see within 30 feet of\
    \ it."
  "name": "Rally the Troops"
"reactions":
- "desc": "The githyanki adds 4 to its AC against one melee attack that would hit\
    \ it. To do so, the githyanki must see the attacker and be wielding a melee weapon."
  "name": "Parry"
"source":
- "MPMM"
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Githyanki%20Kith%27rak.webp"
```
^statblock

```encounter-table
name: Githyanki Kith'rak
creatures:
 - 1: Githyanki Kith'rak
```

## Environment

desert, mountain, urban