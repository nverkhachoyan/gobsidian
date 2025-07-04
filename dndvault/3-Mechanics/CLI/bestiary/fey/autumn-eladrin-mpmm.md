---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/10
- monster/environment/forest
- monster/size/medium
- monster/type/fey/elf
aliases: ["Autumn Eladrin"]
NoteIcon: monster
BestiaryType: fey (elf)
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 115, Mordenkainen's Tome of Foes p. 195
---
# [Autumn Eladrin](3-Mechanics\CLI\bestiary\fey/autumn-eladrin-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 115, Mordenkainen's Tome of Foes p. 195*  

```statblock
"name": "Autumn Eladrin (MPMM)"
"size": "Medium"
"type": "fey"
"subtype": "elf"
"alignment": "Typically  Chaotic Neutral"
"ac": !!int "19"
"ac_class": "natural armor"
"hp": !!int "165"
"hit_dice": "22d8 + 66"
"stats":
- !!int "12"
- !!int "16"
- !!int "16"
- !!int "14"
- !!int "17"
- !!int "18"
"speed": "30 ft."
"skillsaves":
  "Medicine": !!int "7"
  "Insight": !!int "7"
"damage_resistances": "psychic"
"senses": "darkvision 60 ft., passive Perception 13"
"languages": "Common, Elvish, Sylvan"
"cr": "10"
"traits":
- "desc": "The eladrin casts one of the following spells, requiring no material components\
    \ and using Charisma as the spellcasting ability (spell save DC 16):\n\nAt will:\
    \ [hold person](/3-Mechanics/CLI/spells/hold-person.md)\n\n1/day each: [greater\
    \ restoration](/3-Mechanics/CLI/spells/greater-restoration.md), [revivify](/3-Mechanics/CLI/spells/revivify.md)\n\
    \n2/day each: [cure wounds](/3-Mechanics/CLI/spells/cure-wounds.md) (as a\
    \ 5th-level spell), [lesser restoration](/3-Mechanics/CLI/spells/lesser-restoration.md)"
  "name": "Spellcasting"
- "desc": "Any non-eladrin creature that starts its turn within 60 feet of the eladrin\
    \ must make a DC 16 Wisdom saving throw. On a failed save, the creature becomes\
    \ [charmed](/3-Mechanics/CLI/rules/conditions.md#charmed) by the eladrin for 1\
    \ minute. On a successful save, the creature becomes immune to any eladrin's Enchanting\
    \ Presence for 24 hours.\n\nWhenever the eladrin deals damage to the [charmed](/3-Mechanics/CLI/rules/conditions.md#charmed)\
    \ creature, the [charmed](/3-Mechanics/CLI/rules/conditions.md#charmed) creature\
    \ can repeat the saving throw, ending the effect on itself on a success."
  "name": "Enchanting Presence"
- "desc": "The eladrin has advantage on saving throws against spells and other magical\
    \ effects."
  "name": "Magic Resistance"
"actions":
- "desc": "The eladrin makes two Longsword or Longbow attacks. It can replace one\
    \ attack with a use of Spellcasting."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+5 (+5) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d8 + 1|text(5) (1d8 + 1) slashing damage, or dice:1d10 + 1|text(6)\
    \ (1d10 + 1) slashing damage if used with two hands, plus dice:5d8|text(22)\
    \ (5d8) psychic damage."
  "name": "Longsword"
- "desc": "Ranged Weapon Attack: dice: d20+7 (+7) to hit, range 150/600 ft.,\
    \ one target. Hit: dice:1d8 + 3|text(7) (1d8 + 3) piercing damage plus dice:5d8|text(22)\
    \ (5d8) psychic damage."
  "name": "Longbow"
"bonus_actions":
- "desc": "The eladrin teleports, along with any equipment it is wearing or carrying,\
    \ up to 30 feet to an unoccupied space it can see."
  "name": "Fey Step (Recharge 4-6)"
"reactions":
- "desc": "If a creature [charmed](/3-Mechanics/CLI/rules/conditions.md#charmed) by\
    \ the eladrin hits with an attack roll while within 60 feet of the eladrin, the\
    \ eladrin magically causes the attack to miss, provided the eladrin can see the\
    \ attacker."
  "name": "Foster Peace"
"source":
- "MPMM"
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Autumn%20Eladrin.webp"
```
^statblock

```encounter-table
name: Autumn Eladrin
creatures:
 - 1: Autumn Eladrin
```

## Environment

forest