---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/1-8
- monster/environment/hill
- monster/environment/underdark
- monster/size/small
- monster/type/monstrosity
aliases: ["Xvart"]
NoteIcon: monster
BestiaryType: monstrosity
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 267, Volo's Guide to Monsters p. 200
---
# [Xvart](3-Mechanics\CLI\bestiary\monstrosity/xvart-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 267, Volo's Guide to Monsters p. 200*  

```statblock
"name": "Xvart (MPMM)"
"size": "Small"
"type": "monstrosity"
"alignment": "Typically  Chaotic Evil"
"ac": !!int "13"
"ac_class": "[leather armor](/3-Mechanics/CLI/items/leather-armor.md)"
"hp": !!int "7"
"hit_dice": "2d6"
"stats":
- !!int "8"
- !!int "14"
- !!int "10"
- !!int "8"
- !!int "7"
- !!int "7"
"speed": "30 ft."
"skillsaves":
  "Stealth": !!int "4"
"senses": "darkvision 30 ft., passive Perception 8"
"languages": "Abyssal"
"cr": "1/8"
"traits":
- "desc": "The xvart can communicate with ordinary [bats](/3-Mechanics/CLI/bestiary/beast/bat.md)\
    \ and [rats](/3-Mechanics/CLI/bestiary/beast/rat.md), as well as [giant bats](/3-Mechanics/CLI/bestiary/beast/giant-bat.md)\
    \ and [giant rats](/3-Mechanics/CLI/bestiary/beast/giant-rat.md)."
  "name": "Raxivort's Tongue"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+4 (+4) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d6 + 2|text(5) (1d6 + 2) piercing damage. If at least one of\
    \ the xvart's allies is within 5 feet of the target, the xvart can push the target\
    \ 5 feet if the target is a Medium or smaller creature."
  "name": "Shortsword"
- "desc": "Ranged Weapon Attack: dice: d20+4 (+4) to hit, range 30/120 ft.,\
    \ one target. Hit: dice:1d4 + 2|text(4) (1d4 + 2) bludgeoning damage."
  "name": "Sling"
"bonus_actions":
- "desc": "The xvart takes the [Disengage](/3-Mechanics/CLI/rules/actions.md#Disengage)\
    \ action."
  "name": "Low Cunning"
"source":
- "MPMM"
- "VGM"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Xvart.webp"
```
^statblock

```encounter-table
name: Xvart
creatures:
 - 1: Xvart
```

## Environment

hill, underdark