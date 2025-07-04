---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/tce
- monster/cr/
- monster/size/medium
- monster/type/undead
aliases: ["Undead Spirit (Ghostly, 9th-Level Spell)"]
NoteIcon: monster
BestiaryType: undead
SourceType: Bestiary
BookSource: Tasha's Cauldron of Everything p. 114
---
# [Undead Spirit (Ghostly, 9th-Level Spell)](3-Mechanics\CLI\bestiary\undead/undead-spirit-ghostly-9th-level-spell-tce.md)
*Source: Tasha's Cauldron of Everything p. 114*  

```statblock
"name": "Undead Spirit (Ghostly, 9th-Level Spell) (TCE)"
"size": "Medium"
"type": "undead"
"alignment": "Unaligned"
"ac": !!int "20"
"ac_class": "natural armor"
"hp": !!int "90"
"stats":
- !!int "12"
- !!int "16"
- !!int "15"
- !!int "4"
- !!int "10"
- !!int "9"
"speed": "30 ft., fly 40 ft. (ghostly only; hover)"
"damage_immunities": "necrotic, poison"
"condition_immunities": "[exhaustion](/3-Mechanics/CLI/rules/conditions.md#exhaustion),\
  \ [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened), [paralyzed](/3-Mechanics/CLI/rules/conditions.md#paralyzed),\
  \ [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)"
"senses": "darkvision 60 ft., passive Perception 10"
"languages": "understands the languages you speak"
"traits":
- "desc": "The spirit can move through other creatures and objects as if they were\
    \ difficult terrain. If it ends its turn inside an object, it is shunted to the\
    \ nearest unoccupied space and takes dice: 1d10|avg|noform (1d10) force damage\
    \ for every 5 feet traveled."
  "name": "Incorporeal Passage (Ghostly Only)"
- "desc": "Any creature, other than you, that starts its turn within 5 feet of the\
    \ spirit must succeed on a Constitution saving throw against your spell save DC\
    \ or be [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned) until the start\
    \ of its next turn."
  "name": "Festering Aura (Putrid Only)"
"actions":
- "desc": "The spirit makes a number of attacks equal to half this spell's level (rounded\
    \ down)."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: the summoner's spell attack modifier to hit, reach\
    \ 5 ft., one creature. Hit: 1d8 + 3 + summonSpellLevel necrotic damage, and\
    \ the creature must succeed on a Wisdom saving throw against your spell save DC\
    \ or be [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened) of the undead\
    \ until the end of the target's next turn."
  "name": "Deathly Touch (Ghostly Only)"
- "desc": "Ranged Spell Attack: the summoner's spell attack modifier to hit, range\
    \ 150 ft., one target. Hit: 2d4 + 3 + summonSpellLevel necrotic damage."
  "name": "Grave Bolt (Skeletal Only)"
- "desc": "Melee Weapon Attack: the summoner's spell attack modifier to hit, reach\
    \ 5 ft., one target. Hit: 1d6 + 3 + summonSpellLevel slashing damage. If the\
    \ target is [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned), it must\
    \ succeed on a Constitution saving throw against your spell save DC or be [paralyzed](/3-Mechanics/CLI/rules/conditions.md#paralyzed)\
    \ until the end of its next turn."
  "name": "Rotting Claw (Putrid Only)"
"source":
- "TCE"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/TCE/Undead%20Spirit.webp"
```
^statblock

```encounter-table
name: Undead Spirit (Ghostly, 9th-Level Spell)
creatures:
 - 1: Undead Spirit (Ghostly, 9th-Level Spell)
```