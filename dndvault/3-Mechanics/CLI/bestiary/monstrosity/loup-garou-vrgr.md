---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/vrgr
- monster/cr/13
- monster/size/medium
- monster/type/monstrosity/shapechanger
aliases: ["Loup Garou"]
NoteIcon: monster
BestiaryType: monstrosity (shapechanger)
SourceType: Bestiary
BookSource: Van Richten's Guide to Ravenloft p. 237
---
# [Loup Garou](3-Mechanics\CLI\bestiary\monstrosity/loup-garou-vrgr.md)
*Source: Van Richten's Guide to Ravenloft p. 237*  

Loup garou possess a strain of lycanthropy more virulent than that carried by common werewolves. Aside from being deadlier than their werewolf cousins, loup garou aggressively spread the plague of lycanthropy. Only through the death of a loup garou might those afflicted by it escape their curse.

## Loup Garou Lycanthropy

A Humanoid who succumbs to a loup garou's lycanthropy becomes a werewolf. This form of lycanthropy can't be removed while the loup garou that inflicted the curse lives. See the *Monster Manual* for details on lycanthropy.

Once a loup garou is slain, a [remove curse](/3-Mechanics/CLI/spells/remove-curse.md) spell cast during the night of a full moon on any afflicted werewolf it created forces the target to make a DC 17 Constitution saving throw. On a success, the curse is broken, and the target returns to its normal form and gains 3 levels of exhaustion. On a failure, the curse remains, and the target automatically fails any saving throw made to break this curse for 1 month.

```statblock
"name": "Loup Garou (VRGR)"
"size": "Medium"
"type": "monstrosity"
"subtype": "shapechanger"
"alignment": "Unaligned"
"ac": !!int "16"
"ac_class": "natural armor"
"hp": !!int "170"
"hit_dice": "20d8 + 80"
"stats":
- !!int "18"
- !!int "18"
- !!int "18"
- !!int "14"
- !!int "16"
- !!int "16"
"speed": "30 ft. (40 ft. in hybrid form, 50 ft. in dire wolf form)"
"saves":
  "Charisma": !!int "8"
  "Dexterity": !!int "9"
  "Constitution": !!int "9"
"skillsaves":
  "Stealth": !!int "9"
  "Perception": !!int "13"
"condition_immunities": "[charmed](/3-Mechanics/CLI/rules/conditions.md#charmed),\
  \ [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened)"
"senses": "darkvision 120 ft., passive Perception 23"
"languages": "Common (can't speak in wolf form)"
"cr": "13"
"traits":
- "desc": "The loup garou has advantage on attack rolls against a creature that doesn't\
    \ have all its hit points."
  "name": "Blood Frenzy"
- "desc": "When the loup garou fails a saving throw, it can choose to succeed instead."
  "name": "Legendary Resistance (2/Day)"
- "desc": "The loup garou regains 10 hit points at the start of each of its turns.\
    \ If the loup garou takes damage from a silver weapon, this trait doesn't function\
    \ at the start of the loup garou's next turn. The loup garou dies only if it starts\
    \ its turn with 0 hit points and doesn't regenerate."
  "name": "Regeneration"
"actions":
- "desc": "The loup garou makes two attacks: two with its Longsword (humanoid form)\
    \ or one with its Bite and one with its Claws (dire wolf or hybrid form)."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+9 (+9) to hit, reach 5 ft., one target.\
    \ Hit: dice:2d8 + 4|text(13) (2d8 + 4) piercing damage plus dice:4d6|text(14)\
    \ (4d6) necrotic damage. If the target is a Humanoid, it must succeed on a DC\
    \ 17 Constitution saving throw or be cursed with loup garou lycanthropy."
  "name": "Bite (Dire Wolf or Hybrid Form Only)"
- "desc": "Melee Weapon Attack: dice: d20+9 (+9) to hit, reach 5 ft., one target.\
    \ Hit: dice:2d6 + 4|text(11) (2d6 + 4) slashing damage. If the target is\
    \ a creature, it must succeed on a DC 17 Strength saving throw or be knocked [prone](/3-Mechanics/CLI/rules/conditions.md#prone)."
  "name": "Claws (Dire Wolf or Hybrid Form Only)"
- "desc": "Melee Weapon Attack: dice: d20+9 (+9) to hit, reach 5 ft., one target.\
    \ Hit: dice:2d8 + 4|text(13) (2d8 + 4) slashing damage, or dice:2d10 +\
    \ 4|text(15) (2d10 + 4) slashing damage if used with two hands."
  "name": "Longsword (Humanoid Form Only)"
"bonus_actions":
- "desc": "The loup garou polymorphs into a Large wolf-humanoid hybrid or into a Large\
    \ dire wolf, or back into its true form, which appears humanoid. Its statistics,\
    \ other than its size and speed, are the same in each form. Any equipment it is\
    \ wearing or carrying isn't transformed. It reverts to its true form if it dies."
  "name": "Change Shape"
"legendary_actions":
- "desc": "The loup garou makes one Claws attack (dire wolf or hybrid form only) or\
    \ one Longsword attack (humanoid form only)."
  "name": "Swipe"
- "desc": "The loup garou moves up to its speed without provoking opportunity attacks,\
    \ and it can make one Claws attack (dire wolf or hybrid form only) or one Longsword\
    \ attack (humanoid form only) against each creature it moves past."
  "name": "Mauling Pounce (Costs 2 Actions)"
- "desc": "The loup garou changes into hybrid or dire wolf form and then makes one\
    \ Bite attack."
  "name": "Bite (Costs 3 Actions)"
"source":
- "VRGR"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/VRGR/Loup%20Garou.webp"
```
^statblock

```encounter-table
name: Loup Garou
creatures:
 - 1: Loup Garou
```