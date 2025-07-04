---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/tce
- monster/cr/
- monster/size/small
- monster/type/beast
aliases: ["Bestial Spirit (Water, 6th-Level Spell)"]
NoteIcon: monster
BestiaryType: beast
SourceType: Bestiary
BookSource: Tasha's Cauldron of Everything p. 109
---
# [Bestial Spirit (Water, 6th-Level Spell)](3-Mechanics\CLI\bestiary\beast/bestial-spirit-water-6th-level-spell-tce.md)
*Source: Tasha's Cauldron of Everything p. 109*  

```statblock
"name": "Bestial Spirit (Water, 6th-Level Spell) (TCE)"
"size": "Small"
"type": "beast"
"alignment": "Unaligned"
"ac": !!int "17"
"ac_class": "natural armor"
"hp": !!int "50"
"stats":
- !!int "18"
- !!int "11"
- !!int "16"
- !!int "4"
- !!int "14"
- !!int "5"
"speed": "30 ft., climb 30 ft. (land only), fly 60 ft. (air only), swim 30 ft. (water\
  \ only)"
"senses": "darkvision 60 ft., passive Perception 12"
"languages": "understands the languages you speak"
"traits":
- "desc": "The beast can breathe only underwater."
  "name": "Water Breathing (Water Only)"
- "desc": "The beast doesn't provoke opportunity attacks when it flies out of an enemy's\
    \ reach."
  "name": "Flyby (Air Only)"
- "desc": "The beast has advantage on an attack roll against a creature if at least\
    \ one of the beast's allies is within 5 feet of the creature and the ally isn't\
    \ [incapacitated](/3-Mechanics/CLI/rules/conditions.md#incapacitated)."
  "name": "Pack Tactics (Land and Water Only)"
"actions":
- "desc": "The beast makes a number of attacks equal to half this spell's level (rounded\
    \ down)."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: the summoner's spell attack modifier to hit, reach\
    \ 5 ft., one target. Hit: 1d8 + 4 + summonSpellLevel piercing damage."
  "name": "Maul"
"source":
- "TCE"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/TCE/Bestial%20Spirit.webp"
```
^statblock

```encounter-table
name: Bestial Spirit (Water, 6th-Level Spell)
creatures:
 - 1: Bestial Spirit (Water, 6th-Level Spell)
```