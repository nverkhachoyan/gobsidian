---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/11
- monster/environment/coastal
- monster/environment/forest
- monster/environment/swamp
- monster/environment/underdark
- monster/size/large
- monster/type/giant
aliases: ["Spirit Troll"]
NoteIcon: monster
BestiaryType: giant
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 247, Mordenkainen's Tome of Foes p. 244
---
# [Spirit Troll](3-Mechanics\CLI\bestiary\giant/spirit-troll-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 247, Mordenkainen's Tome of Foes p. 244*  

```statblock
"name": "Spirit Troll (MPMM)"
"size": "Large"
"type": "giant"
"alignment": "Typically  Chaotic Evil"
"ac": !!int "17"
"ac_class": "natural armor"
"hp": !!int "130"
"hit_dice": "20d10 + 20"
"stats":
- !!int "1"
- !!int "17"
- !!int "13"
- !!int "8"
- !!int "9"
- !!int "16"
"speed": "30 ft."
"skillsaves":
  "Perception": !!int "3"
"damage_resistances": "acid, cold, fire"
"damage_immunities": "bludgeoning, piercing, slashing from nonmagical attacks"
"condition_immunities": "[exhaustion](/3-Mechanics/CLI/rules/conditions.md#exhaustion),\
  \ [grappled](/3-Mechanics/CLI/rules/conditions.md#grappled), [paralyzed](/3-Mechanics/CLI/rules/conditions.md#paralyzed),\
  \ [petrified](/3-Mechanics/CLI/rules/conditions.md#petrified), [prone](/3-Mechanics/CLI/rules/conditions.md#prone),\
  \ [restrained](/3-Mechanics/CLI/rules/conditions.md#restrained), [unconscious](/3-Mechanics/CLI/rules/conditions.md#unconscious)"
"senses": "darkvision 60 ft., passive Perception 13"
"languages": "Giant"
"cr": "11"
"traits":
- "desc": "The troll can move through other creatures and objects as if they were\
    \ difficult terrain. It takes dice:1d10|text(5) (1d10) force damage if it\
    \ ends its turn inside an object."
  "name": "Incorporeal Movement"
- "desc": "The troll regains 10 hit points at the start of each of its turns. If the\
    \ troll takes psychic or force damage, this trait doesn't function at the start\
    \ of the troll's next turn. The troll dies only if it starts its turn with 0 hit\
    \ points and doesn't regenerate."
  "name": "Regeneration"
"actions":
- "desc": "The troll makes one Bite attack and two Claw attacks."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+7 (+7) to hit, reach 5 ft., one creature.\
    \ Hit: dice:3d10 + 3|text(19) (3d10 + 3) psychic damage, and the target\
    \ must succeed on a DC 15 Wisdom saving throw or be [stunned](/3-Mechanics/CLI/rules/conditions.md#stunned)\
    \ for 1 minute. The [stunned](/3-Mechanics/CLI/rules/conditions.md#stunned) target\
    \ can repeat the saving throw at the end of each of its turns, ending the effect\
    \ on itself on a success."
  "name": "Bite"
- "desc": "Melee Weapon Attack: dice: d20+7 (+7) to hit, reach 5 ft., one creature.\
    \ Hit: dice:3d10 + 3|text(19) (3d10 + 3) psychic damage."
  "name": "Claws"
"source":
- "MPMM"
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Spirit%20Troll.webp"
```
^statblock

```encounter-table
name: Spirit Troll
creatures:
 - 1: Spirit Troll
```

## Environment

coastal, forest, swamp, underdark