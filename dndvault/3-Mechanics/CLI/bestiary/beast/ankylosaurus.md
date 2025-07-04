---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mm
- monster/cr/3
- monster/environment/grassland
- monster/size/huge
- monster/type/beast
aliases: ["Ankylosaurus"]
NoteIcon: monster
BestiaryType: beast
SourceType: Bestiary
BookSource: Monster Manual p. 79. Available in the Basic Rules.
---
# [Ankylosaurus](3-Mechanics\CLI\bestiary\beast/ankylosaurus.md)
*Source: Monster Manual p. 79. Available in the Basic Rules.*  

Thick armor plating covers the body of the plant-eating ankylosaurus, which defends itself against predators with a knobbed tail that delivers a devastating strike. Some varieties of ankylosaurus have spiked tails that deal piercing damage instead of bludgeoning damage.

## Dinosaurs

Dinosaurs, or behemoths, are among the oldest reptiles in the world. Predatory dinosaurs are savage, territorial hunters. Herbivorous dinosaurs are less aggressive, but they might attack to defend their young, or if startled or harassed.

Dinosaurs come in many sizes and shapes. Larger varieties often have drab coloration, while smaller dinosaurs have colorful markings akin to birds. Dinosaurs roam rugged and isolated areas that humanoids seldom visit, including remote mountain valleys, inaccessible plateaus, tropical islands, and deep fens.

```statblock
"name": "Ankylosaurus"
"size": "Huge"
"type": "beast"
"alignment": "Unaligned"
"ac": !!int "15"
"ac_class": "natural armor"
"hp": !!int "68"
"hit_dice": "8d12 + 16"
"stats":
- !!int "19"
- !!int "11"
- !!int "15"
- !!int "2"
- !!int "12"
- !!int "5"
"speed": "30 ft."
"senses": "passive Perception 11"
"languages": ""
"cr": "3"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+7 (+7) to hit, reach 10 ft., one target.\
    \ Hit: dice:4d6 + 4|text(18) (4d6 + 4) bludgeoning damage. If the target\
    \ is a creature, it must succeed on a DC 14 Strength saving throw or be knocked\
    \ [prone](/3-Mechanics/CLI/rules/conditions.md#prone)."
  "name": "Tail"
"source":
- "MM"
- "ToA"
- "CRCotN"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MM/Ankylosaurus.webp"
```
^statblock

```encounter-table
name: Ankylosaurus
creatures:
 - 1: Ankylosaurus
```

## Environment

grassland