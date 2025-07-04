---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/5
- monster/environment/urban
- monster/size/large
- monster/type/monstrosity
aliases: ["Banderhobb"]
NoteIcon: monster
BestiaryType: monstrosity
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 56, Volo's Guide to Monsters p. 122
---
# [Banderhobb](3-Mechanics\CLI\bestiary\monstrosity/banderhobb-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 56, Volo's Guide to Monsters p. 122*  

```statblock
"name": "Banderhobb (MPMM)"
"size": "Large"
"type": "monstrosity"
"alignment": "Typically  Neutral Evil"
"ac": !!int "15"
"ac_class": "natural armor"
"hp": !!int "84"
"hit_dice": "8d10 + 40"
"stats":
- !!int "20"
- !!int "12"
- !!int "20"
- !!int "11"
- !!int "14"
- !!int "8"
"speed": "30 ft."
"skillsaves":
  "Athletics": !!int "8"
  "Stealth": !!int "7"
"condition_immunities": "[charmed](/3-Mechanics/CLI/rules/conditions.md#charmed),\
  \ [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened)"
"senses": "darkvision 120 ft., passive Perception 12"
"languages": "understands Common and the languages of its creator but can't speak"
"cr": "5"
"traits":
- "desc": "If the banderhobb has even a tiny piece of a creature or an object in its\
    \ possession, such as a lock of hair or a splinter of wood, it knows the most\
    \ direct route to that creature or object if it is within 1 mile of the banderhobb."
  "name": "Resonant Connection"
"actions":
- "desc": "The banderhobb makes one Bite attack and one Tongue attack. It can replace\
    \ one attack with a use of Shadow Step."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+8 (+8) to hit, reach 5 ft., one target.\
    \ Hit: dice:3d6 + 5|text(15) (3d6 + 5) piercing damage, and the target is\
    \ [grappled](/3-Mechanics/CLI/rules/conditions.md#grappled) (escape DC 16) if\
    \ it is a Large or smaller creature. Until this grapple ends, the target is [restrained](/3-Mechanics/CLI/rules/conditions.md#restrained),\
    \ and the banderhobb can't use its Bite attack or Tongue attack on another target."
  "name": "Bite"
- "desc": "Melee Weapon Attack: dice: d20+8 (+8) to hit, reach 15 ft., one creature.\
    \ Hit: dice:3d6|text(10) (3d6) necrotic damage, and the target must make\
    \ a DC 16 Strength saving throw. On a failed save, the target is pulled to a space\
    \ within 5 feet of the banderhobb."
  "name": "Tongue"
- "desc": "The banderhobb teleports up to 30 feet to an unoccupied space of dim light\
    \ or darkness that it can see."
  "name": "Shadow Step"
- "desc": "Melee Weapon Attack: dice: d20+8 (+8) to hit, reach 5 ft., one Medium\
    \ or smaller creature [grappled](/3-Mechanics/CLI/rules/conditions.md#grappled)\
    \ by the banderhobb. Hit: dice:3d6 + 5|text(15) (3d6 + 5) piercing damage.\
    \ The creature is also swallowed, and the grapple ends. The swallowed creature\
    \ is [blinded](/3-Mechanics/CLI/rules/conditions.md#blinded) and [restrained](/3-Mechanics/CLI/rules/conditions.md#restrained),\
    \ it has total cover against attacks and other effects outside the banderhobb,\
    \ and it takes dice:3d6|text(10) (3d6) necrotic damage at the start of each\
    \ of the banderhobb's turns. A creature reduced to 0 hit points in this way stops\
    \ taking the necrotic damage and becomes stable.\n\nThe banderhobb can have only\
    \ one creature swallowed at a time. While the banderhobb isn't [incapacitated](/3-Mechanics/CLI/rules/conditions.md#incapacitated),\
    \ it can regurgitate the creature at any time (no action required) in a space\
    \ within 5 feet of it. The creature exits [prone](/3-Mechanics/CLI/rules/conditions.md#prone).\
    \ If the banderhobb dies, it likewise regurgitates a swallowed creature."
  "name": "Swallow"
"bonus_actions":
- "desc": "While in dim light or darkness, the banderhobb takes the [Hide](/3-Mechanics/CLI/rules/actions.md#Hide)\
    \ action."
  "name": "Shadow Stealth"
"source":
- "MPMM"
- "VGM"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Banderhobb.webp"
```
^statblock

```encounter-table
name: Banderhobb
creatures:
 - 1: Banderhobb
```

## Environment

urban