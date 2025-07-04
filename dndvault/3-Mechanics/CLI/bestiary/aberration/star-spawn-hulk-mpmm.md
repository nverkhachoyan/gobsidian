---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/10
- monster/size/large
- monster/type/aberration
aliases: ["Star Spawn Hulk"]
NoteIcon: monster
BestiaryType: aberration
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 227, Mordenkainen's Tome of Foes p. 234
---
# [Star Spawn Hulk](3-Mechanics\CLI\bestiary\aberration/star-spawn-hulk-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 227, Mordenkainen's Tome of Foes p. 234*  

```statblock
"name": "Star Spawn Hulk (MPMM)"
"size": "Large"
"type": "aberration"
"alignment": "Typically  Chaotic Evil"
"ac": !!int "16"
"ac_class": "natural armor"
"hp": !!int "136"
"hit_dice": "13d10 + 65"
"stats":
- !!int "20"
- !!int "8"
- !!int "21"
- !!int "7"
- !!int "12"
- !!int "9"
"speed": "30 ft."
"saves":
  "Dexterity": !!int "3"
  "Wisdom": !!int "5"
"skillsaves":
  "Perception": !!int "5"
"damage_resistances": "bludgeoning, piercing, slashing from nonmagical attacks"
"condition_immunities": "[charmed](/3-Mechanics/CLI/rules/conditions.md#charmed),\
  \ [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened)"
"senses": "darkvision 60 ft., passive Perception 15"
"languages": "Deep Speech"
"cr": "10"
"traits":
- "desc": "If the hulk takes psychic damage, each creature within 10 feet of the hulk\
    \ takes that damage instead; the hulk takes none of the damage. In addition, the\
    \ hulk's thoughts and location can't be discerned by magic."
  "name": "Psychic Mirror"
"actions":
- "desc": "The hulk makes two Slam attacks. If both attacks hit the same target, the\
    \ target also takes dice:2d8|text(9) (2d8) psychic damage and must succeed\
    \ on a DC 17 Constitution saving throw or be [stunned](/3-Mechanics/CLI/rules/conditions.md#stunned)\
    \ until the end of the target's next turn."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+9 (+9) to hit, reach 10 ft., one target.\
    \ Hit: dice:2d8 + 5|text(14) (2d8 + 5) bludgeoning damage."
  "name": "Slam"
- "desc": "The hulk makes a separate Slam attack against each creature within 10 feet\
    \ of it. Each creature that is hit must also succeed on a DC 17 Dexterity saving\
    \ throw or be knocked [prone](/3-Mechanics/CLI/rules/conditions.md#prone)."
  "name": "Reaping Arms (Recharge 5-6)"
"source":
- "MPMM"
- "MTF"
- "BMT"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Star%20Spawn%20Hulk.webp"
```
^statblock

```encounter-table
name: Star Spawn Hulk
creatures:
 - 1: Star Spawn Hulk
```