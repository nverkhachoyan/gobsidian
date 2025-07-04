---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/vgm
- monster/cr/4
- monster/environment/forest
- monster/environment/grassland
- monster/environment/hill
- monster/size/medium
- monster/type/humanoid/goblinoid
aliases: ["Hobgoblin Devastator"]
NoteIcon: monster
BestiaryType: humanoid (goblinoid)
SourceType: Bestiary
BookSource: Volo's Guide to Monsters p. 161
---
# [Hobgoblin Devastator](3-Mechanics\CLI\bestiary\humanoid/hobgoblin-devastator-vgm.md)
*Source: Volo's Guide to Monsters p. 161*  

In hobgoblin society, the Academy of Devastation identifies hobgoblins with a talent for magic and puts them through a grueling training regimen that endows them with the ability to call down fireballs and other destructive magic on the host's behalf. A hobgoblin devastator on the battlefield is simultaneously a boon to all its allies and a threat to every foe around it.

## Into the Fray

While other cultures treat their wizards as cloistered academics, hobgoblins expect their spellcasters to fight. Devastators learn the basics of weapon use, and they measure their deeds by the enemies defeated though their magic.

Devastators have the respect of other members of the host, and they receive obedience and deference from many quarters. Their ability to lay waste to entire formations with a single use of magic allows them to gain far more glory in battle than a single warrior.

Other cultures might view the use of such abilities as a short cut to glory, but to hobgoblins a gift for magic is as valued and useful as a strong sword arm or brilliance in tactics. They are all boons from Maglubiyet that must be cultivated and unleashed upon the enemy.

The Academy of Devastation believes that an academic approach to magic is a sign of weakness and inefficiency. A warrior doesn't need to know about metallurgy to wield a blade, so why should a wizard care about where magic comes from? Devastators love to prove their superiority in battle by seeking out enemy spellcasters and destroying them.

## Only Results Matter

Devastators study a simplified form of evocation magic. Their training lacks the theory and context that other folk study, making them skilled in battle but relatively illiterate on the finer points of how and why their magic works.

```statblock
"name": "Hobgoblin Devastator (VGM)"
"size": "Medium"
"type": "humanoid"
"subtype": "goblinoid"
"alignment": "Lawful Evil"
"ac": !!int "13"
"ac_class": "[studded leather](/3-Mechanics/CLI/items/studded-leather-armor.md)"
"hp": !!int "45"
"hit_dice": "7d8 + 14"
"stats":
- !!int "13"
- !!int "12"
- !!int "14"
- !!int "16"
- !!int "13"
- !!int "11"
"speed": "30 ft."
"skillsaves":
  "Arcana": !!int "5"
"senses": "darkvision 60 ft., passive Perception 11"
"languages": "Common, Goblin"
"cr": "4"
"traits":
- "desc": "The hobgoblin is a 7th-level spellcaster. Its spellcasting ability is Intelligence\
    \ (spell save DC 13, dice: d20+5 (+5) to hit with spell attacks). It has the\
    \ following wizard spells prepared:\n\nCantrips (at will): [acid splash](/3-Mechanics/CLI/spells/acid-splash.md),\
    \ [fire bolt](/3-Mechanics/CLI/spells/fire-bolt.md), [ray of frost](/3-Mechanics/CLI/spells/ray-of-frost.md),\
    \ [shocking grasp](/3-Mechanics/CLI/spells/shocking-grasp.md)\n\n1st level (4\
    \ slots): [fog cloud](/3-Mechanics/CLI/spells/fog-cloud.md), [magic missile](/3-Mechanics/CLI/spells/magic-missile.md),\
    \ [thunderwave](/3-Mechanics/CLI/spells/thunderwave.md)\n\n2nd level (3 slots):\
    \ [gust of wind](/3-Mechanics/CLI/spells/gust-of-wind.md), [Melf's acid arrow](/3-Mechanics/CLI/spells/melfs-acid-arrow.md),\
    \ [scorching ray](/3-Mechanics/CLI/spells/scorching-ray.md)\n\n3rd level (3\
    \ slots): [fireball](/3-Mechanics/CLI/spells/fireball.md), [fly](/3-Mechanics/CLI/spells/fly.md),\
    \ [lightning bolt](/3-Mechanics/CLI/spells/lightning-bolt.md)\n\n4th level (1\
    \ slots): [ice storm](/3-Mechanics/CLI/spells/ice-storm.md)"
  "name": "Spellcasting"
- "desc": "Once per turn, the hobgoblin can deal an extra dice:2d6|text(7) (2d6)\
    \ damage to a creature it hits with a damaging spell attack if that target is\
    \ within 5 feet of an ally of the hobgoblin and that ally isn't [incapacitated](/3-Mechanics/CLI/rules/conditions.md#incapacitated)."
  "name": "Arcane Advantage"
- "desc": "When the hobgoblin casts a spell that causes damage or that forces other\
    \ creatures to make a saving throw, it can choose itself and any number of allies\
    \ to be immune to the damage caused by the spell and to succeed on the required\
    \ saving throw."
  "name": "Army Arcana"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+3 (+3) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d6 + 1|text(4) (1d6 + 1) bludgeoning damage, or dice:1d8 +\
    \ 1|text(5) (1d8 + 1) bludgeoning damage if used with two hands."
  "name": "Quarterstaff"
"source":
- "VGM"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/VGM/Hobgoblin%20Devastator.webp"
```
^statblock

```encounter-table
name: Hobgoblin Devastator
creatures:
 - 1: Hobgoblin Devastator
```

## Environment

grassland, forest, hill