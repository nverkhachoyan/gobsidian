---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mm
- monster/cr/1-8
- monster/environment/urban
- monster/size/medium
- monster/type/humanoid/any-race
aliases: ["Cultist"]
NoteIcon: monster
BestiaryType: humanoid (any race)
SourceType: Bestiary
BookSource: Monster Manual p. 345, Divine Contention, Dragon of Icespire Peak, Storm Lord's Wrath. Available in the SRD and the Basic Rules.
---
# [Cultist](3-Mechanics\CLI\bestiary\humanoid/cultist.md)
*Source: Monster Manual p. 345, Divine Contention, Dragon of Icespire Peak, Storm Lord's Wrath. Available in the SRD and the Basic Rules.*  

Cultists swear allegiance to dark powers such as elemental princes, demon lords, or archdevils. Most conceal their loyalties to avoid being ostracized, imprisoned, or executed for their beliefs. Unlike evil acolytes, cultists often show signs of insanity in their beliefs and practices.

```statblock
"name": "Cultist"
"size": "Medium"
"type": "humanoid"
"subtype": "any race"
"alignment": "Any Non-Good alignment"
"ac": !!int "12"
"ac_class": "[leather armor](/3-Mechanics/CLI/items/leather-armor.md)"
"hp": !!int "9"
"hit_dice": "2d8"
"stats":
- !!int "11"
- !!int "12"
- !!int "10"
- !!int "10"
- !!int "11"
- !!int "10"
"speed": "30 ft."
"skillsaves":
  "Deception": !!int "2"
  "Religion": !!int "2"
"senses": "passive Perception 10"
"languages": "any one language (usually Common)"
"cr": "1/8"
"traits":
- "desc": "The cultist has advantage on saving throws against being [charmed](/3-Mechanics/CLI/rules/conditions.md#charmed)\
    \ or [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened)."
  "name": "Dark Devotion"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+3 (+3) to hit, reach 5 ft., one creature.\
    \ Hit: dice:1d6 + 1|text(4) (1d6 + 1) slashing damage."
  "name": "Scimitar"
"source":
- "MM"
- "CoS"
- "HotDQ"
- "LMoP"
- "PotA"
- "RoT"
- "SKT"
- "WDH"
- "WDMM"
- "GoS"
- "DC"
- "DIP"
- "SLW"
- "BGDIA"
- "RMBRE"
- "IMR"
- "EGW"
- "MOT"
- "IDRotF"
- "CM"
- "JttRC"
- "KftGV"
- "HftT"
- "PaBTSO"
- "AATM"
- "BMT"
- "GHLoE"
- "DoDk"
- "VEoR"
- "QftIS"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MM/Cultist.webp"
```
^statblock

```encounter-table
name: Cultist
creatures:
 - 1: Cultist
```

## Environment

urban