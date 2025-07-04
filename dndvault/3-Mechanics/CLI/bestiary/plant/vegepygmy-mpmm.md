---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/1-4
- monster/environment/forest
- monster/environment/swamp
- monster/size/small
- monster/type/plant
aliases: ["Vegepygmy"]
NoteIcon: monster
BestiaryType: plant
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 252, Volo's Guide to Monsters p. 196
---
# [Vegepygmy](3-Mechanics\CLI\bestiary\plant/vegepygmy-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 252, Volo's Guide to Monsters p. 196*  

```statblock
"name": "Vegepygmy (MPMM)"
"size": "Small"
"type": "plant"
"alignment": "Typically  Neutral"
"ac": !!int "13"
"ac_class": "natural armor"
"hp": !!int "13"
"hit_dice": "3d6 + 3"
"stats":
- !!int "7"
- !!int "14"
- !!int "13"
- !!int "6"
- !!int "11"
- !!int "7"
"speed": "30 ft."
"skillsaves":
  "Stealth": !!int "4"
  "Perception": !!int "2"
"damage_resistances": "lightning, piercing"
"senses": "darkvision 60 ft., passive Perception 12"
"languages": "Vegepygmy"
"cr": "1/4"
"traits":
- "desc": "The vegepygmy has advantage on Dexterity ([Stealth](/3-Mechanics/CLI/rules/skills.md#Stealth))\
    \ checks it makes in any terrain with ample obscuring vegetation."
  "name": "Plant Camouflage"
- "desc": "The vegepygmy regains 3 hit points at the start of its turn. If it takes\
    \ cold, fire, or necrotic damage, this trait doesn't function at the start of\
    \ the vegepygmy's next turn. The vegepygmy dies only if it starts its turn with\
    \ 0 hit points and doesn't regenerate."
  "name": "Regeneration"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+4 (+4) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d6 + 2|text(5) (1d6 + 2) slashing damage."
  "name": "Claws"
- "desc": "Ranged Weapon Attack: dice: d20+4 (+4) to hit, range 30/120 ft.,\
    \ one target. Hit: dice:1d4 + 2|text(4) (1d4 + 2) bludgeoning damage."
  "name": "Sling"
"source":
- "MPMM"
- "VGM"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Vegepygmy.webp"
```
^statblock

```encounter-table
name: Vegepygmy
creatures:
 - 1: Vegepygmy
```

## Environment

forest, swamp