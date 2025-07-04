---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/vgm
- monster/cr/2
- monster/environment/forest
- monster/environment/grassland
- monster/environment/hill
- monster/environment/mountain
- monster/environment/underdark
- monster/size/medium
- monster/type/humanoid/orc
aliases: ["Orc Hand of Yurtrus"]
NoteIcon: monster
BestiaryType: humanoid (orc)
SourceType: Bestiary
BookSource: Volo's Guide to Monsters p. 184
---
# [Orc Hand of Yurtrus](3-Mechanics\CLI\bestiary\humanoid/orc-hand-of-yurtrus-vgm.md)
*Source: Volo's Guide to Monsters p. 184*  

To the common folk of the world, an orc is an orc. They know that any one of these savages can tear an ordinary person to pieces, so no further distinction is necessary.

Orcs know better. Different groups of orcs exist within a tribe, the actions of each dictated by the deity they pay homage to. To complement the various kinds of warriors that spill forth to ravage the countryside, each tribe has members that remain deep inside the lair, seldom if ever seeing what lies outside the darkness of their den.

In addition, orcs have special relationships with two creatures that are sometimes found in their company: the aurochs, a great bull that serves as a mount for warriors that revere Bahgtru, and the tanarukk, a demon-orc crossbreed that is so depraved and destructive that even orcs seek to kill it. The aurochs is described in appendix A. The tanarukk is described below.

## Orc Hand of Yurtrus

Yurtrus is the orc god of death and disease. He is a horrifying abomination covered in rot and infection, except for his perfect, smooth white hands.

Orc priests that oversee the line between life and death are known by the others in the tribe as hands of Yurtrus. They dwell on the fringes of an orc lair, usually communing with other orcs through the auspices of those who follow Luthic. The hands of Yurtrus wear pale gloves made of the bleached skin of other humanoids (preferably elves), symbolizing their connection with Yurtrus, and are sometimes called "white hands" as a result.

Every orc knows that the hands of Yurtrus are the tribe's gateway to the ancestors. Orcs who die having served the tribe well go on to rituals meant to send them to Gruumsh's realm.

As befits followers of a god who doesn't speak, hands of Yurtrus remove their tongues to emulate their deity, for a reason similar to why an eye of Gruumsh puts out one of its eyes.

```statblock
"name": "Orc Hand of Yurtrus (VGM)"
"size": "Medium"
"type": "humanoid"
"subtype": "orc"
"alignment": "Chaotic Evil"
"ac": !!int "12"
"ac_class": "[hide armor](/3-Mechanics/CLI/items/hide-armor.md)"
"hp": !!int "30"
"hit_dice": "4d8 + 12"
"stats":
- !!int "12"
- !!int "11"
- !!int "16"
- !!int "11"
- !!int "14"
- !!int "9"
"speed": "30 ft."
"skillsaves":
  "Medicine": !!int "4"
  "Intimidation": !!int "1"
  "Religion": !!int "2"
  "Arcana": !!int "2"
"senses": "darkvision 60 ft., passive Perception 12"
"languages": "understands Common and Orc but can't speak"
"cr": "2"
"traits":
- "desc": "The orc is a 4th-level spellcaster. Its spellcasting ability is Wisdom\
    \ (spell save DC 12, dice: d20+4 (+4) to hit with spell attacks). It requires\
    \ no verbal components to cast its spells. The orc has the following cleric spells\
    \ prepared:\n\nCantrips (at will): [guidance](/3-Mechanics/CLI/spells/guidance.md),\
    \ [mending](/3-Mechanics/CLI/spells/mending.md), [resistance](/3-Mechanics/CLI/spells/resistance.md),\
    \ [thaumaturgy](/3-Mechanics/CLI/spells/thaumaturgy.md)\n\n1st level (4 slots):\
    \ [bane](/3-Mechanics/CLI/spells/bane.md), [detect magic](/3-Mechanics/CLI/spells/detect-magic.md),\
    \ [inflict wounds](/3-Mechanics/CLI/spells/inflict-wounds.md), [protection from\
    \ evil and good](/3-Mechanics/CLI/spells/protection-from-evil-and-good.md)\n\n\
    2nd level (3 slots): [blindness/deafness](/3-Mechanics/CLI/spells/blindness-deafness.md),\
    \ [silence](/3-Mechanics/CLI/spells/silence.md)"
  "name": "Spellcasting"
- "desc": "As a bonus action, the orc can move up to its speed toward a hostile creature\
    \ that it can see."
  "name": "Aggressive"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+3 (+3) to hit, reach 5 ft., one target.\
    \ Hit: dice:2d8|text(9) (2d8) necrotic damage."
  "name": "Touch of the White Hand"
"source":
- "VGM"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/VGM/Orc%20Hand%20of%20Yurtrus.webp"
```
^statblock

```encounter-table
name: Orc Hand of Yurtrus
creatures:
 - 1: Orc Hand of Yurtrus
```

## Environment

underdark, mountain, grassland, forest, hill