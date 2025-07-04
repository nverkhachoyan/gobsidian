---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mm
- monster/cr/1-2
- monster/environment/underdark
- monster/size/medium
- monster/type/ooze
aliases: ["Gray Ooze"]
NoteIcon: monster
BestiaryType: ooze
SourceType: Bestiary
BookSource: Monster Manual p. 243. Available in the SRD.
---
# [Gray Ooze](3-Mechanics\CLI\bestiary\ooze/gray-ooze.md)
*Source: Monster Manual p. 243. Available in the SRD.*  

A gray ooze is stone turned to liquid by chaos. When it moves, it slithers like a liquid snake, rising to strike.

## Oozes

Oozes thrive in the dark, shunning areas of bright light and extreme temperatures. They flow through the damp underground, feeding on any creature or object that can be dissolved, slinking along the ground, dripping from walls and ceilings, spreading across the edges of underground pools, and squeezing through cracks.

The first warning an adventurer receives of an ooze's presence is often the searing pain of its acidic touch. Oozes are drawn to movement and warmth. Organic material nourishes them, and when prey is scarce they feed on grime, fungus, and offal. Veteran explorers know that an immaculately clean passageway is a likely sign that an ooze lairs nearby.

### Slow Death

An ooze kills its prey slowly. Some varieties, such as black puddings and gelatinous cubes, engulf creatures to prevent escape. The only upside of this torturous death is that a victim's comrades can come to the rescue before it is too late.

Since not every ooze digests every type of substance, some have coins, metal gear, bones, and other debris suspended within their quivering bodies. A slain ooze can be a rich source of treasure for its killers.

Whether this is true or not, the Faceless Lord is one of the few beings that can control oozes and imbue them with a modicum of intelligence. Most of the time, oozes have no sense of tactics or self-preservation. They are direct and predictable, attacking and eating without cunning. Under the control of Juiblex, they exhibit glimmers of sentience and malevolent intent.

### Unwitting Servants

Although an ooze lacks the intelligence to ally itself with other creatures, others that understand an ooze's need to feed might lure it into a location where it can be of use to them. Clever monsters keep oozes around to defend passageways or consume refuse. Likewise, an ooze can be enticed into a pit trap, where its captors feed it often enough to prevent it from coming after them. Crafty creatures place torches and flaming braziers in strategic areas to dissuade an ooze from leaving a particular tunnel or room.

### Spawn of Juiblex

According to the Demonomicon of Iggwilv and other sources, oozes are scattered fragments or offspring of the demon lord Juiblex.

### Ooze Nature

An ooze doesn't require sleep.

```statblock
"name": "Gray Ooze"
"size": "Medium"
"type": "ooze"
"alignment": "Unaligned"
"ac": !!int "8"
"hp": !!int "22"
"hit_dice": "3d8 + 9"
"stats":
- !!int "12"
- !!int "6"
- !!int "16"
- !!int "1"
- !!int "6"
- !!int "2"
"speed": "10 ft., climb 10 ft."
"skillsaves":
  "Stealth": !!int "2"
"damage_resistances": "acid, cold, fire"
"condition_immunities": "[blinded](/3-Mechanics/CLI/rules/conditions.md#blinded),\
  \ [charmed](/3-Mechanics/CLI/rules/conditions.md#charmed), [deafened](/3-Mechanics/CLI/rules/conditions.md#deafened),\
  \ [exhaustion](/3-Mechanics/CLI/rules/conditions.md#exhaustion), [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened),\
  \ [prone](/3-Mechanics/CLI/rules/conditions.md#prone)"
"senses": "blindsight 60 ft. (blind beyond this radius), passive Perception 8"
"languages": ""
"cr": "1/2"
"traits":
- "desc": "The ooze can move through a space as narrow as 1 inch wide without squeezing."
  "name": "Amorphous"
- "desc": "Any nonmagical weapon made of metal that hits the ooze corrodes. After\
    \ dealing damage, the weapon takes a permanent and cumulative −1 penalty to damage\
    \ rolls. If its penalty drops to −5, the weapon is destroyed. Nonmagical ammunition\
    \ made of metal that hits the ooze is destroyed after dealing damage.\n\nThe ooze\
    \ can eat through 2-inch-thick, nonmagical metal in 1 round."
  "name": "Corrode Metal"
- "desc": "While the ooze remains motionless, it is indistinguishable from an oily\
    \ pool or wet rock."
  "name": "False Appearance"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+3 (+3) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d6 + 1|text(4) (1d6 + 1) bludgeoning damage plus dice:2d6|text(7)\
    \ (2d6) acid damage, and if the target is wearing nonmagical metal armor, its\
    \ armor is partly corroded and takes a permanent and cumulative −1 penalty to\
    \ the AC it offers. The armor is destroyed if the penalty reduces its AC to 10."
  "name": "Pseudopod"
"source":
- "MM"
- "CoS"
- "HotDQ"
- "TftYP"
- "WDH"
- "GoS"
- "KftGV"
- "PaBTSO"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MM/Gray%20Ooze.webp"
```
^statblock

```encounter-table
name: Gray Ooze
creatures:
 - 1: Gray Ooze
```

## Environment

underdark