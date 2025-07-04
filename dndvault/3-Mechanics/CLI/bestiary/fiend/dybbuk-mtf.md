---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mtf
- monster/cr/4
- monster/environment/desert
- monster/environment/urban
- monster/size/medium
- monster/type/fiend/demon
aliases: ["Dybbuk"]
NoteIcon: monster
BestiaryType: fiend (demon)
SourceType: Bestiary
BookSource: Mordenkainen's Tome of Foes p. 132
---
# [Dybbuk](3-Mechanics\CLI\bestiary\fiend/dybbuk-mtf.md)
*Source: Mordenkainen's Tome of Foes p. 132*  

## Dybbuk

Dybbuks terrorize mortals on the Material Plane by possessing corpses and giving them a semblance of life, after which the demons use them to engage in a range of sordid activities.

## Puppet Masters

In their natural form, dybbuks appear as translucent flying jellyfish, trailing long tendrils as they move through the air. They rarely travel in this fashion, however. Instead, a dybbuk possesses the first suitable corpse it finds, rousing the body from death so it can then indulge its hideous vices.

## Dark Masquerade

By plundering a corpse's memories and accessing its capabilities, a dybbuk can impersonate the creature as it was in life. But the truth of the matter quickly becomes apparent to those around it, because a dybbuk can't resist pursuing its vices with a maniacal single-mindedness that betrays its true nature. Dybbuks delight in terrorizing other creatures by making their host bodies behave in horrifying ways—throwing up gouts of blood, excreting piles of squirming maggots, and contorting their limbs in impossible ways as they scuttle across the ground.

```statblock
"name": "Dybbuk (MTF)"
"size": "Medium"
"type": "fiend"
"subtype": "demon"
"alignment": "Chaotic Evil"
"ac": !!int "14"
"hp": !!int "37"
"hit_dice": "5d8 + 15"
"stats":
- !!int "6"
- !!int "19"
- !!int "16"
- !!int "16"
- !!int "15"
- !!int "14"
"speed": "0 ft., fly 40 ft. (hover)"
"skillsaves":
  "Intimidation": !!int "4"
  "Deception": !!int "6"
  "Perception": !!int "4"
"damage_resistances": "acid; cold; fire; lightning; thunder; bludgeoning, piercing,\
  \ slashing from nonmagical attacks"
"damage_immunities": "poison"
"condition_immunities": "[charmed](/3-Mechanics/CLI/rules/conditions.md#charmed),\
  \ [exhaustion](/3-Mechanics/CLI/rules/conditions.md#exhaustion), [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened),\
  \ [grappled](/3-Mechanics/CLI/rules/conditions.md#grappled), [paralyzed](/3-Mechanics/CLI/rules/conditions.md#paralyzed),\
  \ [petrified](/3-Mechanics/CLI/rules/conditions.md#petrified), [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned),\
  \ [prone](/3-Mechanics/CLI/rules/conditions.md#prone), [restrained](/3-Mechanics/CLI/rules/conditions.md#restrained)"
"senses": "darkvision 120 ft., passive Perception 14"
"languages": "Abyssal, Common, telepathy 120 ft."
"cr": "4"
"traits":
- "desc": "The dybbuk's innate spellcasting ability is Charisma (spell save DC 12).\
    \ It can innately cast the following spells, requiring no material components:\n\
    \nAt will: [dimension door](/3-Mechanics/CLI/spells/dimension-door.md)\n\n\
    3/day each: [fear](/3-Mechanics/CLI/spells/fear.md), [phantasmal force](/3-Mechanics/CLI/spells/phantasmal-force.md)"
  "name": "Innate Spellcasting"
- "desc": "The dybbuk can move through other creatures and objects as if they were\
    \ difficult terrain. It takes dice:1d10|text(5) (1d10) force damage if it\
    \ ends its turn inside an object."
  "name": "Incorporeal Movement"
- "desc": "The dybbuk has advantage on saving throws against spells and other magical\
    \ effects."
  "name": "Magic Resistance"
- "desc": "The dybbuk can use a bonus action while it is possessing a corpse to make\
    \ it do something unnatural, such as vomit blood, twist its head all the way around,\
    \ or cause a quadruped to move as a biped. Any beast or humanoid that sees this\
    \ behavior must succeed on a DC 12 Wisdom saving throw or become [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened)\
    \ for 1 minute. The [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened)\
    \ creature can repeat the saving throw at the end of each of its turns, ending\
    \ the effect on itself on a success. A creature that succeeds on a saving throw\
    \ against this ability is immune to Violate Corpse for 24 hours."
  "name": "Violate Corpse"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+6 (+6) to hit, reach 5 ft., one target.\
    \ Hit: dice:2d8 + 4|text(13) (2d8 + 4) necrotic damage. If the target is\
    \ a creature, its hit point maximum is also reduced by dice:1d6|text(3) (1d6).\
    \ This reduction lasts until the target finishes a short or long rest. The target\
    \ dies if this effect reduces its hit point maximum to 0."
  "name": "Tendril"
- "desc": "The dybbuk disappears into an intact corpse it can see within 5 feet of\
    \ it. The corpse must be Large or smaller and be that of a beast or a humanoid.\
    \ The dybbuk is now effectively the possessed creature. Its type becomes undead,\
    \ though it now looks alive, and it gains a number of temporary hit points equal\
    \ to the corpse's hit point maximum in life.\n\nWhile possessing the corpse, the\
    \ dybbuk retains its hit points, alignment, Intelligence, Wisdom, Charisma, telepathy,\
    \ and immunity to poison damage, [exhaustion](/3-Mechanics/CLI/rules/conditions.md#exhaustion),\
    \ and being [charmed](/3-Mechanics/CLI/rules/conditions.md#charmed) and [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened).\
    \ It otherwise uses the possessed target's game statistics, gaining access to\
    \ its knowledge and proficiencies but not its class features, if any.\n\nThe possession\
    \ lasts until the temporary hit points are lost (at which point the body becomes\
    \ a corpse once more) or the dybbuk ends its possession using a bonus action.\
    \ When the possession ends, the dybbuk reappears in an unoccupied space within\
    \ 5 feet of the corpse."
  "name": "Possess Corpse (Recharge 6)"
"source":
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MTF/Dybbuk.webp"
```
^statblock

```encounter-table
name: Dybbuk
creatures:
 - 1: Dybbuk
```

## Environment

desert, urban