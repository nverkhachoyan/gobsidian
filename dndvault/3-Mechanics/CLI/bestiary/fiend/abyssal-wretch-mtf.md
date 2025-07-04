---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mtf
- monster/cr/1-4
- monster/size/medium
- monster/type/fiend/demon
aliases: ["Abyssal Wretch"]
NoteIcon: monster
BestiaryType: fiend (demon)
SourceType: Bestiary
BookSource: Mordenkainen's Tome of Foes p. 136
---
# [Abyssal Wretch](3-Mechanics\CLI\bestiary\fiend/abyssal-wretch-mtf.md)
*Source: Mordenkainen's Tome of Foes p. 136*  

```statblock
"name": "Abyssal Wretch (MTF)"
"size": "Medium"
"type": "fiend"
"subtype": "demon"
"alignment": "Chaotic Evil"
"ac": !!int "11"
"hp": !!int "18"
"hit_dice": "4d8"
"stats":
- !!int "9"
- !!int "12"
- !!int "11"
- !!int "5"
- !!int "8"
- !!int "5"
"speed": "20 ft."
"damage_resistances": "cold, fire, lightning"
"damage_immunities": "poison"
"condition_immunities": "[charmed](/3-Mechanics/CLI/rules/conditions.md#charmed),\
  \ [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened), [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)"
"senses": "darkvision 120 ft., passive Perception 9"
"languages": "understands Abyssal but can't speak"
"cr": "1/4"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+3 (+3) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d8 + 1|text(5) (1d8 + 1) slashing damage."
  "name": "Bite"
"source":
- "MTF"
- "BGDIA"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MTF/Abyssal%20Wretch.webp"
```
^statblock

```encounter-table
name: Abyssal Wretch
creatures:
 - 1: Abyssal Wretch
```