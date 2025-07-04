---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/vgm
- monster/cr/1-4
- monster/environment/arctic
- monster/environment/forest
- monster/environment/grassland
- monster/environment/hill
- monster/size/medium
- monster/type/undead
aliases: ["Gnoll Witherling"]
NoteIcon: monster
BestiaryType: undead
SourceType: Bestiary
BookSource: Volo's Guide to Monsters p. 155
---
# [Gnoll Witherling](3-Mechanics\CLI\bestiary\undead/gnoll-witherling-vgm.md)
*Source: Volo's Guide to Monsters p. 155*  

Sometimes gnolls turn against each other, perhaps to determine who rules a war band or because of extreme starvation. Even under ordinary circumstances, gnolls that are deprived of victims for too long can't control their hunger and violent urges. Eventually, they fight among themselves.

The survivors devour the flesh of their slain comrades but preserve the bones. Then, by invoking rituals to Yeenoghu, they bring the remains back to a semblance of life in the form of a gnoll witherling.

Witherlings act much as gnolls do in life, traveling with their comrades and trying to kill anything in their path. They don't eat and aren't motivated by hunger, leaving more flesh for the rest of the war band. Gnoll witherlings are incapable of wielding any weapon more sophisticated than a simple club.

## Undead Nature

A gnoll witherling doesn't require air, food, drink, or sleep.

I heard a warlock sought to magically turn the last witherlings of a destroyed gnoll war band into his own strike force against foes. That worked well until Yeenoghu sent seven war bands after him.

-Volo

```statblock
"name": "Gnoll Witherling (VGM)"
"size": "Medium"
"type": "undead"
"alignment": "Chaotic Evil"
"ac": !!int "12"
"ac_class": "natural armor"
"hp": !!int "11"
"hit_dice": "2d8 + 2"
"stats":
- !!int "14"
- !!int "8"
- !!int "12"
- !!int "5"
- !!int "5"
- !!int "5"
"speed": "30 ft."
"damage_immunities": "poison"
"condition_immunities": "[exhaustion](/3-Mechanics/CLI/rules/conditions.md#exhaustion),\
  \ [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)"
"senses": "darkvision 60 ft., passive Perception 7"
"languages": "understands Gnoll but can't speak"
"cr": "1/4"
"traits":
- "desc": "When the witherling reduces a creature to 0 hit points with a melee attack\
    \ on its turn, it can take a bonus action to move up to half its speed and make\
    \ a bite attack."
  "name": "Rampage"
"actions":
- "desc": "The witherling makes two attacks: one with its bite and one with its club,\
    \ or two with its club."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+4 (+4) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d4 + 2|text(4) (1d4 + 2) piercing damage."
  "name": "Bite"
- "desc": "Melee Weapon Attack: dice: d20+4 (+4) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d4 + 2|text(4) (1d4 + 2) bludgeoning damage."
  "name": "Club"
"reactions":
- "desc": "In response to a gnoll being reduced to 0 hit points within 30 feet of\
    \ the witherling, the witherling makes a melee attack."
  "name": "Vengeful Strike"
"source":
- "VGM"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/VGM/Gnoll%20Witherling.webp"
```
^statblock

```encounter-table
name: Gnoll Witherling
creatures:
 - 1: Gnoll Witherling
```

## Environment

grassland, forest, hill, arctic