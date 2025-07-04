---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/10
- monster/environment/desert
- monster/environment/forest
- monster/size/medium
- monster/type/fey/elf
aliases: ["Summer Eladrin"]
NoteIcon: monster
BestiaryType: fey (elf)
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 116, Mordenkainen's Tome of Foes p. 196
---
# [Summer Eladrin](3-Mechanics\CLI\bestiary\fey/summer-eladrin-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 116, Mordenkainen's Tome of Foes p. 196*  

```statblock
"name": "Summer Eladrin (MPMM)"
"size": "Medium"
"type": "fey"
"subtype": "elf"
"alignment": "Typically  Chaotic Neutral"
"ac": !!int "19"
"ac_class": "natural armor"
"hp": !!int "165"
"hit_dice": "22d8 + 66"
"stats":
- !!int "19"
- !!int "21"
- !!int "16"
- !!int "14"
- !!int "12"
- !!int "18"
"speed": "50 ft."
"skillsaves":
  "Intimidation": !!int "8"
  "Athletics": !!int "8"
"damage_resistances": "fire"
"senses": "darkvision 60 ft., passive Perception 11"
"languages": "Common, Elvish, Sylvan"
"cr": "10"
"traits":
- "desc": "Any non-eladrin creature that starts its turn within 60 feet of the eladrin\
    \ must make a DC 16 Wisdom saving throw. On a failed save, the creature becomes\
    \ [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened) of the eladrin\
    \ for 1 minute. A creature can repeat the saving throw at the end of each of its\
    \ turns, ending the effect on itself on a success. If a creature's saving throw\
    \ is successful or the effect ends for it, the creature is immune to any eladrin's\
    \ Fearsome Presence for the next 24 hours."
  "name": "Fearsome Presence"
- "desc": "The eladrin has advantage on saving throws against spells and other magical\
    \ effects."
  "name": "Magic Resistance"
"actions":
- "desc": "The eladrin makes two Longsword or Longbow attacks."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+8 (+8) to hit, reach 5 ft., one target.\
    \ Hit: dice:2d8 + 4|text(13) (2d8 + 4) slashing damage, or dice:2d10 +\
    \ 4|text(15) (2d10 + 4) slashing damage if used with two hands, plus dice:2d8|text(9)\
    \ (2d8) fire damage."
  "name": "Longsword"
- "desc": "Ranged Weapon Attack: dice: d20+9 (+9) to hit, range 150/600 ft.,\
    \ one target. Hit: dice:2d8 + 5|text(14) (2d8 + 5) piercing damage plus\
    \ dice:2d8|text(9) (2d8) fire damage."
  "name": "Longbow"
"bonus_actions":
- "desc": "The eladrin teleports, along with any equipment it is wearing or carrying,\
    \ up to 30 feet to an unoccupied space it can see."
  "name": "Fey Step (Recharge 4-6)"
"reactions":
- "desc": "The eladrin adds 3 to its AC against one melee attack that would hit it.\
    \ To do so, the eladrin must see the attacker and be wielding a melee weapon."
  "name": "Parry"
"source":
- "MPMM"
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Summer%20Eladrin.webp"
```
^statblock

```encounter-table
name: Summer Eladrin
creatures:
 - 1: Summer Eladrin
```

## Environment

desert, forest