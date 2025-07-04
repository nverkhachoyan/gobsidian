---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mm
- monster/cr/2
- monster/environment/hill
- monster/environment/mountain
- monster/size/medium
- monster/type/monstrosity
aliases: ["Peryton"]
NoteIcon: monster
BestiaryType: monstrosity
SourceType: Bestiary
BookSource: Monster Manual p. 251
---
# [Peryton](3-Mechanics\CLI\bestiary\monstrosity/peryton.md)
*Source: Monster Manual p. 251*  

Although this monstrous carnivore feeds on any creature, it prefers humanoids, especially elves, half-elves, and humans. When it kills a humanoid, a peryton rips out its prey's heart and takes it back to its nest to be devoured.

The peryton is a bizarre creature that blends the body and wings of a bird of prey with the head of a stag. Its strangest feature is its shadow, which appears humanoid rather than reflecting the creature's physical form. Sages postulate that the first perytons were humans transformed by a hideous curse or magical experiment, but bards tell a different tale of a man whose infidelity caused his scorned wife to cut out the heart of her younger, more beautiful rival and consume it in a ritual intended to forever win her husband's heart. The ritual succeeded until the woman's villainy was exposed. She was hanged for her crime, but the lingering magic of her foul ritual caused the carrion birds that feasted on her corpse to transform into the first perytons.

## Unnatural Hunger

A peryton's reproductive cycle depends on the heart of a freshly killed humanoid. The organ must be consumed by a female peryton before she can reproduce. When a peryton consumes a heart, its shadow changes for a brief time to reflect its true monstrous form. When attacking a humanoid, a peryton is single-minded and relentless, fighting until it or its prey dies.

If a peryton is somehow driven away, it stalks lost prey from afar, attacking again when the opportunity arises.

## Bane of the Mountains

Perytons roost atop mountain ridges and lair in high caves. They prey on creatures living or wandering in the vales below, and travelers on lonely mountain roads learn to keep a wary eye on the sky. Because normal weapons are less effective against perytons, the folk of the mountains know to avoid confrontations with these monsters at all costs.

Established settlements are attractive to perytons as a renewable food source. As such, village councils and local nobles often hire adventurers to eliminate peryton nests.

```statblock
"name": "Peryton"
"size": "Medium"
"type": "monstrosity"
"alignment": "Chaotic Evil"
"ac": !!int "13"
"ac_class": "natural armor"
"hp": !!int "33"
"hit_dice": "6d8 + 6"
"stats":
- !!int "16"
- !!int "12"
- !!int "13"
- !!int "9"
- !!int "12"
- !!int "10"
"speed": "20 ft., fly 60 ft."
"skillsaves":
  "Perception": !!int "5"
"damage_resistances": "bludgeoning, piercing, slashing from nonmagical attacks"
"senses": "passive Perception 15"
"languages": "understands Common and Elvish but can't speak"
"cr": "2"
"traits":
- "desc": "If the peryton is flying and dives at least 30 feet straight toward a target\
    \ and then hits it with a melee weapon attack, the attack deals an extra dice:2d8|text(9)\
    \ (2d8) damage to the target."
  "name": "Dive Attack"
- "desc": "The peryton doesn't provoke an opportunity attack when it flies out of\
    \ an enemy's reach."
  "name": "Flyby"
- "desc": "The peryton has advantage on Wisdom ([Perception](/3-Mechanics/CLI/rules/skills.md#Perception))\
    \ checks that rely on sight or smell."
  "name": "Keen Sight and Smell"
"actions":
- "desc": "The peryton makes one gore attack and one talon attack."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+5 (+5) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d8 + 3|text(7) (1d8 + 3) piercing damage."
  "name": "Gore"
- "desc": "Melee Weapon Attack: dice: d20+5 (+5) to hit, reach 5 ft., one target.\
    \ Hit: dice:2d4 + 3|text(8) (2d4 + 3) piercing damage."
  "name": "Talons"
"source":
- "MM"
- "HotDQ"
- "PotA"
- "TftYP"
- "GoS"
- "IDRotF"
- "WBtW"
- "JttRC"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MM/Peryton.webp"
```
^statblock

```encounter-table
name: Peryton
creatures:
 - 1: Peryton
```

## Environment

mountain, hill