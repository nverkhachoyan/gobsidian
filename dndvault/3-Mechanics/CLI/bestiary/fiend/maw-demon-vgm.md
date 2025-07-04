---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/vgm
- monster/cr/1
- monster/environment/underdark
- monster/size/medium
- monster/type/fiend/demon
aliases: ["Maw Demon"]
NoteIcon: monster
BestiaryType: fiend (demon)
SourceType: Bestiary
BookSource: Volo's Guide to Monsters p. 137
---
# [Maw Demon](3-Mechanics\CLI\bestiary\fiend/maw-demon-vgm.md)
*Source: Volo's Guide to Monsters p. 137*  

Demon lords create lesser demons for the purpose of spreading chaos and terror throughout the multiverse.

## Maw Demon

Maw demons share Yeenoghu's ceaseless hunger for carnage and mortal flesh. After a maw demon rests for 8 hours, anything devoured by it is transported directly into the Lord of Savagery's gullet.

Maw demons appear among gnoll war bands, usually summoned as part of ritual offerings of freshly slain humanoids made to Yeenoghu. The gnolls don't command them, but these demons accompany the war band and attack whatever creatures the gnolls fall upon.

```statblock
"name": "Maw Demon (VGM)"
"size": "Medium"
"type": "fiend"
"subtype": "demon"
"alignment": "Chaotic Evil"
"ac": !!int "13"
"ac_class": "natural armor"
"hp": !!int "33"
"hit_dice": "6d8 + 6"
"stats":
- !!int "14"
- !!int "8"
- !!int "13"
- !!int "5"
- !!int "8"
- !!int "5"
"speed": "30 ft."
"damage_resistances": "cold, fire, lightning"
"damage_immunities": "poison"
"condition_immunities": "[charmed](/3-Mechanics/CLI/rules/conditions.md#charmed),\
  \ [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened), [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)"
"senses": "darkvision 60 ft., passive Perception 9"
"languages": "understands Abyssal but can't speak"
"cr": "1"
"traits":
- "desc": "When it reduces a creature to 0 hit points with a melee attack on its turn,\
    \ the maw demon can take a bonus action to move up to half its speed and make\
    \ a bite attack."
  "name": "Rampage"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+4 (+4) to hit, reach 5 ft., one target.\
    \ Hit: dice:2d8 + 2|text(11) (2d8 + 2) piercing damage."
  "name": "Bite"
"source":
- "VGM"
- "GoS"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/VGM/Maw%20Demon.webp"
```
^statblock

```encounter-table
name: Maw Demon
creatures:
 - 1: Maw Demon
```

## Environment

underdark