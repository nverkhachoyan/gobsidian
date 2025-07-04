---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/vgm
- monster/cr/1-2
- monster/environment/forest
- monster/environment/grassland
- monster/environment/hill
- monster/environment/mountain
- monster/environment/underdark
- monster/size/medium
- monster/type/humanoid/orc
aliases: ["Orc Nurtured One of Yurtrus"]
NoteIcon: monster
BestiaryType: humanoid (orc)
SourceType: Bestiary
BookSource: Volo's Guide to Monsters p. 184
---
# [Orc Nurtured One of Yurtrus](3-Mechanics\CLI\bestiary\humanoid/orc-nurtured-one-of-yurtrus-vgm.md)
*Source: Volo's Guide to Monsters p. 184*  

To the common folk of the world, an orc is an orc. They know that any one of these savages can tear an ordinary person to pieces, so no further distinction is necessary.

Orcs know better. Different groups of orcs exist within a tribe, the actions of each dictated by the deity they pay homage to. To complement the various kinds of warriors that spill forth to ravage the countryside, each tribe has members that remain deep inside the lair, seldom if ever seeing what lies outside the darkness of their den.

In addition, orcs have special relationships with two creatures that are sometimes found in their company: the aurochs, a great bull that serves as a mount for warriors that revere Bahgtru, and the tanarukk, a demon-orc crossbreed that is so depraved and destructive that even orcs seek to kill it. The aurochs is described in appendix A. The tanarukk is described below.

## Orc Nurtured One of Yurtrus

When plague strikes a tribe, the hands of Yurtrus isolate the sick. The priests then minister to those who can be saved but not healed. The hands cultivate the sickness of these nurtured ones, turning them into instruments of defense and weapons of war. When orcs go to battle, a band of nurtured ones might charge in first-to give themselves up while softening up the enemy by spreading Yurtrus's vile blessing in its ranks.

```statblock
"name": "Orc Nurtured One of Yurtrus (VGM)"
"size": "Medium"
"type": "humanoid"
"subtype": "orc"
"alignment": "Chaotic Evil"
"ac": !!int "9"
"hp": !!int "30"
"hit_dice": "4d8 + 12"
"stats":
- !!int "15"
- !!int "8"
- !!int "16"
- !!int "7"
- !!int "11"
- !!int "7"
"speed": "30 ft."
"senses": "darkvision 60 ft., passive Perception 10"
"languages": "Common, Orc"
"cr": "1/2"
"traits":
- "desc": "As a bonus action, the orc can move up to its speed toward a hostile creature\
    \ that it can see."
  "name": "Aggressive"
- "desc": "When the orc is reduced to 0 hit points, it explodes, and any creature\
    \ within 10 feet of it must make a DC 13 Constitution saving throw. On a failed\
    \ save, the creature takes dice:4d6|text(14) (4d6) poison damage and becomes\
    \ [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned). On a success, the\
    \ creature takes half as much damage and isn't [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned).\
    \ A creature [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned) by this\
    \ effect can repeat the save at the end of each of its turn, ending the effect\
    \ on itself on a success. While [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)\
    \ by this effect, a creature can't regain hit points."
  "name": "Corrupted Carrier"
- "desc": "The orc has advantage on saving throws against poison and disease."
  "name": "Nurtured One of Yurtrus"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+4 (+4) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d4 + 2|text(4) (1d4 + 2) slashing damage plus dice:1d4|text(2)\
    \ (1d4) necrotic damage."
  "name": "Claws"
- "desc": "The orc reduces itself to 0 hit points, triggering its Corrupted Carrier\
    \ trait."
  "name": "Corrupted Vengeance"
"source":
- "VGM"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/VGM/Orc%20Nurtured%20One%20of%20Yurtrus.webp"
```
^statblock

```encounter-table
name: Orc Nurtured One of Yurtrus
creatures:
 - 1: Orc Nurtured One of Yurtrus
```

## Environment

underdark, mountain, grassland, forest, hill