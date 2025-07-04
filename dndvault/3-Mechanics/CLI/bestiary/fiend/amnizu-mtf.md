---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mtf
- monster/cr/18
- monster/size/medium
- monster/type/fiend/devil
aliases: ["Amnizu"]
NoteIcon: monster
BestiaryType: fiend (devil)
SourceType: Bestiary
BookSource: Mordenkainen's Tome of Foes p. 164
---
# [Amnizu](3-Mechanics\CLI\bestiary\fiend/amnizu-mtf.md)
*Source: Mordenkainen's Tome of Foes p. 164*  

Amnizus lead the infernal legions into battle and command guardians at the gateways to the Hells. Amnizus are arrogant, bullying, and ruthless, but they're also highly intelligent tacticians and unfailingly loyal—qualities that the hellish archdukes value.

## Guarding the River Styx

Some amnizus perform the critical task of watching over the River Styx from fortresses along the river's blighted banks, where it flows through Dis and Stygia. Souls arriving in the form of [lemures](/3-Mechanics/CLI/bestiary/fiend/lemure.md) have no personalities or memories; they're driven only by the desire to commit evil. The amnizus that patrol here drill the rules of the Nine Hells into the new arrivals' pitiful brains and marshal the lemures into legions.

```statblock
"name": "Amnizu (MTF)"
"size": "Medium"
"type": "fiend"
"subtype": "devil"
"alignment": "Lawful Evil"
"ac": !!int "21"
"ac_class": "natural armor"
"hp": !!int "202"
"hit_dice": "27d8 + 81"
"stats":
- !!int "11"
- !!int "13"
- !!int "16"
- !!int "20"
- !!int "12"
- !!int "18"
"speed": "30 ft., fly 40 ft."
"saves":
  "Charisma": !!int "10"
  "Dexterity": !!int "7"
  "Wisdom": !!int "7"
  "Constitution": !!int "9"
"skillsaves":
  "Perception": !!int "7"
"damage_resistances": "cold; bludgeoning, piercing, slashing from nonmagical attacks\
  \ that aren't silvered"
"damage_immunities": "fire, poison"
"condition_immunities": "[charmed](/3-Mechanics/CLI/rules/conditions.md#charmed),\
  \ [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)"
"senses": "darkvision 120 ft., passive Perception 17"
"languages": "Common, Infernal, telepathy 1,000 ft."
"cr": "18"
"traits":
- "desc": "The amnizu's innate spellcasting ability is Intelligence (spell save DC\
    \ 19, dice: d20+11 (+11) to hit with spell attacks). The amnizu can innately\
    \ cast the following spells, requiring no material components:\n\nAt will:\
    \ [charm person](/3-Mechanics/CLI/spells/charm-person.md), [command](/3-Mechanics/CLI/spells/command.md)\n\
    \n1/day each: [dominate monster](/3-Mechanics/CLI/spells/dominate-monster.md),\
    \ [feeblemind](/3-Mechanics/CLI/spells/feeblemind.md)\n\n3/day each: [dominate\
    \ person](/3-Mechanics/CLI/spells/dominate-person.md), [fireball](/3-Mechanics/CLI/spells/fireball.md)"
  "name": "Innate Spellcasting"
- "desc": "Magical darkness doesn't impede the amnizu's darkvision."
  "name": "Devil's Sight"
- "desc": "The amnizu has advantage on saving throws against spells and other magical\
    \ effects."
  "name": "Magic Resistance"
"actions":
- "desc": "The amnizu uses Poison Mind. It also makes two attacks: one with its whip\
    \ and one with its Disruptive Touch."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+11 (+11) to hit, reach 10 ft., one\
    \ target. Hit: dice:2d4 + 5|text(10) (2d4 + 5) slashing damage plus dice:6d10|text(33)\
    \ (6d10) force damage."
  "name": "Taskmaster Whip"
- "desc": "Melee Spell Attack: dice: d20+11 (+11) to hit, reach 5 ft., one target.\
    \ Hit: dice:8d10|text(44) (8d10) necrotic damage."
  "name": "Disruptive Touch"
- "desc": "The amnizu targets one or two creatures that it can see within 60 feet\
    \ of it. Each target must succeed on a DC 19 Wisdom saving throw or take dice:4d12|text(26)\
    \ (4d12) necrotic damage and is [blinded](/3-Mechanics/CLI/rules/conditions.md#blinded)\
    \ until the start of the amnizu's next turn."
  "name": "Poison Mind"
- "desc": "The amnizu targets one creature it can see within 60 feet of it. That creature\
    \ must make a DC 18 Intelligence saving throw and on a failure the target is [stunned](/3-Mechanics/CLI/rules/conditions.md#stunned)\
    \ for 1 minute. A [stunned](/3-Mechanics/CLI/rules/conditions.md#stunned) creature\
    \ repeats the saving throw at the end of each of its turns, ending the effect\
    \ on itself on a success. If the target remains [stunned](/3-Mechanics/CLI/rules/conditions.md#stunned)\
    \ for the full minute, it forgets everything it sensed, experienced, and learned\
    \ during the last 5 hours."
  "name": "Forgetfulness (Recharge 6)"
"reactions":
- "desc": "When a creature within 60 feet of the amnizu makes an attack roll against\
    \ it, and another creature is within the attack's range, the attacker must make\
    \ a DC 19 Wisdom saving throw. On a failed save, the attacker must target the\
    \ creature that is closest to it, not including the amnizu or itself. If multiple\
    \ creatures are closest, the attacker chooses which one to target. If the saving\
    \ throw is successful, the attacker is immune to the amnizu's Instinctive Charm\
    \ for 24 hours."
  "name": "Instinctive Charm"
"source":
- "MTF"
- "BGDIA"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MTF/Amnizu.webp"
```
^statblock

```encounter-table
name: Amnizu
creatures:
 - 1: Amnizu
```