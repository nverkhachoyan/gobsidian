---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/13
- monster/environment/underdark
- monster/environment/urban
- monster/size/medium
- monster/type/monstrosity
aliases: ["Angry Sorrowsworn"]
NoteIcon: monster
BestiaryType: monstrosity
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 222, Mordenkainen's Tome of Foes p. 231
---
# [Angry Sorrowsworn](3-Mechanics\CLI\bestiary\monstrosity/angry-sorrowsworn-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 222, Mordenkainen's Tome of Foes p. 231*  

```statblock
"name": "Angry Sorrowsworn (MPMM)"
"size": "Medium"
"type": "monstrosity"
"alignment": "Typically  Neutral Evil"
"ac": !!int "18"
"ac_class": "natural armor"
"hp": !!int "255"
"hit_dice": "30d8 + 120"
"stats":
- !!int "17"
- !!int "10"
- !!int "19"
- !!int "8"
- !!int "13"
- !!int "6"
"speed": "30 ft."
"skillsaves":
  "Perception": !!int "11"
"damage_resistances": "bludgeoning, piercing, slashing while in dim light or darkness"
"senses": "darkvision 60 ft., passive Perception 21"
"languages": "Common"
"cr": "13"
"traits":
- "desc": "The sorrowsworn has advantage on saving throws against being [blinded](/3-Mechanics/CLI/rules/conditions.md#blinded),\
    \ [charmed](/3-Mechanics/CLI/rules/conditions.md#charmed), [deafened](/3-Mechanics/CLI/rules/conditions.md#deafened),\
    \ [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened), [stunned](/3-Mechanics/CLI/rules/conditions.md#stunned),\
    \ or knocked [unconscious](/3-Mechanics/CLI/rules/conditions.md#unconscious)."
  "name": "Two Heads"
- "desc": "If another creature deals damage to the sorrowsworn, the sorrowsworn's\
    \ attack rolls have advantage until the end of its next turn, and the first time\
    \ it hits with a Hook attack on its next turn, the attack's target takes an extra\
    \ dice:3d12|text(19) (3d12) psychic damage.\n\nOn its turn, the sorrowsworn\
    \ has disadvantage on attack rolls if no other creature has dealt damage to it\
    \ since the end of its last turn."
  "name": "Rising Anger"
"actions":
- "desc": "The sorrowsworn makes two Hook attacks."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+8 (+8) to hit, reach 5 ft., one target.\
    \ Hit: dice:2d12 + 3|text(16) (2d12 + 3) piercing damage."
  "name": "Hook"
"source":
- "MPMM"
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Angry%20Sorrowsworn.webp"
```
^statblock

```encounter-table
name: Angry Sorrowsworn
creatures:
 - 1: Angry Sorrowsworn
```

## Environment

underdark, urban