---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mm
- monster/cr/1-8
- monster/environment/coastal
- monster/environment/desert
- monster/environment/forest
- monster/environment/grassland
- monster/environment/hill
- monster/environment/mountain
- monster/environment/swamp
- monster/environment/underdark
- monster/environment/urban
- monster/size/tiny
- monster/type/beast
aliases: ["Stirge"]
NoteIcon: monster
BestiaryType: beast
SourceType: Bestiary
BookSource: Monster Manual p. 284, Dragon of Icespire Peak. Available in the SRD and the Basic Rules.
---
# [Stirge](3-Mechanics\CLI\bestiary\beast/stirge.md)
*Source: Monster Manual p. 284, Dragon of Icespire Peak. Available in the SRD and the Basic Rules.*  

This horrid flying creature looks like a cross between a large bat and an oversized mosquito. Its legs end in sharp pincers, and its long, needle-like proboscis slashes the air as it seeks its next meal.

Stirges feed on the blood of living creatures, attaching and draining them slowly. Although they pose little danger in small numbers, packs of stirges can be a formidable threat, reattaching as quickly as their weakening prey can pluck them off.

## Blood Drain

A stirge attacks by landing on a victim, finding a vulnerable spot, and plunging its proboscis into the flesh while using its pincer legs to latch on to the victim. Once the stirge has sated itself, it detaches and flies off to digest its meal.

```statblock
"name": "Stirge"
"size": "Tiny"
"type": "beast"
"alignment": "Unaligned"
"ac": !!int "14"
"ac_class": "natural armor"
"hp": !!int "2"
"hit_dice": "1d4"
"stats":
- !!int "4"
- !!int "16"
- !!int "11"
- !!int "2"
- !!int "8"
- !!int "6"
"speed": "10 ft., fly 40 ft."
"senses": "darkvision 60 ft., passive Perception 9"
"languages": ""
"cr": "1/8"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+5 (+5) to hit, reach 5 ft., one creature.\
    \ Hit: dice:1d4 + 3|text(5) (1d4 + 3) piercing damage, and the stirge attaches\
    \ to the target. While attached, the stirge doesn't attack. Instead, at the start\
    \ of each of the stirge's turns, the target loses dice:1d4 + 3|text(5) (1d4\
    \ + 3) hit points due to blood loss.\n\nThe stirge can detach itself by spending\
    \ 5 feet of its movement. It does so after it drains 10 hit points of blood from\
    \ the target or the target dies. A creature, including the target, can use its\
    \ action to detach the stirge."
  "name": "Blood Drain"
"source":
- "MM"
- "HotDQ"
- "LMoP"
- "PotA"
- "TftYP"
- "ToA"
- "WDH"
- "WDMM"
- "GoS"
- "DIP"
- "BGDIA"
- "RMBRE"
- "IMR"
- "EGW"
- "DoSI"
- "KftGV"
- "HftT"
- "PaBTSO"
- "QftIS"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MM/Stirge.webp"
```
^statblock

```encounter-table
name: Stirge
creatures:
 - 1: Stirge
```

## Environment

grassland, forest, swamp, hill, urban, desert, coastal, mountain, underdark