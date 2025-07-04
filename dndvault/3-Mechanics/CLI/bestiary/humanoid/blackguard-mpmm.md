---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/8
- monster/environment/underdark
- monster/environment/urban
- monster/size/medium
- monster/type/humanoid/paladin
aliases: ["Blackguard"]
NoteIcon: monster
BestiaryType: humanoid (paladin)
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 63, Volo's Guide to Monsters p. 211
---
# [Blackguard](3-Mechanics\CLI\bestiary\humanoid/blackguard-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 63, Volo's Guide to Monsters p. 211*  

```statblock
"name": "Blackguard (MPMM)"
"size": "Medium"
"type": "humanoid"
"subtype": "paladin"
"alignment": "Typically  Neutral Evil"
"ac": !!int "18"
"ac_class": "[plate](/3-Mechanics/CLI/items/plate-armor.md)"
"hp": !!int "119"
"hit_dice": "14d8 + 56"
"stats":
- !!int "18"
- !!int "11"
- !!int "18"
- !!int "11"
- !!int "14"
- !!int "15"
"speed": "30 ft."
"saves":
  "Charisma": !!int "5"
  "Wisdom": !!int "5"
"skillsaves":
  "Intimidation": !!int "5"
  "Athletics": !!int "7"
  "Deception": !!int "5"
"senses": "passive Perception 12"
"languages": "any one language (usually Common)"
"cr": "8"
"traits":
- "desc": "The blackguard casts one of the following spells, using Charisma as the\
    \ spellcasting ability (spell save DC 13):\n\n2/day each: [command](/3-Mechanics/CLI/spells/command.md),\
    \ [dispel magic](/3-Mechanics/CLI/spells/dispel-magic.md), [find steed](/3-Mechanics/CLI/spells/find-steed.md)"
  "name": "Spellcasting"
"actions":
- "desc": "The blackguard makes three attacks, using Glaive, Shortbow, or both."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+7 (+7) to hit, reach 10 ft., one target.\
    \ Hit: dice:1d10 + 4|text(9) (1d10 + 4) slashing damage plus dice:2d8|text(9)\
    \ (2d8) necrotic damage."
  "name": "Glaive"
- "desc": "Ranged Weapon Attack: dice: d20+3 (+3) to hit, range 80/320 ft.,\
    \ one target. Hit: dice:1d6 + 2|text(5) (1d6 + 2) piercing damage."
  "name": "Shortbow"
- "desc": "Each enemy within 30 feet of the blackguard must succeed on a DC 13 Wisdom\
    \ saving throw or be [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened)\
    \ of the blackguard for 1 minute. If a [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened)\
    \ target ends its turn more than 30 feet away from the blackguard, the target\
    \ can repeat the saving throw, ending the effect on itself on a success."
  "name": "Dreadful Aspect (Recharges after a Short or Long Rest)"
"bonus_actions":
- "desc": "Immediately after the blackguard hits a target with an attack roll, the\
    \ blackguard can force that target to make a DC 13 Constitution saving throw.\
    \ On a failed save, the target suffers one of the following effects of the blackguard's\
    \ choice:"
  "name": "Smite"
- "desc": "The target is [blinded](/3-Mechanics/CLI/rules/conditions.md#blinded) for\
    \ 1 minute. The [blinded](/3-Mechanics/CLI/rules/conditions.md#blinded) target\
    \ can repeat the save at the end of each of its turns, ending the effect on itself\
    \ on a success."
  "name": "Blind"
- "desc": "The target is pushed up to 10 feet away and knocked [prone](/3-Mechanics/CLI/rules/conditions.md#prone)."
  "name": "Shove"
"source":
- "MPMM"
- "VGM"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Blackguard.webp"
```
^statblock

```encounter-table
name: Blackguard
creatures:
 - 1: Blackguard
```

## Environment

underdark, urban