---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/11
- monster/environment/underdark
- monster/environment/urban
- monster/size/medium
- monster/type/fiend/demon
aliases: ["Alkilith"]
NoteIcon: monster
BestiaryType: fiend (demon)
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 44, Mordenkainen's Tome of Foes p. 130
---
# [Alkilith](3-Mechanics\CLI\bestiary\fiend/alkilith-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 44, Mordenkainen's Tome of Foes p. 130*  

```statblock
"name": "Alkilith (MPMM)"
"size": "Medium"
"type": "fiend"
"subtype": "demon"
"alignment": "Typically  Chaotic Evil"
"ac": !!int "17"
"ac_class": "natural armor"
"hp": !!int "168"
"hit_dice": "16d8 + 96"
"stats":
- !!int "12"
- !!int "19"
- !!int "22"
- !!int "6"
- !!int "11"
- !!int "7"
"speed": "40 ft., climb 40 ft."
"saves":
  "Dexterity": !!int "8"
  "Constitution": !!int "10"
"skillsaves":
  "Stealth": !!int "8"
"damage_resistances": "acid; cold; fire; lightning; bludgeoning, piercing, slashing\
  \ from nonmagical attacks"
"damage_immunities": "poison"
"condition_immunities": "[charmed](/3-Mechanics/CLI/rules/conditions.md#charmed),\
  \ [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened), [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)"
"senses": "darkvision 120 ft., passive Perception 10"
"languages": "understands Abyssal but can't speak"
"cr": "11"
"traits":
- "desc": "If the alkilith surrounds a door, window, or similar opening continuously\
    \ for dice: 6d6|avg|noform (6d6) days, the opening becomes a permanent portal\
    \ to a random layer of the Abyss."
  "name": "Abyssal Rift"
- "desc": "The alkilith can move through a space as narrow as 1 inch wide without\
    \ squeezing."
  "name": "Amorphous"
- "desc": "If the alkilith is motionless at the start of combat, it has advantage\
    \ on its initiative roll. Moreover, if a creature hasn't observed the alkilith\
    \ move or act, that creature must succeed on a DC 18 Intelligence ([Investigation](/3-Mechanics/CLI/rules/skills.md#Investigation))\
    \ check to discern that the alkilith isn't ordinary slime or fungus."
  "name": "False Appearance"
- "desc": "Any creature that isn't a demon that starts its turn within 30 feet of\
    \ the alkilith must succeed on a DC 18 Wisdom saving throw, or it hears a faint\
    \ buzzing in its head for a moment and has disadvantage on its next attack roll,\
    \ saving throw, or ability check.\n\nIf the saving throw against Foment Confusion\
    \ fails by 5 or more, the creature is instead subjected to the [confusion](/3-Mechanics/CLI/spells/confusion.md)\
    \ spell for 1 minute (no [concentration](/3-Mechanics/CLI/rules/conditions.md#concentration)\
    \ required by the alkilith). While under the effect of that confusion, the creature\
    \ is immune to Foment Confusion."
  "name": "Foment Confusion"
- "desc": "The alkilith has advantage on saving throws against spells and other magical\
    \ effects."
  "name": "Magic Resistance"
- "desc": "The alkilith can climb difficult surfaces, such as upside down on ceilings,\
    \ without making an ability check."
  "name": "Spider Climb"
- "desc": "The alkilith doesn't require air, food, drink, or sleep."
  "name": "Unusual Nature"
"actions":
- "desc": "The alkilith makes three Tentacle attacks."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+8 (+8) to hit, reach 15 ft., one target.\
    \ Hit: dice:4d6 + 4|text(18) (4d6 + 4) acid damage."
  "name": "Tentacle"
"source":
- "MPMM"
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Alkilith.webp"
```
^statblock

```encounter-table
name: Alkilith
creatures:
 - 1: Alkilith
```

## Environment

underdark, urban