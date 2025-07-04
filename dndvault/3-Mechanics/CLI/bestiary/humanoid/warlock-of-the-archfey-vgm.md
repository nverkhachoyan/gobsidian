---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/vgm
- monster/cr/4
- monster/environment/arctic
- monster/environment/forest
- monster/environment/mountain
- monster/environment/swamp
- monster/environment/urban
- monster/size/medium
- monster/type/humanoid/any-race
aliases: ["Warlock of the Archfey"]
NoteIcon: monster
BestiaryType: humanoid (any race)
SourceType: Bestiary
BookSource: Volo's Guide to Monsters p. 219
---
# [Warlock of the Archfey](3-Mechanics\CLI\bestiary\humanoid/warlock-of-the-archfey-vgm.md)
*Source: Volo's Guide to Monsters p. 219*  

Warlocks of the archfey gain their powers through magical pacts forged with lords of the Feywild. These warlocks commonly associate with lesser fey creatures such as boggles, quicklings, redcaps, satyrs, and sprites.

```statblock
"name": "Warlock of the Archfey (VGM)"
"size": "Medium"
"type": "humanoid"
"subtype": "any race"
"alignment": "Any alignment"
"ac": !!int "11"
"ac_class": "14 with [mage armor](/3-Mechanics/CLI/spells/mage-armor.md)"
"hp": !!int "49"
"hit_dice": "11d8"
"stats":
- !!int "9"
- !!int "13"
- !!int "11"
- !!int "11"
- !!int "12"
- !!int "18"
"speed": "30 ft."
"saves":
  "Charisma": !!int "6"
  "Wisdom": !!int "3"
"skillsaves":
  "Nature": !!int "2"
  "Deception": !!int "6"
  "Arcana": !!int "2"
  "Persuasion": !!int "6"
"condition_immunities": "[charmed](/3-Mechanics/CLI/rules/conditions.md#charmed)"
"senses": "passive Perception 11"
"languages": "any two languages (usually Sylvan)"
"cr": "4"
"traits":
- "desc": "The warlock's innate spellcasting ability is Charisma. It can innately\
    \ cast the following spells (spell save DC 15), requiring no material components:\n\
    \nAt will: [disguise self](/3-Mechanics/CLI/spells/disguise-self.md), [mage\
    \ armor](/3-Mechanics/CLI/spells/mage-armor.md) (self only), [silent image](/3-Mechanics/CLI/spells/silent-image.md),\
    \ [speak with animals](/3-Mechanics/CLI/spells/speak-with-animals.md)\n\n1/day:\
    \ [conjure fey](/3-Mechanics/CLI/spells/conjure-fey.md)"
  "name": "Innate Spellcasting"
- "desc": "The warlock is an 11th-level spellcaster. Its spellcasting ability is Charisma\
    \ (spell save DC 14, dice: d20+6 (+6) to hit with spell attacks). It regains\
    \ its expended spell slots when it finishes a short or long rest. It knows the\
    \ following warlock spells:\n\nCantrips (at will): [dancing lights](/3-Mechanics/CLI/spells/dancing-lights.md),\
    \ [eldritch blast](/3-Mechanics/CLI/spells/eldritch-blast.md), [friends](/3-Mechanics/CLI/spells/friends.md),\
    \ [mage hand](/3-Mechanics/CLI/spells/mage-hand.md), [minor illusion](/3-Mechanics/CLI/spells/minor-illusion.md),\
    \ [prestidigitation](/3-Mechanics/CLI/spells/prestidigitation.md), [vicious mockery](/3-Mechanics/CLI/spells/vicious-mockery.md)\n\
    \n1st-5th level (3 slots): [blink](/3-Mechanics/CLI/spells/blink.md), [charm\
    \ person](/3-Mechanics/CLI/spells/charm-person.md), [dimension door](/3-Mechanics/CLI/spells/dimension-door.md),\
    \ [dominate beast](/3-Mechanics/CLI/spells/dominate-beast.md), [faerie fire](/3-Mechanics/CLI/spells/faerie-fire.md),\
    \ [fear](/3-Mechanics/CLI/spells/fear.md), [hold monster](/3-Mechanics/CLI/spells/hold-monster.md),\
    \ [misty step](/3-Mechanics/CLI/spells/misty-step.md), [phantasmal force](/3-Mechanics/CLI/spells/phantasmal-force.md),\
    \ [seeming](/3-Mechanics/CLI/spells/seeming.md), [sleep](/3-Mechanics/CLI/spells/sleep.md)"
  "name": "Spellcasting"
"actions":
- "desc": "Melee or Ranged Weapon Attack: dice: d20+3 (+3) to hit, reach 5 ft.\
    \ or range 20/60 ft., one target. Hit: dice:1d4 + 2|text(4) (1d4 + 2) piercing\
    \ damage."
  "name": "Dagger"
"reactions":
- "desc": "In response to taking damage, the warlock turns [invisible](/3-Mechanics/CLI/rules/conditions.md#invisible)\
    \ and teleports up to 60 feet to an unoccupied space it can see. It remains [invisible](/3-Mechanics/CLI/rules/conditions.md#invisible)\
    \ until the start of its next turn or until it attacks, makes a damage roll, or\
    \ casts a spell."
  "name": "Misty Escape (Recharges after a Short or Long Rest)"
"source":
- "VGM"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/VGM/Warlock%20of%20the%20Archfey.webp"
```
^statblock

```encounter-table
name: Warlock of the Archfey
creatures:
 - 1: Warlock of the Archfey
```

## Environment

arctic, forest, mountain, swamp, urban