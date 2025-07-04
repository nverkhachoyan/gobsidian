---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/7
- monster/environment/arctic
- monster/size/medium
- monster/type/fey
aliases: ["Bheur Hag"]
NoteIcon: monster
BestiaryType: fey
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 62, Volo's Guide to Monsters p. 160
---
# [Bheur Hag](3-Mechanics\CLI\bestiary\fey/bheur-hag-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 62, Volo's Guide to Monsters p. 160*  

```statblock
"name": "Bheur Hag (MPMM)"
"size": "Medium"
"type": "fey"
"alignment": "Typically  Chaotic Evil"
"ac": !!int "17"
"ac_class": "natural armor"
"hp": !!int "91"
"hit_dice": "14d8 + 28"
"stats":
- !!int "13"
- !!int "16"
- !!int "14"
- !!int "12"
- !!int "13"
- !!int "16"
"speed": "30 ft., fly 50 ft. (hover, Graystaff magic)"
"saves":
  "Wisdom": !!int "4"
"skillsaves":
  "Nature": !!int "4"
  "Stealth": !!int "6"
  "Perception": !!int "4"
  "Survival": !!int "4"
"damage_immunities": "cold"
"senses": "darkvision 60 ft., passive Perception 14"
"languages": "Auran, Common, Giant"
"cr": "7"
"traits":
- "desc": "While holding or riding the graystaff, the hag casts one of the following\
    \ spells, requiring no material components and using Charisma as the spellcasting\
    \ ability (spell save DC 14):\n\nAt will: [hold person](/3-Mechanics/CLI/spells/hold-person.md)\n\
    \n1/day each: [cone of cold](/3-Mechanics/CLI/spells/cone-of-cold.md), [ice\
    \ storm](/3-Mechanics/CLI/spells/ice-storm.md), [wall of ice](/3-Mechanics/CLI/spells/wall-of-ice.md)"
  "name": "Spellcasting"
- "desc": "The hag can cast the [control weather](/3-Mechanics/CLI/spells/control-weather.md)\
    \ spell, requiring no material components and using Charisma as the spellcasting\
    \ ability."
  "name": "Control Weather (1/Day)"
- "desc": "The hag carries a graystaff, a magic staff. The hag can use its flying\
    \ speed only while astride the staff. If the staff is lost or destroyed, the hag\
    \ must craft another, which takes a year and a day. Only a bheur hag can use a\
    \ graystaff"
  "name": "Graystaff Magic"
- "desc": "The hag can move across and climb icy surfaces without needing to make\
    \ an ability check, and difficult terrain composed of ice or snow doesn't cost\
    \ the hag extra moment."
  "name": "Ice Walk"
"actions":
- "desc": "The hag makes two Slam or Frost Shard attacks."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+4 (+4) to hit, reach 5 ft., one target.\
    \ Hit: dice:2d8 + 1|text(10) (2d8 + 1) bludgeoning damage plus dice:4d8|text(18)\
    \ (4d8) cold damage."
  "name": "Slam"
- "desc": "Ranged Spell Attack: dice: d20+6 (+6) to hit, range 60 ft., one target.\
    \ Hit: dice:6d8 + 3|text(30) (6d8 + 3) cold damage, and the target's speed\
    \ is reduced by 10 feet until the start of the hag's next turn."
  "name": "Frost Shard"
- "desc": "The hag feeds on the corpse of one enemy within reach that died within\
    \ the past minute. Each creature of the hag's choice that is within 60 feet and\
    \ able to see the feeding must succeed on a DC 15 Wisdom saving throw or be [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened)\
    \ of the hag for 1 minute. While [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened)\
    \ in this way, a creature is [incapacitated](/3-Mechanics/CLI/rules/conditions.md#incapacitated),\
    \ can't understand what others say, can't read, and speaks only in gibberish.\
    \ A creature can repeat the saving throw at the end of each of its turns, ending\
    \ the effect on itself on a success. If a creature's saving throw is successful\
    \ or the effect ends for it, the creature is immune to the hag's Horrific Feast\
    \ for the next 24 hours."
  "name": "Horrific Feast"
"source":
- "MPMM"
- "VGM"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Bheur%20Hag.webp"
```
^statblock

```encounter-table
name: Bheur Hag
creatures:
 - 1: Bheur Hag
```

## Environment

arctic