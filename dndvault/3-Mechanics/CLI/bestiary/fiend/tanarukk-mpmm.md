---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/5
- monster/environment/hill
- monster/environment/mountain
- monster/environment/underdark
- monster/size/medium
- monster/type/fiend/demon
aliases: ["Tanarukk"]
NoteIcon: monster
BestiaryType: fiend (demon)
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 240, Volo's Guide to Monsters p. 186
---
# [Tanarukk](3-Mechanics\CLI\bestiary\fiend/tanarukk-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 240, Volo's Guide to Monsters p. 186*  

```statblock
"name": "Tanarukk (MPMM)"
"size": "Medium"
"type": "fiend"
"subtype": "demon"
"alignment": "Typically  Chaotic Evil"
"ac": !!int "14"
"ac_class": "natural armor"
"hp": !!int "95"
"hit_dice": "10d8 + 50"
"stats":
- !!int "18"
- !!int "13"
- !!int "20"
- !!int "9"
- !!int "9"
- !!int "9"
"speed": "30 ft."
"skillsaves":
  "Intimidation": !!int "2"
  "Perception": !!int "2"
"damage_resistances": "fire, poison"
"senses": "darkvision 60 ft., passive Perception 12"
"languages": "Abyssal, Common, plus any one language"
"cr": "5"
"traits":
- "desc": "The tanarukk has advantage on saving throws against spells and other magical\
    \ effects."
  "name": "Magic Resistance"
"actions":
- "desc": "The tanarukk makes one Bite attack and one Greatsword attack."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+7 (+7) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d8 + 4|text(8) (1d8 + 4) piercing damage."
  "name": "Bite"
- "desc": "Melee Weapon Attack: dice: d20+7 (+7) to hit, reach 5 ft., one target.\
    \ Hit: dice:2d6 + 4|text(11) (2d6 + 4) slashing damage."
  "name": "Greatsword"
"bonus_actions":
- "desc": "The tanarukk moves up to its speed toward an enemy that it can see."
  "name": "Aggressive"
"reactions":
- "desc": "In response to being hit by a melee attack, the tanarukk can make one Bite\
    \ or Greatsword attack with advantage against the attacker."
  "name": "Unbridled Fury"
"source":
- "MPMM"
- "VGM"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Tanarukk.webp"
```
^statblock

```encounter-table
name: Tanarukk
creatures:
 - 1: Tanarukk
```

## Environment

hill, mountain, underdark