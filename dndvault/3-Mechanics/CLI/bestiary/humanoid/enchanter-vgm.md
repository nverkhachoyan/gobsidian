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
aliases: ["Enchanter"]
NoteIcon: monster
BestiaryType: humanoid (any race)
SourceType: Bestiary
BookSource: Volo's Guide to Monsters p. 213
---
# [Enchanter](3-Mechanics\CLI\bestiary\humanoid/enchanter-vgm.md)
*Source: Volo's Guide to Monsters p. 213*  

Enchanters are specialist wizards who understand how to alter and control minds using magic. They might be personable and interesting, using magic to manipulate people only when banter and conventional persuasion fails, or they might be rude and demanding, using and relying on [charmed](/3-Mechanics/CLI/rules/conditions.md#charmed), obedient minions.

```statblock
"name": "Enchanter (VGM)"
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
- "desc": "The enchanter is a 9th-level spellcaster. Its spellcasting ability is Intelligence\
    \ (spell save DC 14, dice: d20+6 (+6) to hit with spell attacks). The enchanter\
    \ has the following wizard spells prepared:\n\nCantrips (at will): [friends](/3-Mechanics/CLI/spells/friends.md),\
    \ [mage hand](/3-Mechanics/CLI/spells/mage-hand.md), [mending](/3-Mechanics/CLI/spells/mending.md),\
    \ [message](/3-Mechanics/CLI/spells/message.md)\n\n1st level (4 slots): [charm\
    \ person](/3-Mechanics/CLI/spells/charm-person.md), [mage armor](/3-Mechanics/CLI/spells/mage-armor.md),\
    \ [magic missile](/3-Mechanics/CLI/spells/magic-missile.md)\n\n2nd level (3\
    \ slots): [hold person](/3-Mechanics/CLI/spells/hold-person.md), [invisibility](/3-Mechanics/CLI/spells/invisibility.md),\
    \ [suggestion](/3-Mechanics/CLI/spells/suggestion.md)\n\n3rd level (3 slots):\
    \ [fireball](/3-Mechanics/CLI/spells/fireball.md), [haste](/3-Mechanics/CLI/spells/haste.md),\
    \ [tongues](/3-Mechanics/CLI/spells/tongues.md)\n\n4th level (3 slots): [dominate\
    \ beast](/3-Mechanics/CLI/spells/dominate-beast.md), [stoneskin](/3-Mechanics/CLI/spells/stoneskin.md)\n\
    \n5th level (2 slots): [hold monster](/3-Mechanics/CLI/spells/hold-monster.md)\n\
    \nEnchantment spell of 1st level or higher"
  "name": "Spellcasting"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+2 (+2) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d6 - 1|text(2) (1d6 - 1) bludgeoning damage, or dice:1d8 -\
    \ 1|text(3) (1d8 - 1) bludgeoning damage if used with two hands."
  "name": "Quarterstaff"
"reactions":
- "desc": "The enchanter tries to magically divert an attack made against it, provided\
    \ that the attacker is within 30 feet of it and visible to it. The enchanter must\
    \ decide to do so before the attack hits or misses.\n\nThe attacker must make\
    \ a DC 14 Wisdom saving throw. On a failed save, the attacker targets the creature\
    \ closest to it, other than the enchanter or itself. If multiple creatures are\
    \ closest, the attacker chooses which one to target."
  "name": "Instinctive Charm (Recharges after the Enchanter Casts an Enchantment Spell\
    \ of 1st level or Higher)"
"source":
- "VGM"
- "TftYP"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/VGM/Enchanter.webp"
```
^statblock

```encounter-table
name: Enchanter
creatures:
 - 1: Enchanter
```

## Environment

urban