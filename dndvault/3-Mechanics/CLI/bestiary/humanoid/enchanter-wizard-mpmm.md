---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/5
- monster/environment/urban
- monster/size/medium
- monster/type/humanoid
aliases: ["Enchanter Wizard"]
NoteIcon: monster
BestiaryType: humanoid
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 261, Volo's Guide to Monsters p. 213
---
# [Enchanter Wizard](3-Mechanics\CLI\bestiary\humanoid/enchanter-wizard-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 261, Volo's Guide to Monsters p. 213*  

```statblock
"name": "Enchanter Wizard (MPMM)"
"size": "Medium"
"type": "humanoid"
"alignment": "Any alignment"
"ac": !!int "12"
"ac_class": "15 with [mage armor](/3-Mechanics/CLI/spells/mage-armor.md)"
"hp": !!int "49"
"hit_dice": "11d8"
"stats":
- !!int "9"
- !!int "14"
- !!int "11"
- !!int "17"
- !!int "12"
- !!int "11"
"speed": "30 ft."
"saves":
  "Wisdom": !!int "4"
  "Intelligence": !!int "6"
"skillsaves":
  "History": !!int "6"
  "Arcana": !!int "6"
"senses": "passive Perception 11"
"languages": "any four languages"
"cr": "5"
"traits":
- "desc": "The enchanter casts one of the following spells, using Intelligence as\
    \ the spellcasting ability (spell save DC 14):\n\nAt will: [friends](/3-Mechanics/CLI/spells/friends.md),\
    \ [mage hand](/3-Mechanics/CLI/spells/mage-hand.md), [message](/3-Mechanics/CLI/spells/message.md)\n\
    \n2/day each: [charm person](/3-Mechanics/CLI/spells/charm-person.md), [mage\
    \ armor](/3-Mechanics/CLI/spells/mage-armor.md), [hold person](/3-Mechanics/CLI/spells/hold-person.md),\
    \ [invisibility](/3-Mechanics/CLI/spells/invisibility.md), [suggestion](/3-Mechanics/CLI/spells/suggestion.md),\
    \ [tongues](/3-Mechanics/CLI/spells/tongues.md)"
  "name": "Spellcasting"
"actions":
- "desc": "The enchanter makes three Arcane Burst attacks."
  "name": "Multiattack"
- "desc": "Melee or Ranged Spell Attack: dice: d20+6 (+6) to hit, reach 5 ft.\
    \ or range 120 ft., one target. Hit: dice:3d10 + 3|text(19) (3d10 + 3) psychic\
    \ damage."
  "name": "Arcane Burst"
"reactions":
- "desc": "When a visible creature within 30 feet of the enchanter makes an attack\
    \ roll against it, the enchanter forces the attacker to make a DC 14 Wisdom saving\
    \ throw. On a failed save, the attacker redirects the attack roll to the creature\
    \ closest to it, other than the enchanter or itself. If multiple eligible creatures\
    \ are closest, the attacker chooses which one to target."
  "name": "Instinctive Charm (Recharge 4-6)"
"source":
- "MPMM"
- "VGM"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Enchanter%20Wizard.webp"
```
^statblock

```encounter-table
name: Enchanter Wizard
creatures:
 - 1: Enchanter Wizard
```

## Environment

urban