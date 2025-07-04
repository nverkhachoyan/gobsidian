---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mm
- monster/cr/1-4
- monster/environment/grassland
- monster/environment/hill
- monster/size/large
- monster/type/beast
aliases: ["Axe Beak"]
NoteIcon: monster
BestiaryType: beast
SourceType: Bestiary
BookSource: Monster Manual p. 317. Available in the SRD and the Basic Rules.
---
# [Axe Beak](3-Mechanics\CLI\bestiary\beast/axe-beak.md)
*Source: Monster Manual p. 317. Available in the SRD and the Basic Rules.*  

An axe beak is a tall flightless bird with strong legs and a heavy, wedge-shaped beak. It has a nasty disposition and tends to attack any unfamiliar creature that wanders too close.

```statblock
"name": "Axe Beak"
"size": "Large"
"type": "beast"
"alignment": "Unaligned"
"ac": !!int "11"
"hp": !!int "19"
"hit_dice": "3d10 + 3"
"stats":
- !!int "14"
- !!int "12"
- !!int "12"
- !!int "2"
- !!int "10"
- !!int "5"
"speed": "50 ft."
"senses": "passive Perception 10"
"languages": ""
"cr": "1/4"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+4 (+4) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d8 + 2|text(6) (1d8 + 2) slashing damage."
  "name": "Beak"
"source":
- "MM"
- "SKT"
- "ToA"
- "WDH"
- "GoS"
- "IDRotF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MM/Axe%20Beak.webp"
```
^statblock

```encounter-table
name: Axe Beak
creatures:
 - 1: Axe Beak
```

## Environment

grassland, hill