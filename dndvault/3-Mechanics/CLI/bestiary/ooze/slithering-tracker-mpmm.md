---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/3
- monster/environment/underdark
- monster/environment/urban
- monster/size/medium
- monster/type/ooze
aliases: ["Slithering Tracker"]
NoteIcon: monster
BestiaryType: ooze
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 221, Volo's Guide to Monsters p. 191
---
# [Slithering Tracker](3-Mechanics\CLI\bestiary\ooze/slithering-tracker-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 221, Volo's Guide to Monsters p. 191*  

```statblock
"name": "Slithering Tracker (MPMM)"
"size": "Medium"
"type": "ooze"
"alignment": "Typically  Chaotic Evil"
"ac": !!int "14"
"hp": !!int "39"
"hit_dice": "6d8 + 12"
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
  "Survival": !!int "6"
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
- "desc": "If the slithering tracker is motionless at the start of combat, it has\
    \ advantage on its initiative roll. Moreover, if a creature hasn't observed the\
    \ slithering tracker move or act, that creature must succeed on a DC 18 Intelligence\
    \ ([Investigation](/3-Mechanics/CLI/rules/skills.md#Investigation)) check to discern\
    \ that the slithering tracker isn't a puddle."
  "name": "False Appearance"
- "desc": "The slithering tracker can enter an enemy's space and stop there. It can\
    \ also move through a space as narrow as 1 inch wide without squeezing."
  "name": "Liquid Form"
- "desc": "The slithering tracker can climb difficult surfaces, including upside down\
    \ on ceilings, without needing to make an ability check."
  "name": "Spider Climb"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+5 (+5) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d10 + 3|text(8) (1d10 + 3) bludgeoning damage."
  "name": "Slam"
- "desc": "One Large or smaller creature that the slithering tracker can see within\
    \ 5 feet of it must succeed on a DC 13 Dexterity saving throw or be [grappled](/3-Mechanics/CLI/rules/conditions.md#grappled)\
    \ (escape DC 13). Until this grapple ends, the target is [restrained](/3-Mechanics/CLI/rules/conditions.md#restrained)\
    \ and unable to breathe unless it can breathe water. In addition, the [grappled](/3-Mechanics/CLI/rules/conditions.md#grappled)\
    \ target takes dice:3d10|text(16) (3d10) necrotic damage at the start of each\
    \ of its turns. The slithering tracker can grapple only one target at a time.\n\
    \nWhile grappling the target, the slithering tracker takes only half any damage\
    \ dealt to it (rounded down), and the target takes the other half."
  "name": "Life Leech"
"bonus_actions":
- "desc": "If underwater, the slithering tracker takes the [Hide](/3-Mechanics/CLI/rules/actions.md#Hide)\
    \ action, and it makes the Dexterity ([Stealth](/3-Mechanics/CLI/rules/skills.md#Stealth))\
    \ check with advantage."
  "name": "Watery Stealth"
"source":
- "MPMM"
- "VGM"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Slithering%20Tracker.webp"
```
^statblock

```encounter-table
name: Slithering Tracker
creatures:
 - 1: Slithering Tracker
```

## Environment

underdark, urban