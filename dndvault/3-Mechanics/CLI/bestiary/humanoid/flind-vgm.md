---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/vgm
- monster/cr/9
- monster/environment/arctic
- monster/environment/forest
- monster/environment/grassland
- monster/environment/hill
- monster/size/medium
- monster/type/humanoid/gnoll
aliases: ["Flind"]
NoteIcon: monster
BestiaryType: humanoid (gnoll)
SourceType: Bestiary
BookSource: Volo's Guide to Monsters p. 153
---
# [Flind](3-Mechanics\CLI\bestiary\humanoid/flind-vgm.md)
*Source: Volo's Guide to Monsters p. 153*  

A flind is an exceptionally strong and vicious gnoll that commands and directs the war band it is a part of. It wields a flail imbued with powerful magic by Yeenoghu himself.

A war band can have only one flind, and that creature sets a war band's path. Because of its special connection to Yeenoghu, a flind uses god-given omens and demonic insight to guide the gnolls toward weak prey ripe for slaughter.

Unlike other humanoid leaders that might skulk behind their minions, a flind leads the charge in battle. Its flail causes wracking pain, paralysis, and disorientation in those struck by it.

```statblock
"name": "Flind (VGM)"
"size": "Medium"
"type": "humanoid"
"subtype": "gnoll"
"alignment": "Chaotic Evil"
"ac": !!int "16"
"ac_class": "[chain mail](/3-Mechanics/CLI/items/chain-mail.md)"
"hp": !!int "127"
"hit_dice": "15d8 + 60"
"stats":
- !!int "20"
- !!int "10"
- !!int "19"
- !!int "11"
- !!int "13"
- !!int "12"
"speed": "30 ft."
"saves":
  "Wisdom": !!int "5"
  "Constitution": !!int "8"
"skillsaves":
  "Intimidation": !!int "5"
  "Perception": !!int "5"
"senses": "darkvision 60 ft., passive Perception 15"
"languages": "Abyssal, Gnoll"
"cr": "9"
"traits":
- "desc": "If the flind isn't [incapacitated](/3-Mechanics/CLI/rules/conditions.md#incapacitated),\
    \ any creature with the Rampage trait can make a bite attack as a bonus action\
    \ while within 10 feet of the flind."
  "name": "Aura of Blood Thirst"
"actions":
- "desc": "The flind makes three attacks: one with each of its different flail attacks\
    \ or three with its longbow."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+9 (+9) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d10 + 5|text(10) (1d10 + 5) bludgeoning damage, and the target\
    \ must make a DC 16 Wisdom saving throw. On a failed save, the target must make\
    \ a melee attack against a random target within its reach on its next turn. If\
    \ it has no targets within its reach even after moving, it loses its action on\
    \ that turn."
  "name": "Flail of Madness"
- "desc": "Melee Weapon Attack: dice: d20+9 (+9) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d10 + 5|text(10) (1d10 + 5) bludgeoning damage plus dice:4d10|text(22)\
    \ (4d10) psychic damage."
  "name": "Flail of Pain"
- "desc": "Melee Weapon Attack: dice: d20+9 (+9) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d10 + 5|text(10) (1d10 + 5) bludgeoning damage, and the target\
    \ must succeed on a DC 16 Constitution saving throw or be [paralyzed](/3-Mechanics/CLI/rules/conditions.md#paralyzed)\
    \ until the end of its next turn."
  "name": "Flail of Paralysis"
- "desc": "Ranged Weapon Attack: dice: d20+4 (+4) to hit, range 150/600 ft.,\
    \ one target. Hit: dice:1d8|text(4) (1d8) piercing damage."
  "name": "Longbow"
"source":
- "VGM"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/VGM/Flind.webp"
```
^statblock

```encounter-table
name: Flind
creatures:
 - 1: Flind
```

## Environment

grassland, forest, hill, arctic