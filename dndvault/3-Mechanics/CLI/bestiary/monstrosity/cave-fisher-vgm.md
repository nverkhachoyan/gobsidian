---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/vgm
- monster/cr/3
- monster/environment/underdark
- monster/size/medium
- monster/type/monstrosity
aliases: ["Cave Fisher"]
NoteIcon: monster
BestiaryType: monstrosity
SourceType: Bestiary
BookSource: Volo's Guide to Monsters p. 130
---
# [Cave Fisher](3-Mechanics\CLI\bestiary\monstrosity/cave-fisher-vgm.md)
*Source: Volo's Guide to Monsters p. 130*  

A cave fisher is a subterranean arachnid with a long snout that houses spinnerets, enabling the creature to produce sticky filament, much like the strands of a spider's webbing, which the creature uses to snag prey.

## Ambushers

A cave fisher usually hunts small animals and is fond of bats, so it stretches its filament over an opening that such prey might travel through. It then climbs to a hiding spot and adheres itself to the surface to rest and wait. When prey blunders into the filament, the cave fisher reels in its meal. A group of cave fishers might work together to cover a large area with filaments, but as soon as one captures potential food, every cave fisher in the area competes for the prize. If a victim escapes from the initial ambush, a cave fisher can reclaim its prey by shooting a filament out to capture it again.

A cave fisher instinctively knows that larger targets such as humanoids are more difficult to overcome, so the creatures shy away from attacking such prey unless they come across a solitary target. They might try to pick off a scout moving ahead of a group of travelers or a straggler lagging behind, rather than attracting the attention of the entire group.

## Moving Up in the World

Scarce food might draw a group of cave fishers up to the surface, into a shadowy canyon or a gloomy forest that features both native animal prey and creatures such as explorers or travelers occasionally moving through the area.

## Valuable Innards

Nearly every part of a cave fisher is useful after the creature has been dispatched. Its blood is alcoholic and tastes like strong liquor. Several dwarven spirits include cave fisher blood as part of the recipe, and some dwarves, especially berserkers, drink the blood straight. If they are gathered after being extruded, cave fisher filaments can be woven into rope that is thin, tough, and nearly invisible. Cave fisher meat is edible, tasting much like crab cooked in strong wine. The creature's shell is used in the manufacture of tools, armor, and jewelry.

## Reluctant Servants

While some folk hunt cave fishers to kill them for their filaments and their blood, others capture cave fisher eggs and rear the hatchlings. Cave fishers have a natural aversion to fire, since their blood is flammable. As such, chitines and hobgoblins sometimes use the threat of fire to train cave fishers, then employ them to guard passages or as beasts of war.

```statblock
"name": "Cave Fisher (VGM)"
"size": "Medium"
"type": "monstrosity"
"alignment": "Unaligned"
"ac": !!int "16"
"ac_class": "natural armor"
"hp": !!int "58"
"hit_dice": "9d8 + 18"
"stats":
- !!int "16"
- !!int "13"
- !!int "14"
- !!int "3"
- !!int "10"
- !!int "3"
"speed": "20 ft., climb 20 ft."
"skillsaves":
  "Stealth": !!int "5"
  "Perception": !!int "2"
"senses": "blindsight 60 ft., passive Perception 12"
"languages": ""
"cr": "3"
"traits":
- "desc": "The cave fisher can use its action to extend a sticky filament up to 60\
    \ feet, and the filament adheres to anything that touches it. A creature adhered\
    \ to the filament is [grappled](/3-Mechanics/CLI/rules/conditions.md#grappled)\
    \ by the cave fisher (escape DC 13), and ability checks made to escape this grapple\
    \ have disadvantage. The filament can be attacked (AC 15; 5 hit points; immunity\
    \ to poison and psychic damage), but a weapon that fails to sever it becomes stuck\
    \ to it, requiring an action and a successful DC 13 Strength check to pull free.\
    \ Destroying the filament causes no damage to the cave fisher, which can extrude\
    \ a replacement filament on its next turn"
  "name": "Adhesive Filament"
- "desc": "If the cave fisher drops to half its hit points or fewer, it gains vulnerability\
    \ to fire damage."
  "name": "Flammable Blood"
- "desc": "The cave fisher can climb difficult surfaces, including upside down on\
    \ ceilings, without needing to make an ability check."
  "name": "Spider Climb"
"actions":
- "desc": "The cave fisher makes two attacks with its claws."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+5 (+5) to hit, reach 5 ft., one target.\
    \ Hit: dice:2d6 + 3|text(10) (2d6 + 3) slashing damage."
  "name": "Claw"
- "desc": "One creature [grappled](/3-Mechanics/CLI/rules/conditions.md#grappled)\
    \ by the cave fisher's adhesive filament must make a DC 13 Strength saving throw,\
    \ provided that the target weighs 200 pounds or less. On a failure, the target\
    \ is pulled into an unoccupied space within 5 feet of the cave fisher, and the\
    \ cave fisher makes a claw attack against it as a bonus action. Reeling up the\
    \ target releases anyone else who was attached to the filament. Until the grapple\
    \ ends on the target, the cave fisher can't extrude another filament."
  "name": "Filament"
"source":
- "VGM"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/VGM/Cave%20Fisher.webp"
```
^statblock

```encounter-table
name: Cave Fisher
creatures:
 - 1: Cave Fisher
```

## Environment

underdark