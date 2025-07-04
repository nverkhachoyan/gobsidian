---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mtf
- monster/cr/23
- monster/size/large
- monster/type/fiend/demon
aliases: ["Fraz-Urb'luu"]
NoteIcon: monster
BestiaryType: fiend (demon)
SourceType: Bestiary
BookSource: Mordenkainen's Tome of Foes p. 146
---
# [Fraz-Urb'luu](3-Mechanics\CLI\bestiary\npc/fraz-urbluu-mtf.md)
*Source: Mordenkainen's Tome of Foes p. 146*  

## Fraz-Urb'luu

All demons are liars, but Fraz-Urb'luu is the Prince of Deception and Demon Lord of Illusions. He uses every trick, every ounce of demonic cunning, to manipulate his enemies—mortal and fiend alike—to do his will. Fraz-Urb'luu can create dreamlands and mind-bending fantasies able to deceive the most discerning foes.

Once imprisoned for centuries below Castle Greyhawk on the world of Oerth, Fraz-Urb'luu has slowly rebuilt his power in the Abyss. He seeks the pieces of the legendary staff of power taken from him by those who imprisoned him, and commands his servants to do likewise.

The Prince of Deception's true form is like that of a great gargoyle, some 12 feet tall, with an extended, muscular neck and a smiling face framed by long, pointed ears and lank, dark hair, and bat-like wings are furled against his powerful shoulders. He can assume other forms, however, from the hideous to the beautiful. Often the demon lord becomes so immersed in playing a role that he loses himself in it for a time.

Many of the cultists of Fraz-Urb'luu aren't even aware they serve the Prince of Deception, believing their master is a beneficent being and granter of wishes, some lost god or celestial, or even another fiend. Fraz-Urb'luu wears all these masks and more. He particularly delights in aiding demon-hunters against his demonic adversaries, driving the hunters to greater and greater atrocities in the name of their cause, only to eventually reveal his true nature and claim their souls as his own.

## Fraz-Urb'luu's Lair

Fraz-Urb'luu's lair lies within the Abyssal lair known as Hollow's Heart, a featureless plain of white dust with few structures on it. The lair itself is the city of Zoragmelok, a circular fortress surrounded by adamantine walls topped with razors and hooks. Corkscrew towers loom above twisted domes and vast amphitheaters, just a few examples of the impossible architecture that fills the city.

```statblock
"name": "Fraz-Urb'luu (MTF)"
"size": "Large"
"type": "fiend"
"subtype": "demon"
"alignment": "Chaotic Evil"
"ac": !!int "18"
"ac_class": "natural armor"
"hp": !!int "337"
"hit_dice": "27d10 + 189"
"stats":
- !!int "29"
- !!int "12"
- !!int "25"
- !!int "26"
- !!int "24"
- !!int "26"
"speed": "40 ft., fly 40 ft."
"saves":
  "Dexterity": !!int "8"
  "Wisdom": !!int "14"
  "Intelligence": !!int "15"
  "Constitution": !!int "14"
"skillsaves":
  "Deception": !!int "15"
  "Stealth": !!int "8"
  "Perception": !!int "14"
"damage_resistances": "cold, fire, lightning"
"damage_immunities": "poison; bludgeoning, piercing, slashing that is nonmagical"
"condition_immunities": "[charmed](/3-Mechanics/CLI/rules/conditions.md#charmed),\
  \ [exhaustion](/3-Mechanics/CLI/rules/conditions.md#exhaustion), [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened),\
  \ [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)"
"senses": "truesight 120 ft., passive Perception 24"
"languages": "all, telepathy 120 ft."
"cr": "23"
"traits":
- "desc": "Fraz-Urb'luu's spellcasting ability is Charisma (spell save DC 23). Fraz-Urb'luu\
    \ can innately cast the following spells, requiring no material components:\n\n\
    At will: [alter self](/3-Mechanics/CLI/spells/alter-self.md) (can become Medium\
    \ when changing his appearance), [detect magic](/3-Mechanics/CLI/spells/detect-magic.md),\
    \ [dispel magic](/3-Mechanics/CLI/spells/dispel-magic.md), [phantasmal force](/3-Mechanics/CLI/spells/phantasmal-force.md)\n\
    \n1/day each: [mirage arcane](/3-Mechanics/CLI/spells/mirage-arcane.md), [modify\
    \ memory](/3-Mechanics/CLI/spells/modify-memory.md), [project image](/3-Mechanics/CLI/spells/project-image.md)\n\
    \n3/day each: [confusion](/3-Mechanics/CLI/spells/confusion.md), [dream](/3-Mechanics/CLI/spells/dream.md),\
    \ [mislead](/3-Mechanics/CLI/spells/mislead.md), [programmed illusion](/3-Mechanics/CLI/spells/programmed-illusion.md),\
    \ [seeming](/3-Mechanics/CLI/spells/seeming.md)"
  "name": "Innate Spellcasting"
- "desc": "If Fraz-Urb'luu fails a saving throw, he can choose to succeed instead."
  "name": "Legendary Resistance (3/Day)"
- "desc": "Fraz-Urb'luu has advantage on saving throws against spells and other magical\
    \ effects."
  "name": "Magic Resistance"
- "desc": "Fraz-Urb'luu's weapon attacks are magical."
  "name": "Magic Weapons"
- "desc": "Fraz-Urb'luu can't be targeted by divination magic, perceived through magical\
    \ scrying sensors, or detected by abilities that sense demons or fiends."
  "name": "Undetectable"
"actions":
- "desc": "Fraz-Urb'luu makes three attacks: one with his bite and two with his fists."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+16 (+16) to hit, reach 10 ft., one\
    \ target. Hit: dice:3d6 + 9|text(19) (3d6 + 9) piercing damage."
  "name": "Bite"
- "desc": "Melee Weapon Attack: dice: d20+16 (+16) to hit, reach 10 ft., one\
    \ target. Hit: dice:3d8 + 9|text(22) (3d8 + 9) bludgeoning damage."
  "name": "Fist"
"legendary_actions":
- "desc": "Melee Weapon Attack: dice: d20+16 (+16) to hit, reach 15 ft., one\
    \ target. Hit: dice:2d10 + 9|text(20) (2d10 + 9) bludgeoning damage. If\
    \ the target is a Large or smaller creature, it is also [grappled](/3-Mechanics/CLI/rules/conditions.md#grappled)\
    \ (escape DC 24). The [grappled](/3-Mechanics/CLI/rules/conditions.md#grappled)\
    \ target is also [restrained](/3-Mechanics/CLI/rules/conditions.md#restrained).\
    \ Fraz-Urb'luu can grapple only one creature with his tail at a time."
  "name": "Tail"
- "desc": "Fraz-Urb'luu casts [phantasmal killer](/3-Mechanics/CLI/spells/phantasmal-killer.md),\
    \ no [concentration](/3-Mechanics/CLI/rules/conditions.md#concentration) required."
  "name": "Phantasmal Killer (Costs 2 Actions)"
"lair_actions":
- "desc": "On Initiative count 20 (losing initiative ties), Fraz-Urb'luu can take\
    \ a lair action to cause one of the following effects; he can't use the same effect\
    \ two rounds in a row:"
  "name": ""
- "desc": "- Fraz-Urb'luu causes up to 5 doors within the lair to become walls, and\
    \ an equal number of doors to appear on walls where there previously were none.\
    \  \n- Fraz-Urb'luu chooses one humanoid within the lair and instantly creates\
    \ a simulacrum of that creature (as if created with the [simulacrum](/3-Mechanics/CLI/spells/simulacrum.md)\
    \ spell). This simulacrum obeys Fraz-Urb'luu's commands and is destroyed on the\
    \ next initiative count 20.  \n- Fraz-Urb'luu creates a [wave](/3-Mechanics/CLI/items/wave.md)\
    \ of anguish. Each creature he can see within the lair must succeed on a DC 23\
    \ Wisdom saving throw or take dice:6d10|text(33) (6d10) psychic damage.  "
  "name": ""
"regional_effects":
- "desc": "The region containing Fraz-Urb'luu's lair is warped by his magic, creating\
    \ one or more of the following effects:"
  "name": ""
- "desc": "- Intelligent creatures within 1 mile of the lair frequently see hallucinations\
    \ of long-dead friends and comrades that vanish after only a brief glimpse.  \n\
    - Roads and paths within 6 miles of the lair twist and turn back on themselves,\
    \ making navigation in the area exceedingly difficult.  \n- If a humanoid spends\
    \ at least 1 hour within 1 mile of the lair, that creature must succeed on a DC\
    \ 23 Wisdom saving throw or descend into a madness determined by the Madness of\
    \ Fraz-Urb'luu table. A creature that succeeds on this saving throw can't be affected\
    \ by this regional effect again for 24 hours.  "
  "name": ""
- "desc": "If Fraz-Urb'luu dies, these effects fade over the course of dice: 1d10|avg|noform\
    \ (1d10) days."
  "name": ""
- "desc": "If a creature goes mad in Fraz-Urb'luu's lair or within line of sight of\
    \ the demon lord, roll on the Madness of Fraz-Urb'luu table to determine the nature\
    \ of the madness, which is a character flaw that lasts until cured. See the \"\
    Dungeon Master's Guide\" for more on madness.\n\nMadness of Fraz-Urb'luu\n\
    \ndice: [](fraz-urbluu-mtf.md#^madness-of-fraz-urbluu)\n\n| dice: d100 | Flaw\
    \ (lasts until cured) |\n|------------|--------------------------|\n| 01–20 |\
    \ \"I never let anyone know the truth about my actions or intentions, even if\
    \ doing so would be beneficial to me.\" |\n| 21–40 | \"I have intermittent hallucinations\
    \ and fits of catatonia.\" |\n| 41–60 | \"My mind wanders as I have elaborate\
    \ fantasies that have no bearing on reality. When I return my focus to the world,\
    \ I have a hard time remembering that it was just a daydream.\" |\n| 61–80 | \"\
    I convince myself that things are true, even in the face of overwhelming evidence\
    \ to the contrary.\" |\n| 81–00 | \"My perception of reality doesn't match anyone\
    \ else's. It makes me prone to violent delusions that make no sense to anyone\
    \ else.\" |\n^madness-of-fraz-urbluu"
  "name": "Madness of Fraz-Urb'luu"
"source":
- "MTF"
- "OotA"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MTF/Fraz-Urb%27luu.webp"
```
^statblock

```encounter-table
name: Fraz-Urb'luu
creatures:
 - 1: Fraz-Urb'luu
```