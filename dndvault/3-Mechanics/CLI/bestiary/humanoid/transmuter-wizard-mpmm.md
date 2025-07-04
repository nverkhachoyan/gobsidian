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
aliases: ["Transmuter Wizard"]
NoteIcon: monster
BestiaryType: humanoid
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 265, Volo's Guide to Monsters p. 218
---
# [Transmuter Wizard](3-Mechanics\CLI\bestiary\humanoid/transmuter-wizard-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 265, Volo's Guide to Monsters p. 218*  

```statblock
"name": "Transmuter Wizard (MPMM)"
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
- "desc": "The transmuter casts one of the following spells, using Intelligence as\
    \ the spellcasting ability (spell save DC 14):\n\nAt will: [light](/3-Mechanics/CLI/spells/light.md),\
    \ [message](/3-Mechanics/CLI/spells/message.md), [prestidigitation](/3-Mechanics/CLI/spells/prestidigitation.md)\n\
    \n1/day each: [telekinesis](/3-Mechanics/CLI/spells/telekinesis.md)\n\n2/day\
    \ each: [fireball](/3-Mechanics/CLI/spells/fireball.md), [hold person](/3-Mechanics/CLI/spells/hold-person.md),\
    \ [knock](/3-Mechanics/CLI/spells/knock.md), [mage armor](/3-Mechanics/CLI/spells/mage-armor.md),\
    \ [polymorph](/3-Mechanics/CLI/spells/polymorph.md), [slow](/3-Mechanics/CLI/spells/slow.md)"
  "name": "Spellcasting"
- "desc": "The transmuter carries a magic stone it crafted. The stone grants it one\
    \ of the following benefits while bearing the stone; the transmuter chooses the\
    \ benefit at the end of each long rest:\n\n- Darkvision. The transmuter has\
    \ [darkvision](/3-Mechanics/CLI/rules/senses.md#darkvision) out to a range of\
    \ 60 feet.  \n- Resilience. The transmuter has proficiency in Constitution\
    \ saving throws.   \n- Resistance. Resistance. The transmuter has resistance\
    \ to acid, cold, fire, lightning, or thunder damage (transmuter's choice whenever\
    \ choosing this benefit).  \n- Speed. The transmuter's walking speed is increased\
    \ by 10 feet.  "
  "name": "Transmuter's Stone"
"actions":
- "desc": "The transmuter makes three Arcane Burst attacks."
  "name": "Multiattack"
- "desc": "Melee or Ranged Spell Attack: dice: d20+6 (+6) to hit, reach 5 ft.\
    \ or range 120 ft., one target. Hit: dice:3d10 + 3|text(19) (3d10 + 3) acid\
    \ damage."
  "name": "Arcane Burst"
"bonus_actions":
- "desc": "The transmuter casts [alter self](/3-Mechanics/CLI/spells/alter-self.md)\
    \ or changes the benefit of Transmuter's Stone if bearing the stone."
  "name": "Transmute (Recharge 4-6)"
"source":
- "MPMM"
- "VGM"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Transmuter%20Wizard.webp"
```
^statblock

```encounter-table
name: Transmuter Wizard
creatures:
 - 1: Transmuter Wizard
```

## Environment

urban