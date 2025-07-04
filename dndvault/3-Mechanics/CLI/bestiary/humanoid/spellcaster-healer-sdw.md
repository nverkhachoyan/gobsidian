---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/sdw
- monster/cr/
- monster/size/medium
- monster/type/humanoid/healer
aliases: ["Spellcaster (Healer)"]
NoteIcon: monster
BestiaryType: humanoid (healer)
SourceType: Bestiary
BookSource: Sleeping Dragon's Wake
---
# [Spellcaster (Healer)](3-Mechanics\CLI\bestiary\humanoid/spellcaster-healer-sdw.md)
*Source: Sleeping Dragon's Wake*  

```statblock
"name": "Spellcaster (Healer) (SDW)"
"size": "Medium"
"type": "humanoid"
"subtype": "healer"
"alignment": "Unaligned"
"ac": !!int "13"
"ac_class": "[studded leather](/3-Mechanics/CLI/items/studded-leather-armor.md)"
"hp": !!int "45"
"hit_dice": "10d8"
"stats":
- !!int "10"
- !!int "12"
- !!int "10"
- !!int "15"
- !!int "18"
- !!int "13"
"speed": "30 ft."
"saves":
  "Wisdom": !!int "8"
"skillsaves":
  "Investigation": !!int "6"
  "Religion": !!int "6"
  "Arcana": !!int "6"
"senses": "passive Perception 14"
"languages": "Common, plus one of your choice"
"traits":
- "desc": "The spellcaster's spellcasting ability is Wisdom (spell save DC 16, dice:\
    \ d20+8 (+8) to hit with spell attacks). The spellcaster has following cleric\
    \ spells prepared:\n\nCantrips (at will): [guidance](/3-Mechanics/CLI/spells/guidance.md),\
    \ [light](/3-Mechanics/CLI/spells/light.md), [resistance](/3-Mechanics/CLI/spells/resistance.md),\
    \ [sacred flame](/3-Mechanics/CLI/spells/sacred-flame.md)\n\n1st level (4 slots):\
    \ [bless](/3-Mechanics/CLI/spells/bless.md), [cure wounds](/3-Mechanics/CLI/spells/cure-wounds.md),\
    \ [shield of faith](/3-Mechanics/CLI/spells/shield-of-faith.md)\n\n2nd level\
    \ (3 slots): [aid](/3-Mechanics/CLI/spells/aid.md), [lesser restoration](/3-Mechanics/CLI/spells/lesser-restoration.md)\n\
    \n3rd level (3 slots): [protection from energy](/3-Mechanics/CLI/spells/protection-from-energy.md),\
    \ [revivify](/3-Mechanics/CLI/spells/revivify.md)\n\n4th level (3 slots):\
    \ [banishment](/3-Mechanics/CLI/spells/banishment.md), [death ward](/3-Mechanics/CLI/spells/death-ward.md)\n\
    \n5th level (1 slots): [greater restoration](/3-Mechanics/CLI/spells/greater-restoration.md)"
  "name": "Spellcasting (Healer)"
- "desc": "The spellcaster can add its spellcasting ability modifier to the damage\
    \ it deals with any cantrip."
  "name": "Potent Cantrip"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+4 (+4) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d6|text(3) (1d6) bludgeoning damage, or dice:1d8|text(4)\
    \ (1d8) bludgeoning damage if used with two hands."
  "name": "Quarterstaff"
"source":
- "SDW"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/SDW/Spellcaster%20%28Healer%29.webp"
```
^statblock

```encounter-table
name: Spellcaster (Healer)
creatures:
 - 1: Spellcaster (Healer)
```