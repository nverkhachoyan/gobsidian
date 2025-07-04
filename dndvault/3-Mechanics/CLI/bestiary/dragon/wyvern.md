---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mm
- monster/cr/6
- monster/environment/hill
- monster/environment/mountain
- monster/size/large
- monster/type/dragon
aliases: ["Wyvern"]
NoteIcon: monster
BestiaryType: dragon
SourceType: Bestiary
BookSource: Monster Manual p. 303, Dragon of Icespire Peak, Storm Lord's Wrath. Available in the SRD and the Basic Rules.
---
# [Wyvern](3-Mechanics\CLI\bestiary\dragon/wyvern.md)
*Source: Monster Manual p. 303, Dragon of Icespire Peak, Storm Lord's Wrath. Available in the SRD and the Basic Rules.*  

Travelers in the wild sometimes look to the skies to see the dark-winged shape of a wyvern carrying its prey. These cousins to the great dragons hunt the same tangled forests and caverns as their kin. Their appearance sends ripples of alarm through the borderlands of civilization.

A wyvern has two scaly legs, leathery wings, and a sinewy tail topped with its most potent weapon: a poison stinger. The poison in a wyvern's stinger can kill a creature in seconds. Extremely potent, wyvern poison burns through its victim's bloodstream, disintegrating veins and arteries on its way to the heart. As deadly as wyverns can be, however, hunters and adventurers often track them to claim the venom, which is used in alchemical compounds and to coat weapons.

## Aerial Hunters

A wyvern doesn't fight on the ground unless it can't reach its prey by any other means, or if it has been fooled into a position from which aerial combat isn't an option. If forced into a confrontation on the ground, a wyvern crouches low, keeping its stinger poised above its head as it hisses and growls.

## Aggressive and Reckless

A wyvern intent on its prey backs down only if it sustains serious injury, or if its prey eludes it long enough for another easier potential meal to wander along. If it corners a fleeing creature in an enclosure too small to enter, a wyvern guards where the quarry hides, lashing with its stinger whenever opportunity allows.

Although they possess more cunning than ordinary beasts, wyverns lack the intelligence of their draconic cousins. As such, creatures that maintain their composure as a wyvern hunts them from the air can often elude or trick it. Wyverns follow a direct path to their prey, with no thought given to possible ambushes.

## Tamed Wyverns

A wyvern can be tamed for use as a mount, but doing so presents a difficult and deadly challenge. Raising one as a hatchling offers the best results. However, a wyvern's violent temperament has cost the life of many a would-be master.

```statblock
"name": "Wyvern"
"size": "Large"
"type": "dragon"
"alignment": "Unaligned"
"ac": !!int "13"
"ac_class": "natural armor"
"hp": !!int "110"
"hit_dice": "13d10 + 39"
"stats":
- !!int "19"
- !!int "10"
- !!int "16"
- !!int "5"
- !!int "12"
- !!int "6"
"speed": "20 ft., fly 80 ft."
"skillsaves":
  "Perception": !!int "4"
"senses": "darkvision 60 ft., passive Perception 14"
"languages": ""
"cr": "6"
"actions":
- "desc": "The wyvern makes two attacks: one with its bite and one with its stinger.\
    \ While flying, it can use its claws in place of one other attack."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+7 (+7) to hit, reach 10 ft., one creature.\
    \ Hit: dice:2d6 + 4|text(11) (2d6 + 4) piercing damage."
  "name": "Bite"
- "desc": "Melee Weapon Attack: dice: d20+7 (+7) to hit, reach 5 ft., one target.\
    \ Hit: dice:2d8 + 4|text(13) (2d8 + 4) slashing damage."
  "name": "Claws"
- "desc": "Melee Weapon Attack: dice: d20+7 (+7) to hit, reach 10 ft., one creature.\
    \ Hit: dice:2d6 + 4|text(11) (2d6 + 4) piercing damage. The target must\
    \ make a DC 15 Constitution saving throw, taking dice:7d6|text(24) (7d6) poison\
    \ damage on a failed save, or half as much damage on a successful one."
  "name": "Stinger"
"source":
- "MM"
- "HotDQ"
- "PotA"
- "SKT"
- "WDMM"
- "DIP"
- "SLW"
- "BGDIA"
- "CM"
- "JttRC"
- "ToFW"
- "BMT"
- "QftIS"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MM/Wyvern.webp"
```
^statblock

```encounter-table
name: Wyvern
creatures:
 - 1: Wyvern
```

## Environment

mountain, hill