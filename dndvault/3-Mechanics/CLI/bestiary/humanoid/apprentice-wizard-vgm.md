---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/vgm
- monster/cr/1-4
- monster/environment/urban
- monster/size/medium
- monster/type/humanoid/any-race
aliases: ["Apprentice Wizard"]
NoteIcon: monster
BestiaryType: humanoid (any race)
SourceType: Bestiary
BookSource: Volo's Guide to Monsters p. 209, Dragon of Icespire Peak, Storm Lord's Wrath
---
# [Apprentice Wizard](3-Mechanics\CLI\bestiary\humanoid/apprentice-wizard-vgm.md)
*Source: Volo's Guide to Monsters p. 209, Dragon of Icespire Peak, Storm Lord's Wrath*  

Apprentices are novice arcane spellcasters who serve more experienced wizards or attend school. They perform menial work such as cooking and cleaning in exchange for education in the ways of magic.

```statblock
"name": "Apprentice Wizard (VGM)"
"size": "Medium"
"type": "humanoid"
"subtype": "any race"
"alignment": "Any alignment"
"ac": !!int "10"
"hp": !!int "9"
"hit_dice": "2d8"
"stats":
- !!int "10"
- !!int "10"
- !!int "10"
- !!int "14"
- !!int "10"
- !!int "11"
"speed": "30 ft."
"skillsaves":
  "History": !!int "4"
  "Arcana": !!int "4"
"senses": "passive Perception 10"
"languages": "any one language (usually Common)"
"cr": "1/4"
"traits":
- "desc": "The apprentice is a 1st-level spellcaster. Its spellcasting ability is\
    \ Intelligence (spell save DC 12, dice: d20+4 (+4) to hit with spell attacks).\
    \ It has the following wizard spells prepared:\n\nCantrips (at will): [fire\
    \ bolt](/3-Mechanics/CLI/spells/fire-bolt.md), [mending](/3-Mechanics/CLI/spells/mending.md),\
    \ [prestidigitation](/3-Mechanics/CLI/spells/prestidigitation.md)\n\n1st level\
    \ (2 slots): [burning hands](/3-Mechanics/CLI/spells/burning-hands.md), [disguise\
    \ self](/3-Mechanics/CLI/spells/disguise-self.md), [shield](/3-Mechanics/CLI/spells/shield.md)"
  "name": "Spellcasting"
"actions":
- "desc": "Melee or Ranged Weapon Attack: dice: d20+2 (+2) to hit, reach 5 ft.\
    \ or range 20/60 ft., one target. Hit: dice:1d4|text(2) (1d4) piercing damage."
  "name": "Dagger"
"source":
- "VGM"
- "WDH"
- "ToA"
- "DIP"
- "SLW"
- "ERLW"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/VGM/Apprentice%20Wizard.webp"
```
^statblock

```encounter-table
name: Apprentice Wizard
creatures:
 - 1: Apprentice Wizard
```

## Environment

urban