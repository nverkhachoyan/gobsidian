---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/3
- monster/size/medium
- monster/type/fiend/demon
aliases: ["Bulezau"]
NoteIcon: monster
BestiaryType: fiend (demon)
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 67, Mordenkainen's Tome of Foes p. 131
---
# [Bulezau](3-Mechanics\CLI\bestiary\fiend/bulezau-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 67, Mordenkainen's Tome of Foes p. 131*  

```statblock
"name": "Bulezau (MPMM)"
"size": "Medium"
"type": "fiend"
"subtype": "demon"
"alignment": "Typically  Chaotic Evil"
"ac": !!int "14"
"ac_class": "natural armor"
"hp": !!int "52"
"hit_dice": "7d8 + 21"
"stats":
- !!int "15"
- !!int "14"
- !!int "17"
- !!int "8"
- !!int "9"
- !!int "6"
"speed": "40 ft."
"damage_resistances": "cold, fire, lightning"
"damage_immunities": "poison"
"condition_immunities": "[charmed](/3-Mechanics/CLI/rules/conditions.md#charmed),\
  \ [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened), [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)"
"senses": "darkvision 120 ft., passive Perception 9"
"languages": "Abyssal, telepathy 60 ft."
"cr": "3"
"traits":
- "desc": "When any creature that isn't a demon starts its turn within 30 feet of\
    \ the bulezau, that creature must succeed on a DC 13 Constitution saving throw\
    \ or take dice:1d6|text(3) (1d6) necrotic damage."
  "name": "Rotting Presence"
- "desc": "The bulezau's long jump is up to 20 feet and its high jump is up to 10\
    \ feet, with or without a running start."
  "name": "Standing Leap"
- "desc": "The bulezau has advantage on Strength and Dexterity saving throws made\
    \ against effects that would knock it [prone](/3-Mechanics/CLI/rules/conditions.md#prone)."
  "name": "Sure-Footed"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+4 (+4) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d12 + 2|text(8) (1d12 + 2) piercing damage plus dice:1d8|text(4)\
    \ (1d8) necrotic damage. If the target is a creature, it must succeed on a DC\
    \ 13 Constitution saving throw against disease or become [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)\
    \ until the disease ends. While [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)\
    \ in this way, the target sports festering boils, coughs up flies, and sheds rotting\
    \ skin, and the target must repeat the saving throw after every 24 hours that\
    \ elapse. On a successful save, the disease ends. On a failed save, the target's\
    \ hit point maximum is reduced by dice:1d8|text(4) (1d8). The target dies\
    \ if its hit point maximum is reduced to 0."
  "name": "Barbed Tail"
"source":
- "MPMM"
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Bulezau.webp"
```
^statblock

```encounter-table
name: Bulezau
creatures:
 - 1: Bulezau
```