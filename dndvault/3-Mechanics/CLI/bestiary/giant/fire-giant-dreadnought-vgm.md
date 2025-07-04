---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/vgm
- monster/cr/14
- monster/environment/mountain
- monster/environment/underdark
- monster/size/huge
- monster/type/giant/fire-giant
aliases: ["Fire Giant Dreadnought"]
NoteIcon: monster
BestiaryType: giant (fire giant)
SourceType: Bestiary
BookSource: Volo's Guide to Monsters p. 147
---
# [Fire Giant Dreadnought](3-Mechanics\CLI\bestiary\giant/fire-giant-dreadnought-vgm.md)
*Source: Volo's Guide to Monsters p. 147*  

The ordning for fire giants emphasizes not just strength but also skill at forgecraft. The foundry is the heart of any fire giant community. It is temple, school, proving ground, and political hub rolled into one.

Those who have brawn but little brain are usually consigned to the lowliest of tasks such as working forge bellows or moving coal. However, there is one role at which the strongest among them can excel and gain rank: the dreadnought.

## Weapons of War

Dreadnoughts are massively powerful fire giants who wield two huge shields like plow blades. These shields bear spikes on their exterior and have hollow interiors into which the dreadnought pours hot coals at the first sign of danger. Armed with its two shields, the dreadnought can present a fiery wall to any attacker. When the dreadnought has finished, often all that is left of a foe is a smoking smear on the floor.

When not called on to fight, dreadnoughts maintain their strength by using their shields to shove huge quantities of coal, stone, or ore about the foundry. Occasionally, dreadnoughts are called on by their superiors to accompany a war or diplomatic delegation, The presence of the dreadnoughts presents a fierce face in either case.

```statblock
"name": "Fire Giant Dreadnought (VGM)"
"size": "Huge"
"type": "giant"
"subtype": "fire giant"
"alignment": "Lawful Evil"
"ac": !!int "21"
"ac_class": "[plate armor](/3-Mechanics/CLI/items/plate-armor.md), [shields](/3-Mechanics/CLI/items/shield.md)"
"hp": !!int "187"
"hit_dice": "15d12 + 90"
"stats":
- !!int "27"
- !!int "9"
- !!int "23"
- !!int "8"
- !!int "10"
- !!int "11"
"speed": "30 ft."
"saves":
  "Charisma": !!int "5"
  "Dexterity": !!int "4"
  "Constitution": !!int "11"
"skillsaves":
  "Athletics": !!int "13"
  "Perception": !!int "5"
"damage_immunities": "fire"
"senses": "passive Perception 15"
"languages": "Giant"
"cr": "14"
"traits":
- "desc": "The giant carries two shields, each of which is accounted for in the giant's\
    \ AC. The giant must stow or drop one of its shields to hurl rocks."
  "name": "Dual Shields"
"actions":
- "desc": "The giant makes two fireshield attacks."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+13 (+13) to hit, reach 5 ft., one\
    \ target. Hit: dice:4d6 + 8|text(22) (4d6 + 8) bludgeoning damage plus dice:2d6|text(7)\
    \ (2d6) fire damage plus dice:2d6|text(7) (2d6) piercing damage."
  "name": "Fireshield"
- "desc": "Ranged Weapon Attack: dice: d20+13 (+13) to hit, range 60/240 ft.,\
    \ one target. Hit: dice:4d10 + 8|text(30) (4d10 + 8) bludgeoning damage."
  "name": "Rock"
- "desc": "The giant moves up to 30 feet in a straight line and can move through the\
    \ space of any creature smaller than Huge. The first time it enters a creature's\
    \ space during this move, it makes a fireshield attack against that creature.\
    \ If the attack hits, the target must also succeed on a DC 21 Strength saving\
    \ throw or be pushed ahead of the giant for the rest of this move. If a creature\
    \ fails the save by 5 or more, it is also knocked [prone](/3-Mechanics/CLI/rules/conditions.md#prone)\
    \ and takes dice:3d6 + 8|text(18) (3d6 + 8) bludgeoning damage, or dice:6d6\
    \ + 8|text(29) (6d6 + 8) bludgeoning damage if it was already [prone](/3-Mechanics/CLI/rules/conditions.md#prone)."
  "name": "Shield Charge"
"source":
- "VGM"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/VGM/Fire%20Giant%20Dreadnought.webp"
```
^statblock

```encounter-table
name: Fire Giant Dreadnought
creatures:
 - 1: Fire Giant Dreadnought
```

## Environment

underdark, mountain