---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mm
- monster/cr/2
- monster/environment/forest
- monster/environment/grassland
- monster/size/large
- monster/type/monstrosity
aliases: ["Centaur"]
NoteIcon: monster
BestiaryType: monstrosity
SourceType: Bestiary
BookSource: Monster Manual p. 38, Dragon of Icespire Peak. Available in the SRD and the Basic Rules.
---
# [Centaur](3-Mechanics\CLI\bestiary\monstrosity/centaur.md)
*Source: Monster Manual p. 38, Dragon of Icespire Peak. Available in the SRD and the Basic Rules.*  

> [!quote]- A quote from Batley Summerfoot, a halfling adventurer who never read HOOVES OF FURY, by Iryil Grayborn of Sundown  
> 
> I hear centaurs make excellent mounts!

Reclusive wanderers and omen-readers of the wild, centaurs avoid conflict but fight fiercely when pressed. They roam the vast wilderness, keeping far from borders, laws, and the company of other creatures.

## Wilderness Nomads

Centaur tribes range across lands with mild to hot climates, where a centaur requires only light furs or oiled skins to deal with inclement weather. They are hunter-gatherers and rarely build shelters or even use tents.

Centaur migrations span continents and take decades to repeat, so that a centaur tribe might not retread the same path for generations. These long-ranging patterns can lead to conflict when centaurs encounter settlements of other creatures built along their traditional routes.

## Reluctant Settlers

A centaur that can't keep pace with the rest of its tribe is left behind. Some such centaurs vanish into the wilderness and are never seen again. Those that can bear the loss of their tribe might take up residence among other races. Frontier settlements value the nature knowledge of their centaur residents. Many such communities owe their survival to the insight and acumen of a centaur.

Despite their reclusive nature, centaurs trade with elves and with the caravans of other benevolent humanoids they meet during their wanderings. A trader might save the life of a wounded or an elderly centaur unfit for long travel, escorting it to a settlement where it can peacefully live out the rest of its days.

> [!quote]- A quote from Batley Summerfoot, a halfling adventurer who never read Hooves of Fury by Irvil Grayborn of Sundown  
> 
> I hear centaurs make excellent mounts!


```statblock
"name": "Centaur"
"size": "Large"
"type": "monstrosity"
"alignment": "Neutral Good"
"ac": !!int "12"
"hp": !!int "45"
"hit_dice": "6d10 + 12"
"stats":
- !!int "18"
- !!int "14"
- !!int "14"
- !!int "9"
- !!int "13"
- !!int "11"
"speed": "50 ft."
"skillsaves":
  "Athletics": !!int "6"
  "Perception": !!int "3"
  "Survival": !!int "3"
"senses": "passive Perception 13"
"languages": "Elvish, Sylvan"
"cr": "2"
"traits":
- "desc": "If the centaur moves at least 30 feet straight toward a target and then\
    \ hits it with a pike attack on the same turn, the target takes an extra dice:3d6|text(10)\
    \ (3d6) piercing damage."
  "name": "Charge"
"actions":
- "desc": "The centaur makes two attacks: one with its pike and one with its hooves\
    \ or two with its longbow."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+6 (+6) to hit, reach 10 ft., one target.\
    \ Hit: dice:1d10 + 4|text(9) (1d10 + 4) piercing damage."
  "name": "Pike"
- "desc": "Melee Weapon Attack: dice: d20+6 (+6) to hit, reach 5 ft., one target.\
    \ Hit: dice:2d6 + 4|text(11) (2d6 + 4) bludgeoning damage."
  "name": "Hooves"
- "desc": "Ranged Weapon Attack: dice: d20+4 (+4) to hit, range 150/600 ft.,\
    \ one target. Hit: dice:1d8 + 2|text(6) (1d8 + 2) piercing damage."
  "name": "Longbow"
"source":
- "MM"
- "SKT"
- "TftYP"
- "GoS"
- "DIP"
- "MOT"
- "WBtW"
- "QftIS"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MM/Centaur.webp"
```
^statblock

```encounter-table
name: Centaur
creatures:
 - 1: Centaur
```

## Environment

grassland, forest