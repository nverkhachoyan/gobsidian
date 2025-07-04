---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/vgm
- monster/cr/6
- monster/environment/desert
- monster/environment/hill
- monster/environment/mountain
- monster/environment/underdark
- monster/environment/urban
- monster/size/medium
- monster/type/humanoid/any-race
aliases: ["Warlock of the Great Old One"]
NoteIcon: monster
BestiaryType: humanoid (any race)
SourceType: Bestiary
BookSource: Volo's Guide to Monsters p. 220
---
# [Warlock of the Great Old One](3-Mechanics\CLI\bestiary\humanoid/warlock-of-the-great-old-one-vgm.md)
*Source: Volo's Guide to Monsters p. 220*  

Warlocks of the Great Old One gain their powers through magical pacts forged with eldritch entities from strange and distant realms of existence. Some of these warlocks associate with cultists devoted to these entities, as well as aberrations that share their goals, yet other warlocks of the Great Old One are experts at rooting out the insanity and wickedness inspired by bizarre beings from beyond the stars.

```statblock
"name": "Warlock of the Great Old One (VGM)"
"size": "Medium"
"type": "humanoid"
"subtype": "any race"
"alignment": "Any alignment"
"ac": !!int "12"
"ac_class": "15 with [mage armor](/3-Mechanics/CLI/spells/mage-armor.md)"
"hp": !!int "91"
"hit_dice": "14d8 + 28"
"stats":
- !!int "9"
- !!int "14"
- !!int "15"
- !!int "12"
- !!int "12"
- !!int "18"
"speed": "30 ft."
"saves":
  "Charisma": !!int "7"
  "Wisdom": !!int "4"
"skillsaves":
  "History": !!int "4"
  "Arcana": !!int "4"
"damage_resistances": "psychic"
"senses": "darkvision 60 ft., passive Perception 11"
"languages": "any two languages, telepathy 30 ft."
"cr": "6"
"traits":
- "desc": "The warlock's innate spellcasting ability is Charisma. It can innately\
    \ cast the following spells (spell save DC 15), requiring no material components:\n\
    \nAt will: [detect magic](/3-Mechanics/CLI/spells/detect-magic.md), [jump](/3-Mechanics/CLI/spells/jump.md),\
    \ [levitate](/3-Mechanics/CLI/spells/levitate.md), [mage armor](/3-Mechanics/CLI/spells/mage-armor.md)\
    \ (self only), [speak with dead](/3-Mechanics/CLI/spells/speak-with-dead.md)\n\
    \n1/day each: [arcane gate](/3-Mechanics/CLI/spells/arcane-gate.md), [true\
    \ seeing](/3-Mechanics/CLI/spells/true-seeing.md)"
  "name": "Innate Spellcasting"
- "desc": "The warlock is a 14th-level spellcaster. Its spellcasting ability is Charisma\
    \ (spell save DC 15, dice: d20+7 (+7) to hit with spell attacks). It regains\
    \ its expended spell slots when it finishes a short or long rest. It knows the\
    \ following warlock spells:\n\nCantrips (at will): [chill touch](/3-Mechanics/CLI/spells/chill-touch.md),\
    \ [eldritch blast](/3-Mechanics/CLI/spells/eldritch-blast.md), [guidance](/3-Mechanics/CLI/spells/guidance.md),\
    \ [mage hand](/3-Mechanics/CLI/spells/mage-hand.md), [minor illusion](/3-Mechanics/CLI/spells/minor-illusion.md),\
    \ [prestidigitation](/3-Mechanics/CLI/spells/prestidigitation.md), [shocking grasp](/3-Mechanics/CLI/spells/shocking-grasp.md)\n\
    \n1st-5th level (3 slots): [armor of Agathys](/3-Mechanics/CLI/spells/armor-of-agathys.md),\
    \ [arms of Hadar](/3-Mechanics/CLI/spells/arms-of-hadar.md), [crown of madness](/3-Mechanics/CLI/spells/crown-of-madness.md),\
    \ [clairvoyance](/3-Mechanics/CLI/spells/clairvoyance.md), [contact other plane](/3-Mechanics/CLI/spells/contact-other-plane.md),\
    \ [detect thoughts](/3-Mechanics/CLI/spells/detect-thoughts.md), [dimension door](/3-Mechanics/CLI/spells/dimension-door.md),\
    \ [dissonant whispers](/3-Mechanics/CLI/spells/dissonant-whispers.md), [dominate\
    \ beast](/3-Mechanics/CLI/spells/dominate-beast.md), [telekinesis](/3-Mechanics/CLI/spells/telekinesis.md),\
    \ [vampiric touch](/3-Mechanics/CLI/spells/vampiric-touch.md)"
  "name": "Spellcasting"
- "desc": "At the start of each of the warlock's turns, each creature of its choice\
    \ within 5 feet of it must succeed on a DC 15 Wisdom saving throw or take dice:3d6|text(10)\
    \ (3d6) psychic damage, provided that the warlock isn't [incapacitated](/3-Mechanics/CLI/rules/conditions.md#incapacitated)."
  "name": "Whispering Aura"
"actions":
- "desc": "Melee or Ranged Weapon Attack: dice: d20+5 (+5) to hit, reach 5 ft.\
    \ or range 20/60 ft., one target. Hit: dice:1d4 + 2|text(4) (1d4 + 2) piercing\
    \ damage."
  "name": "Dagger"
"source":
- "VGM"
- "IMR"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/VGM/Warlock%20of%20the%20Great%20Old%20One.webp"
```
^statblock

```encounter-table
name: Warlock of the Great Old One
creatures:
 - 1: Warlock of the Great Old One
```

## Environment

desert, hill, mountain, underdark, urban