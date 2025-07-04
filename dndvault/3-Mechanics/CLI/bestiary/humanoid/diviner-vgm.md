---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/vgm
- monster/cr/8
- monster/environment/urban
- monster/size/medium
- monster/type/humanoid/any-race
aliases: ["Diviner"]
NoteIcon: monster
BestiaryType: humanoid (any race)
SourceType: Bestiary
BookSource: Volo's Guide to Monsters p. 213
---
# [Diviner](3-Mechanics\CLI\bestiary\humanoid/diviner-vgm.md)
*Source: Volo's Guide to Monsters p. 213*  

Diviners are specialist wizards who know that knowledge is power. They might act aloof and mysterious, hinting at omens and secrets, or they might be know-it-alls, spilling secrets and insights to advance their own status or reputation.

```statblock
"name": "Diviner (VGM)"
"size": "Medium"
"type": "humanoid"
"subtype": "any race"
"alignment": "Any alignment"
"ac": !!int "12"
"ac_class": "15 with [mage armor](/3-Mechanics/CLI/spells/mage-armor.md)"
"hp": !!int "67"
"hit_dice": "15d8"
"stats":
- !!int "9"
- !!int "14"
- !!int "11"
- !!int "18"
- !!int "12"
- !!int "11"
"speed": "30 ft."
"saves":
  "Wisdom": !!int "4"
  "Intelligence": !!int "7"
"skillsaves":
  "History": !!int "7"
  "Arcana": !!int "7"
"senses": "passive Perception 11"
"languages": "any four languages"
"cr": "8"
"traits":
- "desc": "The diviner is a 15th-level spellcaster. Its spellcasting ability is Intelligence\
    \ (spell save DC 15, dice: d20+7 (+7) to hit with spell attacks). The diviner\
    \ has the following wizard spells prepared:\n\nCantrips (at will): [fire bolt](/3-Mechanics/CLI/spells/fire-bolt.md),\
    \ [light](/3-Mechanics/CLI/spells/light.md), [mage hand](/3-Mechanics/CLI/spells/mage-hand.md),\
    \ [message](/3-Mechanics/CLI/spells/message.md), [true strike](/3-Mechanics/CLI/spells/true-strike.md)\n\
    \n1st level (4 slots): [detect magic](/3-Mechanics/CLI/spells/detect-magic.md),\
    \ [feather fall](/3-Mechanics/CLI/spells/feather-fall.md), [mage armor](/3-Mechanics/CLI/spells/mage-armor.md)\n\
    \n2nd level (3 slots): [detect thoughts](/3-Mechanics/CLI/spells/detect-thoughts.md),\
    \ [locate object](/3-Mechanics/CLI/spells/locate-object.md), [scorching ray](/3-Mechanics/CLI/spells/scorching-ray.md)\n\
    \n3rd level (3 slots): [clairvoyance](/3-Mechanics/CLI/spells/clairvoyance.md),\
    \ [fly](/3-Mechanics/CLI/spells/fly.md), [fireball](/3-Mechanics/CLI/spells/fireball.md)\n\
    \n4th level (3 slots): [arcane eye](/3-Mechanics/CLI/spells/arcane-eye.md),\
    \ [ice storm](/3-Mechanics/CLI/spells/ice-storm.md), [stoneskin](/3-Mechanics/CLI/spells/stoneskin.md)\n\
    \n5th level (2 slots): [Rary's telepathic bond](/3-Mechanics/CLI/spells/rarys-telepathic-bond.md),\
    \ [seeming](/3-Mechanics/CLI/spells/seeming.md)\n\n6th level (1 slots): [mass\
    \ suggestion](/3-Mechanics/CLI/spells/mass-suggestion.md), [true seeing](/3-Mechanics/CLI/spells/true-seeing.md)\n\
    \n7th level (1 slots): [delayed blast fireball](/3-Mechanics/CLI/spells/delayed-blast-fireball.md),\
    \ [teleport](/3-Mechanics/CLI/spells/teleport.md)\n\n8th level (1 slots):\
    \ [maze](/3-Mechanics/CLI/spells/maze.md)\n\n Divination spell of 1st level or\
    \ higher"
  "name": "Spellcasting"
- "desc": "When the diviner or a creature it can see makes an attack roll, a saving\
    \ throw, or an ability check, the diviner can roll a dice: d20|avg|noform (d20)\
    \ and choose to use this roll in place of the attack roll, saving throw, or ability\
    \ check."
  "name": "Portent (Recharges after the Diviner Casts a Divination Spell of 1st Level\
    \ or Higher)"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+2 (+2) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d6 - 1|text(2) (1d6 - 1) bludgeoning damage, or dice:1d8 -\
    \ 1|text(3) (1d8 - 1) bludgeoning damage if used with two hands."
  "name": "Quarterstaff"
"source":
- "VGM"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/VGM/Diviner.webp"
```
^statblock

```encounter-table
name: Diviner
creatures:
 - 1: Diviner
```

## Environment

urban