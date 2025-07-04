---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mtf
- monster/cr/21
- monster/size/huge
- monster/type/fiend/demon
aliases: ["Molydeus"]
NoteIcon: monster
BestiaryType: fiend (demon)
SourceType: Bestiary
BookSource: Mordenkainen's Tome of Foes p. 134
---
# [Molydeus](3-Mechanics\CLI\bestiary\fiend/molydeus-mtf.md)
*Source: Mordenkainen's Tome of Foes p. 134*  

## Molydeus

The most ruthless and dangerous of demons—more feared than the dreaded balor—the molydeus speaks with the authority of the demon lord it serves as it enforces its master's will. Standing some 12 feet tall, a molydeus has a red-skinned, humanoid body and two heads—one that of a slavering wolf and the other that of a serpent with dripping fangs perched atop a long neck.

Molydei might guard their masters' possessions, roam the battlefields of the Abyss to ensure the loyalty of troops, or bring swift death to their enemies.

## Branded and Bound

When a demon earns the attention of a demon lord through ferocity, cunning, or an act of surprising devotion, the demon lord might reward such service by snatching up the fiend and subjecting it to excruciating torments to remake it into a molydeus.

## Voice of the Master

A demon lord has a direct link to its molydeus and uses the serpent head to communicate its wishes. A molydeus is, therefore, said to utter its master's will, commanding other demons to carry out orders and using violence to ensure they obey. A molydeus must constantly be ready for the scrutiny of its master, for the demon lord can decide at any moment to observe the molydeus through the serpent. Thus, there is no room for treachery in a molydeus.

## Special Weapon

As part of a demon lord's trust in its molydeus, it bestows a powerful weapon upon the guardian demon. The demon lord fashions the weapon from a portion of the fiend's essence, so that the demon and its weapon are forever bound. If the molydeus dies, the weapon dissolves into a pool of foul-smelling slime. It's possible to steal such a weapon, but a molydeus deprived of its weapon will stop at nothing to regain it.

The weapon a molydeus wields reflects the nature of its master. Those that serve Baphomet carry a glaive; Demogorgon, a whip; Fraz-Urb'luu, a battleaxe; Graz'zt, a greatsword; Orcus, a morningstar; and Yeenoghu, a flail. The weapon's form doesn't affect its capabilities.

## Dark Guardians

One of the chief tasks of any molydeus is to help protect its master's amulet—the most prized possession of any demon lord. Each of these dangerous relics allows a demon lord to return to life in the Abyss if the unthinkable occurs and its abyssal form is destroyed. As useful as these amulets are, they are also liabilities—because, armed with an amulet, a creature can coerce the demon lord to which it belongs into doing its bidding, or can strand it in the Abyss if the amulet is destroyed.

```statblock
"name": "Molydeus (MTF)"
"size": "Huge"
"type": "fiend"
"subtype": "demon"
"alignment": "Chaotic Evil"
"ac": !!int "19"
"ac_class": "natural armor"
"hp": !!int "216"
"hit_dice": "16d12 + 112"
"stats":
- !!int "28"
- !!int "22"
- !!int "25"
- !!int "21"
- !!int "24"
- !!int "24"
"speed": "40 ft."
"saves":
  "Charisma": !!int "14"
  "Wisdom": !!int "14"
  "Strength": !!int "16"
  "Constitution": !!int "14"
"skillsaves":
  "Perception": !!int "21"
"damage_resistances": "cold; fire; lightning; bludgeoning, piercing, slashing from\
  \ nonmagical attacks"
"damage_immunities": "poison"
"condition_immunities": "[blinded](/3-Mechanics/CLI/rules/conditions.md#blinded),\
  \ [charmed](/3-Mechanics/CLI/rules/conditions.md#charmed), [deafened](/3-Mechanics/CLI/rules/conditions.md#deafened),\
  \ [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened), [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned),\
  \ [stunned](/3-Mechanics/CLI/rules/conditions.md#stunned)"
"senses": "truesight 120 ft., passive Perception 31"
"languages": "Abyssal, telepathy 120 ft."
"cr": "21"
"traits":
- "desc": "The molydeus's innate spellcasting ability is Charisma (spell save DC 22).\
    \ It can innately cast the following spells, requiring no material components:\n\
    \nAt will: [dispel magic](/3-Mechanics/CLI/spells/dispel-magic.md), [polymorph](/3-Mechanics/CLI/spells/polymorph.md),\
    \ [telekinesis](/3-Mechanics/CLI/spells/telekinesis.md), [teleport](/3-Mechanics/CLI/spells/teleport.md)\n\
    \n1/day: [imprisonment](/3-Mechanics/CLI/spells/imprisonment.md)\n\n3/day:\
    \ [lightning bolt](/3-Mechanics/CLI/spells/lightning-bolt.md)"
  "name": "Innate Spellcasting"
- "desc": "If the molydeus fails a saving throw, it can choose to succeed instead."
  "name": "Legendary Resistance (3/Day)"
- "desc": "The molydeus has advantage on saving throws against spells and other magical\
    \ effects."
  "name": "Magic Resistance"
- "desc": "The molydeus's weapon attacks are magical."
  "name": "Magic Weapons"
"actions":
- "desc": "The molydeus makes three attacks: one with its weapon, one with its wolf\
    \ bite, and one with its snakebite."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+16 (+16) to hit, reach 15 ft., one\
    \ target. Hit: dice:2d10 + 9|text(20) (2d10 + 9) slashing damage. If the\
    \ target has at least one head and the molydeus rolled a 20 on the attack roll,\
    \ the target is decapitated and dies if it can't survive without that head. A\
    \ target is immune to this effect if it takes none of the damage, has legendary\
    \ actions, or is Huge or larger. Such a creature takes an extra dice: 6d8|avg|noform\
    \ (6d8) slashing damage from the hit."
  "name": "Demonic Weapon"
- "desc": "Melee Weapon Attack: dice: d20+16 (+16) to hit, reach 10 ft., one\
    \ target. Hit: dice:2d6 + 9|text(16) (2d6 + 9) piercing damage."
  "name": "Wolf Bite"
- "desc": "Melee Weapon Attack: dice: d20+16 (+16) to hit, reach 15 ft., one\
    \ creature. Hit: dice:1d6 + 9|text(12) (1d6 + 9) piercing damage, and the\
    \ target must succeed on a DC 22 Constitution saving throw or its hit point maximum\
    \ is reduced by an amount equal to the damage taken. This reduction lasts until\
    \ the target finishes a long rest. The target transforms into a [manes](/3-Mechanics/CLI/bestiary/fiend/manes.md)\
    \ if this reduces its hit point maximum to 0. This transformation can be ended\
    \ only by a [wish](/3-Mechanics/CLI/spells/wish.md) spell."
  "name": "Snakebite"
"legendary_actions":
- "desc": "The molydeus makes one attack, either with its demonic weapon or with its\
    \ snakebite."
  "name": "Attack"
- "desc": "The molydeus moves without provoking opportunity attacks."
  "name": "Move"
- "desc": "The molydeus casts one spell from its Innate Spellcasting trait."
  "name": "Cast a Spell"
"source":
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MTF/Molydeus.webp"
```
^statblock

```encounter-table
name: Molydeus
creatures:
 - 1: Molydeus
```