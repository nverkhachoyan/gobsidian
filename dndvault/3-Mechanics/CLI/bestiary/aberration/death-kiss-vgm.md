---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/vgm
- monster/cr/10
- monster/environment/underdark
- monster/size/large
- monster/type/aberration
aliases: ["Death Kiss"]
NoteIcon: monster
BestiaryType: aberration
SourceType: Bestiary
BookSource: Volo's Guide to Monsters p. 124
---
# [Death Kiss](3-Mechanics\CLI\bestiary\aberration/death-kiss-vgm.md)
*Source: Volo's Guide to Monsters p. 124*  

A death kiss is a lesser beholder that might come into being when a beholder has a vivid nightmare about losing blood. Instead of magical eye rays, it has ten long tentacles, each ending in a mouth full of teeth. In coloration and shape it is similar to the beholder that dreamed it into existence, but its hue is more muted.

## Blood Drinker

A death kiss survives solely on ingested blood, which it uses to generate electrical energy inside its body. Paranoid about dying from starvation, it obsessively drains even little creatures such as rats in an effort to stave off this fate for as long as possible. After it drains its prey, it abandons the corpse to scavengers. A death kiss prefers to hunt alone. If it meets another death kiss, it might fight, flee, or team up, depending on its health and pride. When underground, it uses its tentacles as feelers, prodding and examining the environment in all directions. Above ground, it usually keeps its tentacles retracted when on the hunt, then lashes out and extends them to their full length to catch opponents off guard.

## False Tyrant

In poor lighting and with its tentacles extended, a death kiss can be mistaken for a true beholder. It might purposely present itself as a beholder to an ignorant creature, but this behavior is rare, since it usually is focused on hunting and lacks the self-importance and paranoia of a true beholder. It can speak through any of its tentacle-throats, and its voice sounds nasal and high-pitched. A true beholder has little to fear from a death kiss, since it can easily kill or subdue the death kiss long before the death kiss gets into melee range. Thus, out of self-preservation, a death kiss usually submits to the rule of a beholder that it encounters, though it might attempt to escape as soon as its master is preoccupied.

## Simple Tactics

A death kiss lacks the combat finesse and intelligence of a beholder. It might attempt an unusual maneuver to control its prey (such as flying up while grappling), but in most cases, it attaches one or more of its tentacles to a creature and drains blood until its prey collapses. If it is in a superior position and its opponent poses no threat, it might toy with its food, slowly squeezing and draining the life out of a creature.

```statblock
"name": "Death Kiss (VGM)"
"size": "Large"
"type": "aberration"
"alignment": "Neutral Evil"
"ac": !!int "16"
"ac_class": "natural armor"
"hp": !!int "161"
"hit_dice": "17d10 + 68"
"stats":
- !!int "18"
- !!int "14"
- !!int "18"
- !!int "10"
- !!int "12"
- !!int "10"
"speed": "0 ft., fly 30 ft. (hover)"
"saves":
  "Wisdom": !!int "5"
  "Constitution": !!int "8"
"skillsaves":
  "Perception": !!int "5"
"damage_immunities": "lightning"
"condition_immunities": "[prone](/3-Mechanics/CLI/rules/conditions.md#prone)"
"senses": "darkvision 120 ft., passive Perception 15"
"languages": "Deep Speech, Undercommon"
"cr": "10"
"traits":
- "desc": "A creature within 5 feet of the death kiss takes dice:1d10|text(5) (1d10)\
    \ lightning damage whenever it hits the death kiss with a melee attack that deals\
    \ piercing or slashing damage."
  "name": "Lightning Blood"
"actions":
- "desc": "The death kiss makes three tentacle attacks. Up to three of these attacks\
    \ can be replaced by Blood Drain, one replacement per tentacle grappling a creature"
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+8 (+8) to hit, reach 20 ft., one target.\
    \ Hit: dice:3d6 + 4|text(14) (3d6 + 4) piercing damage, and the target is\
    \ [grappled](/3-Mechanics/CLI/rules/conditions.md#grappled) (escape DC 14) if\
    \ it is a Huge or smaller creature. Until this grapple ends, the target is [restrained](/3-Mechanics/CLI/rules/conditions.md#restrained),\
    \ and the death kiss can't use the same tentacle on another target. The death\
    \ kiss has ten tentacles."
  "name": "Tentacle"
- "desc": "One creature [grappled](/3-Mechanics/CLI/rules/conditions.md#grappled)\
    \ by a tentacle of the death kiss must make a DC 16 Constitution saving throw.\
    \ On a failed save, the target takes dice:4d10|text(22) (4d10) lightning damage,\
    \ and the death kiss regains half as many hit points."
  "name": "Blood Drain"
"source":
- "VGM"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/VGM/Death%20Kiss.webp"
```
^statblock

```encounter-table
name: Death Kiss
creatures:
 - 1: Death Kiss
```

## Environment

underdark