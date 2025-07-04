---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/vgm
- monster/cr/3
- monster/environment/coastal
- monster/environment/urban
- monster/size/medium
- monster/type/humanoid/any-race
aliases: ["Swashbuckler"]
NoteIcon: monster
BestiaryType: humanoid (any race)
SourceType: Bestiary
BookSource: Volo's Guide to Monsters p. 217, Dragon of Icespire Peak, Storm Lord's Wrath, Sleeping Dragon's Wake
---
# [Swashbuckler](3-Mechanics\CLI\bestiary\humanoid/swashbuckler-vgm.md)
*Source: Volo's Guide to Monsters p. 217, Dragon of Icespire Peak, Storm Lord's Wrath, Sleeping Dragon's Wake*  

Swashbucklers are charming ne'er-do-wells who live by their own codes of honor. They crave notoriety, often indulge in romantic trysts, and eke out livings as pirates and corsairs, rarely staying in one place for too long.

```statblock
"name": "Swashbuckler (VGM)"
"size": "Medium"
"type": "humanoid"
"subtype": "any race"
"alignment": "Any Non-Lawful alignment"
"ac": !!int "17"
"ac_class": "[leather armor](/3-Mechanics/CLI/items/leather-armor.md)"
"hp": !!int "66"
"hit_dice": "12d8 + 12"
"stats":
- !!int "12"
- !!int "18"
- !!int "12"
- !!int "14"
- !!int "11"
- !!int "15"
"speed": "30 ft."
"skillsaves":
  "Athletics": !!int "5"
  "Acrobatics": !!int "8"
  "Persuasion": !!int "6"
"senses": "passive Perception 10"
"languages": "any one language (usually Common)"
"cr": "3"
"traits":
- "desc": "The swashbuckler can take the Dash or Disengage action as a bonus action\
    \ on each of its turns."
  "name": "Lightfooted"
- "desc": "While the swashbuckler is wearing light or no armor and wielding no shield,\
    \ its AC includes its Charisma modifier."
  "name": "Suave Defense"
"actions":
- "desc": "The swashbuckler makes three attacks: one with a dagger and two with its\
    \ rapier."
  "name": "Multiattack"
- "desc": "Melee or Ranged Weapon Attack: dice: d20+6 (+6) to hit, reach 5 ft.\
    \ or range 20/60 ft., one target. Hit: dice:1d4 + 4|text(6) (1d4 + 4) piercing\
    \ damage."
  "name": "Dagger"
- "desc": "Melee Weapon Attack: dice: d20+6 (+6) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d8 + 4|text(8) (1d8 + 4) piercing damage."
  "name": "Rapier"
"source":
- "VGM"
- "WDH"
- "ToA"
- "DIP"
- "SLW"
- "SDW"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/VGM/Swashbuckler.webp"
```
^statblock

```encounter-table
name: Swashbuckler
creatures:
 - 1: Swashbuckler
```

## Environment

coastal, urban