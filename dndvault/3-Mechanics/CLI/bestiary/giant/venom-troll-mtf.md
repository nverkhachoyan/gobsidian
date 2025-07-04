---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mtf
- monster/cr/7
- monster/environment/forest
- monster/environment/swamp
- monster/environment/underdark
- monster/size/large
- monster/type/giant
aliases: ["Venom Troll"]
NoteIcon: monster
BestiaryType: giant
SourceType: Bestiary
BookSource: Mordenkainen's Tome of Foes p. 245
---
# [Venom Troll](3-Mechanics\CLI\bestiary\giant/venom-troll-mtf.md)
*Source: Mordenkainen's Tome of Foes p. 245*  

## Venom Troll

A troll ravaged by massive doses of poison might mutate into a venom troll. Lingering poison infuses its blood and tissue, and poison leaks from its pores to coat its fangs and claws. These creatures are especially dangerous in close combat, because poison drips off their flesh and sprays out from every wound they receive.

## Trolls

Trolls that are nearly obliterated but survive and regenerate from mere scraps of flesh can display bizarre mutations. One of these warped trolls is especially likely to arise if the creature regenerates in the presence of magical emanations, planar energy, disease, or death on a vast scale, or if its body was damaged by elemental forces. These mutated forms can also be produced and shaped by the ritual magic of evil spellcasters.

```statblock
"name": "Venom Troll (MTF)"
"size": "Large"
"type": "giant"
"alignment": "Chaotic Evil"
"ac": !!int "15"
"ac_class": "natural armor"
"hp": !!int "94"
"hit_dice": "9d10 + 45"
"stats":
- !!int "18"
- !!int "13"
- !!int "20"
- !!int "7"
- !!int "9"
- !!int "7"
"speed": "30 ft."
"skillsaves":
  "Perception": !!int "2"
"damage_immunities": "poison"
"condition_immunities": "[poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)"
"senses": "darkvision 60 ft., passive Perception 12"
"languages": "Giant"
"cr": "7"
"traits":
- "desc": "The troll has advantage on Wisdom ([Perception](/3-Mechanics/CLI/rules/skills.md#Perception))\
    \ checks that rely on smell."
  "name": "Keen Smell"
- "desc": "When the troll takes damage of any type but psychic, each creature within\
    \ 5 feet of the troll takes dice:2d8|text(9) (2d8) poison damage."
  "name": "Poison Splash"
- "desc": "The troll regains 10 hit points at the start of each of its turns. If the\
    \ troll takes acid or fire damage, this trait doesn't function at the start of\
    \ the troll's next turn. The troll dies only if it starts its turn with 0 hit\
    \ points and doesn't regenerate."
  "name": "Regeneration"
"actions":
- "desc": "The troll makes three attacks: one with its bite and two with its claws."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+7 (+7) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d6 + 4|text(7) (1d6 + 4) piercing damage plus dice:1d8|text(4)\
    \ (1d8) poison damage, and the creature is [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)\
    \ until the start of the troll's next turn."
  "name": "Bite"
- "desc": "Melee Weapon Attack: dice: d20+7 (+7) to hit, reach 5 ft., one target.\
    \ Hit: dice:2d6 + 4|text(11) (2d6 + 4) slashing damage plus dice:1d8|text(4)\
    \ (1d8) poison damage."
  "name": "Claws"
- "desc": "The troll slices itself with a claw, releasing a spray of poison in a 15-foot\
    \ cube. The troll takes dice:2d6|text(7) (2d6) slashing damage (this damage\
    \ can't be reduced in any way). Each creature in the area must make a DC 16 Constitution\
    \ saving throw. On a failed save, a creature takes dice:4d8|text(18) (4d8)\
    \ poison damage and is [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)\
    \ for 1 minute. On a successful save, the creature takes half as much damage and\
    \ isn't [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned). A [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)\
    \ creature can repeat the saving throw at the end of each of its turns, ending\
    \ the effect on itself on a success."
  "name": "Venom Spray (Recharge 6)"
"source":
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MTF/Venom%20Troll.webp"
```
^statblock

```encounter-table
name: Venom Troll
creatures:
 - 1: Venom Troll
```

## Environment

forest, swamp, underdark