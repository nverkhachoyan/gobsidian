---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/vgm
- monster/cr/8
- monster/environment/underdark
- monster/environment/urban
- monster/size/medium
- monster/type/humanoid/any-race
aliases: ["Blackguard"]
NoteIcon: monster
BestiaryType: humanoid (any race)
SourceType: Bestiary
BookSource: Volo's Guide to Monsters p. 211, Divine Contention, Dragon of Icespire Peak, Sleeping Dragon's Wake
---
# [Blackguard](3-Mechanics\CLI\bestiary\humanoid/blackguard-vgm.md)
*Source: Volo's Guide to Monsters p. 211, Divine Contention, Dragon of Icespire Peak, Sleeping Dragon's Wake*  

Blackguards are paladins who broke their sacred oaths and now indulge their own dark ambitions. They consort with fiends and undead, and they reject all goodly things from their former lives.

```statblock
"name": "Blackguard (VGM)"
"size": "Medium"
"type": "humanoid"
"subtype": "any race"
"alignment": "Any Non-Good alignment"
"ac": !!int "18"
"ac_class": "[plate armor](/3-Mechanics/CLI/items/plate-armor.md)"
"hp": !!int "153"
"hit_dice": "18d8 + 72"
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
- "desc": "The blackguard is a 10th-level spellcaster. Its spellcasting ability is\
    \ Charisma (spell save DC 13, dice: d20+5 (+5) to hit with spell attacks).\
    \ It has the following paladin spells prepared:\n\n1st level (4 slots): [command](/3-Mechanics/CLI/spells/command.md),\
    \ [protection from evil and good](/3-Mechanics/CLI/spells/protection-from-evil-and-good.md),\
    \ [thunderous smite](/3-Mechanics/CLI/spells/thunderous-smite.md)\n\n2nd level\
    \ (3 slots): [branding smite](/3-Mechanics/CLI/spells/branding-smite.md), [find\
    \ steed](/3-Mechanics/CLI/spells/find-steed.md)\n\n3rd level (2 slots): [blinding\
    \ smite](/3-Mechanics/CLI/spells/blinding-smite.md), [dispel magic](/3-Mechanics/CLI/spells/dispel-magic.md)"
  "name": "Spellcasting"
"actions":
- "desc": "The blackguard makes three attacks with its glaive or its shortbow."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+7 (+7) to hit, reach 10 ft., one target.\
    \ Hit: dice:1d10 + 4|text(9) (1d10 + 4) slashing damage."
  "name": "Glaive"
- "desc": "Ranged Weapon Attack: dice: d20+3 (+3) to hit, range 80/320 ft.,\
    \ one target. Hit: dice:1d6 + 2|text(5) (1d6 + 2) piercing damage."
  "name": "Shortbow"
- "desc": "The blackguard exudes magical menace. Each enemy within 30 feet of the\
    \ blackguard must succeed on a DC 13 Wisdom saving throw or be [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened)\
    \ for 1 minute. If a [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened)\
    \ target ends its turn more than 30 feet away from the blackguard, the target\
    \ can repeat the saving throw, ending the effect on itself on a success."
  "name": "Dreadful Aspect (Recharges after a Short or Long Rest)"
"source":
- "VGM"
- "DC"
- "DIP"
- "SDW"
- "MOT"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/VGM/Blackguard.webp"
```
^statblock

```encounter-table
name: Blackguard
creatures:
 - 1: Blackguard
```

## Environment

underdark, urban