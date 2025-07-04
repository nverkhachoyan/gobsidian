---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mm
- monster/cr/4
- monster/environment/forest
- monster/environment/swamp
- monster/size/medium
- monster/type/humanoid/lizardfolk
aliases: ["Lizard King"]
NoteIcon: monster
BestiaryType: humanoid (lizardfolk)
SourceType: Bestiary
BookSource: Monster Manual p. 205, Tasha's Cauldron of Everything
---
# [Lizard King](3-Mechanics\CLI\bestiary\humanoid/lizard-king.md)
*Source: Monster Manual p. 205, Tasha's Cauldron of Everything*  

Lizardfolk are primitive reptilian humanoids that lurk in the swamps and jungles of the world. Their hut villages thrive in forbidding grottos, half-sunken ruins, and watery caverns.

## Territorial Xenophobes

Lizardfolk deal and trade with other races only rarely. Fiercely territorial, they use camouflaged scouts to guard the perimeter of their domain. When unwelcome visitors are detected, a tribe sends a hunting band to harass or drive the trespassers off, or tricks them into blundering into the lairs of crocodiles and other dangerous creatures.

Lizardfolk have no notion of traditional morality, and they find the concepts of good and evil utterly alien. Truly neutral creatures, they kill when it is expedient and do whatever it takes to survive.

Lizardfolk rarely stray beyond their claimed hunting grounds. Any creature that enters their territory is fair game to be stalked, killed, and devoured. They make no distinction between humanoids, beasts, and monsters. Similarly, lizardfolk don't like reaching too far beyond their borders, where they could easily become the hunted instead of the hunters.

Occasions might arise when lizardfolk will form alliances with their neighbors. These lizardfolk usually learn firsthand that humans, dwarves, halflings, and elves can sometimes prove helpful or trustworthy. Once lizardfolk forge ties with outsiders, they are steadfast and fierce allies.

## Great Feasts and Sacrifices

Lizardfolk are omnivorous, but they have a taste for humanoid flesh. Prisoners are often taken back to their camps to become the centerpieces of great feasts and rites involving dancing, storytelling, and ritual combat.

Victims are either cooked and eaten by the tribe, or are sacrificed to Semuanya, the lizardfolk god.

## Canny Crafters

Though they aren't skilled artisans, lizardfolk craft tools and ornamental jewelry out of the bones of their kills, and they use the hides and shells of dead monsters to create shields.

## Lizardfolk Leaders

Lizardfolk respect and fear magic with a religious awe. Lizardfolk shamans lead their tribes, overseeing rites and ceremonies performed to honor Semuanya. From time to time, however, a lizardfolk tribe produces a powerful figure touched not by Semuanya but by Sess'inek-a reptilian demon lord who seeks to corrupt and control the lizardfolk.

Lizardfolk born in Sess'inek's image are larger and more cunning than other lizardfolk, and are thoroughly evil. These lizard kings and queens dominate lizardfolk tribes, usurping a shaman's authority and inspiring uncharacteristic aggression among their subjects.

## Dragon Worshipers

Lizardfolk speak Draconic, which they are thought to have learned from dragons in ancient times. A tribe that wanders into the territory of a dragon will offer it tribute to win its favor. An evil dragon might exploit lizardfolk for its own vile ends, turning them into raiders and plunderers.

> [!quote]- A quote from A merchant's account of his experience with the lizardfolk tribes of the Lizard Marsh  
> 
> In all my dealings with the lizardfolk, I was never able to tell what they were thinking. Their reptilian eyes belied no hint of their intentions. I gave them supplies. They gave me the willies.


```statblock
"name": "Lizard King"
"size": "Medium"
"type": "humanoid"
"subtype": "lizardfolk"
"alignment": "Chaotic Evil"
"ac": !!int "15"
"ac_class": "natural armor"
"hp": !!int "78"
"hit_dice": "12d8 + 24"
"stats":
- !!int "17"
- !!int "12"
- !!int "15"
- !!int "11"
- !!int "11"
- !!int "15"
"speed": "30 ft., swim 30 ft."
"saves":
  "Wisdom": !!int "2"
  "Constitution": !!int "4"
"skillsaves":
  "Stealth": !!int "5"
  "Perception": !!int "4"
  "Survival": !!int "4"
"condition_immunities": "[frightened](/3-Mechanics/CLI/rules/conditions.md#frightened)"
"senses": "darkvision 60 ft., passive Perception 14"
"languages": "Abyssal, Draconic"
"cr": "4"
"traits":
- "desc": "The lizardfolk can hold its breath for 15 minutes."
  "name": "Hold Breath"
- "desc": "Once per turn, when the lizardfolk makes a melee attack with its trident\
    \ and hits, the target takes an extra dice:3d6|text(10) (3d6) damage, and\
    \ the lizardfolk gains temporary hit points equal to the extra damage dealt."
  "name": "Skewer"
"actions":
- "desc": "The lizardfolk makes two attacks: one with its bite and one with its claws\
    \ or trident or two melee attacks with its trident."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+5 (+5) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d6 + 3|text(6) (1d6 + 3) piercing damage."
  "name": "Bite"
- "desc": "Melee Weapon Attack: dice: d20+5 (+5) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d4 + 3|text(5) (1d4 + 3) slashing damage."
  "name": "Claws"
- "desc": "Melee or Ranged Weapon Attack: dice: d20+5 (+5) to hit, reach 5 ft.\
    \ or range 20/60 ft., one target. Hit: dice:1d6 + 3|text(6) (1d6 + 3) piercing\
    \ damage, or dice:1d8 + 3|text(7) (1d8 + 3) piercing damage if used with two\
    \ hands to make a melee attack."
  "name": "Trident"
"source":
- "MM"
- "PotA"
- "GoS"
- "TCE"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MM/Lizard%20King.webp"
```
^statblock

```encounter-table
name: Lizard King
creatures:
 - 1: Lizard King
```

## Environment

forest, swamp