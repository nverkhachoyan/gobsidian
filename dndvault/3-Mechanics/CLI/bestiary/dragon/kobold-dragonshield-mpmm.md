---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/1
- monster/environment/forest
- monster/environment/hill
- monster/environment/mountain
- monster/environment/underdark
- monster/size/small
- monster/type/dragon
aliases: ["Kobold Dragonshield"]
NoteIcon: monster
BestiaryType: dragon
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 163, Volo's Guide to Monsters p. 165
---
# [Kobold Dragonshield](3-Mechanics\CLI\bestiary\dragon/kobold-dragonshield-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 163, Volo's Guide to Monsters p. 165*  

```statblock
"name": "Kobold Dragonshield (MPMM)"
"size": "Small"
"type": "dragon"
"alignment": "Any alignment"
"ac": !!int "15"
"ac_class": "[leather](/3-Mechanics/CLI/items/leather-armor.md), [shield](/3-Mechanics/CLI/items/shield.md)"
"hp": !!int "44"
"hit_dice": "8d6 + 16"
"stats":
- !!int "12"
- !!int "15"
- !!int "14"
- !!int "8"
- !!int "9"
- !!int "10"
"speed": "20 ft."
"skillsaves":
  "Perception": !!int "1"
"damage_resistances": "see Dragon's Resistance below"
"senses": "darkvision 60 ft., passive Perception 11"
"languages": "Common, Draconic"
"cr": "1"
"traits":
- "desc": "The kobold has resistance to a type of damage based on the color of dragon\
    \ that invested it with power (choose or roll a dice: d10|avg|noform (d10)):\
    \ 1–2, acid (black or copper); 3–4, cold (silver or white); 5–6, fire (brass,\
    \ gold, or red); 7–8, lightning (blue or bronze); 9–10, poison (green)."
  "name": "Dragon's Resistance"
- "desc": "If the kobold is [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened)\
    \ or [paralyzed](/3-Mechanics/CLI/rules/conditions.md#paralyzed) by an effect\
    \ that allows a saving throw, it can repeat the save at the start of its turn\
    \ to end the effect on itself and all kobolds within 30 feet of it. Any kobold\
    \ that benefits from this trait (including the dragonshield) has advantage on\
    \ its next attack roll."
  "name": "Heart of the Dragon"
- "desc": "The kobold has advantage on an attack roll against a creature if at least\
    \ one of the kobold's allies is within 5 feet of the creature and the ally isn't\
    \ [incapacitated](/3-Mechanics/CLI/rules/conditions.md#incapacitated)."
  "name": "Pack Tactics"
- "desc": "While in sunlight, the kobold has disadvantage on attack rolls, as well\
    \ as on Wisdom ([Perception](/3-Mechanics/CLI/rules/skills.md#Perception)) checks\
    \ that rely on sight."
  "name": "Sunlight Sensitivity"
"actions":
- "desc": "The kobold makes two Spear attacks."
  "name": "Multiattack"
- "desc": "Melee or Ranged Weapon Attack: dice: d20+3 (+3) to hit, reach 5 ft.\
    \ or range 20/60 ft., one target. Hit: dice:1d6 + 1|text(4) (1d6 + 1) piercing\
    \ damage, or dice:1d8 + 1|text(5) (1d8 + 1) piercing damage if used with two\
    \ hands to make a melee attack."
  "name": "Spear"
"source":
- "MPMM"
- "VGM"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Kobold%20Dragonshield.webp"
```
^statblock

```encounter-table
name: Kobold Dragonshield
creatures:
 - 1: Kobold Dragonshield
```

## Environment

forest, hill, mountain, underdark