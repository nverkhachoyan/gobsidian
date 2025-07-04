---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/vgm
- monster/cr/8
- monster/environment/arctic
- monster/environment/forest
- monster/environment/grassland
- monster/environment/hill
- monster/size/large
- monster/type/fiend/demon
aliases: ["Shoosuva"]
NoteIcon: monster
BestiaryType: fiend (demon)
SourceType: Bestiary
BookSource: Volo's Guide to Monsters p. 137
---
# [Shoosuva](3-Mechanics\CLI\bestiary\fiend/shoosuva-vgm.md)
*Source: Volo's Guide to Monsters p. 137*  

Demon lords create lesser demons for the purpose of spreading chaos and terror throughout the multiverse.

## Shoosuva

A shoosuva is a hyena-demon gifted by Yeenoghu to an especially powerful gnoll (typically as a fang of Yeenoghu). A shoosuva manifests shortly after a war band achieves a great victory, emerging from a billowing, fetid cloud of smoke as it arrives from the Abyss. In battle, the demon wraps its slavering jaws around one victim while lashing out with the poisonous stinger on its tail to bring down another one. A creature immobilized by the poison becomes easy pickings for any gnolls nearby.

Each shoosuva is bonded to a particular gnoll and fights alongside its master. A gnoll that has been gifted with a shoosuva is second only to a flind in status within a war band.

```statblock
"name": "Shoosuva (VGM)"
"size": "Large"
"type": "fiend"
"subtype": "demon"
"alignment": "Chaotic Evil"
"ac": !!int "14"
"ac_class": "natural armor"
"hp": !!int "110"
"hit_dice": "13d10 + 39"
"stats":
- !!int "18"
- !!int "13"
- !!int "17"
- !!int "7"
- !!int "14"
- !!int "9"
"speed": "40 ft."
"saves":
  "Dexterity": !!int "4"
  "Wisdom": !!int "5"
  "Constitution": !!int "6"
"damage_resistances": "cold; fire; lightning; bludgeoning, piercing, slashing from\
  \ nonmagical attacks"
"damage_immunities": "poison"
"condition_immunities": "[charmed](/3-Mechanics/CLI/rules/conditions.md#charmed),\
  \ [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened), [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)"
"senses": "darkvision 60 ft., passive Perception 12"
"languages": "Abyssal, Gnoll, telepathy 120 ft."
"cr": "8"
"traits":
- "desc": "When it reduces a creature to 0 hit points with a melee attack on its turn,\
    \ the shoosuva can take a bonus action to move up to half its speed and make a\
    \ bite attack."
  "name": "Rampage"
"actions":
- "desc": "The shoosuva makes two attacks: one with its bite and one with its tail\
    \ stinger."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+7 (+7) to hit, reach 5 ft., one target.\
    \ Hit: dice:4d10 + 4|text(26) (4d10 + 4) piercing damage."
  "name": "Bite"
- "desc": "Melee Weapon Attack: dice: d20+7 (+7) to hit, reach 15 ft., one creature.\
    \ Hit: dice:2d8 + 4|text(13) (2d8 + 4) piercing damage, and the target must\
    \ succeed on a DC 14 Constitution saving throw or become [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned).\
    \ While [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned), the target\
    \ is also [paralyzed](/3-Mechanics/CLI/rules/conditions.md#paralyzed). The target\
    \ can repeat the saving throw at the end of each of its turns, ending the effect\
    \ on itself on a success."
  "name": "Tail Stinger"
"source":
- "VGM"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/VGM/Shoosuva.webp"
```
^statblock

```encounter-table
name: Shoosuva
creatures:
 - 1: Shoosuva
```

## Environment

grassland, forest, hill, arctic