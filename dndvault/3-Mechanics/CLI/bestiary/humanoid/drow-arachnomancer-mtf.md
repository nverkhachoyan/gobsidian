---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mtf
- monster/cr/13
- monster/environment/underdark
- monster/size/medium
- monster/type/humanoid/elf
aliases: ["Drow Arachnomancer"]
NoteIcon: monster
BestiaryType: humanoid (elf)
SourceType: Bestiary
BookSource: Mordenkainen's Tome of Foes p. 182
---
# [Drow Arachnomancer](3-Mechanics\CLI\bestiary\humanoid/drow-arachnomancer-mtf.md)
*Source: Mordenkainen's Tome of Foes p. 182*  

> [!quote]- A quote from Mordenkainen  
> 
> Drow mages can be quite learned and skilled in magic. Some of them can even cast my spells.

## Drow Arachnomancer

Drow spellcasters who seek to devote themselves wholly to the Spider Queen sometimes walk the dark path of the arachnomancer. By offering up body and soul to Lolth, they gain tremendous power and a supernatural connection to the ancient spiders of the Demonweb Pits, channeling magic from that dread place.

```statblock
"name": "Drow Arachnomancer (MTF)"
"size": "Medium"
"type": "humanoid"
"subtype": "elf"
"alignment": "Chaotic Evil"
"ac": !!int "15"
"ac_class": "[studded leather](/3-Mechanics/CLI/items/studded-leather-armor.md)"
"hp": !!int "162"
"hit_dice": "25d8 + 50"
"stats":
- !!int "11"
- !!int "17"
- !!int "14"
- !!int "19"
- !!int "14"
- !!int "16"
"speed": "30 ft., climb 30 ft."
"saves":
  "Charisma": !!int "8"
  "Intelligence": !!int "9"
  "Constitution": !!int "7"
"skillsaves":
  "Nature": !!int "9"
  "Stealth": !!int "8"
  "Perception": !!int "7"
  "Arcana": !!int "9"
"damage_resistances": "poison"
"senses": "blindsight 10 ft., darkvision 120 ft., passive Perception 17"
"languages": "Elvish, Undercommon, can speak with spiders"
"cr": "13"
"traits":
- "desc": "The drow's innate spellcasting ability is Charisma (spell save DC 16).\
    \ It can innately cast the following spells, requiring no material components:\n\
    \nAt will: [dancing lights](/3-Mechanics/CLI/spells/dancing-lights.md)\n\n\
    1/day each: [darkness](/3-Mechanics/CLI/spells/darkness.md), [faerie fire](/3-Mechanics/CLI/spells/faerie-fire.md),\
    \ [levitate](/3-Mechanics/CLI/spells/levitate.md) (self only)"
  "name": "Innate Spellcasting"
- "desc": "The drow is a 16th-level spellcaster. Its spellcasting ability is Charisma\
    \ (spell save DC 16, dice: d20+8 (+8) to hit with spell attacks). It regains\
    \ its expended spell slots when it finishes a short or long rest. It knows the\
    \ following warlock spells:\n\n1/day: [eyebite](/3-Mechanics/CLI/spells/eyebite.md),\
    \ [etherealness](/3-Mechanics/CLI/spells/etherealness.md), [dominate monster](/3-Mechanics/CLI/spells/dominate-monster.md)\n\
    \nCantrips (at will): [chill touch](/3-Mechanics/CLI/spells/chill-touch.md),\
    \ [eldritch blast](/3-Mechanics/CLI/spells/eldritch-blast.md), [mage hand](/3-Mechanics/CLI/spells/mage-hand.md),\
    \ [poison spray](/3-Mechanics/CLI/spells/poison-spray.md)\n\n1st-5th level (3\
    \ slots): [conjure animals](/3-Mechanics/CLI/spells/conjure-animals.md) (spiders\
    \ only), [crown of madness](/3-Mechanics/CLI/spells/crown-of-madness.md), [dimension\
    \ door](/3-Mechanics/CLI/spells/dimension-door.md), [dispel magic](/3-Mechanics/CLI/spells/dispel-magic.md),\
    \ [fear](/3-Mechanics/CLI/spells/fear.md), [fly](/3-Mechanics/CLI/spells/fly.md),\
    \ [giant insect](/3-Mechanics/CLI/spells/giant-insect.md), [hold monster](/3-Mechanics/CLI/spells/hold-monster.md),\
    \ [insect plague](/3-Mechanics/CLI/spells/insect-plague.md), [invisibility](/3-Mechanics/CLI/spells/invisibility.md),\
    \ [vampiric touch](/3-Mechanics/CLI/spells/vampiric-touch.md), [web](/3-Mechanics/CLI/spells/web.md),\
    \ [witch bolt](/3-Mechanics/CLI/spells/witch-bolt.md)"
  "name": "Spellcasting"
- "desc": "The drow can use a bonus action to magically polymorph into a [giant spider](/3-Mechanics/CLI/bestiary/beast/giant-spider.md),\
    \ remaining in that form for up to 1 hour. It can revert to its true form as a\
    \ bonus action. Its statistics, other than its size, are the same in each form.\
    \ It can speak and cast spells while in giant spider form. Any equipment it is\
    \ wearing or carrying in humanoid form melds into the giant spider form. It can't\
    \ activate, use, wield, or otherwise benefit from any of its equipment. It reverts\
    \ to its humanoid form if it dies."
  "name": "Change Shape (Recharges after a Short or Long Rest)"
- "desc": "The drow has advantage on saving throws against being [charmed](/3-Mechanics/CLI/rules/conditions.md#charmed),\
    \ and magic can't put the drow to sleep."
  "name": "Fey Ancestry"
- "desc": "The drow can climb difficult surfaces, including upside down on ceilings,\
    \ without needing to make an ability check."
  "name": "Spider Climb"
- "desc": "While in sunlight, the drow has disadvantage on attack rolls, as well as\
    \ on Wisdom ([Perception](/3-Mechanics/CLI/rules/skills.md#Perception)) checks\
    \ that rely on sight."
  "name": "Sunlight Sensitivity"
- "desc": "The drow ignores movement restrictions caused by webbing."
  "name": "Web Walker"
"actions":
- "desc": "The drow makes two poisonous touch attacks or two bite attacks. The first\
    \ of these attacks that hits each round deals an extra dice:4d12|text(26) (4d12)\
    \ poison damage to the target."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+8 (+8) to hit, reach 5 ft., one target.\
    \ Hit: dice:8d6|text(28) (8d6) poison damage."
  "name": "Poisonous Touch (Humanoid Form Only)"
- "desc": "Melee Weapon Attack: dice: d20+7 (+7) to hit, reach 5 ft., one target.\
    \ Hit: dice:2d8 + 3|text(12) (2d8 + 3) piercing damage, and the target must\
    \ make a DC 15 Constitution saving throw, taking dice:4d12|text(26) (4d12)\
    \ poison damage on a failed save, or half as much damage on a successful one.\
    \ If the poison damage reduces the target to 0 hit points, the target is stable\
    \ but [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned) for 1 hour, even\
    \ after regaining hit points, and is [paralyzed](/3-Mechanics/CLI/rules/conditions.md#paralyzed)\
    \ while [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned) in this way."
  "name": "Bite (Giant Spider Form Only)"
- "desc": "Ranged Weapon Attack: dice: d20+8 (+8) to hit, range 30/60 ft., one\
    \ target. Hit: The target is [restrained](/3-Mechanics/CLI/rules/conditions.md#restrained)\
    \ by webbing. As an action, the [restrained](/3-Mechanics/CLI/rules/conditions.md#restrained)\
    \ target can make a DC 15 Strength check, bursting the webbing on a success. The\
    \ webbing can also be attacked and destroyed (AC 10; hp 5; vulnerability to fire\
    \ damage; immunity to bludgeoning, poison, and psychic damage)."
  "name": "Web (Giant Spider Form Only Recharge 5–6)"
"source":
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MTF/Drow%20Arachnomancer.webp"
```
^statblock

```encounter-table
name: Drow Arachnomancer
creatures:
 - 1: Drow Arachnomancer
```

## Environment

underdark