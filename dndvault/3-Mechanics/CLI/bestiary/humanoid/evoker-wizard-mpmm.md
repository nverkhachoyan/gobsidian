---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/9
- monster/environment/urban
- monster/size/medium
- monster/type/humanoid
aliases: ["Evoker Wizard"]
NoteIcon: monster
BestiaryType: humanoid
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 262, Volo's Guide to Monsters p. 214
---
# [Evoker Wizard](3-Mechanics\CLI\bestiary\humanoid/evoker-wizard-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 262, Volo's Guide to Monsters p. 214*  

```statblock
"name": "Evoker Wizard (MPMM)"
"size": "Medium"
"type": "humanoid"
"alignment": "Any alignment"
"ac": !!int "12"
"ac_class": "15 with [mage armor](/3-Mechanics/CLI/spells/mage-armor.md)"
"hp": !!int "121"
"hit_dice": "22d8 + 22"
"stats":
- !!int "9"
- !!int "14"
- !!int "12"
- !!int "17"
- !!int "12"
- !!int "11"
"speed": "30 ft."
"saves":
  "Wisdom": !!int "5"
  "Intelligence": !!int "7"
"skillsaves":
  "History": !!int "7"
  "Arcana": !!int "7"
"senses": "passive Perception 11"
"languages": "any four languages"
"cr": "9"
"traits":
- "desc": "The evoker casts one of the following spells, using Intelligence as the\
    \ spellcasting ability (spell save DC 15):\n\nAt will: [light](/3-Mechanics/CLI/spells/light.md),\
    \ [mage hand](/3-Mechanics/CLI/spells/mage-hand.md), [message](/3-Mechanics/CLI/spells/message.md),\
    \ [prestidigitation](/3-Mechanics/CLI/spells/prestidigitation.md)\n\n1/day each:\
    \ [wall of ice](/3-Mechanics/CLI/spells/wall-of-ice.md)\n\n2/day each: [ice\
    \ storm](/3-Mechanics/CLI/spells/ice-storm.md), [lightning bolt](/3-Mechanics/CLI/spells/lightning-bolt.md),\
    \ [mage armor](/3-Mechanics/CLI/spells/mage-armor.md)"
  "name": "Spellcasting"
"actions":
- "desc": "The evoker makes three Arcane Burst attacks."
  "name": "Multiattack"
- "desc": "Melee or Ranged Spell Attack: dice: d20+7 (+7) to hit, reach 5 ft.\
    \ or range 120 ft., one target. Hit: dice:4d10 + 3|text(25) (4d10 + 3) force\
    \ damage."
  "name": "Arcane Burst"
- "desc": "The evoker unleashes a magical explosion of a particular damage type: cold,\
    \ fire, lightning, or thunder. The magic erupts in a 20-foot-radius sphere centered\
    \ on a point within 150 feet of the evoker. Each creature in that area must make\
    \ a DC 15 Dexterity saving throw. The evoker can select up to three creatures\
    \ it can see in the area to ignore the spell, as the evoker sculpts the spell's\
    \ energy around them. On a failed save, a creature takes dice:9d8|text(40) (9d8)\
    \ damage of the chosen type and is knocked [prone](/3-Mechanics/CLI/rules/conditions.md#prone).\
    \ On a successful save, a creature takes half as much damage and isn't knocked\
    \ [prone](/3-Mechanics/CLI/rules/conditions.md#prone)."
  "name": "Sculpted Explosion (Recharge 4-6)"
"source":
- "MPMM"
- "VGM"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Evoker%20Wizard.webp"
```
^statblock

```encounter-table
name: Evoker Wizard
creatures:
 - 1: Evoker Wizard
```

## Environment

urban