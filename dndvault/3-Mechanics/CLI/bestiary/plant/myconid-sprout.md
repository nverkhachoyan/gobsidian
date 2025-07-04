---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mm
- monster/cr/0
- monster/environment/underdark
- monster/size/small
- monster/type/plant
aliases: ["Myconid Sprout"]
NoteIcon: monster
BestiaryType: plant
SourceType: Bestiary
BookSource: Monster Manual p. 230
---
# [Myconid Sprout](3-Mechanics\CLI\bestiary\plant/myconid-sprout.md)
*Source: Monster Manual p. 230*  

Myconids are intelligent, ambulatory fungi that live in the Underdark, seek enlightenment, and deplore violence. If approached peacefully, myconids gladly provide shelter or allow safe passage through their colonies.

## Circles and Melds

The largest myconid in a colony is its sovereign, which presides over one or more social groups called circles. A circle consists of twenty or more myconids that work, live, and meld together.

A meld is a form of communal meditation that allows myconids to transcend their dull subterranean existence. The myconids' rapport spores bind the participants into a group consciousness. Hallucination spores then induce a shared dream that provides entertainment and social interaction. Myconids consider melding to be the purpose of their existence. They use it in the pursuit of higher consciousness, collective union, and spiritual apotheosis. Myconids also use their rapport spores to communicate telepathically with other sentient creatures.

## Myconid Reproduction

Like other fungi, myconids reproduce by mundane sporing. They carefully control their spores' release to avoid overpopulation.

```statblock
"name": "Myconid Sprout"
"size": "Small"
"type": "plant"
"alignment": "Lawful Neutral"
"ac": !!int "10"
"hp": !!int "7"
"hit_dice": "2d6"
"stats":
- !!int "8"
- !!int "10"
- !!int "10"
- !!int "8"
- !!int "11"
- !!int "5"
"speed": "10 ft."
"senses": "darkvision 120 ft., passive Perception 10"
"languages": ""
"cr": "0"
"traits":
- "desc": "When the myconid takes damage, all other myconids within 240 feet of it\
    \ can sense its pain."
  "name": "Distress Spores"
- "desc": "While in sunlight, the myconid has disadvantage on ability checks, attack\
    \ rolls, and saving throws. The myconid dies if it spends more than 1 hour in\
    \ direct sunlight."
  "name": "Sun Sickness"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+1 (+1) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d4 - 1|text(1) (1d4 - 1) bludgeoning damage plus dice:1d4|text(2)\
    \ (1d4) poison damage."
  "name": "Fist"
- "desc": "A 10-foot radius of spores extends from the myconid. These spores can go\
    \ around corners and affect only creatures with an Intelligence of 2 or higher\
    \ that aren't undead, constructs, or elementals. Affected creatures can communicate\
    \ telepathically with one another while they are within 30 feet of each other.\
    \ The effect lasts for 1 hour."
  "name": "Rapport Spores (3/Day)"
"source":
- "MM"
- "WDMM"
- "DoSI"
- "BMT"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MM/Myconid%20Sprout.webp"
```
^statblock

```encounter-table
name: Myconid Sprout
creatures:
 - 1: Myconid Sprout
```

## Environment

underdark