---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mm
- monster/cr/2
- monster/environment/arctic
- monster/environment/coastal
- monster/environment/desert
- monster/environment/forest
- monster/environment/grassland
- monster/environment/hill
- monster/environment/mountain
- monster/environment/swamp
- monster/environment/underdark
- monster/size/medium
- monster/type/humanoid/any-race
aliases: ["Druid"]
NoteIcon: monster
BestiaryType: humanoid (any race)
SourceType: Bestiary
BookSource: Monster Manual p. 346, Dragon of Icespire Peak, Sleeping Dragon's Wake. Available in the SRD.
---
# [Druid](3-Mechanics\CLI\bestiary\humanoid/druid.md)
*Source: Monster Manual p. 346, Dragon of Icespire Peak, Sleeping Dragon's Wake. Available in the SRD.*  

Druids dwell in forests and other secluded wilderness locations, where they protect the natural world from monsters and the encroachment of civilization. Some are tribal shamans who heal the sick, pray to animal spirits, and provide spiritual guidance.

```statblock
"name": "Druid"
"size": "Medium"
"type": "humanoid"
"subtype": "any race"
"alignment": "Any alignment"
"ac": !!int "11"
"ac_class": "16 with [barkskin](/3-Mechanics/CLI/spells/barkskin.md)"
"hp": !!int "27"
"hit_dice": "5d8 + 5"
"stats":
- !!int "10"
- !!int "12"
- !!int "13"
- !!int "12"
- !!int "15"
- !!int "11"
"speed": "30 ft."
"skillsaves":
  "Medicine": !!int "4"
  "Nature": !!int "3"
  "Perception": !!int "4"
"senses": "passive Perception 14"
"languages": "Druidic plus any two languages"
"cr": "2"
"traits":
- "desc": "The druid is a 4th-level spellcaster. Its spellcasting ability is Wisdom\
    \ (spell save DC 12, dice: d20+4 (+4) to hit with spell attacks). It has the\
    \ following druid spells prepared:\n\nCantrips (at will): [druidcraft](/3-Mechanics/CLI/spells/druidcraft.md),\
    \ [produce flame](/3-Mechanics/CLI/spells/produce-flame.md), [shillelagh](/3-Mechanics/CLI/spells/shillelagh.md)\n\
    \n1st level (4 slots): [entangle](/3-Mechanics/CLI/spells/entangle.md), [longstrider](/3-Mechanics/CLI/spells/longstrider.md),\
    \ [speak with animals](/3-Mechanics/CLI/spells/speak-with-animals.md), [thunderwave](/3-Mechanics/CLI/spells/thunderwave.md)\n\
    \n2nd level (3 slots): [animal messenger](/3-Mechanics/CLI/spells/animal-messenger.md),\
    \ [barkskin](/3-Mechanics/CLI/spells/barkskin.md)"
  "name": "Spellcasting"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+2 (+2) to hit (dice: d20+4 (+4)\
    \ to hit with shillelagh), reach 5 ft., one target. Hit: dice:1d6|text(3)\
    \ (1d6) bludgeoning damage, dice:1d8|text(4) (1d8) bludgeoning damage if\
    \ wielded with two hands, or dice:1d8 + 2|text(6) (1d8 + 2) bludgeoning damage\
    \ with [shillelagh](/3-Mechanics/CLI/spells/shillelagh.md)."
  "name": "Quarterstaff"
"source":
- "MM"
- "CoS"
- "PotA"
- "RoT"
- "SKT"
- "ToA"
- "GoS"
- "DIP"
- "SDW"
- "BGDIA"
- "EGW"
- "MOT"
- "IDRotF"
- "CM"
- "WBtW"
- "CRCotN"
- "JttRC"
- "PSI"
- "SatO"
- "BMT"
- "DoDk"
- "VEoR"
- "QftIS"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MM/Druid.webp"
```
^statblock

```encounter-table
name: Druid
creatures:
 - 1: Druid
```

## Environment

coastal, mountain, grassland, hill, arctic, forest, swamp, underdark, desert