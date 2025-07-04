---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/13
- monster/environment/arctic
- monster/environment/forest
- monster/environment/hill
- monster/environment/mountain
- monster/environment/underdark
- monster/size/huge
- monster/type/giant
aliases: ["Dire Troll"]
NoteIcon: monster
BestiaryType: giant
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 246, Mordenkainen's Tome of Foes p. 243
---
# [Dire Troll](3-Mechanics\CLI\bestiary\giant/dire-troll-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 246, Mordenkainen's Tome of Foes p. 243*  

```statblock
"name": "Dire Troll (MPMM)"
"size": "Huge"
"type": "giant"
"alignment": "Typically  Chaotic Evil"
"ac": !!int "15"
"ac_class": "natural armor"
"hp": !!int "172"
"hit_dice": "15d12 + 75"
"stats":
- !!int "22"
- !!int "15"
- !!int "21"
- !!int "9"
- !!int "11"
- !!int "5"
"speed": "40 ft."
"saves":
  "Charisma": !!int "2"
  "Wisdom": !!int "5"
"skillsaves":
  "Perception": !!int "5"
"damage_resistances": "bludgeoning, piercing, slashing from nonmagical attacks"
"condition_immunities": "[frightened](/3-Mechanics/CLI/rules/conditions.md#frightened),\
  \ [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)"
"senses": "darkvision 60 ft., passive Perception 15"
"languages": "Giant"
"cr": "13"
"traits":
- "desc": "The troll regains 10 hit points at the start of its turn. If the troll\
    \ takes acid or fire damage, it regains only 5 hit points at the start of its\
    \ next turn. The troll dies only if it is hit by an attack that deals 10 or more\
    \ acid or fire damage while the troll has 0 hit points."
  "name": "Regeneration"
"actions":
- "desc": " The troll makes one Bite attack and four Claw attacks."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+11 (+11) to hit, reach 10 ft., one\
    \ target. Hit: dice:1d8 + 6|text(10) (1d8 + 6) piercing damage plus dice:1d10|text(5)\
    \ (1d10) poison damage."
  "name": "Bite"
- "desc": "Melee Weapon Attack: dice: d20+11 (+11) to hit, reach 10 ft., one\
    \ target. Hit: dice:3d6 + 6|text(16) (3d6 + 6) slashing damage."
  "name": "Claws"
- "desc": "Each creature within 10 feet of the troll must make a DC 19 Dexterity saving\
    \ throw, taking dice:8d10|text(44) (8d10) slashing damage on a failed save,\
    \ or half as much damage on a successful one."
  "name": "Whirlwind of Claws (Recharge 5-6)"
"source":
- "MPMM"
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Dire%20Troll.webp"
```
^statblock

```encounter-table
name: Dire Troll
creatures:
 - 1: Dire Troll
```

## Environment

arctic, forest, hill, mountain, underdark