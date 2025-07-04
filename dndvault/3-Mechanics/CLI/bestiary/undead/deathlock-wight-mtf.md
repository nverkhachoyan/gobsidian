---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mtf
- monster/cr/3
- monster/size/medium
- monster/type/undead
aliases: ["Deathlock Wight"]
NoteIcon: monster
BestiaryType: undead
SourceType: Bestiary
BookSource: Mordenkainen's Tome of Foes p. 129
---
# [Deathlock Wight](3-Mechanics\CLI\bestiary\undead/deathlock-wight-mtf.md)
*Source: Mordenkainen's Tome of Foes p. 129*  

## Deathlock Wight

Bereft of much of its magic, a deathlock wight lingers between the warlock it was and the deathly existence of a wight—a special punishment meted out by certain patrons and necromancers.

## Deathlock

The forging of a pact between a warlock and a patron is no minor occasion—at least not for the warlock. The consequences of breaking that pact can be dire and, in some cases, lethal. A warlock who fails to live up to a bargain with an evil patron runs the risk of rising from the dead as a deathlock, a foul undead driven to serve its otherworldly patron from beyond the grave.

An extraordinarily powerful necromancer might also discover the dark methods of creating a deathlock and then bind it to service, acting in this respect as the deathlock's patron.

## Obedient and Obsessed

An overpowering urge to serve consumes the mind of a newly awakened deathlock. All goals and ambitions it had in life that don't please its patron fall away as its master's desires become the purpose that drives the deathlock.

The creature immediately resumes work on its patron's behalf. Accomplishing a difficult goal might mean the deathlock is forced to serve another powerful creature or might entail in gathering servants of its own.

Whatever the goal, it always reflects the patron's interests, ranging from small-scale concerns to matters of cosmic scope. A deathlock in the thrall of a fiend might work to destroy a specific temple dedicated to a good god, while one that serves a Great Old One could be charged with hunting for the materials needed to call forth a horrifying entity into the world.

## Undead Nature

A deathlock doesn't require air, food, drink, or sleep.

```statblock
"name": "Deathlock Wight (MTF)"
"size": "Medium"
"type": "undead"
"alignment": "Lawful Evil"
"ac": !!int "12"
"ac_class": "15 with [mage armor](/3-Mechanics/CLI/spells/mage-armor.md)"
"hp": !!int "37"
"hit_dice": "5d8 + 15"
"stats":
- !!int "11"
- !!int "14"
- !!int "16"
- !!int "12"
- !!int "14"
- !!int "16"
"speed": "30 ft."
"saves":
  "Wisdom": !!int "4"
"skillsaves":
  "Perception": !!int "4"
  "Arcana": !!int "3"
"damage_resistances": "necrotic; bludgeoning, piercing, slashing from nonmagical attacks"
"damage_immunities": "poison"
"condition_immunities": "[exhaustion](/3-Mechanics/CLI/rules/conditions.md#exhaustion),\
  \ [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)"
"senses": "darkvision 60 ft., passive Perception 14"
"languages": "the languages it knew in life"
"cr": "3"
"traits":
- "desc": "The wight's innate spellcasting ability is Charisma (spell save DC 13).\
    \ It can innately cast the following spells, requiring no verbal or material components:\n\
    \nAt will: [detect magic](/3-Mechanics/CLI/spells/detect-magic.md), [disguise\
    \ self](/3-Mechanics/CLI/spells/disguise-self.md), [mage armor](/3-Mechanics/CLI/spells/mage-armor.md)\n\
    \n1/day each: [fear](/3-Mechanics/CLI/spells/fear.md), [hold person](/3-Mechanics/CLI/spells/hold-person.md),\
    \ [misty step](/3-Mechanics/CLI/spells/misty-step.md)"
  "name": "Innate Spellcasting"
- "desc": "While in sunlight, the wight has disadvantage on attack rolls, as well\
    \ as on Wisdom ([Perception](/3-Mechanics/CLI/rules/skills.md#Perception)) checks\
    \ that rely on sight."
  "name": "Sunlight Sensitivity"
"actions":
- "desc": "The wight attacks twice with Grave Bolt."
  "name": "Multiattack"
- "desc": "Ranged Spell Attack: dice: d20+5 (+5) to hit, range 120 ft., one\
    \ target. Hit: dice:1d8 + 3|text(7) (1d8 + 3) necrotic damage."
  "name": "Grave Bolt"
- "desc": "Melee Weapon Attack: dice: d20+4 (+4) to hit, reach 5 ft., one creature.\
    \ Hit: dice:2d6 + 2|text(9) (2d6 + 2) necrotic damage. The target must succeed\
    \ on a DC 13 Constitution saving throw or its hit point maximum is reduced by\
    \ an amount equal to the damage taken. This reduction lasts until the target finishes\
    \ a long rest. The target dies if this effect reduces its hit point maximum to\
    \ 0.\n\nA humanoid slain by this attack rises 24 hours later as a [zombie](/3-Mechanics/CLI/bestiary/undead/zombie.md)\
    \ under the wight's control, unless the humanoid is restored to life or its body\
    \ is destroyed. The wight can have no more than twelve [zombies](/3-Mechanics/CLI/bestiary/undead/zombie.md)\
    \ under its control at one time."
  "name": "Life Drain"
"source":
- "MTF"
- "TftYP"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MTF/Deathlock%20Wight.webp"
```
^statblock

```encounter-table
name: Deathlock Wight
creatures:
 - 1: Deathlock Wight
```