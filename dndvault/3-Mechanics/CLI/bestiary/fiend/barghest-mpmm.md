---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/4
- monster/environment/forest
- monster/environment/grassland
- monster/environment/hill
- monster/environment/mountain
- monster/environment/underdark
- monster/size/large
- monster/type/fiend
aliases: ["Barghest"]
NoteIcon: monster
BestiaryType: fiend
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 60, Volo's Guide to Monsters p. 123
---
# [Barghest](3-Mechanics\CLI\bestiary\fiend/barghest-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 60, Volo's Guide to Monsters p. 123*  

```statblock
"name": "Barghest (MPMM)"
"size": "Large"
"type": "fiend"
"alignment": "Typically  Neutral Evil"
"ac": !!int "17"
"ac_class": "natural armor"
"hp": !!int "60"
"hit_dice": "8d10 + 16"
"stats":
- !!int "19"
- !!int "15"
- !!int "14"
- !!int "13"
- !!int "12"
- !!int "14"
"speed": "60 ft. (30 ft. in goblin form)"
"skillsaves":
  "Intimidation": !!int "4"
  "Deception": !!int "4"
  "Stealth": !!int "4"
  "Perception": !!int "5"
"damage_resistances": "cold; lightning; bludgeoning, piercing, slashing from nonmagical\
  \ attacks"
"damage_immunities": "acid, poison"
"condition_immunities": "[poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)"
"senses": "blindsight 60 ft., darkvision 60 ft., passive Perception 15"
"languages": "Abyssal, Common, Goblin, Infernal, telepathy 60 ft."
"cr": "4"
"traits":
- "desc": "The barghest casts one of the following spells, requiring no material components\
    \ and using Charisma as the spellcasting ability (spell save DC 12):\n\nAt will:\
    \ [levitate](/3-Mechanics/CLI/spells/levitate.md), [minor illusion](/3-Mechanics/CLI/spells/minor-illusion.md),\
    \ [pass without trace](/3-Mechanics/CLI/spells/pass-without-trace.md)\n\n1/day\
    \ each: [charm person](/3-Mechanics/CLI/spells/charm-person.md), [dimension\
    \ door](/3-Mechanics/CLI/spells/dimension-door.md), [suggestion](/3-Mechanics/CLI/spells/suggestion.md)"
  "name": "Spellcasting"
- "desc": "When the barghest starts its turn engulfed in flames that are at least\
    \ 10 feet high or wide, it must succeed on a DC 15 Charisma saving throw or be\
    \ instantly banished to Gehenna"
  "name": "Fire Banishment"
- "desc": "The barghest can feed on the corpse of a Fey or Humanoid it killed within\
    \ the past 10 minutes. This feeding takes at least 1 minute, and it destroys the\
    \ corpse. The victim's soul is trapped in the barghest for 24 hours, after which\
    \ time it is digested and the person is incapable of being revived. If the barghest\
    \ dies before the soul is digested, the soul is released. While a soul is trapped\
    \ in the barghest, any magic that tries to restore the soul to life has a 50%\
    \ chance chance of failing and being wasted."
  "name": "Soul Feeding"
"actions":
- "desc": "The barghest makes one Bite attack and one Claw attack."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+6 (+6) to hit, reach 5 ft., one target.\
    \ Hit: dice:2d8 + 4|text(13) (2d8 + 4) piercing damage."
  "name": "Bite"
- "desc": "Melee Weapon Attack: dice: d20+6 (+6) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d8 + 4|text(8) (1d8 + 4) slashing damage."
  "name": "Claw"
"bonus_actions":
- "desc": "The barghest transforms into a Small goblin or back into its true form.\
    \ Other than its size and speed, its statistics are the same in each form. Any\
    \ equipment it is wearing or carrying isn't transformed. The barghest reverts\
    \ to its true form if it dies."
  "name": "Change Shape"
"source":
- "MPMM"
- "VGM"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Barghest.webp"
```
^statblock

```encounter-table
name: Barghest
creatures:
 - 1: Barghest
```

## Environment

forest, grassland, hill, mountain, underdark