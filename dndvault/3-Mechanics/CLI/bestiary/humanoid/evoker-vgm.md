---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/vgm
- monster/cr/9
- monster/environment/urban
- monster/size/medium
- monster/type/humanoid/any-race
aliases: ["Evoker"]
NoteIcon: monster
BestiaryType: humanoid (any race)
SourceType: Bestiary
BookSource: Volo's Guide to Monsters p. 214, Dragon of Icespire Peak, Storm Lord's Wrath
---
# [Evoker](3-Mechanics\CLI\bestiary\humanoid/evoker-vgm.md)
*Source: Volo's Guide to Monsters p. 214, Dragon of Icespire Peak, Storm Lord's Wrath*  

Evokers are specialist wizards who harness magical energy and elemental forces to destroy. Many tend to be hotheaded and aggressive. Others are cold and reserved, unleashing their power at just the right moment to exploit an opponent's weakness.

```statblock
"name": "Evoker (VGM)"
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
"senses": "passive Perception 11"
"languages": "any four languages"
"cr": "9"
"traits":
- "desc": "The evoker is a 12th-level spellcaster. Its spellcasting ability is Intelligence\
    \ (spell save DC 15, dice: d20+7 (+7) to hit with spell attacks). The evoker\
    \ has the following wizard spells prepared:\n\nCantrips (at will): [fire bolt](/3-Mechanics/CLI/spells/fire-bolt.md),\
    \ [light](/3-Mechanics/CLI/spells/light.md), [prestidigitation](/3-Mechanics/CLI/spells/prestidigitation.md),\
    \ [ray of frost](/3-Mechanics/CLI/spells/ray-of-frost.md)\n\n1st level (4 slots):\
    \ [burning hands](/3-Mechanics/CLI/spells/burning-hands.md), [mage armor](/3-Mechanics/CLI/spells/mage-armor.md),\
    \ [magic missile](/3-Mechanics/CLI/spells/magic-missile.md)\n\n2nd level (3\
    \ slots): [mirror image](/3-Mechanics/CLI/spells/mirror-image.md), [misty step](/3-Mechanics/CLI/spells/misty-step.md),\
    \ [shatter](/3-Mechanics/CLI/spells/shatter.md)\n\n3rd level (3 slots): [counterspell](/3-Mechanics/CLI/spells/counterspell.md),\
    \ [fireball](/3-Mechanics/CLI/spells/fireball.md), [lightning bolt](/3-Mechanics/CLI/spells/lightning-bolt.md)\n\
    \n4th level (3 slots): [ice storm](/3-Mechanics/CLI/spells/ice-storm.md),\
    \ [stoneskin](/3-Mechanics/CLI/spells/stoneskin.md)\n\n5th level (2 slots):\
    \ [Bigby's hand](/3-Mechanics/CLI/spells/bigbys-hand.md), [cone of cold](/3-Mechanics/CLI/spells/cone-of-cold.md)\n\
    \n6th level (1 slots): [chain lightning](/3-Mechanics/CLI/spells/chain-lightning.md),\
    \ [wall of ice](/3-Mechanics/CLI/spells/wall-of-ice.md)\n\nEvocation spell"
  "name": "Spellcasting"
- "desc": "When the evoker casts an evocation spell that forces other creatures it\
    \ can see to make a saving throw, it can choose a number of them equal to 1 +\
    \ the spell's level. These creatures automatically succeed on their saving throws\
    \ against the spell. If a successful save means a chosen creature would take half\
    \ damage from the spell, it instead takes no damage from it."
  "name": "Sculpt Spells"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+3 (+3) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d6 - 1|text(2) (1d6 - 1) bludgeoning damage, or dice:1d8 -\
    \ 1|text(3) (1d8 - 1) bludgeoning damage if used with two hands."
  "name": "Quarterstaff"
"source":
- "VGM"
- "TftYP"
- "DIP"
- "SLW"
- "ERLW"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/VGM/Evoker.webp"
```
^statblock

```encounter-table
name: Evoker
creatures:
 - 1: Evoker
```

## Environment

urban