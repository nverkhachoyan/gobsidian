---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/vgm
- monster/cr/9
- monster/environment/desert
- monster/environment/urban
- monster/size/medium
- monster/type/humanoid/any-race
aliases: ["Necromancer"]
NoteIcon: monster
BestiaryType: humanoid (any race)
SourceType: Bestiary
BookSource: Volo's Guide to Monsters p. 217, Divine Contention, Dragon of Icespire Peak
---
# [Necromancer](3-Mechanics\CLI\bestiary\humanoid/necromancer-vgm.md)
*Source: Volo's Guide to Monsters p. 217, Divine Contention, Dragon of Icespire Peak*  

Necromancers are specialist wizards who study the interaction of life, death, and undeath. Some like to dig up corpses to create undead slaves. A few use their powers for good, becoming hunters of the undead and risking their lives to save others.

```statblock
"name": "Necromancer (VGM)"
"size": "Medium"
"type": "humanoid"
"subtype": "any race"
"alignment": "Any alignment"
"ac": !!int "12"
"ac_class": "15 with [mage armor](/3-Mechanics/CLI/spells/mage-armor.md)"
"hp": !!int "66"
"hit_dice": "12d8 + 12"
"stats":
- !!int "9"
- !!int "14"
- !!int "12"
- !!int "17"
- !!int "12"
- !!int "11"
"speed": "30 ft."
"saves":
  "Wisdom": !!int "5"
  "Intelligence": !!int "7"
"skillsaves":
  "History": !!int "7"
  "Arcana": !!int "7"
"damage_resistances": "necrotic"
"senses": "passive Perception 11"
"languages": "any four languages"
"cr": "9"
"traits":
- "desc": "The necromancer is a 12th-level spellcaster. Its spellcasting ability is\
    \ Intelligence (spell save DC 15, dice: d20+7 (+7) to hit with spell attacks).\
    \ The necromancer has the following wizard spells prepared:\n\nCantrips (at\
    \ will): [chill touch](/3-Mechanics/CLI/spells/chill-touch.md), [dancing lights](/3-Mechanics/CLI/spells/dancing-lights.md),\
    \ [mage hand](/3-Mechanics/CLI/spells/mage-hand.md), [mending](/3-Mechanics/CLI/spells/mending.md)\n\
    \n1st level (4 slots): [false life](/3-Mechanics/CLI/spells/false-life.md),\
    \ [mage armor](/3-Mechanics/CLI/spells/mage-armor.md), [ray of sickness](/3-Mechanics/CLI/spells/ray-of-sickness.md)\n\
    \n2nd level (3 slots): [blindness/deafness](/3-Mechanics/CLI/spells/blindness-deafness.md),\
    \ [ray of enfeeblement](/3-Mechanics/CLI/spells/ray-of-enfeeblement.md), [web](/3-Mechanics/CLI/spells/web.md)\n\
    \n3rd level (3 slots): [animate dead](/3-Mechanics/CLI/spells/animate-dead.md),\
    \ [bestow curse](/3-Mechanics/CLI/spells/bestow-curse.md), [vampiric touch](/3-Mechanics/CLI/spells/vampiric-touch.md)\n\
    \n4th level (3 slots): [blight](/3-Mechanics/CLI/spells/blight.md), [dimension\
    \ door](/3-Mechanics/CLI/spells/dimension-door.md), [stoneskin](/3-Mechanics/CLI/spells/stoneskin.md)\n\
    \n5th level (2 slots): [Bigby's hand](/3-Mechanics/CLI/spells/bigbys-hand.md),\
    \ [cloudkill](/3-Mechanics/CLI/spells/cloudkill.md)\n\n6th level (1 slots):\
    \ [circle of death](/3-Mechanics/CLI/spells/circle-of-death.md)\n\nNecromancy\
    \ spell of 1st level or higher"
  "name": "Spellcasting"
- "desc": "When necromancer kills a creature that is neither a construct nor undead\
    \ with a spell of 1st level or higher, the necromancer regains hit points equal\
    \ to twice the spell's level, or three times if it is a necromancy spell."
  "name": "Grim Harvest (1/Turn)"
"actions":
- "desc": "Melee Spell Attack: dice: d20+7 (+7) to hit, reach 5 ft., one creature.\
    \ Hit: dice:2d4|text(5) (2d4) necrotic damage."
  "name": "Withering Touch"
"source":
- "VGM"
- "TftYP"
- "DC"
- "DIP"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/VGM/Necromancer.webp"
```
^statblock

```encounter-table
name: Necromancer
creatures:
 - 1: Necromancer
```

## Environment

desert, urban