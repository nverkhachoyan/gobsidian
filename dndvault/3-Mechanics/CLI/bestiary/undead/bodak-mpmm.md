---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/6
- monster/environment/swamp
- monster/environment/underdark
- monster/environment/urban
- monster/size/medium
- monster/type/undead
aliases: ["Bodak"]
NoteIcon: monster
BestiaryType: undead
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 64, Volo's Guide to Monsters p. 127
---
# [Bodak](3-Mechanics\CLI\bestiary\undead/bodak-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 64, Volo's Guide to Monsters p. 127*  

```statblock
"name": "Bodak (MPMM)"
"size": "Medium"
"type": "undead"
"alignment": "Typically  Chaotic Evil"
"ac": !!int "15"
"ac_class": "natural armor"
"hp": !!int "58"
"hit_dice": "9d8 + 18"
"stats":
- !!int "15"
- !!int "16"
- !!int "15"
- !!int "7"
- !!int "12"
- !!int "12"
"speed": "30 ft."
"skillsaves":
  "Stealth": !!int "6"
  "Perception": !!int "4"
"damage_resistances": "cold; fire; bludgeoning, piercing, slashing from nonmagical\
  \ attacks"
"damage_immunities": "necrotic, poison"
"condition_immunities": "[charmed](/3-Mechanics/CLI/rules/conditions.md#charmed),\
  \ [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened), [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)"
"senses": "darkvision 120 ft., passive Perception 14"
"languages": "Abyssal, the languages it knew in life"
"cr": "6"
"traits":
- "desc": "When a creature that can see the bodak's eyes starts its turn within 30\
    \ feet of the bodak, the bodak can force it to make a DC 13 Constitution saving\
    \ throw if the bodak isn't [incapacitated](/3-Mechanics/CLI/rules/conditions.md#incapacitated)\
    \ and can see the creature. If the saving throw fails by 5 or more, the creature\
    \ is reduced to 0 hit points unless it is immune to the [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened)\
    \ condition. Otherwise, a creature takes dice:3d10|text(16) (3d10) psychic\
    \ damage on a failed save.\n\nUnless [surprised](/3-Mechanics/CLI/rules/conditions.md#surprised),\
    \ a creature can avert its eyes to avoid the saving throw at the start of its\
    \ turn. If the creature does so, it has disadvantage on attack rolls against the\
    \ bodak until the start of its next turn. If the creature looks at the bodak in\
    \ the meantime, that creature must immediately make the saving throw."
  "name": "Death Gaze"
- "desc": "The bodak takes 5 radiant damage when it starts its turn in sunlight. While\
    \ in sunlight, it has disadvantage on attack rolls and ability checks."
  "name": "Sunlight Hypersensitivity"
- "desc": "The bodak doesn't require air, food, drink, or sleep."
  "name": "Unusual Nature"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+5 (+5) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d4 + 2|text(4) (1d4 + 2) bludgeoning damage plus dice:2d8|text(9)\
    \ (2d8) necrotic damage."
  "name": "Fist"
- "desc": "One creature that the bodak can see within 60 feet of it must make a DC\
    \ 13 Constitution saving throw, taking dice:4d10|text(22) (4d10) necrotic\
    \ damage on a failed save, or half as much damage on a successful one."
  "name": "Withering Gaze"
"bonus_actions":
- "desc": "The bodak activates or deactivates this deathly aura. While active, the\
    \ aura deals 5 necrotic damage to any creature that ends its turn within 30 feet\
    \ of the bodak. Undead and Fiends ignore this effect."
  "name": "Aura of Annihilation"
"source":
- "MPMM"
- "VGM"
- "SatO"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Bodak.webp"
```
^statblock

```encounter-table
name: Bodak
creatures:
 - 1: Bodak
```

## Environment

swamp, underdark, urban