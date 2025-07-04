---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/17
- monster/environment/coastal
- monster/environment/desert
- monster/environment/forest
- monster/environment/swamp
- monster/environment/underdark
- monster/environment/urban
- monster/size/medium
- monster/type/monstrosity/wizard
aliases: ["Nagpa"]
NoteIcon: monster
BestiaryType: monstrosity (wizard)
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 189, Mordenkainen's Tome of Foes p. 215
---
# [Nagpa](3-Mechanics\CLI\bestiary\monstrosity/nagpa-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 189, Mordenkainen's Tome of Foes p. 215*  

```statblock
"name": "Nagpa (MPMM)"
"size": "Medium"
"type": "monstrosity"
"subtype": "wizard"
"alignment": "Typically  Neutral Evil"
"ac": !!int "19"
"ac_class": "natural armor"
"hp": !!int "203"
"hit_dice": "37d8 + 37"
"stats":
- !!int "9"
- !!int "15"
- !!int "12"
- !!int "23"
- !!int "18"
- !!int "21"
"speed": "30 ft."
"saves":
  "Charisma": !!int "11"
  "Wisdom": !!int "10"
  "Intelligence": !!int "12"
"skillsaves":
  "Deception": !!int "11"
  "Insight": !!int "10"
  "Perception": !!int "10"
  "History": !!int "12"
  "Arcana": !!int "12"
"senses": "truesight 120 ft., passive Perception 20"
"languages": "Common plus up to five other languages"
"cr": "17"
"traits":
- "desc": "The nagpa casts one of the following spells, using Intelligence as the\
    \ spellcasting ability (spell save DC 20):\n\nAt will: [detect magic](/3-Mechanics/CLI/spells/detect-magic.md),\
    \ [mage hand](/3-Mechanics/CLI/spells/mage-hand.md), [message](/3-Mechanics/CLI/spells/message.md),\
    \ [minor illusion](/3-Mechanics/CLI/spells/minor-illusion.md)\n\n1/day each:\
    \ [dominate person](/3-Mechanics/CLI/spells/dominate-person.md), [etherealness](/3-Mechanics/CLI/spells/etherealness.md),\
    \ [feeblemind](/3-Mechanics/CLI/spells/feeblemind.md)\n\n2/day each: [fireball](/3-Mechanics/CLI/spells/fireball.md),\
    \ [fly](/3-Mechanics/CLI/spells/fly.md), [hold person](/3-Mechanics/CLI/spells/hold-person.md),\
    \ [suggestion](/3-Mechanics/CLI/spells/suggestion.md), [wall of fire](/3-Mechanics/CLI/spells/wall-of-fire.md)"
  "name": "Spellcasting"
"actions":
- "desc": "The nagpa makes three Staff or Deathly Ray attacks. It can replace one\
    \ attack with a use of Spellcasting."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+8 (+8) to hit, reach 5 ft., one target.\
    \ Hit: dice:2d6 + 2|text(9) (2d6 + 2) bludgeoning damage plus dice:7d6|text(24)\
    \ (7d6) necrotic damage."
  "name": "Staff"
- "desc": "Ranged Spell Attack: dice: d20+12 (+12) to hit, range 120 ft., one\
    \ target. Hit: dice:7d6 + 6|text(30) (7d6 + 6) necrotic damage."
  "name": "Deathly Ray"
"bonus_actions":
- "desc": "The nagpa targets one creature it can see within 90 feet of it. The target\
    \ must make a DC 20 Charisma saving throw. An evil creature makes the save with\
    \ disadvantage. On a failed save, the target is [charmed](/3-Mechanics/CLI/rules/conditions.md#charmed)\
    \ by the nagpa until the start of the nagpa's next turn. On a successful save,\
    \ the target becomes immune to the nagpa's Corruption for the next 24 hours."
  "name": "Corruption"
- "desc": "The nagpa forces each creature within 30 feet of it to make a DC 20 Wisdom\
    \ saving throw, excluding Undead and Constructs. On a failed save, a target is\
    \ [paralyzed](/3-Mechanics/CLI/rules/conditions.md#paralyzed) for 1 minute. A\
    \ [paralyzed](/3-Mechanics/CLI/rules/conditions.md#paralyzed) target can repeat\
    \ the saving throw at the end of each of its turns, ending the effect on itself\
    \ on a success."
  "name": "Paralysis (Recharge 6-6)"
"source":
- "MPMM"
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Nagpa.webp"
```
^statblock

```encounter-table
name: Nagpa
creatures:
 - 1: Nagpa
```

## Environment

coastal, desert, forest, swamp, underdark, urban