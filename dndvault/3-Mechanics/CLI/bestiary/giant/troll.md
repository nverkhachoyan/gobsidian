---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mm
- monster/cr/5
- monster/environment/arctic
- monster/environment/forest
- monster/environment/hill
- monster/environment/mountain
- monster/environment/swamp
- monster/environment/underdark
- monster/size/large
- monster/type/giant
aliases: ["Troll"]
NoteIcon: monster
BestiaryType: giant
SourceType: Bestiary
BookSource: Monster Manual p. 291, Storm Lord's Wrath. Available in the SRD and the Basic Rules.
---
# [Troll](3-Mechanics\CLI\bestiary\giant/troll.md)
*Source: Monster Manual p. 291, Storm Lord's Wrath. Available in the SRD and the Basic Rules.*  

Born with horrific appetites, trolls eat anything they can catch and devour. They have no society to speak of, but they do serve as mercenaries to orcs, ogres, ettins, hags, and giants. As payment, trolls demand food and treasure. Trolls are difficult to control, however, doing as they please even when working with more powerful creatures.

## Regeneration

Smashing a troll's bones and slashing through its rubbery hide only makes it angry. A troll's wounds close quickly. If the monster loses an arm, a leg, or even its head, those dismembered parts can sometimes act with a life of their own. A troll can even reattach severed body parts, untroubled by its momentary disability. Only acid and fire can arrest the regenerative properties of a troll's flesh. The trolls, enraged, will attack individuals making acid and fire attacks against them above all other prey.

## Troll Freaks

Their regenerative capabilities make trolls especially susceptible to mutation. Although uncommon, such transformations can result from what the troll has done or what has been done to it. A decapitated troll might grow two heads from the stump of its neck, while a troll that eats a fey creature might gain one or more of that creature's traits.

```statblock
"name": "Troll"
"size": "Large"
"type": "giant"
"alignment": "Chaotic Evil"
"ac": !!int "15"
"ac_class": "natural armor"
"hp": !!int "84"
"hit_dice": "8d10 + 40"
"stats":
- !!int "18"
- !!int "13"
- !!int "20"
- !!int "7"
- !!int "9"
- !!int "7"
"speed": "30 ft."
"skillsaves":
  "Perception": !!int "2"
"senses": "darkvision 60 ft., passive Perception 12"
"languages": "Giant"
"cr": "5"
"traits":
- "desc": "The troll has advantage on Wisdom ([Perception](/3-Mechanics/CLI/rules/skills.md#Perception))\
    \ checks that rely on smell."
  "name": "Keen Smell"
- "desc": "The troll regains 10 hit points at the start of its turn. If the troll\
    \ takes acid or fire damage, this trait doesn't function at the start of the troll's\
    \ next turn. The troll dies only if it starts its turn with 0 hit points and doesn't\
    \ regenerate."
  "name": "Regeneration"
"actions":
- "desc": "The troll makes three attacks: one with its bite and two with its claws."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+7 (+7) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d6 + 4|text(7) (1d6 + 4) piercing damage."
  "name": "Bite"
- "desc": "Melee Weapon Attack: dice: d20+7 (+7) to hit, reach 5 ft., one target.\
    \ Hit: dice:2d6 + 4|text(11) (2d6 + 4) slashing damage."
  "name": "Claw"
"source":
- "MM"
- "HotDQ"
- "PotA"
- "RoT"
- "SKT"
- "TftYP"
- "ToA"
- "WDH"
- "WDMM"
- "GoS"
- "SLW"
- "EGW"
- "PSZ"
- "SatO"
- "DoDk"
- "QftIS"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MM/Troll.webp"
```
^statblock

```encounter-table
name: Troll
creatures:
 - 1: Troll
```

## Environment

underdark, mountain, forest, swamp, hill, arctic