---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/1-2
- monster/environment/forest
- monster/environment/swamp
- monster/environment/underdark
- monster/environment/urban
- monster/size/small
- monster/type/fey
aliases: ["Darkling"]
NoteIcon: monster
BestiaryType: fey
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 84, Volo's Guide to Monsters p. 134
---
# [Darkling](3-Mechanics\CLI\bestiary\fey/darkling-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 84, Volo's Guide to Monsters p. 134*  

```statblock
"name": "Darkling (MPMM)"
"size": "Small"
"type": "fey"
"alignment": "Typically  Chaotic Neutral"
"ac": !!int "14"
"ac_class": "[leather armor](/3-Mechanics/CLI/items/leather-armor.md)"
"hp": !!int "13"
"hit_dice": "3d6 + 3"
"stats":
- !!int "9"
- !!int "16"
- !!int "12"
- !!int "10"
- !!int "12"
- !!int "10"
"speed": "30 ft."
"skillsaves":
  "Deception": !!int "2"
  "Stealth": !!int "7"
  "Perception": !!int "5"
  "Acrobatics": !!int "5"
"senses": "blindsight 30 ft., darkvision 120 ft., passive Perception 15"
"languages": "Elvish, Sylvan"
"cr": "1/2"
"traits":
- "desc": "When the darkling dies, nonmagical light flashes out from it in a 10-foot\
    \ radius as its body and possessions, other than metal or magic objects, burn\
    \ to ash. Any creature in that area must succeed on a DC 10 Constitution saving\
    \ throw or be [blinded](/3-Mechanics/CLI/rules/conditions.md#blinded) until the\
    \ end of its next turn."
  "name": "Death Flash"
- "desc": "While in bright light, the darkling has disadvantage on attack rolls, as\
    \ well as on Wisdom ([Perception](/3-Mechanics/CLI/rules/skills.md#Perception))\
    \ checks that rely on sight."
  "name": "Light Sensitivity"
"actions":
- "desc": "Melee or Ranged Weapon Attack: dice: d20+5 (+5) to hit, reach 5 ft.\
    \ or range 20/60 ft., one target. Hit: dice:1d4 + 3|text(5) (1d4 + 3) piercing\
    \ damage plus dice:2d6|text(7) (2d6) necrotic damage."
  "name": "Dagger"
"source":
- "MPMM"
- "VGM"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Darkling.webp"
```
^statblock

```encounter-table
name: Darkling
creatures:
 - 1: Darkling
```

## Environment

forest, swamp, underdark, urban