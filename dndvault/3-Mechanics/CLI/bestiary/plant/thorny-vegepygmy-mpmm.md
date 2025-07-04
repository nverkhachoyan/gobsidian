---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/1
- monster/environment/forest
- monster/environment/swamp
- monster/size/medium
- monster/type/plant
aliases: ["Thorny Vegepygmy"]
NoteIcon: monster
BestiaryType: plant
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 253, Volo's Guide to Monsters p. 197
---
# [Thorny Vegepygmy](3-Mechanics\CLI\bestiary\plant/thorny-vegepygmy-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 253, Volo's Guide to Monsters p. 197*  

```statblock
"name": "Thorny Vegepygmy (MPMM)"
"size": "Medium"
"type": "plant"
"alignment": "Typically  Neutral"
"ac": !!int "14"
"ac_class": "natural armor"
"hp": !!int "27"
"hit_dice": "5d8 + 5"
"stats":
- !!int "13"
- !!int "12"
- !!int "13"
- !!int "2"
- !!int "10"
- !!int "6"
"speed": "30 ft."
"skillsaves":
  "Stealth": !!int "3"
  "Perception": !!int "4"
"damage_resistances": "lightning, piercing"
"senses": "darkvision 60 ft., passive Perception 14"
"languages": ""
"cr": "1"
"traits":
- "desc": "The thorny has advantage on Dexterity ([Stealth](/3-Mechanics/CLI/rules/skills.md#Stealth))\
    \ checks it makes in any terrain with ample obscuring vegetation."
  "name": "Plant Camouflage"
- "desc": "The thorny regains 5 hit points at the start of its turn. If it takes cold,\
    \ fire, or necrotic damage, this trait doesn't function at the start of the thorny's\
    \ next turn. The thorny dies only if it starts its turn with 0 hit points and\
    \ doesn't regenerate."
  "name": "Regeneration"
- "desc": "At the start of its turn, the thorny deals dice:1d4|text(2) (1d4) piercing\
    \ damage to any creature grappling it."
  "name": "Thorny Body"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+3 (+3) to hit, reach 5 ft., one target.\
    \ Hit: dice:2d8 + 1|text(10) (2d8 + 1) piercing damage."
  "name": "Bite"
"source":
- "MPMM"
- "VGM"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Thorny%20Vegepygmy.webp"
```
^statblock

```encounter-table
name: Thorny Vegepygmy
creatures:
 - 1: Thorny Vegepygmy
```

## Environment

forest, swamp