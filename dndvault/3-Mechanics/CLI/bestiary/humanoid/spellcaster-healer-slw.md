---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/slw
- monster/cr/
- monster/size/medium
- monster/type/humanoid/healer
aliases: ["Spellcaster (Healer)"]
NoteIcon: monster
BestiaryType: humanoid (healer)
SourceType: Bestiary
BookSource: Storm Lord's Wrath
---
# [Spellcaster (Healer)](3-Mechanics\CLI\bestiary\humanoid/spellcaster-healer-slw.md)
*Source: Storm Lord's Wrath*  

```statblock
"name": "Spellcaster (Healer) (SLW)"
"size": "Medium"
"type": "humanoid"
"subtype": "healer"
"alignment": "Unaligned"
"ac": !!int "13"
"ac_class": "[studded leather](/3-Mechanics/CLI/items/studded-leather-armor.md)"
"hp": !!int "36"
"hit_dice": "8d8"
"stats":
- !!int "10"
- !!int "12"
- !!int "10"
- !!int "15"
- !!int "16"
- !!int "13"
"speed": "30 ft."
"saves":
  "Wisdom": !!int "6"
"skillsaves":
  "Investigation": !!int "5"
  "Religion": !!int "5"
  "Arcana": !!int "5"
"senses": "passive Perception 13"
"languages": "Common, plus one of your choice"
"traits":
- "desc": "The spellcaster's spellcasting ability is Wisdom (spell save DC 14, dice:\
    \ d20+6 (+6) to hit with spell attacks). The spellcaster has following cleric\
    \ spells prepared:\n\nCantrips (at will): [guidance](/3-Mechanics/CLI/spells/guidance.md),\
    \ [light](/3-Mechanics/CLI/spells/light.md), [resistance](/3-Mechanics/CLI/spells/resistance.md),\
    \ [sacred flame](/3-Mechanics/CLI/spells/sacred-flame.md)\n\n1st level (4 slots):\
    \ [bless](/3-Mechanics/CLI/spells/bless.md), [cure wounds](/3-Mechanics/CLI/spells/cure-wounds.md),\
    \ [shield of faith](/3-Mechanics/CLI/spells/shield-of-faith.md)\n\n2nd level\
    \ (3 slots): [aid](/3-Mechanics/CLI/spells/aid.md), [lesser restoration](/3-Mechanics/CLI/spells/lesser-restoration.md)\n\
    \n3rd level (3 slots): [protection from energy](/3-Mechanics/CLI/spells/protection-from-energy.md),\
    \ [revivify](/3-Mechanics/CLI/spells/revivify.md)\n\n4th level (1 slots):\
    \ [death ward](/3-Mechanics/CLI/spells/death-ward.md)"
  "name": "Spellcasting (Healer)"
- "desc": "The spellcaster can add its spellcasting ability modifier to the damage\
    \ it deals with any cantrip."
  "name": "Potent Cantrip"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+3 (+3) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d6|text(3) (1d6) bludgeoning damage, or dice:1d8|text(4)\
    \ (1d8) bludgeoning damage if used with two hands."
  "name": "Quarterstaff"
"source":
- "SLW"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/SLW/Spellcaster%20%28Healer%29.webp"
```
^statblock

```encounter-table
name: Spellcaster (Healer)
creatures:
 - 1: Spellcaster (Healer)
```