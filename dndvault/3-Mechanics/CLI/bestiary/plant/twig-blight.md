---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mm
- monster/cr/1-8
- monster/environment/forest
- monster/size/small
- monster/type/plant
aliases: ["Twig Blight"]
NoteIcon: monster
BestiaryType: plant
SourceType: Bestiary
BookSource: Monster Manual p. 32, Dragon of Icespire Peak. Available in the Basic Rules.
---
# [Twig Blight](3-Mechanics\CLI\bestiary\plant/twig-blight.md)
*Source: Monster Manual p. 32, Dragon of Icespire Peak. Available in the Basic Rules.*  

Twig blights can root in soil, which they do when living prey are scarce. While rooted, they resemble woody shrubs. When it pulls its roots free of the ground to move, a twig blight's branches twist together to form a humanoid-looking body with a head and limbs.

Twig blights seek out campsites and watering holes, rooting there to set up ambushes for potential victims coming to drink or rest. Huddled together in groups, twig blights blend in with an area's natural vegetation or with piles of debris or firewood.

Given how dry they are, twig blights are particularly susceptible to fire.

## Blights

Awakened plants gifted with the powers of intelligence and mobility, blights plague lands contaminated by darkness. Drinking that darkness from the soil, a blight carries out the will of ancient evil and attempts to spread that evil wherever it can.

### Roots of the Gulthias Tree

Legends tell of a vampire named Gulthias who worked terrible magic and raised up an abominable tower called Nightfang Spire. Gulthias was undone when a hero plunged a wooden stake through his heart, but as the vampire was destroyed, his blood infused the stake with a dreadful power. In time, tendrils of new growth sprouted from the wood, growing into a sapling infused with the vampire's evil essence. It is said that a mad druid discovered the sapling, transplanting it to an underground grotto where it could grow. From this Gulthias tree came the seeds from which the first blights were sown.

### Dark Conquest

Wherever a tree or plant is contaminated by a fragment of an evil mind or power, a Gulthias tree can rise to infest and corrupt the surrounding forest. Its evil spreads through root and soil to other plants, which perish or transform into blights. As those blights spread, they poison and uproot healthy plants, replacing them with brambles, toxic weeds, and others of their kind. In time, an infestation of blights can turn any land or forest into a place of corruption.

In forests infested with blights, trees and plants grow with supernatural speed. Vines and undergrowth rapidly spread through buildings and overrun trails and roads. After blights have killed or driven off their inhabitants, whole villages can disappear in the space of days.

### Controlled by Evil

Blights are independent creatures, but most act under a Gulthias tree's control, often displaying the habits and traits of the life force or spirit that spawned them. By attacking their progenitor's old foes or seeking out treasures valuable to it, they carry on the legacy of long-lost evil.

> [!quote]-  
> 
> Behold the legacy of Gulthias the vampire: plants with a taste for blood.


```statblock
"name": "Twig Blight"
"size": "Small"
"type": "plant"
"alignment": "Neutral Evil"
"ac": !!int "13"
"ac_class": "natural armor"
"hp": !!int "4"
"hit_dice": "1d6 + 1"
"stats":
- !!int "6"
- !!int "13"
- !!int "12"
- !!int "4"
- !!int "8"
- !!int "3"
"speed": "20 ft."
"skillsaves":
  "Stealth": !!int "3"
"damage_vulnerabilities": "fire"
"condition_immunities": "[blinded](/3-Mechanics/CLI/rules/conditions.md#blinded),\
  \ [deafened](/3-Mechanics/CLI/rules/conditions.md#deafened)"
"senses": "blindsight 60 ft. (blind beyond this radius), passive Perception 9"
"languages": "understands Common but can't speak"
"cr": "1/8"
"traits":
- "desc": "While the blight remains motionless, it is indistinguishable from a dead\
    \ shrub."
  "name": "False Appearance"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+3 (+3) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d4 + 1|text(3) (1d4 + 1) piercing damage."
  "name": "Claws"
"source":
- "MM"
- "CoS"
- "LMoP"
- "TftYP"
- "GoS"
- "RMBRE"
- "WBtW"
- "PSI"
- "HftT"
- "PaBTSO"
- "DIP"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MM/Twig%20Blight.webp"
```
^statblock

```encounter-table
name: Twig Blight
creatures:
 - 1: Twig Blight
```

## Environment

forest