---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/5
- monster/environment/swamp
- monster/size/large
- monster/type/monstrosity
aliases: ["Catoblepas"]
NoteIcon: monster
BestiaryType: monstrosity
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 70, Volo's Guide to Monsters p. 129
---
# [Catoblepas](3-Mechanics\CLI\bestiary\monstrosity/catoblepas-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 70, Volo's Guide to Monsters p. 129*  

```statblock
"name": "Catoblepas (MPMM)"
"size": "Large"
"type": "monstrosity"
"alignment": "Unaligned"
"ac": !!int "14"
"ac_class": "natural armor"
"hp": !!int "84"
"hit_dice": "8d10 + 40"
"stats":
- !!int "19"
- !!int "12"
- !!int "21"
- !!int "3"
- !!int "14"
- !!int "8"
"speed": "30 ft."
"skillsaves":
  "Perception": !!int "5"
"senses": "darkvision 60 ft., passive Perception 15"
"languages": ""
"cr": "5"
"traits":
- "desc": "Any creature other than a catoblepas that starts its turn within 10 feet\
    \ of the catoblepas must succeed on a DC 16 Constitution saving throw or be [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)\
    \ until the start of the creature's next turn. On a successful saving throw, the\
    \ creature is immune to the Stench of any catoblepas for 1 hour."
  "name": "Stench"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+7 (+7) to hit, reach 10 ft., one target.\
    \ Hit: dice:5d6 + 4|text(21) (5d6 + 4) bludgeoning damage, and the target\
    \ must succeed on a DC 16 Constitution saving throw or be [stunned](/3-Mechanics/CLI/rules/conditions.md#stunned)\
    \ until the start of the catoblepas's next turn."
  "name": "Tail"
- "desc": "The catoblepas targets one creature it can see within 30 feet of it. The\
    \ target must make a DC 16 Constitution saving throw, taking dice:8d8|text(36)\
    \ (8d8) necrotic damage on a failed save, or half as much damage on a successful\
    \ one. If the saving throw fails by 5 or more, the target instead takes 64 necrotic\
    \ damage. The target dies if reduced to 0 hit points by this ray."
  "name": "Death Ray (Recharge 5-6)"
"source":
- "MPMM"
- "VGM"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Catoblepas.webp"
```
^statblock

```encounter-table
name: Catoblepas
creatures:
 - 1: Catoblepas
```

## Environment

swamp