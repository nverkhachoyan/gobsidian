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
aliases: ["War Priest"]
NoteIcon: monster
BestiaryType: humanoid (any race)
SourceType: Bestiary
BookSource: Volo's Guide to Monsters p. 218, Divine Contention, Dragon of Icespire Peak, Sleeping Dragon's Wake
---
# [War Priest](3-Mechanics\CLI\bestiary\humanoid/war-priest-vgm.md)
*Source: Volo's Guide to Monsters p. 218, Divine Contention, Dragon of Icespire Peak, Sleeping Dragon's Wake*  

War priests worship deities of war and combat. They plan tactics, lead soldiers into battle, confront enemy spellcasters, and tend to casualties. A war priest might command an army or serve as a warlord's right hand on the battlefield.

```statblock
"name": "War Priest (VGM)"
"size": "Medium"
"type": "humanoid"
"subtype": "any race"
"alignment": "Any alignment"
"ac": !!int "18"
"ac_class": "[plate armor](/3-Mechanics/CLI/items/plate-armor.md)"
"hp": !!int "117"
"hit_dice": "18d8 + 36"
"stats":
- !!int "16"
- !!int "10"
- !!int "14"
- !!int "11"
- !!int "17"
- !!int "13"
"speed": "30 ft."
"saves":
  "Wisdom": !!int "7"
  "Constitution": !!int "6"
"skillsaves":
  "Intimidation": !!int "5"
  "Religion": !!int "4"
"senses": "passive Perception 13"
"languages": "any two languages"
"cr": "9"
"traits":
- "desc": "The priest is a 9th-level spellcaster. Its spellcasting ability is Wisdom\
    \ (spell save DC 15, dice: d20+7 (+7) to hit with spell attacks). It has the\
    \ following cleric spells prepared:\n\nCantrips (at will): [light](/3-Mechanics/CLI/spells/light.md),\
    \ [mending](/3-Mechanics/CLI/spells/mending.md), [sacred flame](/3-Mechanics/CLI/spells/sacred-flame.md),\
    \ [spare the dying](/3-Mechanics/CLI/spells/spare-the-dying.md)\n\n1st level\
    \ (4 slots): [divine favor](/3-Mechanics/CLI/spells/divine-favor.md), [guiding\
    \ bolt](/3-Mechanics/CLI/spells/guiding-bolt.md), [healing word](/3-Mechanics/CLI/spells/healing-word.md),\
    \ [shield of faith](/3-Mechanics/CLI/spells/shield-of-faith.md)\n\n2nd level\
    \ (3 slots): [lesser restoration](/3-Mechanics/CLI/spells/lesser-restoration.md),\
    \ [magic weapon](/3-Mechanics/CLI/spells/magic-weapon.md), [prayer of healing](/3-Mechanics/CLI/spells/prayer-of-healing.md),\
    \ [silence](/3-Mechanics/CLI/spells/silence.md), [spiritual weapon](/3-Mechanics/CLI/spells/spiritual-weapon.md)\n\
    \n3rd level (3 slots): [beacon of hope](/3-Mechanics/CLI/spells/beacon-of-hope.md),\
    \ [crusader's mantle](/3-Mechanics/CLI/spells/crusaders-mantle.md), [dispel magic](/3-Mechanics/CLI/spells/dispel-magic.md),\
    \ [revivify](/3-Mechanics/CLI/spells/revivify.md), [spirit guardians](/3-Mechanics/CLI/spells/spirit-guardians.md),\
    \ [water walk](/3-Mechanics/CLI/spells/water-walk.md)\n\n4th level (3 slots):\
    \ [banishment](/3-Mechanics/CLI/spells/banishment.md), [freedom of movement](/3-Mechanics/CLI/spells/freedom-of-movement.md),\
    \ [guardian of faith](/3-Mechanics/CLI/spells/guardian-of-faith.md), [stoneskin](/3-Mechanics/CLI/spells/stoneskin.md)\n\
    \n5th level (1 slots): [flame strike](/3-Mechanics/CLI/spells/flame-strike.md),\
    \ [mass cure wounds](/3-Mechanics/CLI/spells/mass-cure-wounds.md), [hold monster](/3-Mechanics/CLI/spells/hold-monster.md)"
  "name": "Spellcasting"
"actions":
- "desc": "The priest makes two melee attacks."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+7 (+7) to hit, reach 5 ft., one target.\
    \ Hit: dice:2d6 + 3|text(10) (2d6 + 3) bludgeoning damage."
  "name": "Maul"
"reactions":
- "desc": "The priest grants a +10 bonus to an attack roll made by itself or another\
    \ creature within 30 feet of it. The priest can make this choice after the roll\
    \ is made but before it hits or misses."
  "name": "Guided Strike (Recharges after a Short or Long Rest)"
"source":
- "VGM"
- "DC"
- "DIP"
- "SDW"
- "MOT"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/VGM/War%20Priest.webp"
```
^statblock

```encounter-table
name: War Priest
creatures:
 - 1: War Priest
```

## Environment

desert, urban