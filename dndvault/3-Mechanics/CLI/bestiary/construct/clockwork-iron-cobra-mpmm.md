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
aliases: ["Clockwork Iron Cobra"]
NoteIcon: monster
BestiaryType: construct
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 79, Mordenkainen's Tome of Foes p. 125
---
# [Clockwork Iron Cobra](3-Mechanics\CLI\bestiary\construct/clockwork-iron-cobra-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 79, Mordenkainen's Tome of Foes p. 125*  

```statblock
"name": "Clockwork Iron Cobra (MPMM)"
"size": "Medium"
"type": "construct"
"alignment": "Unaligned"
"ac": !!int "13"
"hp": !!int "91"
"hit_dice": "14d8 + 28"
"stats":
- !!int "12"
- !!int "16"
- !!int "14"
- !!int "3"
- !!int "10"
- !!int "1"
"speed": "30 ft."
"skillsaves":
  "Stealth": !!int "7"
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
- "desc": "Melee Weapon Attack: dice: d20+5 (+5) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d6 + 3|text(6) (1d6 + 3) piercing damage. If the target is\
    \ a creature, it must succeed on a DC 13 Constitution saving throw or suffer one\
    \ random effect (roll a dice: d6|avg|noform (d6)):\n\n- 1–2 Confusion.\
    \ On its next turn, the target must use its action to make one weapon attack against\
    \ a random creature it can see within 30 feet of it, using whatever weapon it\
    \ has in hand and moving beforehand if necessary to get in range. If it's holding\
    \ no weapon, it makes an unarmed strike. If no creature is visible within 30 feet,\
    \ it takes the [Dash](/3-Mechanics/CLI/rules/actions.md#Dash) action, moving toward\
    \ the nearest creature.  \n- 3–4 Paralysis. The target is [paralyzed](/3-Mechanics/CLI/rules/conditions.md#paralyzed)\
    \ until the end of its next turn.  \n- 5–6 Poison. The target takes dice:3d8|text(13)\
    \ (3d8) poison damage.  "
  "name": "Bite"
"source":
- "MPMM"
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Clockwork%20Iron%20Cobra.webp"
```
^statblock

```encounter-table
name: Clockwork Iron Cobra
creatures:
 - 1: Clockwork Iron Cobra
```

## Environment

forest, grassland, hill, mountain