---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mm
- monster/cr/4
- monster/environment/underdark
- monster/size/large
- monster/type/undead
aliases: ["Bone Naga (Guardian)"]
NoteIcon: monster
BestiaryType: undead
SourceType: Bestiary
BookSource: Monster Manual p. 233
---
# [Bone Naga (Guardian)](3-Mechanics\CLI\bestiary\undead/bone-naga-guardian.md)
*Source: Monster Manual p. 233*  

In response to the long history of conflict between the yuan-ti and the nagas, yuan-ti created a necromantic ritual that could halt a naga's resurrection by transforming the living naga into a skeletal undead servitor. A bone naga retains only a few of the spells it knew in life.

## Nagas

Nagas are intelligent serpents that inhabit the ruins of the past, amassing arcane treasures and knowledge.

The first nagas were created as immortal guardians by a humanoid race long lost to history. When this race died out, the nagas deemed themselves the rightful inheritors of their masters' treasures and magical lore. Industrious and driven, nagas occasionally venture out from their lairs to track down magic items or rare spellbooks.

Nagas never feel the ravages of time or succumb to sickness. Even if it is struck down, a naga's immortal spirit reforms in a new body in a matter of days, ready to continue its eternal work.

### Benevolent Dictators and Brutal Tyrants

A naga rules its domain with absolute authority. Whether it rules with compassion or by terrorizing its subjects, the naga believes itself the master of all other creatures that inhabit its domain.

### Rivalry

Nagas have a long-standing enmity with the yuan-ti, with each race seeing itself as the epitome of serpentine evolution. Though cooperation between them is rare, nagas and yuan-ti sometimes set aside their differences to work toward common objectives. However, yuan-ti always chafe under a naga's authority.

### Immortal Nature

A naga doesn't require air, food, drink, or sleep.

```statblock
"name": "Bone Naga (Guardian)"
"size": "Large"
"type": "undead"
"alignment": "Lawful Evil"
"ac": !!int "15"
"ac_class": "natural armor"
"hp": !!int "58"
"hit_dice": "9d10 + 9"
"stats":
- !!int "15"
- !!int "16"
- !!int "12"
- !!int "15"
- !!int "15"
- !!int "16"
"speed": "30 ft."
"damage_immunities": "poison"
"condition_immunities": "[charmed](/3-Mechanics/CLI/rules/conditions.md#charmed),\
  \ [exhaustion](/3-Mechanics/CLI/rules/conditions.md#exhaustion), [paralyzed](/3-Mechanics/CLI/rules/conditions.md#paralyzed),\
  \ [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)"
"senses": "darkvision 60 ft., passive Perception 12"
"languages": "Common plus one other language"
"cr": "4"
"traits":
- "desc": "The naga is a 5th-level spellcaster (spell save DC 12, dice: d20+4 (+4)\
    \ to hit with spell attacks) that needs only verbal components to cast its spells.\
    \ Its spellcasting ability is Wisdom, and it has the following cleric spells prepared:\n\
    \nCantrips (at will): [mending](/3-Mechanics/CLI/spells/mending.md), [sacred\
    \ flame](/3-Mechanics/CLI/spells/sacred-flame.md), [thaumaturgy](/3-Mechanics/CLI/spells/thaumaturgy.md)\n\
    \n1st level (4 slots): [command](/3-Mechanics/CLI/spells/command.md), [shield\
    \ of faith](/3-Mechanics/CLI/spells/shield-of-faith.md)\n\n2nd level (3 slots):\
    \ [calm emotions](/3-Mechanics/CLI/spells/calm-emotions.md), [hold person](/3-Mechanics/CLI/spells/hold-person.md)\n\
    \n3rd level (2 slots): [bestow curse](/3-Mechanics/CLI/spells/bestow-curse.md)"
  "name": "Spellcasting"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+5 (+5) to hit, reach 10 ft., one creature.\
    \ Hit: dice:2d6 + 3|text(10) (2d6 + 3) piercing damage plus dice:3d6|text(10)\
    \ (3d6) poison damage."
  "name": "Bite"
"source":
- "MM"
- "ToA"
- "WDMM"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MM/Bone%20Naga%20%28Guardian%29.webp"
```
^statblock

```encounter-table
name: Bone Naga (Guardian)
creatures:
 - 1: Bone Naga (Guardian)
```

## Environment

underdark