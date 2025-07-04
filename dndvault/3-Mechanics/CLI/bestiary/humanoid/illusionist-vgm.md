---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/vgm
- monster/cr/3
- monster/environment/urban
- monster/size/medium
- monster/type/humanoid/any-race
aliases: ["Illusionist"]
NoteIcon: monster
BestiaryType: humanoid (any race)
SourceType: Bestiary
BookSource: Volo's Guide to Monsters p. 214
---
# [Illusionist](3-Mechanics\CLI\bestiary\humanoid/illusionist-vgm.md)
*Source: Volo's Guide to Monsters p. 214*  

Illusionists are specialist wizards who twist light, sound, shadow, and even minds to create false and quasi-real effects. They can be flamboyant and use their powers in spectacular and obvious ways, or quiet and subtle, using their magic to conceal the truth.

```statblock
"name": "Illusionist (VGM)"
"size": "Medium"
"type": "humanoid"
"subtype": "any race"
"alignment": "Any alignment"
"ac": !!int "12"
"ac_class": "15 with [mage armor](/3-Mechanics/CLI/spells/mage-armor.md)"
"hp": !!int "38"
"hit_dice": "7d8 + 7"
"stats":
- !!int "9"
- !!int "14"
- !!int "13"
- !!int "16"
- !!int "11"
- !!int "12"
"speed": "30 ft."
"saves":
  "Wisdom": !!int "2"
  "Intelligence": !!int "5"
"skillsaves":
  "History": !!int "5"
  "Arcana": !!int "5"
"senses": "passive Perception 10"
"languages": "any four languages"
"cr": "3"
"traits":
- "desc": "The illusionist is a 7th-level spellcaster. its spellcasting ability is\
    \ Intelligence (spell save DC 13, dice: d20+5 (+5) to hit with spell attacks).\
    \ The illusionist has the following wizard spells prepared:\n\nCantrips (at\
    \ will): [dancing lights](/3-Mechanics/CLI/spells/dancing-lights.md), [mage\
    \ hand](/3-Mechanics/CLI/spells/mage-hand.md), [minor illusion](/3-Mechanics/CLI/spells/minor-illusion.md),\
    \ [poison spray](/3-Mechanics/CLI/spells/poison-spray.md)\n\n1st level (4 slots):\
    \ [color spray](/3-Mechanics/CLI/spells/color-spray.md), [disguise self](/3-Mechanics/CLI/spells/disguise-self.md),\
    \ [mage armor](/3-Mechanics/CLI/spells/mage-armor.md), [magic missile](/3-Mechanics/CLI/spells/magic-missile.md)\n\
    \n2nd level (3 slots): [invisibility](/3-Mechanics/CLI/spells/invisibility.md),\
    \ [mirror image](/3-Mechanics/CLI/spells/mirror-image.md), [phantasmal force](/3-Mechanics/CLI/spells/phantasmal-force.md)\n\
    \n3rd level (3 slots): [major image](/3-Mechanics/CLI/spells/major-image.md),\
    \ [phantom steed](/3-Mechanics/CLI/spells/phantom-steed.md)\n\n4th level (1\
    \ slots): [phantasmal killer](/3-Mechanics/CLI/spells/phantasmal-killer.md)\n\
    \nIllusion spell of 1st level or higher"
  "name": "Spellcasting"
- "desc": "As a bonus action, the illusionist projects an illusion that makes the\
    \ illusionist appear to be standing in a place a few inches from its actual location,\
    \ causing any creature to have disadvantage on attack rolls against the illusionist.\
    \ The effect ends if the illusionist takes damage, it is [incapacitated](/3-Mechanics/CLI/rules/conditions.md#incapacitated),\
    \ or its speed becomes 0."
  "name": "Displacement (Recharges after the Illusionist Casts an Illusion Spell of\
    \ 1st Level or Higher)"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+1 (+1) to hit, reach 5 ft., one target.\
    \ Hit:* dice:1d6 - 1|text(2) (1d6 - 1) bludgeoning damage, or dice:1d8 -\
    \ 1|text(3) (1d8 - 1) bludgeoning damage if used with two hands."
  "name": "Quarterstaff"
"source":
- "VGM"
- "TftYP"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/VGM/Illusionist.webp"
```
^statblock

```encounter-table
name: Illusionist
creatures:
 - 1: Illusionist
```

## Environment

urban