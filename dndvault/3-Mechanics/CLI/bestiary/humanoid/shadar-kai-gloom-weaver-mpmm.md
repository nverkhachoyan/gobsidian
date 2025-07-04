---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/9
- monster/environment/underdark
- monster/environment/urban
- monster/size/medium
- monster/type/humanoid/elf
aliases: ["Shadar-kai Gloom Weaver"]
NoteIcon: monster
BestiaryType: humanoid (elf)
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 213, Mordenkainen's Tome of Foes p. 224
---
# [Shadar-kai Gloom Weaver](3-Mechanics\CLI\bestiary\humanoid/shadar-kai-gloom-weaver-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 213, Mordenkainen's Tome of Foes p. 224*  

```statblock
"name": "Shadar-kai Gloom Weaver (MPMM)"
"size": "Medium"
"type": "humanoid"
"subtype": "elf"
"alignment": "Typically  Neutral Evil"
"ac": !!int "14"
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
- "desc": "The shadar-kai casts one of the following spells, requiring no material\
    \ components and using Charisma as the spellcasting ability (spell save DC 16):\n\
    \nAt will: [arcane eye](/3-Mechanics/CLI/spells/arcane-eye.md), [mage armor](/3-Mechanics/CLI/spells/mage-armor.md),\
    \ [minor illusion](/3-Mechanics/CLI/spells/minor-illusion.md), [prestidigitation](/3-Mechanics/CLI/spells/prestidigitation.md),\
    \ [speak with dead](/3-Mechanics/CLI/spells/speak-with-dead.md)\n\n1/day each:\
    \ [arcane gate](/3-Mechanics/CLI/spells/arcane-gate.md), [bane](/3-Mechanics/CLI/spells/bane.md),\
    \ [confusion](/3-Mechanics/CLI/spells/confusion.md), [darkness](/3-Mechanics/CLI/spells/darkness.md),\
    \ [fear](/3-Mechanics/CLI/spells/fear.md), [major image](/3-Mechanics/CLI/spells/major-image.md),\
    \ [true seeing](/3-Mechanics/CLI/spells/true-seeing.md)"
  "name": "Spellcasting"
- "desc": "Beasts and Humanoids (except elves) have disadvantage on saving throws\
    \ while within 10 feet of the shadar-kai."
  "name": "Burden of Time"
- "desc": "The shadar-kai has advantage on saving throws against being [charmed](/3-Mechanics/CLI/rules/conditions.md#charmed),\
    \ and magic can't put it to sleep."
  "name": "Fey Ancestry"
"actions":
- "desc": "The shadar-kai makes three Shadow Spear attacks. It can replace one attack\
    \ with a use of Spellcasting."
  "name": "Multiattack"
- "desc": "Melee or Ranged Weapon Attack: dice: d20+8 (+8) to hit, reach 5 ft.\
    \ or range 30/120, one target. Hit: dice:1d6 + 4|text(7) (1d6 + 4) piercing\
    \ damage plus dice:4d12|text(26) (4d12) necrotic damage. Hit or Miss: The\
    \ spear magically returns to the shadar-kai's hand immediately after a ranged\
    \ attack."
  "name": "Shadow Spear"
"reactions":
- "desc": "When the shadar-kai takes damage, it turns [invisible](/3-Mechanics/CLI/rules/conditions.md#invisible)\
    \ and teleports, along with any equipment it is wearing or carrying, up to 60\
    \ feet to an unoccupied space it can see. It remains [invisible](/3-Mechanics/CLI/rules/conditions.md#invisible)\
    \ until the start of its next turn or until it attacks or casts a spell."
  "name": "Misty Escape (Recharge 6-6)"
"source":
- "MPMM"
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Shadar-kai%20Gloom%20Weaver.webp"
```
^statblock

```encounter-table
name: Shadar-kai Gloom Weaver
creatures:
 - 1: Shadar-kai Gloom Weaver
```

## Environment

underdark, urban