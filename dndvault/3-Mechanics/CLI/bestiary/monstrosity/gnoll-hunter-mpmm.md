---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/1-2
- monster/environment/arctic
- monster/environment/forest
- monster/environment/grassland
- monster/environment/hill
- monster/size/medium
- monster/type/monstrosity
aliases: ["Gnoll Hunter"]
NoteIcon: monster
BestiaryType: monstrosity
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 144, Volo's Guide to Monsters p. 154
---
# [Gnoll Hunter](3-Mechanics\CLI\bestiary\monstrosity/gnoll-hunter-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 144, Volo's Guide to Monsters p. 154*  

```statblock
"name": "Gnoll Hunter (MPMM)"
"size": "Medium"
"type": "monstrosity"
"alignment": "Typically  Chaotic Evil"
"ac": !!int "13"
"ac_class": "[leather armor](/3-Mechanics/CLI/items/leather-armor.md)"
"hp": !!int "22"
"hit_dice": "4d8 + 4"
"stats":
- !!int "14"
- !!int "14"
- !!int "12"
- !!int "8"
- !!int "12"
- !!int "8"
"speed": "30 ft."
"skillsaves":
  "Stealth": !!int "4"
  "Perception": !!int "3"
"senses": "darkvision 60 ft., passive Perception 13"
"languages": "Gnoll"
"cr": "1/2"
"actions":
- "desc": "The gnoll makes two Bite, Spear, or Longbow attacks."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+4 (+4) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d4 + 2|text(4) (1d4 + 2) piercing damage."
  "name": "Bite"
- "desc": "Melee or Ranged Weapon Attack: dice: d20+4 (+4) to hit, reach 5 ft.\
    \ or range 20/60 ft., one target. Hit: dice:1d6 + 2|text(5) (1d6 + 2) piercing\
    \ damage, or dice:1d8 + 2|text(6) (1d8 + 2) piercing damage when used with\
    \ two hands to make a melee attack."
  "name": "Spear"
- "desc": "Ranged Weapon Attack: dice: d20+4 (+4) to hit, range 150/600 ft.,\
    \ one target. Hit: dice:1d8 + 2|text(6) (1d8 + 2) piercing damage, and the\
    \ target's speed is reduced by 10 feet until the end of its next turn."
  "name": "Longbow"
"bonus_actions":
- "desc": "After the gnoll reduces a creature to 0 hit points with a melee attack\
    \ on its turn, the gnoll moves up to half its speed and makes a Bite attack."
  "name": "Rampage"
"source":
- "MPMM"
- "VGM"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Gnoll%20Hunter.webp"
```
^statblock

```encounter-table
name: Gnoll Hunter
creatures:
 - 1: Gnoll Hunter
```

## Environment

arctic, forest, grassland, hill