---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mm
- monster/cr/2
- monster/environment/underdark
- monster/size/medium
- monster/type/aberration
aliases: ["Nothic"]
NoteIcon: monster
BestiaryType: aberration
SourceType: Bestiary
BookSource: Monster Manual p. 236. Available in the Basic Rules.
---
# [Nothic](3-Mechanics\CLI\bestiary\aberration/nothic.md)
*Source: Monster Manual p. 236. Available in the Basic Rules.*  

A baleful eye peers out from the darkness, its gleam hinting at a weird intelligence and unnerving malevolence. Most times, a nothic is content to watch, weighing and assessing the creatures it encounters. When driven to violence, it uses its horrific gaze to rot the flesh from its enemies' bones.

## Cursed Arcanists

Rather than gaining the godlike supremacy they crave, some wizards who devote their lives to unearthing arcane secrets are reduced to creeping, tormented monsters by a dark curse left behind by Vecna, a powerful lich who, in some worlds, has transcended his undead existence to become a god of secrets. Nothics retain no awareness of their former selves, skulking amid the shadows and haunting places rich in magical knowledge, drawn by memories and impulses they can't quite understand.

## Dark Oracles

Nothics possess a strange magical insight that allows them to extract knowledge from other creatures. This grants them unique understanding of secret and forbidden lore, which they share for a price. A nothic covets magic items, greedily accepting such gifts from creatures that seek out its knowledge.

### Lurkers in Magical Places

Nothics are notorious for infiltrating arcane academies and other places rich in magical learning. They are driven by the vague knowledge that there exists a method to reverse their condition. This isn't a clear sense of purpose, but rather an obsessive tug at the end of the mind. Some nothics are clever enough to realize that this is merely part of the strange lesson for their folly, a false hope to drive them to seek out more arcane secrets.

```statblock
"name": "Nothic"
"size": "Medium"
"type": "aberration"
"alignment": "Neutral Evil"
"ac": !!int "15"
"ac_class": "natural armor"
"hp": !!int "45"
"hit_dice": "6d8 + 18"
"stats":
- !!int "14"
- !!int "16"
- !!int "16"
- !!int "13"
- !!int "10"
- !!int "8"
"speed": "30 ft."
"skillsaves":
  "Stealth": !!int "5"
  "Insight": !!int "4"
  "Perception": !!int "2"
  "Arcana": !!int "3"
"senses": "truesight 120 ft., passive Perception 12"
"languages": "Undercommon"
"cr": "2"
"traits":
- "desc": "The nothic has advantage on Wisdom ([Perception](/3-Mechanics/CLI/rules/skills.md#Perception))\
    \ checks that rely on sight."
  "name": "Keen Sight"
"actions":
- "desc": "The nothic makes two claw attacks."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+4 (+4) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d6 + 3|text(6) (1d6 + 3) slashing damage."
  "name": "Claw"
- "desc": "The nothic targets one creature it can see within 30 feet of it. The target\
    \ must succeed on a DC 12 Constitution saving throw against this magic or take\
    \ dice:3d6|text(10) (3d6) necrotic damage."
  "name": "Rotting Gaze"
- "desc": "The nothic targets one creature it can see within 30 feet of it. The target\
    \ must contest its Charisma ([Deception](/3-Mechanics/CLI/rules/skills.md#Deception))\
    \ check against the nothic's Wisdom ([Insight](/3-Mechanics/CLI/rules/skills.md#Insight))\
    \ check. If the nothic wins, it magically learns one fact or secret about the\
    \ target. The target automatically wins if it is immune to being [charmed](/3-Mechanics/CLI/rules/conditions.md#charmed)."
  "name": "Weird Insight"
"source":
- "MM"
- "CoS"
- "LMoP"
- "PotA"
- "ToA"
- "WDMM"
- "BGDIA"
- "RMBRE"
- "EGW"
- "MOT"
- "IDRotF"
- "JttRC"
- "KftGV"
- "PaBTSO"
- "ToFW"
- "BMT"
- "VEoR"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MM/Nothic.webp"
```
^statblock

```encounter-table
name: Nothic
creatures:
 - 1: Nothic
```

## Environment

underdark