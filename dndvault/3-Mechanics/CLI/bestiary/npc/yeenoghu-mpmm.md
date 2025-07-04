---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/24
- monster/size/huge
- monster/type/fiend/demon
aliases: ["Yeenoghu"]
NoteIcon: monster
BestiaryType: fiend (demon)
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 270, Mordenkainen's Tome of Foes p. 155
---
# [Yeenoghu](3-Mechanics\CLI\bestiary\npc/yeenoghu-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 270, Mordenkainen's Tome of Foes p. 155*  

```statblock
"name": "Yeenoghu (MPMM)"
"size": "Huge"
"type": "fiend"
"subtype": "demon"
"alignment": "Chaotic Evil"
"ac": !!int "20"
"ac_class": "natural armor"
"hp": !!int "333"
"hit_dice": "23d12 + 184"
"stats":
- !!int "29"
- !!int "16"
- !!int "26"
- !!int "26"
- !!int "24"
- !!int "15"
"speed": "50 ft."
"saves":
  "Dexterity": !!int "10"
  "Wisdom": !!int "14"
  "Constitution": !!int "15"
"skillsaves":
  "Intimidation": !!int "9"
  "Perception": !!int "14"
"damage_resistances": "cold, fire, lightning"
"damage_immunities": "bludgeoning, piercing, slashing that is nonmagical"
"condition_immunities": "[charmed](/3-Mechanics/CLI/rules/conditions.md#charmed),\
  \ [exhaustion](/3-Mechanics/CLI/rules/conditions.md#exhaustion), [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened),\
  \ [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)"
"senses": "truesight 120 ft., passive Perception 24"
"languages": "all, telepathy 120 ft."
"cr": "24"
"traits":
- "desc": "Yeenoghu casts one of the following spells, requiring no material components\
    \ and using Charisma as the spellcasting ability (spell save DC 17):\n\nAt will:\
    \ [detect magic](/3-Mechanics/CLI/spells/detect-magic.md)\n\n1/day: [teleport](/3-Mechanics/CLI/spells/teleport.md)\n\
    \n3/day each: [dispel magic](/3-Mechanics/CLI/spells/dispel-magic.md), [fear](/3-Mechanics/CLI/spells/fear.md),\
    \ [invisibility](/3-Mechanics/CLI/spells/invisibility.md)"
  "name": "Spellcasting"
- "desc": "If Yeenoghu fails a saving throw, he can choose to succeed instead."
  "name": "Legendary Resistance (3/Day)"
- "desc": "Yeenoghu has advantage on saving throws against spells and other magical\
    \ effects."
  "name": "Magic Resistance"
"actions":
- "desc": "Yeenoghu makes three Flail attacks."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+16 (+16) to hit, reach 15 ft., one\
    \ target. Hit: dice:2d12 + 9|text(22) (2d12 + 9) force damage. If it's his\
    \ turn, Yeenoghu can cause the target to suffer one of the following additional\
    \ effects, each of which he can apply only once per turn\n\n- Confusion. The\
    \ target must succeed on a DC 17 Wisdom saving throw or be affected by the confusion\
    \ spell until the start of Yeenoghu's next turn.  \n- Force. The target takes\
    \ an extra dice:2d12|text(13) (2d12) force damage.  \n- Paralysis. The\
    \ target must succeed on a DC 17 Constitution saving throw or be [paralyzed](/3-Mechanics/CLI/rules/conditions.md#paralyzed)\
    \ until the start of Yeenoghu's next turn.  "
  "name": "Flail"
- "desc": "Melee Weapon Attack: dice: d20+16 (+16) to hit, reach 10 ft., one\
    \ target. Hit: dice:2d10 + 9|text(20) (2d10 + 9) acid damage."
  "name": "Bite"
"bonus_actions":
- "desc": "When Yeenoghu reduces a creature to 0 hit points with a melee attack, he\
    \ moves up to half his speed and makes one Bite attack."
  "name": "Rampage"
"legendary_actions":
- "desc": "Yeenoghu moves up to his speed."
  "name": "Charge"
- "desc": "Yeenoghu makes one Flail attack. If the attack hits, the target must succeed\
    \ on a DC 24 Strength saving throw or be pushed up to 15 feet in a straight line\
    \ away from Yeenoghu. If the saving throw fails by 5 or more, the target is also\
    \ knocked [prone](/3-Mechanics/CLI/rules/conditions.md#prone)."
  "name": "Swat Away"
- "desc": "Yeenoghu makes a separate Bite attack against each creature within 10 feet\
    \ of him."
  "name": "Savage (Costs 2 Actions)"
"source":
- "MPMM"
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Yeenoghu.webp"
```
^statblock

```encounter-table
name: Yeenoghu
creatures:
 - 1: Yeenoghu
```