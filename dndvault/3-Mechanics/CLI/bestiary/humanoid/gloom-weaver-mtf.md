---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mtf
- monster/cr/9
- monster/environment/underdark
- monster/environment/urban
- monster/size/medium
- monster/type/humanoid/elf
aliases: ["Gloom Weaver"]
NoteIcon: monster
BestiaryType: humanoid (elf)
SourceType: Bestiary
BookSource: Mordenkainen's Tome of Foes p. 224
---
# [Gloom Weaver](3-Mechanics\CLI\bestiary\humanoid/gloom-weaver-mtf.md)
*Source: Mordenkainen's Tome of Foes p. 224*  

## Gloom Weaver

Although a formidable fighter, a gloom weaver is often content to remain hidden in the shadows, watching with rapt attention as its very presence affects its victims. Its dark energy weighs down the heart, causing those within its oppressive aura to feel the approach of death. This torment alone is enough to please its master, the Raven Queen, but should it be detected, a gloom weaver uses its shadow magic to reduce its enemies to ghastly corpses.

## Shadar-kai

In the perpetual gloom of the Shadowfell lives a society that serves the Raven Queen. They were brought into that dusky realm in ages past, so long ago that they're now perfectly adapted to that cheerless environment, both physically and mentally.

## Soul Custodians

Shadar-kai watch over both the Shadowfell and the material world, scouting out choice souls and tragedies that might please their deity. They are rumored to be able to coax worldly events along tragic paths for her amusement. The Raven Queen is famously cryptic even to her most devoted followers, however. The shadar-kai's efforts are rewarded only with vague omens they interpret as best they can.

## Blighted Elves

Shadar-kai were once elves, but eons of exposure to the debilitating influence of the Shadowfell has left them joyless and mournful. In that realm, they have the appearance of withered elves: pale hair, wrinkled gray skin, and swollen joints give them a corpselike aspect. They appear more youthful while on other planes, but their skin always retains its deathly pallor. They dress in dark cloaks and heavy veils, detest mirrors, and avoid keeping things that remind them of their age.

```statblock
"name": "Gloom Weaver (MTF)"
"size": "Medium"
"type": "humanoid"
"subtype": "elf"
"alignment": "Neutral"
"ac": !!int "14"
"ac_class": "17 with [mage armor](/3-Mechanics/CLI/spells/mage-armor.md)"
"hp": !!int "104"
"hit_dice": "16d8 + 32"
"stats":
- !!int "11"
- !!int "18"
- !!int "14"
- !!int "15"
- !!int "12"
- !!int "18"
"speed": "30 ft."
"saves":
  "Dexterity": !!int "8"
  "Constitution": !!int "6"
"damage_immunities": "necrotic"
"condition_immunities": "[charmed](/3-Mechanics/CLI/rules/conditions.md#charmed),\
  \ [exhaustion](/3-Mechanics/CLI/rules/conditions.md#exhaustion)"
"senses": "darkvision 60 ft., passive Perception 11"
"languages": "Common, Elvish"
"cr": "9"
"traits":
- "desc": "The gloom weaver's innate spellcasting ability is Charisma (spell save\
    \ DC 16, dice: d20+8 (+8) to hit with spell attacks). It can innately cast\
    \ the following spells, requiring no material components:\n\nAt will: [arcane\
    \ eye](/3-Mechanics/CLI/spells/arcane-eye.md), [mage armor](/3-Mechanics/CLI/spells/mage-armor.md),\
    \ [speak with dead](/3-Mechanics/CLI/spells/speak-with-dead.md)\n\n1/day each:\
    \ [arcane gate](/3-Mechanics/CLI/spells/arcane-gate.md), [bane](/3-Mechanics/CLI/spells/bane.md),\
    \ [compulsion](/3-Mechanics/CLI/spells/compulsion.md), [confusion](/3-Mechanics/CLI/spells/confusion.md),\
    \ [true seeing](/3-Mechanics/CLI/spells/true-seeing.md)"
  "name": "Innate Spellcasting"
- "desc": "The gloom weaver is a 12th-level spellcaster. Its spellcasting ability\
    \ is Charisma (spell save DC 16, dice: d20+8 (+8) to hit with spell attacks).\
    \ It regains its expended spell slots when it finishes a short or long rest. It\
    \ knows the following warlock spells:\n\nCantrips (at will): [minor illusion](/3-Mechanics/CLI/spells/minor-illusion.md),\
    \ [prestidigitation](/3-Mechanics/CLI/spells/prestidigitation.md), [chill touch](/3-Mechanics/CLI/spells/chill-touch.md),\
    \ [eldritch blast](/3-Mechanics/CLI/spells/eldritch-blast.md)\n\n1st-5th level\
    \ (3 slots): [armor of Agathys](/3-Mechanics/CLI/spells/armor-of-agathys.md),\
    \ [blight](/3-Mechanics/CLI/spells/blight.md), [darkness](/3-Mechanics/CLI/spells/darkness.md),\
    \ [dream](/3-Mechanics/CLI/spells/dream.md), [invisibility](/3-Mechanics/CLI/spells/invisibility.md),\
    \ [fear](/3-Mechanics/CLI/spells/fear.md), [hypnotic pattern](/3-Mechanics/CLI/spells/hypnotic-pattern.md),\
    \ [major image](/3-Mechanics/CLI/spells/major-image.md), [contact other plane](/3-Mechanics/CLI/spells/contact-other-plane.md),\
    \ [vampiric touch](/3-Mechanics/CLI/spells/vampiric-touch.md), [witch bolt](/3-Mechanics/CLI/spells/witch-bolt.md)\n\
    \n3 beams +4 bonus to each damage roll"
  "name": "Spellcasting"
- "desc": "Beasts and humanoids, other than shadar-kai, have disadvantage on saving\
    \ throws while within 10 feet of the gloom weaver."
  "name": "Burden of Time"
- "desc": "The gloom weaver has advantage on saving throws against being [charmed](/3-Mechanics/CLI/rules/conditions.md#charmed),\
    \ and magic can't put it to sleep."
  "name": "Fey Ancestry"
"actions":
- "desc": "The gloom weaver makes two spear attacks and casts one spell that takes\
    \ 1 action to cast."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+8 (+8) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d6 + 4|text(7) (1d6 + 4) piercing damage plus dice:4d12|text(26)\
    \ (4d12) necrotic damage, or dice:1d8 + 4|text(8) (1d8 + 4) piercing damage\
    \ plus dice:4d12|text(26) (4d12) necrotic damage if used with two hands."
  "name": "Shadow Spear"
"reactions":
- "desc": "When the gloom weaver takes damage, it turns [invisible](/3-Mechanics/CLI/rules/conditions.md#invisible)\
    \ and teleports up to 60 feet to an unoccupied space it can see. It remains [invisible](/3-Mechanics/CLI/rules/conditions.md#invisible)\
    \ until the start of its next turn or until it attacks or casts a spell."
  "name": "Misty Escape (Recharges after a Short or Long Rest)"
"source":
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MTF/Gloom%20Weaver.webp"
```
^statblock

```encounter-table
name: Gloom Weaver
creatures:
 - 1: Gloom Weaver
```

## Environment

underdark, urban