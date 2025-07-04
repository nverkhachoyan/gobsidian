---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/vrgr
- monster/cr/9
- monster/size/medium
- monster/type/undead
aliases: ["Jiangshi"]
NoteIcon: monster
BestiaryType: undead
SourceType: Bestiary
BookSource: Van Richten's Guide to Ravenloft p. 236
---
# [Jiangshi](3-Mechanics\CLI\bestiary\undead/jiangshi-vrgr.md)
*Source: Van Richten's Guide to Ravenloft p. 236*  

When a soul becomes trapped within its corpse, its bitterness can reanimate its body, creating a jiangshi. These vengeful dead stalk their descendants and the communities they knew in life, sowing terror and taking retribution for the slights or neglected burial rites that led to their cursed resurrections. Rigor mortis notoriously afflicts the limbs of jiangshi, causing them to hold their arms rigidly and to walk with a stiff gait. This, along with their flight, lead many to call them hopping vampires.

By day, jiangshi lurk within their tombs and hidden ruins to avoid the attention of the living. At night, they emerge to drain life from other creatures, these vital energies sustaining their unnatural existences and granting them greater powers. Humanoids killed by a jiangshi rise as life-hungry corpses and might turn into jiangshi themselves if they feed upon the living.

```statblock
"name": "Jiangshi (VRGR)"
"size": "Medium"
"type": "undead"
"alignment": "Unaligned"
"ac": !!int "16"
"ac_class": "natural armor"
"hp": !!int "119"
"hit_dice": "14d8 + 56"
"stats":
- !!int "18"
- !!int "3"
- !!int "18"
- !!int "17"
- !!int "14"
- !!int "12"
"speed": "20 ft."
"saves":
  "Charisma": !!int "5"
  "Wisdom": !!int "6"
  "Intelligence": !!int "7"
  "Constitution": !!int "8"
"condition_immunities": "[charmed](/3-Mechanics/CLI/rules/conditions.md#charmed),\
  \ [exhaustion](/3-Mechanics/CLI/rules/conditions.md#exhaustion), [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened),\
  \ [paralyzed](/3-Mechanics/CLI/rules/conditions.md#paralyzed), [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)"
"senses": "darkvision 120 ft., passive Perception 12"
"languages": "any languages it knew in life"
"cr": "9"
"traits":
- "desc": "The jiangshi has the following flaws:\n\nFear of Its Own Reflection.\
    \ If the jiangshi sees its own reflection, it immediately uses its reaction, if\
    \ available, to move as far away from the reflection as possible.\n\nSusceptible\
    \ to Holy Symbols. While the jiangshi is wearing or touching a holy symbol, it\
    \ automatically fails saving throws against effects that turn Undead."
  "name": "Jiangshi Weaknesses"
- "desc": "The jiangshi doesn't require air."
  "name": "Unusual Nature"
"actions":
- "desc": "The jiangshi makes three Slam attacks and uses Consume Energy."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+8 (+8) to hit, reach 5 ft., one target.\
    \ Hit: dice:2d8 + 4|text(13) (2d8 + 4) bludgeoning damage."
  "name": "Slam"
- "desc": "The jiangshi draws energy from a creature it can see within 30 feet of\
    \ it. The target makes a DC 16 Constitution saving throw, taking dice:4d8|text(18)\
    \ (4d8) necrotic damage on a failed save, or half as much damage on a successful\
    \ one. The jiangshi regains hit points equal to the amount of necrotic damage\
    \ dealt. After regaining hit points from this action, the jiangshi gains the following\
    \ benefits for 7 days: its walking speed increases to 40 feet, and it gains a\
    \ flying speed equal to its walking speed and can hover.\n\nA Humanoid slain by\
    \ this necrotic damage rises as a wight (see its entry in the Monster Manual)\
    \ at the end of the jiangshi's turn. The wight acts immediately after the jiangshi\
    \ in the initiative order. If this wight slays a Humanoid with its Life Drain,\
    \ the wight transforms into a jiangshi 5 days later."
  "name": "Consume Energy"
- "desc": "The jiangshi polymorphs into a Beast, a Humanoid, or an Undead that is\
    \ Medium or Small or back into its true form. Its statistics, other than its size,\
    \ are the same in each form. Any equipment it is wearing or carrying is absorbed\
    \ or borne by the new form (the jiangshi's choice). It reverts to its true form\
    \ if it dies."
  "name": "Change Shape"
"source":
- "VRGR"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/VRGR/Jiangshi.webp"
```
^statblock

```encounter-table
name: Jiangshi
creatures:
 - 1: Jiangshi
```