---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/2
- monster/environment/forest
- monster/environment/swamp
- monster/environment/underdark
- monster/environment/urban
- monster/size/medium
- monster/type/fey
aliases: ["Darkling Elder"]
NoteIcon: monster
BestiaryType: fey
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 84, Volo's Guide to Monsters p. 134
---
# [Darkling Elder](3-Mechanics\CLI\bestiary\fey/darkling-elder-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 84, Volo's Guide to Monsters p. 134*  

```statblock
"name": "Darkling Elder (MPMM)"
"size": "Medium"
"type": "fey"
"alignment": "Typically  Chaotic Neutral"
"ac": !!int "15"
"ac_class": "[studded leather armor](/3-Mechanics/CLI/items/studded-leather-armor.md)"
"hp": !!int "27"
"hit_dice": "5d8 + 5"
"stats":
- !!int "13"
- !!int "17"
- !!int "12"
- !!int "10"
- !!int "14"
- !!int "13"
"speed": "30 ft."
"skillsaves":
  "Deception": !!int "3"
  "Stealth": !!int "7"
  "Perception": !!int "6"
  "Acrobatics": !!int "5"
"senses": "blindsight 30 ft., darkvision 120 ft., passive Perception 16"
"languages": "Elvish, Sylvan"
"cr": "2"
"traits":
- "desc": "The darkling elder casts [darkness](/3-Mechanics/CLI/spells/darkness.md),\
    \ requiring no spell components and using Wisdom as the spellcasting ability.\n"
  "name": "Darkness (Recharges after a Short or Long Rest)"
- "desc": "When the darkling elder dies, magical light flashes out from it in a 10-foot\
    \ radius as its body and possessions, other than metal or magic objects, burn\
    \ to ash. Any creature in that area must make a DC 11 Constitution saving throw.\
    \ On a failed save, the creature takes dice:2d6|text(7) (2d6) radiant damage\
    \ and is [blinded](/3-Mechanics/CLI/rules/conditions.md#blinded) until the end\
    \ of its next turn. On a successful save, the creature takes half as much damage\
    \ and isn't [blinded](/3-Mechanics/CLI/rules/conditions.md#blinded)."
  "name": "Death Burn"
"actions":
- "desc": "The darkling elder makes two Scimitar attacks."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+5 (+5) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d6 + 3|text(6) (1d6 + 3) slashing damage plus dice:2d6|text(7)\
    \ (2d6) necrotic damage."
  "name": "Scimitar"
"source":
- "MPMM"
- "VGM"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Darkling%20Elder.webp"
```
^statblock

```encounter-table
name: Darkling Elder
creatures:
 - 1: Darkling Elder
```

## Environment

forest, swamp, underdark, urban