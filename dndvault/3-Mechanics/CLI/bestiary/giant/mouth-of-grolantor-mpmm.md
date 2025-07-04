---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/6
- monster/environment/grassland
- monster/environment/hill
- monster/size/huge
- monster/type/giant/hill-giant
aliases: ["Mouth of Grolantor"]
NoteIcon: monster
BestiaryType: giant (hill giant)
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 187, Volo's Guide to Monsters p. 149
---
# [Mouth of Grolantor](3-Mechanics\CLI\bestiary\giant/mouth-of-grolantor-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 187, Volo's Guide to Monsters p. 149*  

```statblock
"name": "Mouth of Grolantor (MPMM)"
"size": "Huge"
"type": "giant"
"subtype": "hill giant"
"alignment": "Typically  Chaotic Evil"
"ac": !!int "14"
"ac_class": "natural armor"
"hp": !!int "105"
"hit_dice": "10d12 + 40"
"stats":
- !!int "21"
- !!int "10"
- !!int "18"
- !!int "5"
- !!int "7"
- !!int "5"
"speed": "50 ft."
"skillsaves":
  "Perception": !!int "1"
"condition_immunities": "[frightened](/3-Mechanics/CLI/rules/conditions.md#frightened)"
"senses": "passive Perception 11"
"languages": "Giant"
"cr": "6"
"traits":
- "desc": "The giant is immune to the [confusion](/3-Mechanics/CLI/spells/confusion.md)\
    \ spell.\n\nOn each of its turns, the giant uses all its movement to move toward\
    \ the nearest creature or whatever else it might perceive as food. Roll a dice:\
    \ d10|avg|noform (d10) at the start of each of the giant's turns to determine\
    \ its action for that turn:\n\n- 1–3. The giant makes three Fist attacks against\
    \ one random creature within reach. If no creatures are within reach, the giant\
    \ flies into a rage and gains advantage on all attack rolls until the end of its\
    \ next turn.  \n- 4–5. The giant makes one Fist attack against each creature\
    \ within reach. If no creatures are within reach, the giant makes one Fist attack\
    \ against itself.  \n- 6–7. The giant makes one Bite attack against one random\
    \ creature within reach. If no other creatures are within reach, its eyes glaze\
    \ over and it is [stunned](/3-Mechanics/CLI/rules/conditions.md#stunned) until\
    \ the start of its next turn.  \n- 8–10. The giant makes one Bite attack and\
    \ two Fist attacks against one random creature within reach. If no creatures are\
    \ within reach, the giant flies into a rage and gains advantage on all attack\
    \ rolls until the end of its next turn.  "
  "name": "Mouth of Chaos"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+8 (+8) to hit, reach 5 ft., one creature.\
    \ Hit: dice:3d6 + 5|text(15) (3d6 + 5) piercing damage, and the giant magically\
    \ regains hit points equal to the damage dealt."
  "name": "Bite"
- "desc": "Melee Weapon Attack: dice: d20+8 (+8) to hit, reach 10 ft., one target.\
    \ Hit: dice:3d8 + 5|text(18) (3d8 + 5) bludgeoning damage."
  "name": "Fist"
"source":
- "MPMM"
- "VGM"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Mouth%20of%20Grolantor.webp"
```
^statblock

```encounter-table
name: Mouth of Grolantor
creatures:
 - 1: Mouth of Grolantor
```

## Environment

grassland, hill