---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/3
- monster/environment/underdark
- monster/size/medium
- monster/type/monstrosity/cleric
aliases: ["Choldrith"]
NoteIcon: monster
BestiaryType: monstrosity (cleric)
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 77, Volo's Guide to Monsters p. 132
---
# [Choldrith](3-Mechanics\CLI\bestiary\monstrosity/choldrith-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 77, Volo's Guide to Monsters p. 132*  

```statblock
"name": "Choldrith (MPMM)"
"size": "Medium"
"type": "monstrosity"
"subtype": "cleric"
"alignment": "Typically  Chaotic Evil"
"ac": !!int "15"
"ac_class": "natural armor"
"hp": !!int "66"
"hit_dice": "12d8 + 12"
"stats":
- !!int "12"
- !!int "16"
- !!int "12"
- !!int "11"
- !!int "14"
- !!int "10"
"speed": "30 ft., climb 30 ft."
"skillsaves":
  "Athletics": !!int "5"
  "Stealth": !!int "5"
  "Religion": !!int "2"
"senses": "darkvision 60 ft., passive Perception 12"
"languages": "Undercommon"
"cr": "3"
"traits":
- "desc": "The choldrith casts one of the following spells, using Wisdom as the spellcasting\
    \ ability (spell save DC 12):\n\nAt will: [guidance](/3-Mechanics/CLI/spells/guidance.md),\
    \ [thaumaturgy](/3-Mechanics/CLI/spells/thaumaturgy.md)\n\n1/day each: [bane](/3-Mechanics/CLI/spells/bane.md),\
    \ [hold person](/3-Mechanics/CLI/spells/hold-person.md)"
  "name": "Spellcasting"
- "desc": "The choldrith has advantage on saving throws against being [charmed](/3-Mechanics/CLI/rules/conditions.md#charmed),\
    \ and magic can't put the choldrith to sleep."
  "name": "Fey Ancestry"
- "desc": "The choldrith can climb difficult surfaces, including upside down on ceilings,\
    \ without needing to make an ability check."
  "name": "Spider Climb"
- "desc": "While in sunlight, the choldrith has disadvantage on attack rolls, as well\
    \ as on Wisdom ([Perception](/3-Mechanics/CLI/rules/skills.md#Perception)) checks\
    \ that rely on sight."
  "name": "Sunlight Sensitivity"
- "desc": "While in contact with a web, the choldrith knows the exact location of\
    \ any other creature in contact with the same web."
  "name": "Web Sense"
- "desc": "The choldrith ignores movement restrictions caused by webbing."
  "name": "Web Walker"
"actions":
- "desc": "Melee or Ranged Weapon Attack: dice: d20+5 (+5) to hit, reach 5 ft.\
    \ or range 20/60 ft., one target. Hit: dice:1d4 + 3|text(5) (1d4 + 3) piercing\
    \ damage plus dice:3d6|text(10) (3d6) poison damage."
  "name": "Dagger"
- "desc": "Ranged Weapon Attack: dice: d20+5 (+5) to hit, range 30/60 ft., one\
    \ Large or smaller creature. Hit: The target is [restrained](/3-Mechanics/CLI/rules/conditions.md#restrained)\
    \ by webbing. As an action, the [restrained](/3-Mechanics/CLI/rules/conditions.md#restrained)\
    \ target can make a DC 11 Strength check, bursting the webbing on a success. The\
    \ webbing can also be attacked and destroyed (AC 10; 5 hit points; vulnerability\
    \ to fire damage; immunity to bludgeoning, poison, and psychic damage)."
  "name": "Web (Recharge 5-6)"
"bonus_actions":
- "desc": "The choldrith conjures a floating, spectral dagger within 60 feet of itself.\
    \ The choldrith can make a melee spell attack (dice: d20+4 (+4) to hit) against\
    \ one creature within 5 feet of the dagger. On a hit, the target takes dice:1d8\
    \ + 2|text(6) (1d8 + 2) force damage.\n\nThe dagger lasts for 1 minute. As\
    \ a bonus action on later turns, the choldrith can move the dagger up to 20 feet\
    \ and repeat the attack against one creature within 5 feet of the dagger."
  "name": "Spectral Dagger (Recharges after a Short or Long Rest)"
"source":
- "MPMM"
- "VGM"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Choldrith.webp"
```
^statblock

```encounter-table
name: Choldrith
creatures:
 - 1: Choldrith
```

## Environment

underdark