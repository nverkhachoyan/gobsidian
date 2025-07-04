---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/tce
- monster/cr/
- monster/size/medium
- monster/type/aberration
aliases: ["Aberrant Spirit (8th-level Spell)"]
NoteIcon: monster
BestiaryType: aberration
SourceType: Bestiary
BookSource: Tasha's Cauldron of Everything p. 109
---
# [Aberrant Spirit (8th-level Spell)](3-Mechanics\CLI\bestiary\aberration/aberrant-spirit-8th-level-spell-tce.md)
*Source: Tasha's Cauldron of Everything p. 109*  

```statblock
"name": "Aberrant Spirit (8th-level Spell) (TCE)"
"size": "Medium"
"type": "aberration"
"alignment": "Unaligned"
"ac": !!int "19"
"ac_class": "natural armor"
"hp": !!int "80"
"stats":
- !!int "16"
- !!int "10"
- !!int "15"
- !!int "16"
- !!int "10"
- !!int "6"
"speed": "30 ft., fly 30 ft. (beholderkin only; hover)"
"damage_immunities": "psychic"
"senses": "darkvision 60 ft., passive Perception 10"
"languages": "Deep Speech, understands the languages you speak"
"traits":
- "desc": "The aberration regains 5 hit points at the start of its turn if it has\
    \ at least 1 hit point."
  "name": "Regeneration (Slaad Only)"
- "desc": "At the start of each of the aberration's turns, each creature within 5\
    \ feet of the aberration must succeed on a Wisdom saving throw against your spell\
    \ save DC or take dice: 2d6|avg|noform (2d6) psychic damage, provided that\
    \ the aberration isn't [incapacitated](/3-Mechanics/CLI/rules/conditions.md#incapacitated)."
  "name": "Whispering Aura (Star Spawn Only)"
"actions":
- "desc": "The aberration makes a number of attacks equal to half this spell's level\
    \ (rounded down)."
  "name": "Multiattack"
- "desc": "Ranged Spell Attack: the summoner's spell attack modifier to hit, range\
    \ 150 ft., one creature. Hit: 1d8 + 3 + summonSpellLevel psychic damage."
  "name": "Eye Ray (Beholderkin Only)"
- "desc": "Melee Weapon Attack: the summoner's spell attack modifier to hit, reach\
    \ 5 ft., one target. Hit: 1d10 + 3 + summonSpellLevel slashing damage. If\
    \ the target is a creature, it can't regain hit points until the start of the\
    \ aberration's next turn."
  "name": "Claws (Slaad Only)"
- "desc": "Melee Spell Attack: the summoner's spell attack modifier to hit, reach\
    \ 5 ft., one creature. Hit: 1d8 + 3 + summonSpellLevel psychic damage."
  "name": "Psychic Slam (Star Spawn Only)"
"source":
- "TCE"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/TCE/Aberrant%20Spirit.webp"
```
^statblock

```encounter-table
name: Aberrant Spirit (8th-level Spell)
creatures:
 - 1: Aberrant Spirit (8th-level Spell)
```