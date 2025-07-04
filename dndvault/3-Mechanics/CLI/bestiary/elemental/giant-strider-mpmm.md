---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/1
- monster/environment/hill
- monster/environment/mountain
- monster/environment/underdark
- monster/size/large
- monster/type/elemental
aliases: ["Giant Strider"]
NoteIcon: monster
BestiaryType: elemental
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 137, Volo's Guide to Monsters p. 143
---
# [Giant Strider](3-Mechanics\CLI\bestiary\elemental/giant-strider-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 137, Volo's Guide to Monsters p. 143*  

```statblock
"name": "Giant Strider (MPMM)"
"size": "Large"
"type": "elemental"
"alignment": "Unaligned"
"ac": !!int "14"
"ac_class": "natural armor"
"hp": !!int "22"
"hit_dice": "3d10 + 6"
"stats":
- !!int "18"
- !!int "13"
- !!int "14"
- !!int "4"
- !!int "12"
- !!int "6"
"speed": "50 ft."
"damage_immunities": "fire"
"senses": "passive Perception 11"
"languages": ""
"cr": "1"
"traits":
- "desc": "Whenever the giant strider is subjected to fire damage, it takes no damage\
    \ and regains a number of hit points equal to half the fire damage dealt."
  "name": "Fire Absorption"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+6 (+6) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d8 + 4|text(8) (1d8 + 4) piercing damage."
  "name": "Bite"
- "desc": "The giant strider hurls a gout of flame at a point it can see within 60\
    \ feet of it. Each creature in a 10-foot-radius sphere centered on that point\
    \ must make a DC 12 Dexterity saving throw, taking dice:4d6|text(14) (4d6)\
    \ fire damage on a failed save, or half as much damage on a successful one. The\
    \ fire spreads around corners, and it ignites flammable objects in that area that\
    \ aren't being worn or carried"
  "name": "Fire Burst (Recharge 5-6)"
"source":
- "MPMM"
- "VGM"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Giant%20Strider.webp"
```
^statblock

```encounter-table
name: Giant Strider
creatures:
 - 1: Giant Strider
```

## Environment

hill, mountain, underdark