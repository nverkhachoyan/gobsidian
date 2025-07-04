---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/19
- monster/environment/mountain
- monster/environment/urban
- monster/size/medium
- monster/type/fiend/devil
aliases: ["Red Abishai"]
NoteIcon: monster
BestiaryType: fiend (devil)
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 40, Mordenkainen's Tome of Foes p. 160
---
# [Red Abishai](3-Mechanics\CLI\bestiary\fiend/red-abishai-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 40, Mordenkainen's Tome of Foes p. 160*  

```statblock
"name": "Red Abishai (MPMM)"
"size": "Medium"
"type": "fiend"
"subtype": "devil"
"alignment": "Typically  Lawful Evil"
"ac": !!int "22"
"ac_class": "natural armor"
"hp": !!int "289"
"hit_dice": "34d8 + 136"
"stats":
- !!int "23"
- !!int "16"
- !!int "19"
- !!int "14"
- !!int "15"
- !!int "19"
"speed": "30 ft., fly 50 ft."
"saves":
  "Wisdom": !!int "8"
  "Strength": !!int "12"
  "Constitution": !!int "10"
"skillsaves":
  "Intimidation": !!int "10"
  "Perception": !!int "8"
"damage_resistances": "cold; bludgeoning, piercing, slashing from nonmagical attacks\
  \ that aren't silvered"
"damage_immunities": "fire, poison"
"condition_immunities": "[frightened](/3-Mechanics/CLI/rules/conditions.md#frightened),\
  \ [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)"
"senses": "darkvision 120 ft., passive Perception 18"
"languages": "Draconic, Infernal, telepathy 120 ft."
"cr": "19"
"traits":
- "desc": "Magical darkness doesn't impede the abishai's [darkvision](/3-Mechanics/CLI/rules/senses.md#darkvision)."
  "name": "Devil's Sight"
- "desc": "The abishai has advantage on saving throws against spells and other magical\
    \ effects."
  "name": "Magic Resistance"
"actions":
- "desc": "The abishai makes one Bite attack and one Claw attack, and it can use Frightful\
    \ Presence or Incite Fanaticism."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+12 (+12) to hit, reach 5 ft., one\
    \ target. Hit: dice:3d10 + 6|text(22) (3d10 + 6) piercing damage plus dice:7d10|text(38)\
    \ (7d10) fire damage."
  "name": "Bite"
- "desc": "Melee Weapon Attack: dice: d20+12 (+12) to hit, reach 5 ft., one\
    \ target. Hit: dice:2d10 + 6|text(17) (2d10 + 6) force damage plus dice:2d10|text(11)\
    \ (2d10) fire damage."
  "name": "Claw"
- "desc": "Each creature of the abishai's choice that is within 120 feet and aware\
    \ of the abishai must succeed on a DC 18 Wisdom saving throw or become [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened)\
    \ of it for 1 minute. A creature can repeat the saving throw at the end of each\
    \ of its turns, ending the effect on itself on a success. If a creature's saving\
    \ throw is successful or the effect ends for it, the creature is immune to the\
    \ abishai's Frightful Presence for the next 24 hours."
  "name": "Frightful Presence"
- "desc": "The abishai chooses up to four other creatures within 60 feet of it that\
    \ can see it. Until the start of the abishai's next turn, each of those creatures\
    \ makes attack rolls with advantage and can't be [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened)."
  "name": "Incite Fanaticism"
- "desc": "The abishai targets one Dragon it can see within 120 feet of it. The Dragon\
    \ must make a DC 18 Charisma saving throw. A chromatic dragon makes this save\
    \ with disadvantage. On a successful save, the target is immune to the abishai's\
    \ Power of the Dragon Queen for 1 hour. On a failed save, the target is [charmed](/3-Mechanics/CLI/rules/conditions.md#charmed)\
    \ by the abishai for 1 hour. While [charmed](/3-Mechanics/CLI/rules/conditions.md#charmed)\
    \ in this way, the target regards the abishai as a trusted friend to be heeded\
    \ and protected. This effect ends if the abishai or its companions deal damage\
    \ to the target."
  "name": "Power of the Dragon Queen"
"source":
- "MPMM"
- "MTF"
- "VEoR"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Red%20Abishai.webp"
```
^statblock

```encounter-table
name: Red Abishai
creatures:
 - 1: Red Abishai
```

## Environment

mountain, urban