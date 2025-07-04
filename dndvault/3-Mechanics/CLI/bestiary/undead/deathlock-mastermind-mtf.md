---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mtf
- monster/cr/8
- monster/environment/urban
- monster/size/medium
- monster/type/undead
aliases: ["Deathlock Mastermind"]
NoteIcon: monster
BestiaryType: undead
SourceType: Bestiary
BookSource: Mordenkainen's Tome of Foes p. 129, Sleeping Dragon's Wake
---
# [Deathlock Mastermind](3-Mechanics\CLI\bestiary\undead/deathlock-mastermind-mtf.md)
*Source: Mordenkainen's Tome of Foes p. 129, Sleeping Dragon's Wake*  

## Deathlock Mastermind

Though deathlocks exist to serve their patrons, they retain some freedom when it comes to devising particular tactics and carrying out their plans. Powerful deathlocks recruit lesser creatures to help them carry out their missions and, in this capacity, become the masterminds behind vast conspiracies and intrigues that culminate in the accomplishment of great acts of evil.

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
"name": "Deathlock Mastermind (MTF)"
"size": "Medium"
"type": "undead"
"alignment": "Lawful Evil"
"ac": !!int "13"
"ac_class": "16 with [mage armor](/3-Mechanics/CLI/spells/mage-armor.md)"
"hp": !!int "110"
"hit_dice": "20d8 + 20"
"stats":
- !!int "11"
- !!int "16"
- !!int "12"
- !!int "15"
- !!int "12"
- !!int "17"
"speed": "30 ft."
"saves":
  "Charisma": !!int "6"
  "Intelligence": !!int "5"
"skillsaves":
  "Perception": !!int "4"
  "History": !!int "5"
  "Arcana": !!int "5"
"damage_resistances": "necrotic; bludgeoning, piercing, slashing from nonmagical attacks\
  \ that aren't silvered"
"damage_immunities": "poison"
"condition_immunities": "[exhaustion](/3-Mechanics/CLI/rules/conditions.md#exhaustion),\
  \ [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)"
"senses": "darkvision 120 ft. (including magical darkness), passive Perception 14"
"languages": "the languages it knew in life"
"cr": "8"
"traits":
- "desc": "The deathlock's innate spellcasting ability is Charisma (spell save DC\
    \ 14). It can innately cast the following spells, requiring no material components:\n\
    \nAt will: [detect magic](/3-Mechanics/CLI/spells/detect-magic.md), [disguise\
    \ self](/3-Mechanics/CLI/spells/disguise-self.md), [mage armor](/3-Mechanics/CLI/spells/mage-armor.md)"
  "name": "Innate Spellcasting"
- "desc": "The deathlock is a 10th-level spellcaster. Its spellcasting ability is\
    \ Charisma (spell save DC 14, dice: d20+6 (+6) to hit with spell attacks).\
    \ It regains its expended spell slots when it finishes a short or long rest. It\
    \ knows the following warlock spells:\n\nCantrips (at will): [chill touch](/3-Mechanics/CLI/spells/chill-touch.md),\
    \ [mage hand](/3-Mechanics/CLI/spells/mage-hand.md), [minor illusion](/3-Mechanics/CLI/spells/minor-illusion.md),\
    \ [poison spray](/3-Mechanics/CLI/spells/poison-spray.md)\n\n1st-5th level (2\
    \ slots): [arms of Hadar](/3-Mechanics/CLI/spells/arms-of-hadar.md), [blight](/3-Mechanics/CLI/spells/blight.md),\
    \ [counterspell](/3-Mechanics/CLI/spells/counterspell.md), [crown of madness](/3-Mechanics/CLI/spells/crown-of-madness.md),\
    \ [darkness](/3-Mechanics/CLI/spells/darkness.md), [dimension door](/3-Mechanics/CLI/spells/dimension-door.md),\
    \ [dispel magic](/3-Mechanics/CLI/spells/dispel-magic.md), [fly](/3-Mechanics/CLI/spells/fly.md),\
    \ [hold monster](/3-Mechanics/CLI/spells/hold-monster.md), [invisibility](/3-Mechanics/CLI/spells/invisibility.md)"
  "name": "Spellcasting"
- "desc": "The deathlock has advantage on saving throws against any effect that turns\
    \ undead."
  "name": "Turn Resistance"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+6 (+6) to hit, reach 5 ft., one target.\
    \ Hit: dice:3d6 + 3|text(13) (3d6 + 3) necrotic damage)."
  "name": "Deathly Claw"
- "desc": "Ranged Spell Attack: dice: d20+6 (+6) to hit, range 120 ft., one\
    \ or two targets. Hit: dice:4d8|text(18) (4d8) necrotic damage. If the target\
    \ is Large or smaller, it must succeed on a DC 16 Strength saving throw or become\
    \ [restrained](/3-Mechanics/CLI/rules/conditions.md#restrained) as shadowy tendrils\
    \ wrap around it for 1 minute. A [restrained](/3-Mechanics/CLI/rules/conditions.md#restrained)\
    \ target can use its action to repeat the saving throw, ending the effect on itself\
    \ on a success."
  "name": "Grave Bolts"
"source":
- "MTF"
- "SDW"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MTF/Deathlock%20Mastermind.webp"
```
^statblock

```encounter-table
name: Deathlock Mastermind
creatures:
 - 1: Deathlock Mastermind
```

## Environment

urban