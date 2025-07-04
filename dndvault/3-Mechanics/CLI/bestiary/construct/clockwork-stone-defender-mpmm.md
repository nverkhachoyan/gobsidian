---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/4
- monster/environment/forest
- monster/environment/grassland
- monster/environment/hill
- monster/environment/mountain
- monster/size/medium
- monster/type/construct
aliases: ["Clockwork Stone Defender"]
NoteIcon: monster
BestiaryType: construct
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 80, Mordenkainen's Tome of Foes p. 126
---
# [Clockwork Stone Defender](3-Mechanics\CLI\bestiary\construct/clockwork-stone-defender-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 80, Mordenkainen's Tome of Foes p. 126*  

```statblock
"name": "Clockwork Stone Defender (MPMM)"
"size": "Medium"
"type": "construct"
"alignment": "Unaligned"
"ac": !!int "16"
"ac_class": "natural armor"
"hp": !!int "105"
"hit_dice": "14d8 + 42"
"stats":
- !!int "19"
- !!int "10"
- !!int "17"
- !!int "3"
- !!int "10"
- !!int "1"
"speed": "30 ft."
"damage_immunities": "poison"
"condition_immunities": "[charmed](/3-Mechanics/CLI/rules/conditions.md#charmed),\
  \ [exhaustion](/3-Mechanics/CLI/rules/conditions.md#exhaustion), [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened),\
  \ [paralyzed](/3-Mechanics/CLI/rules/conditions.md#paralyzed), [petrified](/3-Mechanics/CLI/rules/conditions.md#petrified),\
  \ [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)"
"senses": "darkvision 60 ft., passive Perception 10"
"languages": "understands one language of its creator but can't speak"
"cr": "4"
"traits":
- "desc": "The clockwork has advantage on saving throws against spells and other magical\
    \ effects."
  "name": "Magic Resistance"
- "desc": "The clockwork doesn't require air, food, drink, or sleep."
  "name": "Unusual Nature"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+6 (+6) to hit, reach 5 ft., one target.\
    \ Hit: dice:2d6 + 4|text(11) (2d6 + 4) bludgeoning damage, and if the target\
    \ is Large or smaller, it is knocked [prone](/3-Mechanics/CLI/rules/conditions.md#prone)."
  "name": "Slam"
"reactions":
- "desc": "In response to another creature within 5 feet of it being hit by an attack\
    \ roll, the clockwork gives that creature a +5 bonus to its AC against that attack,\
    \ potentially causing a miss. To use this ability, the clockwork must be able\
    \ to see the creature and the attacker."
  "name": "Intercept Attack"
"source":
- "MPMM"
- "MTF"
- "RtG"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Clockwork%20Stone%20Defender.webp"
```
^statblock

```encounter-table
name: Clockwork Stone Defender
creatures:
 - 1: Clockwork Stone Defender
```

## Environment

forest, grassland, hill, mountain