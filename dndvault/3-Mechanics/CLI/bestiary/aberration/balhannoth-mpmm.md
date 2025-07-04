---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/11
- monster/environment/coastal
- monster/environment/mountain
- monster/environment/underdark
- monster/size/large
- monster/type/aberration
aliases: ["Balhannoth"]
NoteIcon: monster
BestiaryType: aberration
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 55, Mordenkainen's Tome of Foes p. 119
---
# [Balhannoth](3-Mechanics\CLI\bestiary\aberration/balhannoth-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 55, Mordenkainen's Tome of Foes p. 119*  

```statblock
"name": "Balhannoth (MPMM)"
"size": "Large"
"type": "aberration"
"alignment": "Typically  Chaotic Evil"
"ac": !!int "17"
"ac_class": "natural armor"
"hp": !!int "114"
"hit_dice": "12d10 + 48"
"stats":
- !!int "17"
- !!int "8"
- !!int "18"
- !!int "6"
- !!int "15"
- !!int "8"
"speed": "25 ft., climb 25 ft."
"saves":
  "Constitution": !!int "8"
"skillsaves":
  "Perception": !!int "6"
"condition_immunities": "[blinded](/3-Mechanics/CLI/rules/conditions.md#blinded)"
"senses": "blindsight 500 ft. (blind beyond this radius), passive Perception 16"
"languages": "understands Deep Speech, telepathy 1 mile"
"cr": "11"
"traits":
- "desc": "If the balhannoth fails a saving throw, it can choose to succeed instead."
  "name": "Legendary Resistance (2/Day)"
"actions":
- "desc": "The balhannoth makes one Bite attack and two Tentacle attacks."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+7 (+7) to hit, reach 5 ft., one target.\
    \ Hit: dice:3d10 + 3|text(19) (3d10 + 3) piercing damage."
  "name": "Bite"
- "desc": "Melee Weapon Attack: dice: d20+7 (+7) to hit, reach 10 ft., one target.\
    \ Hit: dice:2d6 + 3|text(10) (2d6 + 3) bludgeoning damage, and the target\
    \ is [grappled](/3-Mechanics/CLI/rules/conditions.md#grappled) (escape DC 15)\
    \ and is moved up to 5 feet toward the balhannoth. Until this grapple ends, the\
    \ target is [restrained](/3-Mechanics/CLI/rules/conditions.md#restrained), and\
    \ the balhannoth can't use this tentacle against other targets. The balhannoth\
    \ has four tentacles."
  "name": "Tentacle"
"legendary_actions":
- "desc": "The balhannoth makes one Bite attack against one creature it has [grappled](/3-Mechanics/CLI/rules/conditions.md#grappled)."
  "name": "Bite"
- "desc": "The balhannoth teleports, along with any equipment it is wearing or carrying\
    \ and any creatures it has [grappled](/3-Mechanics/CLI/rules/conditions.md#grappled),\
    \ up to 60 feet to an unoccupied space it can see."
  "name": "Teleport"
- "desc": "The balhannoth magically becomes [invisible](/3-Mechanics/CLI/rules/conditions.md#invisible)\
    \ for up to 10 minutes or until immediately after it makes an attack roll."
  "name": "Vanish"
"source":
- "MPMM"
- "MTF"
- "RtG"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Balhannoth.webp"
```
^statblock

```encounter-table
name: Balhannoth
creatures:
 - 1: Balhannoth
```

## Environment

coastal, mountain, underdark