---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/2
- monster/environment/mountain
- monster/environment/underdark
- monster/size/medium
- monster/type/humanoid/dwarf
aliases: ["Duergar Mind Master"]
NoteIcon: monster
BestiaryType: humanoid (dwarf)
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 108, Mordenkainen's Tome of Foes p. 189
---
# [Duergar Mind Master](3-Mechanics\CLI\bestiary\humanoid/duergar-mind-master-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 108, Mordenkainen's Tome of Foes p. 189*  

```statblock
"name": "Duergar Mind Master (MPMM)"
"size": "Medium"
"type": "humanoid"
"subtype": "dwarf"
"alignment": "Any alignment"
"ac": !!int "14"
"ac_class": "[leather armor](/3-Mechanics/CLI/items/leather-armor.md)"
"hp": !!int "39"
"hit_dice": "6d8 + 12"
"stats":
- !!int "11"
- !!int "17"
- !!int "14"
- !!int "15"
- !!int "10"
- !!int "12"
"speed": "25 ft."
"saves":
  "Wisdom": !!int "2"
"skillsaves":
  "Stealth": !!int "5"
  "Perception": !!int "2"
"damage_resistances": "poison"
"senses": "darkvision 120 ft., truesight 30 ft., passive Perception 12"
"languages": "Dwarvish, Undercommon"
"cr": "2"
"traits":
- "desc": "The duergar has advantage on saving throws against spells and the [charmed](/3-Mechanics/CLI/rules/conditions.md#charmed),\
    \ [paralyzed](/3-Mechanics/CLI/rules/conditions.md#paralyzed), and [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)\
    \ conditions."
  "name": "Duergar Resilience"
- "desc": "While in sunlight, the duergar has disadvantage on attack rolls, as well\
    \ as on Wisdom ([Perception](/3-Mechanics/CLI/rules/skills.md#Perception)) checks\
    \ that rely on sight."
  "name": "Sunlight Sensitivity"
"actions":
- "desc": "The duergar makes two Mind-Poison Dagger attacks. It can replace one attack\
    \ with a use of Mind Mastery."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+5 (+5) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d4 + 3|text(5) (1d4 + 3) piercing damage plus dice:3d6|text(10)\
    \ (3d6) psychic damage, or 1 piercing damage plus dice:3d6|text(10) (3d6)\
    \ psychic damage while under the effect of Reduce."
  "name": "Mind-Poison Dagger"
- "desc": "The duergar magically turns [invisible](/3-Mechanics/CLI/rules/conditions.md#invisible)\
    \ for up to 1 hour or until it attacks, it forces a creature to make a saving\
    \ throw, or its [concentration](/3-Mechanics/CLI/rules/conditions.md#concentration)\
    \ is broken (as if concentrating on a spell). Any equipment the duergar wears\
    \ or carries is [invisible](/3-Mechanics/CLI/rules/conditions.md#invisible) with\
    \ it."
  "name": "Invisibility (Recharge 4-6)"
- "desc": "The duergar targets one creature it can see within 60 feet of it. The target\
    \ must succeed on a DC 12 Intelligence saving throw, or the duergar causes it\
    \ to use its reaction, if available, either to make one weapon attack against\
    \ another creature the duergar can see or to move up to 10 feet in a direction\
    \ of the duergar's choice. Creatures that can't be [charmed](/3-Mechanics/CLI/rules/conditions.md#charmed)\
    \ are immune to this effect."
  "name": "Mind Mastery"
"bonus_actions":
- "desc": "For 1 minute, the duergar magically decreases in size, along with anything\
    \ it is wearing or carrying. While reduced, the duergar is Tiny, reduces its weapon\
    \ damage to 1, and makes attack rolls, ability checks, and saving throws with\
    \ disadvantage if they use Strength. It gains a +5 bonus to all Dexterity ([Stealth](/3-Mechanics/CLI/rules/skills.md#Stealth))\
    \ checks and a +5 bonus to its AC. It can also take a bonus action on each of\
    \ its turns to take the [Hide](/3-Mechanics/CLI/rules/actions.md#Hide) action."
  "name": "Reduce (Recharges after a Short or Long Rest)"
"source":
- "MPMM"
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Duergar%20Mind%20Master.webp"
```
^statblock

```encounter-table
name: Duergar Mind Master
creatures:
 - 1: Duergar Mind Master
```

## Environment

mountain, underdark