---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/vgm
- monster/cr/3
- monster/environment/underdark
- monster/size/medium
- monster/type/monstrosity
aliases: ["Choldrith"]
NoteIcon: monster
BestiaryType: monstrosity
SourceType: Bestiary
BookSource: Volo's Guide to Monsters p. 132
---
# [Choldrith](3-Mechanics\CLI\bestiary\monstrosity/choldrith-vgm.md)
*Source: Volo's Guide to Monsters p. 132*  

Chitines are multiarmed humanoids with arachnid qualities that serve Lolth. They operate in well-organized colonies that prove to be effective fighters in the war against the enemies of the Demon Queen of Spiders. On occasion, Lolth pits chitines against dark elves-even though both groups worship her-as a way of punishing the drow, who created the chitines but displeased their goddess by doing so.

## Unnatural Origin

Long ago, the drow first subjected elf prisoners to horrible rituals that transformed the captives into creatures with both humanoid and spider traits, which their creators dubbed chitines. The dark elves' intention was to create slaves dedicated first of all to the drow and, by association with them, to Lolth. As the drow ultimately discovered, the goddess found this arrangement unacceptable.

The creation process required cooperation between magical disciplines. Drow wizards and warlocks used arcane magic and demonic powers, and drow priestesses invoked Lolth's aid for the divine spark needed to ensure the subject's survival. Lolth watched, expecting at some part of the process to see these new abominations dedicated to her, but no such ritual was performed. In retribution for this lack of respect, the Spider Queen twisted the drow's creation rituals to serve her own purposes.

At first, the drow were unaware that the new creatures were signs of Lolth's displeasure with them. Instead, they were pleased, because choldriths could lay eggs that birthed more chitines (and the rare choldrith) and could direct the chitines in their work. But the dark elves came to realize their mistake-choldriths belonged to Lolth, body and soul. They whispered to the chitines of their adoration of the Spider Queen and their enmity of the drow, and the seeds of a rebellion took root and grew. The chitines and choldriths rose up against their would-be masters; soon afterward most of the creatures were free, and a number of the drow who helped breed and tend them were dead.

Nowadays, drow still create chitines when they have need to. Outside the presence of a choldrith, chitines make good workers for the drow, and they can be useful if the drow find an independent chitine colony and want to infiltrate it. If the creation process yields a choldrith, though, the drow destroy the creature immediately.

As servants of Lolth, choldriths and chitines love spiders and spider-like creatures. They rear spiders and similar arachnids, such as cave fishers. Chitine colonies erect shrines to Lolth that serve as beacons, attracting spiders and other evil, brutish beings that serve her. Anywhere chitines set up a colony quickly becomes a web-shrouded, gloomy, and treacherous place.

A colony can support numerous choldriths, which serve as commanders, priests, and supervisors. The choldriths continually jockey for position, although they rarely confront one another in a way that puts the colony at risk. The colony is ruled by a singular sovereign that determines which colony members perform which tasks, including whether she or any other choldrith is permitted to lay eggs. If this supreme ruler receives a vision from Lolth, she might change her colony's entire course of action. At such times, chitines have emerged from the Underdark to settle in remote, gloomy places on the surface, from where they can wage war on other species, especially drow and elves.

## Lolth's Revenge

As the drow continued to perform the rituals, the process usually transformed the subject into the spindly, stunted creature they expected. Occasionally, though, the elf changed into a monstrosity that was more spider than elf, resembling Lolth in her spider form, and more cunning than a chitine, that the drow dubbed a choldrith.

## Lolth's Chosen

Choldriths are born with a fanatical devotion to Lolth, which leads them to develop some skill in divine magic. They preach that chitines are Lolth's favored people, and that choldriths are the Spider Queen's rightful worldly representatives sent to free the chitines from slavery. Although choldriths and chitines lack sexual characteristics, and choldriths need no mate to lay eggs, these creatures choose the gender identity of their goddess. Choldriths also believe and teach that Lolth's spider form, much like that of a choldrith, is her truest shape. Any idol to Lolth in a chitine colony depicts Lolth in this way.

## Communal Spiders

Chitines and choldriths resemble spiders, but they behave more like social insects such as ants. Chitines are divided into worker and warrior castes, and choldriths occupy the top levels of a colony's hierarchy. Each chitine has a social position that comes with duties related to that rank, and all chitines are expected to willingly sacrifice themselves to protect the choldriths. Every chitine has spinnerets and slowly produces webbing that is used to build floors, walls, structures, objects, and traps that benefit and protect the colony. A warrior might be responsible for crafting web armor (which is as tough as hide or leather), while a group of workers might be tasked to dig a pit trap and cover it with fragile webbing disguised with loose dirt to appear as a solid surface.

```statblock
"name": "Choldrith (VGM)"
"size": "Medium"
"type": "monstrosity"
"alignment": "Chaotic Evil"
"ac": !!int "15"
"ac_class": "[studded leather](/3-Mechanics/CLI/items/studded-leather-armor.md)"
"hp": !!int "66"
"hit_dice": "12d8 + 12"
"stats":
- !!int "12"
- !!int "16"
- !!int "12"
- !!int "11"
- !!int "14"
- !!int "10"
"speed": "30 ft., climb 30 ft."
"skillsaves":
  "Athletics": !!int "5"
  "Stealth": !!int "5"
  "Religion": !!int "2"
"senses": "darkvision 60 ft., passive Perception 12"
"languages": "Undercommon"
"cr": "3"
"traits":
- "desc": "The choldrith is a 4th-level spellcaster. Its spellcasting ability is Wisdom\
    \ (spell save DC 12, dice: d20+4 (+4) to hit with spell attacks). The choldrith\
    \ has the following cleric spells prepared:\n\nCantrips (at will): [guidance](/3-Mechanics/CLI/spells/guidance.md),\
    \ [mending](/3-Mechanics/CLI/spells/mending.md), [resistance](/3-Mechanics/CLI/spells/resistance.md),\
    \ [thaumaturgy](/3-Mechanics/CLI/spells/thaumaturgy.md)\n\n1st level (4 slots):\
    \ [bane](/3-Mechanics/CLI/spells/bane.md), [healing word](/3-Mechanics/CLI/spells/healing-word.md),\
    \ [sanctuary](/3-Mechanics/CLI/spells/sanctuary.md), [shield of faith](/3-Mechanics/CLI/spells/shield-of-faith.md)\n\
    \n2nd level (3 slots): [hold person](/3-Mechanics/CLI/spells/hold-person.md),\
    \ [spiritual weapon](/3-Mechanics/CLI/spells/spiritual-weapon.md) (dagger)"
  "name": "Spellcasting"
- "desc": "The choldrith has advantage on saving throws against being [charmed](/3-Mechanics/CLI/rules/conditions.md#charmed),\
    \ and magic can't put the choldrith to sleep."
  "name": "Fey Ancestry"
- "desc": "The choldrith can climb difficult surfaces, including upside down on ceilings,\
    \ without needing to make an ability check."
  "name": "Spider Climb"
- "desc": "While in sunlight, the choldrith has disadvantage on attack rolls, as well\
    \ as on Wisdom ([Perception](/3-Mechanics/CLI/rules/skills.md#Perception)) checks\
    \ that rely on sight."
  "name": "Sunlight Sensitivity"
- "desc": "While in contact with a web, the choldrith knows the exact location of\
    \ any other creature in contact with the same web."
  "name": "Web Sense"
- "desc": "The choldrith ignores movement restrictions caused by webbing."
  "name": "Web Walker"
"actions":
- "desc": "Melee or Ranged Weapon Attack: dice: d20+5 (+5) to hit, reach 5 ft.\
    \ or range 20/60 ft., one target. Hit: dice:1d4 + 3|text(5) (1d4 + 3) piercing\
    \ damage plus dice:3d6|text(10) (3d6) poison damage."
  "name": "Dagger"
- "desc": "Ranged Weapon Attack: dice: d20+5 (+5) to hit, range 30/60 ft., one\
    \ Large or smaller creature. Hit: The target is [restrained](/3-Mechanics/CLI/rules/conditions.md#restrained)\
    \ by webbing. As an action, the [restrained](/3-Mechanics/CLI/rules/conditions.md#restrained)\
    \ target can make a DC 11 Strength check, bursting the webbing on a success. The\
    \ webbing can also be attacked and destroyed (AC 10; 5 hit points; vulnerability\
    \ to fire damage; immunity to bludgeoning, poison, and psychic damage)."
  "name": "Web (Recharge 5-6)"
"source":
- "VGM"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/VGM/Choldrith.webp"
```
^statblock

```encounter-table
name: Choldrith
creatures:
 - 1: Choldrith
```

## Environment

underdark