---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/vrgr
- monster/cr/2
- monster/size/medium
- monster/type/beast
aliases: ["Swarm of Maggots"]
NoteIcon: monster
BestiaryType: beast
SourceType: Bestiary
BookSource: Van Richten's Guide to Ravenloft p. 247
---
# [Swarm of Maggots](3-Mechanics\CLI\bestiary\beast/swarm-of-maggots-vrgr.md)
*Source: Van Richten's Guide to Ravenloft p. 247*  

Base creatures are among the first to respond to sinister forces at work in a land. As nefarious powers grip an area, populations of maggots, scarabs, and similar scavenging insects explode and become aggressive predators. Roll on the Swarm Behavior table to see how such swarms might manifest.

**Swarm Behavior**

`dice: [](swarm-of-maggots-vrgr.md#^swarm-behavior)`

| dice: d4 | Behavior |
|----------|----------|
| 1 | Crawls on walls in a vaguely bipedal shape |
| 2 | Makes skittering noises that sound like whispered chanting |
| 3 | Skeletal visages, giant eyes, or the faces of nearby creatures appear in relief amid its mass |
| 4 | Occupies and animates a corpse or other debris as if it were alive |
^swarm-behavior

```statblock
"name": "Swarm of Maggots (VRGR)"
"size": "Medium"
"type": "beast"
"alignment": "Unaligned"
"ac": !!int "11"
"hp": !!int "22"
"hit_dice": "5d8"
"stats":
- !!int "3"
- !!int "12"
- !!int "10"
- !!int "1"
- !!int "7"
- !!int "1"
"speed": "20 ft., swim 20 ft."
"damage_resistances": "bludgeoning, piercing, slashing"
"condition_immunities": "[charmed](/3-Mechanics/CLI/rules/conditions.md#charmed),\
  \ [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened), [grappled](/3-Mechanics/CLI/rules/conditions.md#grappled),\
  \ [paralyzed](/3-Mechanics/CLI/rules/conditions.md#paralyzed), [petrified](/3-Mechanics/CLI/rules/conditions.md#petrified),\
  \ [prone](/3-Mechanics/CLI/rules/conditions.md#prone), [restrained](/3-Mechanics/CLI/rules/conditions.md#restrained),\
  \ [stunned](/3-Mechanics/CLI/rules/conditions.md#stunned)"
"senses": "blindsight 10 ft., passive Perception 8"
"languages": ""
"cr": "2"
"traits":
- "desc": "The swarm can occupy another creature's space and vice versa, and the swarm\
    \ can move through any opening large enough for a Tiny maggot. The swarm can't\
    \ regain hit points or gain temporary hit points."
  "name": "Swarm"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+3 (+3) to hit, reach 0 ft., one target\
    \ in the swarm's space. Hit: dice:4d4|text(10) (4d4) piercing damage, or\
    \ dice:2d4|text(5) (2d4) piercing damage if the swarm has half of its hit\
    \ points or fewer. A creature damaged by the swarm must succeed on a DC 12 Constitution\
    \ saving throw or contract a disease.\n\nEach time the diseased creature finishes\
    \ a long rest, roll a dice: d6|avg|noform (d6) to determine the disease's\
    \ effect:\n\n- 1-2. The creature is [blinded](/3-Mechanics/CLI/rules/conditions.md#blinded)\
    \ until it finishes a long rest.  \n- 3-4. The creature's hit point maximum\
    \ decreases by dice:2d4|text(5) (2d4), and the reduction can't be removed\
    \ until the disease ends. The creature dies if its hit point maximum drops to\
    \ 0.  \n- 5-6. The creature has disadvantage on ability checks and attack\
    \ rolls until it finishes its next long rest.  \n\n    The disease lasts until\
    \ it's removed by magic or until the creature rolls the same random effect for\
    \ the disease two long rests in a row.  "
  "name": "Infestation"
"source":
- "VRGR"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/VRGR/Swarm%20of%20Maggots.webp"
```
^statblock

```encounter-table
name: Swarm of Maggots
creatures:
 - 1: Swarm of Maggots
```