---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/vgm
- monster/cr/5
- monster/environment/hill
- monster/environment/mountain
- monster/environment/underdark
- monster/size/medium
- monster/type/fiend/demon
- monster/type/fiend/orc
aliases: ["Tanarukk"]
NoteIcon: monster
BestiaryType: fiend (demon, orc)
SourceType: Bestiary
BookSource: Volo's Guide to Monsters p. 186
---
# [Tanarukk](3-Mechanics\CLI\bestiary\fiend/tanarukk-vgm.md)
*Source: Volo's Guide to Monsters p. 186*  

To the common folk of the world, an orc is an orc. They know that any one of these savages can tear an ordinary person to pieces, so no further distinction is necessary.

Orcs know better. Different groups of orcs exist within a tribe, the actions of each dictated by the deity they pay homage to. To complement the various kinds of warriors that spill forth to ravage the countryside, each tribe has members that remain deep inside the lair, seldom if ever seeing what lies outside the darkness of their den.

In addition, orcs have special relationships with two creatures that are sometimes found in their company: the aurochs, a great bull that serves as a mount for warriors that revere Bahgtru, and the tanarukk, a demon-orc crossbreed that is so depraved and destructive that even orcs seek to kill it. The aurochs is described in appendix A. The tanarukk is described below.

## Tanarukk

When demonic corruption taints a tribe's leadership, orcs might turn to abyssal magic to make tanarukks. Evil humans who control orcs also use such power to bolster their followers' strength.

The demon lord Baphomet gladly shares the secret of creating tanarukks with those who entreat him for power. The process corrupts an unborn orc of the tribe, transforming it at birth into a creature much more savage than an orc.

Although tanarukks are fearsome fighters, they are a threat to their allies off the battlefield. Within the tribe's lair, a tanarukk is destructive and volatile, and best kept imprisoned. Sooner or later, a free tanarukk rampages through the tribe, attempting to take over by force. Most such coups fail, but at great cost to the tribe. If a tanarukk does seize the leadership of a tribe, reckless war is the course it inevitably chooses.

If a tanarukk manages to breed, its blood taints numerous subsequent generations, so its female descendants randomly produce tanarukks. Rather than risk raising a natural-born tanarukk, most tribes slay such abominations.

```statblock
"name": "Tanarukk (VGM)"
"size": "Medium"
"type": "fiend"
"subtype": "demon, orc"
"alignment": "Chaotic Evil"
"ac": !!int "14"
"ac_class": "natural armor"
"hp": !!int "95"
"hit_dice": "10d8 + 50"
"stats":
- !!int "18"
- !!int "13"
- !!int "20"
- !!int "9"
- !!int "9"
- !!int "9"
"speed": "30 ft."
"skillsaves":
  "Intimidation": !!int "2"
  "Perception": !!int "2"
"damage_resistances": "fire, poison"
"senses": "darkvision 60 ft., passive Perception 12"
"languages": "Abyssal, Common, Orc"
"cr": "5"
"traits":
- "desc": "As a bonus action, the tanarukk can move up to its speed toward a hostile\
    \ creature that it can see."
  "name": "Aggressive"
- "desc": "The tanarukk has advantage on saving throws against spells and other magical\
    \ effects."
  "name": "Magic Resistance"
"actions":
- "desc": "The tanarukk makes two attacks: one with its bite and one with its greatsword."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+7 (+7) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d8 + 4|text(8) (1d8 + 4) piercing damage."
  "name": "Bite"
- "desc": "Melee Weapon Attack: dice: d20+7 (+7) to hit, reach 5 ft., one target.\
    \ Hit: dice:2d6 + 4|text(11) (2d6 + 4) slashing damage."
  "name": "Greatsword"
"reactions":
- "desc": "In response to being hit by a melee attack, the tanarukk can make one melee\
    \ weapon attack with advantage against the attacker."
  "name": "Unbridled Fury"
"source":
- "VGM"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/VGM/Tanarukk.webp"
```
^statblock

```encounter-table
name: Tanarukk
creatures:
 - 1: Tanarukk
```

## Environment

underdark, mountain, hill