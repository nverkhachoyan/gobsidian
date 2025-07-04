---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/6
- monster/environment/arctic
- monster/environment/hill
- monster/environment/mountain
- monster/environment/underdark
- monster/environment/urban
- monster/size/medium
- monster/type/humanoid
aliases: ["Warlock of the Great Old One"]
NoteIcon: monster
BestiaryType: humanoid
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 256, Volo's Guide to Monsters p. 220
---
# [Warlock of the Great Old One](3-Mechanics\CLI\bestiary\humanoid/warlock-of-the-great-old-one-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 256, Volo's Guide to Monsters p. 220*  

```statblock
"name": "Warlock of the Great Old One (MPMM)"
"size": "Medium"
"type": "humanoid"
"alignment": "Any alignment"
"ac": !!int "13"
"ac_class": "16 with [mage armor](/3-Mechanics/CLI/spells/mage-armor.md)"
"hp": !!int "91"
"hit_dice": "14d8 + 28"
"stats":
- !!int "9"
- !!int "16"
- !!int "15"
- !!int "12"
- !!int "12"
- !!int "18"
"speed": "30 ft."
"saves":
  "Charisma": !!int "7"
  "Wisdom": !!int "4"
"skillsaves":
  "History": !!int "4"
  "Arcana": !!int "4"
"damage_resistances": "psychic"
"senses": "darkvision 60 ft., passive Perception 11"
"languages": "any two languages, telepathy 30 ft."
"cr": "6"
"traits":
- "desc": "The warlock casts one of the following spells, using Charisma as the spellcasting\
    \ ability (spell save DC 15): \n\nAt will: [detect magic](/3-Mechanics/CLI/spells/detect-magic.md),\
    \ [guidance](/3-Mechanics/CLI/spells/guidance.md), [levitate](/3-Mechanics/CLI/spells/levitate.md),\
    \ [mage armor](/3-Mechanics/CLI/spells/mage-armor.md) (self only), [mage hand](/3-Mechanics/CLI/spells/mage-hand.md),\
    \ [minor illusion](/3-Mechanics/CLI/spells/minor-illusion.md), [prestidigitation](/3-Mechanics/CLI/spells/prestidigitation.md)\n\
    \n1/day each: [arcane gate](/3-Mechanics/CLI/spells/arcane-gate.md), [detect\
    \ thoughts](/3-Mechanics/CLI/spells/detect-thoughts.md), [true seeing](/3-Mechanics/CLI/spells/true-seeing.md)"
  "name": "Spellcasting"
- "desc": "At the start of each of the warlock's turns, each creature of its choice\
    \ within 10 feet of it must succeed on a DC 15 Wisdom saving throw or take dice:3d6|text(10)\
    \ (3d6) psychic damage, provided that the warlock isn't [incapacitated](/3-Mechanics/CLI/rules/conditions.md#incapacitated)."
  "name": "Whispering Aura"
"actions":
- "desc": "The warlock makes two Dagger attacks."
  "name": "Multiattack"
- "desc": "Melee or Ranged Weapon Attack: dice: d20+6 (+6) to hit, reach 5 ft.\
    \ or range 20/60 ft., one target. Hit: dice:1d4 + 3|text(5) (1d4 + 3) piercing\
    \ damage plus dice:3d6|text(10) (3d6) psychic damage."
  "name": "Dagger"
- "desc": "The warlock opens a momentary extraplanar rift within 60 feet of it. The\
    \ rift is a scream-filled, 20-foot cube. Each creature in that area must make\
    \ a DC 15 Wisdom saving throw. On a failed save, a creature takes dice:2d8|text(9)\
    \ (2d8) psychic damage and is [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened)\
    \ of the warlock until the start of the warlock's next turn. On a successful save,\
    \ a creature takes half as much damage and isn't [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened)."
  "name": "Howling Void"
"source":
- "MPMM"
- "VGM"
- "BMT"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Warlock%20of%20the%20Great%20Old%20One.webp"
```
^statblock

```encounter-table
name: Warlock of the Great Old One
creatures:
 - 1: Warlock of the Great Old One
```

## Environment

arctic, hill, mountain, underdark, urban