---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/vgm
- monster/cr/3
- monster/environment/underdark
- monster/environment/urban
- monster/size/medium
- monster/type/ooze
aliases: ["Slithering Tracker"]
NoteIcon: monster
BestiaryType: ooze
SourceType: Bestiary
BookSource: Volo's Guide to Monsters p. 191
---
# [Slithering Tracker](3-Mechanics\CLI\bestiary\ooze/slithering-tracker-vgm.md)
*Source: Volo's Guide to Monsters p. 191*  

The quest for revenge sometimes leads one to undergo a ritual whereby they transform into a body of semiliquid sentience known as a slithering tracker. Innocuous and insidious at the same time, a tracker flows into places where a normal creature can't go and brings its own brand of watery death down upon its quarry.

## Vengeance at Any Cost

The ritual for creating a slithering tracker is known to hags, liches, and priests who worship gods of vengeance. It can only be performed on a willing creature that hungers for revenge. The ritual sucks all the moisture from the person's body, killing it. Yet the mind lives on in the puddle of liquid that issues forth from the remains, and so too does the subject's insatiable need for retribution.

## Stealthy Assassins

A slithering tracker tastes the ground it courses over, seeking any trace of its prey. To kill, a slithering tracker rises up and enshrouds a creature, attempting to drown the prey while also draining it of blood. A slithering tracker that has killed in this fashion becomes much easier to locate for a time, since its liquid form becomes tinged with blood and its body leaves a visible trail of the stuff behind it.

## Descent into Madness

Achieving revenge against its target doesn't end a slithering tracker's existence, nor its hunger for blood. Some slithering trackers remain aware of their purpose and extend their quest for vengeance to others, such as anyone who supported or befriended the original target. Most of the time, though, a tracker's mind can't cope with being trapped in liquid form, unable to communicate, and driven by the desire for blood: after a tracker fulfills its duty, insanity takes over the creature, and it attacks indiscriminately until it is destroyed.

```statblock
"name": "Slithering Tracker (VGM)"
"size": "Medium"
"type": "ooze"
"alignment": "Chaotic Evil"
"ac": !!int "14"
"hp": !!int "32"
"hit_dice": "5d8 + 10"
"stats":
- !!int "16"
- !!int "19"
- !!int "15"
- !!int "10"
- !!int "14"
- !!int "11"
"speed": "30 ft., climb 30 ft., swim 30 ft."
"skillsaves":
  "Stealth": !!int "8"
"damage_vulnerabilities": "cold, fire"
"damage_resistances": "bludgeoning, piercing, slashing from nonmagical attacks"
"condition_immunities": "[blinded](/3-Mechanics/CLI/rules/conditions.md#blinded),\
  \ [deafened](/3-Mechanics/CLI/rules/conditions.md#deafened), [exhaustion](/3-Mechanics/CLI/rules/conditions.md#exhaustion),\
  \ [grappled](/3-Mechanics/CLI/rules/conditions.md#grappled), [paralyzed](/3-Mechanics/CLI/rules/conditions.md#paralyzed),\
  \ [petrified](/3-Mechanics/CLI/rules/conditions.md#petrified), [prone](/3-Mechanics/CLI/rules/conditions.md#prone),\
  \ [restrained](/3-Mechanics/CLI/rules/conditions.md#restrained), [unconscious](/3-Mechanics/CLI/rules/conditions.md#unconscious)"
"senses": "blindsight 120 ft., passive Perception 12"
"languages": "understands languages it knew in its previous form but can't speak"
"cr": "3"
"traits":
- "desc": "In the first round of a combat, the slithering tracker has advantage on\
    \ attack rolls against any creature it [surprised](/3-Mechanics/CLI/rules/conditions.md#surprised)."
  "name": "Ambusher"
- "desc": "While grappling a creature, the slithering tracker takes only haIf the\
    \ damage dealt to it, and the creature it is grappling takes the other half."
  "name": "Damage Transfer"
- "desc": "While the slithering tracker remains motionless, it is indistinguishable\
    \ from a puddle, unless an observer succeeds on a DC 18 Intelligence ([Investigation](/3-Mechanics/CLI/rules/skills.md#Investigation))\
    \ check."
  "name": "False Appearance"
- "desc": "The slithering tracker has advantage on Wisdom checks to track prey."
  "name": "Keen Tracker"
- "desc": "The slithering tracker can enter an enemy's space and stop there. It can\
    \ also move through a space as narrow as 1 inch wide without squeezing."
  "name": "Liquid Form"
- "desc": "The slithering tracker can climb difficult surfaces, including upside down\
    \ on ceilings, without needing to make an ability check."
  "name": "Spider Climb"
- "desc": "While underwater, the slithering tracker has advantage on Dexterity ([Stealth](/3-Mechanics/CLI/rules/skills.md#Stealth))\
    \ checks made to hide, and it can take the Hide action as a bonus action."
  "name": "Watery Stealth"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+5 (+5) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d10 + 3|text(8) (1d10 + 3) bludgeoning damage."
  "name": "Slam"
- "desc": "One Large or smaller creature that the slithering tracker can see within\
    \ 5 feet of it must succeed on a DC 13 Dexterity saving throw or be [grappled](/3-Mechanics/CLI/rules/conditions.md#grappled)\
    \ (escape DC 13). Until this grapple ends, the target is [restrained](/3-Mechanics/CLI/rules/conditions.md#restrained)\
    \ and unable to breathe unless it can breathe water. In addition, the [grappled](/3-Mechanics/CLI/rules/conditions.md#grappled)\
    \ target takes dice:3d10|text(16) (3d10) necrotic damage at the start of each\
    \ of its turns. The slithering tracker can grapple only one target at a time."
  "name": "Life Leech"
"source":
- "VGM"
- "ToA"
- "IMR"
- "MOT"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/VGM/Slithering%20Tracker.webp"
```
^statblock

```encounter-table
name: Slithering Tracker
creatures:
 - 1: Slithering Tracker
```

## Environment

underdark, urban