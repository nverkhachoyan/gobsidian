---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mtf
- monster/cr/4
- monster/environment/urban
- monster/size/medium
- monster/type/undead
aliases: ["Deathlock"]
NoteIcon: monster
BestiaryType: undead
SourceType: Bestiary
BookSource: Mordenkainen's Tome of Foes p. 128
---
# [Deathlock](3-Mechanics\CLI\bestiary\undead/deathlock-mtf.md)
*Source: Mordenkainen's Tome of Foes p. 128*  

## Deathlock

The forging of a pact between a warlock and a patron is no minor occasion—at least not for the warlock. The consequences of breaking that pact can be dire and, in some cases, lethal. A warlock who fails to live up to a bargain with an evil patron runs the risk of rising from the dead as a deathlock, a foul undead driven to serve its otherworldly patron from beyond the grave.

An extraordinarily powerful necromancer might also discover the dark methods of creating a deathlock and then bind it to service, acting in this respect as the deathlock's patron.

## Obedient and Obsessed

An overpowering urge to serve consumes the mind of a newly awakened deathlock. All goals and ambitions it had in life that don't please its patron fall away as its master's desires become the purpose that drives the deathlock. The creature immediately resumes work on its patron's behalf. Accomplishing a difficult goal might mean the deathlock is forced to serve another powerful creature or might entail in gathering servants of its own.

Whatever the goal, it always reflects the patron's interests, ranging from small-scale concerns to matters of cosmic scope. A deathlock in the thrall of a fiend might work to destroy a specific temple dedicated to a good god, while one that serves a Great Old One could be charged with hunting for the materials needed to call forth a horrifying entity into the world.

## Undead Nature

A deathlock doesn't require air, food, drink, or sleep.

```statblock
"name": "Deathlock (MTF)"
"size": "Medium"
"type": "undead"
"alignment": "Lawful Evil"
"ac": !!int "12"
"ac_class": "15 with [mage armor](/3-Mechanics/CLI/spells/mage-armor.md)"
"hp": !!int "36"
"hit_dice": "8d8"
"stats":
- !!int "11"
- !!int "15"
- !!int "10"
- !!int "14"
- !!int "12"
- !!int "16"
"speed": "30 ft."
"saves":
  "Charisma": !!int "5"
  "Intelligence": !!int "4"
"skillsaves":
  "History": !!int "4"
  "Arcana": !!int "4"
"damage_resistances": "necrotic; bludgeoning, piercing, slashing from nonmagical attacks\
  \ that aren't silvered"
"damage_immunities": "poison"
"condition_immunities": "[exhaustion](/3-Mechanics/CLI/rules/conditions.md#exhaustion),\
  \ [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)"
"senses": "darkvision 60 ft., passive Perception 11"
"languages": "the languages it knew in life"
"cr": "4"
"traits":
- "desc": "The deathlock's innate spellcasting ability is Charisma (spell save DC\
    \ 13). It can innately cast the following spells, requiring no material components:\n\
    \nAt will: [detect magic](/3-Mechanics/CLI/spells/detect-magic.md), [disguise\
    \ self](/3-Mechanics/CLI/spells/disguise-self.md), [mage armor](/3-Mechanics/CLI/spells/mage-armor.md)"
  "name": "Innate Spellcasting"
- "desc": "The deathlock is a 5th-level spellcaster. Its spellcasting ability is Charisma\
    \ (spell save DC 13, dice: d20+5 (+5) to hit with spell attacks). It regains\
    \ its expended spell slots when it finishes a short or long rest. It knows the\
    \ following warlock spells:\n\nCantrips (at will): [chill touch](/3-Mechanics/CLI/spells/chill-touch.md),\
    \ [eldritch blast](/3-Mechanics/CLI/spells/eldritch-blast.md), [mage hand](/3-Mechanics/CLI/spells/mage-hand.md)\n\
    \n1st-3rd level (2 slots): [arms of Hadar](/3-Mechanics/CLI/spells/arms-of-hadar.md),\
    \ [dispel magic](/3-Mechanics/CLI/spells/dispel-magic.md), [hold person](/3-Mechanics/CLI/spells/hold-person.md),\
    \ [hunger of Hadar](/3-Mechanics/CLI/spells/hunger-of-hadar.md), [invisibility](/3-Mechanics/CLI/spells/invisibility.md),\
    \ [spider climb](/3-Mechanics/CLI/spells/spider-climb.md)"
  "name": "Spellcasting"
- "desc": "The deathlock has advantage on saving throws against any effect that turns\
    \ undead."
  "name": "Turn Resistance"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+4 (+4) to hit, reach 5 ft., one target.\
    \ Hit: dice:2d6 + 2|text(9) (2d6 + 2) necrotic damage."
  "name": "Deathly Claw"
"source":
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MTF/Deathlock.webp"
```
^statblock

```encounter-table
name: Deathlock
creatures:
 - 1: Deathlock
```

## Environment

urban