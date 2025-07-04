---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/vgm
- monster/cr/5
- monster/environment/urban
- monster/size/medium
- monster/type/humanoid/any-race
aliases: ["Transmuter"]
NoteIcon: monster
BestiaryType: humanoid (any race)
SourceType: Bestiary
BookSource: Volo's Guide to Monsters p. 218
---
# [Transmuter](3-Mechanics\CLI\bestiary\humanoid/transmuter-vgm.md)
*Source: Volo's Guide to Monsters p. 218*  

Transmuters are specialist wizards who embrace change, rail against the status quo, and view magical transmutation as a path to riches, enlightenment, or apotheosis.

```statblock
"name": "Transmuter (VGM)"
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
"cr": "5"
"traits":
- "desc": "The transmuter is a 9th-level spellcaster. Its spellcasting ability is\
    \ Intelligence (spell save DC 14, dice: d20+6 (+6) to hit with spell attacks).\
    \ The transmuter has the following wizard spells prepared:\n\nCantrips (at will):\
    \ [light](/3-Mechanics/CLI/spells/light.md), [mending](/3-Mechanics/CLI/spells/mending.md),\
    \ [prestidigitation](/3-Mechanics/CLI/spells/prestidigitation.md), [ray of frost](/3-Mechanics/CLI/spells/ray-of-frost.md)\n\
    \n1st level (4 slots): [chromatic orb](/3-Mechanics/CLI/spells/chromatic-orb.md),\
    \ [expeditious retreat](/3-Mechanics/CLI/spells/expeditious-retreat.md), [mage\
    \ armor](/3-Mechanics/CLI/spells/mage-armor.md)\n\n2nd level (3 slots): [alter\
    \ self](/3-Mechanics/CLI/spells/alter-self.md), [hold person](/3-Mechanics/CLI/spells/hold-person.md),\
    \ [knock](/3-Mechanics/CLI/spells/knock.md)\n\n3rd level (3 slots): [blink](/3-Mechanics/CLI/spells/blink.md),\
    \ [fireball](/3-Mechanics/CLI/spells/fireball.md), [slow](/3-Mechanics/CLI/spells/slow.md)\n\
    \n4th level (3 slots): [polymorph](/3-Mechanics/CLI/spells/polymorph.md),\
    \ [stoneskin](/3-Mechanics/CLI/spells/stoneskin.md)\n\n5th level (1 slots):\
    \ [telekinesis](/3-Mechanics/CLI/spells/telekinesis.md)\n\nTransmutation spell\
    \ of 1st level or higher"
  "name": "Spellcasting"
- "desc": "The transmuter carries a magic stone it crafted that grants its bearer\
    \ one of the following effects:\n\nDarkvision out to a range of 60 feet\n\nAn\
    \ extra 10 feet of speed while the bearer is unencumbered\n\nProficiency with\
    \ Constitution saving throws\n\nResistance to acid, cold, fire, lightning, or\
    \ thunder damage (transmuter's choice whenever the transmuter chooses this benefit)\n\
    \nIf the transmuter has the stone and casts a transmutation spell of 1st level\
    \ or higher, it can change the effect of the stone."
  "name": "Transmuter's Stone"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+2 (+2) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d6 - 1|text(2) (1d6 - 1) bludgeoning damage, or dice:1d8 -\
    \ 1|text(3) (1d8 - 1) bludgeoning damage if used with two hands."
  "name": "Quarterstaff"
"source":
- "VGM"
- "TftYP"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/VGM/Transmuter.webp"
```
^statblock

```encounter-table
name: Transmuter
creatures:
 - 1: Transmuter
```

## Environment

urban