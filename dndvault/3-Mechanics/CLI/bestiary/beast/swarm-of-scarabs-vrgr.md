---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/vrgr
- monster/cr/3
- monster/size/medium
- monster/type/beast
aliases: ["Swarm of Scarabs"]
NoteIcon: monster
BestiaryType: beast
SourceType: Bestiary
BookSource: Van Richten's Guide to Ravenloft p. 247
---
# [Swarm of Scarabs](3-Mechanics\CLI\bestiary\beast/swarm-of-scarabs-vrgr.md)
*Source: Van Richten's Guide to Ravenloft p. 247*  

Base creatures are among the first to respond to sinister forces at work in a land. As nefarious powers grip an area, populations of maggots, scarabs, and similar scavenging insects explode and become aggressive predators. Roll on the Swarm Behavior table to see how such swarms might manifest.

**Swarm Behavior**

`dice: [](swarm-of-scarabs-vrgr.md#^swarm-behavior)`

| dice: d4 | Behavior |
|----------|----------|
| 1 | Crawls on walls in a vaguely bipedal shape |
| 2 | Makes skittering noises that sound like whispered chanting |
| 3 | Skeletal visages, giant eyes, or the faces of nearby creatures appear in relief amid its mass |
| 4 | Occupies and animates a corpse or other debris as if it were alive |
^swarm-behavior

```statblock
"name": "Swarm of Scarabs (VRGR)"
"size": "Medium"
"type": "beast"
"alignment": "Unaligned"
"ac": !!int "13"
"ac_class": "natural armor"
"hp": !!int "27"
"hit_dice": "5d8 + 5"
"stats":
- !!int "3"
- !!int "14"
- !!int "13"
- !!int "1"
- !!int "12"
- !!int "1"
"speed": "30 ft., burrow 30 ft., climb 30 ft."
"damage_resistances": "bludgeoning, piercing, slashing"
"condition_immunities": "[charmed](/3-Mechanics/CLI/rules/conditions.md#charmed),\
  \ [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened), [grappled](/3-Mechanics/CLI/rules/conditions.md#grappled),\
  \ [paralyzed](/3-Mechanics/CLI/rules/conditions.md#paralyzed), [petrified](/3-Mechanics/CLI/rules/conditions.md#petrified),\
  \ [prone](/3-Mechanics/CLI/rules/conditions.md#prone), [restrained](/3-Mechanics/CLI/rules/conditions.md#restrained),\
  \ [stunned](/3-Mechanics/CLI/rules/conditions.md#stunned)"
"senses": "tremorsense 60 ft., passive Perception 11"
"languages": ""
"cr": "3"
"traits":
- "desc": "The swarm can occupy another creature's space and vice versa, and the swarm\
    \ can move through any opening large enough for a Tiny scarab. The swarm can't\
    \ regain hit points or gain temporary hit points."
  "name": "Swarm"
- "desc": "If the swarm starts its turn in the same space as a dead creature that\
    \ is Large or smaller, the corpse is destroyed, leaving behind only equipment\
    \ and bones (or exoskeleton)."
  "name": "Skeletonize"
- "desc": "The swarm can climb difficult surfaces, including upside down on ceilings,\
    \ without needing to make an ability check."
  "name": "Spider Climb"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+4 (+4) to hit, reach 0 ft., one target\
    \ in the swarm's space. Hit: dice:4d6|text(14) (4d6) piercing damage, or\
    \ dice:2d6|text(7) (2d6) piercing damage if the swarm has half of its hit\
    \ points or fewer. If the target is a creature, scarabs burrow into its body,\
    \ and the creature takes dice:1d6|text(3) (1d6) piercing damage at the start\
    \ of each of its turns. Any creature can use an action to kill or remove the scarabs\
    \ with fire or a weapon that deals piercing damage, causing 1 damage of the appropriate\
    \ type to the target. A creature reduced to 0 hit points by the swarm's piercing\
    \ damage dies."
  "name": "Ravenous Bites"
"source":
- "VRGR"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/VRGR/Swarm%20of%20Scarabs.webp"
```
^statblock

```encounter-table
name: Swarm of Scarabs
creatures:
 - 1: Swarm of Scarabs
```