---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mtf
- monster/cr/2
- monster/environment/mountain
- monster/environment/underdark
- monster/size/medium
- monster/type/humanoid/dwarf
aliases: ["Duergar Mind Master"]
NoteIcon: monster
BestiaryType: humanoid (dwarf)
SourceType: Bestiary
BookSource: Mordenkainen's Tome of Foes p. 189
---
# [Duergar Mind Master](3-Mechanics\CLI\bestiary\humanoid/duergar-mind-master-mtf.md)
*Source: Mordenkainen's Tome of Foes p. 189*  

> [!quote]-  
> 
> Absence of light results in darkness. So, what results from the absence of Good? The underdark is a place with very little of either.

## Duergar Mind Master

The feared duergar mind masters usually operate as spies, both inside and beyond a duergar stronghold. Their psionically augmented abilities enable them to see through illusions with ease and shrink down to miniature size to spy on their targets.

> [!quote]- A quote from Mordenkainen  
> 
> You think dwarves are dour and unpleasant? Wait until you meet a duergar.

The cruel duergar plot not only to defeat other dwarves, but to cast down the entire dwarven pantheon in revenge for their being abandoned and left to be enslaved by mind flayers. To this end, duergar train warriors to fulfill a variety of roles.

```statblock
"name": "Duergar Mind Master (MTF)"
"size": "Medium"
"type": "humanoid"
"subtype": "dwarf"
"alignment": "Lawful Evil"
"ac": !!int "14"
"ac_class": "[leather armor](/3-Mechanics/CLI/items/leather-armor.md); 19 while reduced"
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
- "desc": "The duergar has advantage on saving throws against poison, spells, and\
    \ illusions, as well as to resist being [charmed](/3-Mechanics/CLI/rules/conditions.md#charmed)\
    \ or [paralyzed](/3-Mechanics/CLI/rules/conditions.md#paralyzed)."
  "name": "Duergar Resilience"
- "desc": "While in sunlight, the duergar has disadvantage on attack rolls, as well\
    \ as on Wisdom ([Perception](/3-Mechanics/CLI/rules/skills.md#Perception)) checks\
    \ that rely on sight."
  "name": "Sunlight Sensitivity"
"actions":
- "desc": "The duergar makes two melee attacks. It can replace one of those attacks\
    \ with a use of Mind Mastery."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+5 (+5) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d4 + 3|text(5) (1d4 + 3) piercing damage and dice:3d6|text(10)\
    \ (3d6) psychic damage, or 1 piercing damage and dice:4d6|text(14) (4d6)\
    \ psychic damage while reduced."
  "name": "Mind-Poison Dagger"
- "desc": "The duergar magically turns [invisible](/3-Mechanics/CLI/rules/conditions.md#invisible)\
    \ for up to 1 hour or until it attacks, it casts a spell, it uses its Reduce,\
    \ or its [concentration](/3-Mechanics/CLI/rules/conditions.md#concentration) is\
    \ broken (as if concentrating on a spell). Any equipment the duergar wears or\
    \ carries is [invisible](/3-Mechanics/CLI/rules/conditions.md#invisible) with\
    \ it."
  "name": "Invisibility (Recharge 4-6)"
- "desc": "The duergar targets one creature it can see within 60 feet of it. The target\
    \ must succeed on a DC 12 Intelligence saving throw, or the duergar causes it\
    \ to use its reaction either to make one weapon attack against another creature\
    \ the duergar can see or to move up to 10 feet in a direction of the duergar's\
    \ choice. Creatures that can't be [charmed](/3-Mechanics/CLI/rules/conditions.md#charmed)\
    \ are immune to this effect."
  "name": "Mind Mastery"
- "desc": "For 1 minute, the duergar magically decreases in size, along with anything\
    \ it is wearing or carrying. While reduced, the duergar is Tiny, reduces its weapon\
    \ damage to 1, and makes attacks, checks, and saving throws with disadvantage\
    \ if they use Strength. It gains a +5 bonus to all Dexterity ([Stealth](/3-Mechanics/CLI/rules/skills.md#Stealth))\
    \ checks and a +5 bonus to its AC. It can also take a bonus action on each of\
    \ its turns to take the Hide action."
  "name": "Reduce (Recharges after a Short or Long Rest)"
"source":
- "MTF"
- "IDRotF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MTF/Duergar%20Mind%20Master.webp"
```
^statblock

```encounter-table
name: Duergar Mind Master
creatures:
 - 1: Duergar Mind Master
```

## Environment

mountain, underdark