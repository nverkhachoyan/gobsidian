---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/7
- monster/environment/arctic
- monster/environment/desert
- monster/environment/underdark
- monster/environment/urban
- monster/size/medium
- monster/type/humanoid
aliases: ["Warlock of the Fiend"]
NoteIcon: monster
BestiaryType: humanoid
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 255, Volo's Guide to Monsters p. 219
---
# [Warlock of the Fiend](3-Mechanics\CLI\bestiary\humanoid/warlock-of-the-fiend-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 255, Volo's Guide to Monsters p. 219*  

```statblock
"name": "Warlock of the Fiend (MPMM)"
"size": "Medium"
"type": "humanoid"
"alignment": "Any alignment"
"ac": !!int "13"
"ac_class": "16 with [mage armor](/3-Mechanics/CLI/spells/mage-armor.md)"
"hp": !!int "78"
"hit_dice": "12d8 + 24"
"stats":
- !!int "10"
- !!int "16"
- !!int "15"
- !!int "12"
- !!int "12"
- !!int "18"
"speed": "30 ft."
"saves":
  "Charisma": !!int "7"
  "Wisdom": !!int "4"
"skillsaves":
  "Deception": !!int "7"
  "Religion": !!int "4"
  "Arcana": !!int "4"
  "Persuasion": !!int "7"
"senses": "darkvision 60 ft., passive Perception 11"
"languages": "any two languages (usually Abyssal or Infernal)"
"cr": "7"
"traits":
- "desc": "The warlock casts one of the following spells, using Charisma as the spellcasting\
    \ ability (spell save DC 15): \n\nAt will: [alter self](/3-Mechanics/CLI/spells/alter-self.md),\
    \ [mage armor](/3-Mechanics/CLI/spells/mage-armor.md) (self only), [mage hand](/3-Mechanics/CLI/spells/mage-hand.md),\
    \ [minor illusion](/3-Mechanics/CLI/spells/minor-illusion.md), [prestidigitation](/3-Mechanics/CLI/spells/prestidigitation.md)\n\
    \n1/day each: [banishment](/3-Mechanics/CLI/spells/banishment.md), [plane\
    \ shift](/3-Mechanics/CLI/spells/plane-shift.md), [suggestion](/3-Mechanics/CLI/spells/suggestion.md)"
  "name": "Spellcasting"
- "desc": "When the warlock makes an ability check or saving throw, it can add a dice:\
    \ d10|avg|noform (d10) to the roll. It can do this after the roll is made but\
    \ before any of the roll's effects occur."
  "name": "Dark One's Own Luck (Recharges after a Short or Long Rest)"
"actions":
- "desc": "The warlock makes three Scimitar attacks."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+6 (+6) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d6 + 3|text(6) (1d6 + 3) slashing damage plus dice:4d6|text(14)\
    \ (4d6) fire damage."
  "name": "Scimitar"
- "desc": "Green flame explodes in a 10-foot-radius sphere centered on a point within\
    \ 120 feet of the warlock. Each creature in that area must make a DC 15 Dexterity\
    \ saving throw, taking dice:3d10|text(16) (3d10) fire damage and dice:2d10|text(11)\
    \ (2d10) necrotic damage on a failed save, or half as much damage on a successful\
    \ one."
  "name": "Hellfire"
"reactions":
- "desc": "In response to being damaged by a visible creature within 60 feet of it,\
    \ the warlock forces that creature to make a DC 15 Constitution saving throw,\
    \ taking dice:4d10|text(22) (4d10) necrotic damage on a failed save, or half\
    \ as much damage on a successful one."
  "name": "Fiendish Rebuke (3/Day)"
"source":
- "MPMM"
- "VGM"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Warlock%20of%20the%20Fiend.webp"
```
^statblock

```encounter-table
name: Warlock of the Fiend
creatures:
 - 1: Warlock of the Fiend
```

## Environment

arctic, desert, underdark, urban