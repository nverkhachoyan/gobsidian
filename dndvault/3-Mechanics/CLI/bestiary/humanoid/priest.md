---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mm
- monster/cr/2
- monster/environment/urban
- monster/size/medium
- monster/type/humanoid/any-race
aliases: ["Priest"]
NoteIcon: monster
BestiaryType: humanoid (any race)
SourceType: Bestiary
BookSource: Monster Manual p. 348, Divine Contention, Dragon of Icespire Peak, Storm Lord's Wrath, Sleeping Dragon's Wake, Tasha's Cauldron of Everything. Available in the SRD and the Basic Rules.
---
# [Priest](3-Mechanics\CLI\bestiary\humanoid/priest.md)
*Source: Monster Manual p. 348, Divine Contention, Dragon of Icespire Peak, Storm Lord's Wrath, Sleeping Dragon's Wake, Tasha's Cauldron of Everything. Available in the SRD and the Basic Rules.*  

Priests bring the teachings of their gods to the common folk. They are the spiritual leaders of temples and shrines and often hold positions of influence in their communities. Evil priests might work openly under a tyrant, or they might be the leaders of religious sects hidden in the shadows of good society, overseeing depraved rites. A priest typically has one or more acolytes to help with religious ceremonies and other sacred duties.

```statblock
"name": "Priest"
"size": "Medium"
"type": "humanoid"
"subtype": "any race"
"alignment": "Any alignment"
"ac": !!int "13"
"ac_class": "[chain shirt](/3-Mechanics/CLI/items/chain-shirt.md)"
"hp": !!int "27"
"hit_dice": "5d8 + 5"
"stats":
- !!int "10"
- !!int "10"
- !!int "12"
- !!int "13"
- !!int "16"
- !!int "13"
"speed": "30 ft."
"skillsaves":
  "Medicine": !!int "7"
  "Religion": !!int "5"
  "Persuasion": !!int "3"
"senses": "passive Perception 13"
"languages": "any two languages"
"cr": "2"
"traits":
- "desc": "The priest is a 5th-level spellcaster. Its spellcasting ability is Wisdom\
    \ (spell save DC 13, dice: d20+5 (+5) to hit with spell attacks). The priest\
    \ has the following cleric spells prepared:\n\nCantrips (at will): [light](/3-Mechanics/CLI/spells/light.md),\
    \ [sacred flame](/3-Mechanics/CLI/spells/sacred-flame.md), [thaumaturgy](/3-Mechanics/CLI/spells/thaumaturgy.md)\n\
    \n1st level (4 slots): [cure wounds](/3-Mechanics/CLI/spells/cure-wounds.md),\
    \ [guiding bolt](/3-Mechanics/CLI/spells/guiding-bolt.md), [sanctuary](/3-Mechanics/CLI/spells/sanctuary.md)\n\
    \n2nd level (3 slots): [lesser restoration](/3-Mechanics/CLI/spells/lesser-restoration.md),\
    \ [spiritual weapon](/3-Mechanics/CLI/spells/spiritual-weapon.md)\n\n3rd level\
    \ (2 slots): [dispel magic](/3-Mechanics/CLI/spells/dispel-magic.md), [spirit\
    \ guardians](/3-Mechanics/CLI/spells/spirit-guardians.md)"
  "name": "Spellcasting"
- "desc": "As a bonus action, the priest can expend a spell slot to cause its melee\
    \ weapon attacks to magically deal an extra dice:3d6|text(10) (3d6) radiant\
    \ damage to a target on a hit. This benefit lasts until the end of the turn. If\
    \ the priest expends a spell slot of 2nd level or higher, the extra damage increases\
    \ by dice: 1d6|avg|noform (1d6) for each level above 1st."
  "name": "Divine Eminence"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+2 (+2) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d6|text(3) (1d6) bludgeoning damage."
  "name": "Mace"
"source":
- "MM"
- "CoS"
- "HotDQ"
- "PotA"
- "SKT"
- "TftYP"
- "ToA"
- "WDH"
- "WDMM"
- "GoS"
- "DC"
- "DIP"
- "SLW"
- "SDW"
- "BGDIA"
- "ERLW"
- "EGW"
- "MOT"
- "IDRotF"
- "TCE"
- "CM"
- "CRCotN"
- "JttRC"
- "DSotDQ"
- "PSI"
- "SatO"
- "ToFW"
- "BMT"
- "DoDk"
- "VEoR"
- "QftIS"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MM/Priest.webp"
```
^statblock

```encounter-table
name: Priest
creatures:
 - 1: Priest
```

## Environment

urban