---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/5
- monster/environment/desert
- monster/size/large
- monster/type/monstrosity
aliases: ["Tlincalli"]
NoteIcon: monster
BestiaryType: monstrosity
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 242, Volo's Guide to Monsters p. 193
---
# [Tlincalli](3-Mechanics\CLI\bestiary\monstrosity/tlincalli-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 242, Volo's Guide to Monsters p. 193*  

```statblock
"name": "Tlincalli (MPMM)"
"size": "Large"
"type": "monstrosity"
"alignment": "Typically  Neutral"
"ac": !!int "15"
"ac_class": "natural armor"
"hp": !!int "85"
"hit_dice": "10d10 + 30"
"stats":
- !!int "16"
- !!int "13"
- !!int "16"
- !!int "8"
- !!int "12"
- !!int "8"
"speed": "40 ft."
"skillsaves":
  "Stealth": !!int "4"
  "Perception": !!int "4"
  "Survival": !!int "4"
"senses": "darkvision 60 ft., passive Perception 14"
"languages": "Tlincalli"
"cr": "5"
"actions":
- "desc": "The tlincalli makes one Longsword or Spiked Chain attack and one Sting\
    \ attack."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+6 (+6) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d8 + 3|text(7) (1d8 + 3) slashing damage, or dice:1d10 + 3|text(8)\
    \ (1d10 + 3) slashing damage if used with two hands."
  "name": "Longsword"
- "desc": "Melee Weapon Attack: dice: d20+6 (+6) to hit, reach 10 ft., one target.\
    \ Hit: dice:1d6 + 3|text(6) (1d6 + 3) piercing damage, and the target is\
    \ [grappled](/3-Mechanics/CLI/rules/conditions.md#grappled) (escape DC 11) if\
    \ it is a Large or smaller creature. Until this grapple ends, the target is [restrained](/3-Mechanics/CLI/rules/conditions.md#restrained),\
    \ and the tlincalli can't use the spiked chain against another target."
  "name": "Spiked Chain"
- "desc": "Melee Weapon Attack: dice: d20+6 (+6) to hit, reach 5 ft., one creature.\
    \ Hit: dice:1d6 + 3|text(6) (1d6 + 3) piercing damage plus dice:4d6|text(14)\
    \ (4d6) poison damage, and the target must succeed on a DC 14 Constitution saving\
    \ throw or be [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned) for 1\
    \ minute. If it fails the saving throw by 5 or more, the target is also [paralyzed](/3-Mechanics/CLI/rules/conditions.md#paralyzed)\
    \ while [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned). The target\
    \ can repeat the saving throw at the end of each of its turns, ending the effect\
    \ on itself on a success."
  "name": "Sting"
"source":
- "MPMM"
- "VGM"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Tlincalli.webp"
```
^statblock

```encounter-table
name: Tlincalli
creatures:
 - 1: Tlincalli
```

## Environment

desert