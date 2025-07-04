---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/tce
- monster/cr/
- monster/size/medium
- monster/type/elemental
aliases: ["Elemental Spirit (6th-level Spell)"]
NoteIcon: monster
BestiaryType: elemental
SourceType: Bestiary
BookSource: Tasha's Cauldron of Everything p. 111
---
# [Elemental Spirit (6th-level Spell)](3-Mechanics\CLI\bestiary\elemental/elemental-spirit-6th-level-spell-tce.md)
*Source: Tasha's Cauldron of Everything p. 111*  

```statblock
"name": "Elemental Spirit (6th-level Spell) (TCE)"
"size": "Medium"
"type": "elemental"
"alignment": "Unaligned"
"ac": !!int "17"
"ac_class": "natural armor"
"hp": !!int "70"
"stats":
- !!int "18"
- !!int "15"
- !!int "17"
- !!int "4"
- !!int "10"
- !!int "16"
"speed": "40 ft., burrow 40 ft. (earth only), fly 40 ft. (air only; hover), swim 40\
  \ ft. (water only)"
"damage_resistances": "lightning, thunder (Air only)"
"damage_immunities": "poison; fire (Fire only)"
"condition_immunities": "[exhaustion](/3-Mechanics/CLI/rules/conditions.md#exhaustion),\
  \ [paralyzed](/3-Mechanics/CLI/rules/conditions.md#paralyzed), [petrified](/3-Mechanics/CLI/rules/conditions.md#petrified),\
  \ [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned), [unconscious](/3-Mechanics/CLI/rules/conditions.md#unconscious)"
"senses": "darkvision 60 ft., passive Perception 10"
"languages": "Primordial, understands the languages you speak"
"traits":
- "desc": "The elemental can move through a space as narrow as 1 inch wide without\
    \ squeezing."
  "name": "Amorphous Form (Air, Fire, and Water Only)"
"actions":
- "desc": "The elemental makes a number of attacks equal to half this spell's level\
    \ (rounded down)."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: the summoner's spell attack modifier to hit, reach\
    \ 5 ft., one target. Hit: 1d10 + 4 + summonSpellLevel bludgeoning damage (Air,\
    \ Earth, and Water only) or fire damage (Fire only)."
  "name": "Slam"
"source":
- "TCE"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/TCE/Elemental%20Spirit.webp"
```
^statblock

```encounter-table
name: Elemental Spirit (6th-level Spell)
creatures:
 - 1: Elemental Spirit (6th-level Spell)
```