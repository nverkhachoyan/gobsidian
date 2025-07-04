---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/23
- monster/size/huge
- monster/type/fiend/demon
aliases: ["Baphomet"]
NoteIcon: monster
BestiaryType: fiend (demon)
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 58, Mordenkainen's Tome of Foes p. 143
---
# [Baphomet](3-Mechanics\CLI\bestiary\npc/baphomet-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 58, Mordenkainen's Tome of Foes p. 143*  

```statblock
"name": "Baphomet (MPMM)"
"size": "Huge"
"type": "fiend"
"subtype": "demon"
"alignment": "Chaotic Evil"
"ac": !!int "22"
"ac_class": "natural armor"
"hp": !!int "319"
"hit_dice": "22d12 + 176"
"stats":
- !!int "30"
- !!int "14"
- !!int "26"
- !!int "18"
- !!int "24"
- !!int "16"
"speed": "40 ft."
"saves":
  "Dexterity": !!int "9"
  "Wisdom": !!int "14"
  "Constitution": !!int "15"
"skillsaves":
  "Intimidation": !!int "17"
  "Perception": !!int "14"
"damage_resistances": "cold, fire, lightning"
"damage_immunities": "poison; bludgeoning, piercing, slashing that is nonmagical"
"condition_immunities": "[charmed](/3-Mechanics/CLI/rules/conditions.md#charmed),\
  \ [exhaustion](/3-Mechanics/CLI/rules/conditions.md#exhaustion), [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened),\
  \ [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)"
"senses": "truesight 120 ft., passive Perception 24"
"languages": "all, telepathy 120 ft."
"cr": "23"
"traits":
- "desc": "Baphomet casts one of the following spells, requiring no material components\
    \ and using Charisma as the spellcasting ability (spell save DC 18):\n\n1/day:\
    \ [teleport](/3-Mechanics/CLI/spells/teleport.md)\n\n3/day each: [dispel magic](/3-Mechanics/CLI/spells/dispel-magic.md),\
    \ [dominate beast](/3-Mechanics/CLI/spells/dominate-beast.md), [maze](/3-Mechanics/CLI/spells/maze.md),\
    \ [wall of stone](/3-Mechanics/CLI/spells/wall-of-stone.md)"
  "name": "Spellcasting"
- "desc": "Baphomet can perfectly recall any path he has traveled, and he is immune\
    \ to the [maze](/3-Mechanics/CLI/spells/maze.md) spell."
  "name": "Labyrinthine Recall"
- "desc": "If Baphomet fails a saving throw, he can choose to succeed instead."
  "name": "Legendary Resistance (3/Day)"
- "desc": "Baphomet has advantage on saving throws against spells and other magical\
    \ effects."
  "name": "Magic Resistance"
"actions":
- "desc": "Baphomet makes one Bite attack, one Gore attack, and one Heartcleaver attack.\
    \ He also uses Frightful Presence."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+17 (+17) to hit, reach 10 ft., one\
    \ target. Hit: dice:2d8 + 10|text(19) (2d8 + 10) piercing damage."
  "name": "Bite"
- "desc": "Melee Weapon Attack: dice: d20+17 (+17) to hit, reach 10 ft., one\
    \ target. Hit: dice:2d6 + 10|text(17) (2d6 + 10) piercing damage. If Baphomet\
    \ moved at least 10 feet straight toward the target immediately before the hit,\
    \ the target takes an extra dice:3d10|text(16) (3d10) piercing damage. If\
    \ the target is a creature, it must succeed on a DC 25 Strength saving throw or\
    \ be pushed up to 10 feet away and knocked [prone](/3-Mechanics/CLI/rules/conditions.md#prone)."
  "name": "Gore"
- "desc": "Melee Weapon Attack: dice: d20+17 (+17) to hit, reach 15 ft., one\
    \ target. Hit: dice:2d10 + 10|text(21) (2d10 + 10) force damage."
  "name": "Heartcleaver"
- "desc": "Each creature of Baphomet's choice within 120 feet of him and aware of\
    \ him must succeed on a DC 18 Wisdom saving throw or become [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened)\
    \ for 1 minute. A [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened)\
    \ creature can repeat the saving throw at the end of each of its turns, ending\
    \ the effect on itself on a success. These later saves have disadvantage if Baphomet\
    \ is within line of sight of the creature.\n\nIf a creature succeeds on any of\
    \ these saves or the effect ends on it, the creature is immune to Baphomet's Frightful\
    \ Presence for the next 24 hours."
  "name": "Frightful Presence"
"legendary_actions":
- "desc": "Baphomet makes one Heartcleaver attack."
  "name": "Heartcleaver Attack"
- "desc": "Baphomet moves up to his speed without provoking opportunity attacks, then\
    \ makes a Gore attack."
  "name": "Charge (Costs 2 Actions)"
"source":
- "MPMM"
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Baphomet.webp"
```
^statblock

```encounter-table
name: Baphomet
creatures:
 - 1: Baphomet
```