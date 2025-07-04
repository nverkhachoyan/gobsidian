---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/5
- monster/environment/forest
- monster/size/medium
- monster/type/plant
aliases: ["Wood Woad"]
NoteIcon: monster
BestiaryType: plant
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 266, Volo's Guide to Monsters p. 198
---
# [Wood Woad](3-Mechanics\CLI\bestiary\plant/wood-woad-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 266, Volo's Guide to Monsters p. 198*  

```statblock
"name": "Wood Woad (MPMM)"
"size": "Medium"
"type": "plant"
"alignment": "Typically  Lawful Neutral"
"ac": !!int "18"
"ac_class": "natural armor, [shield](/3-Mechanics/CLI/items/shield.md)"
"hp": !!int "75"
"hit_dice": "10d8 + 30"
"stats":
- !!int "18"
- !!int "12"
- !!int "16"
- !!int "10"
- !!int "13"
- !!int "8"
"speed": "30 ft., climb 30 ft."
"skillsaves":
  "Athletics": !!int "7"
  "Stealth": !!int "4"
  "Perception": !!int "4"
"damage_vulnerabilities": "fire"
"damage_resistances": "bludgeoning, piercing"
"condition_immunities": "[charmed](/3-Mechanics/CLI/rules/conditions.md#charmed),\
  \ [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened)"
"senses": "darkvision 60 ft., passive Perception 14"
"languages": "Sylvan"
"cr": "5"
"traits":
- "desc": "The wood woad has advantage on Dexterity ([Stealth](/3-Mechanics/CLI/rules/skills.md#Stealth))\
    \ checks it makes in any terrain with ample obscuring vegetation."
  "name": "Plant Camouflage"
- "desc": "The wood woad regains 10 hit points at the start of its turn if it is in\
    \ contact with the ground. If the wood woad takes fire damage, this trait doesn't\
    \ function at the start of the wood woad's next turn. The wood woad dies only\
    \ if it starts its turn with 0 hit points and doesn't regenerate."
  "name": "Regeneration"
- "desc": "Once on each of its turns, the wood woad can use 10 feet of its movement\
    \ to step magically into one living tree within 5 feet of it and emerge from a\
    \ second living tree within 60 feet of it that it can see, appearing in an unoccupied\
    \ space within 5 feet of the second tree. Both trees must be Large or bigger."
  "name": "Tree Stride"
"actions":
- "desc": "The wood woad makes two Club attacks."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+7 (+7) to hit, reach 5 ft., one target.\
    \ Hit: dice:4d4 + 4|text(14) (4d4 + 4) force damage."
  "name": "Club"
"source":
- "MPMM"
- "VGM"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Wood%20Woad.webp"
```
^statblock

```encounter-table
name: Wood Woad
creatures:
 - 1: Wood Woad
```

## Environment

forest