---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mm
- monster/cr/5
- monster/environment/grassland
- monster/environment/hill
- monster/environment/mountain
- monster/size/large
- monster/type/monstrosity
aliases: ["Bulette"]
NoteIcon: monster
BestiaryType: monstrosity
SourceType: Bestiary
BookSource: Monster Manual p. 34. Available in the SRD.
---
# [Bulette](3-Mechanics\CLI\bestiary\monstrosity/bulette.md)
*Source: Monster Manual p. 34. Available in the SRD.*  

A bulette is a massive predator that terrorizes any lands it inhabits. Also called a "land shark," it lives only to feed. Irascible and rapacious, bulettes fear no other creature, and they attack with no regard for superior numbers or strength.

## Underground Hunters

Bulettes use their powerful claws to tunnel through the earth when they hunt. Heedless of obstruction, they uproot trees, cause landslides in loose slopes, and leave sinkholes behind them. When vibrations in the soil and rock alert a bulette to movement, it shoots to the surface, its jaws spread wide as it attacks.

## Wandering Monster

A bulette ranges across temperate lands, feeding on any animals and humanoids it comes across. These creatures dislike dwarf and elf flesh, although they often kill them before realizing what they are. A bulette loves halfling meat the most, and it is never happier than when chasing plump halflings across an open field.

A bulette has no lair, but roams a hunting territory up to thirty miles wide. Its sole criterion for territory is availability of food, and when it has eaten everything in an area, a bulette moves on. These creatures often home in on humanoid settlements, terrorizing them until their panicked residents have fled, or until the bulette is slain.

All creatures shun bulettes, which treat anything that moves as food-even other predators and bulettes. Bulettes come together only to mate, resulting in a bloody act of claws and teeth that usually ends with the male's death and consumption.

## Arcane Creation

Some sages believe the bulette is the result of a mad wizard's experiments at crossbreeding snapping turtles and armadillos, within fusions of demon ichor. Bulettes have been thought to be extinct at different times, but after years without a sighting, the creatures inevitably reappear. Because their young are almost never seen, some sages suspect that bulettes maintain secret nesting grounds from which adults strike out into the world.

```statblock
"name": "Bulette"
"size": "Large"
"type": "monstrosity"
"alignment": "Unaligned"
"ac": !!int "17"
"ac_class": "natural armor"
"hp": !!int "94"
"hit_dice": "9d10 + 45"
"stats":
- !!int "19"
- !!int "11"
- !!int "21"
- !!int "2"
- !!int "10"
- !!int "5"
"speed": "40 ft., burrow 40 ft."
"skillsaves":
  "Perception": !!int "6"
"senses": "darkvision 60 ft., tremorsense 60 ft., passive Perception 16"
"languages": ""
"cr": "5"
"traits":
- "desc": "The bulette's long jump is up to 30 feet and its high jump is up to 15\
    \ feet, with or without a running start."
  "name": "Standing Leap"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+7 (+7) to hit, reach 5 ft., one target.\
    \ Hit: dice:4d12 + 4|text(30) (4d12 + 4) piercing damage."
  "name": "Bite"
- "desc": "If the bulette jumps at least 15 feet as part of its movement, it can then\
    \ use this action to land on its feet in a space that contains one or more other\
    \ creatures. Each of those creatures must succeed on a DC 16 Strength or Dexterity\
    \ saving throw (target's choice) or be knocked [prone](/3-Mechanics/CLI/rules/conditions.md#prone)\
    \ and take dice:3d6 + 4|text(14) (3d6 + 4) bludgeoning damage plus dice:3d6\
    \ + 4|text(14) (3d6 + 4) slashing damage. On a successful save, the creature\
    \ takes only half the damage, isn't knocked [prone](/3-Mechanics/CLI/rules/conditions.md#prone),\
    \ and is pushed 5 feet out of the bulette's space into an unoccupied space of\
    \ the creature's choice. If no unoccupied space is within range, the creature\
    \ instead falls [prone](/3-Mechanics/CLI/rules/conditions.md#prone) in the bulette's\
    \ space."
  "name": "Deadly Leap"
"source":
- "MM"
- "PotA"
- "SKT"
- "IDRotF"
- "JttRC"
- "LoX"
- "DSotDQ"
- "QftIS"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MM/Bulette.webp"
```
^statblock

```encounter-table
name: Bulette
creatures:
 - 1: Bulette
```

## Environment

mountain, grassland, hill