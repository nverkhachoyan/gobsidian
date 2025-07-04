---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/sdw
- monster/cr/
- monster/size/medium
- monster/type/humanoid
aliases: ["Warrior"]
NoteIcon: monster
BestiaryType: humanoid
SourceType: Bestiary
BookSource: Sleeping Dragon's Wake
---
# [Warrior](3-Mechanics\CLI\bestiary\humanoid/warrior-sdw.md)
*Source: Sleeping Dragon's Wake*  

```statblock
"name": "Warrior (SDW)"
"size": "Medium"
"type": "humanoid"
"alignment": "Unaligned"
"ac": !!int "20"
"ac_class": "[plate](/3-Mechanics/CLI/items/plate-armor.md), [shield](/3-Mechanics/CLI/items/shield.md)"
"hp": !!int "65"
"hit_dice": "10d8 + 20"
"stats":
- !!int "18"
- !!int "14"
- !!int "14"
- !!int "10"
- !!int "12"
- !!int "10"
"speed": "30 ft."
"saves":
  "Constitution": !!int "6"
"skillsaves":
  "Athletics": !!int "8"
  "Perception": !!int "5"
  "Survival": !!int "5"
"senses": "passive Perception 15"
"languages": "Common, plus one of your choice"
"traits":
- "desc": "The warrior has advantage on initiative rolls."
  "name": "Battle Readiness"
- "desc": "The warrior's attack rolls score a critical hit on a roll of 19 or 20 on\
    \ the dice: d20|avg|noform (d20)."
  "name": "Improved Critical"
- "desc": "The warriorcan reroll a saving throw that it fails, but it must use the\
    \ new result."
  "name": "Indomitable (1/Day)"
- "desc": "The warrior has one of the following traits of your choice:\n\n- Attacker.\
    \ The warrior gains a +2 bonus to attack rolls.  \n- Defender. The warrior\
    \ gains the Protection reaction below.  "
  "name": "Martial Role"
- "desc": "The warrior can use a bonus action on its turn to regain hit points equal\
    \ to dice: 1d10|avg|noform (1d10) + its level."
  "name": "Second Wind (Recharges after a Short or Long Rest)"
"actions":
- "desc": "The warrior can attack twice, instead of once, whenever it takes the attack\
    \ action on its turn."
  "name": "Extra Attack"
- "desc": "Melee Weapon Attack: dice: d20+8 (+8) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d8 + 4|text(8) (1d8 + 4) slashing damage, or dice:1d10 + 4|text(9)\
    \ (1d10 + 4) slashing damage if used with two hands."
  "name": "Longsword"
- "desc": "Ranged Weapon Attack: dice: d20+6 (+6) to hit, range 150/600 ft.,\
    \ one target. Hit: dice:1d8 + 2|text(6) (1d8 + 2) piercing damage."
  "name": "Longbow"
"reactions":
- "desc": "When a creature the warrior can see attacks a target other than the warrior\
    \ that is within 5 feet of the warrior, the warrior can use their reaction to\
    \ impose disadvantage on the attack roll. The warrior must be wielding a shield."
  "name": "Protection (Defender Only)"
"source":
- "SDW"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/SDW/Warrior.webp"
```
^statblock

```encounter-table
name: Warrior
creatures:
 - 1: Warrior
```