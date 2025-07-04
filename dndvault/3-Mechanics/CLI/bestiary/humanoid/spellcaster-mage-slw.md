---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/slw
- monster/cr/
- monster/size/medium
- monster/type/humanoid/mage
aliases: ["Spellcaster (Mage)"]
NoteIcon: monster
BestiaryType: humanoid (mage)
SourceType: Bestiary
BookSource: Storm Lord's Wrath
---
# [Spellcaster (Mage)](3-Mechanics\CLI\bestiary\humanoid/spellcaster-mage-slw.md)
*Source: Storm Lord's Wrath*  

```statblock
"name": "Spellcaster (Mage) (SLW)"
"size": "Medium"
"type": "humanoid"
"subtype": "mage"
"alignment": "Unaligned"
"ac": !!int "13"
"ac_class": "[studded leather](/3-Mechanics/CLI/items/studded-leather-armor.md)"
"hp": !!int "36"
"hit_dice": "8d8"
"stats":
- !!int "10"
- !!int "12"
- !!int "10"
- !!int "16"
- !!int "14"
- !!int "13"
"speed": "30 ft."
"saves":
  "Wisdom": !!int "5"
"skillsaves":
  "Investigation": !!int "6"
  "Religion": !!int "6"
  "Arcana": !!int "6"
"senses": "passive Perception 12"
"languages": "Common, plus one of your choice"
"traits":
- "desc": "The spellcaster's spellcasting ability is Intelligence (spell save DC 14,\
    \ dice: d20+6 (+6) to hit with spell attacks). The spellcaster has following\
    \ wizard spells prepared:\n\nCantrips (at will): [fire bolt](/3-Mechanics/CLI/spells/fire-bolt.md),\
    \ [light](/3-Mechanics/CLI/spells/light.md), [mage hand](/3-Mechanics/CLI/spells/mage-hand.md),\
    \ [minor illusion](/3-Mechanics/CLI/spells/minor-illusion.md)\n\n1st level (4\
    \ slots): [burning hands](/3-Mechanics/CLI/spells/burning-hands.md), [shield](/3-Mechanics/CLI/spells/shield.md),\
    \ [sleep](/3-Mechanics/CLI/spells/sleep.md)\n\n2nd level (3 slots): [flaming\
    \ sphere](/3-Mechanics/CLI/spells/flaming-sphere.md), [invisibility](/3-Mechanics/CLI/spells/invisibility.md)\n\
    \n3rd level (3 slots): [fireball](/3-Mechanics/CLI/spells/fireball.md), [fly](/3-Mechanics/CLI/spells/fly.md)\n\
    \n4th level (1 slots): [wall of fire](/3-Mechanics/CLI/spells/wall-of-fire.md)"
  "name": "Spellcasting (Mage)"
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
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/SLW/Spellcaster%20%28Mage%29.webp"
```
^statblock

```encounter-table
name: Spellcaster (Mage)
creatures:
 - 1: Spellcaster (Mage)
```