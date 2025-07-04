---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/vgm
- monster/cr/3
- monster/environment/urban
- monster/size/medium
- monster/type/humanoid/any-race
aliases: ["Martial Arts Adept"]
NoteIcon: monster
BestiaryType: humanoid (any race)
SourceType: Bestiary
BookSource: Volo's Guide to Monsters p. 216
---
# [Martial Arts Adept](3-Mechanics\CLI\bestiary\humanoid/martial-arts-adept-vgm.md)
*Source: Volo's Guide to Monsters p. 216*  

Martial arts adepts are disciplined monks with extensive training in hand-to-hand combat. Some protect monasteries; others travel the world seeking enlightenment or new forms of combat to master. A few become bodyguards, trading their combat prowess and loyalty for food and lodging.

```statblock
"name": "Martial Arts Adept (VGM)"
"size": "Medium"
"type": "humanoid"
"subtype": "any race"
"alignment": "Any alignment"
"ac": !!int "16"
"hp": !!int "60"
"hit_dice": "11d8 + 11"
"stats":
- !!int "11"
- !!int "17"
- !!int "13"
- !!int "11"
- !!int "16"
- !!int "10"
"speed": "40 ft."
"skillsaves":
  "Stealth": !!int "5"
  "Insight": !!int "5"
  "Acrobatics": !!int "5"
"senses": "passive Perception 13"
"languages": "any one language (usually Common)"
"cr": "3"
"traits":
- "desc": "While the adept is wearing no armor and wielding no shield, its AC includes\
    \ its Wisdom modifier."
  "name": "Unarmored Defense"
"actions":
- "desc": "The adept makes three unarmed strikes or three dart attacks."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+5 (+5) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d8 + 3|text(7) (1d8 + 3) bludgeoning damage. If the target\
    \ is a creature, the adept can choose one of the following additional effects:\n\
    \n- The target must succeed on a DC 13 Strength saving throw or drop one item\
    \ it is holding (adept's choice).  \n- The target must succeed on a DC 13 Dexterity\
    \ saving throw or be knocked [prone](/3-Mechanics/CLI/rules/conditions.md#prone).\
    \  \n- The target must succeed on a DC 13 Constitution saving throw or be [stunned](/3-Mechanics/CLI/rules/conditions.md#stunned)\
    \ until the end of the adept's next turn.  "
  "name": "Unarmed Strike"
- "desc": "Ranged Weapon Attack: dice: d20+5 (+5) to hit, range 20/60 ft., one\
    \ target. Hit: dice:1d4 + 3|text(5) (1d4 + 3) piercing damage."
  "name": "Dart"
"reactions":
- "desc": "In response to being hit by a ranged weapon attack, the adept deflects\
    \ the missile. The damage it takes from the attack is reduced by dice: 1d10 +\
    \ 3|avg|noform (1d10 + 3). If the damage is reduced to 0, the adept catches\
    \ the missile if it's small enough to hold in one hand and the adept has a hand\
    \ free."
  "name": "Deflect Missile"
"source":
- "VGM"
- "WDH"
- "TftYP"
- "ToA"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/VGM/Martial%20Arts%20Adept.webp"
```
^statblock

```encounter-table
name: Martial Arts Adept
creatures:
 - 1: Martial Arts Adept
```

## Environment

urban