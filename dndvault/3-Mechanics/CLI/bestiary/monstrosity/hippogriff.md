---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mm
- monster/cr/1
- monster/environment/grassland
- monster/environment/hill
- monster/environment/mountain
- monster/size/large
- monster/type/monstrosity
aliases: ["Hippogriff"]
NoteIcon: monster
BestiaryType: monstrosity
SourceType: Bestiary
BookSource: Monster Manual p. 184. Available in the SRD and the Basic Rules.
---
# [Hippogriff](3-Mechanics\CLI\bestiary\monstrosity/hippogriff.md)
*Source: Monster Manual p. 184. Available in the SRD and the Basic Rules.*  

A beast whose magical origins are lost to history, a hippogriff possesses the wings and forelimbs of an eagle, the hindquarters of a horse, and a head that combines the features of both animals.

Reclusive and omnivorous, hippogriffs mate for life and seldom venture more than a few miles from their nest. When defending its mate or its young, a hippogriff fights to the death. Hippogriffs don't lay eggs but give birth to live young.

Dragons, griffons, and wyverns have a taste for hippogriff meat and frequently prey on these creatures.

## Flying Mounts

A hippogriff raised in captivity can be trained to be a faithful companion and mount. Of all the creatures that can serve as flying mounts, hippogriffs are among the easiest to train and the most loyal once trained properly.

```statblock
"name": "Hippogriff"
"size": "Large"
"type": "monstrosity"
"alignment": "Unaligned"
"ac": !!int "11"
"hp": !!int "19"
"hit_dice": "3d10 + 3"
"stats":
- !!int "17"
- !!int "13"
- !!int "13"
- !!int "2"
- !!int "12"
- !!int "8"
"speed": "40 ft., fly 60 ft."
"skillsaves":
  "Perception": !!int "5"
"senses": "passive Perception 15"
"languages": ""
"cr": "1"
"traits":
- "desc": "The hippogriff has advantage on Wisdom ([Perception](/3-Mechanics/CLI/rules/skills.md#Perception))\
    \ checks that rely on sight."
  "name": "Keen Sight"
"actions":
- "desc": "The hippogriff makes two attacks: one with its beak and one with its claws."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+5 (+5) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d10 + 3|text(8) (1d10 + 3) piercing damage."
  "name": "Beak"
- "desc": "Melee Weapon Attack: dice: d20+5 (+5) to hit, reach 5 ft., one target.\
    \ Hit: dice:2d6 + 3|text(10) (2d6 + 3) slashing damage."
  "name": "Claws"
"source":
- "MM"
- "PotA"
- "SKT"
- "WDH"
- "ERLW"
- "EGW"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MM/Hippogriff.webp"
```
^statblock

```encounter-table
name: Hippogriff
creatures:
 - 1: Hippogriff
```

## Environment

mountain, grassland, hill