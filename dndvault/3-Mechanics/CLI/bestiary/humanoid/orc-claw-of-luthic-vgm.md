---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/vgm
- monster/cr/2
- monster/environment/mountain
- monster/environment/underdark
- monster/size/medium
- monster/type/humanoid/orc
aliases: ["Orc Claw of Luthic"]
NoteIcon: monster
BestiaryType: humanoid (orc)
SourceType: Bestiary
BookSource: Volo's Guide to Monsters p. 183
---
# [Orc Claw of Luthic](3-Mechanics\CLI\bestiary\humanoid/orc-claw-of-luthic-vgm.md)
*Source: Volo's Guide to Monsters p. 183*  

To the common folk of the world, an orc is an orc. They know that any one of these savages can tear an ordinary person to pieces, so no further distinction is necessary.

Orcs know better. Different groups of orcs exist within a tribe, the actions of each dictated by the deity they pay homage to. To complement the various kinds of warriors that spill forth to ravage the countryside, each tribe has members that remain deep inside the lair, seldom if ever seeing what lies outside the darkness of their den.

In addition, orcs have special relationships with two creatures that are sometimes found in their company: the aurochs, a great bull that serves as a mount for warriors that revere Bahgtru, and the tanarukk, a demon-orc crossbreed that is so depraved and destructive that even orcs seek to kill it. The aurochs is described in appendix A. The tanarukk is described below.

## Orc Claw of Luthic

Luthic is Gruumsh's wife and the paragon of maternity to all orcs. She is the Cave Mother, a fierce dweller in the darkness who raises new broods of orcs to be vicious and strong. Her symbol is the cave bear, and orc females raise such bears alongside orc whelps. Females particularly attracted to Luthic grow long nails and lacquer them, learning to use these claws as weapons much as Luthic uses her own.

Orc females devoted to Luthic are in charge of fortifying and maintaining an orc stronghold. They help to guarantee the survival of the tribe, and most are skilled in the healing arts. The most powerful among Luthic's disciples are the claws of Luthic, which can use the Cave Mother's magic to heal, protect, and curse.

```statblock
"name": "Orc Claw of Luthic (VGM)"
"size": "Medium"
"type": "humanoid"
"subtype": "orc"
"alignment": "Chaotic Evil"
"ac": !!int "14"
"ac_class": "[hide armor](/3-Mechanics/CLI/items/hide-armor.md)"
"hp": !!int "45"
"hit_dice": "6d8 + 18"
"stats":
- !!int "14"
- !!int "15"
- !!int "16"
- !!int "10"
- !!int "15"
- !!int "11"
"speed": "30 ft."
"skillsaves":
  "Medicine": !!int "4"
  "Intimidation": !!int "2"
  "Survival": !!int "4"
"senses": "darkvision 60 ft., passive Perception 12"
"languages": "Common, Orc"
"cr": "2"
"traits":
- "desc": "The orc is a 5th-level spellcaster. Its spellcasting ability is Wisdom\
    \ (spell save DC 12, dice: d20+4 (+4) to hit with spell attacks). The orc\
    \ has the following cleric spells prepared:\n\nCantrips (at will): [guidance](/3-Mechanics/CLI/spells/guidance.md),\
    \ [mending](/3-Mechanics/CLI/spells/mending.md), [resistance](/3-Mechanics/CLI/spells/resistance.md),\
    \ [thaumaturgy](/3-Mechanics/CLI/spells/thaumaturgy.md)\n\n1st level (4 slots):\
    \ [bane](/3-Mechanics/CLI/spells/bane.md), [cure wounds](/3-Mechanics/CLI/spells/cure-wounds.md),\
    \ [guiding bolt](/3-Mechanics/CLI/spells/guiding-bolt.md)\n\n2nd level (3 slots):\
    \ [augury](/3-Mechanics/CLI/spells/augury.md), [warding bond](/3-Mechanics/CLI/spells/warding-bond.md)\n\
    \n3rd level (2 slots): [bestow curse](/3-Mechanics/CLI/spells/bestow-curse.md),\
    \ [create food and water](/3-Mechanics/CLI/spells/create-food-and-water.md)"
  "name": "Spellcasting"
- "desc": "As a bonus action, the orc can move up to its speed toward a hostile creature\
    \ that it can see."
  "name": "Aggressive"
"actions":
- "desc": "The orc makes two claw attacks, or four claw attacks if it has fewer than\
    \ half of its hit points remaining."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+4 (+4) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d8 + 2|text(6) (1d8 + 2) slashing damage."
  "name": "Claw"
"source":
- "VGM"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/VGM/Orc%20Claw%20of%20Luthic.webp"
```
^statblock

```encounter-table
name: Orc Claw of Luthic
creatures:
 - 1: Orc Claw of Luthic
```

## Environment

underdark, mountain