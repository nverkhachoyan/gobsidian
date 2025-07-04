---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/16
- monster/environment/coastal
- monster/environment/desert
- monster/environment/mountain
- monster/environment/underwater
- monster/size/huge
- monster/type/giant
aliases: ["Storm Giant Quintessent"]
NoteIcon: monster
BestiaryType: giant
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 235, Volo's Guide to Monsters p. 151
---
# [Storm Giant Quintessent](3-Mechanics\CLI\bestiary\giant/storm-giant-quintessent-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 235, Volo's Guide to Monsters p. 151*  

```statblock
"name": "Storm Giant Quintessent (MPMM)"
"size": "Huge"
"type": "giant"
"alignment": "Typically  Chaotic Good"
"ac": !!int "12"
"hp": !!int "230"
"hit_dice": "20d12 + 100"
"stats":
- !!int "29"
- !!int "14"
- !!int "20"
- !!int "17"
- !!int "20"
- !!int "19"
"speed": "50 ft., fly 50 ft. (hover), swim 50 ft."
"saves":
  "Charisma": !!int "9"
  "Wisdom": !!int "10"
  "Strength": !!int "14"
  "Constitution": !!int "10"
"skillsaves":
  "Perception": !!int "10"
  "History": !!int "8"
  "Arcana": !!int "8"
"damage_resistances": "cold; bludgeoning, piercing, slashing from nonmagical attacks"
"damage_immunities": "lightning, thunder"
"senses": "truesight 60 ft., passive Perception 20"
"languages": "Common, Giant"
"cr": "16"
"traits":
- "desc": "The giant can breathe air and water."
  "name": "Amphibious"
- "desc": "If the giant fails a saving throw, it can choose to succeed instead."
  "name": "Legendary Resistance (1/Day)"
"actions":
- "desc": "The giant makes two Lightning Sword attacks, or it uses Wind Javelin twice."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+14 (+14) to hit, reach 15 ft., one\
    \ target. Hit: dice:9d6 + 9|text(40) (9d6 + 9) lightning damage."
  "name": "Lightning Sword"
- "desc": "The giant coalesces wind into a javelin-like form and hurls it at a creature\
    \ it can see within 600 feet of it. The javelin deals dice:3d6 + 9|text(19)\
    \ (3d6 + 9) force damage to the target, striking unerringly. The javelin disappears\
    \ after it hits."
  "name": "Wind Javelin"
"legendary_actions":
- "desc": "The giant targets a creature it can see within 60 feet of it and creates\
    \ a magical gust of wind around the target, who must succeed on a DC 18 Strength\
    \ saving throw or be moved up to 20 feet in any horizontal direction the giant\
    \ chooses."
  "name": "Gust"
- "desc": "The giant hurls a thunderbolt at a creature it can see within 600 feet\
    \ of it. The target must make a DC 18 Dexterity saving throw, taking dice:4d10|text(22)\
    \ (4d10) thunder damage on a failed save, or half as much damage on a successful\
    \ one."
  "name": "Thunderbolt (Costs 2 Actions)"
- "desc": "The giant vanishes, dispersing itself into the storm surrounding its lair.\
    \ The giant can end this effect at the start of any of its turns, becoming a giant\
    \ once more and appearing in any location it chooses within its lair. While dispersed,\
    \ the giant can't take any actions other than lair actions, and it can't be targeted\
    \ by attacks, spells, or other effects. The giant can't use this ability outside\
    \ its lair, nor can it use this ability if another creature is using a [control\
    \ weather](/3-Mechanics/CLI/spells/control-weather.md) spell or similar magic\
    \ to quell the storm."
  "name": "One with the Storm (Costs 3 Actions)"
"source":
- "MPMM"
- "VGM"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Storm%20Giant%20Quintessent.webp"
```
^statblock

```encounter-table
name: Storm Giant Quintessent
creatures:
 - 1: Storm Giant Quintessent
```

## Environment

coastal, desert, mountain, underwater