---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/4
- monster/environment/grassland
- monster/environment/hill
- monster/environment/mountain
- monster/size/large
- monster/type/giant
aliases: ["Ogre Battering Ram"]
NoteIcon: monster
BestiaryType: giant
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 200, Mordenkainen's Tome of Foes p. 220
---
# [Ogre Battering Ram](3-Mechanics\CLI\bestiary\giant/ogre-battering-ram-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 200, Mordenkainen's Tome of Foes p. 220*  

```statblock
"name": "Ogre Battering Ram (MPMM)"
"size": "Large"
"type": "giant"
"alignment": "Typically  Chaotic Evil"
"ac": !!int "11"
"ac_class": "[hide armor](/3-Mechanics/CLI/items/hide-armor.md)"
"hp": !!int "76"
"hit_dice": "9d10 + 27"
"stats":
- !!int "19"
- !!int "8"
- !!int "16"
- !!int "5"
- !!int "7"
- !!int "7"
"speed": "40 ft."
"senses": "darkvision 60 ft., passive Perception 8"
"languages": "Common, Giant"
"cr": "4"
"traits":
- "desc": "The ogre deals double damage to objects and structures."
  "name": "Siege Monster"
"actions":
- "desc": "The ogre makes two Bash attacks."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+6 (+6) to hit, reach 5 ft., one target.\
    \ Hit: dice:2d10 + 4|text(15) (2d10 + 4) bludgeoning damage, and the ogre\
    \ can push the target 5 feet away if the target is Huge or smaller."
  "name": "Bash"
"reactions":
- "desc": "When a creature enters a space within 5 feet of the ogre, the ogre makes\
    \ a Bash attack against that creature. If the attack hits, the target's speed\
    \ is reduced to 0 until the start of the ogre's next turn."
  "name": "Block the Path"
"source":
- "MPMM"
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Ogre%20Battering%20Ram.webp"
```
^statblock

```encounter-table
name: Ogre Battering Ram
creatures:
 - 1: Ogre Battering Ram
```

## Environment

grassland, hill, mountain