---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/1
- monster/environment/hill
- monster/environment/underdark
- monster/size/small
- monster/type/monstrosity
aliases: ["Xvart Warlock of Raxivort"]
NoteIcon: monster
BestiaryType: monstrosity
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 267, Volo's Guide to Monsters p. 200
---
# [Xvart Warlock of Raxivort](3-Mechanics\CLI\bestiary\monstrosity/xvart-warlock-of-raxivort-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 267, Volo's Guide to Monsters p. 200*  

```statblock
"name": "Xvart Warlock of Raxivort (MPMM)"
"size": "Small"
"type": "monstrosity"
"alignment": "Typically  Chaotic Evil"
"ac": !!int "12"
"hp": !!int "22"
"hit_dice": "5d6 + 5"
"stats":
- !!int "8"
- !!int "14"
- !!int "12"
- !!int "8"
- !!int "11"
- !!int "12"
"speed": "30 ft."
"skillsaves":
  "Stealth": !!int "3"
"senses": "darkvision 30 ft., passive Perception 10"
"languages": "Abyssal"
"cr": "1"
"traits":
- "desc": "The xvart casts one of the following spells, requiring no material components\
    \ and using Charisma as the spellcasting ability (spell save DC 11):\n\nAt will:\
    \ [detect magic](/3-Mechanics/CLI/spells/detect-magic.md), [mage armor](/3-Mechanics/CLI/spells/mage-armor.md)\
    \ (self only), [mage hand](/3-Mechanics/CLI/spells/mage-hand.md), [minor illusion](/3-Mechanics/CLI/spells/minor-illusion.md),\
    \ [prestidigitation](/3-Mechanics/CLI/spells/prestidigitation.md)\n\n1/day each:\
    \ [burning hands](/3-Mechanics/CLI/spells/burning-hands.md), [invisibility](/3-Mechanics/CLI/spells/invisibility.md)"
  "name": "Spellcasting"
- "desc": "When the xvart reduces an enemy to 0 hit points, the xvart gains 4 temporary\
    \ hit points."
  "name": "Raxivort's Blessing"
- "desc": "The xvart can communicate with ordinary [bats](/3-Mechanics/CLI/bestiary/beast/bat.md)\
    \ and [rats](/3-Mechanics/CLI/bestiary/beast/rat.md), as well as [giant bats](/3-Mechanics/CLI/bestiary/beast/giant-bat.md)\
    \ and [giant rats](/3-Mechanics/CLI/bestiary/beast/giant-rat.md)."
  "name": "Raxivort's Tongue"
"actions":
- "desc": "The xvart makes two Scimitar or Raxivort's Bite attacks."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+4 (+4) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d6 + 2|text(5) (1d6 + 2) slashing damage."
  "name": "Scimitar"
- "desc": "Ranged Spell Attack: dice: d20+3 (+3) to hit, range 30 ft., one creature.\
    \ Hit: dice:1d10 + 2|text(7) (1d10 + 2) poison damage."
  "name": "Raxivort's Bite"
"bonus_actions":
- "desc": "The xvart takes the [Disengage](/3-Mechanics/CLI/rules/actions.md#Disengage)\
    \ action."
  "name": "Low Cunning"
"source":
- "MPMM"
- "VGM"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Xvart%20Warlock%20of%20Raxivort.webp"
```
^statblock

```encounter-table
name: Xvart Warlock of Raxivort
creatures:
 - 1: Xvart Warlock of Raxivort
```

## Environment

hill, underdark