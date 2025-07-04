---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/vgm
- monster/cr/6
- monster/environment/urban
- monster/size/medium
- monster/type/humanoid/any-race
aliases: ["Conjurer"]
NoteIcon: monster
BestiaryType: humanoid (any race)
SourceType: Bestiary
BookSource: Volo's Guide to Monsters p. 212
---
# [Conjurer](3-Mechanics\CLI\bestiary\humanoid/conjurer-vgm.md)
*Source: Volo's Guide to Monsters p. 212*  

Conjurers are specialist wizards who summon creatures from other planes and create materials out of thin air. Some conjurers use their magic to bolster armies or destroy enemies on battlefields, while others use summoned creatures to guard their lairs.

```statblock
"name": "Conjurer (VGM)"
"size": "Medium"
"type": "humanoid"
"subtype": "any race"
"alignment": "Any alignment"
"ac": !!int "12"
"ac_class": "15 with [mage armor](/3-Mechanics/CLI/spells/mage-armor.md)"
"hp": !!int "40"
"hit_dice": "9d8"
"stats":
- !!int "9"
- !!int "14"
- !!int "11"
- !!int "17"
- !!int "12"
- !!int "11"
"speed": "30 ft."
"saves":
  "Wisdom": !!int "4"
  "Intelligence": !!int "6"
"skillsaves":
  "History": !!int "6"
  "Arcana": !!int "6"
"senses": "passive Perception 11"
"languages": "any four languages"
"cr": "6"
"traits":
- "desc": "The conjurer is a 9th-level spellcaster. Its spellcasting ability is intelligence\
    \ (spell save DC 14, dice: d20+6 (+6) to hit with spell attacks). The conjurer\
    \ has the following wizard spells prepared:\n\nCantrips (at will): [acid splash](/3-Mechanics/CLI/spells/acid-splash.md),\
    \ [mage hand](/3-Mechanics/CLI/spells/mage-hand.md), [poison spray](/3-Mechanics/CLI/spells/poison-spray.md),\
    \ [prestidigitation](/3-Mechanics/CLI/spells/prestidigitation.md)\n\n1st level\
    \ (4 slots): [mage armor](/3-Mechanics/CLI/spells/mage-armor.md), [magic missile](/3-Mechanics/CLI/spells/magic-missile.md),\
    \ [unseen servant](/3-Mechanics/CLI/spells/unseen-servant.md)\n\n2nd level\
    \ (3 slots): [cloud of daggers](/3-Mechanics/CLI/spells/cloud-of-daggers.md),\
    \ [misty step](/3-Mechanics/CLI/spells/misty-step.md), [web](/3-Mechanics/CLI/spells/web.md)\n\
    \n3rd level (3 slots): [fireball](/3-Mechanics/CLI/spells/fireball.md), [stinking\
    \ cloud](/3-Mechanics/CLI/spells/stinking-cloud.md)\n\n4th level (3 slots):\
    \ [Evard's black tentacles](/3-Mechanics/CLI/spells/evards-black-tentacles.md),\
    \ [stoneskin](/3-Mechanics/CLI/spells/stoneskin.md)\n\n*5th level (2 slots):\
    \ [cloudkill](/3-Mechanics/CLI/spells/cloudkill.md), [conjure elemental](/3-Mechanics/CLI/spells/conjure-elemental.md)\n\
    \nConjuration spell of 1st level or higher"
  "name": "Spellcasting"
- "desc": "As a bonus action, the conjurer teleports up to 30 feet to an unoccupied\
    \ space that it can see. If it instead chooses a space within range that is occupied\
    \ by a willing Small or Medium creature, they both teleport, swapping places."
  "name": "Benign Transportation (Recharges after the Conjurer Casts a Conjuration\
    \ Spell of 1st Level or Higher)"
"actions":
- "desc": "Melee or Ranged Weapon Attack: dice: d20+5 (+5) to hit, reach 5 ft.\
    \ or range 20/60 ft., one target. Hit: dice:1d4 + 2|text(4) (1d4 + 2) piercing\
    \ damage."
  "name": "Dagger"
"source":
- "VGM"
- "TftYP"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/VGM/Conjurer.webp"
```
^statblock

```encounter-table
name: Conjurer
creatures:
 - 1: Conjurer
```

## Environment

urban