---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/4
- monster/environment/forest
- monster/size/large
- monster/type/monstrosity
aliases: ["Girallon"]
NoteIcon: monster
BestiaryType: monstrosity
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 139, Volo's Guide to Monsters p. 152
---
# [Girallon](3-Mechanics\CLI\bestiary\monstrosity/girallon-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 139, Volo's Guide to Monsters p. 152*  

```statblock
"name": "Girallon (MPMM)"
"size": "Large"
"type": "monstrosity"
"alignment": "Unaligned"
"ac": !!int "13"
"hp": !!int "59"
"hit_dice": "7d10 + 21"
"stats":
- !!int "18"
- !!int "16"
- !!int "16"
- !!int "5"
- !!int "12"
- !!int "7"
"speed": "40 ft., climb 40 ft."
"skillsaves":
  "Stealth": !!int "5"
  "Perception": !!int "5"
"senses": "darkvision 60 ft., passive Perception 15"
"languages": ""
"cr": "4"
"actions":
- "desc": "The girallon makes one Bite attack and four Claw attacks."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+6 (+6) to hit, reach 5 ft., one creature.\
    \ Hit: dice:1d6 + 4|text(7) (1d6 + 4) piercing damage."
  "name": "Bite"
- "desc": "Melee Weapon Attack: dice: d20+6 (+6) to hit, reach 10 ft., one target.\
    \ Hit: dice:1d6 + 4|text(7) (1d6 + 4) slashing damage."
  "name": "Claw"
"bonus_actions":
- "desc": "The girallon moves up to its speed toward a hostile creature that it can\
    \ see."
  "name": "Aggressive"
"source":
- "MPMM"
- "VGM"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Girallon.webp"
```
^statblock

```encounter-table
name: Girallon
creatures:
 - 1: Girallon
```

## Environment

forest